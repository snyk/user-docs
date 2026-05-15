# Snyk IaC-specific CI/CD strategies

The best way to implement Snyk Infrastructure as Code in your pipeline is as part of the stages, but after the SCA and the Containers stage.

Snyk Infrastructure as Code supports:

* Deployments, Pods, and Services
* CronJobs, Jobs, StatefulSet, ReplicaSet, DaemonSet, and ReplicationController

See [Kubernetes files scanning with CLI for IaC ](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/test-your-iac-files/kubernetes-files.md)for more details.
