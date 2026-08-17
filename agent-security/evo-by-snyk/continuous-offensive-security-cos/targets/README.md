# Targets

Configure the applications you want Snyk Continuous Offensive Security to test. A target is the persistent record of one application—its URL, scope, and credentials. You create a target once and scan it as often as you like.

Use the pages in this section to add a target and manage existing ones.

## What is a target?

A target is the persistent record of one application you want to test. It holds everything the AI Pentesting engine needs to assess that application, so you configure it once and scan it as often as you like.

### What a target holds

| Element                   | Purpose                                                                                                                                          |
| ------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| **Name**                  | A display name so you can identify the target in lists and reports                                                                               |
| **Main URL**              | The primary entry point of the application. Always in scope                                                                                      |
| **Scope**                 | The additional hosts and paths agents may reach, plus any paths they must never touch. See [Define the target scope](define-the-target-scope.md) |
| **Users and credentials** | One or more users, ideally with different roles, that agents authenticate as. See [Configure authentication](configure-authentication.md)        |
| **Extra instructions**    | Custom headers, cookies, and free-text context about the application. See [Before you begin](before-you-begin.md)                                |
| **Linked repository**     | An optional source repository. When present, agents use source code context and remediation guidance becomes file and line specific              |

### Targets and scans

A target is a configuration. A scan is one execution of that configuration.

```
Target (configured once)
  ├── Scan 1  → findings, report
  ├── Scan 2  → findings, report
  └── Scan 3  → findings, report
```

This separation is what makes findings trackable over time. Because each scan runs against the same target definition, Snyk Continuous Offensive Security can tell you whether a finding is new, still open, or no longer reproducible. That is what drives the first found and last found dates on every finding, and the fix rate on your target risk overview view.

Every scan a target has ever run stays available on the target, so you can revisit a historical result or regenerate its report without rescanning.

### APIs are tested through the target, not as their own target

A target is an application, and the APIs behind it are tested as part of that application.

This is why the scope matters as much as it does. Agents observe the API calls the application makes and test those endpoints directly, but only where the API hostname is in scope. An API on a separate hostname that you did not add to the scope is rejected at the proxy, so it is never tested, even though the front end that calls it was scanned thoroughly.

Add every hostname your application depends on to the scope, and the API surface is covered. See [Define the target scope](define-the-target-scope.md).

An endpoint backed by an LLM is detected automatically and triggers red teaming coverage for AI-native threats such as prompt injection and tool abuse. This does not require a separate target either.

### How many targets should you create?

Create one target per application, per environment. Staging and production are different targets even when they run the same code, because their scope, credentials, and risk tolerance differ.

Do not split one application across several targets just because it spans multiple hostnames. Use the scope to include the additional hosts instead, so agents can follow the flows that cross between them.
