# Restart your Broker for a new environment variable

If you change an environment variable, you must restart your Broker.

You must use `-e BROKER_SERVER_URL=https://broker.REGION.snyk.io \` for regional Snyk instances. For details, see [Broker URLs](../../../../working-with-snyk/regional-hosting-and-data-residency.md#broker-urls).

| <pre><code>docker run --restart=always \
    -p 8000:8000 \
    -e ACCEPT_CODE=true \
    -e DEPLOYMENT_ID=&#x3C;DEPLOYMENTID> \
    -e CLIENT_ID=&#x3C;CLIENTID> \
    -e CLIENT_SECRET=&#x3C;CLIENTSECRET> \
    -e MY_GITHUB_TOKEN=GITHUB_TOKEN_VALUE \
    -e PORT=8000 \
snyk/broker:universal
</code></pre> |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |

At this point, the Broker will display a message like the following:

| <pre><code>{"name":"my github connection","hostname":"ae7d64e0edac","pid":1,"level":30,"id":"12345678-1234-1234-1234-123456789012","msg":"Connection (my github connection) not in use by any orgs. Will check periodically and create connection when in use.","time":"2024-06-18T20:21:24.382Z","v":0}
</code></pre> |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |

