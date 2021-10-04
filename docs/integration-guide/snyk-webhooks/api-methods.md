# API Methods

{% hint style="warning" %}
**Official** Snyk API is available at [https://snyk.docs.apiary.io/\#](https://snyk.docs.apiary.io/#)
{% endhint %}

{% api-method method="post" host="https://snyk.io/api/v1/org/:orgId" path="/webhooks" %}
{% api-method-summary %}
Create a webhook
{% endapi-method-summary %}

{% api-method-description %}
This endpoint allows you to create a webhook.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="orgId" type="string" required=true %}
Organization ID that uniquely identifies your organization
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
token API\_KEY
{% endapi-method-parameter %}

{% api-method-parameter name="Content-Type" type="string" required=true %}
application/json; charset=utf-8
{% endapi-method-parameter %}
{% endapi-method-headers %}

{% api-method-body-parameters %}
{% api-method-parameter name="url" type="string" required=true %}
Webhook handler endpoint
{% endapi-method-parameter %}

{% api-method-parameter name="secret" type="string" required=true %}
Webhook handler token
{% endapi-method-parameter %}
{% endapi-method-body-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Webhook successfully created.
{% endapi-method-response-example-description %}

```javascript
{
    "id": ​"d3cf26b3-2d77-497b-bce2-23b33cc15362"​,
    "url": ​"https://my.app.com/webhook-handler/snyk123"​
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://snyk.io/api/v1/org/:orgId" path="/webhooks" %}
{% api-method-summary %}
List webhooks
{% endapi-method-summary %}

{% api-method-description %}
This endpoint allows you to get a list of all webhooks
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="orgId" type="string" required=true %}
Organization ID that uniquely identifies your organization
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
token API\_KEY
{% endapi-method-parameter %}

{% api-method-parameter name="Content-Type" type="string" required=true %}
application/json; charset=utf-8
{% endapi-method-parameter %}
{% endapi-method-headers %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
{
    "results": [
        {
            "id": ​"d3cf26b3-2d77-497b-bce2-23b33cc15362"​, 
            "url": ​"https://my.app.com/webhook-handler/snyk123",​
        },
    ],
    "total": ​1
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://snyk.io/api/v1/org/:orgId" path="/webhooks/:webhookId" %}
{% api-method-summary %}
Retrieve a webhook
{% endapi-method-summary %}

{% api-method-description %}
This endpoint allows you to get details for a specific webhook
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="webhookId" type="string" required=true %}
Snyk webhook ID
{% endapi-method-parameter %}

{% api-method-parameter name="orgId" type="string" required=true %}
Organization ID that uniquely identifies your organization
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
token API\_KEY
{% endapi-method-parameter %}

{% api-method-parameter name="Content-Type" type="string" required=true %}
application/json; charset=utf-8
{% endapi-method-parameter %}
{% endapi-method-headers %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
{
    "id": ​"d3cf26b3-2d77-497b-bce2-23b33cc15362"​,
    "url": ​"https://my.app.com/webhook-handler/snyk123"​
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="post" host="https://snyk.io/api/v1/org/:orgId" path="/webhooks/:webhookId/ping" %}
{% api-method-summary %}
Ping a webhook
{% endapi-method-summary %}

{% api-method-description %}

{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="webhookId" type="string" required=true %}
Snyk webhook ID
{% endapi-method-parameter %}

{% api-method-parameter name="orgId" type="string" required=true %}
Organization ID that uniquely identifies your organization
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
token API\_KEY
{% endapi-method-parameter %}

{% api-method-parameter name="Content-Type" type="string" required=true %}
application/json; charset=utf-8
{% endapi-method-parameter %}
{% endapi-method-headers %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```
Ok
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="delete" host="https://snyk.io/api/v1/org/:orgId" path="/webhooks/:webhookId" %}
{% api-method-summary %}
Delete a webhook
{% endapi-method-summary %}

{% api-method-description %}

{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="webhookId" type="string" required=true %}
Snyk webhook ID
{% endapi-method-parameter %}

{% api-method-parameter name="orgId" type="string" required=true %}
Organization ID that uniquely identifies your organization
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
token API\_KEY
{% endapi-method-parameter %}

{% api-method-parameter name="Content-Type" type="string" required=true %}
application/json; charset=utf-8
{% endapi-method-parameter %}
{% endapi-method-headers %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```
Ok
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

