---
description: How to use one Snyk Broker token across multiple Organizations with the same SCM
---

# Enabling Broker across multiple Organizations

You can use the same Git service across multiple Organizations in Snyk with the same Broker token.

To do this, create the token for an Organization and then create a new Organization based on the original one. This clones the token, and you can now enable the Broker for it.

To do this retroactively for existing Organizations, you can use the endpoint [Clone an integration (with settings and credentials) ](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/snyk-api/reference/integrations-v1#org-orgid-integrations-integrationid-clone)to clone a specific integration, including the Broker token.

Unless you do this, you must generate a new Broker token for the Organization, as each integration and Organization has its own unique Broker token.
