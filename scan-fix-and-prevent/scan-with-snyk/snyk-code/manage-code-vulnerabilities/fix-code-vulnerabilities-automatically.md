# Fix code vulnerabilities automatically

| Supported  | Limited support |
| ---------- | --------------- |
| Java       | APEX            |
| JavaScript | C#              |
| Python     | Go              |
| TypeScript | C/C++           |

<table><thead><tr><th width="211">Stage</th><th>Subsystem<select><option value="144c7d0e56c649fdaffeeef234193541" label="Static Code Analysis Engine" color="blue"></option><option value="453931e7eaf94118b3ea6ec945dfce7f" label="Neural Network (Generative LLM)" color="blue"></option></select></th><th>Details</th></tr></thead><tbody><tr><td>Code scan and discovery of issues</td><td><span data-option="144c7d0e56c649fdaffeeef234193541">Static Code Analysis Engine</span></td><td>Corresponds to a normal flow of scanning the code from IDE.</td></tr><tr><td>Code preprocessing and minimization with respect to the data flow of the particular issue <span class="math">\mathcal{I}</span></td><td><span data-option="144c7d0e56c649fdaffeeef234193541">Static Code Analysis Engine</span></td><td>Data flow of <span class="math">\mathcal{I}</span> is analyzed and code is minimized, keeping the relevant context only.</td></tr><tr><td>Generating <span class="math">k</span> candidate fixes for the given issue <span class="math">\mathcal{I}</span></td><td><span data-option="453931e7eaf94118b3ea6ec945dfce7f">Neural Network (Generative LLM)</span></td><td>Here, <span class="math">k</span> is an implementation parameter.</td></tr><tr><td>Candidate fixes ranking and self-assessment</td><td><span data-option="144c7d0e56c649fdaffeeef234193541">Static Code Analysis Engine</span></td><td>Each of the <span class="math">k</span> fixes is assessed by the Code Engine, filtering out those rendering invalid code or failing to fix the issue (the issue persists).</td></tr><tr><td>Returning the best candidate fix</td><td></td><td>The system has finished.</td></tr></tbody></table>

{% hint style="info" %}
DeepCode AI Fix is now Snyk Agent Fix.

As of May 2026, Snyk Agent Fix has been upgraded to a new agentic architecture for significantly higher fix accuracy and broader language support.
{% endhint %}

Snyk Agent Fix provides production-ready code fixes to address security vulnerabilities and code quality flaws detected by Snyk Code. It offers full rule coverage for all supported languages.

Snyk Agent Fix uses an agentic architecture that combines Snyk proprietary security intelligence with advanced large language models (LLMs). Key advantages include:

* Dynamic few-shot prompting: Instead of relying on fine-tuning, the architecture uses the Snyk database of more than 35,000 expert-written fixes to provide real-world context to the LLM during inference. Every sample includes vulnerable code from real open-source projects and fixes written by Snyk security experts.
* Agentic retries: If a generated fix fails a Snyk Code scan, the system analyzes the error, feeds it back into the model, and generates a corrected version.

Snyk Agent Fix remediates vulnerabilities across your entire stack without language-specific fine-tuning. By using a prompt-based agentic reasoning model, Snyk Agent Fix supports all languages supported by Snyk Code: Apex, C, C++, C#, Go, Java, JavaScript, PHP, Python, Ruby, Swift, and TypeScript.

## How Snyk Agent Fix works

{% hint style="info" %}
Snyk Agent Fix does not use customer code to train underlying models, add to datasets, or improve performance.

For more information, see [How Snyk handles your data](https://app.gitbook.com/s/ELvljsaLKPkSpffOkmsQ/how-snyk-handles-your-data).
{% endhint %}

The agentic flow involves a feedback loop between the static analysis engine and the generative model.

<table><thead><tr><th width="211">Stage</th><th>Subsystem</th><th>Details</th></tr></thead><tbody><tr><td>Discovery</td><td>Static Code Analysis Engine</td><td>Identifies a vulnerability <span class="math">\mathcal{I}</span> during a standard scan.</td></tr><tr><td>Prompt enrichment</td><td>Snyk Intelligence DB</td><td>Retrieves relevant human-written fix examples for the specific CWE from our 35,000+ pair database.</td></tr><tr><td>Generation</td><td>Agentic LLM</td><td>Generates <span class="math">k</span> candidate fixes using dynamic few-shot prompting.</td></tr><tr><td>Verification</td><td>Static Code Analysis Engine</td><td>Checks each candidate to ensure the vulnerability is gone and no new ones have been introduced.</td></tr><tr><td>Agentic retry</td><td>Agentic Loop</td><td>If a fix fails verification, the system extracts the error, feeds it back to the agent, and attempts a corrected fix.</td></tr><tr><td>Final delivery</td><td>Snyk Interface</td><td>Presents the final, verified candidate to the developer.</td></tr></tbody></table>

## Enable Snyk Agent Fix in the Snyk web UI

Before enabling Snyk Agent Fix, ensure you:

* Enable [Snyk Code](https://app.gitbook.com/s/L7HyJj9FsK1W4pNt8Gzl/implementation-guides/enterprise-implementation-guide/create-a-template-organization/connect-your-development-tools#enable-snyk-code).
* Install the Snyk IDE plugin for [VS Code](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner-preview), [Visual Studio](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner-vs-2022), [Eclipse](https://marketplace.eclipse.org/content/snyk-security-code%E2%80%8B-open-source%E2%80%8B-iac-configurations), or [JetBrains IDEs](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions/jetbrains-plugin).

To enable the feature for only a specific Organization, use the Organization-level settings.

1. Navigate to **Settings** > **Snyk Agent Fix** for your **Group** or **Organization**.
2. Enable **Snyk Agent Fix**.

## Apply an automated fix

Snyk automatically generates fixes for eligible vulnerabilities in your codebase. A zap icon marks issues eligible for an automated fix.

To see the latest fix suggestions:

1. Enable automated fixes in Snyk Preview in your [IDE plugin or extension](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions).
2. Save your files and [trigger a scan](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/scan-source-code-with-snyk-code-using-the-cli).

To generate a fix:

1. Open your codebase in your IDE and navigate to the Snyk panel or use Code Lens to find vulnerabilities.
2. Click **Generate AI fix** and review the suggested fix.
3. Apply the fix, save the file, and rescan to confirm the fix resolves the vulnerability.

## Considerations

The agentic architecture reduces errors through the retry loop, but the following considerations apply:

* Human review required: Review suggestions to ensure they align with the broader application architecture.
* Complex inter-file logic: Snyk Agent Fix focuses on local-file fixes. It does not automatically fix complex vulnerabilities that span multiple files.
* Latency: The agentic retry loop takes time. Fix requests take up to two minutes if the initial responses need correction.
* Filtering: Snyk does not show a suggestion if the agentic loop cannot produce a fix that meets security and functional benchmarks.
