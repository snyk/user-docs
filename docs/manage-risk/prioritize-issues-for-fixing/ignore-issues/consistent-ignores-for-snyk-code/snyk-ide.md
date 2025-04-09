# Snyk IDE

{% hint style="info" %}
**Release status**

Snyk Code Consistent Ignores is in Early Access and available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).

To make sure Snyk Code Consistent Ignores Early Access meets your needs and requirements, review [Known limitations](known-limitations.md) and [FAQ](consistent-ignores-for-snyk-code-faqs.md) sections.
{% endhint %}

When you run tests in any of the [four supported Snyk IDE plugins](../../../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/), the plugins will take into account your ignores.

## **Minimum version required**

Snyk Code Consistent Ignores works best with the latest IDE plugin versions.

| IDE           | Minimum version required |
| ------------- | ------------------------ |
| VS Code       | 2.5.0                    |
| IntelliJ      | 2.7.14                   |
| Visual Studio | 2.0.0                    |
| Eclipse       | 3.0.0                    |

## **Set up the Organization**

To take ignores into account, specify the Organization where the ignores reside. [Group-level policies also cascade down to all Organizations](broken-reference). See [How to select the Organization to use in IDE plugins (Visual Studio Code example)](../../../../snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli.md).

## Snyk IDE default ignore behavior

The IDE display output hides ignored results by default to maintain developer focus.&#x20;

## View ignores in Snyk IDE

You can apply filters in the plugin settings to show ignored results alongside open results or in isolation. When you set ignored issues to display, the issues and their details will appear in the plugin.

<figure><img src="../../../../.gitbook/assets/snyk-code-ignored-issue-ide.png" alt=""><figcaption><p>View ignores in Snyk IDE</p></figcaption></figure>
