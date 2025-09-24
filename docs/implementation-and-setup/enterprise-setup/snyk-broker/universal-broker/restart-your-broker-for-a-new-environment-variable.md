# Restart your Broker for a new environment variable

If you change an environment variable, you must restart your Broker, except in Kubernetes deployments.

You must use `-e BROKER_SERVER_URL=https://broker.REGION.snyk.io \` for regional Snyk instances. For details, see [Broker URLs](../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).

<table data-header-hidden><thead><tr><th></th></tr></thead><tbody><tr><td><pre><code>docker run --restart=always \
    -p 8000:8000 \
    -e ACCEPT_CODE=true \
    -e DEPLOYMENT_ID=&#x3C;DEPLOYMENTID> \
    -e CLIENT_ID=&#x3C;CLIENTID> \
    -e CLIENT_SECRET=&#x3C;CLIENTSECRET> \
    -e MY_GITHUB_TOKEN=GITHUB_TOKEN_VALUE \
    -e PORT=8000 \
snyk/broker:universal
</code></pre></td></tr></tbody></table>

At this point, the Broker will display a message like the following:

<table data-header-hidden><thead><tr><th></th></tr></thead><tbody><tr><td><pre><code>{"name":"my github connection","hostname":"ae7d64e0edac","pid":1,"level":30,"id":"12345678-1234-1234-1234-123456789012","msg":"Connection (my github connection) not in use by any orgs. Will check periodically and create connection when in use.","time":"2024-06-18T20:21:24.382Z","v":0}
</code></pre></td></tr></tbody></table>

In Kubernetes deployments with hot-loaded secrets, when you edit a secret, typically using a vault system or something similar, the values in the secrets are automatically updated into the mounted secret file. This allows the Broker to trigger a reloading while running, hot-loading the new value without needing to restart the container.
