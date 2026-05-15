# Coverage and coverage gap policies

You can use the Coverage filter to verify if an asset has ever been tested by the product and the Coverage gap filter to verify if the asset meets the coverage requirements set in a Set coverage control policy.

{% hint style="info" %}
The Coverage Gap filter is not the opposite of the Coverage filter.

An asset can be covered (if it was scanned a month ago) but, at the same time, can still have a coverage gap, if the requirement is a daily scan.
{% endhint %}

## Coverage

### Include all selected products

Identify assets that have been scanned by both Snyk Code and Snyk OS simultaneously. Identifying these assets does not mean that they meet the coverage requirements.

### Include at least one of the selected products

Identify assets that have been scanned by either Snyk Code or Snyk OS. Identifying these assets does not mean that they meet the coverage requirements.

## Coverage gap

### Pre-set coverage control policy

Snyk Essentials includes a default asset policy to highlight potential coverage gaps in your inventory for Snyk Open Source and Code. Use this default policy to ensure you have immediate visibility into repositories that might not be scanned, specifically by identifying repositories with relevant language tags that are not currently covered through scanning.

Identifying coverage gaps using this policy is SCM integration dependent, meaning that it does not flag coverage gaps when there are no language tags found.

You can use this policy as a template, edit, update and delete it.

### Exclude all products

Identify all assets that are not part of the policy. For this, you must create a Coverage gap filter covering all the available products.

### Filter non-compliant assets

Identify the assets that don’t meet the coverage requirements of Snyk Code or Snyk OS, or both, simultaneously.\
Identify assets that meet the coverage requirements of both Snyk Code and Snyk OS.

{% hint style="info" %}
You can use the Coverage and Coverage gap filters to monitor coverage and coverage gaps of image assets. This provides a more complete view of the risks of the assets and the capability to take action when needed.
{% endhint %}
