# Kubernetes integration architecture diagram

![](<../../../.gitbook/assets/System Diagram-Kubernetes integration.jpg>)

1. Customer's Snyk Org is enabled for the Kubernetes integration.
2. Customer can install Snyk Controller into their Kubernetes cluster.
3. Snyk Controller will read image information and pull images from container registries.
4. Snyk Controller will scan the images.
5. Snyk Controller will send scan results to Snyk Platform to analyze issues.
6. Customer can view the vulnerability issues on Snyk Platform.
