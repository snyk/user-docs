# Snyk Code - Clone capability with Broker for Docker

Brokered Snyk Code enables the Broker to accept code files. The Broker then scans between the SCM system and Snyk.

By default, the Git clone capabilities required by Snyk Code are disabled.

To grant Broker access to perform a Git clone of your repository, add the environment variable: `ACCEPT_CODE=true`

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

This adds the necessary `accept` rules for your Git server.&#x20;

After this is done, you can follow the Broker instructions for your SCM system. For details, see [Install and configure Snyk Broker](../).
