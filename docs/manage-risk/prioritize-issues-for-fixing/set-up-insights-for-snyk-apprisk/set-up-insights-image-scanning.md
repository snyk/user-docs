# Set up Insights: image scanning

{% hint style="info" %}
The Set up Insights option is available only with Snyk AppRisk.
{% endhint %}

To determine the risk factors and prioritize your Code, Open Source, and Container issues, you must scan your container images using [Snyk Container](../../../scan-with-snyk/snyk-container/).&#x20;

The container image is at the center of the application model that powers Snyk AppRisk. A container image includes your source code and dependencies, and it is deployed to your running environment, enabling Snyk AppRisk to use the container image to bridge the development and deployment states.\
\
Snyk AppRisk will identify any deployed container images using the [Kubernetes Connector](set-up-insights-kubernetes-connector.md) and compare the deployed container images to the list of scanned images you have scanned using [Snyk Container](../../../scan-with-snyk/snyk-container/).&#x20;

{% hint style="info" %}
Snyk recommends that you scan each image using at least one of the Snyk Container integrations.
{% endhint %}

## Snyk Container scanning with the CLI

To ensure the image names match, specify the full name of the image as referenced in your Kubernetes deployment.&#x20;

Example:

`snyk container monitor gcr.io/my-company/my-image:latest`

For details, see [Snyk CLI for container security](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/).

## Snyk Container registry scanning

The names will match if you are importing images to Snyk from the same container registry that you are referencing in your Kubernetes deployments.

For details, see [Snyk Container - Integrations](../../../scan-with-snyk/snyk-container/container-registry-integrations/).

## Snyk Container scanning with Kubernetes integration

The names of the container images will match because the deployed image is scanned by Snyk and created as a Project.

{% hint style="info" %}
To ensure you have set up your Kubernetes Connector properly, navigate to the **Set up Insights** tab on the **Issues** page and check the **Image coverage** section to view the data Insights has access to.
{% endhint %}

For details, see [Kubernetes integration](../../../scan-with-snyk/snyk-container/kubernetes-integration/).
