---
description: How to import repositories into Snyk from your SCM integration to start scanning your Projects
nav_context: new
---

# Import repositories

{% include "../../../.gitbook/includes/pilot-guide-navigation.md" %}

After you set up your SCM integration, import repositories to Snyk.

1. In the Snyk Web UI, use the scope selector at the top of the page to select your Organization.
2. Select **Projects** in the side menu, then click **Add projects**.

<figure><img src="../../../.gitbook/assets/import-repositories.png" alt="Starting a repository import in the Snyk Web UI"><figcaption></figcaption></figure>

3. Select the integration that you configured. Alternatively, navigate to **Settings** > **Integrations** > **All integrations** and select the integration from there.

<figure><img src="../../../.gitbook/assets/configure-organization-level-integration.png" alt="Starting an import from the integrations list"><figcaption></figcaption></figure>

4. Select the repositories that you want to import and click **Add selected repositories**.

<figure><img src="../../../.gitbook/assets/select-repositories-want-import-click-add-selected.png" alt="Selecting the repositories to import"><figcaption></figcaption></figure>

The import starts immediately. To monitor progress, click **View import logs** on the **Projects** page.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXcAbYAUF2UkQGgLSwX6fogOh42iosfEe_vyNjhY9wH-SOM_HZCQRxQNQRiI8jPGtcOaHP8ts3C8GoZpfRBLislwqtjgS_TuwUf01rH9gf6W0xxdC0Mq2Tflw3qDdTomfd5n6121?key=i_CNrr-DvB8PGUAzq09BT3pc" alt="Import progress shown in the import logs"><figcaption></figcaption></figure>

All imported repositories appear on the **Projects** page.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXeaN56eVc58YuGp2c_pQ2Eo_6D5G3ms6bWCs17pk1zYHXCrgPDY6mH6T-0wWCfmdnp9ot55q6f9TznZMBtpl_KYAsUxp78NBdiu1zraOY9fSp7ArsfANKxKoDBMkLqA3hVyBM9j?key=i_CNrr-DvB8PGUAzq09BT3pc" alt="Imported repositories listed on the Projects page"><figcaption></figcaption></figure>

Continue monitoring the import in the import logs, or select an imported repository to start reviewing its vulnerabilities.

{% hint style="info" %}
Additional resources

* [Import Project repository](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-projects/import-project-repository)
* [API-driven imports](https://docs.snyk.io/developer-tools/snyk-apps/tool-snyk-api-import)
{% endhint %}
