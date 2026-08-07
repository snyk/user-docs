# Secrets scanning in the SCM

{% hint style="info" %}
To use Snyk Secrets in the SCM, consider these system parameters:

* Maximum file size: 1 MB
* Maximum repository or bundle size: 70 GB
* Maximum number of files in a repository or bundle: 4.5 million
* Once Snyk Secrets is enabled, pre-existing projects are only automatically scanned for secrets when Snyk receives a push event from the SCM. Scans can also be triggered by re-importing the project.
{% endhint %}

Snyk Secrets scanning detects hard-coded secrets, credentials, and API keys directly in your source control management (SCM) repositories before they are exposed.

Snyk includes secret detection during routine pull request checks and manual SCM imports.

## View Snyk Secrets results

After you enable secret scanning and test your repositories, you can review the discovered secrets in the Snyk Web UI.

1. Navigate to **Projects**.
2. Select the SCM target (for example, your GitHub repository) that you want to review.
3. Click the **Secrets** Project nested under that repository.
4. Under the **Issue Type** filter, select **Hardcoded Secrets** to isolate secret-related vulnerabilities.
5. Click a specific secret finding to expand the issue card and reveal the following details:
   * **Issue Description**: The type of secret detected (for example, AWS Access Key or GitHub Token).
   * **Data flow**: The exact file names and line numbers where the secret is exposed.
   * **Remediation Advice**: Best practices on how to revoke the exposed secret and securely inject it using environment variables or a secrets manager.

{% hint style="warning" %}
Revoke any valid secrets immediately after discovery. Snyk highlights where the secret is exposed, but cannot automatically revoke external credentials.
{% endhint %}

## Exclude files and directories from SCM secrets scans

To exclude paths from secrets scanning, commit a `.snyk` file to the root of your repository. Snyk applies the exclusions before it retrieves your files, so Snyk never retrieves or scans the excluded files, and they produce no findings. Use exclusions when a repository contains a large number of known, non-production secrets that you do not want to review individually.

{% hint style="info" %}
Excluding a path is not the same as ignoring an issue. Snyk does not scan excluded paths at all, so they produce no findings and do not appear in the Snyk Web UI as ignored issues. To suppress a specific finding but continue to scan the file, use ignores instead.
{% endhint %}

Snyk Secrets applies the patterns from the following `exclude` sections of the `.snyk` file:

* `global`: Applies to Snyk Secrets and to the other Snyk products that support the `global` section.
* `secrets`: Applies to Snyk Secrets only.

The following example excludes the `vendor` directory from all the supported Snyk products, and excludes the PEM test fixtures and the `examples` directory from Snyk Secrets only:

```yaml
# Snyk (https://snyk.io) policy file
exclude:
  global:
    - vendor/**
  secrets:
    - "fixtures/**/*.pem"
    - examples/**
```

{% hint style="info" %}
* For SCM secrets scans, Snyk reads the `.snyk` file at the root of the repository only. Unlike Snyk Code, Snyk Secrets does not apply the `.snyk` files that are located in subdirectories.
* Snyk Secrets applies the `global` and `secrets` sections only. It does not apply the `code` and `iac-drift` sections, which apply to Snyk Code and Snyk IaC.
* Patterns follow the `.gitignore` pattern syntax, are matched against the file paths relative to the root of the repository, and are case-sensitive.
* List the patterns as plain strings. For SCM secrets scans, Snyk does not support the `expires` and `reason` fields on an exclusion pattern; if a pattern includes them, Snyk cannot parse the file and the scan continues without the exclusions.
* Wrap any pattern that begins with a special character, such as an asterisk (`*`), in double quotation marks.
* If the `.snyk` file is missing, empty, or cannot be parsed, the scan continues without the exclusions.
* Snyk supports a maximum of 1,000 exclusion patterns for each scan, including the patterns that Snyk applies by default.
* Snyk Secrets also skips file types that cannot contain readable secrets, such as binaries, archives, media files, fonts, and dependency lockfiles. Snyk skips these regardless of your `.snyk` file.

For the full exclusion pattern syntax and formatting rules, see [Exclusion syntax of the `.snyk` file](https://app.gitbook.com/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/import-project-repository/exclude-directories-and-files-from-project-import#exclusion-syntax-of-the-.snyk-file).
{% endhint %}

After you add or change the `.snyk` file, Snyk applies the new exclusions on the next scan of the repository. To apply them immediately, push the change to your repository or re-import the Project.

To exclude paths when you scan locally with the Snyk CLI, see [Exclude files and directories from a scan](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-secrets/secrets-scanning-in-the-snyk-cli.md#exclude-files-and-directories-from-a-scan).
