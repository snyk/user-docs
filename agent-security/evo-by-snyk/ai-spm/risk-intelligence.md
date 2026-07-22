---
nav_context: agnostic
---

# Risk intelligence

The **Risk intelligence agent** scores the AI models in your inventory so you can judge their risk. Snyk tests each model independently across five security categories using hundreds of test suites, rather than relying on self-reported vendor claims.

## Risk index

The Risk index is a score from 0 to 1000 that measures the severity-weighted failure rate of a model across Snyk's security tests. Higher scores indicate greater risk. A model with critical vulnerabilities scores higher than a model with only low-severity issues, even at a similar pass rate.

The index is absolute. A Risk index of 300 represents a 30% weighted failure rate, and the score does not change when Snyk adds new models. As Snyk tests new models and updates its assessments, Risk index scores update automatically.

A Risk index of 300 or above in a category raises a high-severity issue for the model. To learn how Evo raises and tracks issues, visit Policies and issues.

## Security categories

Snyk scores each model in five categories:

* Insecure code generation: how often the model generates vulnerable code, for example, SQL injection, cross-site scripting (XSS), operating system command injection, and path traversal.
* Bias and discrimination: whether the model produces discriminatory or stereotyping content across gender, race, and religion.
* Attack reconnaissance: whether the model reveals system prompts or internal configuration through extraction and adversarial escalation.
* Safety guardrail bypass: how easily a user circumvents the model's safety guardrails to produce harmful or policy-violating content.
* Sensitive data exposure: whether the model reproduces personally identifiable information (PII) from its training data.
