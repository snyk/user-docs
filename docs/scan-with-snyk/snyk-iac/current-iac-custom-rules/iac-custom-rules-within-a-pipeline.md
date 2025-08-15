# IaC custom rules within a pipeline

Using a CI/CD such as [GitHub Actions](https://github.com/features/actions) is ideal for managing, distributing, and enforcing your custom rules.

## Overview of IaC custom rules in a pipeline

This example shows how a security team can:

* Store their rules in a GitHub repository
* Use GitHub Actions to add different development-time steps to their pipelines
* Configure a different GitHub repository to run a GitHub Action pipeline that uses the custom rules to gate changes.

Snyk uses the [snyk/custom-rules-example](https://github.com/snyk/custom-rules-example) repository for the example; this repo contains all the custom rules written while [getting started with the SDK](writing-rules-using-the-sdk/).

Aims: Configure our pipeline to:

* Verify that new rules or changes to the existing rules do nor break existing functionality.
* Publish the rules in `main` to an OCI registry.
* Enforce the use of custom rules in other pipelines.
* Optionally: Configure the custom rules using environment variables.

## Adding PR checks using GitHub Action

An example of a PR check can be seen in [https://github.com/snyk/custom-rules-example/pull/5](https://github.com/snyk/custom-rules-example/pull/5) where there is an attempt to add a new rule called `my_rule`. This is the same rule shown for [learning how to write a rule](writing-rules-using-the-sdk/writing-a-rule.md).

To verify that this rule works as expected, unit tests were implemented. To run the unit tests as part of PR checks, a GitHub Action was configured previously under `.github/workflows` called `test.yml`:

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

* It was configured to run on all non-`main` branches, so that it runs when PRs are open.
* Steps were added steps to set up a Node.js environment to enable installing the `snyk-iac-rules` SDK using [npm](install-the-sdk.md#install-the-sdk-with-npm).
* A step was added to run `snyk-iac-rules test`, which will cause the PR check to fail if any of the tests fail.

{% hint style="info" %}
You must configure your `main` branch under `Settings` -> `Branches`first, so that no one can push directly to `main`.
{% endhint %}

## Snyk IaC GitHub Action

Another way to test the rules is by testing the contract with the [Snyk CLI](../../../developer-tools/snyk-cli/) by using the [Snyk IaC GitHub Action](https://github.com/snyk/actions/tree/master/iac), making sure the generated bundle can be read by the CLI.

To do this, you will need a step for installing the Snyk CLI and a `SNYK_TOKEN`, which can be found in your Snyk Account Settings.

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

You can also expand these tests to use [Shellspec](https://github.com/shellspec/shellspec) and verify that the desired vulnerabilities get triggered, but Snyk recommends using the unit tests for this.

## Publishing the custom rules

Once a PR passes its checks from the previous section and gets merged into the `main` branch, you can publish the Snyk rules to an OCI registry. This allows you to configure a separate pipeline, download the custom rules bundle from this location, and run the custom rules in order to catch misconfigurations.

For this, add another workflow under `.github/workflows` called `publish.yml`:

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
          OCI_REGISTRY_URL: "${{ secrets.OCI_REGISTRY_NAME }}:v1"
```
{% endcode %}

It looks like the previous workflow, but there are a few things to note about this one:

* It was configured to run only on `main` branches, so that it runs when PRs are merged.
* A step was added to authenticate with Docker Hub, our chosen OCI registry. For a list of supported registries, read about [pushing bundles](writing-rules-using-the-sdk/pushing-a-bundle.md). Use the [docker/login-action](https://github.com/docker/login-action) GitHub Action to do that and be sure to configure the GitHub secrets under `Settings` -> `Secrets`.
* A step was added to run `snyk-iac-rules build` followed by `snyk-iac-rules push`, which will publish the generated custom rules bundle to an OCI registry.

## Versioning rules

If you want to release an experimental version of the custom rules without affecting all our CI/CD pipelines, use tagging to version the bundles.

Start trialing bundle `v2-beta` while still using `v1` in most of our services:

{% code title=".github/workflows/publish.yml" %}
```
      - name: Publish experimental rules
        run: snyk-iac-rules push --registry $OCI_REGISTRY_URL bundle.tar.gz
        env:
          OCI_REGISTRY_URL: "${{ secrets.OCI_REGISTRY_NAME }}:v1"
      - name: Publish rules
        run: snyk-iac-rules push --registry $OCI_REGISTRY_URL bundle.tar.gz
        env:
          OCI_REGISTRY_URL: "${{ secrets.OCI_REGISTRY_NAME }}:v2-beta"
```
{% endcode %}

{% hint style="info" %}
Ensure that the OCI\_REGISTRY\_NAME configured in the GitHub Secrets does not already contain the tag or the protocol if you want to use this workflow.
{% endhint %}

## Enforcing the custom rules

After publishing the custom rules to an OCI registry, you can configure a separate pipeline to use these rules.

One way to do this is by using the API endpoint [Update the Infrastructure as Code settings for a group](../../../snyk-api/reference/iacsettings.md#patch-groups-group_id-settings-iac).

This means configuring the GitHub Action above with another job for updating Snyk to use the configured custom rules bundle:

```
      - name: Update Snyk
        run: |
          curl --location --request PATCH 'https://api.snyk.io/rest/groups/<group id>/settings/iac/?version=2021-11-03~beta' \
          --header 'Content-Type: application/vnd.api+json' \
          --header 'Authorization: token ${{ secrets.SNYK_TOKEN }}' \
          --data-raw '{
            "data": {
                  "type": "iac_settings",
                  "attributes": {
                    "custom_rules": {
                      "oci_registry_url": "registry-1.${{ secrets.OCI_REGISTRY_NAME }}",
                      "oci_registry_tag": "v1",
                      "is_enabled": true
                    }
                }
            }
          }'
```

This API call will update the chosen Snyk Group and all the Organizations underneath it to use the configured custom rules bundle.

{% hint style="info" %}
To configure an Organization to use a different bundle, such as the `v2-beta` one, use the Snyk Settings page. You can either configure a new bundle or disable custom rules to allow using environment variables in the CI/CD pipeline to run the custom rules.
{% endhint %}

In a different repository, authenticate with one of the Organizations underneath this Group and add the Snyk IaC GitHub Action to a workflow:

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
```

The result is that the GitHub action will fail until the generated misconfigurations have been resolved:

```
Testing example.tf...


Infrastructure as code issues:
  ✗ IAM Role missing one of the required tags: owner, description or type [Medium Severity] [CUSTOM-RULE-8]
    introduced by input > resource > aws_iam_role[new_role] > tags

  ✗ Vendor or Service must have either owneralternate or ticket group or both tags. [Medium Severity] [CUSTOM-RULE-9]
    introduced by input > resource > aws_iam_role[new_role] > tags
```

## Configuring the custom rules

In case using an API or the Snyk Settings page seems too restrictive, Snyk also provides a way to configure the custom rules by using environment variables.

You can use the Snyk IaC GitHub Action with the `SNYK_CFG_OCI_REGISTRY_URL`, `SNYK_CFG_OCI_REGISTRY_USERNAME`, and `SNYK_CFG_OCI_REGISTRY_PASSWORD` environment variables to scan your configuration files for any custom rules that may have been breached.

The GitHub Action reads these environment variables and pulls down the bundle pushed in the previous step to the configured OCI registry. The GitHub action looks like this:

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
