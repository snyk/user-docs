# Fix code vulnerabilities automatically

{% hint style="warning" %}
**Release status**&#x20;

DeepCode AI Fix is in [Early Access](../../../getting-started/snyk-release-process.md#early-access) for anyone to try in the IDE.

To enable the feature, see [Enable DeepCode AI Fix](fix-code-vulnerabilities-automatically.md#enable-deepcode-ai-fix).
{% endhint %}

Fix the security issues and quality flaws in the source code through an automated flow. DeepCode AI Fix calculates the most suitable solution for your issues and applies it automatically.

## Why use DeepCode AI Fix?

DeepCode AI Fix combines the power of a thorough program analysis engine with the abilities of an in-house deep learning-based large language model. This combination allows for compiling large amounts of unstructured language information [from open-source code](fix-code-vulnerabilities-automatically.md#what-data-does-deepcode-ai-fix-suggestions-collect).

Key features set DeepCode AI Fix apart. It has a neural network trained on millions of lines of code, allowing for greater versatility and creativity. The [Snyk Code engine](../snyk-code-local-engine.md) rigorously checks the suggestions from the neural network, ensuring all automated fixes are small and targeted to each vulnerability or code issue.

## What issues can you fix automatically?

You can address various issues detected by the Snyk Code engine in terms of quality, promoting best code practices, and security vulnerabilities. DeepCode AI Fix currently does not support inter-file fixes.&#x20;

## DeepCode AI Fix language support

DeepCode AI Fix supports the following languages:

* Javascript and Typescript
* Java
* Python
* C/C++
* Go (Limited support)
* C# (Limited support)
* APEX (Limited support)

What is the difference between supported and limited support?&#x20;

* Supported languages provide remediation for 10 or more rules covering the OWASP Top 10,&#x20;
* Limited support languages provide remediation for less than 10 rules.

## What data does DeepCode AI Fix collect?

### Customer data

DeepCode AI Fix does not collect customer data for training purposes nor send customer data to third parties.

### Training data

The Large Language Model (LLM) is trained exclusively on public repositories with **permissive licenses**. If a license for a repository changes after the initial scrape, the repository is immediately excluded from the training data. DeepCode AI Fix does not use customer data for training purposes.

The data collection process is thorough and includes the following:

* Static analysis of permissive public repositories
* Automated assessment of the suggested fix qualities
* Partial in-house labeling by humans

The training data is ensured to be of the highest quality to optimize the performance of the LLM.

For more information on how Snyk manages data, see [How Snyk handles your data](../../../working-with-snyk/how-snyk-handles-your-data.md).

## How DeepCode AI Fix works

A representation of information flow involved in fixing one issue is presented in the following table.

<table><thead><tr><th width="211">Stage</th><th>Subsystem<select><option value="144c7d0e56c649fdaffeeef234193541" label="Static Code Analysis Engine" color="blue"></option><option value="453931e7eaf94118b3ea6ec945dfce7f" label="Neural Network (Generative LLM)" color="blue"></option></select></th><th>Details</th></tr></thead><tbody><tr><td>Code scan and   discovery of issues</td><td><span data-option="144c7d0e56c649fdaffeeef234193541">Static Code Analysis Engine</span></td><td>Corresponds to a normal flow of scanning the code from IDE.</td></tr><tr><td>Code preprocessing and minimization with respect to the data flow of the particular issue <span class="math">\mathcal{I}</span></td><td><span data-option="144c7d0e56c649fdaffeeef234193541">Static Code Analysis Engine</span></td><td>Data flow of <span class="math">\mathcal{I}</span> is analyzed and code is minimized, keeping the relevant context only.</td></tr><tr><td>Generating <span class="math">k</span> candidate fixes for the given issue <span class="math">\mathcal{I}</span></td><td><span data-option="453931e7eaf94118b3ea6ec945dfce7f">Neural Network (Generative LLM)</span></td><td>Here, <span class="math">k</span> is an implementation parameter.</td></tr><tr><td>Candidate fixes ranking and self-assessment</td><td><span data-option="144c7d0e56c649fdaffeeef234193541">Static Code Analysis Engine</span></td><td>Each of the <span class="math"> k</span> fixes is assessed by the Code Engine, filtering out those rendering invalid code or failing to fix the issue (the issue persists).</td></tr><tr><td>Returning the best candidate fix </td><td></td><td>The system has finished.</td></tr></tbody></table>

## Requirements for DeepCode AI Fix

* [Snyk Code Enterprise plan](../../../implement-snyk/enterprise-implementation-guide/trial-limitations.md).
* Snyk IDE Plugin for  [VS Code](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner-preview), [Eclipse](https://marketplace.eclipse.org/content/snyk-security-code%E2%80%8B-open-source%E2%80%8B-iac-configurations), or [IntelliJ](../../../scm-ide-and-ci-cd-workflow-and-integrations/snyk-ide-plugins-and-extensions/jetbrains-plugins/)

## Enable DeepCode AI Fix

Enable DeepCode AI Fix  for your Organization in Snyk Web UI by navigating to **Settings** > **Snyk Preview**.

<figure><img src="../../../.gitbook/assets/enable_fix_suggestions_snyk_preview.png" alt="DeepCodeAI Fix Suggestions settings in Snyk Preview"><figcaption><p>DeepCodeAI Fix Suggestions settings in Snyk Preview</p></figcaption></figure>

## Fix code issues automatically

{% hint style="info" %}
**Before you begin**

* Make sure you have Snyk IDE with Automated fixes.
* Save the files and [scan your code](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/scan-source-code-with-snyk-code-using-the-cli.md) to generate a fresh set of results.
* You should see a zap icon:zap:next to all Snyk Code issues that can be automatically fixed.
{% endhint %}

1. Open your code base.
2. Find and fix issues through the panel or by clicking **Fix this issue** in Code Lens.
3. After a fix has been applied,  save and rescan.

## Example: Fix a code issue automatically

DeepCode AI Fix highlights all identified vulnerabilities that can be automatically fixed. These are highlighted with a zap icon :zap:. For example, in this scenario, we have identified a Cross-Site Request Forgery (CSRF).

Opening the vulnerability gives us details on where the issue is and allows us to generate a fix using DeepCode AI Fix.

<figure><img src="../../../.gitbook/assets/image (444).png" alt=""><figcaption><p>Opening the Snyk Code vulnerability panel</p></figcaption></figure>

Once you click on Generate fix using Snyk DeepCode AI, the machines will start turning and up to 5 fixes will be generated. To ensure we have fixed the vulnerability and DeepCode AI has not hallucinated and added a new vulnerability, we automatically retest all fixes with Snyk Code's engine.

The result, in this case, is 5 fixes, which you can navigate through to decide which one is best for you. The first one is importing and using `csrf`, should solve this issue.

<figure><img src="../../../.gitbook/assets/image (443).png" alt=""><figcaption><p>5 fixes have been generated</p></figcaption></figure>

When you apply the fix in Code Lens, you will be guided to where the new code has been introduced. After you save and rescan, the vulnerability will disappear.

<figure><img src="../../../.gitbook/assets/image (447).png" alt=""><figcaption><p>Vulnerability has been fixed</p></figcaption></figure>

