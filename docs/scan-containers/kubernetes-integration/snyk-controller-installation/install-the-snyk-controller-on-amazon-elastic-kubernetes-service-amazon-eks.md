# Install the Snyk controller on Amazon Elastic Kubernetes Service (Amazon EKS)

{% hint style="danger" %}
Before following the steps on this page, review the [prerequisites page](prerequisites-for-snyk-controller.md).
{% endhint %}

{% hint style="info" %}
These installation steps **work best** for EKS and ECR with the same AWS account. If you have a different setup, submit a request to [Snyk support](https://support.snyk.io/hc/en-us/requests/new).
{% endhint %}

Installing the Snyk Controller enables you to import and test your running EKS workloads and identify vulnerabilities in their associated images and configurations that might make those workloads less secure. After the workload is imported, Snyk continues to monitor the workload, identifying additional security issues as new images are deployed and the workload configuration changes.

This page provides instructions for configuring Snyk Controller to pull and scan private images from ECR.

Follow these steps to install Amazon EKS.

1\. Access your Kubernetes environment and run the following command in order to add the Snyk Charts repository to Helm:

```
helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor --force-update
```

2\. After the repository is added, create a unique namespace for the Snyk Controller:

```
kubectl create namespace snyk-monitor
```

{% hint style="info" %}
Use a unique namespace to isolate the controller resources easily. This is generally **good practice** for Kubernetes applications. Notice the namespace is called **snyk-monitor**; you will need this later when configuring other resources.
{% endhint %}

3\. Create a file named **dockercfg.json** and ensure it matches the following example:

```
{
  "credsStore": "ecr-login"
}
```

For additional setup for Private Registries,to see [Private Container Registry authentication](private-container-registry-authentication.md).

4\. Create a secret with your **Integration ID**, **Service Account Token** and **dockercfg.json** file added:

```
kubectl create secret generic snyk-monitor \
        -n snyk-monitor --from-file=dockercfg.json \
        --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234 \
        --from-literal=serviceAccountApiToken=bdca4123-dbca-4343-bbaa-1313cbad4231
```

5\. Attach **policies** or **roles** for nodes.

**Option 1: Attach policies for worker nodes**

a. Attach the **NodeInstanceRole** policy from [Using Amazon ECR Images with Amazon EKS](https://docs.aws.amazon.com/AmazonECR/latest/userguide/ECR\_on\_EKS.html)

b. Attach the **AmazonEC2ContainerRegistryReadOnly** policy to your EKS worker nodes.\
The Snyk Controller should now be able to pull private images when running on those worker nodes.

**Option 2: Create an EKS node role for your Node Group**

a. Follow the instructions on the page [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html); check your existing node role. Make sure you have attached the policy **AmazonEC2ContainerRegistryReadOnly.**

b. Select the **Details** tab on your EKS node group page, where you should see **Node IAM Role ARN.** It should look something like this:

```
arn:aws:iam::<role-id>:role/<role-name>
```

c. Create a **\<newFile>.yaml** with the following content:

```
volumes:
  projected:
    serviceAccountToken: true
    
securityContext:
  fsGroup: 65534

rbac:
  serviceAccount:
    annotations:
      eks.amazonaws.com/role-arn: <Node IAM Role ARN>
```

6\. Install the Snyk Controller.

After creating the IAM role for your Service Account, you can install your Snyk Controller with this newly created YAML file to overwrite the values in the Helm chart.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName=<ENTER_CLUSTER_NAME> \
             -f <newFile>.yaml
```
