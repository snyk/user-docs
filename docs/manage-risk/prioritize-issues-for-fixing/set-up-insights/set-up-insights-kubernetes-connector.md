# Set up Insights: Kubernetes connector

## What is Kubernetes connector?

One of the goals is to identify risk factors for workloads that are publicly accessible through a network configuration. To do that, Snyk needs to understand what images are deployed on which workloads and how they are configured.

Thus, Snyk needs to collect the following information:

* The list of images and their IDs and SHAs that are deployed.
* The services and associated configuration, for example, the networking services you are using.

The Kubernetes connector is the agent deployed in your Kubernetes clusters to collect this information.

## Getting started

### Prerequisites

Before you can deploy the Kubernetes connector in your Kubernetes clusters, ensure you have the following:

* A Snyk Organization to which the Kubernetes information collected will be sent to be stored. This could be a new Organization.
* A Snyk service account created specifically to be used with the Kubernetes connector. For instructions on creating a service account, see [Service accounts](../../../implementation-and-setup/enterprise-setup/service-accounts/). For the roles and permissions, Snyk recommends:
  * Creating a new specific role for this service account
  * Taking a least privilege approach, granting the new specific role the sole permission required to **Publish Kubernetes Resources**.

### Create a Snyk Organization

If you create a separate Organization for the Kubernetes connector, follow the steps in the Snyk documentation to [create a Snyk Organization](../../../snyk-platform-administration/groups-and-organizations/organizations/create-and-delete-organizations.md#create-an-organization). The new Snyk Organization must be in the same Snyk Group as your other Snyk Organization.

If you are not creating a separate Snyk Organization, continue with the next step.

### Create a new role

To create a new role, see [create a new role](../../../snyk-platform-administration/user-roles/user-role-management.md#create-a-custom-role).

This example illustrates creating a new role called **Kubernetes connector**

<figure><img src="../../../.gitbook/assets/image (94).png" alt="Create the Kubernetes connector for Insights role"><figcaption><p>Create the Kubernetes connector for role</p></figcaption></figure>

### Assign permissions to this role

Navigate to the newly created role and [select edit](../../../snyk-platform-administration/user-roles/user-role-management.md#edit-a-custom-role); you will also be taken to this page immediately after creating the role.

Scroll to the bottom of the page, tick the **Publish Kubernetes Resources** permission, and save the changes by clicking the **Update Role Permissions** button.

<figure><img src="../../../.gitbook/assets/image (35).png" alt="Publish Kubernetes Resources permission"><figcaption><p>Publish Kubernetes Resources permission</p></figcaption></figure>

### Create a service account and assign it to a role

Next, create a new [service account](../../../implementation-and-setup/enterprise-setup/service-accounts/) for Kubernetes connector integration.

{% hint style="info" %}
Snyk recommends creating this service account for the Snyk Organization used or created for the Kubernetes agent.
{% endhint %}

Navigate to that **Snyk Organization** > **Settings** > **Service Account.**

Create a new service account with your chosen name, and from the dropdown, select the role you created in the previous step.

<figure><img src="../../../.gitbook/assets/image (32).png" alt="Select the Insights Kubernetes Agent role"><figcaption><p>Select the Kubernetes Agent role</p></figcaption></figure>

After the service account is created, you will see the API token. Copy the API token and save it somewhere safe; you will need this to configure the agent in the Helm chart.

### Install the Kubernetes connector in your Kubernetes clusters

Snyk recommends using the Helm Chart to deploy the agent; the Helm Chart will create the associated permissions for the agent to run on your cluster. The user installing the Helm Chart needs sufficient permissions on the Kubernetes cluster to create new roles.\
\
[Follow the instructions on the Kubernetes-scanner GitHub repo](https://github.com/snyk/kubernetes-scanner) to use the Helm Chart to deploy the [latest released version](https://github.com/snyk/kubernetes-scanner/releases).

{% hint style="info" %}
To ensure you have set up your Kubernetes connector properly, navigate to the **Set up Insights** tab on the **Issues** page and check the **Kubernetes resources** table to view the data that Insights has access to.
{% endhint %}

## FAQ

#### **What is the difference between the** [**Kubernetes monitor**](../../../scan-with-snyk/snyk-container/kubernetes-integration/overview-of-kubernetes-integration/) **(also called Snyk Controller or Snyk-Monitor) and the Kubernetes connector?**

* The Kubernetes monitor extracts images from a Kubernetes cluster’s workloads and scans them for vulnerabilities. The Kubernetes monitor reports the **Deployed** risk factor.
* The Kubernetes connector extracts workload configurations from a Kubernetes cluster. The Kubernetes connector reports the **Public facing** and **Deployed** risk factors.

{% hint style="warning" %}
Risk factors are available only based on the integration option you implemented.
{% endhint %}

#### **If I’m a customer and already use the existing agent, do I also need to install the Kubernetes connector?**

If you want to view both the **Deployed** and **Public Facing** risk factors, you must install the Kubernetes connector into your Kubernetes clusters.

If you have only the existing agent installed, Snyk can compute only the **Deployed** risk factor.

#### **What workload data is the Kubernetes connector collecting?**

[The list of data Snyk is collecting](https://github.com/snyk/kubernetes-scanner/blob/main/helm/kubernetes-scanner/values.yaml) is on the Kubernetes scanner GitHub repo.

**What happens to my data if I delete the Kubernetes connector?**

If you choose to delete the Kubernetes connector, your data and risk factors will remain available for 48 hours.

**What happens to my data if the Kubernetes connector stops working?**

If, for any reason, the Kubernetes connector stops working, your data and risk factors will remain available for 48 hours.
