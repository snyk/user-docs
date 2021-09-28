# Fix Configuration

From the Snyk app, we will go to the **Projects** menu where we can expand each integration and have a holistic view of our project. Here we will select the _**Kubernetes deployment**_ under our Amazon EKS cluster.

![](../../../../.gitbook/assets/snyk-projects-03.png)

Here we see that our application was deployed to our cluster, however, several _security context_ properties were either misconfigured or not configured at all.

![](../../../../.gitbook/assets/snyk-eks-integration-03.png)

If we navigate back to our imported Bitbucket project and review the `goof-deployment-template.yaml` file results, we are able to see additional context about these issues.

![](../../../../.gitbook/assets/snyk-k8s-report.png)

To fix these, we simply add `securityContext` to our manifest and explicitly define these properties. We will copy the following block into our manifest file.

```yaml
          securityContext:
            allowPrivilegeEscalation: false
            readOnlyRootFilesystem: true
            runAsNonRoot: true
            capabilities:
              drop:
                - all
```

Let's once again, navigate to our Bitbucket repository and edit `./deployment/goof-deployment-template.yaml` with Bitbucket's built-in editor so that it looks as follows:

![](../../../../.gitbook/assets/bitbucket-edit-manifest.png)

Click **Commit**.

