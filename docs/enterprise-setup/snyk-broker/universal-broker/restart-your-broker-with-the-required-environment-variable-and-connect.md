# Restart your Broker with the required environment variable and connect

## Restart your Broker with the required environment variable. <a href="#restart-your-broker-with-required-environment-variable" id="restart-your-broker-with-required-environment-variable"></a>

| <pre><code>docker run --restart=always \
    -p 8000:8000 \
    -e ACCEPT_CODE=true \
    -e DEPLOYMENT_ID=&#x3C;DEPLOYMENTID> \
    -e CLIENT_ID=&#x3C;CLIENTID> \
    -e CLIENT_SECRET=&#x3C;CLIENTSECRET> \
    -e UNIVERSAL_BROKER_ENABLED=true \
    -e MY_GITHUB_TOKEN=GITHUB_TOKEN_VALUE \
    -e PORT=8000 \
    -e BROKER_SERVER_URL=https://broker.pre-prod.snyk.io \
    -e BROKER_HA_MODE_ENABLED=true \
snyk/broker:universal-dev
</code></pre> |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |

At this point, the Broker will display the following message:

| <pre><code>{"name":"my github connection","hostname":"ae7d64e0edac","pid":1,"level":30,"id":"12345678-1234-1234-1234-123456789012","msg":"Connection (my github connection) not in use by any orgs. Will check periodically and create connection when in use.","time":"2024-06-18T20:21:24.382Z","v":0}
</code></pre> |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |

You can now use the connection in an organization integration.

## &#x20;Connect your organization integration to use a connection <a href="#id-5-connect-your-org-integration-to-use-a-connection" id="id-5-connect-your-org-integration-to-use-a-connection"></a>

| <pre><code>curl --location --request POST 'https://api.pre-prod.snyk.io/rest/tenants/TENANT_ID/brokers/connections/CONNECTION_ID/orgs/ORG_ID/integration?version=2024-02-08~experimental' \
--header 'Content-Type: application/vnd.api+json' \
--header 'Authorization: token YOUR_SNYK_TOKEN' \
--data-raw '{
    "data": {
                "integration_id": "INTEGRATION_ID",
                "type": "github"
    }
}'
</code></pre> |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |

Repeat connecting your organization for as many integrations as needed.
