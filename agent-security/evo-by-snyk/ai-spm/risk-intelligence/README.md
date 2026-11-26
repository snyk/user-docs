---
nav_context: agnostic
last_reviewed: 2026-09-16T00:00:00.000Z
last_reviewed_source: bulk-sync
description: How the Risk intelligence agent scores AI models using adversarial testing
---

# Risk intelligence

The **Risk intelligence agent** scores the AI models in your inventory so you can judge their risk. Snyk tests each model independently with benchmark tasks and adversarial attacks, rather than relying on self-reported vendor claims.

A model that Snyk has scored has a **Risk profile** tab showing where that model can fail, and how much it matters when it does. With risk intelligence, you can:

* Read a model's risk at the level of a specific attacker goal, or rolled up by risk category
* Compare models on one scale, including models that are not in your repositories, in the Global Risk Database
* Set policies that raise an issue when a model crosses a score you choose
* Apply remediation advice matched to the kind of risk

Scored models appear under **Inventory** > **Code** after you import the repositories you want AI-SPM to scan. To import repositories, visit [AI asset visibility](../ai-asset-visibility.md).

### Coverage and limitations

Snyk assesses risk at the model variant level and shares it across every version of that variant. For example, `claude-3-5-sonnet-20240620` and `claude-3-5-sonnet-20241022` are both versions of the `claude-3-5-sonnet` variant, so both carry the same scores and the same **Risk profile**.

If you use a version Snyk has not tested directly, Evo matches it to a tested model in the same family and indicates on the profile that the data comes from a related model. A family groups related variants: `claude-3-5-sonnet`, `claude-sonnet-4`, and `claude-sonnet-5` all belong to the `claude-sonnet` family.

When Evo cannot match a model in your inventory to a tested model, the result is no score rather than a score of zero, and Evo shows no risk data for that model.

Snyk does not score every risk category for every model, and it adds categories over time. A category without a score has not been tested for that model, which does not mean the model passed.

Every score describes behavior under the conditions Snyk tested. It does not establish the probability of harm in your deployment, so treat a score as a lower bound on the real risk.

Snyk publishes the score and the method behind it. It does not publish the individual attack prompts or the per-goal success rates.

### Model Risk Score

The Model Risk Score is a value from 0 to 1,000 that combines two things:

* Likelihood: how often the model fails a test, for example when an attack succeeds or the model takes an unsafe action. Snyk measures this by running each test repeatedly across many tasks and averaging the results.
* Impact: how serious the consequence is when the model fails. Snyk assesses impact once per attacker goal, and it is the same for every model.

Multiplying the two is what makes scores comparable. A failure that happens often but causes little harm scores no higher than **Medium**, and a severe failure that almost never happens scores low. Only a failure that both happens often and causes real harm scores high.

Scores share one 0 to 1,000 scale, so you can compare scores across attacker goals and set thresholds that apply to all of them.

Several tests can measure the same attacker goal. Snyk measures each test separately and gives every test equal weight in the goal's score. Snyk leaves a test with no result out of the score.

{% hint style="info" %}
The Model Risk Score is not a percentage or a failure rate. A score of 300 does not mean the model failed 30% of its tests. Evo shows the resulting score rather than the underlying success rate.
{% endhint %}

