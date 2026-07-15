# Rule Extensions

You can enhance Snyk Code by selectively extending its security rules about your project's unique context. A Rule Extension allows you to add your custom functions to an existing security rule, helping the engine understand your code's specific logic for more accurate findings.

A primary use for this feature is to reduce false positives. For example, you can register your Project's in-house [sanitizer functions](custom-sanitizers.md) so the scanner recognizes them as valid ways to clean data.

Rule Extensions is available to Enterprise customers and can be managed through the Snyk Web UI or the public REST API. Access is granted through a custom role — see [Rule Extensions permissions](rule-extensions-permissions.md) for the permissions required for each.
