# Pushing a bundle

Optionally, once you have generated your custom rules bundle, you can also distribute automatically to one of our supported container registries by using the `push` command:

```
snyk-iac-rules push -r docker.io/example/test bundle.tar.gz
```

Amongst our supported registries we have [DockerHub](https://hub.docker.com), [Azure Container Registry](https://azure.microsoft.com/en-us/services/container-registry/), and more to come.

This will push your bundle with the `latest` tag, but you can also provide your own tagging to version the bundle.

```
snyk-iac-rules push -r docker.io/example/test:v0.0.1 bundle.tar.gz
```

&#x20;You can now [run snyk iac test with your newly built custom bundle. ](../how-to-run-custom-rules-with-the-snyk-cli.md)
