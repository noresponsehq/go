# Changelog

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
