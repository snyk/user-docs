# View Snyk Cloud Environments

To view all [Snyk Cloud](./) and [Snyk Integrated IaC](../snyk-infrastructure-as-code/integrated-infrastructure-as-code/) Environments in an Organization, navigate to your Organization's **Settings (cog icon) > Cloud environments**.

The cloud environments table displays the following information for each environment:

* Name
* Native ID (e.g., AWS account ID, Google project ID, CLI)
* Kind (e.g., AWS, Google, CLI)
* Date onboarded

<figure><img src="../../.gitbook/assets/snyk-cloud-environments-page.png" alt="The Snyk Cloud Environments page in the Snyk Web UI"><figcaption><p>The Snyk Cloud Environments page in the Snyk Web UI</p></figcaption></figure>

## Add an environment

To add an environment, select the **Add environment** drop-down and select the cloud provider. Follow the steps in [Snyk Cloud for AWS: Web UI](getting-started-with-snyk-cloud-aws/snyk-cloud-for-aws-web-ui/) or [Snyk Cloud for Google: Web UI](getting-started-with-snyk-cloud-google/snyk-cloud-for-google-web-ui/) to create the environment.

<figure><img src="../../.gitbook/assets/snyk-cloud-environments-page-add-env.png" alt="Add an environment in the Snyk Web UI"><figcaption><p>Add an environment in the Snyk Web UI</p></figcaption></figure>

You can also add an environment using the Snyk API:

* [Snyk Cloud for AWS: API](getting-started-with-snyk-cloud-aws/snyk-cloud-for-aws-api/)
* [Snyk Cloud for Google: API](getting-started-with-snyk-cloud-google/snyk-cloud-for-google-api/)

## Remove an environment

To remove an environment:

1. In the **Actions** column, select the `...` icon for the environment you want to remove.
2. Select **Remove**.
3. In the confirmation modal, select **Yes, remove**.

<figure><img src="../../.gitbook/assets/snyk-cloud-remove-env-ui.png" alt="Remove an environment in the Snyk Web UI"><figcaption><p>Remove an environment in the Snyk Web UI</p></figcaption></figure>

You can also remove an environment using the [Snyk API](remove-a-snyk-cloud-environment.md#api).
