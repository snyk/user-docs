---
description: How Snyk uses generative AI across its platform, including the third-party models that power its AI features
nav_context: agnostic
---

# How Snyk incorporates generative AI into the platform

Snyk’s AI Security Platform uses generative AI to enhance automation, efficiency, and innovation for developers and security teams. Snyk’s generative AI features are powered by third-party large language models (LLMs) from established AI providers.

This document explains what generative AI technologies Snyk uses and how data flows through our systems. It also describes the measures we take to protect your data. The field of AI is changing quickly. As a result, the AI technologies we use may change when we introduce new features or update existing ones.

## Core principles

Snyk places the utmost importance on data security and integrity.

* **No training on customer code**: Snyk does not use customer proprietary software code to train, optimize, fine-tune, or improve any AI models, and does not use or incorporate any third-party AI models into the platform unless they make the same commitments.
* **Contractual protection**: All of the AI functionality described in this document forms part of Snyk’s services. Your use of this functionality is governed by your existing agreements with Snyk and benefits from the same contractual protections. No separate in-service terms, addenda, or amendments to your existing agreements with Snyk are required.

## AI models

Snyk uses LLMs from established AI providers, including OpenAI and Anthropic, through API connections and cloud services like AWS Bedrock and GCP Vertex.

## Product-specific AI implementations

### Agent Fix & Explain

<table><thead><tr><th width="226.24609375">Attribute</th><th>Details</th></tr></thead><tbody><tr><td><strong>Purpose</strong></td><td><p>Designed to help developers:</p><ul><li>Fix their code faster by suggesting fixes to vulnerabilities identified by Snyk Code; and</li><li>Better understand findings and suggestions returned by Snyk by providing detailed explanations on demand.</li></ul></td></tr><tr><td><strong>AI models / deployment</strong></td><td>Anthropic’s Claude models through AWS Bedrock or GCP Vertex.</td></tr><tr><td><strong>Data processed</strong></td><td>Code snippets containing only the relevant scope of the vulnerability.</td></tr><tr><td><strong>Data retention</strong></td><td>Customer proprietary software code is not retained by the provider of these AI models.</td></tr><tr><td><strong>Additional information</strong></td><td>More information about Agent Fix is available <a href="https://app.gitbook.com/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically">here</a>.</td></tr></tbody></table>

### Snyk Assist for Snyk Learn

<table><thead><tr><th width="230.103515625">Attribute</th><th>Details</th></tr></thead><tbody><tr><td><strong>Purpose</strong></td><td><p>AI powered chat assistant designed to help developers and Snyk users:</p><ul><li>Obtain contextually relevant assistance when navigating the information and resources available within Snyk Learn; and</li><li>Get immediate customized answers to specific application security, secure coding and Snyk product usage questions.</li></ul></td></tr><tr><td><strong>AI models / deployment</strong></td><td>Anthropic's Claude models through GCP Vertex.</td></tr><tr><td><strong>Data Processed</strong></td><td>User input, in the form of chat-based questions submitted by developers and Snyk users.</td></tr><tr><td><strong>Safeguards</strong></td><td><p>Snyk has implemented:</p><ul><li>Technical safeguards designed to check for code in user input — if found, code is not sent to the AI model or stored by Snyk.</li><li>Measures designed to handle inappropriate user input, for your safety and that of Snyk.</li></ul></td></tr><tr><td><strong>Data retention</strong></td><td>Anonymized user inputs are retained by Snyk for a reasonable period for monitoring and managing service performance, after which they are permanently deleted.</td></tr><tr><td><strong>Additional information</strong></td><td>More information about Snyk Assist for Learn is available <a href="https://app.gitbook.com/s/L7HyJj9FsK1W4pNt8Gzl/developer-education-with-snyk-learn/snyk-learn/snyk-assist">here</a>.</td></tr></tbody></table>

### Snyk Assist for Support

