# Classification policy - Use case

You can use the **Set Asset Class** action from the Policies view to classify the assets based on importance, where class A is the most important and class D is the least important.&#x20;

You can set the asset class based on:&#x20;

* the repository name &#x20;
* the asset tags

{% hint style="info" %}
Snyk AppRisk identifies GitHub and GitLab topics as asset tags.
{% endhint %}

Use the classification policy to give business context to your application. When you set up a classification policy, all your assets are automatically classified.

If you just started using the classification policy, the recommendation is to focus first on the Class D assets, since they are the least important.

The following example filters the assets that contain `sandbox`, `test`, and `to-delete` in their names. In Snyk AppRisk, GitHub and GitLab topics are pulled in from the SCM integration and applied to repository assets, so if topics like `PCI-Compliance` have been added to repos in the SCM, Snyk can take those tags in Snyk AppRisk and classify those assets as Class A.

<figure><img src="../../../../.gitbook/assets/image (7).png" alt="AppRisk - Setting up filters for a classification policy"><figcaption><p>Snyk AppRisk - Setting up filters for a classification policy</p></figcaption></figure>

After you set up the filters, you need to apply a Class D asset classification to those assets.

<figure><img src="../../../../.gitbook/assets/image (8).png" alt="AppRisk - Setting up actions for a classification policy"><figcaption><p>Snyk AppRisk - Setting up actions for a classification policy</p></figcaption></figure>

You can apply a similar pattern and create actions for Class A, B, and C assets, within the same policy.

<figure><img src="../../../../.gitbook/assets/image (9).png" alt="AppRisk - Setting up multiple actions for a classification policy"><figcaption><p>Snyk AppRisk - Setting up multiple actions for a classification policy</p></figcaption></figure>

