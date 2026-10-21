---
description: >-
  Create remediation campaigns in Evo by Snyk to fix open-source vulnerabilities
  across repositories and open pull requests on demand or on a schedule
---

# Remediation campaigns

{% hint style="info" %}
Remediation campaigns are in [Early Access](https://docs.snyk.io/snyk-release-process#early-access-features). Snyk enables the feature for each Tenant. To request access, contact your Snyk account team.
{% endhint %}

Use remediation campaigns in Evo by Snyk to fix open-source vulnerabilities across your repositories. Choose the repositories and vulnerabilities to target, then run a campaign once or on a schedule. The [Remediation Agent](../remediation-agent.md) generates and verifies fixes in Snyk-hosted infrastructure, then opens pull requests in your source code management (SCM) system for you to review and merge. Snyk provides the model, so there is no model provider or API key to set up.

## Prerequisites

To create a campaign, you must have:

* Evo by Snyk, with remediation campaigns enabled for your Tenant.
* Workspaces enabled. A Group or Organization admin can enable Workspaces on the **Integrations** settings page. For details, visit [Workspaces for SCM integrations](https://docs.snyk.io/developer-tools/integrations/scm-integrations/workspaces).
* Repositories imported into Snyk through an SCM integration, with Snyk Open Source Projects. Campaigns support GitHub, GitHub Enterprise, GitHub Cloud App, GitHub Server App, GitLab, Azure Repos, Bitbucket Cloud, Bitbucket Cloud App, and Bitbucket Data Center/Server. The integration needs write access to the repositories to open fix pull requests. For details, visit [Organization-level integrations](https://docs.snyk.io/developer-tools/integrations/scm-integrations/organization-level-integrations).
* Access to Evo by Snyk and to the Organization that owns the repositories. If Snyk needs to install its App in the Organization, you also need permission to install Apps in that Organization. The **Org Admin** role includes this permission.
* Repositories that use a [supported package manager](#supported-package-managers).

Campaigns do not support private package registries. For details, see [Limitations](#limitations).

## Supported package managers

Remediation campaigns support the following package managers and build tools:

| Package manager or build tool | Language |
|---|---|
| Maven | Java |
| npm | JavaScript |
| Yarn | JavaScript |
| NuGet | .NET |
| Gradle with the Groovy DSL | Java |
| Gradle with the Kotlin DSL | Java |
| Go Modules | Go |
| pnpm | JavaScript |
| sbt | Scala |
| Bundler | Ruby |

Bundler support is limited to relock operations.

## How campaigns work

A campaign defines what to fix. A campaign run carries out that instruction, with a separate repo run for each repository that has vulnerabilities to fix.

| Term | Meaning |
|---|---|
| Campaign | Defines the repositories, vulnerabilities, and run schedule. Each campaign covers one Organization. |
| Campaign run | One execution of a campaign. It starts a repo run for each repository in scope that has vulnerabilities to fix. |
| Repo run | Fixes vulnerabilities in one repository in an isolated sandbox. It combines its fixes into one pull request against the repository's default branch. |

Choose a campaign type based on when you want the agent to run:

| Campaign type | When it runs | Use it to |
|---|---|---|
| **One-off** | When you create it. To run it again, click **Run now** on the campaign. | Address a specific set of vulnerabilities, such as a newly disclosed critical vulnerability. |
| **Regular campaign** | On your schedule, until you pause it or it reaches its end condition. | Fix vulnerabilities across the same repositories at regular intervals. |

### What a repo run skips

Before starting a repo run, Snyk excludes:

* Vulnerabilities that you have ignored in Snyk or that are no longer open.
* Vulnerabilities that Snyk already tracks in a pending or active repo run, or in an open campaign pull request, when it checks for existing work. These appear as **Skipped issues** on the campaign.
* Vulnerabilities that exceed the Tenant's [daily fix allowance](#daily-fix-allowance). These appear as **Over the daily allowance**.

### Repository instructions

The agent reads the `AGENTS.md` file at the root of each repository, or `CLAUDE.md` if there is no `AGENTS.md`. Use this file to tell the agent how to install, build, and test the project. Campaigns have no separate field for custom instructions.

### Breaking change assessment

Snyk assesses the breaking-change risk of proposed dependency fixes before applying them. If you select **Non-breaking fixes only**, the agent selects fixes assessed as low risk. It excludes fixes with a medium, high, or unknown risk assessment. This filter applies to each proposed fix, rather than rejecting an entire repository's pull request because one fix has a higher risk.

A low-risk assessment does not guarantee that a fix cannot break your build. Review the changes and validation results before merging. If no eligible fix produces a pull request, the repo run ends with **No fix produced**.

## Create a campaign

Snyk recommends starting with a few repositories. Review the pull requests before adding more repositories to the campaign. Each run is limited by the Tenant's [daily fix allowance](#daily-fix-allowance).

{% stepper %}
{% step %}
### Start a campaign

In Evo by Snyk, navigate to **Agentic Remediation** > **Campaigns** and click **New campaign**.
{% endstep %}

{% step %}
### Define remediation scope

1. Under **Scope**, click **Choose repositories** and select the repositories to include. After you select the first repository, the picker shows only repositories in the same Organization.
2. Under **Detection rules**, configure the vulnerabilities to fix:
   * **Issue type**: select **Open source (SCA)**.
   * **Issue severities**: select one or more severities. The campaign targets exactly these severities. For example, selecting **High** severity does not include **Critical** vulnerabilities.
   * **Breakability**: select **Non-breaking fixes only** or **Breaking allowed, flagged for review**. For details, see [Breaking change assessment](#breaking-change-assessment).

If another campaign covers the same repositories, the form shows **Already covered by** and the campaign name. Review the existing campaign before creating another with the same scope. Snyk checks for vulnerabilities already covered by tracked runs or open campaign pull requests, but campaigns that start close together can still attempt the same vulnerability.
{% endstep %}

{% step %}
### Set the run schedule

Select **Run once** for a one-off campaign, or **Regular campaign** to run on a schedule.

For a regular campaign, set:

* **Repeats every**: choose an interval in hours, days, weeks, or months. For a weekly schedule, select the days in **Repeats on**.
* **Starts**: start now or choose a date and time. You cannot change the start date or time after you create the campaign.
* **Ends**: select **Never, keep running**, **After** a number of campaign runs, or **On** a date.
{% endstep %}

{% step %}
### Review and create

For a regular campaign, enter a **Campaign name** and review **What to expect**. Click **Create campaign** to use a future start date, or **Create and run now** to start immediately.

For a one-off campaign, review **What to expect** and click **Run now**.
{% endstep %}
{% endstepper %}

## Daily fix allowance

During Early Access, your Tenant has a daily allowance for fixing vulnerabilities, shared across all its Organizations. **Fixes available** in the navigation shows the remaining allowance.

Each campaign run uses the remaining allowance, starting with the most severe vulnerabilities. Vulnerabilities beyond the allowance remain unfixed for that run. A later run can pick them up when allowance is available.

* Fixed vulnerabilities and vulnerabilities in progress count toward the allowance. If a repo run ends without fixing a vulnerability, Snyk restores that vulnerability's allowance.
* Scheduled runs still start when the allowance is exhausted, but they fix no vulnerabilities.
* The allowance resets daily at 12:00 AM UTC. Unused fixes do not carry over.
* A repo run counts against the allowance for the UTC day it starts, even if it finishes after the reset.
* You cannot increase the allowance.

## Review the results

Navigate to **Agentic Remediation** to review campaigns and their results:

| Page | What you can review |
|---|---|
| **Campaign runs** | Campaign runs, their repo runs, and run statuses. |
| **Campaigns** | Campaign scope, the next run, and the pull request merge rate. |
| **Pull requests** | Pull requests opened by campaigns and their status in your SCM system. |
| **Metrics** | Runs, vulnerabilities fixed, and pull requests merged over time. |

Open a repo run to read its **Agent journal**, which records the agent's actions step by step.

### Repo run outcomes

| Outcome | Meaning |
|---|---|
| **Fix produced** | The agent fixed at least one vulnerability and opened a pull request. |
| **No fix produced** | The repo run completed without an error but produced no fix it could safely include in a pull request. |
| **Failed** | An error stopped the repo run. Open the repo run to read the reason. |

### Pull request status

Snyk shows the most recent status it received from your SCM system:

* **Open**: the pull request is open.
* **Merged**: the pull request was merged.
* **Closed**: the pull request was closed without merging.

### What a pull request contains

The pull request description lists the vulnerabilities the agent fixed and did not fix. It also describes the changes, how the agent validated them, and the risk of breaking changes. AI generates parts of the description. Review the changes and validation results before merging.

The agent excludes new files and changes to dotfiles, such as `.npmrc`, from its pull requests.

## Edit or pause a campaign

* To change a campaign, open it and click **Edit**. Changes apply to future runs. Previous runs retain their original configuration.
* To stop scheduled runs, click **Pause** on a regular campaign. The campaign retains its configuration so you can resume it later.
* To start a run manually, click **Run now**. A campaign can have only one run in progress. If a run is already in progress, Snyk does not start another.

You cannot delete a campaign. Its runs and pull requests remain available.

## Differences from the Snyk CLI

Campaigns run the same agent as `snyk fix --agentic` in the [Snyk CLI](../remediation-agent.md#fix-vulnerabilities-from-the-snyk-cli). Choose the entry point that fits your workflow:

| Capability | Remediation campaigns | Snyk CLI |
|---|---|---|
| Where the agent runs | Snyk-hosted sandbox | Your machine or pipeline |
| Model configuration | Snyk supplies the model | You configure a provider and credentials |
| Approval | The agent proceeds automatically. You review its pull request. | Interactive, unless you use `--auto-approve` |
| Private package registries | Not supported | Supported when your local environment can reach them |
| Scan types | Snyk Open Source | Snyk Open Source, Snyk Code, and Snyk Container |


{% hint style="info" %}
* The agent fetches dependencies from public registries only. Repositories that depend on private registries can fail to build in the sandbox.
* The agent can fix breaking changes that require small code edits. It does not perform large migrations, such as a major framework upgrade.
* Campaigns do not disable [Snyk Fix PRs](../snyk-pull-or-merge-requests/README.md). If you use both for a repository, both can open pull requests for the same vulnerability.
{% endhint %}

## Data handling

Neither Snyk nor its third-party AI providers use your proprietary code to train or improve AI models. For details on how Snyk uses data with generative AI, visit [How Snyk incorporates generative AI into the platform](https://docs.snyk.io/snyk-data-and-governance/how-snyk-incorporates-generative-ai-into-the-platform).

Your campaign history, agent journals, and pull requests remain available to you in Snyk and your SCM system.

## Get help

Send feedback through your Snyk account team or [Snyk Support](https://support.snyk.io).
