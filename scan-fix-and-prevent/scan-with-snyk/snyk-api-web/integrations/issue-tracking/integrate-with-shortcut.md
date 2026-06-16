# Integrate with Shortcut

You can synchronize findings with your Shortcut storyboard by connecting Snyk API & Web to Shortcut. This synchronization can be done automatically or manually, finding by finding.

Synchronization is one-way, and changes made to Shortcut items will not be propagated to their respective findings on Snyk API & Web. Manual changes to Snyk API & Web findings are also not synced to Shortcut. Only changes caused by a scan are synced.

## Generate a Shortcut token

1. In Shortcut, go to **Your Account > Settings > API Tokens**.
2. Name and generate a new token.
3. Copy and save the token somewhere secure. When you generate a new token, the value is displayed only once, so write it down.

## Authenticate and configure the integration

Once you have the token, use it to authenticate with Snyk API & Web.

1. Log in to your account and choose **Integrations** from the side menu.
2. Scroll to the Shortcut option.
3. Insert your token and click **Save**.
4. Next, go to your target's **Settings** and access the **Integrations** tab to set up the configuration.
5. Fill out the required fields from the Shortcut configuration form. It includes:
   * **Project**
   * **Story type**
   * **Priority mapping**
   * **Severity mapping**

This information is required for Snyk API & Web to start synchronizing findings.

1. Check the box "Automatically sync all findings". Otherwise, Snyk API & Web will not start synchronizing them.

If you do not wish to synchronize all the findings or if you prefer to hand-pick some of them, instead of selecting this checkbox you can manually configure the synchronization for certain findings. To do that:

1. Go to that target and choose a finding you want to synchronize with your Shortcut project and board.
2. Check the "Sync finding" box to get updates from Snyk API & Web to Shortcut.
