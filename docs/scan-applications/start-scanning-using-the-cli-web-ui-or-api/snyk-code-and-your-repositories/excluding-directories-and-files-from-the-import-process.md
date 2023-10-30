# Exclude directories and files from the import process

When you import a repository to be tested by Snyk Code, you can exclude certain directories and files from the import by using the `.snyk` file. The `.snyk` file is a YAML policy file that can contain shell matching patterns (regular expressions), which allow you to specify the directories and files you want to exclude from the import process. The `.snyk` file should be created in the repository you intend to import.

{% hint style="warning" %}
* In Snyk Code, the`.snyk` file can only be used to exclude directories and files from import. It cannot be used to ignore vulnerabilities or for any other action, as in other Snyk products.
* Currently, the `Exclude` option in the`.snyk` file applies only to the Snyk Web UI and CLI Environments. The `Exclude` option cannot be used when working with Snyk Code in an IDE environment.
* In certain situations, your excluded files may not be excluded if there is an invalid`.snyk` file. In these situations, the scan continues without the`.snyk` file.
{% endhint %}

{% hint style="info" %}
Consider excluding directories and files only if you do not publish or compile them into production. If a trace goes through an excluded file or directory with existing vulnerabilities, Snyk might miss potential issues.
{% endhint %}

You can also use the instructions in this section to exclude directories and files from the [Snyk Code CLI test](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/using-snyk-code-from-the-cli/testing-your-source-code-using-the-cli.md), by creating the `.snyk` file in your tested repository.

## **Exclusion syntax of the .snyk file**

### Snytax to use to exclude files and directories from Snyk Code testing

Use the following syntax to exclude files and directories by using the `.snyk` file:

<pre class="language-yaml"><code class="lang-yaml"># Snyk (https://snyk.io) policy file

exclude:
    global:
<strong>        # Exclude a single file. For example, - test.spec.js
</strong>            - file_name.ext
<strong>        # Exclude a single directory. For example, - src/lib
</strong>            - source/directory_name
        # Exclude any file with a specific extension in the specific directory. For example, - tests/.js
<strong>            - directory_name/.ext
</strong>        # Exclude files with a specific ending in any directory. For example, - “*.spec.js”
            - "*.ending.ext"
        # Exclude files in directories that have the same name with a different ending, like “test” and “tests”. The last character before the question mark is optional. For example, - tests?/
            - directory_name?/
        # Exclude all files and directories in a specific directory. For example, - tests/
<strong>            - directory_name/**
</strong></code></pre>

{% hint style="info" %}
You can use `global` or `code.` Either will exclude the specified directories and files from Snyk code tests. `code` applies only to Snyk Code analysis.`global` currently applies only to analysis using Snyk Code but may apply to other products in the future.
{% endhint %}

### **Considerations in creating the `.snyk` file**

* The path in the rule should be relative to the `.snyk` file location.
* All rules must have a preceding dash to be valid: `- <Exclusion_rule>`
* For rules beginning with special characters and patterns, such as an asterisk character `*`, you must wrap them in double quotes (`" "`). This ensures they are treated as a single entity, avoiding potential misinterpretation or unintended behavior. For example, `"*/src"`
* The following are considerations in using indentations:
  * When using the syntax in the `.snyk` YAML file, pay careful attention to new lines and their indentation. Using the wrong indentation will prevent the execution of your excluding specification.
  * Do NOT use tabs for indentation. Use only spaces for indentation.
  * To verify that you are using the syntax correctly, you can use a YAML Validator, like [YAML Lint](http://www.yamllint.com/). Be aware that some YAML Validators do not differentiate between the use of tabs and spaces for indentation. If you use tabs, a Validator may approve the syntax, but the exclude specifications will not be executed.
* For more information on the syntax of shell matching patterns, see, for example. the following:
  * GNU Org - [Shell Pattern Matching](https://www.gnu.org/software/findutils/manual/html\_node/find\_html/Shell-Pattern-Matching.html)
  * Docstore - [Pattern Matching Quick Reference with Examples](https://docstore.mik.ua/orelly/unix/upt/ch26\_10.htm)

## **Use the `.snyk` file to exclude directories and files from import**

Follow these steps to exclude directories and files from the import process using the .`snyk` file:

1\. In the repository you want to import, create a YAML file called `.snyk`, for example`:`

<figure><img src="../../../.gitbook/assets/Snyk Code - Exlude from Import - .snyk file creation - 2.png" alt=".snyk file in a repository"><figcaption><p><code>.snyk</code> file in a repository</p></figcaption></figure>

2\. In the `.snyk` file, specify the directories or files or both that you want to exclude from import according to the following syntax:

```yaml
# Snyk (https://snyk.io) policy file
exclude:
 global:
   - <Exclusion_rule>
   - <Exclusion_rule>
```

For example:

```yaml
# Snyk (https://snyk.io) policy file
exclude:
 global:
   - todolist-goof/** 
```

3\. From the Snyk Web UI, import your repository in one of the following ways:

*   If the repository was already imported to Snyk, retest the repository as follows:

    On the **Projects** page, click the **Code analysis** Project of the repository. Then, on the **Code Analysis** page, click the **Retest now** option:

<figure><img src="../../../.gitbook/assets/image (459).png" alt="Retest now option"><figcaption><p>Retest now option</p></figcaption></figure>

* If the repository was not imported yet to Snyk, [import the repository](./#import-repository-to-snyk).

Your repository is imported to Snyk, without the directories and/or files you selected to exclude.

## **Example**: **Excluding two files from Snyk Code analysis**

We have a repository called `snyk-goof`, which we want to test for vulnerabilities using Snyk Code. After we import his repository to Snyk, we get a list of ten detected vulnerability issues, which were found in three files:

<figure><img src="../../../.gitbook/assets/Snyk Code - Exlude from Import - Example - Before Exclude.png" alt="Vulnerabilities detected found in three files"><figcaption><p>Vulnerabilities detected found in three files</p></figcaption></figure>

Now we want to exclude the `app.js` and `db.js` files from the Snyk Code analysis. To achieve that, we do the following:

1\. We create a `.snyk` file in the **snyk-goof** repository in GitHub:

<figure><img src="../../../.gitbook/assets/Snyk Code - Exlude from Import - Example - .snyk file creation.png" alt=".snyk file in snyk-goof repository"><figcaption><p><code>.snyk</code> file in snyk-goof repository</p></figcaption></figure>

2\. In the `.snyk` file, we enter the following commands to exclude the `app.js` and `db.js` files from the import:

<figure><img src="../../../.gitbook/assets/Snyk Code - Exlude from Import - Example - Command.png" alt=".snyk file commands"><figcaption><p><code>.snyk</code> file commands</p></figcaption></figure>

3\. We retest the **snyk-goof** repository by clicking the **Retest now** option on the **Code Analysis** page for the repository.

The `app.js` and `db.js` files are excluded from the retest and therefore are not tested by Snyk Code and do not appear in the Code Analysis results. Now only five vulnerability issues are detected:

<figure><img src="../../../.gitbook/assets/image (545) (1).png" alt="Detected issues after files are excluded"><figcaption><p>Detected issues after files are excluded</p></figcaption></figure>
