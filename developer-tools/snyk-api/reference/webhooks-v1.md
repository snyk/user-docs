# Webhooks (v1)

{% hint style="info" %}
This document uses the v1 API, which will eventually be deprecated, as further Snyk developments are now focused on the REST API. For more details, see the [v1 API](../v1-api.md).
{% endhint %}

{% hint style="warning" %}
The Webhooks API is in beta. While the API is in beta, Snyk may change the API and the structure of webhook payloads at any time, without notice. During this beta, Webhooks are available only in the Snyk US-01, US-02, EU-01, and AU-01 regions.
{% endhint %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/webhooks" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/webhooks" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/webhooks/{webhookId}" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/webhooks/{webhookId}" method="delete" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/webhooks/{webhookId}/ping" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}
