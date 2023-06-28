# Ingress options with Snyk Broker Helm installation

When you are setting up the Broker using Helm, you may need to configure the `brokerClientUrl` parameter. This parameter enables PR Checks if you are connecting to an SCM and enables connecting to Container Registries.

For this connection to happen, there must be inbound connectivity from the SCM or Container Registry to the Broker. Kubernetes manages this inbound connectivity through ingress.

Ingress is a way to route incoming network traffic to specific services in a Kubernetes cluster. For more information on how to set up ingress see the [Kubernetes guide to ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/).

There are two options available for ingress traffic. By default, the pods are not accessible from outside the cluster.

To enable a **load balancer**, add the `--set service.<service-type>=LoadBalancer`. Allowed values are `brokertype`, `crType`, and `caType`. Servicet ype refers to the type of Broker you are running.

Example for Github:

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --set service.brokerType=LoadBalancer \
             -n snyk-broker --create-namespace
```

To add an **ingress**, enable it in the `values.yaml` file and add the relevant configuration parameters, following the example values in the values file. For this to work, your cluster must have an ingress controller configured properly.
