# Review scan coverage

{% include "../../../../.gitbook/includes/pilot-guide-toc.md" %}

## Check ‌the Group-level SCM integration

Ensure that the Group-level SCM integration is configured. Navigate to the Group level, open the Integrations tab, and find your SCM in the Integration Hub. Follow the instructions for your SCM.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXfVcJKkJhYj17m0dkLnJ1DQNQsyDyfPkRJ-9Gjkf84-XGQuqVuEOY2NCKWd8E_24KByHwmAZY2kUVs8jWzYX_pCpVMRbuwskqNCiAjz7e1dsDmaLzOmWTQL0je9gaf2IHvNFE8dfA?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

## Review the Inventory

The first page of the Inventory provides an overview of your most important repositories and identifies coverage gaps, showing which repos have been tested with Snyk and which have not. See the [Inventory menu](../../../../manage-assets/manage-assets.md) page for more details.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXe9jbqqPfkZH6PvNz2pBcNtNKgfphy1GayWURCQnmGxFvQG5cw4vcYcsOpPF78ztH8xjcJKkRWKty4lNlC63bN0S8qP-qcU_EOpvbnbBEaPIgadfRnnVilbYlvR8Uk7U6n6fWyhRw?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

The **All Assets** page shows a complete list of all repositories, including the number of issues for each repo, which Snyk tests have run, tags ingested from the SCM, developers who contribute to this repo, and the repository freshness.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXcmljy-ooiInsutXG8MAz5nwwTDGFlODl6YGUnbdx942g-RUiuVDMExkkAXG0cCPkcbsh6uT-eJdURlItkQSUZfxGHbYUhLlxqNMI0IFDIX2paJE45ywN6kX3zB2SMlZ_rg4cqb?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

Click on the 'Not tested' section of the first pie chart on the overview page, or use coverage filters on the 'All Assets' page to view all repositories that the selected Snyk product has not tested.&#x20;

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXdlF4lPPJwbjSsf6AuZGzPERNXWKLi6gUWCF7JO0jOiYCgNYtHsqkvzxCIkx6-Ea5Kl9pF1VItX100Eo-ZN8MquDnUNfzUC2_C6f2p29TIv_5zOLNBfPumkBg0BrSf3bmCLDk2tXQ?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

## Classifying Assets

The Class column is available for each repository. This class is meant to reflect the business criticality of the asset from A (most critical) to D (least critical). Try setting a few of your most important repos manually to Class A. This attribute can be used in reporting to help focus on issues from your company’s most important repositories.

![](<../../../../.gitbook/assets/image (432).png>)\


While you can set the class manually, you can also create a policy to automatically classify the repository. For example, Snyk ingests metadata like GitHub Topics and Custom Properties that can be used in a policy.

In the following example, the policy filters on all repositories where there is a compliance tag, which in this case comes from the GitHub Custom Property, and sets the asset class to ‘A’.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXcC7X0ZTEwCAix3-FOwBYV5pKnyCTGigFGmVLy0VNgkPcW4JZpgzdEB3YtuyJWZ6I_f1VSWXlQfIoKnGw1Em3u3Wqnf1-LVSYhL1Im189lpjZLpJi1kTFf7zw9Fn03x3BC0xZbVSg?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

The asset class is a great way to filter on the most critical repositories in your organization, and help take a risk-based approach to prioritization by accounting for business impact. The asset class can be used when reviewing scan coverage and in all available Snyk reports.

{% hint style="info" %}
Additional Resources

* [Manage assets](../../../../manage-assets/manage-assets.md)
* [Product Training: Snyk Essentials](https://learn.snyk.io/catalog/?type=product-training\&topics=Snyk+Essentials)
{% endhint %}
