# Snyk Infrastructure as code for self-hosted git (with Broker)

Snyk Broker enables you to connect your local git server to Snyk if the git server is not internet- reachable. By default Snyk Broker allows only information for Snyk Open Source and Docker files to go through. If you also want to analyze Infrastructure as Code files such as `.tf` or `.yaml` then additional configuration is required.

To learn how to configure Snyk Broker see the [Broker documentation](broken-reference).

## Snyk Infrastructure as Code configuration

You must grant Broker access to particular files in the repository. This requires specific API permissions. These API permissions are slightly different depending on which source control system you are using.

To learn what to configure and how, see [Snyk Broker Infrastructure as Code detection](../../integrations/snyk-broker/snyk-broker-infrastructure-as-code-detection/) documentation.
