---
description: Compare every AI model Snyk has scored, by risk and capability
---

# Global Risk Database

## Global Risk Database

The Global Risk Database is a catalog of every AI model that Snyk has scored, whether or not the model appears in your repositories. Use it to compare models on risk and capability before you adopt one, and to see how the models in your code compare with the alternatives.

Snyk tests each model in a secure harness with benchmark tasks and adversarial attacks, and scores the results. To learn how scoring works, visit [Risk intelligence](./). The Risk profile tab on a model in your inventory covers models found in your code. The Global Risk Database covers all scored models.

With the Global Risk Database, you can:

* Compare models on risk against capability.
* Find models by name, provider, or cost tier.
* Compare two to five models side by side, category by category.
* Read the full risk profile of any scored model, including models you do not use.

The Global Risk Database contains models only. MCP servers and skills are not included.

### Prerequisites

Your Tenant must have Evo AI-SPM. Anyone who can open Evo in the Tenant can open the Global Risk Database.

To open the page, navigate to **Research** > **Global Risk Database**.

### Weigh risk against capability

A more capable model is not always a riskier one, and a lower-risk model is not always less capable. The Global Risk Database plots both, so you can see the trade-off.

Capability comes from the Epoch Capabilities Index (ECI), a third-party measure of how capable a model is. Risk comes from the overall risk score, which aggregates a model's [risk scores](./#model-risk-score). Snyk calculates risk scores from its own testing, on a scale of 0 to 1,000. The chart labels the overall score as the aggregated risk score. Snyk maps every score to a severity:

| Severity | Score     |
| -------- | --------- |
| Low      | 0–249     |
| Medium   | 250–499   |
| High     | 500–749   |
| Critical | 750–1,000 |

Read the two together. Between two models with the same ECI, the one with the lower score carries less risk for the same capability. To check the models you already use, limit the view to the variants in your repositories.

The chart uses the overall risk score, so treat it as a starting point. Snyk measures risk per attacker goal, and the goal scores are the evidence behind every overall score. Before you rely on one, review the attacker goal scores, as described in [Read a score](global-risk-database.md#read-a-score).

Snyk scores models by variant. A variant groups the versions of a model that behave alike. Snyk tests one version and shares the result across the variant. Variants that have no ECI or no risk score do not appear on the chart. To learn how Evo handles versions that Snyk has not tested, visit [Coverage and limitations](global-risk-database.md#limitations).

### Compare models

Compare two to five models to see where their risk differs, not only which score is lower. The comparison breaks each model's score into risk categories, and into the attacker goals within each category, which is where Snyk measures risk.

Two models with the same overall score can fail in different ways. The category view shows which risks matter for your use.

Each model's page shows the same scores in more detail, along with its capabilities: the tasks it supports, its modalities, and its context length. For a model that your Tenant uses, the page also lists the repositories that use the model and the issues that Snyk raised.

### Cost tiers

Snyk groups models into five cost tiers so you can filter by what a model costs to run. For text models, the tier reflects the price per 1 million tokens, blended at three parts input to one part output.

| Tier       | Price per 1 million tokens |
| ---------- | -------------------------- |
| Free       | No per-token charge        |
| Low        | Under $0.50                |
| Medium     | $0.50 to under $3          |
| High       | $3 to under $15            |
| Enterprise | $15 or more                |

Models without a published price show no tier.

### Read a score

The overall risk score is an entry point. Use it to shortlist and rank models, then check the attacker goals before you decide.

Snyk measures risk at the attacker goal level. An attacker goal is a specific risk that Snyk assesses through defined tests, such as a model revealing personal data. The Model Risk Score for a goal combines likelihood, how often the model fails the test, with impact, how serious the consequence is when it does. Snyk groups related attacker goals into categories, and the overall risk score aggregates all of them into one number.

Because the overall score aggregates, it can hide a serious weakness. A model can have a moderate overall score and still have a High or Critical severity on one attacker goal. If that goal matters for your use, the overall score understates your risk.

A low score also has more than one meaning. A model can resist the attacks, or a successful attack can carry limited impact. The attacker goal descriptions explain what is at stake.

The overall score reflects only the categories that have data. A model can have a score in some categories and none in others. A score of 0 in a category means Snyk tested the model and the attacks did not succeed.

Scores reflect the attacks that Snyk tested, so treat them as a lower bound on the real risk. Scores change when Snyk re-tests a model or adds a risk category. Rankings change as Snyk tests more models, but existing scores do not. To learn how Snyk calculates scores, visit [Risk intelligence](./).

### Ask Evo about models

You can ask [Evo chat](../../platform-surfaces/evo-chat.md) about the models in the Global Risk Database. Example prompts include:

* "List the 5 models in the Global Risk Database with the highest overall risk score."
* "List every OpenAI model in the Global Risk Database, including models we don't use."

{% hint style="info" %}
Mention the Global Risk Database when you want models beyond the ones in your repositories.
{% endhint %}

### Limitations

* The Global Risk Database contains models only.
* Snyk has not tested every model. A model that is not in the database has no score. A missing score does not mean the model is low risk.
* Snyk re-tests models and adds risk categories over time, which can change scores. Snyk also adds models, which can change rankings but not existing scores.
