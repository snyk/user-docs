# Initial team rollout

Invite your stakeholders to explore Snyk features and integrate security into their workflows. After this step, your teams can fix issues, monitor pipelines, and manage vulnerabilities using integrations like Jira.

Follow these steps to roll out Snyk to your teams:

1. [Configure notifications:](initial-team-rollout.md#configure-notifications) Set up email alerts to ensure users receive relevant information without being overwhelmed.
2. [Announce Snyk:](initial-team-rollout.md#announce-snyk-to-your-teams) Communicate the rollout to developers using standardized templates.
3. [Provide training:](initial-team-rollout.md#provide-developer-training) Direct users to Snyk Learn for product and security education.
4. [Deploy IDE plugins:](initial-team-rollout.md#engage-development-with-ide-plugins) Enable developers to find and fix issues locally before they commit code.

## Configure notifications

{% hint style="success" %}
**Key decision**: Determine the notification volume. Snyk recommends disabling all email notifications during the initial import to prevent alert fatigue.
{% endhint %}

Managing notifications ensures that developers only see high-priority issues that require action.

* Instruct administrators to manually enable the critical alerts through their personal settings if they need to monitor progress.
* Once the environment is stable, enable notifications in bulk for **High** and **Critical** severities only.
* Disable all email notifications for new Organizations.

Navigate to **Group** > **Settings** to view the notification defaults overview.

## Announce Snyk to your teams

{% hint style="success" %}
**Key decision**: Assess the current comfort level of your development team. If they are new to security automation, disable intrusive features like Automatic PRs until they have completed initial training.
{% endhint %}

Use these templates to introduce Snyk. Replace the bracketed text with your specific details.

### Email template

|                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <p><em>To: Developers</em></p><p><em>Subject: Launching Snyk at [Company name]</em></p><p><em>Hi all,</em></p><p><em>I’m excited to announce that we’re implementing Snyk at [Company name]</em></p><p><em>[optional: add personalized video, if desired]</em></p><p><em>Snyk will help us [enter your goal(s)].</em></p><p><em>As part of the launch process, we’ll invite you to a short “Intro to Snyk” and Q&#x26;A session to learn more about Snyk and the products we’re implementing. You’ll also have the opportunity to attend a developer training session and get access to Snyk Learn for self-paced tutorials to help you get started.</em></p><p><em>We’re looking forward to building secure applications together, with less frustration and interruption to your workflows for addressing security issues.</em></p><p><em>More info can be found at [hyperlink to your internal resource page/wiki with more info].</em></p><p><em>Regards,</em></p><p></p><p><em>_____ [Sender]</em></p> |

### Instant message template

|                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <p><em>To: Developers</em></p><p><em>Subject: Launching Snyk at [Company name]</em></p><p><em>Hi all,</em></p><p><em>I’m excited to announce that we’re implementing Snyk at [Company name]</em></p><p><em>[optional: add personalized video, if desired]</em></p><p><em>Snyk will help us [enter your goal(s)].</em></p><p><em>As part of the launch process, we’ll invite you to a short “Intro to Snyk” and Q&#x26;A session to learn more about Snyk and the products we’re implementing. You’ll also have the opportunity to attend a developer training session and get access to Snyk Learn for self-paced tutorials to help you get started.</em></p><p><em>We’re looking forward to building secure applications together, with less frustration and interruption to your workflows for addressing security issues.</em></p><p><em>More info can be found at [hyperlink to your internal resource page/wiki with more info].</em></p><p><em>Regards,</em></p><p><br></p><p><em>_____ [Sender]</em></p> |

## Provide developer training

{% hint style="success" %}
**Key decision**: Match the training to the user persona. Use security education for novices and product training for those needing to master Snyk-specific workflows.
{% endhint %}

* [Snyk Learn security education](https://learn.snyk.io/catalog/security-education/?type=security-education): Teaches general security concepts, such as NoSQL injection and server-side request forgery.
* [Snyk Learn product training](https://learn.snyk.io/catalog/product-training/?type=product-training): Provides role-based learning paths for developers and administrators on the Snyk platform.

## Engage development with IDE plugins

{% hint style="success" %}
**Key decision**: Evaluate the maturity of your AppSec program. For new programs, introduce plugins as a tool to validate fixes for prioritized issues. For mature programs, provide immediate access to prevent new issues from entering the codebase.
{% endhint %}

Snyk IDE plugins allow developers to find and fix vulnerabilities before they reach the CI/CD pipeline. This shift-left approach reduces the time spent on security reviews.

1. Identify the primary IDEs used by your teams (VS Code, JetBrains, Visual Studio, or Eclipse).
2. Provide installation guides for the relevant Snyk IDE extension.
3. Configure regional hosting: If your application is on the EU or AU data center, specify the regional URL in the plugin settings.
