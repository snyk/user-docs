# Framework mappings

Map Snyk findings to your existing governance frameworks. Snyk cross-references each identified attacker goal with four industry frameworks. This allows you to align security findings with the controls, control families, or risk categories your governance process already tracks.

## The frameworks

* MITRE ATLAS, version 5.6.0. Adversarial tactics and techniques against AI systems. Identifiers take the form `AML.Txxxx`, with sub-techniques as `AML.Txxxx.yyy`.
* OWASP Top 10 for Large Language Model Applications, 2025. Identifiers take the form `LLMxx`.
* OWASP Top 10 for Agentic Applications, 2026, from the Agentic Security Initiative. Identifiers take the form `ASIxx`.
* The NIST AI Risk Management Framework, does not include risk identifiers. Mapping therefore use two NIST risk taxonomies: `AI600-1:2.N` refers to a subsection of section 2 of NIST AI 600-1, covering content, safety, and privacy risks. `NISTAML.NNN` refers to an attack class in NIST AI 100-2 E2025.

## How the mappings relate

Three properties of the mapping explain most of the gaps you see in the tables.

Every indirect attacker goal shares indirect prompt injection as its delivery vector, so every indirect row also maps to `AML.T0051.001`, `LLM01`, and `NISTAML.015`. The indirect table lists only the additional identifiers specific to each goal.

There is no direct equivalent shared vector. `AML.T0051.000`, direct prompt injection, describes an attacker overriding model safety behavior to achieve their goals. This vector does not apply to every direct goal, so Snyk lists it per row instead of factoring it out.

The OWASP Top 10 for LLM Applications has no harmful-content category, because it is a security taxonomy rather than a content-safety one. Where a harmful-content goal lists `LLM01`, that identifier describes the jailbreak or injection used to bypass the model's guardrails, not the goal itself. Those rows are marked with an asterisk. The goal itself maps to NIST AI 600-1 and to the MITRE ATLAS external-harms and jailbreak techniques.

{% hint style="info" %}
Insecure code generation is intentionally unmapped to MITRE ATLAS. No adversary operates against the model, so no ATLAS technique applies.
{% endhint %}

### Direct attack surface

On the direct surface, the user is the adversary and controls the prompt.

| Attacker goal                                                                                | MITRE ATLAS                                        | OWASP LLM      | OWASP ASI    | NIST AI 100-2 / 600-1                 |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------- | -------------- | ------------ | ------------------------------------- |
| Chemical & biological                                                                        | AML.T0054, AML.T0051.000, AML.T0048.002            | LLM01\*        | None         | NISTAML.04, AI600-1:2.1               |
| Cybercrime & intrusion                                                                       | AML.T0054, AML.T0051.000, AML.T0016.002, AML.T0102 | LLM01\*        | None         | NISTAML.04, AI600-1:2.9               |
| Harassment & bullying                                                                        | AML.T0054, AML.T0051.000, AML.T0048.003            | LLM01\*        | None         | NISTAML.04, AI600-1:2.3, AI600-1:2.11 |
| Harmful                                                                                      | AML.T0054, AML.T0051.000, AML.T0048                | LLM01\*        | None         | NISTAML.04, AI600-1:2.3               |
| Illegal activities                                                                           | AML.T0054, AML.T0051.000, AML.T0048.002            | LLM01\*        | None         | NISTAML.04, AI600-1:2.3               |
| Misinformation & disinformation                                                              | AML.T0054, AML.T0048.002                           | LLM01\*, LLM09 | None         | NISTAML.04, AI600-1:2.8               |
| Copyright reproduction                                                                       | AML.T0054, AML.T0057                               | LLM02          | None         | AI600-1:2.10                          |
| Gender bias                                                                                  | AML.T0048.002                                      | None           | None         | AI600-1:2.6                           |
| Race bias                                                                                    | AML.T0048.002                                      | None           | None         | AI600-1:2.6                           |
| Religion bias                                                                                | AML.T0048.002                                      | None           | None         | AI600-1:2.6                           |
| Insecure code generation                                                                     | None                                               | LLM05          | ASI05        | NISTAML.027, AI600-1:2.9              |
| System prompt extraction                                                                     | AML.T0051.000, AML.T0056, AML.T0069.002            | LLM07          | None         | NISTAML.035, AI600-1:2.9              |
| Tool extraction                                                                              | AML.T0051.000, AML.T0084.001                       | LLM07          | None         | AI600-1:2.9                           |
| Model identification                                                                         | AML.T0014, AML.T0069                               | LLM07          | None         | AI600-1:2.9                           |
| PII extraction                                                                               | AML.T0051.000, AML.T0057                           | LLM02          | None         | NISTAML.038, AI600-1:2.4              |
| Role / purpose deviation                                                                     | AML.T0054, AML.T0051.000                           | LLM01          | ASI01        | NISTAML.018, AI600-1:2.7              |
| Unauthorized tool use                                                                        | AML.T0051.000, AML.T0053                           | LLM06          | ASI02, ASI03 | AI600-1:2.9                           |
| \* `LLM01` describes the elicitation vector rather than the goal. See the preceding section. |                                                    |                |              |                                       |

