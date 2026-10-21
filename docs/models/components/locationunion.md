# LocationUnion


## Supported Types

### 

```go
locationUnion := components.CreateLocationUnionStr(string{/* values here */})
```

### 

```go
locationUnion := components.CreateLocationUnionNumber(float64{/* values here */})
```

### 

```go
locationUnion := components.CreateLocationUnionBoolean(bool{/* values here */})
```

### Location

```go
locationUnion := components.CreateLocationUnionLocation(components.Location{/* values here */})
```

### 

```go
locationUnion := components.CreateLocationUnionArrayOfAny([]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch locationUnion.Type {
	case components.LocationUnionTypeStr:
		// locationUnion.Str is populated
	case components.LocationUnionTypeNumber:
		// locationUnion.Number is populated
	case components.LocationUnionTypeBoolean:
		// locationUnion.Boolean is populated
	case components.LocationUnionTypeLocation:
		// locationUnion.Location is populated
	case components.LocationUnionTypeArrayOfAny:
		// locationUnion.ArrayOfAny is populated
	default:
		// Unknown type - use locationUnion.GetUnknownRaw() for raw JSON
}
```
