# Integrations (v1)

{% hint style="info" %}
This document uses the v1 API, which will eventually be deprecated, as further Snyk developments are now focused on the REST API. For more details, see the [v1 API](../v1-api.md).
{% endhint %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/integrations" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/integrations -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/integrations" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/integrations
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/integrations/{type}" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/integrations/{type}
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/integrations/{integrationId}" method="put" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/integrations/{integrationId} -X PUT \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/integrations/{integrationId}/settings" method="put" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/integrations/{integrationId}/settings -X PUT \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/integrations/{integrationId}/settings" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/integrations/{integrationId}/settings
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/integrations/{integrationId}/clone" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/integrations/{integrationId}/clone -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/integrations/{integrationId}/authentication" method="delete" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/integrations/{integrationId}/authentication -X DELETE
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/integrations/{integrationId}/authentication/switch-token" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/integrations/{integrationId}/authentication/switch-token -X POST
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/integrations/{integrationId}/authentication/provision-token" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/integrations/{integrationId}/authentication/provision-token -X POST
```