{% hint style="info" %}
The Model Risk Score applies to AI models in Evo. It is a different measure from the Snyk platform Risk Score, which prioritizes vulnerabilities in Snyk Open Source and Snyk Container. The two share a 0 to 1,000 range but are not comparable. To learn about the platform score, visit [Risk Score](https://docs.snyk.io/scan-fix-and-prevent/fix/prioritize-issues-for-fixing/risk-score).
{% endhint %}

#### Severity

Snyk maps every score to a severity:

| Severity | Score     |
| -------- | --------- |
| Low      | 0–249     |
| Medium   | 250–499   |
| High     | 500–749   |
| Critical | 750–1,000 |

These ranges are fixed. Because the score already accounts for both likelihood and impact, the same range applies to every attacker goal and every risk category.

Not every attacker goal can reach every severity. When the consequence of a failure is limited, the score has a ceiling. A model that reveals its system prompt, for example, scores no higher than **Medium**, however often it does so. Only a few attacker goals can reach **Critical**. This is by design, not a gap in coverage.

A score's severity range does not limit the issues you raise. If a failure matters more in your environment than its score suggests, create a custom policy for that attacker goal with a lower threshold and the issue severity you choose, including **Critical**. To learn how, visit [Policies and issues](../../platform-surfaces/policies-and-issues.md).

### Guardrails used in testing

Snyk runs its prompt injection and jailbreak attacks against every model with a baseline guardrail in place. Before each attack, the model receives standard system-prompt hardening.

Scores from these attacks therefore describe a defended model. They are not worst-case figures for a model running with no protection. A high score means the attacks succeeded against the guarded model often enough to carry real impact.

Bias measurement is the one exception and runs without the hardening defense.

### Attacker goals and risk categories

Snyk organizes model risk by what goes wrong, not by how the failure is triggered. There are three levels:

* A risk category is the top level, for example, Privacy & Confidentiality or Code Security.
* A sub-category is a group of related attacker goals, for example, External Data Exfiltration.
* An attacker goal is one specific risk that Snyk assesses through defined tests, for example, Sensitive Data Leakage or Unauthorized Command Execution. This is where Snyk measures the score.

An attacker goal is not always something an attacker does. Some tests use an adversary, such as a user who tries to jailbreak the model, or instructions hidden in a web page or a tool result that the model reads. Others use an ordinary request, such as asking the model to write code, and check whether the result is safe. A model that writes vulnerable code in response to an ordinary coding task fails a test without anyone attacking it.

Snyk calculates each sub-category and risk category score from the scores beneath it. One high-risk score raises its parent without pulling it up to its own level, unless it is the parent's only score. Snyk leaves out a missing score rather than counting it as zero.

A model's overall score works the same way across every risk category Snyk has scored, including Code Security. To see overall scores and learn how to read them, visit Global Risk Database.

#### Framework mappings

Snyk cross-references attacker goals to MITRE ATLAS, the OWASP Top 10 for Large Language Model Applications, the OWASP Top 10 for Agentic Applications, and NIST publications on AI risk. If your company already governs AI risk through one of those frameworks, you can locate a model risk finding inside it rather than tracking Snyk findings as a separate vocabulary.

A mapping links related concepts. It does not establish that a model, or your use of it, complies with a framework.

To look up the identifiers for a specific attacker goal, visit [Framework mappings](framework-mappings.md).

### View a model's risk profile

1. Navigate to **Inventory** > **Code**.
2. Select a model.
3. Select the **Risk profile** tab.

The tab shows:

* the severity ranges, so you can read any score against them
* a **Risk score** table listing each attacker goal, its score, and a description of what goes wrong if the model fails, with the sub-category under the attacker goal name
* an **Attacker goals** and **Risk categories** toggle, where attacker goals show the most specific view and risk categories group them

### Model metadata

Alongside its scores, each model profile carries identity and capability information: vendor, model family and variant, release date, parameters, context length, modalities, and license.

Two fields help with procurement. **License type** identifies the model's license category, such as open source, proprietary, or research-only. **Cost tier** groups models by price. You can write a policy against a model's license, for example, to flag any model with a license you do not allow.

### Policies and issues

Scores describe a model. Policies decide what to do about them.

A policy raises an issue when a model matches its conditions. A model risk condition checks the score of one attacker goal or one sub-category against a threshold. You can add conditions on other attributes, such as the model's license or Epoch Capabilities Index (ECI) score. Snyk provides default policies, and you can create your own. To learn how Evo raises and tracks issues, visit [Policies and issues](../../platform-surfaces/policies-and-issues.md).

#### Default policies

Snyk default policies target attacker goals. Each default policy raises a **High** issue from a score of 500, a **Critical** issue from a score of 750, or both, depending on the attacker goal.

You cannot edit or delete default policies, but you can disable them for your whole Tenant. To learn how, visit [Disable or enable a policy](../../platform-surfaces/policies-and-issues.md#disable-or-enable-a-policy).

{% hint style="info" %}
A score in a severity range does not always raise an issue. If no default policy exists for that attacker goal at that severity, the score appears on the **Risk profile** tab with no issue attached. You can create your own policy against any attacker goal or sub-category at any threshold.
{% endhint %}

#### Custom policies

Create a custom policy when you want a threshold Snyk does not provide by default, or when you want to act at the sub-category level.

Target a sub-category when its attacker goals are interchangeable for you: they share a fix and an owner, and the specific goal does not change what you do about it. Bias is the common example, where you care that a model shows bias rather than whether it is gender, racial, or religious bias.

Avoid creating one policy for an attacker goal and another for the sub-category that contains it. Evo raises an issue for each policy, so one problem raises two issues.

### Remediation

Snyk does not change the model. Each policy carries advice on what to change on your side, and the advice depends on the risk:

* For harmful content, copyright, and misinformation, choose a model that scores lower, then add input and output safeguards and review of consequential output.
* For bias, compare models on cases that match your users, remove unnecessary demographic cues, apply consistent decision criteria, and review consequential outcomes independently.
* For data disclosure, data leakage, command execution, destructive actions, and sandbox escape, limit what the model can reach: keep sensitive data out of its context, use task-scoped access, restrict outbound destinations, isolate execution, and require approval for irreversible actions.
* For supply chain tampering, security control weakening, and unauthorized transactions, limit what the model can change: allow only approved packages and registries, route changes to security settings through independent review, and bind approvals to the exact transaction.
* For decision manipulation, keep authorization and decision criteria outside model-generated text, and require independent review of consequential decisions.
* For code security risks, check what the model produces: run functional tests, static analysis, and dependency checks, require independent review of sensitive changes, and protect your review and build gates from changes by the model.

### How scores change over time

Snyk re-evaluates models as its testing evolves, so scores update. A model's score can move without anything changing in your environment.

Within a taxonomy version, the scale is stable. A score of 600 means the same thing this month as it did last month, and adding new models to the Global Risk Database does not shift existing scores.

Snyk also adds risk categories over time. A model's overall score reflects only the categories Snyk has scored for that model, so it can change when Snyk scores a new category.

The risk taxonomy changed when Snyk moved to Risk Taxonomy v3. Snyk combined some attacker goals, so scores under the earlier taxonomy are not directly comparable with scores under the current one.

### Getting help

If a score looks wrong for a model, contact Snyk Support with the model name and the attacker goal.
