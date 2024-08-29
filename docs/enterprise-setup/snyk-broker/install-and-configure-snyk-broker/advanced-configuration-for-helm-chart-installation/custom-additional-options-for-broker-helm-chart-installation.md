# Custom additional options for Broker Helm Chart installation

If you need to inject additional option(s) using environment variables, use the `override.yaml` value file to add any additional environment variable(s) you may need.

Adding the `--values override.yaml` will load those values into your deployment. For example:

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --values override.yaml \
             -n snyk-broker --create-namespace
```

You can do the same inline without the override.yaml file if it is more convenient.

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --set env[0].name=myEnvVarName \
             --set env[0].value=myEnvVarValue \
             --set env[1].name=myOtherEnvVarName \
             --set env[1].value=myOtherEnvVarValue \
             -n snyk-broker --create-namespace
```

You can add custom Kubernetes resources and objects to the chart by adding them to the values file.

Various combinations of Kubernetes options and objects are available at various levels of the specifications hierarchy, deployment, container, and pod-specific. These are listed in the [default `values.yaml` ](https://github.com/snyk/snyk-broker-helm/blob/a805f97235ba6b004df7a38c93ee94e399b699b7/charts/snyk-broker/values.yaml#L403)file as `extraObjects`, `extraVolumes`, `extraVolumeMounts`, and `extraPodSpecs`.

Be careful to use the right syntax and validate the rendered `yaml` using `helm template` command.
