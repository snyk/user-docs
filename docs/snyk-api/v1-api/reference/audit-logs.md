# Audit logs

Get audit logs of your group or organization. Logs are only available for past 3 months. Note that the API returns personally identifiable information and requires the use of either a personal Snyk API token or a Snyk service account token with Group Admin level permission.

{% swagger src="../../../.gitbook/assets/spec.yaml" path="/group/{groupId}/audit" method="post" %}
[spec.yaml](../../../.gitbook/assets/spec.yaml)
{% endswagger %}

{% swagger src="../../../.gitbook/assets/spec.yaml" path="/org/{orgId}/audit" method="post" %}
[spec.yaml](../../../.gitbook/assets/spec.yaml)
{% endswagger %}
