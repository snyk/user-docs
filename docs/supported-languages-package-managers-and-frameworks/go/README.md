# Go

## Applicability

Snyk for Go is supported for Snyk Open Source and Snyk Code.

Check the language availability to be imported as an application, tested, or monitored using the Snyk products.&#x20;

Available functions:

* SCM import, available for Snyk Open Source and Snyk Code.&#x20;
* Test or monitor your app through CLI and IDE, available for Snyk Open Source and Snyk Code.
* Test your app's SBOM using `pkg:golang`&#x20;
* Test your app's packages using `pkg:golang`

## Package managers and supported file extensions

Snyk for Go supports Go Modules and dep as package managers, and the package registry uses multiple sources.

Snyk for Go supports the following file formats:

* Snyk Open:
  * For Go Modules: `go.mod`
  * For dep: `gopkg.lock`
* Snyk Code: `.go`

## Frameworks and libraries

The following frameworks and libraries are supported in Snyk for Go:&#x20;

* Azure/azure-sdk-for-go/sdk/ai/azopenai - Comprehensive&#x20;
* gage-technologies/mistral-go - Comprehensive&#x20;
* Gin - Partial&#x20;
* Go Standard Library - Comprehensive&#x20;
* google/generative-ai-go/genai - Comprehensive&#x20;
* GORM library - Partial&#x20;
* grpc-go - Comprehensive
* labstack/echo - Partial&#x20;
* sashabaranov/go-openai - Comprehensive&#x20;
* spf13/pflag - Comprehensive

## Features

The following features are supported in Snyk for Go:

| Snyk Open Source                                                     | Snyk Code                                                                 |
| -------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| <ul><li>PR checks</li><li>License scanning</li><li>Reports</li></ul> | <ul><li>Reports</li><li>Custom rules</li><li>Interfile analysis</li></ul> |

{% hint style="info" %}
If the **Snyk FixPR** feature is enabled, this means that you will be notified if the PR checks fail when the following conditions are met:&#x20;

* The **PR checks** feature is enabled and configured to **Only fail when the issues found have a fix available.**
* "**Fixed in" available** is set to **Yes.**
{% endhint %}
