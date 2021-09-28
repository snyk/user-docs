# Interpret scan results

The `Projects` page will contain an inventory of all projects added and a high level summary of findings. You can expand on a particular project to learn more about vulnerabilities that may have been found and guidance on how to fix these or optimizations. Let's walk through some examples.

We can see that the results of our ACR image scan is identical to the image scan for our Kubernetes deployment. This is expected. We are gradually taking steps to secure our workflow. The `redis` image we pulled from our `manifest` is the same image from our `Dockerfile`. 

![](../../../.gitbook/assets/snyk_scan_08.png)

We will discuss fixing the image in the next section.

