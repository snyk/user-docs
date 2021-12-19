# Use Cases

## Reduce security debt over time

### Automatically monitor your repositories to track issues over time

You can seamlessly and continuously monitor your repositories, and discover security flaws in the source code files in them using static analysis.

This integration allows you to:

* Manage Code projects using your existing native import flow and tools
* View and prioritize security issues found in the source code
* Run a retest of a project and see history snapshots of a project

### Prioritize the right issues using priority score and issue grouping

Every issue identified by Snyk Code is assigned with a Priority Score which reflects prevalence, risk and estimated effort and allows you to focus on what matters.

In addition, you can group issues by vulnerability type or file to save time and fix several issues at once.

### Only scan what matters with smart exclude and ignore mechanisms

Snyk Code uses a smart fingerprinting algorithm that won't resurface previously-ignored issues in case you moved some code around or changed variable or function names.

Code also allows you to exclude files and folders from being analyzed, allowing you to filter the noise and focus on the code that matters.

## Identify and fix issues as they emerge

### Review newly introduced issues in build and review time

You can use Code in the IDE or in a PR check (currently in beta) to identify issues as they are introduced. Snyk Code doesn't require compilation and analyzes the source code itself, so you can see results immediately as you write new code.

### Quickly understand how problems flow across your applications

Some issues are one-liners, but they are not the majority. Most SAST issues are multi-step, multi-file and sometimes multi-language. To best represent that, Code presents a full data-flow visualization that allows you to navigate through your source code from source (user input) to sink (the operation that needs to receive clean input and can be exploited otherwise).

### Fix more issues by utilizing the experience of others

When Snyk Code identifies an issue, it includes real-world fix examples based on issues identified in similar patterns and data-flows in other projects. These examples provide relevant inspiration for how to remediate the issue and save lots of expensive research time.

### Learn how to prevent issues from coming back

On every issue, Snyk Code provides a curated overview of the vulnerability which includes how it's created, what the risk is, what are possible mitigation strategies and other bite-size lessons and educational content. This allows developers to improve their security knowledge and write more secure code overtime.

## Scale and govern security processes

### Easily onboard Code across teams and departments

Remove triage bottleneck from security team by enabling triage and fix in every step of the SDLC. Shift left and enable developers to “self serve” on security.

Using Contributing Developers you can see how many developers are utilizing Snyk Code.

### Review and share progress using reporting

Improving your company's security posture is an ongoing process, and you need the right reporting to be able to answer questions like:

* How many issues have been **fixed** this month?
* How many issues have been **introduced** this month?
* How many issues have been **ignored** this month?

### Gate deployments by using the CLI from your CI/CD

You can make sure no high severity code flaws reach production by integrating the Snyk CLI with your CI/CD processes and failing builds upon a certain criteria.
