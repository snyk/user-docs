# Enabling PR Checks for Snyk Code

{% hint style="info" %}
Before enabling the Automatic PR Checks feature, your selected Git repository must be integrated with Snyk; see [Set up an integration](../../../getting-started/quickstart/set-up-an-integration.md).
{% endhint %}

You can configure Snyk Code to automatically check your PRs for security vulnerabilities on the level of [an entire Organization](enabling-pr-checks-for-snyk-code.md#enabling-the-automatic-pull-request-checks-for-an-entire-organization) or [a specific Project](enabling-pr-checks-for-snyk-code.md#enabling-the-automatic-pull-request-checks-for-a-specific-project). In both levels, the configuration is done for a specific integrated SCM.

{% hint style="info" %}
One Snyk Organization can have multiple Git repository integrations. However, only the integration that has the Automatic PR Checks feature enabled in its Settings, can be used for the PR Checks by the Organization.
{% endhint %}

By default, the Project Settings inherit the Organization Settings. However, when the Settings on the Organization and Project levels are different, the Project Settings override the Organization Settings, unless you selected to apply the Organization Settings to all its Projects when saving them.

### Enabling the Automatic PR Checks for an entire Organization

When enabling the Automatic PR Checks feature for an entire Organization, all the imported repositories in this Organization will be scanned by Snyk Code once new PRs will be created in them. However, if in this Organization there are Projects that have a Custom Settings for this feature, the Custom Settings of the Projects will override the Organization Settings, unless you select to the save the Organization Settings as **Apply changes to all overridden projects** (see step 6 below).

**To enable the automatic PR Checks for an entire Organization:**

1\. On the Snyk Web UI, in your Organization, select <img src="../../../.gitbook/assets/Snyk Code - CLI - Org Settings button - Icon.png" alt="" data-size="line"> **Settings** > **Integrations**.

2\. Locate your configured and required SCM, and click the **Edit settings** option at the end of its row:

<figure><img src="../../../.gitbook/assets/image (352) (1).png" alt="Accessing the Edit settings pagefor your SCM integration"><figcaption><p>Accessing the Edit settings page for your SCM integration</p></figcaption></figure>

3\. On the **Settings** page of the selected integration, scroll down to the **Code Analysis** section:

<figure><img src="../../../.gitbook/assets/Snyk Code - PR Checks - Integrations page - PR Checks section.png" alt="Scroll to the Code Analysis section of your integration&#x27;s Setting page"><figcaption><p>Scroll to the Code Analysis section of your integration's Setting page</p></figcaption></figure>

4\. On the **Code Analysis** section, perform the following:

* Slider – change to **Enabled**.
* **Fail conditions** – leave as is. You currently can configure the Automatic PR Checks to **Only fail PR on newly introduced issues**. This means that Snyk Code scans only the changes that are included in the new PRs. If the repository has older issues that are not part of the new PR, these issues will not fail the PR merge.
* **Minimal severity to fail PR check** – select from the drop-down list the minimal severity level of the discovered issue in the PR, which will cause the PR to automatically fail. For example, if you selected **Medium**, all PR Checks that have issues with severity level of Medium and higher will fail. PR Checks with issues that have a Low severity level will be merged.

5\. To save and apply your changes, click one of the following:

* **Apply changes to all overridden projects** option – your changes are saved, and will be applied to all the Projects in the Organization. Projects that already have Custom Settings will inherit these new Organization Settings, and their Custom Settings will be overridden. After you click this option, a message appears, asking you to confirm the override action:

<figure><img src="../../../.gitbook/assets/Snyk Code - PR Checks - Integration Settings - Automatic Upgrade section - Override message.png" alt="Confirming the override action for your SCM projects"><figcaption><p>Confirming the override action for your SCM projects</p></figcaption></figure>

Click **OK** to complete the Save with Override action.

* **Save changes** button – your changes are saved, and will be applied to all the Projects in the Organization that are configured to inherit these Settings from the Organization. Projects that have Custom Settings will not be influenced by this change.

From now on, the new PRs in the imported repositories of the Organization will be automatically checked for security vulnerabilities by Snyk Code.

### Enabling the Automatic PR Checks for a specific Project

By default, the Settings on the Project level override the Settings on the Organization level. However, the Settings on the Organization level can override the Custom Settings of a Project, if they are configured after the Project level customization, and are saved with the option - **Apply changes to all overridden projects** (see Step 6 above).

**To configure the Automatic PR Checks for a specific Project:**

1\. In the Organization that includes the required Project, select the **Projects** menu option.

2\. Locate the **Code analysis** Project for which you want to enable the Automatic PR Checks. Then, click the **Settings** button on the right side of its row:

<figure><img src="../../../.gitbook/assets/enabling_auto_pr-checks_15dec2022 (1).png" alt="Locate the Code analysis project for automatic PR checks"><figcaption><p>Locate the Code analysis project for automatic PR checks</p></figcaption></figure>

3\. On the Project **Settings** page, click the relevant **integration** tab on left menu:

<figure><img src="../../../.gitbook/assets/image (3).png" alt="Click the associated integration tab in the left menu"><figcaption><p>Click the associated integration tab in the left menu</p></figcaption></figure>

4\. On the **integration** page > **Snyk Code for pull requests** section, select one of the following:

* **Inherit from Integration settings** – apply the Integration Settings of the Organization to the selected Project. Note that if the Automatic PR Checks feature is disabled for the Organization, this option will also be disabled for the Project when the Project inherits its Settings from the Organization. If you selected this option, move to Step 7 below.
* **Custom** – apply specific settings of the Automatic PR Checks feature to the Project. If you selected this option, move to the next step.

5\. \[For the **Custom** option only] If you selected the **Custom** option, perform the following:

* Slider – change to **Enabled**.
* **Fail conditions** – leave as is. You currently can configure the Automatic PR Checks to **Only fail PR on newly introduced issues**. This means that Snyk Code scans only the changes that are included in the new PRs. If the repository has older issues that are not part of the new PR, these issues will not fail the PR merge.
* **Minimal severity to fail PR check** – select from the drop-down list the minimal severity level of the discovered issue in the PR, which will cause the PR to automatically fail. For example, if you selected **Medium**, all PR Checks that have issues with severity level of Medium and higher will fail. PR Checks with issues that have a Low severity level will be merged.

6\. Click the **Update Snyk code pull request settings** button to save your settings.

From now on, the new PRs in the selected Project will be automatically checked for security vulnerabilities by Snyk Code.
