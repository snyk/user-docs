# Configure AWS ECR

It is convenient to deploy your container image to a container registry, and we use AWS ECR because we also deploy the container to AWS EKS in the next module.  Other pipelines may use other registries and Kubernetes environments, and the concepts here are still applicable because of the focus on the configuration parameters.

We start by setting up [AWS ECR](../../../getting-started/aws-integrations/aws-ecr.md).  From those instructions, you will need this information for future steps:

* The name of the Container registry.
* The URI.  _For now, AWS ECR integrations only support private registries utilizing a standard naming convention and we don't use this information.  However, as a newer feature, we anticipate public ECR registries will be supported in the near future._
* Your AWS Keys.
* Your AWS Region (e.g. us-east-1).

