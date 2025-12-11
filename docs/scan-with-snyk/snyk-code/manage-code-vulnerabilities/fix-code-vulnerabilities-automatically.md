# Fix code vulnerabilities automatically

{% hint style="info" %}
DeepCode AI Fix has a new name: Snyk Agent Fix.
{% endhint %}

Fix security issues and quality flaws in the source code through an automated flow. Snyk Agent Fix calculates the most suitable solution for your issues and applies it automatically.

## Why use Snyk Agent Fix?

Snyk Agent Fix combines the power of a thorough program analysis engine with the abilities of an in-house deep-learning-based large language model. This combination allows for compiling large amounts of unstructured language information [from open-source code](fix-code-vulnerabilities-automatically.md#what-data-does-snyk-agent-fix-collect).

Key features set Snyk Agent Fix apart. It has a neural network trained on millions of lines of code, allowing for greater versatility and creativity. The [Snyk Code engine](../snyk-code-local-engine.md) rigorously checks the suggestions from the neural network, ensuring all automated fixes are small and targeted to each vulnerability or code issue.

## What issues can you fix automatically?

You can address various issues detected by the Snyk Code engine in terms of quality, promoting best code practices, and security vulnerabilities. Snyk Agent Fix currently does not support inter-file fixes.

## Snyk Agent Fix language support

| Supported  | Limited support |
| ---------- | --------------- |
| Java       | APEX            |
| JavaScript | C#              |
| Python     | Go              |
| TypeScript | C/C++           |

What is the difference between supported and limited support?

* Supported: Snyk Agent Fix consistently generates fixes for the majority of Snyk Code rules.
* Limited support: Snyk Agent Fix may not generate fixes as consistently, covering fewer Snyk Code rules.

## What data does Snyk Agent Fix collect?

### Training data

Snyk trains its Large Language Model (LLM) using permissively-licensed public repositories. Snyk does not use code input by customers to train its LLM.

The data collection process is thorough and includes the following:

* Static analysis of permissive public repositories.
* Automated assessment of the suggested fix qualities.
* Partial in-house labeling by humans.

The training data is regularly checked for quality to optimize the performance of the LLM.

### Customer data

Snyk does not use customer code submitted to Snyk Agent Fix for training purposes.

For more information on how Snyk manages data, see [How Snyk handles your data](../../../snyk-data-and-governance/how-snyk-handles-your-data.md).

## How Snyk Agent Fix works

A representation of information flow involved in fixing one issue is presented in the following table.

<table><thead><tr><th width="211">Stage</th><th>Subsystem<select><option value="144c7d0e56c649fdaffeeef234193541" label="Static Code Analysis Engine" color="blue"></option><option value="453931e7eaf94118b3ea6ec945dfce7f" label="Neural Network (Generative LLM)" color="blue"></option></select></th><th>Details</th></tr></thead><tbody><tr><td>Code scan and discovery of issues</td><td><span data-option="144c7d0e56c649fdaffeeef234193541">Static Code Analysis Engine</span></td><td>Corresponds to a normal flow of scanning the code from IDE.</td></tr><tr><td>Code preprocessing and minimization with respect to the data flow of the particular issue <span class="math">\mathcal{I}</span></td><td><span data-option="144c7d0e56c649fdaffeeef234193541">Static Code Analysis Engine</span></td><td>Data flow of <span class="math">\mathcal{I}</span> is analyzed and code is minimized, keeping the relevant context only.</td></tr><tr><td>Generating <span class="math">k</span> candidate fixes for the given issue <span class="math">\mathcal{I}</span></td><td><span data-option="453931e7eaf94118b3ea6ec945dfce7f">Neural Network (Generative LLM)</span></td><td>Here, <span class="math">k</span> is an implementation parameter.</td></tr><tr><td>Candidate fixes ranking and self-assessment</td><td><span data-option="144c7d0e56c649fdaffeeef234193541">Static Code Analysis Engine</span></td><td>Each of the <span class="math">k</span> fixes is assessed by the Code Engine, filtering out those rendering invalid code or failing to fix the issue (the issue persists).</td></tr><tr><td>Returning the best candidate fix</td><td></td><td>The system has finished.</td></tr></tbody></table>

## Requirements for Snyk Agent Fix

* [Snyk Code](../../../implementation-and-setup/enterprise-implementation-guide/trial-limitations.md) enabled
* Snyk IDE Plugin for [VS Code](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner-preview), [Visual Studio](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner-vs-2022), [Eclipse](https://marketplace.eclipse.org/content/snyk-security-code%E2%80%8B-open-source%E2%80%8B-iac-configurations), or [JetBrains IDEs](../../../developer-tools/snyk-ide-plugins-and-extensions/jetbrains-plugin/)

## Enable Snyk Agent Fix

Enable Snyk Agent Fix for your Group or Organization in the Snyk Web UI by navigating to **Group/Organization** > **Settings** > **Snyk Agent Fix**.

## Fix code issues automatically

{% hint style="info" %}
**Before you begin**

* Ensure you have automated fixes enabled in Snyk Preview to work with your [Snyk IDE plugin or extension](../../../developer-tools/snyk-ide-plugins-and-extensions/).
* Save the files and [scan your code](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/scan-source-code-with-snyk-code-using-the-cli.md) to generate a fresh set of results.
* You should see a zap icon next to all Snyk Code issues that can be automatically fixed.
{% endhint %}

1. Open your codebase.
2. Find and fix issues through the panel or by clicking **Fix this issue** in Code Lens.
3. After a fix has been applied, save and rescan.

## Example: Fix a code issue automatically

Snyk Agent Fix highlights all identified vulnerabilities that can be automatically fixed. These are highlighted with a zap icon. For example, in this scenario, we have identified an Information Exposure vulnerability.

Opening the vulnerability gives us details on where the issue is and allows us to generate a fix using Snyk Agent Fix.

<figure><img src="../../../.gitbook/assets/image (24) (3).png" alt=""><figcaption><p>Opening the Snyk Code vulnerability panel</p></figcaption></figure>

After you select Generate Fix using Snyk DeepCode AI, the system will analyze your code and generate up to five potential fixes. After you apply a fix, Snyk Agent Fix automatically retests the fix for quality using Snyk Code's engine.

The result, in this case, is five fixes, which you can navigate through to decide which one is best for you. The first fix is adding Helmet middleware package that disables the `X-Powered-By` header by default, preventing attackers from knowing that the app is running Express.

<figure><img src="../../../.gitbook/assets/Snyk AI Fix 5 examples.png" alt=""><figcaption><p>5 fixes have been generated</p></figcaption></figure>

When you apply the fix, you will be guided to where the new code has been introduced. After you save and rescan, the vulnerability will disappear.

<figure><img src="../../../.gitbook/assets/image (26) (1).png" alt=""><figcaption><p>Vulnerability has been fixed</p></figcaption></figure>

## Limitations

Snyk Agent Fix is at the forefront of AI, but there are still limitations based on the AI engine. Users must always review the Snyk Agent Fix suggestions to ensure that the resulting implementation of the fix does not break their application

* Snyk Agent Fix suggestions might generate code that results in the application not working properly
* Snyk Agent Fix suggestions might generate code that is not syntactically correct
* Snyk Agent Fix verifies that the vulnerability is fixed and no new vulnerabilities are introduced. This means that sometimes, a valid suggestion will not be made because the AI engine has not generated a good enough result.
