# Security policy actions

An action defines what you want to happen when the [security policy conditions](security-policies-conditions.md) are matched.

{% hint style="info" %}
You cannot stack multiple actions in the same rule. To have multiple actions with a rule, create a new rule block with the same conditions, and specify a different action.
{% endhint %}

These are the actions that can currently be applied:

<table><thead><tr><th width="164">Action</th><th>Definitions</th></tr></thead><tbody><tr><td>Change severity to…</td><td><p>Changes the severity of whatever issues match the conditions. This can be set to <strong>Low</strong>, <strong>Medium</strong>, <strong>High</strong> or <strong>Critical</strong>.<br></p><p>Issues with a changed severity:</p><ul><li>Have their <a href="../../priorities-for-fixing-issues/priority-score.md">priority score </a>updated to reflect the new severity.</li><li>A note appears on the <a href="../../snyk-projects/issue-card-information.md">issue card</a> indicating that the issue's severity has been changed by a policy.</li><li>The severity icon will also be "stacked", showing the original severity behind the new one.</li></ul></td></tr><tr><td>Ignore current and future instances</td><td><p>Ignore all vulnerabilities that match the conditions. For example, ignore all issues that have no known exploits in projects with a <strong>business criticality</strong> attribute of <strong>low</strong>.</p><p></p><p>After an ignore policy is applied, ignores will happen every time the relevant Project is re-scanned, with ignored issues marked as <strong>ignored by Security Policy</strong>.</p><p></p><p>When setting the action, you can select <strong>won't fix</strong> and <strong>not vulnerable</strong> as ignore types, and add a reason, which appears on the issue card, alongside the ignore.</p><p></p><p>Policy-based ignores have the same behavior as issues that are manually ignored. Like manual ignores, automatic PRs are not raised on issues ignored by a security policy, or included in the issue count in reporting.</p></td></tr></tbody></table>
