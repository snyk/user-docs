# Set up the Snyk Webhook

You will set up the Snyk Webhook using the [Snyk API v1](https://snyk.docs.apiary.io/#reference/webhooks/webhook-collection/create-a-webhook) including the built-in console.

To set up the Snyk Webhook, follow these steps:

1.  Copy your **Organization ID** from the **Organization** settings in the Snyk Web UI.\


    <figure><img src="https://lh3.googleusercontent.com/n5_nk9_s2lIb982FQV8LwIQzgYWxC6xeDKEiZMnN_TvrAuDV5oWvCR2RO15XMzyhvVpQwpg1IcL97ljvhis1Q3hfynm91EEqRQvaA7mdkeholt_JvmKPeq1eVmgmnQu5Iaahmdl4UC_8oPP4A6kSGUBO7iz0YPrBca4hbhXOLndO_DLK0NkPPK4dmQ" alt="Snyk Web UI copy Organization ID"><figcaption><p>Snyk Web UI copy Organization ID</p></figcaption></figure>
2.  Get your organization admin **API Token** from the Snyk Web UI, either from a Service Account or your own account.\


    <figure><img src="https://lh4.googleusercontent.com/MMFNYicHGcrUS0OivVU5TzjZLM90tyjlULOXxl1lDov1vCBwg93f6w2X-3TA2fP-cQSyOCLNHCwPit1EijUGgtYv9lxgJqNC2sU67a0fOSDn9nwi-bULgLPJ2x3l8EeWflr8w8CcccL2ahfLRFwVcSZHgj3XJLqB3cd_9XzO4gcVzlsPhCiJyZPMWA" alt="Snyk Web UI get API Token"><figcaption><p>Snyk Web UI get API Token</p></figcaption></figure>
3.  Switch to **Console** in the Snyk API v1 and add your organization ID as a parameter.\


    <figure><img src="https://lh3.googleusercontent.com/-sXMkOgM3GdCYP-15KqxtZ5DhxZlV3coqUZLYNdNnpVSdCFMH7wZApPhJAr9_8JxzAqyZOFGdIpqjT1t5Jpj570jQ67ykj_L3db4Gph3s74QOXdXjTwEJdRHRfWW0jpY14_lBAOinKC4x1An7yIIfHI-lk-cMULUosb8uDxC_z9mleGNkbdwUC3zVA" alt="Snyk API v1 add orgId to POST to Create a webhook"><figcaption><p>Snyk API v1 add orgId to POST to Create a webhook</p></figcaption></figure>


4.  In the Headers section add your Snyk API key to the **Authorization**.\


    <figure><img src="https://lh6.googleusercontent.com/nhlX0u7hJZSTue4rK01FLvComCMVmEQc1uE_z0nsnQ2_uK0ew5TFryBrTBkL24AKj03NjwKZvK5DsoN6j3fdKu0K9lX2a6SN2JP30m5-ST_Fj-IlMYO4Nu6PwDaDMeQH0ZPzyCF7__zc77iIaHRxxV2_57JDmgv7NbCeJi3Ti3LwP5K9UyYpkrma1A" alt="Snyk API v1 add API key to Authorization"><figcaption><p>Snyk API v1 add API key to Authorization</p></figcaption></figure>
5.  In the Body section add your values\
    `{` \
    &#x20;    `“url”: “value-of-your-apigw-url”,`\
    &#x20;    `“secret”: “value-of-your-lambda-secret-environment-variable”` \
    `}`\


    <figure><img src="https://lh5.googleusercontent.com/VXsSM6NFIWtWa_D4t_pJsWMUm3jHLMxSTEH8N7uLmb7IX98oxfm80_nPg0F6SGd-ffqth-iH3a2afcRQvE58hl5YoAP0NfvfaSPeUP6osRYdnPiPd1-ZOGUajvFk3vvOfXye_khV6lOylFC-T-47nLjclQD7ls8soL-EbWa8KAznWZJeLtj05eshSQ" alt="Snyk API v1 POST body"><figcaption><p>Snyk API v1 POST body</p></figcaption></figure>
6. Click **Call Resource**.

With this request done, your connection from Snyk to Slack will be completed. Every time there is a new vulnerability, you will get a new notification.
