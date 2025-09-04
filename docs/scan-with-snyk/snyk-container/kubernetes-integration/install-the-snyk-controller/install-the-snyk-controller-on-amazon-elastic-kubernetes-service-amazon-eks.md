# Install the Snyk Controller on Amazon Elastic Kubernetes Service (Amazon EKS)

{% hint style="info" %}
Ensure you have reviewed the [prerequisites for installing the Snyk Controller](./#prerequisites-for-installing-the-snyk-controller).
{% endhint %}

{% hint style="info" %}
The installation steps work best for EKS and ECR with the same AWS account. If you have a different setup, contact [Snyk support](https://support.snyk.io).
{% endhint %}

Installing the Snyk Controller enables you to import and test your running EKS workloads and identify vulnerabilities in their associated images and configurations that can make the workloads less secure. After the workload is imported, Snyk continues to monitor the workload, identifying additional security issues, as new images are deployed and the workload configuration changes.

The steps described below provide instructions for configuring the Snyk Controller to pull and scan private images from ECR.

To install Amazon EKS:

1\. Access your Kubernetes environment. Run the following command in order to add the Snyk Charts repository to Helm:

```
helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor --force-update
```

2\. After the repository is added, create a unique namespace for the Snyk Controller:

```
kubectl create namespace snyk-monitor
```

{% hint style="info" %}
As a good practice for Kubernetes applications, use a unique namespace to isolate the controller resources easily.

Ensure you remember the namespace `snyk-monitor`. You will use it when configuring other resources.
{% endhint %}

3\. Create a file named "dockercfg.json" and ensure it matches the following example:

```
{
  "credsStore": "ecr-login"
}
```

For additional setup for private registries, see [Authenticate to private container registries](authenticate-to-private-container-registries.md).

4\. Create a Kubernetes secret containing your Integration ID, service account token, and dockercfg.json file:

```
kubectl create secret generic snyk-monitor \
        -n snyk-monitor --from-file=dockercfg.json \
        --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234 \
        --from-literal=serviceAccountApiToken=bdca4123-dbca-4343-bbaa-1313cbad4231
```

5\. Attach policies or roles for nodes. You can do this using one of the below options:

#### Attach policies for worker nodes

1. Attach the `NodeInstanceRole` policy. See [Using Amazon ECR Images with Amazon EKS](https://docs.aws.amazon.com/AmazonECR/latest/userguide/ECR_on_EKS.html).
2. Attach the `AmazonEC2ContainerRegistryReadOnly` policy to your EKS worker nodes.\
   The Snyk Controller is now able to pull private images when running on those worker nodes.

#### **Create an EKS node role for your Node Group**

1. **Create an EKS node role for your Node Group**
   
   Follow the instructions on [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html) and check your existing node role. Ensure you have attached the policy `AmazonEC2ContainerRegistryReadOnly`.

2. **Add the Trust Relationship for the IAM Role**
   
   For the Snyk Controller to successfully assume the IAM role and pull images from ECR, the role's trust policy must be configured to trust your EKS cluster's OIDC provider.

   In the AWS Management Console, navigate to the IAM role created for your EKS node group.

   Select the Trust relationships tab and click Edit trust policy.

   Add the following JSON to the policy document. Replace the placeholders `xxx` and `oidc.eks.us-east-1.amazonaws.com/id/xxx` with your specific account ID and OIDC provider URL.

   ```json
   {
     "Version": "2012-10-17",
     "Statement": [
       {
         "Effect": "Allow",
         "Principal": {
           "Federated": "arn:aws:iam::xxx:oidc-provider/oidc.eks.us-east-1.amazonaws.com/id/xxx"
         },
         "Action": "sts:AssumeRoleWithWebIdentity",
         "Condition": {
           "StringEquals": {
             "oidc.eks.us-east-1.amazonaws.com/id/xxx:sub": "system:serviceaccount:snyk-monitor:snyk-monitor"
           }
         }
       }
     ]
   }
   ```

   Click Update policy to save your changes. 
3. Navigate to the **Details** tab on your EKS node group page, where you will see `Node IAM Role ARN`

```
arn:aws:iam::<role-id>:role/<role-name>
```

4. Create a \<newFile>.yaml with the following content:

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

5. Install the Snyk Controller.

After creating the IAM role for your service account, you can install your Snyk Controller with the newly created YAML file to overwrite the values in the Helm chart.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName=<ENTER_CLUSTER_NAME> \
             -f <newFile>.yaml
```
