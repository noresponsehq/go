# Changelog

## 7.1.6

* `Sdk.Integrations.Create()`: 
  *  `request.Request.Platform.Enum(supabase)` **Added**
  *  `response.Platform.Enum(supabase)` **Added**
* `Sdk.Integrations.List()`:  `response.Items[].Platform.Enum(supabase)` **Added**
* `Sdk.Integrations.Rotate()`:  `response.Platform.Enum(supabase)` **Added**

## 7.1.5

* `Sdk.Assessments.Create()`:  `response.Account.Risk.Reasons[].Code` **Changed**
* `Sdk.Assessments.Get()`:  `response.Account.Risk.Reasons[].Code` **Changed**
* `Sdk.Assessments.Query()`:  `response.Items[].Risk.Reasons[].Code` **Changed**
* `Sdk.Accounts.ListChanged()`:  `response.Items[].Risk.Reasons[].Code` **Changed**
* `Sdk.Accounts.Get()`:  `response.Risk.Reasons[].Code` **Changed**
* `Sdk.Accounts.ListAssessments()`:  `response.Items[].Risk.Reasons[].Code` **Changed**

## 7.1.4

No changes detected.

## 7.1.3

No changes detected.

## 7.1.2

No changes detected.

## 7.1.1

* `Sdk.Accounts.RecordEnforcement()`: 
  *  `request.Body.Action.Enum(activate)` **Added**
  *  `response.Action.Enum(activate)` **Added**

## 7.1.0

* `Sdk.Keys.Create()`: **Added**
* `Sdk.Keys.List()`: **Added**
* `Sdk.Keys.Revoke()`: **Added**
* `Sdk.Settings.Get()`: **Added**
* `Sdk.Settings.Update()`: **Added**
* `Sdk.Assessments.Create()`:  `error.Code` **Changed**
* `Sdk.Assessments.Get()`:  `error.Code` **Changed**
* `Sdk.Assessments.Delete()`:  `error.Code` **Changed**
* `Sdk.Assessments.RecordOutcome()`:  `error.Code` **Changed**
* `Sdk.Assessments.Query()`:  `error.Code` **Changed**
* `Sdk.Accounts.ListChanged()`:  `error.Code` **Changed**
* `Sdk.Accounts.Get()`:  `error.Code` **Changed**
* `Sdk.Accounts.Delete()`:  `error.Code` **Changed**
* `Sdk.Accounts.ListAssessments()`:  `error.Code` **Changed**
* `Sdk.Accounts.RecordEnforcement()`:  `error.Code` **Changed**
* `Sdk.KeyUsage.Record()`:  `error.Code` **Changed**
* `Sdk.Keys.Rotate()`: 
  *  `response.Scopes[].Enum(management)` **Added**
  *  `error.Code` **Changed**
* `Sdk.Integrations.Create()`:  `error.Code` **Changed**
* `Sdk.Integrations.List()`:  `error.Code` **Changed**
* `Sdk.Integrations.Delete()`:  `error.Code` **Changed**
* `Sdk.Integrations.Rotate()`:  `error.Code` **Changed**
* `Sdk.WebhookEndpoints.List()`:  `error.Code` **Changed**
* `Sdk.WebhookEndpoints.Create()`:  `error.Code` **Changed**
* `Sdk.WebhookEndpoints.Delete()`:  `error.Code` **Changed**
* `Sdk.WebhookEndpoints.RotateSigningKey()`:  `error.Code` **Changed**
* `Sdk.WebhookEndpoints.SendTestEvent()`:  `error.Code` **Changed**
* `Sdk.WebhookEndpoints.ListDeliveries()`:  `error.Code` **Changed**
* `Sdk.WebhookEndpoints.ReplayDelivery()`:  `error.Code` **Changed**

## 7.0.5

No changes detected.

## 7.0.4

* `Sdk.Assessments.Create()`:  `response.Account.Risk.Reasons[].Code` **Changed**
* `Sdk.Assessments.Get()`:  `response.Account.Risk.Reasons[].Code` **Changed**
* `Sdk.Assessments.Query()`:  `response.Items[].Risk.Reasons[].Code` **Changed**
* `Sdk.Accounts.ListChanged()`:  `response.Items[].Risk.Reasons[].Code` **Changed**
* `Sdk.Accounts.Get()`:  `response.Risk.Reasons[].Code` **Changed**
* `Sdk.Accounts.ListAssessments()`:  `response.Items[].Risk.Reasons[].Code` **Changed**

## 7.0.3

* `Sdk.Assessments.Create()`:  `response.Detail.IncompleteSources[].Lane.Enum(domain)` **Added**
* `Sdk.Assessments.Get()`:  `response.Detail.IncompleteSources[].Lane.Enum(domain)` **Added**
* `Sdk.Assessments.Query()`:  `response.Items[].IncompleteSources[].Lane.Enum(domain)` **Added**
* `Sdk.Accounts.ListAssessments()`:  `response.Items[].IncompleteSources[].Lane.Enum(domain)` **Added**

