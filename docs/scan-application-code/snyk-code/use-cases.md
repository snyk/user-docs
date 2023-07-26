# Snyk Code - Benefits and Capabilities

## Identify and fix issues as they emerge

### Review newly introduced issues in build and review time

You can use Snyk Code in your IDE or in your SCM (using the [Automatic PR Checks](introducing-snyk-code/broken-reference/) feature) to identify issues as they are introduced into your code. Snyk Code does not require compilation, and it analyzes the source code itself, so you can see the results immediately as you write new code.

### Understand how problems flow across your applications

Some issues are one-liners, but they are not the majority. Most [SAST](https://snyk.io/learn/application-security/sast-vs-dast/) issues are multi-steps, multi-files, and sometimes multi-languages. To best represent that, Snyk Code presents a full data-flow visualization, which allows you to navigate through your source code from Source (user input) to Sink (the operation that needs to receive clean input and could be otherwise exploited).

### Fix more issues by using the experience of others

When Snyk Code identifies an issue, it includes real-world fix examples based on issues identified in similar patterns and data-flows in other projects. These examples provide relevant inspiration on how to fix the issue and minimize expensive research time.

### Learn how to prevent issues from coming back

Snyk Code provides a curated overview on every issue of the vulnerability - this includes how it is created, what the risk is, what the possible mitigation strategies are, and other bite-size, educational content. This allows developers to improve their security knowledge and write secure code in real-time.

## Reduce security debt over time

### Monitor your repositories to track issues over time

You can seamlessly and continuously monitor your repositories and discover security flaws in the source code files using static analysis provided by Snyk Code test scans.

### Prioritize the right issues to focus on

Every issue identified by Snyk Code is assigned with a Priority Score which reflects prevalence, risk, and estimated effort. This allows you to focus on the issues that present the greatest risk to your code.

In addition, you can group issues according to vulnerability, type, or file to save time and fix several issues at once.

### See what matters with smart exclude and ignore mechanisms

Snyk Code uses a smart fingerprinting algorithm that won't resurface previously-ignored issues in case you moved some code around or changed variable or function names.

Code also allows you to exclude files and folders from being analyzed, allowing you to filter the noise and focus on the code that matters most to your team.

## Scale and govern security processes

### Easily onboard Code across teams and departments

Remove the code security triage bottleneck from security teams by enabling automatic triage and fix in every step of the SDLC. Shift-left and enable developers to “self serve” on security.

Using Contributing Developers you can see how many developers are using Snyk Code.

### Review and share progress using reporting

Improving your company's security posture is an ongoing process and you need the right reporting to be able to answer questions like:

* How many issues have been **fixed** this month?
* How many issues have been **introduced** this month?
* How many issues have been **ignored** this month?

### Build automation to Gate deployments and extend Code

You can make sure no high severity code issues reach production by integrating the Snyk CLI with your CI/CD processes and failing builds based upon a certain, configurable criteria.

You can also use our public APIs to manage projects, query analysis findings, and build automatic workflows on top of them.
