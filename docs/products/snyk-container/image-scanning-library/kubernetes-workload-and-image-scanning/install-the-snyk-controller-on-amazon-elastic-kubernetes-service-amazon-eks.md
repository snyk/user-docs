# Install the Snyk controller on Amazon Elastic Kubernetes Service (Amazon EKS)

Installing the Snyk controller enables you to import and test your running EKS workloads and identify vulnerabilities in their associated images and configurations that might make those workloads less secure. Once imported, Snyk continues to monitor those workloads, identifying additional security issues as new images are deployed and the workload configuration changes.

### [Open the AWS Quick Start Guide](https://aws.amazon.com/quickstart/architecture/eks-snyk/)

You have the option of deploying the Snyk controller for Amazon EKS as an official [AWS Quick Start](https://aws.amazon.com/quickstart/architecture/eks-snyk/). This option eliminates the need for manual configuration. Deploying this Quick Start with default parameters into an existing Amazon EKS cluster builds the following environment.

![](<../../../../.gitbook/assets/architecture (1).png>)

There are three deployment options to match most common use cases. These are as follows:

**1.** If you already have an Amazon EKS cluster running in your AWS account

[![cloudformation-launch-stack.png - REPLACE THIS IMAGE - ZENDESK IMAGE - UPDATE ME!](../../../../.gitbook/assets/cloudformation-launch-stack.png)](https://us-east-2.console.aws.amazon.com/cloudformation/home?region=us-east-2#/stacks/create/template?stackName=Snyk-EKS\&templateURL=https://aws-quickstart.s3.us-east-1.amazonaws.com/quickstart-amazon-eks/submodules/quickstart-eks-snyk/templates/eks-snyk.template.yaml)

**2.** If you already have an Amazon Virtual Private Cloud (Amazon VPC) but need an Amazon EKS cluster _**with**_ the Snyk controller deployed to the cluster

[![cloudformation-launch-stack.png - REPLACE THIS IMAGE - ZENDESK IMAGE - UPDATE ME!](../../../../.gitbook/assets/cloudformation-launch-stack.png)](https://us-east-2.console.aws.amazon.com/cloudformation/home?region=us-east-2#/stacks/create/template?stackName=Amazon-EKS-with-Snyk\&templateURL=https://aws-quickstart.s3.us-east-1.amazonaws.com/quickstart-amazon-eks/templates/amazon-eks-master-existing-vpc.template.yaml)

**3.** If you have neither an Amazon VPC or Amazon EKS cluster and need all services _**with**_ the Snyk controller deployed to the cluster

[![cloudformation-launch-stack.png - REPLACE THIS IMAGE - ZENDESK IMAGE - UPDATE ME!](../../../../.gitbook/assets/cloudformation-launch-stack.png)](https://us-east-2.console.aws.amazon.com/cloudformation/home?region=us-east-2#/stacks/create/template?stackName=Amazon-EKS-with-Snyk\&templateURL=https://aws-quickstart.s3.us-east-1.amazonaws.com/quickstart-amazon-eks/templates/amazon-eks-master-existing-vpc.template.yaml)

**NOTE:** _Please review the_ [_parameter reference_](https://github.com/aws-quickstart/quickstart-eks-snyk#parameter-reference) _prior to deployment._

## **Prerequisites**

{% hint style="info" %}
**Feature availability**\
This feature is available with all paid plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

* An administrator account for your Snyk organization.
* A minimum of 50 GB of storage must be available in the form of an [emptyDir](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir) on the cluster.
* Your Kubernetes cluster needs to be able to communicate with Snyk outbound over HTTPS.
* When configuring Snyk to integrate with an Amazon Elastic Kubernetes Services (EKS) cluster, if you wish to scan images hosted on your Amazon Elastic Container Registry (ECR), you may also deploy our Quick Start, [Snyk Security on AWS](https://aws.amazon.com/quickstart/architecture/snyk-security/) to enable this integration.

![](../../../../.gitbook/assets/snyk\_rocket.png)

## [Deployment Guide](https://aws-quickstart.github.io/quickstart-eks-snyk/)

## Configure Snyk Controller to pull and scan private images from ECR

Ensure your **dockerconfig.json** matches the example below:

```
{
    "credsStore": "ecr-login"
  }
```

### Attach policies for worker nodes

For all the options above, attach the **NodeInstanceRole** policy that can be found [here](https://docs.aws.amazon.com/AmazonECR/latest/userguide/ECR\_on\_EKS.html) with the **AmazonEC2ContainerRegistryReadOnly** policy to your EKS worker nodes. The Snyk Controller should now be able to pull private images when running on those worker nodes.&#x20;

Alternatively, you can also use the IAM role for Service Accounts by **creating EKS node role for your Node Group**, and configure the Snyk Controller as follows:

### Create an EKS node role for your Node Group

* Following the instruction [here](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html), check your existing node role. Make sure you have attached the policy **AmazonEC2ContainerRegistryReadOnly.**
* Select the **Details** tab on your EKS node group page, where you should see **Node IAM Role ARN.** It should look something like this:

```
arn:aws:iam::<role-id>:role/<role-name>
```

* Create a YAML file with the following content:

```
 volumes:
  projected:
    serviceAccountToken: true

rbac:
  serviceAccount:
    annotations:
      eks.amazonaws.com/role-arn: <Node IAM Role ARN>
```

### Install Snyk Controller&#x20;

After creating the IAM role for your Service Account, you can now install your Snyk Controller with this newly created YAML file to overwrite the values in the Helm chart.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             -f <newFile>.yaml
```
