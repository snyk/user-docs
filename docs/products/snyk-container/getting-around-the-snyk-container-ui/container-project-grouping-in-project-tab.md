# Container project grouping in Project tab

Snyk allows the customer to import images using the following methods:

1. Snyk Container CLI
2. Container Registry integration
3. Kubernetes integration

Different methods result in different project groupings in the project tab.

### CLI Monitor import

Snyk groups images and the applications found in the image. However, Snyk CLI doesn't use image tags for grouping, so Snyk **will not do sub-grouping** for the different image tags. Thus images from the same repo with different image tags are all grouped.

![](<../../../.gitbook/assets/image (220).png>)

### Container Registry Integration

If the customer imports images with container registry integration, then in the project tab Snyk does **sub-grouping per image tag** for each image name.

![](<../../../.gitbook/assets/Screenshot 2022-08-23 at 15.13.18.png>)

### Kubernetes Integration

If the source is Kubernetes integration, the top clickable item is the workload in the cluster, and the grouping is based on the image in the workload **without sub-grouping per image tag**.

![](<../../../.gitbook/assets/Screenshot 2022-08-22 at 19.37.56.png>)
