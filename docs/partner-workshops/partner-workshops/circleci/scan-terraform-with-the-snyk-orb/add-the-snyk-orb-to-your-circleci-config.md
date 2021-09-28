# Add the Snyk Orb to your CircleCI Config

In the CircleCI Academy Orbs module you learned about Orbs, packages of configuration that simplify your builds. [Snyk's Orb](https://circleci.com/developer/orbs/orb/snyk/snyk) exposes the [Snyk CLI](https://support.snyk.io/hc/en-us/articles/360003812578-CLI-reference), allowing you to find and fix known vulnerabilities in app dependencies, container images, and infrastructure as code. 

## Add the Snyk Orb to your CircleCI configuration YML

In your fork of the [learn-iac](https://github.com/datapunkz/learn_iac) repo, open the `.circleci/config.yml` file. Add the Snyk Orb to the top replacing `@x.y.z` with the latest version of the Snyk Orb from the Orb Registry. 

{% embed url="https://circleci.com/developer/orbs/orb/snyk/snyk\\" caption="Snyk Orb in CircleCI Orb Registry" %}

```text
version: 2.1

orbs:
  snyk: snyk/snyk@x.y.z
  
jobs:
  ...
```

Adding the Orb exposes the `snyk` commands and jobs to your workflow. Consider your requirements when choosing where in the workflow to add them. 

## Add the Scan IAC Job to the Workflow

For this example, add the `snyk/scan-iac` job before the `gke-create-cluster` job to check Terraform files are correctly configured before creating the cloud infrastructure. The `args` parameter points to which files to check for misconfigurations and can also be used to pass [other Snyk CLI arguments](https://support.snyk.io/hc/en-us/articles/360018728618-Test-your-configuration-files).

```text
workflows:
  build_test:
    jobs:
      - run_tests
      - build_docker_image
      - snyk/scan-iac:
          args: part03/iac_gke_cluster/
      - gke_create_cluster:
          requires:
            - run_tests
            - build_docker_image
            - snyk/scan-iac
```

{% hint style="info" %}
Snyk Infrastructure as Code can also check Kubernetes and AWS CloudFormation files for misconfigurations. Learn more in the [Snyk IAC documentation](https://support.snyk.io/hc/en-us/articles/360018728618-Test-your-configuration-files). 
{% endhint %}

When ready, commit and merge your changes to trigger the workflow run.

## Working with results

When the workflow runs, the output will be displayed in your CircleCI project run. The job fails because issues are found in the `main.tf` file scanned.

![Snyk Orb output in the CircleCI UI](../../../.gitbook/assets/iac-job-run-fail.png)

{% hint style="info" %}
Visit [Understanding configuration scan issues in the Snyk Docs](https://support.snyk.io/hc/en-us/articles/360012499738-Understanding-configuration-scan-issues) to learn more about interpreting the output of the Snyk IAC CLI powering the Snyk Orb.
{% endhint %}

### Viewing results in the Snyk UI

Import your fork of the `learn-iac` repo to the Snyk UI using the GitHub integration. Visit the [Snyk IAC documentation](https://support.snyk.io/hc/en-us/articles/360011018938-Configure-your-integration-to-find-security-issues-in-your-Terraform-files) to learn how. Once imported, you'll see the manifest files in the Snyk UI.

![](../../../.gitbook/assets/imported-iac-project.png)

Clicking on the `main.tf` file will show you an in-line view of the issues found, with additional information such as the impact of the configuration and how to fix it.

![](../../../.gitbook/assets/iac-result-details.png)

In the next section we'll show how you can tune this analysis to adjust the test's pass/fail criteria.

