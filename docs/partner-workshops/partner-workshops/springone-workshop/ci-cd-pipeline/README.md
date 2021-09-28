# CI/CD Pipeline

## Getting started with the GitHub Actions Pipeline

Each step of the SDLC moves your application closer to production, and these steps represent another opportunity to assess and evaluate the vulnerabilities and licensing compliance of your application. Snyk provides numerous integration points along the SDLC to enforce security gates. A security gate is an organizational policy to help AppSec teams assess the risks of releasing code to production. In the exercise below, we will review the GitHub Actions pipeline and examine how to apply the security gates to your pipeline. When the pipeline executes it will send Synk test results to the Snyk UI, we will review the results in further exercises.

GitHub workflow diagram:

![](../../../.gitbook/assets/github_workflow_diagram_56-.png)

### Workshop exercises

We will complete the following steps. 

1. Populate the pipeline with GitHub secrets. We will add a Synk API token and Docker Hub API token.
2. Edit pipeline environment variables to build or container image, commit changes to the pipeline, and execute the pipeline.

