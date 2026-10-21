# SocialProfileUnion


## Supported Types

### 

```go
socialProfileUnion := components.CreateSocialProfileUnionStr(string{/* values here */})
```

### 

```go
socialProfileUnion := components.CreateSocialProfileUnionNumber(float64{/* values here */})
```

### 

```go
socialProfileUnion := components.CreateSocialProfileUnionBoolean(bool{/* values here */})
```

### SocialProfile

```go
socialProfileUnion := components.CreateSocialProfileUnionSocialProfile(components.SocialProfile{/* values here */})
```

### 

```go
socialProfileUnion := components.CreateSocialProfileUnionArrayOfAny([]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch socialProfileUnion.Type {
	case components.SocialProfileUnionTypeStr:
		// socialProfileUnion.Str is populated
	case components.SocialProfileUnionTypeNumber:
		// socialProfileUnion.Number is populated
	case components.SocialProfileUnionTypeBoolean:
		// socialProfileUnion.Boolean is populated
	case components.SocialProfileUnionTypeSocialProfile:
		// socialProfileUnion.SocialProfile is populated
	case components.SocialProfileUnionTypeArrayOfAny:
		// socialProfileUnion.ArrayOfAny is populated
	default:
		// Unknown type - use socialProfileUnion.GetUnknownRaw() for raw JSON
}
```
