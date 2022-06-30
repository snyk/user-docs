# Enabling the Automatic PR Checks

**Note**: Before enabling the Automatic PR Checks feature, your selected SCM must be integrated with Snyk.

You can configure Snyk Code to automatically check your PRs for security vulnerabilities on the level of [an entire organization](enabling-the-automatic-pr-checks.md#enabling-the-automatic-pull-request-checks-for-an-entire-organization) or [a specific Project](enabling-the-automatic-pr-checks.md#enabling-the-automatic-pull-request-checks-for-a-specific-project). In both levels, the configuration is done for a specific integrated SCM.

**Note**: One organization can be integrated with several SCMs. However, only the SCM that has the Automatic PR Checks feature enabled in its Settings, can be used for the PR Checks by the organization.

By default, the Project Settings inherit the organization Settings. However, when the Settings on the organization and Project levels are different, the Project Settings override the organization Settings, unless you selected to apply the organization Settings to all its Projects when saving them.

### Enabling the Automatic PR Checks for an entire organization

When enabling the Automatic PR Checks feature for an entire organization, all the imported repositories in this organization will be scanned by Snyk Code once new PRs will be created in them. However, if in this organization there are Projects that have a Custom Settings for this feature, the Custom Settings of the Projects will override the organization Settings, unless you select to the save the organization Settings as **Apply changes to all overridden projects** (see step 6 below).

**To enable the automatic PR Checks for an entire organization:**

1\. On the Snyk Web UI, open the required organization:

![](<../../../.gitbook/assets/OS - Automatic Dependency Upgrade - Selecting Organization (1) (1) (1) (1) (1) (3).png>)

2\. Once the required organization is open, click the **Org Settings** button \<img src="../../../.gitbook/assets/Snyk Code - PR Checks - Org Settings button - Icon.png" alt="A black and white logo

Description automatically generated with low confidence" data-size="line"> on the top menu:

![](<../../../.gitbook/assets/OS - Automatic Dependency Upgrade - Org Settings button (1) (1) (1) (1) (1) (3).png>)

3\. On the Org **Settings** page, select **Integrations** on the left menu. Then, locate your configured and required SCM, and click the **Edit settings** option at the end of its row:

![](<../../../.gitbook/assets/OS - Automatic Dependency Upgrade - Organization - Integrations page.png>)

4\. On the **Settings** page of the selected integration, scroll down to the **Code Analysis** section:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Integrations page - PR Checks section.png>)

5\. On the **Code Analysis** section, perform the following:

* Slider – change to **Enabled**.
* **Fail conditions** – leave as is. You currently can configure the Automatic PR Checks to **Only fail PR on newly introduced issues**. This means that Snyk Code scans only the changes that are included in the new PRs. If the repository has older issues that are not part of the new PR, these issues will not fail the PR merge.
* **Minimal severity to fail PR check** – select from the drop-down list the minimal severity level of the discovered issue in the PR, which will cause the PR to automatically fail. For example, if you selected **Medium**, all PR Checks that have issues with severity level of Medium and higher will fail. PR Checks with issues that have a Low severity level will be merged.

6\. To save \*\*\*\* and apply your changes, click one of the following:

* **Apply changes to all overridden projects** option – your changes are saved, and will be applied to all the Projects in the organization. Projects that already have Custom Settings will inherit these new organization Settings, and their Custom Settings will be override. After you click this option, a message appears, asking you to confirm the override action:

!\[Graphical user interface, text, application

Description automatically generated]\(<../../../.gitbook/assets/Snyk Code - PR Checks - Integration Settings - Automatic Upgrade section - Override message.png>)

Click **OK** to complete the Save with Override action.

* **Save changes** button – your changes are saved, and will be applied to all the Projects in the organization that are configured to inherit these Settings from the organization. Projects that have Custom Settings will not be influenced by this change.

From now on, the new PRs in the imported repositories of the organization will be automatically checked for security vulnerabilities by Snyk Code.

### Enabling the Automatic PR Checks for a specific Project

By default, the Settings on the Project level override the Settings on the organization level. However, the Settings on the Organization level can override the Custom Settings of a Project, if they are configured after the Project level customization, and are saved with the option - **Apply changes to all overridden projects** (see Step 6 above).

**To configure the Automatic PR Checks for a specific Project:**

1\. On the Snyk Web UI, open the organization that includes the required Project:

![](<../../../.gitbook/assets/OS - Automatic Dependency Upgrade - Selecting Organization (1) (1) (1) (1) (1) (3).png>)

2\. Move to the **Projects** page by clicking the **Projects** tab on the top menu.

3\. Locate the **Code analysis** Project for which you want to enable the Automatic PR Checks. Then, click the **Settings** button on the right side of its row:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Project - Project Settings button.png>)

4\. On the Project **Settings** page, click the relevant **integration** tab on left menu\*\*:\*\*

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Project - Settings - Integration page.png>)

5\. On the **integration** page -> **Snyk Code for pull requests** section, select one of the following:

* **Inherit from Integration settings** – apply the Integration Settings of the organization to the selected Project. Note that if the Automatic PR Checks feature is disabled for the organization, this option will also be disabled for the Project when the Project inherits its Settings from the organization. If you selected this option, move to Step 7 below.
* **Custom** – apply specific settings of the \*\*\*\* Automatic PR Checks feature to the Project. If you selected this option, move to the next step.

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Project - Settings - Integration - Inhertinace.png>)

6\. \[For the **Custom** option only] If you selected the **Custom** option, perform the following:

* Slider – change to **Enabled**.
* **Fail conditions** – leave as is. You currently can configure the Automatic PR Checks to **Only fail PR on newly introduced issues**. This means that Snyk Code scans only the changes that are included in the new PRs. If the repository has older issues that are not part of the new PR, these issues will not fail the PR merge.
* **Minimal severity to fail PR check** – select from the drop-down list the minimal severity level of the discovered issue in the PR, which will cause the PR to automatically fail. For example, if you selected **Medium**, all PR Checks that have issues with severity level of Medium and higher will fail. PR Checks with issues that have a Low severity level will be merged.

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Project - Settings - Integration - Custom.png>)

7\. Click the **Update Snyk code pull request settings** button to save your settings.

From now on, the new PRs in the selected Project will be automatically checked for security vulnerabilities by Snyk Code.
