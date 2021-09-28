# Module 3

## Red Hat OpenShift container platform

In order to complete these steps, you **must** have a running OpenShift 4 cluster. In **Getting Started**, we listed prerequisites for this workshop and the steps found in [Red Hat's Get started with OpenShift](https://www.openshift.com/try) guide should be followed if you don't already have a cluster to work on.

{% hint style="warning" %}
You will need to install the [OpenShift Container Platform command-line interface \(CLI\)](https://docs.openshift.com/container-platform/4.2/cli_reference/openshift_cli/getting-started-cli.html) before you continue.
{% endhint %}

### Securing Kubernetes workloads

#### Authenticating to your cluster

Once you have installed the CLI, let's start with authenticating to our cluster:

```bash
oc login
```

#### Deploying the application

Next, we will [create a project ](https://docs.openshift.com/container-platform/4.2/cli_reference/openshift_cli/getting-started-cli.html#creating-a-project)using the `oc new-project` command:

```bash
oc new-project demo
```

Next, [create a new app](https://docs.openshift.com/container-platform/4.2/cli_reference/openshift_cli/getting-started-cli.html#creating-a-new-app) with the `oc new-app` command. Here we will create our backend using the `mongo` container image and pass the `--name` parameter to name our app `goof-mongo`:

```bash
oc new-app mongo --name goof-mongo
```

Next, we can _optionally_ watch the status bu running the `oc logs` command against our deployment configuration:

```bash
oc logs -f dc/goof-mongo
```

We can also _optionally_ view the status of the service created by running the `oc get svc` command:

```bash
oc get svc goof-mongo
```

Now, we are ready to deploy our frontend application. Here we will do a few things:

1. Pull the container image we previously pushed to our quay repository.
2. Set the `COMPONENT_BACKEND_HOST` & `COMPONENT_BACKEND_PORT` to allow our frontend to talk to our backend.

```bash
oc new-app quay.io/<QUAY_USER>/goof \
--name goof -e COMPONENT_BACKEND_HOST=goof-mongo \
-e COMPONENT_BACKEND_PORT=27017
```

We will also need to set an environment variable for our frontend app by running the `oc set env` command against our frontend deployment configuration.

```bash
oc set env dc/goof DOCKER=1
```

Similar to what we did for our backend, we can _optionally_ run `oc get svc` to see the status of our frontend service.

```bash
oc get svc goof
```

Now that our frontend is deployed, we will need to expose this so that we can access our app from a web browser. For this, we will run the `oc expose` command against our frontend service.

```bash
oc expose svc/goof
```

Lastly, we will run `oc get routes` to get the URL which we will copy and paste into our browser.

```bash
oc get routes
```

You can see the results of these commands by logging into the OpenShift console and reviewing the **Deployment Configs** for your project.

![](../../.gitbook/assets/openshift-deployment-config.png)

#### Installing the Snyk controller with OpenShift 4 and OperatorHub

We mentioned in **Getting Started** that you will need to install the Snyk controller to your cluster. If you have not already done so, now is a good time to go through the steps documented in our [guide](https://support.snyk.io/hc/en-us/articles/360006548317-Install-the-Snyk-controller-with-OpenShift-4-and-OperatorHub).

#### Monitoring Kubernetes workloads

By logging into your Snyk account, you will be able to [add Kubernetes workloads for security scanning](https://support.snyk.io/hc/en-us/articles/360003947117-Adding-Kubernetes-workloads-for-security-scanning). In this example, we are going to monitor our recently deployed **goof** application.

![](../../.gitbook/assets/kubernetes-integration-01.png)

Once you **Add selected workloads** you will be taken back to the **Projects** page where you can see a high level overview of all imported projects including your Kubernetes application.

![](../../.gitbook/assets/kubernetes-integration-02.png)

By clicking on the the `demo/replicationcontroller/goof-1` project, you will see additional details specific to the application security configuration. In this example, we see a number of issues have been found. 

As we now know, Kubernetes is not secure by default. It is beyond the scope of these exercises to provide a deep-dive on this topic, but detailed documentation on [securing a cluster](https://kubernetes.io/docs/tasks/administer-cluster/securing-a-cluster/) is readily available. Additional reading is recommended on configuring a [security context](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/) for a pod, [pod security](https://kubernetes.io/docs/concepts/policy/pod-security-policy/) policies, and setting [capabilities](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-capabilities-for-a-container) for a container.

![](../../.gitbook/assets/kubernetes-integration-03.png)

