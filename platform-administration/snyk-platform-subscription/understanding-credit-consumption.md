---
description: >-
  How to track Snyk Platform Subscription credit consumption in the Credit
  breakdown dashboard: contract pacing, monthly charts, and the Billable Events
  Log
last_reviewed: 2026-10-30
last_reviewed_source: focused-edit
---

# Track credit consumption

The **Credit breakdown** dashboard reports your credit consumption, calculated from the rate card linked in your order form. For details about your rate card, visit [Rate card and your contract](snyk-platform-subscription-mechanics.md#rate-card-and-your-contract).

## Credit breakdown dashboard

In your Tenant settings, navigate to **Plan and billing** > **Credit breakdown** to view reports that track your consumption. The dashboard has three parts: pacing against your contract, monthly consumption charts, and the Billable Events Log.

{% hint style="info" %}
All consumption data on the **Credit breakdown** page has a 24-hour lag while Snyk processes the latest data. The timestamp beside the page title, for example **Updated Today at 1:00 AM**, shows when the data was last updated.
{% endhint %}

### Contract pacing

The dashboard shows cumulative credit usage to date against the number of days into the contract. These pacing data points show the percentage of credits you have consumed against the linear share of your credit balance for the contract to date. For example, if you have a balance of 120,000 credits for a one-year term and have consumed 30,000 credits after three months from the contract start, your credit efficiency is 100%: 30,000 / ((3 / 12) × 120,000) = 100%.

### Monthly consumption charts

The dashboard also contains charts that show monthly credit consumption by Group contribution and by capability contribution for your entire Tenant. You can filter these charts by Group or capability to see consumption in more detail.

### Billable Events Log

For the most granular breakdown of consumption data, select **Billable Events Log** to open an exportable log of the underlying billable events.

The event table shows each consumption event by day as its own row, mapped to its Organization and Group. To export the billable event logs, click **Download CSV**.
