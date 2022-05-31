# Install the Snyk controller on Amazon Elastic Kubernetes Service (Amazon EKS)

{% hint style="danger" %}
Before following this installation page, please review the [prerequisite setting page](prerequisite-setting.md).
{% endhint %}

{% hint style="info" %}
These installation steps **work best** for EKS and ECR with the same AWS account. If you have a different setup, please contact [Snyk support](https://snyk.zendesk.com/agent/dashboard).
{% endhint %}

Installing the Snyk controller enables you to import and test your running EKS workloads and identify vulnerabilities in their associated images and configurations that might make those workloads less secure. Once imported, Snyk continues to monitor those workloads, identifying additional security issues as new images are deployed and the workload configuration changes.

Follow these instructions to **configure Snyk Controller to pull and scan private images from ECR**.

**Installation steps**

1\. Access your Kubernetes environment and run the following command in order to add the Snyk Charts repository to Helm:

```
helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor --force-update
```

2\. Once the repository is added, create a unique namespace for the Snyk controller:

```
kubectl create namespace snyk-monitor
```

**Tip:** Use a unique namespace to isolate the controller resources easily. This is generally **good practice** for Kubernetes applications. Notice our namespace is called **snyk-monitor**; you’ll need this later when configuring other resources.

3\. Create a file named **dockercfg.json** and ensure it matches the following example:

```
{
  "credsStore": "ecr-login"
}
```

4\. Create a secret with the **dockercfg.json** file added:

```
kubectl create secret generic snyk-monitor \
        -n snyk-monitor --from-file=dockercfg.json \
        --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234
```

5\. Attach **policies** or **roles** for nodes

<mark style="color:purple;">**Option 1:**</mark> <mark style="color:purple;"></mark><mark style="color:purple;">Attach policies for worker nodes</mark>

a. attach the **NodeInstanceRole** policy from [Using Amazon ECR Images with Amazon EKS](https://docs.aws.amazon.com/AmazonECR/latest/userguide/ECR\_on\_EKS.html)&#x20;

b.  attach **AmazonEC2ContainerRegistryReadOnly** policy to your EKS worker nodes. \
The Snyk Controller should now be able to pull private images when running on those worker nodes.

<mark style="color:purple;">**Option 2:**</mark> <mark style="color:purple;"></mark><mark style="color:purple;">Creating an EKS node role for your Node Group</mark>

a. Follow the instructions on [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html), check your existing node role. Make sure you have attached the policy **AmazonEC2ContainerRegistryReadOnly.**

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

6\. Install Snyk Controller

After creating the IAM role for your Service Account, you can install your Snyk Controller with this newly created YAML file to overwrite the values in the Helm chart.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             -f <newFile>.yaml
```

