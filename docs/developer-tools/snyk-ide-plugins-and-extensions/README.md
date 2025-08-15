# Snyk IDE plugins and extensions

{% hint style="info" %}
If Snyk is hosting your data in a region other than `SNYK US-01`, you may set the custom endpoint in the IDE. For more information, see [IDEs URLS](../../snyk-data-and-governance/regional-hosting-and-data-residency.md#ides-urls) on the [Regional hosting and data residency](../../snyk-data-and-governance/regional-hosting-and-data-residency.md) page.\
Multi-tenant users who do not belong to the `SNYK-US-01` region will be automatically redirected to the correct domain for the email with which the user authenticated. The redirect will not occur for cases where the users are expected to use a custom URL, such as companies with single-tenant setups.
{% endhint %}

Snyk Security plugins and extensions find and fix security vulnerabilities and issues in Snyk Projects. This helps individual developers, open source contributors, and code maintainers to pass security reviews, avoid costly fixes later in development, and reduce time to develop and deliver secure code.

The results of a vulnerability scan show issues with context, impact, and fix guidance in your IDE, where the fix for the vulnerability can be done right in the IDE itself.

The Snyk IDE plugins and extensions rely on the [Snyk CLI](../snyk-cli/) to perform many functions. For details, refer to the documentation for each IDE.  When you are troubleshooting, it is always helpful to run the CLI for the same action with the debug option, `-d`.

The Snyk IDE plugins and extensions also rely on the [Snyk Vulnerability Database](https://security.snyk.io/). For more information, see the documentation for the [Snyk Vulnerability Database](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/snyk-vulnerability-database.md).

The following Snyk plugins and extensions are available.

* [Visual Studio Code extension](visual-studio-code-extension/)
* [JetBrains plugin](jetbrains-plugin/)
* [Visual Studio extension](visual-studio-extension/)
* [Eclipse plugin](eclipse-plugin/)

Snyk also offers a [Language Server](snyk-language-server/).&#x20;

The following summarizes the versions of each IDE supported by the Snyk plugin or extension. Snyk recommends always using the latest version of the IDE plugin or extension.

| Snyk plugin or extension     | IDE version supported                                                                                                                                                                                                |
| ---------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Visual Studio Code extension | The latest version of the Snyk Visual Studio Code extension supports use with Visual Studio Code version 1.76.0 and later.                                                                                           |
| JetBrains plugin             | <p>The latest Snyk JetBrains plugin supports use with all JetBrains IDEs 2024.2 or newer.</p><p>Older plugin versions may support use with JetBrains IDEs 2020.3 or newer.</p>                                       |
| Visual Studio extension      | <p>The latest version of the Snyk Visual Studio extension supports use with Visual Studio 2022 (version 17.5 and above).</p><p>Older extension versions may support use with Visual Studio 2015, 2017, and 2019.</p> |
| Eclipse plugin               | <p>The latest Snyk Eclipse plugin supports use with Eclipse 2024-03 or newer.</p><p>Older plugin versions may support use with Eclipse 2023-03 or newer.</p>                                                         |

To learn more, see the Snyk Learn lessons about [Using Snyk in an IDE](https://learn.snyk.io/lesson/snyk-in-an-ide/?ecosystem=general).\
