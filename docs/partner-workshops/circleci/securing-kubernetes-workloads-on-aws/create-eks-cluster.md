# Create EKS cluster

There are several options available for creating an [Amazon Elastic Kubernetes Service \(Amazon EKS\)](https://aws.amazon.com/eks/) cluster on AWS. Among those are the following:

* The [Modular and Scalable Amazon EKS Architecture](https://aws.amazon.com/quickstart/architecture/amazon-eks/) reference deployment developed by AWS solutions architects.
* The CircleCI [`aws-eks`](https://circleci.com/orbs/registry/orb/circleci/aws-eks#jobs-create-cluster) Orb.
* The `eksctl` command-line tool.

For the purposes of these exercises, `eksctl` will suffice. To create the cluster, run the following command:

```bash
eksctl create cluster \
--name snyk-circleci-eks \
--version 1.16 \
--region us-west-2
```

{% hint style="info" %}
You can replace `us-west-2` with any [Amazon EKS Fargate supported Region](https://docs.aws.amazon.com/eks/latest/userguide/fargate.html). However, **do not** replace the name as it will be referenced in subsequent instructions.
{% endhint %}

It can take approximately 15 minutes for your cluster to be provisioned. When your cluster is ready, you will be able to run `kubectl get svc` and see results.

![](../../../.gitbook/assets/kubectl_get_svc.gif)

