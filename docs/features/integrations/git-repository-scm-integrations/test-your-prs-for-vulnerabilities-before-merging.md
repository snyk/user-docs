# Test your PRs for vulnerabilities before merging

Snyk integrates with your preferred Git repository to scan your manifest files for any new code and potential vulnerabilities whenever you submit a pull request (PR), protecting the security of your code before you ever merge it with the main branch.

By default, Snyk scans every pull request submitted on your monitored repositories, displaying the results and recommendations grouped together in a single security check and a single license check. The following statuses can appear on your Snyk checks:

* Success - no issues are identified and all checks pass
* Processing - this status appears until the Snyk test ends
* Failure - when issues are identified that must be fixed in order for the check to pass
* Error - an error occurs when your manifest file is out of sync, Snyk couldn't read the manifest file, or Snyk couldn't find the manifest file
* Canceled - Snyk test can't run because you've reached your monthly test limit

Administrators can manage settings for Snyk PR tests from our app on both the organization and the project levels, configuring whether the feature is on (enabled by default) and under what conditions Snyk should fail your PR checks.

### **Prerequisites**

* You must already have a Snyk account, you must be the owner or administrator of the specific organization and you must already have set up the integration for the Git repository with which you'd like to work.

{% hint style="info" %}
Settings on the project level override the settings on the organization level. Currently, we support all languages supported by the Git repositories that we integrate with: GitHub, GitLab, Bitbucket and Azure repos.
{% endhint %}

## Configure the pull request test settings for your organization

1. Log in to your account and navigate to the relevant group and organization that you want to manage.
2. Click on settings ![](../../../.gitbook/assets/cog\_icon.png) > **Integrations**.
3. Click **Edit Settings** for integration required.
4. Navigate to the **Default Snyk test for pull requests** section:
5. Choose settings from the dropdown list as follows:
   1. **Only fail when the PR is adding a dependency with issues** - only fail the license or security check on the pull request if the pull request aims to add a new dependency that contains issues
   2. **Fail if the repo has any issues** - fail the license or security check on the pull request if there are any issues at all in the repository
6. Check mark any relevant settings from the dropdown list as follows:
   1. **Only fail for high severity issues** - only fail the check for high severity issues, based on the option you chose from the dropdown list above
   2. **Only fail when the issues found have a fix available** - only fail the check for issues if there is also a fix for those issues, based on the option you chose from the dropdown list above
7. Click **Update settings** to update settings and apply them to all future projects and to all existing projects currently using your default settings. Click **Update settings & apply to existing projects** to update settings and apply them to all existing and new projects, including those projects that had been customized.
8. Settings are saved. Every time any of your collaborators submit pull requests, checks are run based on these settings, combined with the settings you've configured on the side of your Git repository.

{% hint style="info" %}
We recommend that you set Snyk status checks to be mandatory for merging pull requests from the relevant repository. See your Git repository documentation for additional help with this.
{% endhint %}

### Configure pull request test settings for a project

1. Log in to your account and navigate to the relevant group and organization that you want to manage.
2. Click on settings ![](../../../.gitbook/assets/cog\_icon.png) > **Integrations**,
3. Click **Edit Settings** for integration required.
4. Navigate to the **Default Snyk test for pull requests** section:
5. Choose:
   1. Inherit from **Integration** settings for the project to use the settings set at the Organization level
   2. **Custom** and then click the slider to enable the functionality.
6. From the options that appear, choose settings from the dropdown list as follows:
   1. **Only fail when the PR is adding a dependency with issues** - only fail the license or security check on the pull request if the pull request aims to add a new dependency that contains issues
   2. **Fail if the repo has any issues** - fail the license or security check on the pull request if there are any issues at all in the repository
7. Check mark any relevant settings from the dropdown list as follows:
   1. **Only fail for high or critical severity issues** - only fail the check for high or critical severity issues, based on the option you chose from the dropdown list above
   2. **Only fail when the issues found have a fix available** - only fail the check for issues if there is also a fix for those issues, based on the option you chose from the dropdown list above
8. Click **Update settings** to update settings and apply them to all future projects.
9. Settings are saved. Every time any of your collaborators submit pull requests for this specific manifest file, checks are run based on these settings, combined with the settings you've configured on the side of your Git repository.
