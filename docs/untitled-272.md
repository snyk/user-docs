# Introduction to Snyk Code

* [ Introduction to Snyk Code](/hc/en-us/articles/360017059758-Introduction-to-Snyk-Code)
* [ Snyk Code language and framework support](/hc/en-us/articles/360016973477-Snyk-Code-language-and-framework-support)
* [ Using Snyk Code \(web\)](/hc/en-us/articles/360017147558-Using-Snyk-Code-web-)

##  Introduction to Snyk Code

#### Integrated platform

Snyk Code expands the Snyk cloud-native application security platform, supporting all elements for a modern software supply chain. For a cloud-native application, this provides security for: 

* [**Snyk Code**](https://support.snyk.io/hc/en-us/categories/360003257537-Snyk-Code): the application’s proprietary code.
* [**Snyk Open Source**](https://support.snyk.io/hc/en-us/categories/360003049458-Snyk-Open-Source): the open-source libraries it leverages.
* [**Snyk Container**](https://support.snyk.io/hc/en-us/categories/360000583498-Snyk-Container): the container it runs in.
* [**Snyk Infrastructure as Code**](https://support.snyk.io/hc/en-us/categories/360001342678-Infrastructure-as-code): the infrastructure as code that provisions it.

#### How does static application security testing \(SAST\) work?

1. Checks and reads for DeepCode/Snyk ignore specific files `.gitignore` `.dcignore` \(if they exist\).
2. Using the information obtained in step 1, we are filtering to get only the following source code files:
   1. `.es,.es6,.htm,.html,.js,.jsx,.ts,.tsx,.vue,.java,.CS,.Cs,.cs,.py`
   2. We are accessing only the files in the project directory. We do not go above the current project directory.
3. Files which size is less than 4 MB found in step 2 are bundled and the bundle is sent to Snyk.

#### Developer-first approach

Snyk Code is developer-first, embedding SAST as part of the development process, enabling developers to build software securely during development, not trying to find and fix problems after the code is compiled. Snyk Code works in the IDEs and SCMs developers use to build and review software and provides fast, actionable, meaningful results to fix issues in real-time.

#### Snyk Code AI engine

Snyk Code is powered by machine learning based on an AI engine. This engine learns from millions of open-source commits and is paired with Snyk’s Security Intelligence database, creating a continually growing code security knowledge base.

The Snyk Code engine expresses rules that capture results of analysis, plus meta information such as explanations and examples.

#### Fast and accurate code flaws detection

Snyk Code is significantly faster and more accurate than other conventional SAST products:

* The semantic analysis engine added via Snyk Code, trained on Snyk’s Vulnerability Database, reduces false positives to near-zero.
* Due to our proprietary engine, Snyk Code is up to 50x faster than traditional SAST solutions, enabling security during the development process, with no delays.

#### Secret detection in source code

Snyk Code includes secret detection capabilities that scan and highlight secrets like keys, credentials, PII, and sensitive information in your source code. Unlike tools that use entropy checks or regular expressions, Snyk Code uses machine learning and is able to learn from experience, improving the odds of accurately detecting secrets while minimizing false positives.

#### Issue code flow visualization

Snyk products all provide a developer-friendly experience, so Snyk Code helps developers to quickly understand the problem, learn the background, and how to approach it. Snyk Code helps you understand the dangerous code flow step-by-step.

For every issue, Code also provides a link to the lines in the relevant files, to view more details on the problem like the CWE, and how to approach it.

Snyk Code also provides rich curated additional information for many suggestions.

Snyk Code integrates with the cloud versions of:

* [GitHub](https://support.snyk.io/hc/en-us/articles/360004032117-GitHub-integration)
* [GitLab](https://support.snyk.io/hc/en-us/articles/360004002238-GitLab-integration)
* [Bitbucket](https://support.snyk.io/hc/en-us/articles/360004032097-Bitbucket-Cloud-integration) 

#### Using integrations

You can seamlessly import and continuously monitor your repositories, and discover security flaws in the source code files in them using static analysis.

This integration allows you to:

* Manage Code projects using your existing native import flow and tools
* View and prioritize security issues found in the source code
* Run a retest of a project and see history snapshots of a project

### Snyk Code IDE plugins

IDE integrations use Snyk Code’s fast analysis and response, allowing you to spot an issue, understand and learn more about it, and fix it, as you write the code before you check the code in. So you can find possible security flaws in your code as you write it, on a line-by-line basis.

#### JetBrains IDE plugin

Snyk Code supports a JetBrains plugin to support issue finding and fixing, directly from the IDE:

For more details, see [JetBrains IDE Plugins](https://support.snyk.io/hc/en-us/articles/360004032317-JetBrains-IDE-Plugins).

#### VS Code IDE plugin

Snyk Code supports a Visual Studio Code plugin to support issue finding and fixing, directly from the IDE:

For more details, see the [Visual Studio Code extension for Snyk Code](https://support.snyk.io/hc/en-us/articles/360018585717-Visual-Studio-Code-extension-for-Snyk-Code-).

