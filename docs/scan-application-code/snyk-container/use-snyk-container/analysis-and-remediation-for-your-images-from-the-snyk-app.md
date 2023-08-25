# Analysis and fixes for your images from the Snyk Web UI

You can import container Projects into Snyk using the CLI command `snyk monitor`. Alternatively, you can import Projects directly from a supported container registry using the [Snyk Web UI](../../../getting-started/quickstart/create-a-snyk-account/logging-in-to-an-existing-account.md).

Snyk then scans your Project, testing for vulnerabilities, and imports a snapshot of your Project. Based on your configuration, daily or weekly, Snyk regularly scans the image snapshot dependencies, which in turn refers to its tag. Based on your configuration, Snyk sends you an update by email or Slack when any new vulnerabilities are identified.

If the tag for an image changes and the original tag is used for a different image, Snyk continues to scan the image associated with the original tag, meaning Snyk scans the new image on recurring tests. If you want to continue testing an image with a different tag, import the relevant tag.

On the **Projects** page, if the Project is imported from a registry integration, it is marked with the relevant registry icon; if the Project is imported from the CLI, it is marked with a CLI icon.

You can filter for all container Projects, as shown in the following example:

<figure><img src="../../../.gitbook/assets/Analysis of container.png" alt="Snyk Projects list with filters"><figcaption><p>Snyk Projects list with filters</p></figcaption></figure>

When you open any container Project, the resulting analysis and fix advice appears in the Snyk Web UI:

<figure><img src="../../../.gitbook/assets/image (315) (1).png" alt="Analysis and fix advice for a Project"><figcaption><p>Analysis and fix advice for a Project</p></figcaption></figure>

The following information is displayed:

* Project summary: general project details, including these unique details:
  * Image ID - derived from the container image digest
  * Image tag
  * Base Image
  * Total dependencies with known vulnerabilities and the total number of vulnerabilities
* Fix advice: If you included your Dockerfile for monitoring, a,navailable actionable fix advice is displayed. To view all advice, click the **Show more upgrade types** link. The advice offered depends on available fixes, as illustrated in the following image:

<figure><img src="../../../.gitbook/assets/image (115) (1) (2) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Recommendations for upgrading the base image"><figcaption><p>Recommendations for upgrading the base image</p></figcaption></figure>

* Upgrade suggestions can include the following:
  * Minor upgrades: the safest and best minor upgrade available
  * Major upgrades: an option for a major upgrade that will reduce more vulnerabilities but with greater risk
  * Alternative upgrades: viable image options for replacing your current base image with possible different base images that provide the least amount of vulnerabilities.
  * Recommendation to rebuild your base image if it is outdated
* Upgrade recommendations include these details:
  * the name of the recommended base image version
  * the number of vulnerabilities existent in the recommended upgrade
  * a summary of the vulnerability severities accordingly.
* Filters are available. When you view a container Project, you can filter by the filters available for all supported Project types, and these filters:
  * A specific binary or OS packages for binaries and packages containing issues
  * Dockerfile instructions: If you attach your Dockerfile, you can filter to view issues associated only with the base image, or view Dockerfile-related advice (user instruction), or both

<figure><img src="../../../.gitbook/assets/image (195) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2).png" alt="Filters for binaries and images"><figcaption><p>Filters for binaries and images</p></figcaption></figure>

{% hint style="info" %}
If there is only one category of issues in your container, for example, Node binary vulnerabilities or OS packages, this filter does not appear.\
If there is no Dockerfile attached for additional advice, the Dockerfile instruction filter does not appear
{% endhint %}

* The Issues tab provides a list of vulnerabilities, including origins, paths, and an overview of the vulnerabilities
* The Dependencies tab provides a tree view of the package hierarchy inside the image.
