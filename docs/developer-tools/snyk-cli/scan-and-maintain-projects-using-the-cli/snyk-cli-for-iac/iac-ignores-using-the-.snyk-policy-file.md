# IaC ignores using the .snyk policy file

When you scan IaC configuration files using the Snyk CLI `iac test` command, you can ignore issues that are not relevant to you by using the [`.snyk` policy file](../../../../manage-risk/policies/the-.snyk-file.md). Snyk recommends that you store and version the `.snyk` file in the root of the working directory where you store your IaC configuration files. This file can be created with the `snyk ignore` command. For details see [Ignore vulnerabilities using Snyk CLI](../ignore-vulnerabilities-using-the-snyk-cli.md).

## Ignore paths

For tests run using the Snyk CLI, only issues defined in the `.snyk` file are ignored.

For tests run from imported Git repositories, issues can be ignored in the Snyk UI. Note that these ignores apply only to scans done using the Snyk UI.

{% hint style="warning" %}
Ignores in the `.snyk` file and ignores created in the Snyk UI are not synchronized.
{% endhint %}

## `.snyk` file semantics

The `.snyk` file has some limitations for IaC Projects. See [The `.snyk` file](../../../../manage-risk/policies/the-.snyk-file.md) for standard functionality.

* The **patches** section is not yet supported and is ignored.
* There are no IaC-supported **language settings**. This section is ignored.

When you run `snyk iac test` against a directory, either by passing in one or more directories or using the default argument of the current working directory, the Snyk CLI looks for a file named `.snyk` in each of those directories.

The syntax of the policy file is as follows:

```
version: v1.19.0
ignore:
  SNYK-CC-K8S-1:
    - '*':
        reason: None Given
        expires: 2021-08-26T08:40:35.249Z
        created: 2021-07-27T08:40:35.251Z
```

The `*` object key causes the CLI to ignore all instances of the `SNYK-CC-K8S-1` vulnerability. You can add multiple entries, keyed by the IaC issue ID, to ignore multiple vulnerabilities.

## Ignoring a single file

Ignore rules can be scoped more narrowly. To scope the ignore to a single file, change the `*` to the path of that single file relative to the directory being tested that contains the `.snyk` policy file.

You can specify scoped ignore rules either by using the `ignore` command in the Snyk CLI or manually modifying the `.snyk` file.

In the following example, an issue is being ignored with the `SNYK-CC-K8S-1` ID in two specific files:

* `staging/deployment.yaml`
* `staging/cronjob.yaml`

You can generate the scoped ignore rules with the Snyk CLI by running the following commands:

```
snyk ignore --id=SNYK-CC-K8S-1 --path='staging/cronjob.yaml > *'
snyk ignore --id=SNYK-CC-K8S-1 --path='staging/deployment.yaml > *'
```

Alternatively, manually modify the `.snyk` policy file as follows:

```
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

For more information about the Snyk CLI ignore command, see [Ignore vulnerabilities using Snyk CLI](../ignore-vulnerabilities-using-the-snyk-cli.md).

## Ignore instances of a vulnerability

Individual instances of a vulnerability within a file can be ignored. To do this, take the “resource path” from the output of `snyk iac test`, and add it to the file path.

For example, from the following output snippet (line break added for ease of reading):

```
Testing production/deployment.yaml...Infrastructure as code issues:
  ✗ Container is running in privileged mode [High Severity] [SNYK-CC-K8S-1] in Deployment
    introduced by [DocId: 0] > input > spec > template > spec > containers[web] 
    > securityContext > privileged
```

you could generate the scoped ignore rule with the Snyk CLI by running the following command:

```
 snyk ignore --id=SNYK-CC-K8S-1 --path='production/deployment.yaml > [DocId:1] > spec > template > spec > containers[web] 
 > securityContext > privileged'
```

Alternatively, manually modify the policy file as follows:

```
version: v1.19.0
ignore:
  SNYK-CC-K8S-1:
    - 'production/deployment.yaml > [DocId:1] > spec > template > spec > containers[web] > securityContext > privileged':
        reason: None Given
        expires: 2021-08-26T08:40:35.249Z
        created: 2021-07-27T08:40:35.251Z
```

For more information about the Snyk CLI ignore command, see [Ignore vulnerabilities using Snyk CLI](../ignore-vulnerabilities-using-the-snyk-cli.md).

## Policy flags and policy file notes

You cannot have more than one `.snyk` policy file for each test. For example, the command `snyk iac test dir1/ dir2/` loads `dir1/.snyk` and `dir2/.snyk`, but if the file `dir1/foo/bar/.snyk` exists, the CLI does not load it.

When you run `snyk iac test`, the CLI loads `$PWD/.snyk`. One common pattern is to use a single `.snyk` policy file per repository in the root of that repository.

The CLI accepts an option, `--policy-path=...`, which overrides the location of `.snyk` policy files. The path can either be a directory containing a file named `.snyk` or the path to a file named `.snyk`. The name of the policy file must be `.snyk`.

Policies are not loaded automatically when the argument to `snyk iac test` is a file rather than a directory. In this case, `--policy-path` must be specified in order to load policies.

The CLI accepts the option `--ignore-policy`, which causes any `.snyk` policy files that are found to be ignored.
