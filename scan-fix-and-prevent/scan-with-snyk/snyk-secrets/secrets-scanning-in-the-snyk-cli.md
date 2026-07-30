---
description: Scan codebases for exposed secrets with the Snyk CLI.
---

# Secrets scanning in the Snyk CLI

{% hint style="info" %}
To use Snyk Secrets in the CLI, consider these system parameters:

* Maximum file size: 1 MB
* Maximum repository or bundle size: 2 GB
* Maximum number of files in a repository or bundle: 300,000
* Maximum file path length: 256 characters
{% endhint %}

Scan your codebase for hard-coded secrets, including API keys and passwords. Use the `snyk secrets test` command to run a secrets scan from the command line.

{% columns %}
{% column %}
### Command reference

Review the command syntax, options, and exit codes.

[snyk secrets test command reference](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/commands/secrets-test)
{% endcolumn %}

{% column %}
### Scan guidance

Learn how to scan for secrets using the Snyk CLI.

[Secrets scanning in the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/secrets-scanning-in-the-snyk-cli)
{% endcolumn %}
{% endcolumns %}

