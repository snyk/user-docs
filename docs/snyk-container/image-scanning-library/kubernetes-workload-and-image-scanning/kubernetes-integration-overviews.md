# Kubernetes integration overviews

Snyk integrates with Kubernetes, enabling you to import and test your running workloads and identify vulnerabilities in their associated images and configurations that might make those workloads less secure. Once imported, Snyk continues to monitor those workloads, identifying additional security issues as new images are deployed and the workload configuration changes.

{% hint style="info" %}
**Feature availability**  
This feature is available in Business and Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

**How it works**

1. Your administrator installs a controller on your cluster, authenticating the integration with a unique ID generated from the Snyk account. Install the controller with either of these options:
   * [Install the Snyk controller with Helm](https://support.snyk.io/hc/articles/360003916158#UUID-753328ea-3d73-0eeb-4301-c22522273797)
   * [Install the Snyk controller with OpenShift and OperatorHub](https://support.snyk.io/hc/articles/360006548317#UUID-7b1c8c43-51a6-d807-5623-e2338f830623)
   * [Install the Snyk controller on Amazon Elastic Kubernetes Service \(Amazon EKS\)](https://support.snyk.io/hc/en-us/articles/360011128137)
2. The controller communicates with the Kubernetes API to determine which workloads \(for instance the Deployment, ReplicationController, CronJob, etc.\) are running on the cluster, find their associated images, and scan them directly on the cluster for vulnerabilities. **Note**: dynamically injected sidecar containers are currently unsupported, and cannot be detected and scanned for vulnerabilities. These sidecars are normally injected by mutating admission controllers.
3. From Snyk, collaborators select which workloads to import, or workloads can be imported automatically using annotations. These options are as described in [Adding Kubernetes workloads for security scanning](https://support.snyk.io/hc/articles/360003947117#UUID-a0526554-0943-3363-6977-7a11f766ede2).
4. For each workload that your collaborators import, Snyk displays the vulnerabilities found in each image as well as a summary of configuration issues identified with the workload.
5. Snyk monitors your imported workloads on an ongoing basis, reporting on new vulnerabilities as they are disclosed whenever they affect your projects.
6. Based on your configurations, if vulnerabilities are found, Snyk notifies you via email or Slack so that you can take immediate action.

**Terms and conditions**

_The Snyk Container Kubernetes integration uses Red Hat UBI \(Universal Base Image\)._

_Before downloading or using this application, you must agree to the Red Hat subscription agreement located at redhat.com/licenses. If you do not agree with these terms, do not download or use the application. If you have an existing Red Hat Enterprise Agreement \(or other negotiated agreement with Red Hat\) with terms that govern subscription services associated with Containers, then your existing agreement will control._

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

