# How Snyk incorporates generative AI into the platform

Snyk’s AI Trust Platform uses generative AI to enhance automation, efficiency, and innovation for developers and security teams. Snyk uses a mix of solutions, including proprietary, self-hosted models and third-party large language models (LLMs).

This document explains what generative AI technologies Snyk uses and how data flows through our systems. It also describes the measures we take to protect your data. The field of AI is changing quickly. As a result, the AI technologies we use may change when we introduce new features or update existing ones.

### Core principles &#x20;

Snyk places the utmost importance on data security and integrity.

* **No training on customer code**: Snyk does not use customer proprietary software code to train, optimize, fine tune or improve any of its AI models, and does not use or incorporate any third-party AI models into the platform unless they make the same commitments.
* **Contractual protection**: All of the AI functionality described in this document forms part of Snyk’s services. Your use of this functionality is governed by your existing agreements with Snyk and benefits from the same contractual protections. No separate in-service terms, addenda, or amendments to your existing agreements with Snyk are required.

### AI models&#x20;

Snyk uses multiple AI deployment strategies to balance performance, security, and data protection:

* **Proprietary / self-hosted models**: Snyk’s core generative AI model is proprietary and maintained entirely within our controlled environment. This model runs on dedicated infrastructure and powers our fundamental product functionality of identifying issues and proposing fixes to those issues. &#x20;
* **Hybrid models**: For certain products or features, Snyk uses both its proprietary self-hosted model and open-source model/s. In this case, these open-source models are hosted and maintained entirely within our controlled environment. &#x20;
* **Third-party LLMs**: For certain products or features, Snyk uses LLMs from established AI providers, including OpenAI and Anthropic, through secure API connections and cloud services like AWS Bedrock and GCP Vertex.&#x20;

### Product-specific AI implementations

#### Agent Fix & Explain

<table data-header-hidden><thead><tr><th width="226.24609375"></th><th></th></tr></thead><tbody><tr><td><strong>Attribute</strong></td><td><strong>Details</strong></td></tr><tr><td><strong>Purpose</strong></td><td><p>Designed to help developers:</p><p></p><ul><li>Fix their code faster by suggesting fixes to vulnerabilities identified by Snyk Code; and</li><li>Better understand findings and suggestions returned by Snyk by providing detailed explanations on demand.</li></ul></td></tr><tr><td><strong>AI models / deployment</strong></td><td>A combination of Snyk’s proprietary DeepCode AI engine and other open-source models that may be fine-tuned on Snyk’s existing datasets (which do not include any customer proprietary software code) and which are maintained and hosted entirely within our controlled environment.</td></tr><tr><td><strong>Data processed</strong></td><td>Code snippets containing only the relevant scope of the vulnerability.</td></tr><tr><td><strong>Data retention</strong></td><td>Because these models are entirely Snyk hosted, no customer proprietary software code is retained by them.</td></tr><tr><td><strong>Additional information</strong> </td><td>More information about Agent Fix is available <a href="../scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically.md">here</a>.</td></tr></tbody></table>

#### &#x20;Snyk Assist for Snyk Learn

<table data-header-hidden><thead><tr><th width="230.103515625"></th><th></th></tr></thead><tbody><tr><td><strong>Attribute</strong></td><td><strong>Details</strong></td></tr><tr><td><strong>Purpose</strong></td><td><p>AI powered chat assistant designed to help developers and Snyk users:</p><p></p><ul><li>Obtain contextually relevant assistance when navigating the information and resources available within Snyk Learn; and</li><li>Get immediate customized answers to specific application security, secure coding and Snyk product usage questions.</li></ul></td></tr><tr><td><strong>AI models / deployment</strong></td><td>OpenAI’s GPT-4o model, accessed via secure API connections.</td></tr><tr><td><strong>Data Processed</strong></td><td>User input, in the form of chat-based questions submitted by developers and Snyk users.</td></tr><tr><td><strong>Safeguards</strong></td><td><p>Snyk has implemented:</p><p></p><ul><li>Technical safeguards designed to check for code in user input; if found, code is not sent to the AI model or stored by Snyk; and</li><li>Measures designed to handle inappropriate user input, for your safety and that of Snyk.</li></ul></td></tr><tr><td><strong>Data retention</strong></td><td>Anonymized user inputs are retained by Snyk for a reasonable period for monitoring and managing service performance, after which they are permanently deleted.</td></tr><tr><td><strong>Training</strong></td><td>In addition to the training restriction described above, the content of any Snyk Assist prompts are not used to train OpenAI’s models.</td></tr><tr><td><strong>Additional information</strong></td><td>More information about Snyk Assist is available <a href="../snyk-learn/snyk-assist.md">here</a>.</td></tr></tbody></table>

#### &#x20;Snyk API & Web (False Positive Reduction (FPR))

<table data-header-hidden><thead><tr><th width="230.103515625"></th><th></th></tr></thead><tbody><tr><td><strong>Attribute</strong></td><td><strong>Details</strong></td></tr><tr><td><strong>Purpose</strong></td><td>Designed to help classify findings to reduce manual review and improve efficiency.</td></tr><tr><td><strong>AI models / deployment</strong></td><td>Anthropic’s Claude models through AWS Bedrock.</td></tr><tr><td><strong>Data Processed</strong></td><td>Parts of HTTP requests and responses (i.e. components of web communications that are analyzed to detect and classify vulnerabilities).</td></tr><tr><td><strong>Data retention</strong></td><td>Customer proprietary software code is not passed to or retained by the provider of these AI models.</td></tr><tr><td><strong>Additional information</strong></td><td>More information about Snyk API &#x26; Web is available <a href="https://help.probely.com/en/">here</a>.</td></tr></tbody></table>

### &#x20;Additional AI safeguards and controls

Snyk has taken a proactive approach to AI governance by implementing robust policies, procedures and technical controls to encompass AI-specific considerations. In addition to Snyk's internal policies and controls, we maintain an overarching AI Governance Program managed by our cross-functional AI Advisory Board.

Snyk does not develop general-purpose AI models. Our proprietary AI is purpose-built to support the same functionality as our underlying platform: identifying vulnerabilities in code, proposing fixes to those vulnerabilities, and promoting security within the software development lifecycle. Additionally, our AI governance incorporates key principles of emerging AI regulations. This includes validating our training datasets for quality and copyright compliance, and ongoing testing of output quality. Snyk's AI capabilities are designed to enable our customers' to assess AI-related risks and vulnerabilities, including governance mechanisms, transparency measures, and security controls.

### How Snyk handles data generally

View [How Snyk handles your data](how-snyk-handles-your-data.md) for more general information about Snyk’s data management practices.&#x20;
