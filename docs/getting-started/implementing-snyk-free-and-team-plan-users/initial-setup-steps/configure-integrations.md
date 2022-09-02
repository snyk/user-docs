# Configure integrations



Under **Settings > Integrations**, select the appropriate code repository and complete the necessary authorization steps outlined in the setup guide.

See [Snyk Integrations](../../../integrations/) for more details.

### Initial disable PR tests

As part of your rollout process you may decide to globally disable **Snyk test for pull request** checks at the start, only enabling them on your most important projects initially.

1. Select **Disable Pull Request check** at the Git integration screen for your Git platform/tool.
2. For Projects where you want to enable these tests, navigate to the Project settings and enable **Snyk test for pull request** **checks** for that Project.

### Configure Jira integration

For Team Plan Users, additionally you can configure [Jira Integration](https://docs.snyk.io/integrations/notifications-ticketing-system-integrations/jira) to assist with logging tickets and addressing backlogged security issues.