## 7.0.2

No changes detected.

## 7.0.1

* `Sdk.Assessments.Create()`:  `response.Account.Risk.Decision.Enum(held)` **Added**
* `Sdk.Assessments.Get()`:  `response.Account.Risk.Decision.Enum(held)` **Added**
* `Sdk.Accounts.ListChanged()`:  `response.Items[].Risk.Decision.Enum(held)` **Added**
* `Sdk.Accounts.Get()`:  `response.Risk.Decision.Enum(held)` **Added**

## 7.0.0

No changes detected.

## 6.0.1

* `Sdk.Accounts.ListChanged()`: **Added**
* `Sdk.Accounts.RecordEnforcement()`: **Added**
* `Sdk.Assessments.Create()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.Assessments.Get()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.Assessments.Delete()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.Assessments.RecordOutcome()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.Assessments.Query()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.Accounts.Get()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.Accounts.Delete()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.Accounts.ListAssessments()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.KeyUsage.Record()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.Keys.Rotate()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.Integrations.Create()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.Integrations.List()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.Integrations.Delete()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.Integrations.Rotate()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.WebhookEndpoints.List()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.WebhookEndpoints.Create()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.WebhookEndpoints.Delete()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.WebhookEndpoints.RotateSigningKey()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.WebhookEndpoints.SendTestEvent()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.WebhookEndpoints.ListDeliveries()`:  `error.Code.Enum(riskVersionAhead)` **Added**
* `Sdk.WebhookEndpoints.ReplayDelivery()`:  `error.Code.Enum(riskVersionAhead)` **Added**

## 6.0.0

No changes detected.

## 5.1.1

* `Sdk.Assessments.Create()`:  `response.Account.RiskVersion` **Added**
* `Sdk.Assessments.Get()`:  `response.Account.RiskVersion` **Added**
* `Sdk.Accounts.Get()`:  `response.RiskVersion` **Added**

## 5.1.0

* `Sdk.Keys.Rotate()`:  `response.Scopes[].Enum(delete)` **Added**

## 5.0.1

No changes detected.

## 5.0.0

