---
description: Learn how Snyk Secrets works in Pull Request checks.
---

# Secrets Pull Request checks

{% hint style="info" %}
Snyk enables Secrets Pull Request checks by default for Organizations that use Snyk Secrets. Configure or turn off these checks in **Integration Settings**.
{% endhint %}

Snyk Secrets adds hard-coded secret detection to Pull Request checks. It scans pull request changes for exposed credentials, tokens, and API keys before merge.

This page covers only Secrets-specific behavior. For setup, supported source code management (SCM) integrations, status checks, and branch protection, visit the main [Pull Request checks](https://docs.snyk.io/scan-fix-and-prevent/prevent/pull-request-checks) documentation.

## About Secrets Pull Request checks

{% hint style="info" %}
Manage ignore requests within the Snyk UI. To create an ignore request, click the link in the inline comment to open the finding in Snyk.
{% endhint %}

Secrets Pull Request checks detect new hard-coded secrets introduced in pull request changes.

Use them to:

* Prevent valid credentials from reaching the default branch
* Review findings in pull request context
* Open the matching finding in Snyk for more details

## Prerequisites

Secrets Pull Request checks require Snyk Secrets to be enabled for the Organization.

After you enable Snyk Secrets for an Organization, reimport your repositories to create the underlying Project. This step is required for Secrets scanning to work, including recurring tests and Pull Request checks.

## What you see in a pull request

When Secrets is enabled, and a repository has been reimported, Snyk scans that repository's pull requests for secrets. On a pull request with new findings, you see:

* A `secrets/snyk` SCM check reporting the status of the Secrets scan.
*   A Secrets row in the Snyk pull request summary comment, alongside the other Snyk checks.

    ![Snyk pull request summary comment with a Secrets result row](https://2479055233-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2FQTiGglVLHTj5smrGKzeI%2Fuploads%2F7DocXqUYRRmzFh06Mpc1%2Funknown.png?alt=media\&token=27913490-705c-4c8b-8dea-15484150b938)
* An inline comment on each identified secret, detailing the finding. The inline comment includes a link to the Snyk UI, where you can create a request to ignore the finding.

## Review and remediate findings

Snyk reports findings as **Hardcoded Secrets** and labels each finding with the specific secret type detected. Review each finding in the pull request's inline comments or in the nested Secrets Project under the repository in Snyk. Finding details include the secret type, file path, line number, and remediation guidance.

Secrets findings often require more than a code change. If a finding exposes a valid credential, revoke or rotate it first. Then remove it from the codebase and replace it with a secure secret source.

For related guidance, visit [Snyk Secrets](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-secrets), [Secrets scanning in the SCM](https://docs.snyk.io/developer-tools/integrations/scm-integrations/secrets-scanning-in-the-scm), and [Secrets scanning in the Snyk CLI](https://docs.snyk.io/developer-tools/integrations/scm-integrations/snyk-secrets/secrets-scanning-in-the-snyk-cli).
