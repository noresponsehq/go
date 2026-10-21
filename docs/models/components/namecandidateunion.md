# NameCandidateUnion


## Supported Types

### 

```go
nameCandidateUnion := components.CreateNameCandidateUnionStr(string{/* values here */})
```

### 

```go
nameCandidateUnion := components.CreateNameCandidateUnionNumber(float64{/* values here */})
```

### 

```go
nameCandidateUnion := components.CreateNameCandidateUnionBoolean(bool{/* values here */})
```

### NameCandidate

```go
nameCandidateUnion := components.CreateNameCandidateUnionNameCandidate(components.NameCandidate{/* values here */})
```

### 

```go
nameCandidateUnion := components.CreateNameCandidateUnionArrayOfAny([]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch nameCandidateUnion.Type {
	case components.NameCandidateUnionTypeStr:
		// nameCandidateUnion.Str is populated
	case components.NameCandidateUnionTypeNumber:
		// nameCandidateUnion.Number is populated
	case components.NameCandidateUnionTypeBoolean:
		// nameCandidateUnion.Boolean is populated
	case components.NameCandidateUnionTypeNameCandidate:
		// nameCandidateUnion.NameCandidate is populated
	case components.NameCandidateUnionTypeArrayOfAny:
		// nameCandidateUnion.ArrayOfAny is populated
	default:
		// Unknown type - use nameCandidateUnion.GetUnknownRaw() for raw JSON
}
```
