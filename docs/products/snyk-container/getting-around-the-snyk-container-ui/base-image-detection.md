# Base image detection

Detecting vulnerable base images allows you to identify the source of your vulnerabilities and fix by updating the base image according to recommendation.

After you configure a [container integration](https://docs.snyk.io/snyk-container) (such as the CLI or a container registry integration), you can detect your base image.

{% hint style="info" %}
For a list of supported container registries, consult our [image scanning library.](../image-scanning-library/)
{% endhint %}

## Identification methods

To identify vulnerable base images, you can use one of these methods:

* **Auto-detection**: When Snyk analyses your container image, relevant metadata is extracted from the image manifest, and the base image is automatically detected. This method analyses **rootfs** layers from the image manifest, which can be equivalent to more than one images/ image tags in DockerHub.
* **Dockerfile**: Snyk can also detect vulnerable base images with the use of your Dockerfile. It can either be attached with a **--file** flag to your CLI **snyk container** scan, linked from an SCM through your project settings, or the Dockerfile can be detected and scanned when you import a Git repository. This method can be more accurate, but requires an additional step compared to auto-detection. For multi-stage Dockerfiles, Snyk will detect the base image included in the image introduced in the final `FROM` line. Per [Docker's multi-stage build documentation](https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds), this is because using multiple `FROM` statements lets you "selectively copy artifacts from one stage to another, leaving behind everything you don't want in the final image."

For either method, a “project” in the Snyk UI is created.

{% hint style="info" %}
Snyk currently only supports a subset of official Docker images. The list of Snyk's supported images can be found [here](https://snyk.io/docker-images/).
{% endhint %}

## Base image recommendations

If the base image is a Snyk supported image from [this](https://snyk.io/docker-images/) list, the results include recommendations for upgrades to resolve some of the discovered vulnerabilities.

This allows you to see vulnerability counts in minor and major upgrades, as well as in alternative base images which might have fewer vulnerabilities. With that, you can decide whether to upgrade your base image, and which one will be the best fit.

![](../../../.gitbook/assets/base-image2.png)

See [Analysis and fixes for your images from the Snyk Web UI](https://docs.snyk.io/products/snyk-container/getting-around-the-snyk-container-ui/analysis-and-remediation-for-your-images-from-the-snyk-app) for more details.

You can find the base image vulnerabilities in your project, among the vulnerabilities added by your instructions, sorted by their priority score. You can also filter only the base image vulnerabilities, with the **Base image** option under the **Image Layer** handy filter. See [image layer information](image-layer-information.md) for more details.
