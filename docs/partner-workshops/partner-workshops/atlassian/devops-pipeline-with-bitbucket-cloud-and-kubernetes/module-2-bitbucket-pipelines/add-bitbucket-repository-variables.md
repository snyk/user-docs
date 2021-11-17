---
description: Tailor your workshop with details specific to you
---

# Add Bitbucket Repository Variables

In this section, we'll add variables to customize the workshop for your instance.  Below, we list the variable name, whether is it a secured variable, and we describe its purpose.

### Add Repository Variables

Once you enable pipelines in Bitbucket, you will need to add [Repository Variables](../../../setting-up-your-workshops/atlassian-integrations/atlassian-bitbucket-pipeline-variables.md) to permit parameterized operations. &#x20;

You will need permissions to add variables in your repository.

When we review the pipeline definition, you will see variable references to scan code, and publish results to AWS services and these are the repository variables you will need:

* Snyk API token for authenticating with your Snyk account: `SNYK_TOKEN.` We'll come back to this variable after we enable integrations in the next step.
* This value is a secured variable
* This is the service account API token you created
* AWS Identity & Access Management User [key and secret](https://docs.aws.amazon.com/IAM/latest/UserGuide/id\_credentials\_access-keys.html) for secure authenticated interactions with the AWS API: `AWS_ACCESS_KEY_ID` & `AWS_SECRET_ACCESS_KEY`
  * These are secured variables.
  * These are the API keys for your access into AWS.
* AWS region you will be deploying to: `AWS_DEFAULT_REGION`
  * This is not a secured variable.
  * This is a region to where your ECR registry is accessible from.  For example - us-east-1.
* Amazon ECR [URL](https://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html) for your repository: `AWS_ECR_URI`
  * This is not a secured variable.
  * This is the private ECR URI for your repository of this form: [https://aws\_account\_id.dkr.ecr.region.amazonaws.com](https://aws\_account\_id.dkr.ecr.region.amazonaws.com)
* Container image name: `IMAGE`
  * This is not a secured variable.
  * This is the name of your image, such as java-goof.
* Amazon EKS name of your cluster: `AWS_EKS_CLUSTER`
  * This value is not a secured variable.
  * This is the name of your EKS cluster.
  * This workshop does not cover the EKS cluster, but you may set this value if you continue with the example and add a deployment to a cluster.

Once you have configured the variables, we can now review the pipeline and trigger a run in the next section.