<table><thead><tr><th width="230.103515625">Attribute</th><th>Details</th></tr></thead><tbody><tr><td><strong>Purpose</strong></td><td>AI-powered assistant designed to help users obtain support through the enterprise search and virtual agent capabilities within the Snyk support portal.</td></tr><tr><td><strong>AI models / deployment</strong></td><td>OpenAI, accessed through secure API connections, for question moderation only; Anthropic's models through GCP Vertex for answer generation.</td></tr><tr><td><strong>Data Processed</strong></td><td><ul><li>For question moderation: User input in the form of chat-based questions (i.e. question text only); and</li><li>For answer generation: User input in the form of chat-based questions, and limited user context including user email, account name, account type, and group name.</li></ul></td></tr><tr><td><strong>Data retention</strong></td><td><ul><li>OpenAI: Snyk has enabled OpenAI’s <a href="https://platform.openai.com/docs/guides/your-data">Zero Data Retention</a> control;</li><li>GCP Vertex: Data may be cached in-memory for up to 24 hours to reduce latency; no data is stored at rest. Prompts may be logged for up to 30 days where automated safety classifiers detect suspicious activity requiring investigation into potential policy violations.</li></ul></td></tr><tr><td><strong>Additional information</strong></td><td>Snyk Assist for Support is available <a href="https://support.snyk.io/s/">here</a>.</td></tr></tbody></table>

### Snyk API & Web (False Positive Reduction (FPR))

<table><thead><tr><th width="230.103515625">Attribute</th><th>Details</th></tr></thead><tbody><tr><td><strong>Purpose</strong></td><td>Designed to help classify findings to reduce manual review and improve efficiency.</td></tr><tr><td><strong>AI models / deployment</strong></td><td>Anthropic’s Claude models through AWS Bedrock.</td></tr><tr><td><strong>Data Processed</strong></td><td>Parts of HTTP requests and responses (i.e. components of web communications that are analyzed to detect and classify vulnerabilities).</td></tr><tr><td><strong>Data retention</strong></td><td>Customer proprietary software code is not passed to or retained by the provider of these AI models.</td></tr><tr><td><strong>Additional information</strong></td><td>More information about Snyk API &#x26; Web is available <a href="https://app.gitbook.com/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-api-web">here</a>.</td></tr></tbody></table>

### Snyk API & Web - Broken Object Level Authorization for APIs

<table><thead><tr><th width="229.3359375">Attribute</th><th>Details</th></tr></thead><tbody><tr><td><strong>Purpose</strong></td><td>Designed to help identify authorization vulnerabilities in APIs, including Broken Object Level Authorization.</td></tr><tr><td><strong>AI models / deployment</strong></td><td>Anthropic’s Claude models through AWS Bedrock.</td></tr><tr><td><strong>Data processed</strong></td><td>Parts of HTTP requests and responses (i.e. components of web communications that are analyzed to detect vulnerabilities and classify content).</td></tr><tr><td><strong>Data retention</strong></td><td>HTTP requests and responses are not retained by the provider of these AI models.</td></tr><tr><td><strong>Additional information</strong></td><td>More information about Snyk API &#x26; Web is available <a href="https://app.gitbook.com/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-api-web">here</a>.</td></tr></tbody></table>

### Snyk SAST / DAST Correlation

<table><thead><tr><th width="230.103515625">Attribute</th><th>Details</th></tr></thead><tbody><tr><td><strong>Purpose</strong></td><td>Designed to correlate Snyk Code and Snyk API &#x26; Web findings by highlighting the source code that triggered a given API &#x26; Web Vulnerability.</td></tr><tr><td><strong>AI models / deployment</strong></td><td>Google’s Gemini and Anthropic’s Claude models through AWS Bedrock and GCP Vertex.</td></tr><tr><td><strong>Data Processed</strong></td><td>Customer source code; DAST scan information (for example, vulnerability types, endpoints, and parameters); and SAST scan information (for example, vulnerability types, file names, and code locations), to help correlate source code with runtime vulnerabilities.</td></tr><tr><td><strong>Data retention</strong></td><td>Customer proprietary software code is not retained by the provider of these AI models.</td></tr><tr><td><strong>Additional information</strong></td><td>More information about Snyk API &#x26; Web is available <a href="https://app.gitbook.com/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-api-web">here</a>.</td></tr></tbody></table>

### AI-SPM

#### Evo Chat Agent

