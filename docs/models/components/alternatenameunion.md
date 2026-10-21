# AlternateNameUnion


## Supported Types

### 

```go
alternateNameUnion := components.CreateAlternateNameUnionStr(string{/* values here */})
```

### 

```go
alternateNameUnion := components.CreateAlternateNameUnionNumber(float64{/* values here */})
```

### 

```go
alternateNameUnion := components.CreateAlternateNameUnionBoolean(bool{/* values here */})
```

### AlternateName

```go
alternateNameUnion := components.CreateAlternateNameUnionAlternateName(components.AlternateName{/* values here */})
```

### 

```go
alternateNameUnion := components.CreateAlternateNameUnionArrayOfAny([]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch alternateNameUnion.Type {
	case components.AlternateNameUnionTypeStr:
		// alternateNameUnion.Str is populated
	case components.AlternateNameUnionTypeNumber:
		// alternateNameUnion.Number is populated
	case components.AlternateNameUnionTypeBoolean:
		// alternateNameUnion.Boolean is populated
	case components.AlternateNameUnionTypeAlternateName:
		// alternateNameUnion.AlternateName is populated
	case components.AlternateNameUnionTypeArrayOfAny:
		// alternateNameUnion.ArrayOfAny is populated
	default:
		// Unknown type - use alternateNameUnion.GetUnknownRaw() for raw JSON
}
```
