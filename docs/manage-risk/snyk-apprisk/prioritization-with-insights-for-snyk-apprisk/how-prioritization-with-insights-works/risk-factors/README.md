# Risk factors

By understanding your images, packages, and Kubernetes resources as "application context", Snyk can compute the following risk factors:

* [Deployed](deployed.md)
* [Loaded package](loaded-package.md)
* [OS condition](os-condition.md)
* [Public facing](public-facing.md)

{% hint style="warning" %}
**Release status** \
The **Loaded package** risk factor is currently in Closed Beta and available only for Snyk AppRisk Pro plans.&#x20;

Contact your account manager if you are interested in Snyk AppRisk Pro.
{% endhint %}

You can enable and disable all of these "application context" risk factors through the Group **Settings**, on the **Insights** UI tab. If you choose to disable a risk factor, a provider selection, or the Kubernetes cluster mapping, Snyk will no longer compute them.&#x20;

Depending on the integration options enabled for your application, risk factors are applied differently. You can [prioritize your integrations](../../prioritization-setup/#prioritize-your-integrations) by customizing the available Insights options from the Group settings.

<figure><img src="../../../../../.gitbook/assets/image (1).png" alt="Insights tab in the Group settings"><figcaption><p>Insights tab in the Group settings</p></figcaption></figure>

{% hint style="info" %}
Risk factor settings may take up to four hours to take effect.
{% endhint %}
