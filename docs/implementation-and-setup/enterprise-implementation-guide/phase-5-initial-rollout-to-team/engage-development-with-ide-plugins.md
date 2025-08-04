# Engage development with IDE plugins

## Introduce Snyk IDE plugins to your developers

Encouraging your developers to install a [Snyk IDE plugin](../../../developer-tools/snyk-ide-plugins-and-extensions/) is a key step in improving developer adoption and achieving “shift left,” where developers are more empowered and responsible for fixing issues before they are committed to pull requests or added to builds.

Snyk IDE plugins help developers to find and fix security vulnerabilities and issues. This allows individual developers to pass security reviews, avoid costly fixes later in development, and reduce time to develop and deliver secure code.

The webinar [Workflows for Developers using Snyk in the IDE and CLI](https://www.youtube.com/watch?v=jzUJS6S6H48) covers IDE usage in more detail, including a demo using an IDE plugin to review an issue.

## Implementation tips

Companies just getting started with security should start with the visibility and prevention stages to give clear direction on what issues must be fixed. As developers are fixing prioritized issues or starting new work, a natural opportunity arises to introduce an IDE plugin to validate fixes and proactively test to eliminate potential friction introduced by the security gates. This is where AppSec programs often move from reactive to proactive approaches.&#x20;

Companies with a mature AppSec program may give developers earlier access to Snyk IDE plugins as they are already familiar with security concepts. So rollout is less about resolving a backlog, and more about prevention of new issues and validation of fixes.

## Available plugins

The following Snyk plugins and extensions are available:

* [Eclipse plugin](../../../developer-tools/snyk-ide-plugins-and-extensions/eclipse-plugin/)
* [JetBrains plugins](../../../developer-tools/snyk-ide-plugins-and-extensions/jetbrains-plugin/)
* [Visual Studio extension](../../../developer-tools/snyk-ide-plugins-and-extensions/visual-studio-extension/)
* [Visual Studio Code extension](../../../developer-tools/snyk-ide-plugins-and-extensions/visual-studio-code-extension/)

The training session [Introduction to using Snyk in an IDE](https://learn.snyk.io/lesson/snyk-in-an-ide/) provides further guidance.

{% hint style="info" %}
&#x20;If your Snyk application is hosted on the Snyk EU or AU data center, then you must specify this in your IDE plugin settings with the [URL relevant to your region](../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#ides-urls).\
See [Regional hosting and data residency](../../../snyk-data-and-governance/regional-hosting-and-data-residency.md) for more details.
{% endhint %}
