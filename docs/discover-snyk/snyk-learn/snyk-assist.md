# Snyk Assist for Learn

{% hint style="info" %}
Snyk Assist is available only in the Learning Management add-on offering. For more information, contact your Snyk account team.
{% endhint %}

Snyk Assist for Learn is an AI-powered learning assistant integrated into the Snyk Learn platform. It is designed to answer your Snyk product and application security questions instantly, helping you learn faster and resolve queries efficiently directly within your learning environment.

Snyk Assist enhances your learning experience within the Snyk Learn platform by:

* Providing immediate answers to questions about Snyk products, features, and general application security concepts.
* Delivering context-aware replies based on Snyk's extensive knowledge base.
* Suggesting relevant follow-on learning opportunities available within Snyk Learn, Snyk Docs and the Snyk Blog.

<figure><img src="../../.gitbook/assets/Screenshot 2025-04-24 at 14.52.22.png" alt=""><figcaption><p>Snyk Assist on Snyk Learn</p></figcaption></figure>

## How Snyk Assist Works

Snyk Assist utilizes Generative AI to respond to questions based on information retrieved from trusted Snyk sources:

* [Snyk Learn lessons](https://learn.snyk.io/catalog/?format=lesson)
* [Snyk Documentation](../../)

Snyk Assist **does not** have access to your Snyk tenant data, such as issue and asset information, or Snyk Learn learning history. While interacting with Snyk Assist, we collect and process your queries and conversation history to provide responses and improve our service.

The technology leverages Snyk-managed services and selected third-party Large Language Model (LLM) providers to respond with helpful and informative answers.

{% hint style="info" %}
**Capabilities and Limitations**

* Snyk Assist **is designed** to answer questions regarding Snyk functionality and security concepts based on its knowledge base derived from Snyk Learn and Snyk Docs.
* Snyk Assist **does not** perform security analysis, check your code for vulnerabilities, comment on code quality, or provide specific code samples. For code scanning and analysis, continue using the Snyk CLI, Snyk IDE extensions, or recurring SCM tests integrated with your repositories.
* Snyk Assist provides information from its body of knowledge, but responses should not be considered legal, security, or compliance advice. While we strive for accuracy, Snyk is not liable for decisions made based on information provided by Snyk Assist. Always verify critical security information through official Snyk documentation and tools.

See the usage disclaimer [here](https://snyk.io/policies/snyk-assist-disclaimer).
{% endhint %}

## Using Snyk Assist

Snyk Assist is accessible directly within the Snyk Learn interface for users with access to the platform, and where your Snyk admin has enabled the functionality.

1. Navigate to [Snyk Learn](https://learn.snyk.io).
2. Click the **Snyk Assist** icon <img src="../../.gitbook/assets/Screenshot 2025-04-25 at 14.50.54.png" alt="" data-size="line"> to open the chat window. This is found in the bottom right of the page.
3. Type your questions about Snyk products or application security concepts into the chat prompt.

Snyk Assist will answer based on its knowledge base, potentially including links to relevant documentation, Learn lessons, and Snyk blogs.

{% hint style="info" %}
If you do not see the chat icon, your Snyk admin may not have enabled this functionality for your Group. Additionally, ensure that the default Org in your Snyk settings is set to your company's Snyk Org and not a personal Snyk Org.
{% endhint %}

## Enabling Snyk Assist

Snyk Group Admins can control the availability of Snyk Assist for users within their scope.

To enable Snyk Assist for your Group:

1. Log in to the Snyk Web UI.
2. Navigate to the Group settings page for the Group you wish to manage.
3. Toggle on **Snyk Assist** to enable Snyk Assist for users within that Group.

After Snyk Assist is enabled, it is visible and accessible on Snyk Learn for the users belonging to that specific Group.

<figure><img src="../../.gitbook/assets/Screenshot 2025-04-29 at 17.53.20.png" alt=""><figcaption><p>Snyk Assist settings page</p></figcaption></figure>

## Data Handling and Safeguards

See [How Snyk handles your data](../../snyk-data-and-governance/how-snyk-handles-your-data.md#snyk-learn) for more information on this topic.
