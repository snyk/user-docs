# Supported workloads, container registries, languages, and operating systems

## Supported workloads

The Snyk Controller can detect the following workloads in the cluster:

* Deployment
* ReplicaSet
* ReplicationController
* DaemonSet
* StatefulSet
* Job
* CronJob
* DeploymentConfig (OpenShift)
* Pod, when these Pods have no parent or owner references.

The Controller detects these workloads by tracing the chain of owner references starting from individual Pods until it reaches the topmost workload.

For example, given a Pod, the Controller can detect that it is owned by a ReplicaSet, which in turn is owned by a Deployment, which in turn has no other owners. The workload detected is thus the Deployment.

In cases where a workload is owned by a Custom Resource, the `snyk monitor` currently cannot proceed and must assume that the current workload is the topmost parent that the Controller was able to process.

## Supported container registries

Snyk Controller supports the following container registries:

| Container Registry                                                                                                                                 | Installation Document                                                                                                                                                                                     |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <p>ACR<br>GCR<br>DigitalOcean</p><p>ECR<br>Github Container Registry<br>Gitlab Container Registry<br>Harbor<br>JFrog Artifactory<br>Docker Hub</p> | [install-the-snyk-controller-with-helm-azure-and-google-cloud-platform.md](../install-the-snyk-controller/install-the-snyk-controller-with-helm-azure-and-google-cloud-platform.md "mention")             |
| OCR                                                                                                                                                | [install-the-snyk-controller-with-openshift-4-and-operatorhub.md](../install-the-snyk-controller/install-the-snyk-controller-with-openshift-4-and-operatorhub.md "mention")                               |
| <p>ECR<br>JFrog Artifactory</p>                                                                                                                    | [install-the-snyk-controller-on-amazon-elastic-kubernetes-service-amazon-eks.md](../install-the-snyk-controller/install-the-snyk-controller-on-amazon-elastic-kubernetes-service-amazon-eks.md "mention") |

## Supported languages

The Snyk Controller supports the following languages:

* Node
* Go
* Java
* PHP
* Python

## Supported operating systems

See [Operating system distributions supported by Snyk Container](../../how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md).
