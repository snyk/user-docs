# Snyk Code - Clone capability with Broker for Docker

{% hint style="info" %}
**Feature availability**

The environment variable to enable Git clone capabilities is currently in closed beta. Contact your Snyk account management team to find out more.
{% endhint %}

By default, the Git clone capabilities required by Snyk Code are disabled. To grant Broker access to perform a Git clone of your repository, add the environment variable: `ACCEPT_CODE=true`.

Example:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e GITHUB_TOKEN=secret-github-token \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=http://my.broker.client:8000 \
           -e ACCEPT_CODE=true
       snyk/broker:github-com
```

This adds the necessary `accept` rules for your choice of Git server. These rules are in the [client templates in the Broker GitHub repository](https://github.com/snyk/broker/tree/master/client-templates).

If custom `accept` rules are required, you can provide a custom `accept.json`.

Example:

```console
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e GITHUB_TOKEN=secret-github-token \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=http://my.broker.client:8000 \
           -e ACCEPT=/myFolder/accept.json
       snyk/broker:github-com
```

If you are using a custom `accept` file, from a separate folder, with the ACCEPT environment variable, you cannot use the ACCEPT\_CODE mechanism.
