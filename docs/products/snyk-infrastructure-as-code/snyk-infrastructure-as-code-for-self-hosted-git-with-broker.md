# Snyk Infrastructure as code for self-hosted git (with broker)

Snyk Broker enables you to connect your local git server to Snyk if the git server is not internet reachable. By default, Snyk Broker only allows information for Snyk Opensource and Dockerfiles to go through. If you want to also analyze Infrastructure as Code files such as `.tf` or `.yaml` then additional configuration is required.&#x20;

To learn how to configure Snyk broker see the [broker documentation](../../features/integrations/snyk-broker/broker-introduction.md) for details.

## Snyk Infrastructure as code configuration

You will need to grant the broker access to particular files in the repository. This requires specific API permissions. These API permissions are slightly different depending on which source control system you are using.&#x20;

To learn what to configure and how, see [Snyk Broker Infrastructure as Code detection ](broken-reference)documentation for more details.
