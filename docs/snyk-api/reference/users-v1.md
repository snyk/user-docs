# Users (v1)

{% hint style="info" %}
This document uses the v1 API, which will eventually be deprecated, as further Snyk developments are now focused on the REST API. For more details, see the [v1 API](../v1-api.md).
{% endhint %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/user/{userId}" method="get" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/user/me" method="get" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/user/me/notification-settings/org/{orgId}" method="put" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/user/me/notification-settings/org/{orgId}" method="get" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/user/me/notification-settings/org/{orgId}/project/{projectId}" method="put" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/user/me/notification-settings/org/{orgId}/project/{projectId}" method="get" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}
