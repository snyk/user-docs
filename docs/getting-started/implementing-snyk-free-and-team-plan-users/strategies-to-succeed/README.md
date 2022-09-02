# Strategies to succeed

### Before you start

Snyk has a lot of capabilities and is very flexible; this matters because companies tend to roll out Snyk in phases.

While you may start using the [Git repository integration](../../walkthrough-code-repository-projects/) or [CI/CD integrations](../../../integrations/ci-cd-integrations/) immediately, how you use it will change in the early stages of using Snyk.&#x20;

Similarly, after scan results start getting generated, you’ll want to deal with issues in a phased approach, especially if you have a large collection of applications which have not been tested.

### Successful approaches

We have found the most successful customers:&#x20;

* Pilot Snyk on an initial set of important projects to get [**visibility**](visibility.md) right away on the problems you need to be aware of
  * Start fixing the worst issues right away (using priority score or filters to help focus where to start).
  * If you find a lot of issues, you will use a phased approach to addressing them:
    * ****[**Prevention**](prevention.md): Adopt tools and enable features to make sure the problem does not get any worse.
    * ****[**Addressing the backlog**](addressing-the-backlog.md): Address lower priority issues over time or through new development work.
* Invite the rest of the team after the initial piloting team has identified what works best.
* Start turning on the controls (Pull Request Checks/Break builds) after an agreed period of time.&#x20;
* Have a clear stakeholder or set of stakeholders who can force merge PRs, mark pull request checks as successful in Snyk (Snyk Admins), and adjust build settings as needed. Additionally, determine who can use the Ignore option, if this is limited to admins.

### What's next?

To address vulnerabilities, we will discuss the phased approaches of **visibility**, **prevention** and **addressing the backlog**, in addition to the relevant settings, in the following sections.
