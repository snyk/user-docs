# Integrating IaC custom rules within a pipeline

The ideal scenario for managing, distributing, and enforcing your custom rules is to use a CI/CD like [GitHub Actions](https://github.com/features/actions).&#x20;

### Overview

This example shows how a security team can:

* Store their rules in a GitHub repository
* Use GitHub Actions to add different development-time to their pipelines
* Configure a different GitHub repository to run a GitHub Action pipeline that uses the custom rules to gate changes.

We use the [snyk/custom-rules-example](https://github.com/snyk/custom-rules-example) repository for the example; this repo contains all the custom rules written while [getting started with the SDK](getting-started-with-the-sdk/).

#### Aims

We want to configure our pipeline to achieve two things:

1. Verify that new rules or changes to the existing rules don't break existing functionality.
2. Publish the rules in `main` to an OCI registry.
3. Enforce the usage of custom rules in other pipelines.

### Adding PR checks using GitHub Action

An example of a PR check can be seen in [https://github.com/snyk/custom-rules-example/pull/5](https://github.com/snyk/custom-rules-example/pull/5) where we attempt to add a new rule called `my_rule`&#x20;

(**note**: this is the same rule we showed when [learning how to write a rule](getting-started-with-the-sdk/writing-a-rule.md))

To verify that this rule works as expected, we have implemented unit tests. To run the unit tests as part of PR checks, we previously configured a GitHub Action under `.github/workflows` called `test.yml`:

{% code title=".github/workflows/test.yml" %}
```
name: Test Custom Rules

on:
  push:
    branches:
      - '**'        # matches every branch
      - '!main'     # excludes main

jobs:
  unit_test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2

      - uses: actions/setup-node@v1
        with:
          node-version: 15

      - name: Install snyk-iac-rules
        run: npm i -g snyk-iac-rules

      - name: Run unit tests
        run: snyk-iac-rules test
```
{% endcode %}

A few things to note about this workflow:

* We configured it to run on all non-`main` branches, so that it runs when PRs are open.
* We added steps to setup a Node.js environment, so that we can then install the `snyk-iac-rules` SDK using [npm](install-the-sdk.md#install-the-sdk-with-npm).
* We added a step to run `snyk-iac-rules test`, which will cause the PR check to fail if any of the tests fail.

{% hint style="info" %}
You need to configure your `main` branch under `Settings` -> `Branches`first, so that no one can push directly to `main`.
{% endhint %}

### Snyk IaC GitHub Action

Another way to test the rules is by testing the contract with the [Snyk CLI](../../../features/snyk-cli/) by using the [Snyk IaC GitHub Action](https://github.com/snyk/actions/tree/master/iac), making sure the generated bundle can be read by the CLI.&#x20;

To do this, you will need a step for installing the Snyk CLI and a `SNYK_TOKEN`, which can be found in your Snyk Account Settings.&#x20;

{% code title=".github/workflows/test.yml" %}
```
jobs:
  contract_test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2

      - uses: actions/setup-node@v1
        with:
          node-version: 15

      - name: Install snyk-iac-rules
        run: npm i -g snyk-iac-rules

      - name: Build bundle
        run: snyk-iac-rules build .

      - name: Run contract with Snyk to check Infrastructure as Code files for issues
        continue-on-error: true
        uses: snyk/actions/iac@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
        with:
          args: --rules=bundle.tar.gz
```
{% endcode %}

You can also expand these tests to use [Shellspec](https://github.com/shellspec/shellspec) and verify that the desired vulnerabilities get triggered, but we recommend using the unit tests for this.

### Publishing the custom rules

Once a PR passes its checks from the previous section and gets merged into the `main` branch, you can publish our rules to an OCI registry. This allows you to configure a separate pipeline, to download the custom rules bundle from this location, and run the custom rules in order to catch misconfigurations.

For this, we will add another workflow under `.github/workflows` called `publish.yml`:

{% code title=".github/workflows/publish.yml" %}
```
name: Publish Custom Rules

on:
  push:
    branches:
      - 'main'

jobs:
  publish:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2

      - uses: actions/setup-node@v1
        with:
          node-version: 15

      - name: Install snyk-iac-rules
        run: npm i -g snyk-iac-rules
        
      - name: Build bundle
        run: snyk-iac-rules build .
        
      - name: Login to Docker Hub
        uses: docker/login-action@v1
        with:
          username: ${{ secrets.OCI_REGISTRY_USERNAME }}
          password: ${{ secrets.OCI_REGISTRY_PASSWORD }}

      - name: Publish rules
        run: snyk-iac-rules push --registry $OCI_REGISTRY_URL bundle.tar.gz
        env:
          OCI_REGISTRY_URL: ${{ secrets.OCI_REGISTRY_URL }}
```
{% endcode %}

It looks similar to the previous workflow, but there are a few things to note about this one:

* We configured it to run only on `main` branches, so that it runs when PRs are merged.
* We added a step to authenticate with Docker Hub, our chosen OCI registry. For a list of supported registries read about [pushing bundles](getting-started-with-the-sdk/pushing-a-bundle.md). Use the [docker/login-action](https://github.com/docker/login-action) GitHub Action to do that and make sure to configure the GitHub secrets under `Settings` -> `Secrets`.
* We added a step to run `snyk-iac-rules build` followed by `snyk-iac-rules push`, which will publish our generated custom rules bundle to an OCI registry.

#### Versioning rules

If we want to release an experimental version of the custom rules without affecting all our CI/CD pipelines, we can use tagging to version our bundles.&#x20;

So, we can start trialing bundle `v2-beta` while still using `v1` in most of our services:

{% code title=".github/workflows/publish.yml" %}
```
      - name: Publish experimental rules
        run: snyk-iac-rules push --registry $OCI_REGISTRY_URL:v2-beta bundle.tar.gz
        env:
          OCI_REGISTRY_URL: ${{ secrets.OCI_REGISTRY_URL }}
      - name: Publish rules
        run: snyk-iac-rules push --registry $OCI_REGISTRY_URL:v1 bundle.tar.gz
        env:
          OCI_REGISTRY_URL: ${{ secrets.OCI_REGISTRY_URL }}
```
{% endcode %}

{% hint style="info" %}
Make sure that the `OCI_REGISTRY_URL` configured in the GitHub Secrets does not already contain the tag if you want to use this workflow.
{% endhint %}

### Enforcing the custom rules

After publishing custom rules to an OCI registry, you can configure a separate pipeline to use these rules. For example, see [https://github.com/snyk/infrastructure-as-code-goof/pull/67](https://github.com/snyk/infrastructure-as-code-goof/pull/67).

You can use the Snyk IaC GitHub Action with the `SNYK_CFG_OCI_REGISTRY_URL`, `SNYK_CFG_OCI_REGISTRY_USERNAME`, and `SNYK_CFG_OCI_REGISTRY_PASSWORD` environment variables to scan your configuration files for any custom rules which may have been breached.&#x20;

The GitHub Action reads these environment variables and pulls down the bundle pushed in the previous step to the configured OCI registry. The GitHub action will look similar to this:

```
name: Snyk Infrastructure as Code Custom Rules

on:
  push:

jobs:
  snyk-iac-security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2

      - name: Run Snyk to check Infrastructure as Code files for issues
        continue-on-error: false
        uses: snyk/actions/iac@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
          SNYK_CFG_OCI_REGISTRY_URL: ${{ secrets.OCI_REGISTRY_URL }}
          SNYK_CFG_OCI_REGISTRY_USERNAME: ${{ secrets.OCI_REGISTRY_USERNAME }}
          SNYK_CFG_OCI_REGISTRY_PASSWORD: ${{ secrets.OCI_REGISTRY_PASSWORD }}
```

The result is that the GitHub action will fail until the generated misconfigurations have been resolved:

```
Testing example.tf...


Infrastructure as code issues:
  ✗ Non-encrypted Redshift DB at rest [Medium Severity] [SNYK-CC-TF-108] in Redshift
    introduced by resource > aws_redshift_cluster[test] > encrypted

  ✗ Missing a description and an owner from tag, or owner tag does not comply with email requirements [Medium Severity] [CUSTOM-RULE-4]
    introduced by input > resource > aws_redshift_cluster[test] > tags

  ✗ Missing a description or an owner from the tag [Medium Severity] [CUSTOM-RULE-3]
    introduced by input > resource > aws_redshift_cluster[test] > tags

  ✗ Missing an owner from tag [Medium Severity] [CUSTOM-RULE-1]
    introduced by input > resource > aws_redshift_cluster[test] > tags

  ✗ Missing a description and an owner from the tag [Medium Severity] [CUSTOM-RULE-2]
    introduced by input > resource > aws_redshift_cluster[test] > tags

  ✗ Redshift cluster logging disabled [Low Severity] [SNYK-CC-TF-136] in Redshift
    introduced by resource > aws_redshift_cluster[test] > logging
```

