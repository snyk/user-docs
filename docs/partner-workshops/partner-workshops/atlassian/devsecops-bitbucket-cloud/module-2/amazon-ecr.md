# Amazon ECR

To enable the integration between Amazon ECR and Snyk, we will take advantage of the [**Snyk: Developer-First Security on the AWS Cloud**](https://github.com/aws-quickstart/quickstart-snyk-security) **AWS Quick Start**.

![](../../../../.gitbook/assets/quickstart-snyk-security-ecr.png)

{% hint style="success" %}
Snyk's integration for Amazon Elastic Container Registry \(ECR\) is enabled between one Amazon ECR registry and one Snyk organization. To integrate with multiple registries, create a unique Snyk organization for each registry.
{% endhint %}

## Obtain your Snyk API token

From the Snyk console, navigate to **Settings** and under the **General** menu `Copy` your **Organization ID**.

![](../../../../.gitbook/assets/snyk-api-token.png)

## Enable the integration

You have the option of establishing cross-account access to enable Snyk's Amazon ECR integration as a 1-click deployment. This options is available as an official AWS Quick Start and eliminates the need for manual configuration. By clicking on the **Launch Stack** button below, you will be redirected to the AWS CloudFormation console where you will be prompted to complete the following steps:

* Create stack, click **Next**
* Specify stack details, click **Next**
* Configure stack options, click **Next**
* Scroll to bottom section under **Capabilities** and check the box and click **Create stack** 

When you are ready, [click here to deploy](https://us-west-2.console.aws.amazon.com/cloudformation/home?region=us-west-2#/stacks/create/template?stackName=Snyk-Security-ECR&templateURL=https://aws-quickstart.s3.amazonaws.com/quickstart-snyk-security/templates/snyk-ecr.yaml)!

{% hint style="warning" %}
The installation takes approximately 1 minute to complete.
{% endhint %}

### Gathering outputs

When complete, the AWS CloudFormation template will provide two necessary values in the **Outputs** tab. You will copy these values.

![](../../../../.gitbook/assets/snyk-ecr-integration-01.png)

### Accessing the integrations menu

From the Snyk app, navigate to the **Integrations** menu then click **ECR**.

![](../../../../.gitbook/assets/snyk-ecr-integration-02.png)

### Inputting values

Paste the two values perviously copied from the CloudFormation console's **Outputs** tab into the respective fields, then click **Save**.

![](../../../../.gitbook/assets/snyk-ecr-integration-03.png)

### Adding images

Once successfully connected, you will receive a confirmation message and a button to **Add your ECR images to Snyk**. Click the button.

![](../../../../.gitbook/assets/snyk-ecr-integration-04.png)

### Scan repositories

You will be able to browse all repositories associated with the AWS region selected when the integration was enabled. Select the desired repository, then click the **Add selected repositories** button.

![](../../../../.gitbook/assets/snyk-ecr-integration-05.png)

Let's proceed to the next section.

