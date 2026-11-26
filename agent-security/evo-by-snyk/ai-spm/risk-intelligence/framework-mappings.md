---
last_reviewed: 2026-08-21T00:00:00.000Z
last_reviewed_source: focused-edit
description: How Snyk maps attacker goals to four industry governance frameworks
---

# Framework mappings

Snyk cross-references each attacker goal it scores to industry frameworks, so you can locate a model risk finding inside a control framework you already use instead of treating it as a separate vocabulary.

### The frameworks

* MITRE ATLAS, version 5.6.0: adversarial tactics and techniques against AI systems. Identifiers take the form `AML.Txxxx`, with sub-techniques as `AML.Txxxx.yyy`.
* OWASP Top 10 for Large Language Model Applications, 2025: identifiers take the form `LLMxx`.
* OWASP Top 10 for Agentic Applications, 2026, from the Agentic Security Initiative: identifiers take the form `ASIxx`.
* NIST publications on AI risk: mappings target two NIST documents. `AI600-1:2.N` refers to a subsection of section 2 of NIST AI 600-1, which lists risks unique to or exacerbated by generative AI. `NISTAML.xx` is a Snyk label for a concept in NIST AI 100-2 E2025, not an identifier that NIST issues.

### How to read the mappings

A mapping connects an attacker goal to a related concept in a framework. It does not claim the two are the same risk.

Not every attacker goal has an entry in every framework. A cell that reads None means Snyk has not mapped a related entry.

{% hint style="info" %}
The mappings help you locate a finding. They do not establish that your use of a model complies with a framework.
{% endhint %}

### Attacker goals

#### Privacy & Confidentiality

| Attacker goal              | MITRE ATLAS                                        | OWASP LLM | OWASP Agentic (ASI) | NIST AI 600-1 / 100-2   |
| -------------------------- | -------------------------------------------------- | --------- | ------------------- | ----------------------- |
| Sensitive Data Disclosure  | AML.T0057                                          | LLM02     | None                | AI600-1:2.4             |
| Agent Internals Disclosure | AML.T0056, AML.T0084.001, AML.T0014, AML.T0069.002 | LLM07     | None                | AI600-1:2.9             |
| Sensitive Data Leakage     | AML.T0057, AML.T0086, AML.T0098                    | LLM02     | None                | AI600-1:2.4, NISTAML.03 |
| Agent Internals Leakage    | AML.T0056, AML.T0086                               | LLM07     | None                | AI600-1:2.9             |

#### System Integrity & Availability

| Attacker goal                       | MITRE ATLAS          | OWASP LLM | OWASP Agentic (ASI) | NIST AI 600-1 / 100-2               |
| ----------------------------------- | -------------------- | --------- | ------------------- | ----------------------------------- |
| Unauthorized Command Execution      | AML.T0050, AML.T0053 | LLM06     | ASI05               | AI600-1:2.9                         |
| Supply Chain & Dependency Tampering | None                 | None      | None                | AI600-1:2.9                         |
| Security Control Weakening          | AML.T0081            | LLM06     | ASI03               | AI600-1:2.9                         |
| Destructive Data Operations         | AML.T0101            | LLM06     | ASI02               | NISTAML.01, NISTAML.02, AI600-1:2.9 |
| Unauthorized Service Transactions   | None                 | LLM06     | ASI02               | AI600-1:2.9                         |

#### Code Security

| Attacker goal            | MITRE ATLAS | OWASP LLM | OWASP Agentic (ASI) | NIST AI 600-1 / 100-2    |
| ------------------------ | ----------- | --------- | ------------------- | ------------------------ |
| Code Backdoor Insertion  | None        | None      | None                | NISTAML.027, AI600-1:2.9 |
| Insecure Code Generation | None        | None      | None                | AI600-1:2.9              |

#### Content Safety & Integrity

| Attacker goal                    | MITRE ATLAS                         | OWASP LLM | OWASP Agentic (ASI) | NIST AI 600-1 / 100-2                 |
| -------------------------------- | ----------------------------------- | --------- | ------------------- | ------------------------------------- |
| Chemical & Biological Harm       | AML.T0048.002                       | None      | None                | NISTAML.04, AI600-1:2.1               |
| Cybercrime Assistance            | AML.T0102                           | None      | None                | NISTAML.04, AI600-1:2.9               |
| Harassment & Bullying            | AML.T0048.003                       | None      | None                | NISTAML.04, AI600-1:2.3, AI600-1:2.11 |
| General Harmful Content          | AML.T0048                           | None      | None                | NISTAML.04, AI600-1:2.3               |
| Illegal Activity Assistance      | AML.T0048.002                       | None      | None                | NISTAML.04, AI600-1:2.3               |
| Misinformation & Disinformation  | None                                | LLM09     | None                | AI600-1:2.8                           |
| Copyrighted Content Reproduction | None                                | None      | None                | AI600-1:2.10                          |
| Gender Bias                      | None                                | None      | None                | AI600-1:2.6                           |
| Religious Bias                   | None                                | None      | None                | AI600-1:2.6                           |
| Racial Bias                      | None                                | None      | None                | AI600-1:2.6                           |
| Decision Manipulation            | AML.T0067, AML.T0052.000, AML.T0073 | LLM09     | ASI09               | AI600-1:2.7, AI600-1:2.8              |

#### Agent Behavior & Control Risks

| Attacker goal          | MITRE ATLAS | OWASP LLM | OWASP Agentic (ASI) | NIST AI 600-1 / 100-2 |
| ---------------------- | ----------- | --------- | ------------------- | --------------------- |
| Sandbox Escape         | None        | None      | ASI03, ASI05        | AI600-1:2.9           |
| Instruction Violations | None        | LLM06     | ASI02               | None                  |

To learn how Snyk scores attacker goals, visit [Risk intelligence](./).
