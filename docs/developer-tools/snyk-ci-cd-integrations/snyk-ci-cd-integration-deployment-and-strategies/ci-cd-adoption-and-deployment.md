# CI/CD adoption and deployment

When deciding to use a Snyk integration, compare the advantages of source control management (SCM) integrations and CI/CD integrations.&#x20;

## Typical stages in adopting CI/CD Integration

Developer teams typically adopt Snyk in the following stages (example commands are Snyk Open Source):

1. [Expose vulnerabilities](ci-cd-adoption-and-deployment.md#stage-1-expose-vulnerabilities-snyk-monitor) (`snyk monitor`)
2. [Use Snyk as a gatekeeper](ci-cd-adoption-and-deployment.md#stage-2-use-snyk-as-a-gatekeeper-snyk-test) (`snyk test`)
3. [Continuous monitoring](ci-cd-adoption-and-deployment.md#stage-3-continuous-monitoring-snyk-test-and-snyk-monitor) (`snyk test` and `snyk monitor`)

### Stage 1: Expose vulnerabilities (`snyk monitor`)

A typical approach is to use Snyk results to expose vulnerabilities during the development process. This increases visibility of vulnerabilities among the members of your team.

When you first implement Snyk in your pipeline, Snyk recommends using only the `snyk monitor` command. If you use one of the Snyk CI plugins, Snyk recommends that you configure the plugin to NOT fail the build.

This is because all Projects have vulnerabilities, and after you set Snyk to fail the build, every build fails because of Snyk. This may cause problems with your team being quickly overwhelmed with failure messages.

Using `snyk monitor` to expose results provides information without disrupting processes.

For information about `snyk monitor`, see the [`monitor` command help](../../snyk-cli/commands/monitor.md).

### Stage 2: Use Snyk as a gatekeeper (`snyk test`)

Using Snyk as a gatekeeper prevents the introduction of new vulnerabilities (sometimes referred to as "stopping the bleeding").

After your teams understand the vulnerabilities in their applications and develop a process for fixing them early in the development cycle, you can configure Snyk to fail your builds to prevent introducing vulnerabilities into your applications.

Add `snyk test` to your build or enable the fail functionality to make Snyk fail your builds, providing the results output to the console. Your developers or DevOps teams can use the results to decide whether to stop or continue the build.

For information about `snyk test`, see the [`test` command help](../../snyk-cli/commands/test.md).

### Stage 3: Continuous monitoring (`snyk test` and `snyk monitor`)

After you configure Snyk to fail the build when vulnerabilities are detected, you can configure Snyk to send a snapshot of your Project's successful builds to Snyk for ongoing monitoring.

To do this, configure your pipeline to run `snyk monitor` if your `snyk test` returns a successful exit code.

## CI/CD deployment methods

{% hint style="info" %}
All of these methods provide the same results, as they all rely on the same Snyk engine. Thus, the same arguments or options apply regardless of the deployment method you select.
{% endhint %}

There are various ways to configure Snyk within your pipeline. Choose a method depending on your environment and preferences. You can expect all methods to lead to a successful run.

### **Use Snyk native plugins**

Snyk native plugins are available for most common CI/CD tools. You can use these plugins to set up and get started. The plugins include the most common parameters and options in the user interface.

### **Deploy the Snyk CLI using the npm method**

Follow steps similar to those for [installing the CLI](../../snyk-cli/install-or-update-the-snyk-cli/) locally. You must be able to run an npm command in the pipeline script. This method has the advantage of completely aligning with the CLI experience so that you can easily troubleshoot and configure.

### **Deploy the Snyk CLI binary version**

The advantage of the binary setup is that it has no dependency with the local environment. For example, binary setup is useful if you cannot run an npm command in your pipeline.

CLI binaries are available on the [CLI GitHub repository](https://github.com/snyk/cli/tags).

Snyk has Linux, Windows, and other versions.

### **Deploy a Snyk container**

You may also deploy Snyk in your pipeline using one of the Snyk images in [Docker Hub](https://hub.docker.com/r/snyk/snyk).

## Examples of Snyk CI/CD Integrations

This repository shows some examples of binary and npm integrations for various CI/CD tools: [CI/CD examples](https://github.com/snyk-labs/snyk-cicd-integration-examples).