* `Sdk.Assessments.Create()`: 
  *  `request.Body.Flow` **Changed** (Breaking ⚠️)
  *  `response` **Changed** (Breaking ⚠️)
  *  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.Assessments.Get()`: 
  *  `response` **Changed** (Breaking ⚠️)
  *  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.Assessments.Delete()`:  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.Assessments.RecordOutcome()`:  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.Assessments.Query()`: 
  *  `response.Items[].Risk` **Changed** (Breaking ⚠️)
  *  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.Accounts.Get()`: 
  *  `response` **Changed** (Breaking ⚠️)
  *  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.Accounts.Delete()`:  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.Accounts.ListAssessments()`: 
  *  `response.Items[].Risk` **Changed** (Breaking ⚠️)
  *  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.KeyUsage.Record()`:  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.Keys.Rotate()`:  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.Integrations.Create()`:  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.Integrations.List()`:  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.Integrations.Delete()`:  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.Integrations.Rotate()`:  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.WebhookEndpoints.List()`:  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.WebhookEndpoints.Create()`:  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.WebhookEndpoints.Delete()`:  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.WebhookEndpoints.RotateSigningKey()`:  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.WebhookEndpoints.SendTestEvent()`:  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.WebhookEndpoints.ListDeliveries()`:  `error.Code` **Changed** (Breaking ⚠️)
* `Sdk.WebhookEndpoints.ReplayDelivery()`:  `error.Code` **Changed** (Breaking ⚠️)

## 4.0.1

No changes detected.

## 4.0.0

* `Sdk.WebhookEndpoints.ListDeliveries()`:  `response.Items[].EventType` **Changed** (Breaking ⚠️)
* `Sdk.WebhookEndpoints.ReplayDelivery()`:  `response.EventType` **Changed** (Breaking ⚠️)

## 3.0.8

No changes detected.

## 3.0.7

No changes detected.

## 3.0.6

No changes detected.

## 3.0.5

No changes detected.

## 3.0.4

No changes detected.

## 3.0.3

* `Sdk.KeyUsage.Record()`: **Added**
* `Sdk.Assessments.Create()`:  `response.Account.Risk.Reasons[].Code.Enum(sharedApiKey)` **Added**
* `Sdk.Assessments.Get()`:  `response.Account.Risk.Reasons[].Code.Enum(sharedApiKey)` **Added**
* `Sdk.Assessments.Query()`:  `response.Items[].Risk.Reasons[].Code.Enum(sharedApiKey)` **Added**
* `Sdk.Accounts.Get()`:  `response.Risk.Reasons[].Code.Enum(sharedApiKey)` **Added**
* `Sdk.Accounts.ListAssessments()`:  `response.Items[].Risk.Reasons[].Code.Enum(sharedApiKey)` **Added**

## 3.0.2

No changes detected.

## 3.0.1

No changes detected.

## 3.0.0

* `Sdk.Assessments.Create()`: 
  *  `request.Body` **Changed** (Breaking ⚠️)
  *  `response` **Changed** (Breaking ⚠️)
* `Sdk.Assessments.Get()`:  `response` **Changed** (Breaking ⚠️)

## 2.0.30

* `Sdk.Assessments.Delete()`: **Added**

## 2.0.29

No changes detected.

## 2.0.28

No changes detected.

## 2.0.27

No changes detected.

## 2.0.26

* `Sdk.Accounts.Delete()`: **Added**

## 2.0.25

* `Sdk.Assessments.Create()`:  `request.Body` **Changed**

## 2.0.24

No changes detected.

## 2.0.23

* `Sdk.WebhookEndpoints.ListDeliveries()`: **Added**
* `Sdk.WebhookEndpoints.ReplayDelivery()`: **Added**

## 2.0.22

No changes detected.

## 2.0.21

No changes detected.

## 2.0.20

* `Sdk.Assessments.Create()`:  `response.Risk.Reasons[].Code.Enum(testScenario)` **Added**
* `Sdk.Assessments.Get()`:  `response.Risk.Reasons[].Code.Enum(testScenario)` **Added**
* `Sdk.Assessments.Query()`:  `response.Items[].Risk.Reasons[].Code.Enum(testScenario)` **Added**
* `Sdk.Accounts.Get()`:  `response.Risk.Reasons[].Code.Enum(testScenario)` **Added**
* `Sdk.Accounts.ListAssessments()`:  `response.Items[].Risk.Reasons[].Code.Enum(testScenario)` **Added**

## 2.0.19

No changes detected.

## 2.0.18

No changes detected.

## 2.0.17

No changes detected.

## 2.0.16

* `Sdk.Accounts.Get()`: **Added**
* `Sdk.Accounts.ListAssessments()`: **Added**

## 2.0.15

No changes detected.

## 2.0.14

* `Sdk.Assessments.Create()`:  `response.AuthContext` **Added**
* `Sdk.Assessments.Get()`:  `response.AuthContext` **Added**
* `Sdk.Assessments.Query()`:  `response.Items[].AuthContext` **Added**

## 2.0.13

* `Sdk.WebhookEndpoints.List()`: **Added**
* `Sdk.WebhookEndpoints.Create()`: **Added**
* `Sdk.WebhookEndpoints.Delete()`: **Added**
* `Sdk.WebhookEndpoints.RotateSigningKey()`: **Added**
* `Sdk.WebhookEndpoints.SendTestEvent()`: **Added**
* `Sdk.Assessments.Create()`:  `error.Code` **Changed**
* `Sdk.Assessments.Get()`:  `error.Code` **Changed**
* `Sdk.Assessments.RecordOutcome()`:  `error.Code` **Changed**
* `Sdk.Assessments.Query()`:  `error.Code` **Changed**
* `Sdk.Keys.Rotate()`:  `error.Code` **Changed**
* `Sdk.Integrations.Create()`:  `error.Code` **Changed**
* `Sdk.Integrations.List()`:  `error.Code` **Changed**
* `Sdk.Integrations.Delete()`:  `error.Code` **Changed**
* `Sdk.Integrations.Rotate()`:  `error.Code` **Changed**

## 2.0.12

* `Sdk.Integrations.Create()`: **Added**
* `Sdk.Integrations.List()`: **Added**
* `Sdk.Integrations.Delete()`: **Added**
* `Sdk.Integrations.Rotate()`: **Added**
* `Sdk.Assessments.Create()`:  `error.Code.Enum(tooManyIntegrations)` **Added**
* `Sdk.Assessments.Get()`:  `error.Code.Enum(tooManyIntegrations)` **Added**
* `Sdk.Assessments.RecordOutcome()`:  `error.Code.Enum(tooManyIntegrations)` **Added**
* `Sdk.Assessments.Query()`:  `error.Code.Enum(tooManyIntegrations)` **Added**
* `Sdk.Keys.Rotate()`: 
  *  `response.Scopes[].Enum(integrations)` **Added**
  *  `error.Code.Enum(tooManyIntegrations)` **Added**

## 2.0.11

* `Sdk.Assessments.Create()`: 
  *  `request.Body.IssuedCredential` **Added**
  *  `response.Risk` **Changed**
* `Sdk.Assessments.Get()`:  `response.Risk` **Changed**
* `Sdk.Assessments.Query()`:  `response.Items[].Risk` **Changed**

## 2.0.10

No changes detected.

## 2.0.9

* `Sdk.Assessments.Create()`: 
  *  `request.Body.AuthContext` **Changed**

## 2.0.8

No changes detected.

## 2.0.7

No changes detected.

## 2.0.6

No changes detected.

## 2.0.5

No changes detected.

## 2.0.4

No changes detected.

## 2.0.3

No changes detected.

## 2.0.2

API changes since the previous release are not listed: no earlier document was recorded.
