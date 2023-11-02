# Manage IaC+ and cloud rules

{% hint style="info" %}
The name of the rule settings page differs based on the products enabled for your Organization.
{% endhint %}

If [Snyk IaC+](../iac+-code-to-cloud-capabilities/) is enabled for your Organization, you can view a list of all Snyk Cloud and IaC rules on the Organization **Settings > Snyk Cloud & IaC** page.

Each rule links to detailed fix advice on the [Cloud Security Rules](https://security.snyk.io/rules/cloud/) site.

A rule has a **Cloud** tag if it applies to Snyk cloud context and an **IaC** tag if it applies to Snyk IaC. Most rules apply to both. Exceptions include cloud-only rules that check for missing resources, such as [SNYK-CC-00168](https://security.snyk.io/rules/cloud/SNYK-CC-00168), "CloudWatch log metric filter and alarm should be set for Config configuration changes."

<figure><img src="../../../.gitbook/assets/snyk-cloud-and-iac-settings-page (1).png" alt="The Snyk Cloud and IaC settings page"><figcaption><p>The Snyk Cloud and IaC settings page</p></figcaption></figure>

## Set custom severity level

To set a custom severity level for a rule:

1. Navigate to **Settings > Snyk Cloud & IaC**.
2. In the **Severity settings** section, select the tab for the desired cloud provider.
3. Find the rule you want to update and select the new severity level from the drop-down menu:

<figure><img src="../../../.gitbook/assets/snyk-cloud-and-iac-set-custom-severity-ui (1).png" alt="Select the new rule severity level from the drop-down menu"><figcaption><p>Select the new rule severity level from the drop-down menu</p></figcaption></figure>

Changes take effect for an environment after its next scan.

To reset all custom severities, select **Reset Custom Settings**.

## Filter rules by cloud or IaC+ area

{% hint style="info" %}
This section applies to Organizations with both cloud and IaC+ enabled, or IaC+ only.
{% endhint %}

By default, all rules are shown. Under the **Product Area** section, you can uncheck the Cloud box to hide Cloud-only rules, or the IaC box to hide IaC-only rules.

<figure><img src="../../../.gitbook/assets/snyk-cloud-iac-rules-select-by-product.png" alt="The Product Area section allows you to filter rules by product area"><figcaption><p>The Product Area section allows you to filter rules by product area</p></figcaption></figure>
