---
description: How to configure automatic asset archiving in Snyk API and Web
nav_context: classic
---

# Configure automatic asset archiving

To keep your asset list current, Snyk API & Web automatically archives assets that are no longer detected in discovery scans.

This guide explains how to configure the automatic archiving settings for your account.

## Prerequisites

To access this feature, your Snyk plan must include the Asset Discovery entitlement.

## Configure automatic archiving

1. In your Snyk API & Web account, navigate to **Settings** > **Scan Settings** and locate the **Archive discovered assets** module.
2. Use the toggle to turn the automatic archiving feature on or off. By default, this feature is on.
3. In the input field, enter the number of days (from one to 365) an asset can be undetected before Snyk automatically archives it. The default is 30 days.
4. Click **Save** to apply your changes.

## Verify the outcome

After you configure the feature, Snyk automatically moves any asset to an **Archived** state if it is not detected in any discovery scan for the number of days you specified.

You can view your archived assets at any time:

1. Navigate to the **Discovery** page.
2. Select the **State** dropdown filter.
3. Select **Archived** to see a list of all assets in that state.

## Manage archived assets

You do not need to manually reactivate assets. If a discovery scan detects a previously archived asset again, Snyk automatically returns it to the **Active** state.
