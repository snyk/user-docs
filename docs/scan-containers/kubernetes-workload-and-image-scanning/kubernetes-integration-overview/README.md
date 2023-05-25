# Kubernetes integration overview

Snyk integrates with Kubernetes, enabling you to import and test your running workloads and identify vulnerabilities in their associated images and configurations that might make those workloads less secure. Once workloads are imported, Snyk continues to monitor those workloads, identifying additional security issues as new images are deployed and the workload configuration changes.

{% hint style="info" %}
**Feature availability**\
This feature is available in Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

## **How Kubernetes integration works**

1. Your administrator installs a controller on your cluster, authenticating the integration with a unique integration ID and a service account token with the needed permissions generated from the Snyk account. For more information including the permissions needed, see the [prerequisites](../../kubernetes-integration/snyk-controller-installation/prerequisites-for-snyk-controller.md).
2. Install the controller with one of these options:
   * [Install the Snyk controller with Helm (Azure and Google Cloud Platform)](../../kubernetes-integration/snyk-controller-installation/install-the-snyk-controller-with-helm.md)
   * [Install the Snyk controller with OpenShift and OperatorHub](../../kubernetes-integration/snyk-controller-installation/install-the-snyk-controller-with-openshift-4-and-operatorhub.md)
   * [Install the Snyk controller on Amazon Elastic Kubernetes Service (Amazon EKS)](../../kubernetes-integration/snyk-controller-installation/install-the-snyk-controller-on-amazon-elastic-kubernetes-service-amazon-eks.md)
3. The controller communicates with the Kubernetes API to determine which workloads (for instance the Deployment, ReplicationController, CronJob, etc.) are running on the cluster, find their associated images, and scan them directly on the cluster for vulnerabilities.
4. From Snyk, collaborators select which workloads to import, or workloads can be imported automatically using annotations. These options are as described in [Adding Kubernetes workloads for security scanning](../../kubernetes-integration/kubernetes-integration-features/adding-kubernetes-workloads-for-security-scanning.md).
5. For each workload that your collaborators import, Snyk displays the vulnerabilities found in each image as well as a summary of configuration issues identified with the workload.
6. Snyk monitors your imported workloads on an ongoing basis, reporting on new vulnerabilities as they are disclosed whenever they affect your projects.
7. Based on your configurations, if vulnerabilities are found, Snyk notifies you via email or Slack so that you can take immediate action.

{% hint style="warning" %}
In order to maintain the health of the database, any information that relates to a workload that has not been **changed or updated for eight (8) days** will be removed. This can lead to failure on **retesting** the workload.

If an image and its corresponding project is removed, and a reimport of the same workload is initiated during the eight (8) days when the metadata still resides in the database, the project may be created again.
{% endhint %}

## **Terms and conditions**

_The Snyk Container Kubernetes integration uses Red Hat UBI (Universal Base Image)._

_Before downloading or using this application, you must agree to the Red Hat subscription agreement located at redhat.com/licenses. If you do not agree with these terms, do not download or use the application. If you have an existing Red Hat Enterprise Agreement (or other negotiated agreement with Red Hat) with terms that govern subscription services associated with Containers, then your existing agreement will control._
