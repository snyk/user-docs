# Understanding credit consumption

Your rate card is the one linked in your order form. Every Enterprise order form that includes credits links a specific, versioned rate card, and that version governs your contract for its full term. Refer to it for the consumption rates and units of measure that apply to you.

Snyk publishes the current rate card at [Credit-based billing policies](https://snyk.io/policies/credit-based-billing/). This page always shows the latest version, which can differ from the version linked in your order form.

## Credit breakdown dashboard

In your Tenant settings, navigate to **Plan and billing** > **Credit breakdown** to see reporting that tracks your consumption. The dashboard has three parts: pacing against your contract, monthly consumption charts, and the Billable Events Log.

{% hint style="info" %}
All consumption data on the **Credit breakdown** page has a 24-hour lag while Snyk processes the latest data. The timestamp beside the page title, for example **Updated Today at 1:00 AM**, shows when the data was last updated.
{% endhint %}

### Contract pacing

The dashboard shows cumulative credit usage to date against the number of days into the contract. These pacing data points show the percentage of credits you have consumed against the linear share of your credit balance for the contract to date. For example, if you have a balance of 120,000 credits for a one-year term and have consumed 30,000 credits after three months from the contract start, your credit efficiency is 100%: 30,000 / ((3 / 12) × 120,000) = 100%.

### Monthly consumption charts

The dashboard also contains charts that show monthly credit consumption by Group contribution and by capability contribution for your entire Tenant. You can filter these charts by Group or capability for more consumption visibility.

### Billable Events Log

For the most granular breakdown of consumption data, select **Billable Events Log** to open an exportable log of the underlying billable events.

The event table shows each consumption event by day as its own row, mapped to its Organization and Group. To export the billable event logs, click **Download CSV**.
