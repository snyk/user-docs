# Organizations (v1)

{% hint style="info" %}
This document uses the v1 API, which will eventually be deprecated, as further Snyk developments are now focused on the REST API. For more details, see the [v1 API](../v1-api.md).
{% endhint %}

{% hint style="info" %}
Code Consistent Ignores, used in org-level policies, are only available in the REST API.
{% endhint %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/orgs" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/orgs
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}" method="delete" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId} -X DELETE
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/settings" method="put" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/settings -X PUT \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/settings" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/settings
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/provision" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/provision -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/provision" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/provision
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/provision" method="delete" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/provision -X DELETE
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/notification-settings" method="put" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/notification-settings -X PUT \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/notification-settings" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/notification-settings
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/members" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/members
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/members/{userId}" method="put" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/members/{userId} -X PUT \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/members/{userId}" method="delete" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/members/{userId} -X DELETE
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/members/update/{userId}" method="put" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/members/update/{userId} -X PUT \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/invite" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/invite -X POST \
  --input body.json
```