<table><thead><tr><th width="230.15625">Attribute</th><th>Details</th></tr></thead><tbody><tr><td><strong>Purpose</strong></td><td>AI-powered security assistant designed to help users explore the data available to them as part of AI-SPM (e.g. understanding AI assets, policy results, creating and updated policies, generating basic reporting).</td></tr><tr><td><strong>AI models / deployment</strong></td><td>Google’s Gemini and Anthropic’s Claude models through GCP Vertex and AWS Bedrock (depending on the deployment region).</td></tr><tr><td><strong>Data processed</strong></td><td><ul><li>User context (user ID, tenant ID, and authentication token) used to scope requests and authorise downstream API calls.</li><li>User chat messages and agent responses within a session.</li><li>AI asset metadata (e.g. model name, type, vendor, repository, organisation), asset relationships, security policies, and policy violation issues.</li></ul></td></tr><tr><td><strong>Data retention</strong></td><td><p>AWS Bedrock: Prompts and completions are not retained;<br></p><p>Google Vertex: Data may be cached in-memory for up to 24 hours to reduce latency; no data is stored at rest. Prompts may be logged for up to 30 days where automated safety classifiers detect suspicious activity requiring investigation into potential policy violations.</p></td></tr><tr><td><strong>Additional information</strong></td><td>More information about Evo Chat Agent is available via the Customer portal.</td></tr></tbody></table>

#### Evo AI-BOM Pattern Detection

<table><thead><tr><th width="230.15625">Attribute</th><th>Details</th></tr></thead><tbody><tr><td><strong>Purpose</strong></td><td>Allows the EVO AI-BOM to detect customer-specific AI assets, such as fine-tuned models.</td></tr><tr><td><strong>AI models / deployment</strong></td><td>Anthropic’s Claude models through GCP Vertex or AWS Bedrock.</td></tr><tr><td><strong>Data processed</strong></td><td>Code snippets used to label whether patterns are AI assets.</td></tr><tr><td><strong>Data retention</strong></td><td>Customer proprietary software code is not retained by the provider of these AI models.</td></tr><tr><td><strong>Additional information</strong></td><td>More information about Evo AI-BOM Pattern Detection is available via the Customer portal.</td></tr></tbody></table>

#### Evo AI-BOM Agentic Enrichment

<table><thead><tr><th width="230.15625">Attribute</th><th>Details</th></tr></thead><tbody><tr><td><strong>Purpose</strong></td><td>AI-powered agent that operates in two parts: (1) Agentic Discovery - discovers first-party AI components (agents, tools, MCP servers/clients/resources, models) and the relations between them that static analysis misses; (2) Agentic Metadata Gathering - extracts tool descriptions and agent system prompts from already-discovered components by reading their source code.</td></tr><tr><td><strong>AI models / deployment</strong></td><td>Anthropic’s Claude models through GCP Vertex or AWS Bedrock.</td></tr><tr><td><strong>Data processed</strong></td><td><ul><li>Repository source code</li><li>Existing inventory of AI components and relations from static analysis</li><li>Source code locations of already-discovered tools and agents for metadata extraction</li></ul></td></tr><tr><td><strong>Data retention</strong></td><td>Customer proprietary software code is not retained by the provider of these AI models. Tool descriptions and agent system prompts are stored in Snyk’s database as properties of their respective components.</td></tr><tr><td><strong>Additional information</strong></td><td>More information about Evo AI-BOM Agentic Enrichment is available via the Customer portal.</td></tr></tbody></table>

## Additional AI safeguards and controls

Snyk has taken a proactive approach to AI governance by implementing robust policies, procedures, and technical controls to encompass AI-specific considerations. In addition to Snyk's internal policies and controls, we maintain an overarching AI Governance Program managed by our cross-functional AI Advisory Board.

Snyk’s AI capabilities are specifically designed to support the same functionality as our underlying platform: identifying vulnerabilities in code, proposing fixes to those vulnerabilities, and promoting security within the software development lifecycle. Our AI governance incorporates key principles of emerging AI regulations. This includes validating our deterministic training datasets for quality and copyright compliance, and ongoing testing of output quality. Snyk's AI capabilities are designed to enable our customers to assess AI-related risks and vulnerabilities, including governance mechanisms, transparency measures, and security controls.

## How Snyk handles data generally

View [How Snyk handles your data](how-snyk-handles-your-data.md) for more general information about Snyk’s data management practices.
