# IaC ignores using the .snyk policy file

When scanning your IaC configuration files using the Snyk CLI with **snyk iac test** you can ignore issues that are not relevant to you.

You can do this by using the [.snyk policy file](fixing-and-prioritizing-issues/policies/the-.snyk-file), which we recommend is stored and versioned in the root of your working directory for where you store your IaC configuration files.

## Ignore paths

For rest runs using the Snyk CLI, only issues defined in the **.snyk** file are ignored.

For test runs from imported git repositories:

* Issues can be ignored in the Snyk UI - note these ignores will only apply to scans conducted using the Snyk UI.
* Important: These two sources of ignores are not synchronized.

## .snyk file semantics

{% hint style="info" %}
The **.snyk** file has some limitations for IaC projects \(see [The .snyk file](fixing-and-prioritizing-issues/policies/the-.snyk-file) for standard functionality\):

* The **patches** section is not yet supported and will be ignored.
* There are no IaC-supported language settings. This section will be ignored.
{% endhint %}

When running **snyk iac test** against a directory, either by passing in one or more directories or using the default argument of the current working directory, the Snyk CLI looks for a file named **.snyk** in each of those directories.

The syntax of the policy file is as follows:

```text
version: v1.19.0
ignore:
  SNYK-CC-K8S-1:
    - '*':
        reason: None Given
        expires: 2021-08-26T08:40:35.249Z
        created: 2021-07-27T08:40:35.251Z
```

This file can be created with the **snyk ignore** command, See [Ignore vulnerabilities](snyk-cli/fix-vulnerabilities-from-the-cli/ignore-vulnerabilities-using-snyk-cli).

The \`\*\` object key causes the CLI to ignore all instances of the SNYK-CC-K8S-1 vulnerability. You can add multiple entries, keyed by the IaC issue ID, to ignore multiple vulnerabilities.

## Ignoring a single file

Ignore rules can be scoped more narrowly. To scope the ignore to a single file, change the \`\*\` to the path of that file relative to the directory under test that contains the policy file:

```text
version: v1.19.0
ignore:
  SNYK-CC-K8S-1:
    - 'staging/deployment.yaml > *':
        reason: None Given
        expires: 2021-08-26T08:40:35.249Z
        created: 2021-07-27T08:40:35.251Z
  - 'staging/cronjob.yaml > *':
        reason: None Given
        expires: 2021-08-26T08:40:35.249Z
        created: 2021-07-27T08:40:35.251Z
```

In the example above we are ignoring an issue in 2 specific files.

## Ignore instances of a vulnerability

Individual instances of a vulnerability within a file can be ignored. To do this, you’ll need to take the “resource path” from the output of **snyk iac test**, and add it to the file path.

For example, from the following output snippet:

```text
Testing production/deployment.yaml...Infrastructure as code issues:
  ✗ Container is running in privileged mode [High Severity] [SNYK-CC-K8S-1] in Deployment
    introduced by [DocId: 0] > input > spec > template > spec > containers[web] > securityContext > privileged
```

You could modify the policy file as follows:

```text
version: v1.19.0
ignore:
  SNYK-CC-K8S-1:
    - 'production/deployment.yaml > [DocId:1] > spec > template > spec > containers[web] > securityContext > privileged':
        reason: None Given
        expires: 2021-08-26T08:40:35.249Z
        created: 2021-07-27T08:40:35.251Z
```

## Policy flags and policy file notes

There cannot be more than one policy file per directory under test. For example, **snyk iac test dir1/ dir2/** will load **dir1/.snyk** and **dir2/.snyk**, but if the file **dir1/foo/bar/.snyk** exists, the CLI will not load it.

When running **snyk iac test**, the CLI loads **$PWD/.snyk**. One common pattern is to use a single policy file per repository, in the root of that repository.

The CLI accepts a flag **--policy-path=..**., which overrides the location of policy files. The path can either be a directory containing a file named **.snyk**, or the path to a file named **.snyk**. The policy file’s actual name must be **.snyk**.

Policies are not loaded automatically when the argument to **snyk iac test** is a file rather than a directory. In this case, **--policy-path** must be specified in order to load policies.

The CLI accepts a flag **--ignore-policy**, which will cause any found **.snyk** policy files to be ignored.

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

