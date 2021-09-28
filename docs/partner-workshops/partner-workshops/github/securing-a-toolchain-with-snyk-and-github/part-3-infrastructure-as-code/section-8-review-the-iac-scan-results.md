# Section 8: Review the IaC Scan Results

## Step 1: Review the IaC Scan results in GitHub <a id="step-1-review-the-iac-scan-results-in-github"></a>

‌Like our Container workflow, the Snyk IaC Action uploads scan results to GitHub Security Code Scanning, allowing us to view misconfiguration risks within the GitHub UI. To view the results, head to Security -&gt; Code Scanning Alerts -&gt; Snyk Infrastructure as Code.​

![](../../../../.gitbook/assets/gh-iac-codescanningresults.png)

This gives you a fast initial glimpse into the risks present in our deployment manifests. Clicking each issue provides the line of code where it's introduced, a description of the issue, and the commit where it first appeared.​

![](../../../../.gitbook/assets/gh-iac-issuedetail.png)

## Step 2: Review IaC Scan results in the Snyk UI <a id="step-2-review-iac-scan-results-in-the-snyk-ui"></a>

We can add our deployment manifests to Snyk to view information about the issues found and help us start taking action.‌

### Import the deployment YAML into Snyk <a id="import-the-deployment-yaml-into-snyk"></a>

In the Snyk UI, add the `goof-deployment.yaml` and `goof-service.yaml` files into the Project we created before. Click the + sign on the Project entry, and provide the path to the files.​

![](../../../../.gitbook/assets/snyk-iac-addfiletoproject.png)

Once they import, the files appear in the Project list. Let's explore our service definition first, since it only has a single issue.​

![](../../../../.gitbook/assets/snyk-iac-selectservice.png)

### Review configuration issues and fix suggestions <a id="review-configuration-issues-and-fix-suggestions"></a>

The Project issues view shows the issue, its impact, and how it can be resolved. Snyk also highlights the line of code where the change can be introduced.​

![](../../../../.gitbook/assets/snyk-iac-viewissuedetails.png)

{% hint style="info" %}
Tip: Snyk IaC allows you to [set your own Severity scores](https://support.snyk.io/hc/en-us/articles/360006402818#UUID-c1919782-6bfa-b84b-a638-3913cee39fc5) for the rules it checks against.
{% endhint %}

In the next section, we'll take action to remediate these issues and unblock our Snyk Gate to we can merge these changes into our PROD branch. Proceed to Section 3 when ready!

