# Install the Snyk controller on Amazon Elastic Kubernetes Service (Amazon EKS)

{% hint style="info" %}
Before following this installation page, please review the [prerequisite setting page](../prerequisite-setting.md).
{% endhint %}

Installing the Snyk controller enables you to import and test your running EKS workloads and identify vulnerabilities in their associated images and configurations that might make those workloads less secure. Once imported, Snyk continues to monitor those workloads, identifying additional security issues as new images are deployed and the workload configuration changes.

Follow these instructions to **configure Snyk Controller to pull and scan private images from ECR**.

{% hint style="info" %}
These installation steps **work best** for EKS and ECR with the same AWS account. If you have a different setup, please contact [Snyk support](https://snyk.zendesk.com/agent/dashboard).
{% endhint %}

Ensure your **dockerconfig.json** matches the following example:

```
{
  "credsStore": "ecr-login"
}
```

## Attach policies for worker nodes

For all the preceding options, attach the **NodeInstanceRole** policy from [Using Amazon ECR Images with Amazon EKS](https://docs.aws.amazon.com/AmazonECR/latest/userguide/ECR\_on\_EKS.html) and the **AmazonEC2ContainerRegistryReadOnly** policy to your EKS worker nodes. The Snyk Controller should now be able to pull private images when running on those worker nodes.

Alternatively, you can also use the IAM role for Service Accounts by **creating an EKS node role for your Node Group** (no need to set up an extra OIDC service account role), and configure the Snyk Controller as follows.

## Create an EKS node role for your Node Group

* Following the instructions on [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html), check your existing node role. Make sure you have attached the policy **AmazonEC2ContainerRegistryReadOnly.**
* Select the **Details** tab on your EKS node group page, where you should see **Node IAM Role ARN.** It should look something like this:

```
arn:aws:iam::<role-id>:role/<role-name>
```

* Create a YAML file with the following content:

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

## Install Snyk Controller

After creating the IAM role for your Service Account, you can install your Snyk Controller with this newly created YAML file to overwrite the values in the Helm chart.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             -f <newFile>.yaml
```
