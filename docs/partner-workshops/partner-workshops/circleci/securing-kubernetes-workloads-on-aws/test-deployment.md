# Test deployment

Now that our job ran successfully, let's validate our application is running. To do this, we will run `kubectl get svc` once more.

![](../../../.gitbook/assets/kubectl_get_svc_external-ip.gif)

This time, we should see _**two**_ ****additional services: **goof** and **goof-mongo** which correspond to the deployment and services configurations from our Kubernetes manifest. Recall that in our `./deployment/goof-service.yaml` we defined this as our `frontend` app with type `LoadBalancer` listening on the standard `http` port `80`.

#### Install jq

You can just as well copy and paste results into your clipboard, but we will take a different approach and use some command-line magic to achieve better results. Let's start by installing a command-line JSON processor [`jq`](https://formulae.brew.sh/formula/jq) using Homebrew.

```bash
brew update && brew install jq
```

Next let's create an environment variable `GOOF_HOST` and pass the value of `EXTERNAL-IP` for our **goof** service in one command.

```bash
GOOF_HOST=$(kubectl get svc goof -o json | \
jq -r '.status.loadBalancer.ingress[].hostname')
```

Let's make sure it worked and type `echo $GOOF_HOST` on the terminal command-line and see the results. We should see the a return value that matches our `EXTERNAL-IP`. Next, we are going to perform a simple test to see if our application is serving requests. For this, we will run a simple `curl` command and browse `./public/about.html`. Your results should look similar to the following:

![](../../../.gitbook/assets/goof_curl_about.gif)

If you prefer, you can also launch your favorite web browser and navigate to the `EXTERNAL-IP` of our `LoadBalancer` as shown below:

![](../../../.gitbook/assets/circleci_goof.png)



