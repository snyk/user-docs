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

To exclude paths from secrets scanning, commit a `.snyk` file to the root of your repository. Snyk applies the exclusions before it retrieves your files, so excluded files are never retrieved for scanning and never appear in the results. Use exclusions when a repository contains a large number of known, non-production secrets that you do not want to review individually.

{% hint style="info" %}
Excluding a path is not the same as ignoring an issue. Snyk does not scan excluded paths at all, so they produce no findings and do not appear in the Snyk Web UI as ignored issues. To suppress a specific finding while continuing to scan the file, use ignores instead.
{% endhint %}

Snyk Secrets applies the patterns in the following `exclude` sections of the `.snyk` file:

* `global`: Applies to Snyk Secrets and to other Snyk products that support the `global` section.
* `secrets`: Applies only to Snyk Secrets.

The following example excludes the `vendor` directory from all supported Snyk products, and excludes PEM test fixtures and the `examples` directory from Snyk Secrets only:

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
* For SCM secrets scans, Snyk reads only the `.snyk` file at the root of the repository. Unlike Snyk Code, Snyk Secrets does not apply `.snyk` files that are located in subdirectories.
* Snyk Secrets applies only the `global` and `secrets` sections. It ignores the `code` and `iac-drift` sections, which apply to Snyk Code and Snyk IaC.
* Wrap any pattern that begins with a special character, such as an asterisk (`*`), in double quotation marks.
* If the `.snyk` file is missing, empty, or cannot be parsed, the scan continues without the `.snyk` exclusions.
* Snyk Secrets also skips file types that cannot contain readable secrets, such as binaries, archives, media files, fonts, and dependency lockfiles. This happens regardless of your `.snyk` file.

For the full exclusion pattern syntax and formatting rules, see [Exclusion syntax of the `.snyk` file](https://app.gitbook.com/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/import-project-repository/exclude-directories-and-files-from-project-import#exclusion-syntax-of-the-.snyk-file).
{% endhint %}

After you add or change the `.snyk` file, Snyk applies the new exclusions on the next scan of the repository. To apply them immediately, push the change to your repository or re-import the Project.

To exclude paths when you scan locally with the Snyk CLI, see [Secrets scanning in the Snyk CLI](../../snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/secrets-scanning-in-the-snyk-cli.md#exclude-files-and-directories-from-a-scan).
