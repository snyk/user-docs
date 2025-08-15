# Overview of Kubernetes integration

{% hint style="info" %}
**Feature availability**\
Kubernetes integration is available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

Snyk is able to integrate with Kubernetes, enabling you to import and scan your running workloads. This helps you identify vulnerabilities in their associated images and configurations that can make those workloads less secure. After workloads are imported, Snyk continues to monitor them and to identify additional security issues as new images are deployed and the workload configuration changes.

## **How Kubernetes integration works**

Kubernetes integration follows this process:

1. Your administrator installs a controller on your cluster, authenticating the integration with a unique integration ID and a service account token with the needed permissions generated from your Snyk account. For more information, including the permissions needed, see [Prerequisites for installing the Snyk Controller](../install-the-snyk-controller/#prerequisites-for-installing-the-snyk-controller).
2. You install the controller with one of the options:
   * [Install the Snyk Controller with Helm (Azure and Google Cloud Platform)](../install-the-snyk-controller/install-the-snyk-controller-with-helm-azure-and-google-cloud-platform.md)
   * [Install the Snyk Controller on Amazon Elastic Kubernetes Service (Amazon EKS)](../install-the-snyk-controller/install-the-snyk-controller-on-amazon-elastic-kubernetes-service-amazon-eks.md)
   * [Install the Snyk Controller with OpenShift and OperatorHub](../install-the-snyk-controller/install-the-snyk-controller-with-openshift-4-and-operatorhub.md)
3. The Controller communicates with the Kubernetes API to determine which workloads are running on the cluster (for example, the Deployment, ReplicationController, CronJob, and so on), find their associated images, and scan them for vulnerabilities directly on the cluster.
4. For each workload that your collaborators import, Snyk displays the vulnerabilities found in each image, as well as a summary of the configuration issues that have been identified with the workload.
5. Snyk continuously monitors your imported workloads and reports new vulnerabilities it identifies whenever they affect your Projects.
6. Based on your configurations, if vulnerabilities are found, Snyk notifies you through email or Slack so that you can take immediate action.

{% hint style="info" %}
To maintain the health of the database, Snyk removes any information relating to a workload that has not been changed or updated for eight days. This can lead to failure when rescanning the workload.

If an image and its corresponding Project are removed, and you reimport the same workload during the eight days when the metadata still resides in the Snyk database, you can create the Project again.
{% endhint %}

## **The architecture of Snyk and Kubernetes integration**

<figure><img src="../../../../.gitbook/assets/System Diagram-Kubernetes integration.jpg" alt="Kubernetes integration architecture diagram"><figcaption><p>Kubernetes integration architecture diagram</p></figcaption></figure>

Based on the process exemplified in this diagram:

1. The customer's Snyk Organization is enabled for Kubernetes integration.
2. The customer installs the Snyk Controller into their Kubernetes cluster.
3. The Snyk Controller reads image information and pulls images from container registries.
4. The Snyk Controller scans the images.
5. The Snyk Controller sends scan results to the Snyk Platform to analyze issues.
6. The customer views the vulnerability issues on the Snyk Platform.

## **Terms and conditions**

Snyk Container Kubernetes integration uses Red Hat UBI (Universal Base Image).

Before downloading or using this application, you must agree to the Red Hat subscription agreement located on the [Red Hat website](https://www.redhat.com/en/about/agreements). If you do not agree with these terms, do not download or use the application.

If you have an existing Red Hat Enterprise Agreement (or other negotiated agreement with Red Hat) with terms that govern subscription services associated with Containers, then your existing agreement will be the controlling agreement.

## Kubernetes integration troubleshooting

For troubleshooting information on the Kubernetes integration or the Snyk Controller, as well as for common onboarding errors, see the [Kubernetes integration troubleshooting](https://support.snyk.io/s/article/Kubernetes-Integration-troubleshooting) support page.
