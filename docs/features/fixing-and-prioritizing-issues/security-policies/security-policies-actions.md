# Security policies: Actions

An action is what you want to happen when the conditions in the policy rule are matched. These are the actions that can currently be applied:

**Customize severity**

| Action | Definitions |
| :--- | :--- |
| Change severity to… | Changes the severity of whatever issues match the conditions. This can be set to Low, Medium, High or Critical. Issues that have a changed severity will have their priority score updated to reflect the new severity. A note will also appear on the issue card indicating that the issue's severity has been changed by a policy. The severity icon will also be "stacked", showing the original severity behind the new one. The issue’s priority score will also be recalculated, based on the new severity. |
| Ignore current and future instances | The ignore action ignores all vulnerabilities that match the conditions specified in the rule. It will ignore all existing vulnerabilities that match the conditions, and all future instances of it that are found. An example rule could be to ignore all issues that match a specific CVE. Or ignore all issues that have no known exploits in projects that have a "business criticality" attribute of "low" Once an ignore policy is applied, ignores will happen the next time the project is re-tested \(either via a manual retest, or automatically on the next test run which happens daily if the project is set to retest daily\). Once that's happened, the ignored issues will be marked as "ignored by Security Policy". When setting the action, you can select "won't fix" and "not vulnerable" as ignore types, and add a reason you'd like to appear alongside the ignore. This will show on the issue card in your project. Policy-based ignores have the same behaviour as issues that are manually ignored. Like manual ignores, automatic pull requests will not be raised on issues ignored by a security policy, or included in the issue count in reporting. |

{% hint style="info" %}
You cannot stack multiple actions in the same rule. To have multiple actions with a rule, create a new rule block with the same conditions, and specify a different action.
{% endhint %}

