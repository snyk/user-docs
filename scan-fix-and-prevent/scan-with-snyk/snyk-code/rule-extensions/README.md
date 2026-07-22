---
description: How Snyk Code Rule Extensions add custom functions to existing security rules to match the unique context of your Project
nav_context: agnostic
---

# Rule Extensions

You can extend Snyk Code's security rules to fit your Project's unique context. Rule Extensions add custom functions to existing security rules, which helps the engine understand your code's specific logic and provides more accurate findings.

A primary use for this feature is to reduce false positives. For example, you can register your Project's in-house [sanitizer functions](custom-sanitizers.md) so the scanner recognizes them as valid ways to clean data.

Rule Extensions are available to Enterprise customers and can be managed through the Snyk Web UI or the public REST API. Access is granted through a custom role — see [Rule Extensions permissions](rule-extensions-permissions.md) for the permissions required for each.
