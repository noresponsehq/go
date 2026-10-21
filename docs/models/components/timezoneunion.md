# TimezoneUnion


## Supported Types

### 

```go
timezoneUnion := components.CreateTimezoneUnionStr(string{/* values here */})
```

### 

```go
timezoneUnion := components.CreateTimezoneUnionNumber(float64{/* values here */})
```

### 

```go
timezoneUnion := components.CreateTimezoneUnionBoolean(bool{/* values here */})
```

### Timezone

```go
timezoneUnion := components.CreateTimezoneUnionTimezone(components.Timezone{/* values here */})
```

### 

```go
timezoneUnion := components.CreateTimezoneUnionArrayOfAny([]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch timezoneUnion.Type {
	case components.TimezoneUnionTypeStr:
		// timezoneUnion.Str is populated
	case components.TimezoneUnionTypeNumber:
		// timezoneUnion.Number is populated
	case components.TimezoneUnionTypeBoolean:
		// timezoneUnion.Boolean is populated
	case components.TimezoneUnionTypeTimezone:
		// timezoneUnion.Timezone is populated
	case components.TimezoneUnionTypeArrayOfAny:
		// timezoneUnion.ArrayOfAny is populated
	default:
		// Unknown type - use timezoneUnion.GetUnknownRaw() for raw JSON
}
```
