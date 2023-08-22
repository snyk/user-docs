# Supported workloads, Container Registry, languages, OS for Kubernetes integration

## Supported workloads

The Snyk controller can detect the following workloads in the cluster:

* Deployment
* ReplicaSet
* ReplicationController
* DaemonSet
* StatefulSet
* Job
* CronJob
* DeploymentConfig (OpenShift)
* Pod, when these Pods have no parent or owner references

The controller detects these workloads by tracing the chain of owner references starting from individual Pods until it reaches the topmost workload. For example, given a Pod, the controller may detect that it is owned by a ReplicaSet, which in turn is owned by a Deployment, which in turn has no other owners - the workload detected is therefore the Deployment.

In cases where the a workload is owned by a Custom Resource, the snyk monitor currently cannot proceed and must assume that the current workload is the topmost parent that the controller was able to process.

## Supported container registries

Snyk Controller supports the following container registries.

| Container Registry                                                                                                                                 | Installation Document                                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <p>ACR<br>GCR<br>DigitalOcean</p><p>ECR<br>Github Container Registry<br>Gitlab Container Registry<br>Harbor<br>JFrog Artifactory<br>Docker Hub</p> | [install-the-snyk-controller-with-helm.md](../snyk-controller-installation/install-the-snyk-controller-with-helm.md "mention")                                                                             |
| OCR                                                                                                                                                | [install-the-snyk-controller-with-openshift-4-and-operatorhub.md](../snyk-controller-installation/install-the-snyk-controller-with-openshift-4-and-operatorhub.md "mention")                               |
| <p>ECR<br>JFrog Artifactory</p>                                                                                                                    | [install-the-snyk-controller-on-amazon-elastic-kubernetes-service-amazon-eks.md](../snyk-controller-installation/install-the-snyk-controller-on-amazon-elastic-kubernetes-service-amazon-eks.md "mention") |

## Supported languages

Snyk Controller supports the following languages.

* Node
* Go
* Java
* PHP
* Python

## Supported operating systems

You can find an up-to-date list of supported operating systems in [Snyk Container: Supported operating system distributions](../../../../scan-containers/how-snyk-container-works/supported-operating-system-distributions.md).
