# Snyk IaC-specific CI/CD strategies

The best way to implement Snyk Infrastructure as Code in your pipeline is as part of the stages, but after the SCA and the Containers stage.

Snyk Infrastructure as Code supports:

* Deployments, Pods, and Services
* CronJobs, Jobs, StatefulSet, ReplicaSet, DaemonSet, and ReplicationController

See [Test your Kubernetes files with Snyk CLI](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/test-your-kubernetes-files-with-our-cli-tool) for more details.
