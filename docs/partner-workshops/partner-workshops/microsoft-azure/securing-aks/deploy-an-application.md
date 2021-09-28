# Deploy an application

### Kubernetes manifest

We will deploy our sample application using a [Kubernetes manifest](https://docs.microsoft.com/en-us/azure/aks/concepts-clusters-workloads#deployments-and-yaml-manifests) file. A sample file named [`azure-vote.yaml`](https://github.com/snyk-partners/snyk-azure-resources/blob/master/templates/azure-vote.yaml) is provided for your convenience in the repository. The manifest includes two Kubernetes deployments: one for a sample Azure Vote Python application and the other for a [Redis](https://redislabs.com/) instance. Two Kubernetes [services](https://docs.microsoft.com/en-us/azure/aks/concepts-network#services) are also created: one is an internal service for the Redis instance and the other is an external service to allow access to the application from the internet.

If you have cloned the repository and are working from the local directory, run the following command to deploy the application from the `YAML` manifest:

```bash
kubectl apply -f kube-manifests/azure-vote.yaml
```

If successful, you should see output similar to the following:

```text
deployment "azure-vote-back" created
service "azure-vote-back" created
deployment "azure-vote-front" created
service "azure-vote-front" created
```

### Test the application

We will invoke the [kubectl get service](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#get) from our CLI with the `--watch` argument to monitor the application deployment and obtain the `EXTERNAL-IP` of the `LoadBalancer`.

Run the following command:

```bash
kubectl get service azure-vote-front --watch
```

While the application is deploying you may see output similar to the following:

```text
NAME               TYPE           CLUSTER-IP   EXTERNAL-IP   PORT(S)        AGE
azure-vote-front   LoadBalancer   10.0.37.27   <pending>     80:30572/TCP   6s
```

Note that `EXTERNAL-IP` displays a status of `pending`. Wait until this displays a valid public IP address then copy and paste this value into your web browser.

Your browser should resolve and display the following:

![](../../../.gitbook/assets/azure_voting_app.png)

## 

