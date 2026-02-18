# Using the API to set up a GitHub connection

This page provides an example of using the API to set up a GitHub connection with the Universal Broker. Repeat connecting your Organization for as many integrations as needed.

Using the `snyk-broker-config` CLI tool is recommended for an easier experience. The API allows for automation and more control, and also requires a clear understanding of Broker deployments, credentials, connections, and integrations.

{% hint style="info" %}
In any of the calls that follow, replace `api.snyk.io` with your regional equivalent if necessary, for example, `api.eu.snyk.io`. For a list of URLs, see [API URLs](../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#api-urls).
{% endhint %}

### Prerequisites

* You must be a **Snyk Tenant admin**.
* You must use a **personal Snyk API token**. Service account tokens do not work.
* You need **Org admin** access for the Organization where you install the Broker app.

{% hint style="info" %}
All examples on this page use REST API version `2025-11-05`.
{% endhint %}

## Install Broker for your Organization <a href="#id-1-install-the-broker-app-on-your-org" id="id-1-install-the-broker-app-on-your-org"></a>

Universal Broker facilitates the secure connection and communication with the Broker server through OAuth.

Install the Broker app at the Organization level. Group-level installation is not supported.

1. Identify the **Snyk Broker App ID** for your region.
   * For most regions, the app ID is `cb43d761-bd17-4b44-9b6c-e5b8ad077d33`.
   * If you are using Snyk for Government, contact your account team for the app ID.
2. Install the Broker on the Organization you use to administer Broker:

```bash
curl --location --request POST 'https://api.snyk.io/rest/orgs/ORG_ID/apps/installs?version=2025-11-05' \
--header 'Content-Type: application/vnd.api+json' \
--header 'Authorization: token YOUR_SNYK_TOKEN' \
--data-raw '{
  "data": {
    "type": "app_install",
    "relationships": {
      "app": {
        "data": {
          "id": "SNYK_BROKER_APP_ID",
          "type": "app"
        }
      }
    }
  }
}'
```

Use the endpoint [Install a Snyk App to this Organization](../../../../snyk-api/reference/apps.md#post-orgs-org_id-apps-installs) as the source of truth for the exact request and response schema.

Save the values needed in later steps:

* `INSTALL_ID` is the `data.id` from the response.
* The **Broker App ID** is shown under `data.relationships.app.data.id`.
* `CLIENT_ID` and `CLIENT_SECRET` (returned by the install call) are needed to run the Broker client.

{% hint style="info" %}
The installation displays the `CLIENT_SECRET` only one time. You can rotate the secret using an API call.
{% endhint %}

## Create your deployment <a href="#id-2-create-your-deployment" id="id-2-create-your-deployment"></a>

Use the following call to create your deployment.

```
curl --location --request POST 'https://api.snyk.io/rest/tenants/TENANT_ID/brokers/installs/INSTALL_ID/deployments?version=2025-11-05' \
--header 'Content-Type: application/vnd.api+json' \
--header 'Authorization: token YOUR_SNYK_TOKEN' \
--data-raw '{
    "data": {
        "type": "broker_deployment",
        "attributes": {
            "broker_app_installed_in_org_id":"ORG_ID_WHERE_APP_WAS_INSTALLED",
            "metadata": {
                "deployment_name": "My Universal Broker Deployment",
                "cluster": "Cluster X Region Y or whatever you need to not lose your deployment."
            }
        }
    }
}'
```

This returns the DEPLOYMENT\_ID (data.id), for example:

<table data-header-hidden><thead><tr><th></th></tr></thead><tbody><tr><td><pre><code>{
    ...
    "data": {
        "id": "12345678-1234-1234-1234-123456789012",
        "type": "broker_deployment",
        "attributes": {
            "install_id": "12345678-1234-1234-1234-123456789012",
            "metadata": {
                "deployment_name": "My Universal Broker Deployment",
                "cluster": "Cluster X Region Y or whatever you need to not lose your deployment."
            }
        }
    },
    ...
}
</code></pre></td></tr></tbody></table>

At this point, you can start running the Broker client.

## Run your Broker deployment <a href="#run-your-broker-deployment_1" id="run-your-broker-deployment_1"></a>

Target your desired environment with the usual `-e BROKER_SERVER_URL=https://broker.REGION.snyk.io \` if needed. For a list of URLs, see [Broker server URLs](../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).

Use the `CLIENT_ID` and `CLIENT_SECRET` values returned when you installed the Broker app.

<table data-header-hidden><thead><tr><th></th></tr></thead><tbody><tr><td><pre><code>docker run --restart=always \
    -p 8000:8000 \
    -e ACCEPT_CODE=true \
    -e DEPLOYMENT_ID=&#x3C;DEPLOYMENTID> \
    -e CLIENT_ID=&#x3C;CLIENTID> \
    -e CLIENT_SECRET=&#x3C;CLIENTSECRET> \
    -e UNIVERSAL_BROKER_ENABLED=true \
    -e PORT=8000 \
    -e BROKER_HA_MODE_ENABLED=true \
snyk/broker:universal
</code></pre></td></tr></tbody></table>

When the command is running, you should get the following message in the output:

<table data-header-hidden><thead><tr><th></th></tr></thead><tbody><tr><td><pre><code>{"name":"snyk-broker","hostname":"c918a4e1535a","pid":1,"level":30,"msg":"Found deployment 12345678-1234-1234-1234-123456789012. Waiting for connections. (polling)","time":"2024-06-18T20:15:20.572Z","v":0}
</code></pre></td></tr></tbody></table>

## Create your credentials reference or references <a href="#id-3-create-your-credentials-references" id="id-3-create-your-credentials-references"></a>

<table data-header-hidden><thead><tr><th></th></tr></thead><tbody><tr><td><pre><code>curl --location --request POST 'https://api.snyk.io/rest/tenants/TENANT_ID/brokers/installs/INSTALL_ID/deployments/DEPLOYMENT_ID/credentials?version=2025-11-05' \
--header 'Content-Type: application/vnd.api+json' \
--header 'Authorization: token YOUR_SNYK_TOKEN' \
--data-raw '{
    "data": {
        "type": "deployment_credential",
        "attributes": [{
            "comment": "This is github token service account XYZ (example)",
            "environment_variable_name": "MY_GITHUB_TOKEN",
            "type": "github"
        }]
    }
}'
</code></pre></td></tr></tbody></table>

This returns the credential reference ID (data. id), for example:

<table data-header-hidden><thead><tr><th></th></tr></thead><tbody><tr><td><pre><code>{
    ...
    "data": {
        "id": "12345678-1234-1234-1234-123456789012",
        "type": "deployment_credential",
        "attributes": {
            "comment": "test cred",
            "deployment_id": "12345678-1234-1234-1234-123456789012",
            "environment_variable_name": "MY_GITHUB_TOKEN",
            "type": "github"
        }
    },
    ...
}
</code></pre></td></tr></tbody></table>

You can create a maximum of ten credentials in one call by adding more objects in attributes. These objects can be of different types.

<table data-header-hidden><thead><tr><th></th></tr></thead><tbody><tr><td><pre><code>...
"attributes": [{
            "comment": "This is github token service account XYZ (example)",
            "environment_variable_name": "MY_GITHUB_TOKEN",
            "type": "github"
        },
        {
            "comment": "This is my other github token service account ABC (example)",
            "environment_variable_name": "MY_OTHER_GITHUB_TOKEN",
            "type": "github"
        }]
...
</code></pre></td></tr></tbody></table>

## Create your connection or connections <a href="#id-4-create-your-connections" id="id-4-create-your-connections"></a>

<table data-header-hidden><thead><tr><th></th></tr></thead><tbody><tr><td><pre><code>curl --location --request POST 'https://api.snyk.io/rest/tenants/TENANT_ID/brokers/installs/INSTALL_ID/deployments/DEPLOYMENT_ID/connections?version=2025-11-05' \
--header 'Content-Type: application/vnd.api+json' \
--header 'Authorization: token YOUR_SNYK_TOKEN' \
--data-raw '{
    "data": {
        "type": "broker_connection",
        "attributes": {
            "name": "my github connection",
            "configuration": {
                "required": {
                    "github_token": "CRED_REFERENCE_UUID",
                    "broker_client_url":"http://your.broker.hostname:port"
                },
                "type": "github"
            },
            "deployment_id": "DEPLOYMENT_ID"
        }
    }
}'
</code></pre></td></tr></tbody></table>

This returns a Connection ID (data.id), as shown in the example. This call returns the credential reference directly, ready for use, rather than the credential reference UUID.

<table data-header-hidden><thead><tr><th></th></tr></thead><tbody><tr><td><pre><code>{
    ...
    "data": {
        "id": "12345678-1234-1234-1234-123456789012",
        "type": "broker_connection",
        "attributes": {
            "deployment_id": "12345678-1234-1234-1234-123456789012",
            "name": "my github connection",
            "secrets": {
                "primary": {
                    "encrypted": "not-yet-implemented",
                    "expires_at": "1970-01-01T00:00:00.000Z",
                    "nonce": "not-yet-implemented"
                },
                "secondary": {
                    "encrypted": "not-yet-implemented",
                    "expires_at": "1970-01-01T00:00:00.000Z",
                    "nonce": "not-yet-implemented"
                }
            },
            "configuration": {
                "required": {
                    "broker_client_url": "dummy",
                    "github_token": "${MY_GITHUB_TOKEN}"
                },
                "type": "github"
            }
        }
    },
    ...
}
</code></pre></td></tr></tbody></table>

If your credential reference is missing, you get the following message:

<table data-header-hidden><thead><tr><th></th></tr></thead><tbody><tr><td><pre><code>{"name":"snyk-broker","hostname":"029cda64bd98","pid":1,"level":50,"connection":"my github connection","msg":"Connection is missing environment variable value MY_GITHUB_TOKEN. Connection is disabled till value is provided. Restart broker once added.","time":"2024-06-18T14:29:06.910Z","v":0}
{"name":"my github connection","hostname":"029cda64bd98","pid":1,"level":50,"id":"12345678-1234-1234-1234-123456789012","msg":"Connection is disabled due to (a) missing environment variable(s). Please provide the value and restart the broker client.","time":"2024-06-18T14:29:06.911Z","v":0}
</code></pre></td></tr></tbody></table>

When you update an environment variable, you must restart your Broker. See [Restart your Broker for a new environment variable](../../../../implementation-and-setup/enterprise-setup/snyk-broker/universal-broker/restart-your-broker-for-a-new-environment-variable.md).

You can now use the connection in an Organization integration.

## Connect your Organization integration to use a connection <a href="#id-5-connect-your-org-integration-to-use-a-connection" id="id-5-connect-your-org-integration-to-use-a-connection"></a>

<table data-header-hidden><thead><tr><th></th></tr></thead><tbody><tr><td><pre><code>curl --location --request POST 'https://api.snyk.io/rest/tenants/TENANT_ID/brokers/connections/CONNECTION_ID/orgs/ORG_ID/integration?version=2025-11-05' \
--header 'Content-Type: application/vnd.api+json' \
--header 'Authorization: token YOUR_SNYK_TOKEN' \
--data-raw '{
    "data": {
                "integration_id": "INTEGRATION_ID",
                "type": "github"
    }
}'
</code></pre></td></tr></tbody></table>

Repeat the process of connecting your Organization for as many integrations as needed.
