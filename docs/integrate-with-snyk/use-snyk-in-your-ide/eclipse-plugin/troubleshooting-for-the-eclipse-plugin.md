# Troubleshooting for the Eclipse plugin

## Logs

To determine where plugin logs are stored, navigate to **Preferences** > **Language Servers** > **Logs** and find the **Snyk Language Server** row in Eclipse. As it can be disabled, you may need to enable it to retrieve the logs. You will find the logs either in the console or in the file based on the preference set.

To see additional plugin error logs:

1. Navigate to **Window** > **Show View** > **Others...**.
2. In **type text filer** search for **Error Log**.
3. Click **Open** to see the error log tab. If you group the tab view by plugin (kebab menu in the top right corner > **Group By** > **Plug-in**), the `io.snyk.eclipse.plugin` row shows any plugin errors.

## Unable to install plugin in Eclipse 2022.12 or newer

The following explains **failure of installation because of LSP4e dependency**.

The Snyk Eclipse plugin relies on the Eclipse LSP4e project. This project has decided to require JDK 17 as a runtime. In order to maintain backwards compatibility for a longer time, Snyk has put a limit on the maximum version, which is LSP4e release 0.20.6. This does not seem to be included in the standard Eclipse 2022.12 distribution.

Before you install the Snyk plugin, add the Eclipse LSP4e update site [https://download.eclipse.org/lsp4e/releases/0.20.6/](https://download.eclipse.org/lsp4e/releases/0.20.6/) to your Eclipse installation:

<figure><img src="../../../.gitbook/assets/image (424) (1).png" alt="Adding the Eclipse LSP4e update site"><figcaption><p>Adding the Eclipse LSP4e update site</p></figcaption></figure>

Then try or retry installation of the plugin, for example, through the Snyk Marketplace client, as described on the [Eclipse plugin ](./)main page.
