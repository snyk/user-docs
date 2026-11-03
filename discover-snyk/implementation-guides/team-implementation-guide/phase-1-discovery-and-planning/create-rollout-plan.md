---
description: How to create a Snyk rollout plan paced to the security maturity and compliance needs of your team
nav_context: agnostic
---

# Create rollout plan

Every business is different. If your teams have already used security tools and are in a heavily compliance-focused industry, you can turn on controls relatively quickly. However, if security as part of development is new to your teams, Snyk strongly recommends rolling out tools and controls in phases.

## Suggested onboarding approach

When you are introducing Snyk to your business, Snyk recommends the following phased rollout after you configure integrations.

### 1. Kick off with a pilot team

Start by selecting a small group of engaged pilot users from:

* Application security teams (if applicable)
* Project teams building new applications
* Developers of business-critical applications

This allows you to:

* Thoroughly onboard initial users
* Gather feedback to refine processes
* Identify issues before the broader rollout
* Build success stories to promote Snyk

Typically, importing everything using a repository integration for visibility and working through the rollout with a small pilot team allows you to identify the best processes and ways to implement Snyk in your environment.

### 2. Gain visibility with SCM integration

Next, set up Snyk integrations across your Git repositories to gain broad visibility into your security posture.

{% hint style="info" %}
To reduce noise, disable notifications before import if you have onboarded all your users.
{% endhint %}

The key advantages of using this process are:

* Widespread scanning across your codebase
* Automatic scanning triggered on code changes
* A convenient way to gain coverage

### 3. Prioritize key applications

Have your pilot team focus on securing priority applications using targeted Snyk CLI scans.

The key advantages of using this process are:

* Enhanced visibility into critical apps
* Fine-tuned CLI testing for precision
* Reduced repo noise for a focused view

### 4. Expand access

With priorities addressed and processes refined, start expanding access more broadly across teams.

This phased approach allows thoughtful onboarding while rapidly gaining visibility and control.

### 5. Turn on gating

After the first month, gradually turn on gating measures.

* Pull request (PR) and merge request (MR) checks using criteria such as `severity` and `is fixable`.
* Fail builds based on criteria such as `High` or `Critical`, `CVSS`, `Mature Exploit` for Open Source and other criteria using the [Snyk Filter](https://github.com/snyk-labs/snyk-filter) plugin.

Snyk recommends starting with a few applications, especially during the pilot team phase. Work through the processes, and then roll out more widely.

## Exception handling

Ensure that an exception process is in place and that users are aware of it. For example:

* If Snyk fails a PR or MR, let users know which Snyk admin can override it.
* Similarly, if Snyk fails in CI/CD, let users know who can create an ignore rule, authorize it to progress, or configure CI/CD to run without the Snyk `test` or set it to `monitor` only.

