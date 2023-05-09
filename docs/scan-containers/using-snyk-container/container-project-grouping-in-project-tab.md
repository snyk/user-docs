# Container Project grouping in Project tab

Snyk allows the customer to import images using the following methods:

1. Snyk Container CLI
2. Container Registry integration
3. Kubernetes integration

Different methods result in different Project groupings in the Project tab.

## CLI Monitor import

Snyk groups images and the applications found in the image. However, Snyk CLI does not use image tags for grouping, so Snyk **will not do sub-grouping** for the different image tags. Thus images from the same repo with different image tags are all grouped.

<figure><img src="../../.gitbook/assets/image (152).png" alt="Images with different image tags in one group"><figcaption><p>Images with different image tags in one group</p></figcaption></figure>

## Container Registry Integration

If the customer imports images with container registry integration, then in the Project tab, Snyk does **sub-grouping per image tag** for each image name.

<figure><img src="../../.gitbook/assets/Screenshot 2022-08-23 at 15.13.18.png" alt="Images with different image tags in sub-groups"><figcaption><p>Images with different image tags in sub-groups</p></figcaption></figure>

## Kubernetes Integration

If the source is Kubernetes integration, the top clickable item is the workload in the cluster, and the grouping is based on the image in the workload **without sub-grouping per image tag**.

<figure><img src="../../.gitbook/assets/Screenshot 2022-08-22 at 19.37.56.png" alt="Grouping of images from Kubernetes integration"><figcaption><p>Grouping of images from Kubernetes integration</p></figcaption></figure>
