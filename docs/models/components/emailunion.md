# EmailUnion


## Supported Types

### 

```go
emailUnion := components.CreateEmailUnionStr(string{/* values here */})
```

### 

```go
emailUnion := components.CreateEmailUnionNumber(float64{/* values here */})
```

### 

```go
emailUnion := components.CreateEmailUnionBoolean(bool{/* values here */})
```

### Email

```go
emailUnion := components.CreateEmailUnionEmail(components.Email{/* values here */})
```

### 

```go
emailUnion := components.CreateEmailUnionArrayOfAny([]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch emailUnion.Type {
	case components.EmailUnionTypeStr:
		// emailUnion.Str is populated
	case components.EmailUnionTypeNumber:
		// emailUnion.Number is populated
	case components.EmailUnionTypeBoolean:
		// emailUnion.Boolean is populated
	case components.EmailUnionTypeEmail:
		// emailUnion.Email is populated
	case components.EmailUnionTypeArrayOfAny:
		// emailUnion.ArrayOfAny is populated
	default:
		// Unknown type - use emailUnion.GetUnknownRaw() for raw JSON
}
```
