# Ensure deployment manifests are free from misconfiguration

In the previous sections we fixed vulnerabilities introduced by the container base image and the application dependencies, and tested our changes by deploying to Kubernetes.

Even with 0 vulnerabilities, we can expose ourselves to risk by how our application's deployment manifests are written. In this section we'll use Snyk to find and fix issues in those files.

## Investigate issues in deployment manifests

In the Snyk project imported earlier, we can see the two Kubernetes manifests for the goof application. One deploys the application, and the other deploys the mongo database it needs.

![](../../../.gitbook/assets/snyk-iac-dockerlab.png)

Clicking into either one of these brings you to the configuration risks Snyk identified in the file. Starting with the `goof-mongo-deployment.yaml` file, we see the following issues:

![](../../../.gitbook/assets/goof-mongo-issues.png)

For each issue, Snyk calls out the issue identified, its impact, and how it can be resolved. It also highlights the line of code where the issue exists. In the example below, we see a Low Severity configuration risk that can be addressed by adding an `ImagePullPolicy` to the deployment.

![](../../../.gitbook/assets/iac-pullpolicyissue.png)

With this information, developers can either ignore the issue or make the necessary changes to their deployment manifests to fix them.

{% hint style="danger" %}
These are examples. It's possible your application requires administrative privileges or other explicitly stated privileges. Know your application and its dependencies before  making changes to your files; you can break your deployment if you're not careful.
{% endhint %}

## Fixing configuration issues in deployment manifests

As the warning states, configuration issues are very nuanced. While there are documented best practices we check against, your workloads might require those in order to function correctly. For this reason, we allow you to [change the severities for the configuration checks](https://support.snyk.io/hc/en-us/articles/360006402818#UUID-c1919782-6bfa-b84b-a638-3913cee39fc5) Snyk runs.

In this example, we'll fix a few issues we deem important enough to warrant a fix. 

### Issue: Container could be running with outdated image

Let's look once again at the issue from above.

![](../../../.gitbook/assets/iac-pullpolicyissue.png)

Reading the issue description, this issue can be fixed by adding an `ImagePullPolicy` attribute to the `containers` spec. In your IDE, or text editor, add the attribute to `goof-mongo-deployment.yaml`.

```yaml
spec:
  containers:
    - image: mongo
      name: goof-mongo
      imagePullPolicy: Always
      ports:
```

### Issue: Container is running without root user control

Let's look at another issue. This time, it has to do with the permissions given to the container.

![](../../../.gitbook/assets/iac-runasnoonroot.png)

This one can be fixed by setting a `securityContext` on the `container`. To set this, add the following to the `goof-deployment-mongo.yaml` file.

```yaml
spec:
  containers:
    - image: mongo
      name: goof-mongo
      imagePullPolicy: Always
      securityContext:
        runAsNonRoot: true
```

## Test your config files to ensure they work

Any time you make a local change to a configuration file, it's a good idea to test it locally before committing the changes to GitHub. Tear down the current deployment by deleting the namespace:

```yaml
kubectl delete ns snyk-docker
```

Once it's gone, re-deploy the application using the following commands.

```yaml
# Create a namespace
kubectl create ns snyk-docker

# Set the current context to use the new namespace
kubectl config set-context --current --namespace snyk-docker

# Spin up the goof deployment and service
kubectl create -f goof-deployment.yaml,goof-mongo-deployment.yaml
```

If it spins up correctly, test the app by navigating to [http://localhost:3001](http://localhost:3001). Once you've confirmed it works, commit the changes to GitHub! Snyk will re-test the files and show updated issue counts.

You reached the end! We hope you enjoyed it. Check out [Recap and Next Steps](recap-and-next-steps.md) for extra resources.

