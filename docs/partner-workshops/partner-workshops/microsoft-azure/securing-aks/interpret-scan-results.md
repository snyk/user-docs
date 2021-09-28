# Interpret scan results

The `Projects` page will contain an inventory of all projects added and a high level summary of findings. You can expand on a particular project to learn more about vulnerabilities that may have been found and guidance on how to fix these or optimizations. Let's walk through some examples.

The figure below shows the two workloads we imported in the previous section. These have been expanded to show that the results are including two key areas: Kubernetes configuration and container image scan results. Clicking on each of these will present you with additional insights.

![](../../../.gitbook/assets/snyk_scan_01.png)

Let's begin with examining our Kubernetes configuration by click on each respective workload.

![](../../../.gitbook/assets/snyk_scan_02.png)

![](../../../.gitbook/assets/snyk_scan_03.png)

The results above yield some interesting findings. From this view, we are able to see a summary of `Vulnerabilities`, number of `Dependencies`, and our `Security configuration`. Let's take a closer look and interpret what these results mean.

If we examine the `azure-vote.yaml` manifest we applied to our Kubernetes cluster, we can see that we defined some parameters such as `cpu` and `memory` limits. As a result, these were not flagged during the scan.

```yaml
resources:
  requests:
    cpu: 100m
    memory: 128Mi
  limits:
    cpu: 250m
    memory: 256Mi
```

However, our manifest did not set `securityContext` parameters such as `readOnlyRootFilesystem`, `runAsNonRoot`, `allowPrivilegeEscalation`, and `capabilities`. As a result, we see this in our findings with the `FAIL` flag.

* Run as non-root: Whether any containers in the workload have `container.securityContext.runAsNonRoot` set to `false` or unset.
* Read-only root file system: Whether any containers in the workload have `container.securityContext.readOnlyFilesystem` set to `false` or unset.
* Drop capabilities: Whether all capabilities are dropped and `CAP_SYS_ADMIN` is not added.

Next, let's take a closer look at our container image results:

![](../../../.gitbook/assets/snyk_scan_04.png)

Here we are immediately advised that a `Dockerfile` is missing. We know this to be true, because our deployment consisted of a manifest file that is pulling an image from a public registry. In the case of our `vote-back` and `vote-front` applications we are pulling the [image](https://kubernetes.io/docs/concepts/containers/images/) from `redis` and `microsoft/azure-vote-front:v1` respectively.

A best practice here would be to include a `Dockerfile` in our Git repo and pull the image from our private registry. That way we can resolve any issues with our base images and also scan & monitor those through Snyk's integrations to [Azure Repos](https://support.snyk.io/hc/en-us/articles/360004002198-Azure-Repos-integration) and [Azure Container Registry \(ACR\)](https://support.snyk.io/hc/en-us/articles/360003946957-Container-security-with-ACR-integrate-and-test).

We will dive into these in a later workshop.

