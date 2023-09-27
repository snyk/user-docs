# Insights setup: Kubernetes connector

## What is Kubernetes connector for Insights?

One of the goals of Insights is to identify risk factors for workloads that are publicly accessible via network configuration. To do that, Snyk needs to understand what images are deployed on which workloads, and how they are configured.&#x20;

So Snyk needs to collect the following information:

* The list of images and their IDs and SHAs that are deployed.
* The services and associated configuration, for example, the networking services you are using.

The Kubernetes connector for Insights is the agent deployed in your Kubernetes clusters to collect this information.&#x20;

## Getting started with Kubernetes connector for Insights

### Prerequisites

Before you can deploy the Kubernetes connector for Insights in your Kubernetes clusters, be sure you have the following:

* **Snyk Organization**: You need a Snyk Organization to which the Kubernetes information collected will be sent to be stored. This could be a new Organization, it does not have to be the same one containing the Snyk Projects you wish to use with Insights, but it must be in the same Snyk Group.&#x20;
* **Snyk service account:** You need to create a service account specifically to be used with the Kubernetes connector for Insights. For instructions on creating a service account, see [Service accounts](../../../enterprise-setup/service-accounts/).&#x20;
  * For the roles and permissions, Snyk recommends creating a new specific role for this service account and taking a least privilege approach, granting that role the sole permission required to **Publish Kubernetes Resources**.

### Step 1: Create a Snyk Organization

If you create a separate Organization for the Kubernetes connector for Insights, follow the steps in the Snyk documentation to [create a Snyk Organization](../../../snyk-admin/manage-groups-and-organizations/manage-organizations.md#create-an-organization). The new Snyk Organization must be in the same Snyk Group as your other Snyk Organization.&#x20;

If you are not creating a separate Snyk Organization, go to step 2.

### Step 2: Create a new role

Follow the steps in this documentation to [create a new role](../../../snyk-admin/manage-permissions-and-roles/manage-member-roles.md#create-a-role).

This example illustrates creating a new role called **Kubernetes connector for Insights.**

<figure><img src="../../../.gitbook/assets/image (14) (1) (1).png" alt="Create the Kubernetes connector for Insights role"><figcaption><p>Create the Kubernetes connector for Insights role</p></figcaption></figure>

### Step 3: Assign permissions to this role

Navigate to the newly created role and [select edit](../../../snyk-admin/manage-permissions-and-roles/manage-member-roles.md#edit-a-role); you will also be taken to this page immediately after creating the role.&#x20;

Scroll to the bottom of the page, tick the **Publish Kubernetes Resources** permission, and save the changes by clicking the **Update Role Permissions** button.&#x20;

<figure><img src="../../../.gitbook/assets/image (12) (1) (1) (1).png" alt="Publish Kubernetes Resources permission"><figcaption><p>Publish Kubernetes Resources permission</p></figcaption></figure>

### Step 4: Create a service account and assign it to a role

Next, you need to create a new [service account](../../../enterprise-setup/service-accounts/) for this integration.

{% hint style="info" %}
Snyk recommends creating this service account for the Snyk Organization used or created for the Kubernetes agent.&#x20;
{% endhint %}

Navigate to that **Snyk Organization -> Settings -> Service Account.**

Create a new service account with your chosen name, and from the drop-down, select the role you created in the previous step.

<figure><img src="../../../.gitbook/assets/image (11) (2) (1).png" alt="Select the Insights k8s Agent role"><figcaption><p>Select the Insights k8s Agent role</p></figcaption></figure>

After the service account is created, you will be shown the API token. Copy this down and store it somewhere safe; you’ll need this to configure the agent in the Helm chart.

### Step 5: Install Kubernetes connector for Insights in your Kubernetes clusters

Snyk recommends using the Helm Chart to deploy the agent; the Helm Chart will create the associated permissions for the agent to run on your cluster. The user installing the Helm Chart needs sufficient permissions on the Kubernetes cluster to create new roles. \
\
[Follow the instructions on the kubernetes-scanner GitHub repo](https://github.com/snyk/kubernetes-scanner) to use the Helm Chart to deploy the [latest released version](https://github.com/snyk/kubernetes-scanner/releases).

{% hint style="warning" %}
To ensure you have set up your Kubernetes connector properly, navigate to the **Set up Insights** tab on the **Insights** page and check the **Kubernetes resources** table to view the data that Insights has access to.
{% endhint %}

## FAQ

#### **What is the difference between the** [**Kubernetes monitor**](../../../scan-applications/snyk-container/kubernetes-integration/overview-of-the-kubernetes-integration/) **(also called Snyk Controller or Snyk-Monitor) and the Kubernetes connector for Insights?**

The Kubernetes **monitor** extracts images from a Kubernetes cluster’s workloads and scans them for vulnerabilities. The Kubernetes **connector** for Insights extracts workload configurations from a Kubernetes cluster.

#### **For Insights to work, do I need both or only the Kubernetes connector for Insights?**

You need only the Kubernetes connector for Insights installed in your Kubernetes clusters.

#### **If I’m a customer and already use the existing agent, do I also need to install the Kubernetes connector for Insights?**

If you want Insights to be able to compute both the Deployed and Public Facing risk factors, you must install the Kubernetes connector into your Kubernetes clusters.

If you have only the existing agent installed, Insights is able to compute only the Deployed risk factor.

#### **What workload configuration the Kubernetes connector for Insights is collecting?**

[Here](https://github.com/snyk/kubernetes-scanner/blob/main/helm/kubernetes-scanner/values.yaml) is the list of data Snyk is collecting.

**What happens to my data if I delete the Kubernetes connector?**

If you choose to delete the Kubernetes connector, your data and risk factors will still be available for 48 hours.

**What happens to my data if the Kubernetes connector stops working?**

If, for any reason, the Kubernetes connector stops working, your data and risk factors will still be available for 48 hours.
