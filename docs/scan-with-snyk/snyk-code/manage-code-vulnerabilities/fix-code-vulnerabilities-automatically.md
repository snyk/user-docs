# Fix code vulnerabilities automatically

{% hint style="info" %}
DeepCode AI Fix has a new name: Snyk Agent Fix.

Note: As of May 2026, Snyk Agent Fix has been upgraded to a new agentic architecture for significantly higher fix accuracy and broader language support.
{% endhint %}

Fix security issues in source code with Snyk’s security expertise by your side.

## Why use Snyk Agent Fix?

Snyk Agent Fix leverages a state-of-the-art agentic architecture that moves beyond static AI models. It combines Snyk's proprietary security intelligence with advanced Large Language Models (LLMs) to provide developers with production-ready code fixes.

Key advantages include:

* **Dynamic Few-Shot Prompting**: Instead of relying on fine-tuning, the new architecture uses Snyk’s database of 35,000+ security expert written fixes to provide real-world context into the LLM at the moment of inference. Every sample includes vulnerable code taken from real Open Source projects as well as fixes written by Snyk security experts.
* **Agentic Retries**: If a generated fix fails a Snyk Code scan, the system automatically analyzes the error, feeds it back into the model, and re-thinks the solution to generate a corrected version.

## What issues can you fix automatically?

You can address the vast majority of issues detected by Snyk Code, including security vulnerabilities and code quality flaws. With the new architecture, Snyk Agent Fix now offers full rule coverage for all supported languages.

## Snyk Agent Fix language support

Snyk Agent Fix now provides comprehensive support for all languages supported by Snyk Code. This includes, but is not limited to:

| Category | Supported Languages |
| :--- | :--- |
| **Full Support** | Apex, C, C++, C#, Go, Java, JavaScript, PHP, Python, Ruby, Swift, TypeScript |

By moving to a prompt-based agentic reasoning model, Snyk Agent Fix can now remediate vulnerabilities across your entire stack without the need for language-specific fine-tuning.

## What data does Snyk Agent Fix collect?

### Training and Prompting Data

Snyk does not use code input by customers to train its underlying models, to add to its dataset, or to improve its performance in any way. To ensure high-quality suggestions, Snyk utilizes:

* **Curated Intelligence**: A proprietary database of over 35,000 unique, human-written vulnerability-fix pairs.
* **Public Repositories**: Use of permissively-licensed public code for benchmarking and context.

### Customer Data

Snyk does not use customer code submitted to Snyk Agent Fix for training purposes. For more information, see [How Snyk handles your data](../../../snyk-data-and-governance/how-snyk-handles-your-data.md).

## How Snyk Agent Fix works

The agentic flow involves a feedback loop between the static analysis engine and the generative model.

<table><thead><tr><th width="211">Stage</th><th>Subsystem</th><th>Details</th></tr></thead><tbody><tr><td>1. Discovery</td><td><span data-option="144c7d0e56c649fdaffeeef234193541">Static Code Analysis Engine</span></td><td>Identifies a vulnerability <span class="math">\mathcal{I}</span> during a standard scan.</td></tr><tr><td>2. Prompt Enrichment</td><td>Snyk Intelligence DB</td><td>Retrieves relevant human-written fix examples for the specific CWE from our 35,000+ pair database.</td></tr><tr><td>3. Generation</td><td><span data-option="453931e7eaf94118b3ea6ec945dfce7f">Agentic LLM</span></td><td>Generates <span class="math">k</span> candidate fix using dynamic few-shot prompting.</td></tr><tr><td>4. Verification</td><td><span data-option="144c7d0e56c649fdaffeeef234193541">Static Code Analysis Engine</span></td><td>Each candidate is checked to make sure the vulnerability is gone and no new ones have been introduced.</td></tr><tr><td>5. Agentic Retry</td><td>Agentic Loop</td><td>If a fix fails verification, the system extracts the error, feeds it back to the agent, and attempts a corrected fix.</td></tr><tr><td>6. Final Delivery</td><td>Snyk Interface</td><td>The final, verified candidate is presented to the developer.</td></tr></tbody></table>

## Requirements for Snyk Agent Fix

* [Snyk Code](../../../implementation-and-setup/enterprise-implementation-guide/trial-limitations.md) enabled
* Snyk IDE Plugin for [VS Code](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner-preview), [Visual Studio](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner-vs-2022), [Eclipse](https://marketplace.eclipse.org/content/snyk-security-code%E2%80%8B-open-source%E2%80%8B-iac-configurations), or [JetBrains IDEs](../../../developer-tools/snyk-ide-plugins-and-extensions/jetbrains-plugin/)
* Agent Fix enabled

## Enable Snyk Agent Fix

Enable Snyk Agent Fix for your Group or Organization in the Snyk Web UI by navigating to **Group/Organization** > **Settings** > **Snyk Agent Fix**.

* **Group level**: Enables the feature for all Organizations in that Group.
* **Organization level**: Enables the feature for a specific Organization only.

## Fix code issues automatically

{% hint style="info" %}
**Before you begin**

* Ensure you have automated fixes enabled in Snyk Preview in your [IDE plugin or extension](../../../developer-tools/snyk-ide-plugins-and-extensions/).
* Save your files and [trigger a scan](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/scan-source-code-with-snyk-code-using-the-cli.md) to see the latest fix suggestions.
* Issues eligible for an automated fix are marked with a zap icon.
{% endhint %}

1. Open your codebase in your IDE.
2. Navigate to the Snyk panel or use Code Lens to find vulnerabilities.
3. Click **Generate AI fix** (or use the zap icon).
4. Review the suggested fix.
5. Apply the fix, save, and rescan to confirm the vulnerability is resolved.

## Example: Fix a code issue automatically

Snyk Agent Fix highlights identified vulnerabilities with a zap icon. For example, if an Information Exposure vulnerability is found:

Opening the vulnerability provides details and the option to **Generate AI fix**. The system analyzes the code and uses agentic reasoning to provide the best possible remediation.

In this example, the agent might suggest adding the Helmet middleware to an Express app to disable the `X-Powered-By` header.

<figure><img src="../../../.gitbook/assets/Snyk AI Fix 5 examples.png" alt=""><figcaption><p>The Agentic engine generates and verifies multiple fix options.</p></figcaption></figure>

## Limitations

While the agentic architecture significantly reduces errors through the retry loop, some limitations remain:

* **Human Review Required**: Developers must always review suggestions to ensure they align with the broader application architecture.
* **Complex Inter-file Logic**: Snyk Agent Fix currently focuses on local-file fixes; complex vulnerabilities spanning many files are not automatically fixable.
* **Latency**: The agentic retry loop takes time; you may see requests for fixes taking up to two minutes if we find problems with the model’s initial responses.
* **Filtering**: If the agentic loop cannot produce a fix that meets Snyk's high security and functional benchmarks, no suggestion will be shown to avoid introducing new risks.