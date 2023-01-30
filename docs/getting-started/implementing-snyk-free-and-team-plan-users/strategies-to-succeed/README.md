# Strategies to succeed

## Introduction

Snyk products have a wide range of flexible functions. This matters because companies tend to roll out Snyk in phases, and the way Snyk functions are used can change over time. For example, you may start using the [Git repository integration](../../walkthrough-code-repository-projects/) or [CI/CD integrations](../../../integrations/ci-cd-integrations/) immediately, but how you use these functions may change during your Snyk rollout.

### Successful approaches

After scan results start getting generated, you’ll want to deal with issues in a phased approach, especially if you have a large collection of applications which have not been tested.

We have found that successful customers take approaches based on Visibility, Prevention, and Addressing the Backlog.

#### Stage 1: Visibility

Roll out Snyk on an initial set of important projects, to get immediate awareness about the problems you need to be aware of. This allows you to fix the worst issues right away, using priority scores or filters to help focus on where to start.

See [Visibility](visibility.md) for more details.

#### Stage 2: Prevention

You will need to use a phased approach to addressing issues found during the Visibility stage.

Prevention is the process of adopting tools and enabling features to make sure the problem does not get any worse.

See [Prevention](prevention.md) for more details.

#### Stage 3: Addressing the backlog

After Visibility and Prevention, you can now address your backlog, by examining and managing lower-priority issues over time or through new development work.

See [Addressing the backlog](addressing-the-backlog.md) for more details.

#### Other tips

* Invite the rest of the teams after the initial piloting team has identified what works best.
* Start turning on the controls (for example, [PR Checks](../../../scan-application-code/run-pr-checks/)) after an agreed period of time.
* Allocate Snyk admin roles: Have a clear stakeholder or set of stakeholders who can force merge PRs, mark PR checks as successful in Snyk, and adjust build settings as needed. See [Snyk admin](../../../user-and-group-management/) for details.
* Define who can [Ignore issues](../../../manage-issues/issue-management/ignore-issues.md), if this function is limited to Snyk Admins.

### What's next?

Let's start with gaining [Visibility](visibility.md).
