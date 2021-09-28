# Getting Started

## Securing your Red Hat OpenShift workloads

![](../../.gitbook/assets/redhat-snyk-pipeline%20%281%29.png)

### Welcome!

This workshop will take you through a series of exercises created with one purpose in mind: to provide you with hands-on examples that demonstrate how you can integrate Snyk into your Red Hat workflows in order to identify and fix potential vulnerabilities in your applications.

### Prerequisites

The examples presented in these modules will require some supporting infrastructure deployed and available for you to use. This will consist of a [Red Hat OpenShift](http://try.openshift.com/) cluster, a [Red Hat Quay.io](https://quay.io/) private registry, a [Snyk account](https://app.snyk.io/login), and some supporting sample code available in our [GitHub repository](https://github.com/snyk-partners/patterns-library-redhat).

{% hint style="danger" %}
It is **NOT** recommended that you use production systems for this workshop.
{% endhint %}

#### Red Hat OpenShift cluster

The recommended deployment method for this workshop is to install Red Hat OpenShift 4 in your account on any of the supported public cloud providers. Detailed guidance on the steps needed to do this are available on [Red Hat's Get started with OpenShift](https://www.openshift.com/try) site.

#### Red Hat Quay.io private container registry

There are a few ways to deploy [Quay](https://quay.io/). While functionally the steps contained in these modules will be the same irrespective of how you deploy this registry, we have opted for the cloud.

#### Snyk controller with OpenShift 4

To get vulnerability details about your Kubernetes workloads running on OpenShift, you must first install the Snyk controller onto your cluster. The Snyk monitor requires some minimal configuration items in order to work correctly. The necessary steps are detailed in [Snyk's Knowledge Center](https://support.snyk.io/hc/en-us/articles/360006548317#UUID-7b1c8c43-51a6-d807-5623-e2338f830623).

#### Microsoft Visual Studio Code with Dependency Analytics extension

If you do not already have VSCode, you should [download](https://code.visualstudio.com/download) it for free. We will leverage [Red Hat Dependency Analytics](https://marketplace.visualstudio.com/items?itemName=redhat.fabric8-analytics) extension available in the Visual Studio Marketplace. Dependency Analytics is powered by [Snyk Intel Vulnerability DB](https://snyk.io/product/vulnerability-database/), it is the most advanced and accurate open source vulnerability database in the industry. That adds value with the latest, fastest and more number of vulnerabilities derived from numerous sources.

#### Sample application and supplemental resources

In our examples, we will build a container image for Snyk's vulnerable demo app, Goof. You will need to [`git clone`](https://github.com/snyk-partners/patterns-library-redhat.git) the repository in order to complete these exercises.

