# Security policy actions

An action defines what you want to happen when the [security policy conditions](security-policies-conditions.md) are matched.

{% hint style="info" %}
You cannot stack multiple actions in the same rule. To have multiple actions with a rule, create a new rule block with the same conditions, and specify a different action.
{% endhint %}

These are the actions that can currently be applied:

<table><thead><tr><th width="164">Action</th><th>Definitions</th></tr></thead><tbody><tr><td>Change severity to…</td><td><p>Changes the severity of whatever issues match the conditions. This can be set to <code>Low</code>, <code>Medium</code>, <code>High</code>, or <code>Critical</code>.<br></p><p>Issues with a changed severity have their <a href="../../prioritize-issues-for-fixing/priority-score.md">priority score </a>updated to reflect the new severity. A note appears on the <a href="../../../snyk-platform-administration/snyk-projects/issue-card-information.md">issue card</a> indicating that the severity of the issue has been changed by a policy. The severity icon will also be "stacked," showing the original severity behind the new one.</p></td></tr><tr><td>Ignore current and future instances</td><td><p>Ignore all vulnerabilities that match the conditions. For example, ignore all issues that have no known exploits in Projects with a <code>business</code> <code>criticality</code> attribute of <code>low</code>.</p><p></p><p>After an ignore policy is applied, ignores will occur every time the relevant Project is re-scanned, with ignored issues marked as <code>ignored by Security Policy</code>.</p><p></p><p>When setting the action, you can select <code>won't fix</code> and <code>not vulnerable</code> as ignore types, and add a reason, which appears with the ignore on the issue card </p><p></p><p>Policy-based ignores have the same behavior as issues that are manually ignored. As with manual ignores, automatic PRs are not raised on issues ignored by a security policy or included in the issue count in reporting.</p></td></tr></tbody></table>

{% hint style="info" %}
Security policies are applicable to Snyk Open Source and Snyk Container Projects.
{% endhint %}
