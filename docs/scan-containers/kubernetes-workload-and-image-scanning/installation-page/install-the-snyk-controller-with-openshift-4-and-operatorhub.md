# Install the Snyk controller with OpenShift 4 and OperatorHub

## Prerequisites

{% hint style="danger" %}
Before following this installation page, please review the [prerequisite setting page](prerequisite-setting.md)
{% endhint %}

To get vulnerability details about your Kubernetes workloads, a Snyk admin must first install the Snyk controller onto your cluster. For most Kubernetes platforms, the Snyk monitor requires only some minimal configuration items in order to work correctly.

As with any Kubernetes deployment, the Snyk controller runs within a single namespace.

This section covers:

* The Snyk controller is installed using OperatorHub through RedHat OpenShift 4.
* A minimum of 50 GB of storage must be available in the form of an [emptyDir](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir) on the cluster and the person configuring the cluster must be an administrator.

## Installation steps

1.  Access your Kubernetes environment and run the following command in order to add the Snyk Charts repository to Helm:

    ```
    helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor --force-update
    ```
2. Ensure that your Kubernetes command-line tool (`kubectl`) points to the relevant cluster.
3.  Create a unique namespace for the Snyk controller with Cluster scope to enable the controller to monitor all of its deployments:

    ```
    kubectl create namespace snyk-monitor
    ```

    **Tip**\
    **1.** Use a unique namespace to isolate the controller resources more easily as generally good practice for Kubernetes applications. Notice our namespace is called `snyk-monitor`, you’ll need this later when configuring other resources.\
    **2.** Cluster scope is the default scope and we recommend you use this when installing the namespace so that we can scan the entire cluster. You can choose namespace scope, in which case the Snyk controller will watch for workloads only in the namespace in which it is deployed.
4.  The Snyk monitor runs by using your Snyk **Integration ID**, and using a `dockercfg.json` file.\
    \
    <mark style="color:purple;">**Public Container Registry**</mark>\
    Create a Kubernetes secret called `snyk-monitor` containing the Snyk **Integration ID** from the previous step and run the following command:

    ```
    kubectl create secret generic snyk-monitor -n snyk-monitor --from-literal=dockercfg.json={} --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234
    ```

    **Note**: The secret must be called `snyk-monitor` in order for the integration to work.\
    \
    <mark style="color:purple;">**Private Container Registry**</mark>\
    Create a file named `dockercfg.json` and provide credentials to access those registries by creating a secret (which must be called `snyk-monitor`) using both the Snyk **Integration ID** as well as a `dockercfg.json` file.\
    \
    The `dockercfg.json` file is necessary to allow the monitor to look up images in private registries. Usually a copy of the `dockercfg.json` resides in `$HOME/.docker/config.json`.\
    \
    Create a `dockercfg.json` configuration file:

    ```
    {
      "auths": {
        "gcr.io": {
          "auth": "BASE64-ENCODED-AUTH-DETAILS"
        }
        // Add other registries as necessary
      }
    }
    ```

    \
    If you are using **Artifactory Container Registry to host multiple private repositories,** your `dockercfg.json` file should look like this:

    ```
    {
      "auths": {
        "<registry1>": {
            "auth": "BASE64-ENCODED-AUTH-DETAILS"
           },
        "<registry2>": {
           "auth": "BASE64-ENCODED-AUTH-DETAILS"
        }
      }
    }
    ```

    \
    Create a secret with the file added:

    ```
    kubectl create secret generic snyk-monitor -n snyk-monitor --from-file=dockercfg.json --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234
    ```
5. Log in to your OpenShift Container Platform (OCP) web console, navigate to OperatorHub and then search for **Snyk** to install the desired version of **Snyk Operator:**
   * The **Marketplace** version is certified by Red Hat and listed in the OpenShift marketplace, and is the recommended installation option.
   * The **Community** version is released more frequently and may be more up-to-date, but may not yet have been certified for OpenShift.
6. Double-check installation was successful from the **Installed Operators** area
7.  Now, from the **Subscription** tab, create an Operator Subscription for the **Snyk controller**. Select "A specific namespace on the cluster" or leave the default “All namespaces on the cluster” based on your needs.

    **Note**:\
    Snyk always uses a stable channel when scanning your workloads

    1. Leave the remaining default configurations as they are.
    2. Click Subscribe to make the Operator available to the namespaces on this OpenShift Container Platform cluster.
8. Now, create an instance of the **Snyk Monitor**. From the **Snyk Monitor** custom resource, click Create instance.
9. Double-check successful installation from the cluster.
10. After successfully installing the **Snyk Operator** and the instance of a **Snyk Monitor**, you can also view your cluster in Snyk.

![Example of successful installation from the cluster.](<../../../.gitbook/assets/image (40) (1).png>)

## Optional installation steps

For any additional Snyk Controller optional steps that fit your environment refer to [Optional installation steps for Snyk Controller with Helm](optional-installation-steps-for-snyk-controller-with-helm.md)[.](optional-installation-steps-for-snyk-controller-with-helm.md)
