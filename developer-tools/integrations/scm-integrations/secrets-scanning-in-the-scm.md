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
