# Running your Universal Broker client

Run your Broker deployment on your container engine or Kubernetes cluster.

If you are not using broker.snyk.io, target the Broker server for your region by using the command  `-e BROKER_SERVER_URL=https://broker.region.snyk.io \` . For details, see [Broker URLs](../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).

Add the environment variable or variables as defined in your credentials references with the associated values. If references are missing, the connection will not be established, and an error entry will be logged in the Broker client logs.

```
docker run --restart=always 
-p 8000:8000 
-e DEPLOYMENT_ID=<DEPLOYMENTID> 
-e CLIENT_ID=<CLIENTID> 
-e CLIENT_SECRET=<CLIENTSECRET> 
-e PORT=8000 
-e <YOUR_CREDENTIALS_REFERENCE>=<secret value> 
snyk/broker:universal
```

A [Helm chart](https://github.com/snyk/snyk-universal-broker-helm) is available for use on Kubernetes clusters. Refer to the readme for details.

Ensure that you first pull the Helm chart:

`helm pull oci://registry-1.docker.io/snyk/snyk-universal-broker`

Then run:

```
helm install my-snyk-broker oci://registry-1.docker.io/snyk/snyk-universal-broker \
  --set deploymentId='YOUR_DEPLOYMENT_ID' \
  --set clientId='YOUR_CLIENT_ID' \
  --set clientSecret='YOUR_CLIENT_SECRET' \
  --set credentialReferences.MY_GITHUB_TOKEN='YOUR_GITHUB_PAT' \
```
