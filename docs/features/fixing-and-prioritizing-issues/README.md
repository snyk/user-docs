# Fixing and prioritizing issues

Snyk goes beyond finding vulnerabilities and license compliance issues. Our priority scoring, reporting, and policy setting capabilities help you prioritize and fix the most critical issues first.

### Use priority scoring

The [Snyk Priority Score](starting-to-fix-vulnerabilities/snyk-priority-score.md) prioritizes issues based on a number of industry-standard criteria, including Kubernetes configuration data and signals from running containers.

![](<../../.gitbook/assets/image (86) (2).png>)

### Apply project attributes <a href="#h.r3thgse7qt7n" id="h.r3thgse7qt7n"></a>

Control prioritization at a granular level by applying [project attributes](../../getting-started/introduction-to-snyk-projects/view-project-information/project-attributes.md) such as:

* Lifecycle stage
* Business criticality
* Environment

### Organize issues

Snyk [reporting](../../introducing-snyk/snyks-core-concepts/reporting.md) keeps you up to date on the status of the issues, dependencies, and licenses you need, while letting you [ignore the issues](issue-management/ignore-issues.md) you don’t.

![](<../../.gitbook/assets/image (66) (4).png>)

### Assess reachability <a href="#h.ts3kx23p4m7p" id="h.ts3kx23p4m7p"></a>

Gauge risk by identifying whether a vulnerable function is being called by your application, with Snyk’s [reachable vulnerability](prioritizing-issues/reachable-vulnerabilities.md) scanning.

### Set security policies

Prioritize or de-prioritize specific vulnerabilities automatically with our customizable [security policies](security-policies/).

![](<../../.gitbook/assets/image (82) (3).png>)

### Fix the issues

After you decide issue priority, it's time to start [fixing the vulnerabilities](starting-to-fix-vulnerabilities/).

Snyk can generate automatic PRs for upgrades and suggest recommended fixes. See [remediate-your-vulnerabilities.md](issue-management/remediate-your-vulnerabilities.md "mention") for more information on how Snyk helps you maintain code security through patches and direct dependency upgrades.
