# Users (v1)

{% hint style="info" %}
This document uses the v1 API, which will eventually be deprecated, as further Snyk developments are now focused on the REST API. For more details, see the [v1 API](../v1-api.md).
{% endhint %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/user/{userId}" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/user/{userId}
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/user/me" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/user/me
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/user/me/notification-settings/org/{orgId}" method="put" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/user/me/notification-settings/org/{orgId} -X PUT \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/user/me/notification-settings/org/{orgId}" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/user/me/notification-settings/org/{orgId}
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/user/me/notification-settings/org/{orgId}/project/{projectId}" method="put" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/user/me/notification-settings/org/{orgId}/project/{projectId} -X PUT \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/user/me/notification-settings/org/{orgId}/project/{projectId}" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/user/me/notification-settings/org/{orgId}/project/{projectId}
```
