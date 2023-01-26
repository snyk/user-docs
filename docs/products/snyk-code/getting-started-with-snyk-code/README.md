# Getting started with Snyk Code

This section describes how to get started with Snyk Code using the Snyk Web UI Environment.

You can work with Snyk Code using the following Snyk Environments:

* [Snyk Web UI ](../../../getting-started/quickstart/create-a-snyk-account/snyk-web-ui.md)(including [PR Checks](../../../scan-application-code/run-pr-checks/pr-checks-for-snyk-code/))
* [Snyk for IDEs](../../../ide-tools/)
* [Snyk CLI](../../../snyk-cli/)
* [Snyk API](../../../snyk-api-info/)

The prerequisites for using Snyk Code are the same for all Snyk Environments:

* If you are using the Snyk Web UI, follow the step-by-step instructions in this section to activate and use Snyk Code.
* If you are using the Snyk IDE, CLI, or API Environments, use the Snyk Web UI to [enable the **Snyk Code** option](activating-snyk-code-using-the-web-ui/step-1-enabling-the-snyk-code-option.md). Then, continue to test your code with Snyk Code in your selected Environment.

### **What is required for getting started with Snyk Code?**

To get started with Snyk Code, all you need is a Snyk Account, a supported SCM, and source code that is written in one of the supported platforms and languages.

{% hint style="info" %}
For more information, see [Prerequisites for Snyk Code](prerequisites-for-snyk-code.md)
{% endhint %}

### **What is the activation process of Snyk Code?**

First, you need to enable the **Snyk Code** option in your Snyk Org Settings via the Web UI. If you intend on using the IDE, CLI, or API Environments, this is the only step you need to do in order to start working with Snyk Code.

{% hint style="info" %}
If you are working with the API Environment WITHOUT the CLI, you also need to [integrate your SCM via the Web UI](activating-snyk-code-using-the-web-ui/step-2-integrating-your-source-control-system-with-snyk-code.md)
{% endhint %}

When working with the Snyk Web UI, the next step in the activation process is integrating your Snyk Account with the SCM that contains the repositories you want to test. Then, you can import to your Snyk Account the required repositories, and Snyk Code will automatically analyze them and present to you its analysis results.

{% hint style="info" %}
For more information, see [Activating Snyk Code using the Web UI](activating-snyk-code-using-the-web-ui/).
{% endhint %}

### **Where are the Snyk Code Results?**

During the import process of your selected repositories, Snyk Code automatically analyzes your imported code in search of potential vulnerabilities. All the vulnerability findings that Snyk Code detects in the code of one imported repository, are aggregated in one Snyk Project, called **Code analysis**:

<figure><img src="../../../.gitbook/assets/SnykCode1.png" alt=""><figcaption></figcaption></figure>

**Note**: Unlike other Snyk products, which create a separate Snyk Project for each imported file, Snyk Code creates one Snyk Project for all the imported files of one repository. This way, all the vulnerabilities that were detected in the repository code are aggregated in one Project, and the Snyk Code results can present the data flow of a vulnerability across multiple files.

To view all the security vulnerabilities that Snyk Code detected in your imported code, click the **Code analysis** Project, and explore the details of each vulnerability:

<figure><img src="../../../.gitbook/assets/SnykCode2.png" alt=""><figcaption></figcaption></figure>

If your Snyk Account is already integrated with your SCM, and it already contains imported repositories, Snyk Code may be already active and running in your Org Settings. In this case, you can check if your existing repositories are already being tested by Snyk Code, by searching for the **Code analysis** Project in your Target folders. If a **Code analysis** Project exists in your imported repositories, you can skip this _Getting Started_ section and move to the [Exploring and working with the Snyk Code results](../exploring-and-working-with-the-snyk-code-results/) section. However, you may want to learn more about the following topics:

* [Importing additional repositories to Snyk](activating-snyk-code-using-the-web-ui/step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/importing-additional-repositories-to-snyk.md).
* [Excluding directories and files from the import process](activating-snyk-code-using-the-web-ui/step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/excluding-directories-and-files-from-the-import-process.md).
* [Removing imported repositories from the Snyk Code test](activating-snyk-code-using-the-web-ui/step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/removing-imported-repositories-from-the-snyk-code-test.md).

##