### Indirect attack surface

On the indirect surface, the attacker hides instructions in content the model reads while doing its job, and the user is the victim. Every goal in the following table also maps to `AML.T0051.001`, `LLM01`, and `NISTAML.015`.

| Attacker goal                            | MITRE ATLAS                         | OWASP LLM | OWASP ASI    | NIST AI 100-2 / 600-1                 |
| ---------------------------------------- | ----------------------------------- | --------- | ------------ | ------------------------------------- |
| Credential / secret theft                | AML.T0098, AML.T0055, AML.T0086     | LLM02     | ASI01, ASI03 | NISTAML.039, AI600-1:2.9              |
| System prompt extraction (via injection) | AML.T0056, AML.T0069.002            | LLM07     | ASI01        | NISTAML.035, AI600-1:2.9              |
| Source code exfiltration                 | AML.T0086, AML.T0048.004            | LLM02     | ASI01        | NISTAML.039, AI600-1:2.10             |
| Private data leakage                     | AML.T0057, AML.T0086, AML.T0077     | LLM02     | ASI01        | NISTAML.036, AI600-1:2.4              |
| Command execution                        | AML.T0050, AML.T0053, AML.T0102     | LLM05     | ASI01, ASI05 | AI600-1:2.9                           |
| Data / file destruction                  | AML.T0053, AML.T0101                | LLM06     | ASI01, ASI02 | AI600-1:2.9                           |
| Other unauthorized actions               | AML.T0053, AML.T0048.000            | LLM06     | ASI01, ASI02 | AI600-1:2.9                           |
| Code backdoors                           | AML.T0102                           | LLM05     | ASI05, ASI09 | NISTAML.027, AI600-1:2.9              |
| Supply chain poisoning                   | AML.T0010, AML.T0011.001, AML.T0060 | None      | ASI04        | AI600-1:2.12                          |
| Security control degradation             | AML.T0053, AML.T0081                | LLM06     | ASI01, ASI02 | AI600-1:2.9                           |
| Decision override                        | AML.T0067, AML.T0031                | LLM09     | ASI01, ASI09 | NISTAML.027, AI600-1:2.8              |
| Phishing / credential harvesting         | AML.T0052.000, AML.T0067            | None      | ASI09        | AI600-1:2.9                           |
| Trust exploitation                       | AML.T0067, AML.T0073                | LLM09     | ASI09        | NISTAML.027, AI600-1:2.7, AI600-1:2.8 |

{% hint style="info" %}
Learn how Snyk scores attacker goals and rolls up scores in [Risk intelligence](./).

Raise an issue when a model exceeds a selected score in [Policies & issues](../../platform-surfaces/policies-and-issues.md).
{% endhint %}
