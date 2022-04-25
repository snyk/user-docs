# How Snyk Controller handles your data

Once Snyk Controller is installed in your Kubernetes cluster. Snyk Controller will pull the images in the container registry; this is the Snyk pull mechanism.

1. Snyk Controller has read-only access to the Kubernetes workloads and container registries.&#x20;
2. Snyk Controller pulls images to disk, scans them, and then deletes the image once scanning is done.
3. Snyk Controller collects only the minimum relevant information to do vulnerability analysis, which is the list of dependencies in the image.&#x20;
