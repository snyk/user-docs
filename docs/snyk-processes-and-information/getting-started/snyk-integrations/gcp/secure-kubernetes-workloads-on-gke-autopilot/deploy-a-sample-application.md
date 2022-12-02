# Deploy a sample application

{% hint style="danger" %}
**DO NOT** deploy the following sample application in a production environment. This application is used for demo purposes and contains a number of vulnerabilities. It is **highly recommended** that you promptly uninstall the application after you've completed the tutorial.
{% endhint %}

Snyk provides a demo application called **goof** which may be used for this tutorial. You may deploy any application you desire. To deploy the **goof** app, clone the repository to your local environment with the following command:

```bash
git clone https://github.com/snyk-partners/k8s-goof.git
```

Next, change directory or `cd` to where you cloned as follows:

```bash
cd k8s-goof/
```

This will bring us to where we have conveniently placed some Kubernetes manifests to help with your deployments. Run the following commands:

```bash
kubectl create namespace goof && \
kubectl apply -f ./manifests -n goof
```

It may take a few minutes for your application to be ready, but you can check the status with the following command:

```bash
kubectl get pods -n goof
```

When these display that they are in a **ready** state `1/1` then you should be able to pull the external IP and past that into a browser window to access the application. You can get the external IP for the application with the following command:

```bash
kubectl get svc -n goof
```
