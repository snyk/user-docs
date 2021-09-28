# Interpret scan results

The [`Projects`](https://solutions.snyk.io/snyk-academy/open-source/import-scm-project) page will contain an inventory of all projects added and a high level summary of findings. You can expand on a particular project to view details about vulnerabilities found as well as guidance on how to fix those. For our examples, we will want to configure **three** integrations:

1. Source code management with [GitHub](https://support.snyk.io/hc/en-us/articles/360004032117-GitHub-integration)
2. Container registry with [Amazon Elastic Container Registry \(ECR\)](https://support.snyk.io/hc/en-us/articles/360003947077-Amazon-Elastic-Container-Registry-ECR-add-images-to-Snyk)
3. Cloud native applications on [Kubernetes](https://support.snyk.io/hc/en-us/articles/360003947117-Adding-Kubernetes-workloads-for-security-scanning)

#### Source Code Management

A scan of our Git repository will yield any potential vulnerabilities in our applications open source dependencies.

![](../../../../.gitbook/assets/circleci_source_scan.png)

#### Container Registry

Scanning container images in our private registry will analyze our base image and provide upgrade recommendations to reduce known vulnerabilities.

![](../../../../.gitbook/assets/circleci_ecr_scan.png)

#### Kubernetes

Enabling the Kubernetes integration will provide insights and guidance on fixing security misconfigurations in your deployments.

![](../../../../.gitbook/assets/circleci_eks_scan.png)

