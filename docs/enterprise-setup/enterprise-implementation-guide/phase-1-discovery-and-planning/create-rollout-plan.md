# Create rollout plan

Every business is different. If your teams have already used security tools and are in a heavily compliance-focused industry, controls may be turned on more quickly. However, if security as part of development is new, rolling out tools and controls in phases is strongly suggested.

## Suggested onboarding approach

When introducing Snyk to your business, we suggest the following phased rollout after implementing access (using SSO), and configuring integrations.

### 1. Kick Off with a Pilot Team

Start by selecting a small group of engaged pilot users from:

* Application security teams
* Project teams building new applications
* Developers of business-critical applications

This allows you to:

* Thoroughly onboard initial users
* Gather feedback to refine processes
* Identify issues before the broader rollout
* Build success stories to promote Snyk

Typically within a large enterprise, importing everything via repository integration for visibility and working through the rollout with a small pilot team allows you to identify the best processes and ways to implement Snyk within your environment.

### 2. Gain visibility with Git repository integration

Next, set up Snyk integrations across your Git repositories to gain broad visibility into your security posture.&#x20;

{% hint style="warning" %}
Disable notifications before import if you've onboarded all your users to reduce noise.
{% endhint %}

Key advantages:

* Widespread scanning across your codebase
* Automatic scanning triggered on code changes
* Convenient way to gain coverage

### 3. Prioritize Key Applications

Have your pilot team focus on securing priority applications using targeted Snyk CLI scans.

Key advantages:

* Enhanced visibility into critical apps
* Fine-tuned CLI testing for precision
* Remove repo noise for a focused view

### 4. Expand Access

With priorities addressed and processes refined, start expanding access more broadly across teams.

This phased approach allows thoughtful onboarding while rapidly gaining visibility and control.

### 5. Turn on Gating

After the first month, gradually turn on gating measures.

* Pull Request/Merge Request Checks using criteria such as severity, and is fixable.
* Fail builds based on criteria such as High/Critical, CVSS, Mature Exploit (for Open Source), and others using the [Snyk Filter](https://github.com/snyk-labs/snyk-filter) plugin.

It's recommended to start with a few applications, especially during the pilot team phase, work through the processes then roll out more widely.



## Exception Handling

Ensure there is an exception process in place and users are aware. For example&#x20;

* If a pull request/merge request is failed by Snyk, let them know who is the Snyk admin who can override it.&#x20;
* Similarly, if Snyk fails in CI/CD, who can create an ignore rule or authorize it to progress or configure CI/CD to run without the Snyk test or set it to monitor only?

