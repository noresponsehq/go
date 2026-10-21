# WebsiteUnion


## Supported Types

### 

```go
websiteUnion := components.CreateWebsiteUnionStr(string{/* values here */})
```

### 

```go
websiteUnion := components.CreateWebsiteUnionNumber(float64{/* values here */})
```

### 

```go
websiteUnion := components.CreateWebsiteUnionBoolean(bool{/* values here */})
```

### Website

```go
websiteUnion := components.CreateWebsiteUnionWebsite(components.Website{/* values here */})
```

### 

```go
websiteUnion := components.CreateWebsiteUnionArrayOfAny([]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch websiteUnion.Type {
	case components.WebsiteUnionTypeStr:
		// websiteUnion.Str is populated
	case components.WebsiteUnionTypeNumber:
		// websiteUnion.Number is populated
	case components.WebsiteUnionTypeBoolean:
		// websiteUnion.Boolean is populated
	case components.WebsiteUnionTypeWebsite:
		// websiteUnion.Website is populated
	case components.WebsiteUnionTypeArrayOfAny:
		// websiteUnion.ArrayOfAny is populated
	default:
		// Unknown type - use websiteUnion.GetUnknownRaw() for raw JSON
}
```
