---
nav_context: agnostic
---

# Risk intelligence

The **Risk intelligence agent** scores the AI models in your inventory so you can judge their risk. Snyk tests each model independently with adversarial attacks, rather than relying on self-reported vendor claims.

Every model has a **Risk profile** tab showing what an attacker can make that model do, and how much it matters when they succeed. With risk intelligence, you can:

* Read a model's risk at the level of a specific attacker goal, or rolled up by risk category.
* Compare the models found in your repositories on one scale.
* Set policies that raise an issue when a model crosses a score you choose.
* Apply remediation advice matched to the kind of risk.

Scored models appear under **Inventory** > **Code** after you import the repositories you want AI-SPM to scan. To import repositories, visit [AI asset visibility](../ai-asset-visibility.md).

## Coverage and limitations

Snyk assesses risk at the model variant level and shares it across every version of that variant. For example, `claude-3-5-sonnet-20240620` and `claude-3-5-sonnet-20241022` are both versions of the `claude-3-5-sonnet` variant, so both carry the same scores and the same **Risk profile**.

If you use a version Snyk has not tested directly, Evo matches it to a tested model in the same family and indicates on the profile that the data comes from a related model. A family groups related variants: `claude-3-5-sonnet`, `claude-sonnet-4`, and `claude-sonnet-5` all belong to the `claude-sonnet` family.

When Evo cannot match a model in your inventory to a tested model, the result is no score rather than a score of zero, and the **Risk profile** tab shows no data.

Snyk publishes the score and the method behind it. The individual attack prompts and the per-goal success rates are not published.

## Model Risk Score

The Model Risk Score is a value from 0 to 1,000 that combines two things:

* Likelihood: how often an attack succeeds against the model. Snyk measures this by running the attack many times and recording the results.
* Impact: how serious the consequence is when that attack succeeds. Snyk assesses impact once per attacker goal, and it is the same for every model.

Multiplying the two is what makes scores comparable. An attack that succeeds often but causes little harm scores low, and so does a severe attack that almost never works. Only an attack that both succeeds often and causes real harm scores high.

A score of 600 represents the same amount of risk wherever it appears, so you can compare scores across attacker goals and set thresholds that mean the same thing everywhere.

{% hint style="info" %}
The Model Risk Score is not a percentage or a failure rate. A score of 300 does not mean the model failed 30% of its tests. Evo shows the resulting score rather than the underlying attack success rate.
{% endhint %}

