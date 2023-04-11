# Managing Snyk Cloud rules

{% hint style="info" %}
The name of the rule settings page differs based on the products enabled for your Organization. See [Rule settings page name](managing-snyk-cloud-rules.md#rule-settings-page-name).
{% endhint %}

If [Snyk Integrated IaC](../snyk-infrastructure-as-code/integrated-infrastructure-as-code/) is enabled for your Organization, you can view a list of all Snyk Cloud and IaC rules on the Organization's **Settings > Snyk Cloud & IaC** page.

Each rule links to detailed fix advice on the [Cloud Security Rules](https://snyk.io/security-rules/cloud) site.

A rule has a **Cloud** tag if it applies to Snyk Cloud and an **IaC** tag if it applies to Snyk IaC. Most rules apply to both products. Exceptions include Cloud-only rules that check for missing resources, such as [SNYK-CC-00168](https://snyk.io/security-rules/cloud/SNYK-CC-00168/cloudwatch-log-metric-filter-and-alarm-should-be-set-for-config-configuration-changes/), "CloudWatch log metric filter and alarm should be set for Config configuration changes."

<figure><img src="../../.gitbook/assets/snyk-cloud-and-iac-settings-page (1).png" alt=""><figcaption><p>The Snyk Cloud &#x26; IaC settings page</p></figcaption></figure>

## Set custom severity level

To set a custom severity level for a rule:

1. Navigate to **Settings > Snyk Cloud & IaC**.
2. In the **Severity settings** section, select the tab for the desired cloud provider.
3. Find the rule you want to update and select the new severity level from the drop-down menu:

<figure><img src="../../.gitbook/assets/snyk-cloud-and-iac-set-custom-severity-ui (1).png" alt=""><figcaption><p>Select the new rule severity level from the drop-down menu</p></figcaption></figure>

Changes take effect for an environment after its next scan.

To reset all custom severities, select **Reset Custom Settings**.

## Filter rules by product area

{% hint style="info" %}
This section only applies to Organizations with both Snyk Cloud and Integrated IaC enabled, or only Integrated IaC.
{% endhint %}

By default, all rules are shown. Under the **Product Area** section, you can uncheck the Cloud box to hide Cloud-only rules, or the IaC box to hide IaC-only rules.

<figure><img src="../../.gitbook/assets/snyk-cloud-iac-rules-select-by-product.png" alt=""><figcaption><p>The Product Area section allows you to filter rules by product area</p></figcaption></figure>

## Rule settings page name

The name of the rule settings page differs based on the products you have enabled:

| Enabled products              | Settings page name                                                 |
| ----------------------------- | ------------------------------------------------------------------ |
| Snyk Cloud only               | Snyk Cloud                                                         |
| Snyk Cloud and Integrated IaC | Snyk Cloud & IaC                                                   |
| Snyk Cloud and Current IaC    | <p>Separate pages:<br>- Infrastructure as code<br>- Snyk Cloud</p> |
| Integrated IaC only           | Snyk Cloud & IaC                                                   |
| Current IaC only              | Infrastructure as code                                             |
