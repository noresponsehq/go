package noresponse

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"testing"
)

var firstGoBlock = regexp.MustCompile("(?s)```go\n(.*?)```")

// noImporter fails every import; createUsesNewKey only needs local names.
type noImporter struct{}

func (noImporter) Import(path string) (*types.Package, error) {
	return nil, fmt.Errorf("import %q not loaded", path)
}

// isNewKeyCall reports whether expr is noresponse.NewIdempotencyKey().
func isNewKeyCall(expr ast.Expr) bool {
	call, ok := expr.(*ast.CallExpr)
	if !ok || len(call.Args) != 0 {
		return false
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "NewIdempotencyKey" {
		return false
	}
	pkg, ok := sel.X.(*ast.Ident)
	return ok && pkg.Name == "noresponse"
}

// isCreateCall reports whether call is <x>.Assessments.Create(...).
func isCreateCall(call *ast.CallExpr) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "Create" {
		return false
	}
	recv, ok := sel.X.(*ast.SelectorExpr)
	return ok && recv.Sel.Name == "Assessments"
}

// createUsesNewKey parses a Go program and reports whether it calls
// Assessments.Create at least once and every call passes either
// noresponse.NewIdempotencyKey() or a variable whose last assignment before
// the call is noresponse.NewIdempotencyKey(), and no generated key is sent
// with two different request bodies. Comments are not code, so they cannot
// satisfy it.
func createUsesNewKey(src string) bool {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		return false
	}

	// Resolve every identifier to the variable it names, so a key declared in
	// an inner block is never mistaken for the one passed to Create. Imports
	// are not loaded: local declarations resolve without them, and go vet
	// compiles the full program separately.
	info := &types.Info{
		Defs: map[*ast.Ident]types.Object{},
		Uses: map[*ast.Ident]types.Object{},
	}
	conf := types.Config{
		Importer: noImporter{},
		Error:    func(error) {},
	}
	_, _ = conf.Check("main", fset, []*ast.File{file}, info)
	objectOf := func(ident *ast.Ident) types.Object {
		if obj := info.Defs[ident]; obj != nil {
			return obj
		}
		return info.Uses[ident]
	}

	// Every assignment of a single value to a variable, in source order.
	type assignment struct {
		pos   token.Pos
		obj   types.Object
		value ast.Expr
	}
	var assignments []assignment
	var creates []*ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.AssignStmt:
			for i, lhs := range n.Lhs {
				ident, ok := lhs.(*ast.Ident)
				if !ok {
					continue
				}
				var value ast.Expr
				if len(n.Rhs) == len(n.Lhs) {
					value = n.Rhs[i]
				}
				assignments = append(assignments, assignment{n.Pos(), objectOf(ident), value})
			}
		case *ast.ValueSpec:
			for i, ident := range n.Names {
				var value ast.Expr
				if len(n.Values) == len(n.Names) {
					value = n.Values[i]
				}
				assignments = append(assignments, assignment{n.Pos(), objectOf(ident), value})
			}
		case *ast.CallExpr:
			if isCreateCall(n) {
				creates = append(creates, n)
			}
		}
		return true
	})

	if len(creates) == 0 {
		return false
	}
	// The request body sent with each generated key. One key may be sent again
	// with the same body (a retry), never with a different body.
	bodyFor := map[ast.Expr]string{}
	for _, call := range creates {
		if len(call.Args) < 3 {
			return false
		}
		body := types.ExprString(call.Args[2])
		key := call.Args[1]
		if isNewKeyCall(key) {
			continue
		}
		ident, ok := key.(*ast.Ident)
		if !ok {
			return false
		}
		obj := info.Uses[ident]
		if obj == nil {
			return false
		}
		var last ast.Expr
		for _, a := range assignments {
			if a.obj == obj && a.pos < call.Pos() {
				last = a.value
			}
		}
		if last == nil || !isNewKeyCall(last) {
			return false
		}
		if sent, ok := bodyFor[last]; ok && sent != body {
			return false
		}
		bodyFor[last] = body
	}
	return true
}

