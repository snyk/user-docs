# Risk factor: Deployed

{% hint style="info" %}
The Deployed risk factor is available only for Snyk AppRisk users.
{% endhint %}

Any deployed code increases the risk of exploitation of your application and business.

Understanding what code is deployed and where enables you to adopt a remediation strategy that reduces your risk surface area from running code.

## Types of integration

The Deployed risk factor works with the Kubernetes Connector, the Snyk Runtime Sensor, and third-party integrations.

### Kubernetes **Connector**

Snyk AppRisk determines that a container image is deployed by looking for a match between the running images on the Kubernetes cluster and the created Snyk Container Projects.

Snyk AppRisk uses Kubernetes state information to extract Docker image identifiers that are being executed. The status of a Kubernetes container has image names that are being run by the Kubernetes runtime. A search is performed on the database of known Docker images to find matching names. When the image names are matched, Snyk can display this information in a graph. The graph shows the relationship between the image and the container. \


<figure><img src="https://lh6.googleusercontent.com/BoYMeFGbzjUmNmXbmtrklBcl9LLm9S94mwJWkrFA_5E5WIO07BsS3Zv-fbGBlXkNAx4oGnbBtzFijWTxUQbsnlzJI2QqprUJWPevpwBybhmwtzQayYnmW6_Qvhddgz1_vdy-NDZgQKUQhmxnY54xkrI" alt="A vulnerability in a deployed image"><figcaption><p>A vulnerability in a deployed image</p></figcaption></figure>

Kubernetes is very [specific](https://kubernetes.io/docs/concepts/containers/images/#image-names) about how images are managed. Snyk uses the same logic to map the images Snyk knows about. Whenever you scan an Image with Snyk Container, Snyk collects information about the image name and image ID. Snyk uses this information to map images against information from Kubernetes.

{% hint style="info" %}
Snyk adheres to the defined naming standards as documented for [Kubernetes](https://kubernetes.io/docs/concepts/containers/images/#image-names) and Docker [images](https://docs.docker.com/engine/reference/commandline/images/) to ensure consistency with Kubernetes.&#x20;
{% endhint %}

<figure><img src="../../../.gitbook/assets/Screenshot 2023-07-12 at 02.01.48.png" alt="Naming standards"><figcaption><p>Naming standards</p></figcaption></figure>

Consider the following examples:\


<table><thead><tr><th width="267.3333333333333">User-provided name in Kubernetes manifest </th><th>Name used in matching</th><th>Changed (yes/no)</th></tr></thead><tbody><tr><td>gcr.io/my-company/my-app:production</td><td>gcr.io/my-company/my-app:production</td><td>No</td></tr><tr><td>gcr.io/my-company/my-app:latest</td><td>gcr.io/my-company/my-app:latest</td><td>No</td></tr><tr><td>gcr.io/my-company/my-app</td><td>gcr.io/my-company/my-app:latest</td><td>Yes - latest tag appended</td></tr><tr><td>my-app</td><td>docker.io/my-app/my-app:latest</td><td>Yes - defaulted to Docker public registry, latest tag appended</td></tr></tbody></table>

The matching uses the following order of precedence, where the first step must pass for at least one Snyk Container Project, and the subsequent steps further validate the match.&#x20;

1. Match the image name, for example, `gcr.io/my-company/my-app:latest`.
2. Match the image digest.
3. Group Snyk Container Projects by image digest.

Consider the examples that follow.

#### **Example one: Using Snyk Container CLI**

Result: Image successfully matched, and risk factor applied

<figure><img src="../../../.gitbook/assets/Screenshot 2023-07-12 at 02.04.31.png" alt="Image matecjed"><figcaption><p>Image matched</p></figcaption></figure>

The container image is scanned using the Snyk Container CLI only, referencing the full name of the image, including the registry. Snyk recommends doing this after the image is built and before it is deployed to your cluster.

An example scan follows:

&#x20;`$ snyk container monitor gcr.io/my-company/my-app:latest`

The image is deployed to your Kubernetes cluster with this example manifest:\


`spec:`

&#x20; `containers:`

&#x20; `- name: my-app`

&#x20;   `image: gcr.io/my-company/my-app:latest`

This enables Insights to successfully match the image name and apply the Deployed risk factor to all issues associated with this Snyk Container Project.

#### **Example two: Using Snyk Container CLI and a container Registry**&#x20;

Result: Image successfully matched, and risk factor applied

<figure><img src="../../../.gitbook/assets/Screenshot 2023-07-12 at 02.05.31.png" alt="Matching of names"><figcaption><p>Matching of names</p></figcaption></figure>

The container image is scanned referencing a partial name, omitting the container registry.&#x20;

An example scan follows:

`$ snyk container monitor my-app:lates`t

Insights is not able to match this Project as the names do not match.&#x20;

The image is deployed to your Kubernetes cluster with this example manifest.

`spec:`

&#x20; `containers:`

&#x20; `- name: my-app`

&#x20;   `image: gcr.io/my-company/my-app:latest`

In addition, the same image is scanned by the container registry.&#x20;

This creates a Project with the full name of the image, including the registry, enabling a match to be made.&#x20;

This also includes the image digest.&#x20;

Insights is then able to group all Snyk Container Projects by image digest, enabling a deployment match on the CLI Project with the partial name through the container registry Project. &#x20;

This enables Insights to successfully match both Snyk Container Projects and apply the Deployed risk factor to all issues associated with those Projects.&#x20;

{% hint style="info" %}
Snyk recommends always specifying the full name of the image in your CLI command. If you are not able to do this, Snyk recommends also scanning the same image using a second integration.&#x20;
{% endhint %}

### Snyk Runtime Sensor

The Snyk Runtime Sensor applies the Deployed risk factor by constantly monitoring the active containers within a Kubernetes cluster. It does this by capturing real-time Kubernetes events and taking hourly snapshots of the cluster's state. These real-time data and snapshots help identify which images are currently being deployed. By comparing this information with Snyk Projects and known vulnerabilities, the Snyk Runtime Sensor accurately assesses deployed risks. This ensures that security insights are up-to-date and reflect the current deployment environment.

### Third-party integrations

The Deployed risk factor is extended to third-party integrations through real-time monitoring of Kubernetes clusters, along with data from external sources. This enables a comprehensive assessment of deployed risks across diverse environments. By integrating with third-party tools, Snyk can cross-reference vulnerabilities from various sources, ensuring that the deployed risk factor reflects the most accurate and current threat landscape. This holistic approach allows Snyk to provide actionable insights and maintain robust security postures across all integrated platforms.

## Technical details for Snyk Insights Deployed risk factor

The Kubernetes Connector continuously monitors the Kubernetes events. These events are streamed to the Snyk platform continuously.&#x20;

On a regular schedule, every hour, the data pipeline performs a reconciliation of the state of the cluster to create a snapshot. This snapshot is used to determine what images are being run in the cluster.

At the same interval, the data pipeline takes a snapshot of all Snyk Projects and data sources and extrapolates packages and images. This snapshot is used to determine which images and packages are known to Snyk for any given customer.&#x20;

Both snapshots are then compared, and evidence graphs are generated to determine the deployed facts at that point in time.&#x20;