{% hint style="info" %}
The Model Risk Score applies to AI models in Evo. It is a different measure from the Snyk platform Risk Score, which prioritizes Snyk Open Source and Snyk Container issues. The two share a 0 to 1,000 range but are not comparable. To learn about the platform score, visit [Risk Score](https://app.gitbook.com/s/BJO0IZx7zB6bOkotxQP2/fix/prioritize-issues-for-fixing/risk-score).
{% endhint %}

### Severity

Snyk maps every score to a severity:

| Severity | Score     |
| -------- | --------- |
| Low      | 0–249     |
| Medium   | 250–499   |
| High     | 500–749   |
| Critical | 750–1,000 |

These ranges are fixed. Because the score already accounts for both likelihood and impact, the same range applies to every attacker goal and every risk category.

Some attacker goals cannot reach **High** or **Critical**. When the consequence of an attack is limited, for example, a model revealing which model it is, the score cannot enter the **High** range no matter how often the attack succeeds. This is intended rather than a gap in coverage.

## Guardrails used in testing

Snyk tests every model with a baseline guardrail in place. Before each attack, the model receives standard system-prompt hardening, the kind of instruction-level defense a reasonable deployment applies.

Scores therefore describe a defended model. They are not worst-case figures for a model running with no protection. A high score means the attacks succeeded at that rate against a model that was defended.

Bias measurement is the one exception and runs without the hardening defense, because bias is a property of the model's output rather than something an attacker breaks through.

## Attacker goals and risk categories

Snyk organizes model risk by what goes wrong, not by how the attack is delivered. There are three levels:

* Risk category: the top level, for example, Data Exfiltration or Unsafe Content. Each answers one question about the model.
* Sub-category: a group of related attacker goals, for example, Credentials & Secrets.
* Attacker goal: one specific thing an attacker is trying to achieve, for example, PII extraction or Command execution. This is where Snyk measures the score.

Sub-category and risk category scores are calculated from the attacker goals they contain. A single high-risk attacker goal raises its parent scores without pulling them all the way up to its own level, so one severe result is neither hidden nor allowed to dominate.

### Attack surfaces

The taxonomy separates two kinds of attack, because they call for different defenses:

* Direct: the attacker controls the prompt. The user is the adversary, typing something designed to break the model. Jailbreaks and extraction attempts are direct attacks.
* Indirect: the attacker hides instructions in content the model reads while doing its job, such as a web page, a document, or a tool result. The user is the victim rather than the adversary, and often does not know an attack is happening. This is prompt injection.

You address direct attacks mostly by choosing a model with better safety training. You address indirect attacks mostly in your own architecture, because the shared root cause is prompt injection and no model resists it reliably.

### Framework mappings

Every attacker goal is cross-referenced to MITRE ATLAS, the OWASP Top 10 for Large Language Model Applications, the OWASP Top 10 for Agentic Security Initiative, and the NIST AI Risk Management Framework. If your organization already governs AI risk through one of those frameworks, you can locate a model risk finding inside it rather than tracking Snyk findings as a separate vocabulary.

To look up the identifiers for a specific attacker goal, visit [Framework mappings](framework-mappings.md).

### Code security

Snyk scores whether a model writes vulnerable code when asked normally, for example, code containing SQL injection, cross-site scripting (XSS), or path traversal.

Evo reports this score separately from the attack-resistance categories. No attacker is involved: the model writes insecure code in response to an ordinary request. Keeping it separate means a model's code quality does not distort its resistance to attack, or the reverse. Both matter, and Evo reports both.

## View a model's risk profile

1. Navigate to **Inventory** > **Code**.
2. Select a model.
3. Select the **Risk profile** tab.

The tab shows:

* The severity ranges, so you can read any score against them.
* A **Risk score** table listing each attacker goal, its score, and a description of what an attacker achieves if they succeed. The sub-category appears under the attacker goal name.
* An **Attacker goals** and **Risk categories** toggle. Attacker goals show the most specific view, and risk categories show the same data rolled up.

## Model metadata

Alongside its scores, each model profile carries identity and capability information: vendor, model family and variant, release date, parameters, context length, modalities, and license.

Two fields are useful for procurement policies. **License type** identifies whether a model is open source, proprietary, or research-only. **Cost tier** groups models by price. You can write a policy against either, for example, to flag any model that is not open source.

## Policies and issues

Scores describe a model. Policies decide what to do about them.

A policy watches one attacker goal or one sub-category and raises an issue when a model's score for it crosses a threshold. Snyk provides default policies, and you can create your own. To learn how Evo raises and tracks issues, visit [Policies & issues](../../platform-surfaces/policies-and-issues.md).

### Default policies

Snyk default policies target attacker goals, because that is the level where the risk is specific enough to act on and where remediation advice applies. They raise:

* a **High** issue from a score of 500
* a **Critical** issue from a score of 750

{% hint style="info" %}
A score in a severity range does not always raise an issue. If no default policy exists for that attacker goal at that severity, the score appears on the **Risk profile** tab with no issue attached. You can create your own policy against any attacker goal or sub-category at any threshold.
{% endhint %}

### Custom policies

Create a custom policy when you want a threshold Snyk does not provide by default, or when you want to act at the sub-category level.

Target a sub-category when its attacker goals are interchangeable for you: the same fix, the same owner, and the specific goal does not change what you do about it. Bias is the common example, where you care that a model shows bias rather than whether it is gender, race, or religion bias.

Avoid pairing an attacker goal with the sub-category that contains it. Both fire, and you get two issues for one problem.

## Remediation

Snyk does not change the model. Each policy carries advice on what to change on your side, and that advice falls into two groups.

Some risks are properties of the model itself, including harmful content, bias, and copyright reproduction. Your main option is to choose a model that scores lower on that goal. Input and output classifiers help as a second layer, because they filter the result rather than change the model.

The rest are risks you control through your architecture, including data exfiltration, command execution, destructive actions, and decision override. Model choice matters less here. Assume the model can be hijacked and limit what it can reach: least-privilege access, egress controls, execution sandboxing, and approval gates for consequential actions. These hold regardless of which model you run.

## How scores change over time

Snyk re-evaluates models as its testing evolves, so scores update. A model's score can move without anything changing in your environment.

The scale itself is stable. A score of 600 means the same thing this month as it did last month, and adding new models to the Risk Database does not shift existing scores.

## Getting help

If a score looks wrong for a model, contact Snyk Support with the model name and the attacker goal.
