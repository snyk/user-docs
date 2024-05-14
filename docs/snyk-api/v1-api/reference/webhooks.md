# Webhooks

### Intro

> Warning: the webhooks feature is currently in beta. While in this status, we may change the API and the structure of webhook payloads at any time, without notice.

Webhooks allow you to be notified of events taking place in the Snyk system and react to changes in your projects.

Webhooks associate an event type with a URL. When something triggers that event type, Snyk sends an HTTP POST request to the URL with a payload containing information about the event. Currently supported targets/scan types are Open Source and container.

### Who can access this feature?

Only Business and Enterprise customers.

### Configuring webhooks

Webhooks can be configured using our API at organization level, by organization admins.

{% swagger src="../../../.gitbook/assets/spec.yaml" path="/org/{orgId}/webhooks" method="get" %}
[spec.yaml](../../../.gitbook/assets/spec.yaml)
{% endswagger %}

{% swagger src="../../../.gitbook/assets/spec.yaml" path="/org/{orgId}/webhooks" method="post" %}
[spec.yaml](../../../.gitbook/assets/spec.yaml)
{% endswagger %}

{% swagger src="../../../.gitbook/assets/spec.yaml" path="/org/{orgId}/webhooks/{webhookId}" method="get" %}
[spec.yaml](../../../.gitbook/assets/spec.yaml)
{% endswagger %}

{% swagger src="../../../.gitbook/assets/spec.yaml" path="/org/{orgId}/webhooks/{webhookId}" method="delete" %}
[spec.yaml](../../../.gitbook/assets/spec.yaml)
{% endswagger %}

{% swagger src="../../../.gitbook/assets/spec.yaml" path="/org/{orgId}/webhooks/{webhookId}/ping" method="post" %}
[spec.yaml](../../../.gitbook/assets/spec.yaml)
{% endswagger %}
