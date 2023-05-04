# Snyk Code - Clone capability with Broker

{% hint style="info" %}
**Feature availability**

The environment variable to enable Git clone capabilities is currently in closed beta. Contact your Snyk account management team to find out more.
{% endhint %}

By default, the Git clone capabilities required by Snyk Code are disabled. To grant Broker access to perform a Git clone of your repository, you can add an environment variable: `ACCEPT_CODE=true`.

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

If you are using a custom `accept` file, from a separate folder, with the ACCEPT environment variable, you cannot use the ACCEPT\_CODE mechanism.
