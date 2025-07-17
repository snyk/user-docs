# Coverage and coverage gap policies

You can use the Coverage filter to verify if an asset has ever been tested by the product and the Coverage gap filter to verify if the asset meets the coverage requirements set in a Set coverage control policy.

{% hint style="info" %}
The Coverage Gap filter is not the opposite of the Coverage filter.&#x20;

An asset can be covered (if it was scanned a month ago) but, at the same time, can still have a coverage gap, if the requirement is a daily scan.
{% endhint %}

The following video presents an overview of the coverage and coverage gap filters from the Snyk Web UI.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1737656987/snyk-learn/product-training-videos/Snyk_Essentials_and_Snyk_AppRisk_-4b_-_v1_-_Coverage_vs_Coverage_Gap.mp4" %}
Coverage and coverage gap when filtering or creating asset policies
{% endembed %}

## Coverage

### Include all selected products

Identify assets that have been scanned by both Snyk Code and Snyk OS simultaneously. Identifying these assets does not mean that they meet the coverage requirements.

<figure><img src="https://lh7-us.googleusercontent.com/1aKKSl4O03NT8YL3qR0K1vpcfEMtlCw9pLYrKJ3Q2OdtVYTqdMbsbtWr7Jq32TzMBKEo1t53c7gaEndbiFVqLObxPcUcw7vmmaaSHO5K7UsgtjVu6FO3kLCp6cT_-CX1CzX5Anst0acYqVom89K9y14" alt="Set Coverage filters"><figcaption><p>Set coverage filters</p></figcaption></figure>

### Include at least one of the selected products

Identify assets that have been scanned by either Snyk Code or Snyk OS. Identifying these assets does not mean that they meet the coverage requirements.

<figure><img src="https://lh7-us.googleusercontent.com/V9uzAQdi6GRne6GXxQ5cQLYXrMD6BD-HMcDIX5ebRk6OWpgxgkU7JSWf49CsNwciu2WZtCoKY7Eg4gk_7mQOXtsGRRns-Z0z96L4aDQQzT_CD17RVEVr57TJK-mMgYiCZW64z4EK71BjvldkWF8iLe4" alt="Set Coverage filters to include at least one of the selected products"><figcaption><p>Set coverage filters to include at least one of the selected products</p></figcaption></figure>

## Coverage gap

### Pre-set coverage control policy

Snyk Essentials includes a default asset policy to highlight potential coverage gaps in your inventory for Snyk Open Source and Code. Use this default policy to ensure you have immediate visibility into repositories that might not be scanned, specifically by identifying repositories with relevant language tags that are not currently covered through scanning.

Identifying coverage gaps using this policy is SCM integration dependent, meaning that it does not flag coverage gaps when there are no language tags found.

You can use this policy as a template, edit, update and delete it.

### Exclude all products

Identify all assets that are not part of the policy. For this, you must create a Coverage gap filter covering all the available products.

<figure><img src="https://lh7-us.googleusercontent.com/RcfoCkR_1a6-L44Bf55ed7xSX8Loyr57KKyI4oX4yh0j6ce3Oj4fu0XL67v9Ij1XKTES-uwTMgqJBFicBtLwaHKilj1orTi_LU0_dEllCvUE2jhfpJimlXIfRON8-0_DF_Qe__tmFLuKmSTOJoFOxCk" alt="Set Coverage gap filters"><figcaption><p>Set Coverage gap filters</p></figcaption></figure>

### Filter non-compliant assets

Identify the assets that don’t meet the coverage requirements of Snyk Code or Snyk OS, or both, simultaneously.

<figure><img src="https://lh7-us.googleusercontent.com/guCzWVv9SP7H1h6WYSFGwHEVvW3c0DVvg26mHAdxkorPgZI3gYCIH93QN0fXNl71ZDZxucfpROkkjruxuQ_vu5QCjS7_ImROEZlBTYIh-hxZnsM3comPaQpQbsy7s_3MDuFVEiljw2G8szWddXjqPgQ" alt="Set Coverage gap filters to include at least one of the selected products"><figcaption><p>Set Coverage gap filters to include at least one of the selected products</p></figcaption></figure>

\
Identify assets that meet the coverage requirements of both Snyk Code and Snyk OS.&#x20;

<figure><img src="https://lh7-us.googleusercontent.com/-Ys7HZ5UcthgyDyPbBNG572CTM04RJ_Tcc1JTa9GrltfSVUM5gvFLrxpNRlV6ZNqRJQOw5hL0QFworAAOVbGYCMM4J-H4z9G8L3BiU3-PEU79GqxAalKB5UvdXxKUIgNEszwH0jUN_7kpos8HLSXvo8" alt="Set Coverage filters to exclude the selected products"><figcaption><p>Set coverage filters to exclude the selected products</p></figcaption></figure>

{% hint style="info" %}
You can use the Coverage and Coverage gap filters to monitor coverage and coverage gaps of image assets. This provides a more complete view of the risks of the assets and the capability to take action when needed.
{% endhint %}