// TestDocQuickStartsCompile vets the first Go block of README.md and USAGE.md
// as a program inside this module, so the published quick starts cannot drift
// from the client or fall back to a fixed idempotency key.
func TestDocQuickStartsCompile(t *testing.T) {
	gobin, err := exec.LookPath("go")
	if err != nil {
		t.Skip("go toolchain not on PATH")
	}
	for _, doc := range []string{"README.md", "USAGE.md"} {
		t.Run(doc, func(t *testing.T) {
			text, err := os.ReadFile(doc)
			if err != nil {
				t.Fatal(err)
			}
			match := firstGoBlock.FindSubmatch(text)
			if match == nil {
				t.Fatalf("%s has no Go block", doc)
			}
			code := string(match[1])
			if !createUsesNewKey(code) {
				t.Fatalf("%s quick start does not pass a key from noresponse.NewIdempotencyKey() to every Assessments.Create call", doc)
			}
			dir, err := os.MkdirTemp(".", "doccheck-")
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { _ = os.RemoveAll(dir) })
			if err := os.WriteFile(filepath.Join(dir, "main.go"), []byte(code), 0o600); err != nil {
				t.Fatal(err)
			}
			out, err := exec.Command(gobin, "vet", "./"+filepath.ToSlash(dir)).CombinedOutput()
			if err != nil {
				t.Fatalf("%s quick start does not compile: %v\n%s", doc, err, out)
			}
		})
	}
}

func TestCreateUsesNewKey(t *testing.T) {
	cases := []struct {
		body string
		want bool
	}{
		{"client.Assessments.Create(ctx, noresponse.NewIdempotencyKey(), req)", true},
		{"key := noresponse.NewIdempotencyKey()\nclient.Assessments.Create(ctx, key, req)", true},
		{"var key = noresponse.NewIdempotencyKey()\nclient.Assessments.Create(ctx, key, req)", true},
		{"key := noresponse.NewIdempotencyKey()\nclient.Assessments.Create(ctx, key, req)\nkey = \"fixed\"", true},
		{`client.Assessments.Create(ctx, "signup-acct_7f3a9c", req)`, false},
		{"_ = noresponse.NewIdempotencyKey()\nclient.Assessments.Create(ctx, \"signup-1\", req)", false},
		{"key := \"signup-1\"\n_ = noresponse.NewIdempotencyKey()\nclient.Assessments.Create(ctx, key, req)", false},
		{"key := noresponse.NewIdempotencyKey()\nkey = \"fixed\"\nclient.Assessments.Create(ctx, key, req)", false},
		{"key := \"fixed\"\n// key = noresponse.NewIdempotencyKey()\nclient.Assessments.Create(ctx, key, req)", false},
		{"client.Assessments.Create(ctx, noresponse.NewIdempotencyKey(), req)\nclient.Assessments.Create(ctx, \"fixed\", req)", false},
		{"client.Assessments.Create(ctx, key, req)", false},
		{"key := \"fixed\"\n{\nkey := noresponse.NewIdempotencyKey()\n_ = key\n}\nclient.Assessments.Create(ctx, key, req)", false},
		{"var key string\n{\nkey = noresponse.NewIdempotencyKey()\n}\nclient.Assessments.Create(ctx, key, req)", true},
		{"key := noresponse.NewIdempotencyKey()\nclient.Assessments.Create(ctx, key, req)\nclient.Assessments.Create(ctx, key, req)", true},
		{"key := noresponse.NewIdempotencyKey()\nclient.Assessments.Create(ctx, key, req)\nclient.Assessments.Create(ctx, key, other)", false},
		{"key := noresponse.NewIdempotencyKey()\nclient.Assessments.Create(ctx, key, req)\nkey = noresponse.NewIdempotencyKey()\nclient.Assessments.Create(ctx, key, other)", true},
		{"client.Assessments.Get(ctx, \"asm_1\", nil)", false},
	}
	for _, c := range cases {
		src := "package main\n\nfunc main() {\n" + c.body + "\n}\n"
		if got := createUsesNewKey(src); got != c.want {
			t.Errorf("createUsesNewKey(%q) = %v, want %v", c.body, got, c.want)
		}
	}
}
