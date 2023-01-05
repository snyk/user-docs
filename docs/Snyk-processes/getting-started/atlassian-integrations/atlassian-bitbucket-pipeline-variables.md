# Atlassian Bitbucket Pipeline Variables

Atlassian Bitbucket Pipelines are a feature your Bitbucket Cloud administrators enable to enable automated tasks.  Snyk workshops rely on Repository Variables to parameterize operations, and this section describes what variables are needed.

Atlassian provides [documentation to get started](https://support.atlassian.com/bitbucket-cloud/docs/get-started-with-bitbucket-pipelines/) with Bitbucket Pipelines.

In  your workshop, the bitbucket repositories we provide will usually contain a starter pipeline for the workshop you will automatically use in the onboarding of your repository.  These repositories are described in the respective workshop documentation.&#x20;

## Repository variables

You will need to define [repository variables](https://support.atlassian.com/bitbucket-cloud/docs/variables-in-pipelines/#Repository-variables) at the repository level which will later be [referenced in your pipeline](https://support.atlassian.com/bitbucket-cloud/docs/variables-in-pipelines/).

This screenshot is representative of the secured and unsecured variables, specific to each workshop and its integrations.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/bitubucket-repo-vars.png)

See the documentation for your workshop for the required variables.
