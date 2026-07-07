# Integrate with Shortcut

You can synchronize findings with your Shortcut storyboard by connecting Snyk API & Web to Shortcut. Snyk can do this synchronization automatically or manually, finding by finding.

Synchronization is one-way. Snyk does not propagate changes made to Shortcut items to their respective findings in Snyk. Snyk also does not sync manual changes to Snyk findings to Shortcut. Snyk syncs only changes caused by a scan.

## Generate a Shortcut token

1. In Shortcut, navigate to **Your Account > Settings > API Tokens**.
2. Name and generate a new token.
3. Copy and save the token somewhere secure. When you generate a new token, the value appears only once, so write it down.

## Authenticate and configure the integration

After you have the token, use it to authenticate with Snyk.

1. Log in to your account and choose **Integrations** from the side menu.
2. Scroll to the Shortcut option.
3. Enter your token and click **Save**.
4. Next, navigate to your target's **Settings** and access the **Integrations** tab to set up the configuration.
5. Fill out the required fields in the Shortcut configuration form. It includes:
   * **Project**
   * **Story type**
   * **Priority mapping**
   * **Severity mapping**

Snyk requires this information to start synchronizing findings.

1. Select the **Automatically sync all findings** box. Otherwise, Snyk does not start synchronizing them.

If you do not want to synchronize all the findings, or if you prefer to hand-pick some of them, you can manually configure the synchronization for certain findings instead of selecting this box. To do that:

1. Navigate to that target and choose a finding you want to synchronize with your Shortcut project and board.
2. Select the **Sync finding** box to get updates from Snyk to Shortcut.
