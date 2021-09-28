# Snyk checks in GitHub

## Snyk test on pull request checks

Snyk will test your pull request for vulnerability and license checks in the merge process. This creates an opportunity to establish a security gate and prevent pull requests from adding new vulnerabilities or non-compliant open source libraries that don't meet the license policy of the organization to the source code baseline. The result of the Snyk test check is configurable providing flexibility in the security gate process while minimizing the impact and pace to the software development teams.

In the previous step, we completed the pull request workflow which directed us to the screenshot below. As the PR workflow completed, Snyk validated the vulnerability and license policy set for the project. Based on the policy, the checks either passed or failed and are represented in the GitHub UI. 

In the case of our sample application, all checks will pass and we can merge the pull request with the master branch using the merge pull request button.

{% hint style="info" %}
For more details on the [Snyk Test for Pull Request](https://support.snyk.io/hc/en-us/articles/360004032117-GitHub-integration#UUID-58e66c47-1931-675e-6437-c48fc9b71438_section-5dc5a2318c45e-idm44771256643696) please see the product documentation
{% endhint %}

![](../../../.gitbook/assets/screen-shot-2020-08-22-at-1.08.53-pm.png)

