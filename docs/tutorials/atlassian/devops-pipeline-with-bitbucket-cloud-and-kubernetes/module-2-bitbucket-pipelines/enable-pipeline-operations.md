---
description: Run a Bitbucket Pipeline to see automated results and deploy your application
---

# Enable Pipeline Operations

### Enable Bitbucket Pipelines

You may need to enable [Bitbucket Pipelines](../../../getting-started/atlassian-integrations/atlassian-bitbucket-pipeline-variables.md) for the repository when the default is to start with pipelines disabled.

The [Snyk Bitbucket Pipelines Docs](../../../../features/integrations/ci-cd-integrations/bitbucket-pipelines-integration-overview.md) and [Snyk Bitbucket Pipelines Video](broken-reference) contain additional integration information, and you will need permissions in your repository to enable pipeline operations.

When you enable pipelines, we'll focus on the main branch, named either main or master. The reference repository already contains a bitbucket\_pipelines.yaml at the root directory, and we'll need to configure some variables to make it work. You can modify variables once the pipeliene is enabled, and we cover that in the next section.
