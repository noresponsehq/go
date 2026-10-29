# Actor


## Supported Types

### UserActor

```go
actor := components.CreateActorUser(components.UserActor{/* values here */})
```

### KeyActor

```go
actor := components.CreateActorAPIKey(components.KeyActor{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch actor.Type {
	case components.ActorTypeUser:
		// actor.UserActor is populated
	case components.ActorTypeAPIKey:
		// actor.KeyActor is populated
	default:
		// Unknown type - use actor.GetUnknownRaw() for raw JSON
}
```
