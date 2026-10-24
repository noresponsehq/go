# AssessmentUpdatedData

The assessment, or the minimal notice that it expired.


## Supported Types

### Assessment

```go
assessmentUpdatedData := components.CreateAssessmentUpdatedDataComplete(components.Assessment{/* values here */})
```

### Assessment

```go
assessmentUpdatedData := components.CreateAssessmentUpdatedDataPartial(components.Assessment{/* values here */})
```

### Assessment

```go
assessmentUpdatedData := components.CreateAssessmentUpdatedDataPending(components.Assessment{/* values here */})
```

### AssessmentExpired

```go
assessmentUpdatedData := components.CreateAssessmentUpdatedDataExpired(components.AssessmentExpired{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch assessmentUpdatedData.Type {
	case components.AssessmentUpdatedDataTypeComplete:
		// assessmentUpdatedData.Assessment is populated
	case components.AssessmentUpdatedDataTypePartial:
		// assessmentUpdatedData.Assessment is populated
	case components.AssessmentUpdatedDataTypePending:
		// assessmentUpdatedData.Assessment is populated
	case components.AssessmentUpdatedDataTypeExpired:
		// assessmentUpdatedData.AssessmentExpired is populated
}
```
