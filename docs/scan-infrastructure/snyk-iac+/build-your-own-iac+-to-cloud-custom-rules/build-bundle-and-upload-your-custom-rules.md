# Build, bundle, and upload your Custom Rules

Use the following CLI command:

```
snyk iac rules push
```

This command builds, bundles, and uploads your Custom Rules to the Snyk App for use with your IaC files across all of the Snyk workflows in your Organization.

This command also creates or updates the push key pair (custom\_rule\_id and organization\_id) in the `manifest.json` file in your Project. With this push key pair, your Project must be checked into your source code management system so that the custom rules can be properly updated in the future.
