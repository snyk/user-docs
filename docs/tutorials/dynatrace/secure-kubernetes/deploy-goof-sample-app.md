# Deploy Goof Sample App

{% hint style="danger" %}
**DO NOT** deploy this sample application to any production environment.
{% endhint %}

For these examples, we are providing a sample vulnerable Node.js application which you can deploy to your lab environment. The app is available in a [public GitHub repository](https://github.com/snyk-partners/k8s-goof) which can be **cloned** to your local development machine.

From a terminal, navigate to the location where you cloned the repository and execute the following commands:

```bash
 kubectl apply -f goof-service.yaml
```

```bash
kubectl apply -f goof-deployment.yaml
```

Once successfully deployed you will see results as follows:

```bash
deployment.apps/goof created
deployment.apps/goof-mongo created
```

If you would like to access the public endpoint for this app you may do so by running the following command and copy and pasting the **EXTERNAL-IP** output into a web browser:

```bash
kubectl get svc
```

This step is not necessary for this exercise. You may proceed to the next section.

