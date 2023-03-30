# Ingress options with Snyk Broker Helm installation

There are two options available for ingress traffic. By default, the pods are not accessible from outside the cluster.

To enable a **load balancer**, add the `--set service.<service-type>=LoadBalancer`. Allowed values are `brokertype`, `crType`, and `caType`

Example for Github:

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --set service.type=LoadBalancer \
             -n snyk-broker --create-namespace
```
