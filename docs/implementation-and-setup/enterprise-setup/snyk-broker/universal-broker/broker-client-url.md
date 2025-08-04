# Broker client URL

A webhook from an SCM to Snyk allows the SCM integration to communicate with Snyk. When code is updated, whether the update is to an open-source manifest file or to code files, and the developer creates a pull request, the SCM notifies Snyk through the webhook.

The webhook can be set to reach out to Snyk directly to indicate that files have changed, and Snyk will request the files through Broker. The SCM then sends those files to Snyk, and Snyk scans for code, security, and license issues based on your integrations and the products you have, and returns a pass/fail determination for the SCM check.

Sometimes the webhook cannot communicate directly with Snyk because the infrastructure prevents the repository from reaching outside of the private cloud or data center. Then the Broker facilitates the communication, sending notification of changes to the code, sending the request for files to the SCM, and conveying the files to Snyk, which returns the results of the scan. You can then continue work: review the status, ask for more information, see the issues in the PR if you are using the inline comments capability, and view the details in the Snyk portal.

The same webhook mechanism is used to facilitate communication when Snyk creates a pull request, for example, when you see an issue in the Snyk interface and use the button to fix the vulnerability.

When you set up a connection, the webhook target endpoint is defined by the value of the `broker_client_url` value. The webhook can point directly to Snyk or to the Broker client container, which will relay the webhook to Snyk.&#x20;

Normally, and by default for SCM integrations, the regional Snyk API endpoint is used. For the list of URLs, see [Broker client URLs](../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-client-urls). Sometimes, however, your environment may prohibit SCM webhooks from leaving the private cloud or data center. In that case, the webhook and thus the SCM must point to the Snyk Broker container running in your environment. In these cases, the`broker_client_url` has to reflect the hostname and port of the Broker client.&#x20;

Non-SCM integrations, notably the container registries integrations, require the Broker client URL to be the Broker client address. Snyk recommends using the DNS hostname (`http://my.broker.client`), but you can also use IP addresses (`http://192.168.0.1`). Append the port, for example, `http://my.broker.client:8000`.

Note that the `https` webhook calls requires additional setup in the Broker client to bring a TLS certificate and mount it into the container. For details, see [HTTPS for Broker client with Docker](../https-for-broker-client-with-docker.md).
