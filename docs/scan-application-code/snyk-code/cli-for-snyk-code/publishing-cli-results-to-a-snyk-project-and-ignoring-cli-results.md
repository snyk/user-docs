# Publishing CLI results to a Snyk Project and ignoring CLI results (beta)

{% hint style="info" %}
**Feature availability**\
Publishing and ignoring CLI results is in [Closed Beta](../../../more-info/snyk-feature-release-process.md#closed-beta) and only available by invitation from Snyk.\
\
Minimum supported CLI version: v1.1138.0
{% endhint %}

Snyk Code CLI supports publishing the results to a Snyk Project in the Web UI and respecting issues that were ignored in a Snyk Project in the Web UI so you can filter them from the analysis results.

This enables Code to be used as a blocking CI/CD gate to test and block builds at the main branch level and then have developers review the results in the Web UI, fix any newly introduced vulnerabilities, or ignore irrelevant ones.

## **Publishing CLI results to a Snyk Project**

In the terminal, enter:

```
snyk code test --report --project-name="PROJECT_NAME"
```

* PROJECT\_NAME must be in double quotation marks. Single quotes or missing quotes will result in an error.
* The Project name must contain only alphanumeric characters, forward slashes (/), dashes (-), underscores (\_), and square brackets (\[]), and be no longer than 64 characters
* There is a (temporary) limit of 3MB for the resulting payload, meaning that for large Projects with many issues, the process will not complete.

Running the `snyk code test` command with the `--report` option as shown returns the results to the terminal window, along with a URL to the Snyk Code Project where the results have been published. Refer to the following screenshot.

<figure><img src="../../../.gitbook/assets/image (2) (7).png" alt="Snyk code test results with --report option"><figcaption><p>Snyk code test results with --report option</p></figcaption></figure>

If a CLI-based Snyk Code Project does not yet exist for the value provided in the `--project-name` option, Snyk creates a new CLI-based project. If a CLI-based Project already exists, a new snapshot is created under the same Project

## **Ignoring CLI results**

You can ignore issues from CLI results in the Web UI by using the ignore button:

<figure><img src="../../../.gitbook/assets/image (1) (6).png" alt="Ignoring issues in the Web UI"><figcaption><p>Ignoring issues in the Web UI</p></figcaption></figure>

Issues that are ignored in the Web UI are ignored in CLI tests when you use the following command:

```
snyk code test --report --project-name="PROJECT_NAME"
```

Ignores that have been applied to the Project with a `PROJECT_NAME` suppress the issue the next time that the CLI runs for the same `PROJECT_NAME`.

\
