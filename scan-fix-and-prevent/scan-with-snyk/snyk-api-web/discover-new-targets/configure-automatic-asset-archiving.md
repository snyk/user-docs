# Configure automatic asset archiving

To keep your asset list current, Snyk API & Web automatically archives assets that are no longer detected in discovery scans.

This guide explains how to configure the automatic archiving settings for your account.

## Prerequisites

To access this feature, your Snyk API & Web plan must include the Asset Discovery entitlement.

## Configure automatic archiving

1. In your Snyk API & Web account, navigate to **Settings** > **Scan Settings** and locate the **Archive discovered assets** module.
2. Use the toggle to enable or disable the automatic archiving feature. By default, this feature is enabled.
3. In the input field, enter the number of days (from one to 365) an asset should be undetected before it is automatically archived. The default is 30 days.
4. Click **Save** to apply your changes.

## Verify the outcome

Once configured, Snyk API & Web automatically moves any asset to an **Archived** state if it is not detected in any discovery scan for the number of days you specified.

You can view your archived assets at any time:

1. Navigate to the **Discovery** page.
2. Select the **State** dropdown filter.
3. Select **Archived** to see a list of all assets in that state.

## Manage archived assets

You do not need to manually reactivate assets. If a previously archived asset is detected again in a new discovery scan, Snyk API & Web automatically returns it to the **Active** state.
