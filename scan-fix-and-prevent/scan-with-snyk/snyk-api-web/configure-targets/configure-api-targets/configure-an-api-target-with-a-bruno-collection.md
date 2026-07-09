# Configure an API target with a Bruno Collection

Use Bruno Collections to define API endpoints for scanning with Snyk API & Web.

{% hint style="info" %}
Snyk API & Web supports OpenCollection YAML (`.yml` files), the recommended format for Bruno v3.0.0 and later.
{% endhint %}

To configure Snyk API & Web to scan API endpoints with a Bruno Collection, complete these three steps:

1. Prepare the Bruno Collection.
2. Add an API target using the Bruno Collection.
3. Configure the API target with Bruno environment variables.

## Prepare the Bruno Collection

Prepare the Bruno Collection so the full sequence of requests runs from start to end without errors. This ensures the collection works correctly when you import it into Snyk API & Web.

Export the collection in OpenCollection format as a ZIP file for the next step. If the collection is already in a folder on your local machine, you can use the folder without zipping it.

## Add a Bruno Collection target

After you prepare the Bruno Collection, add an API target.

1. Navigate to **Targets** and click **Add**.
2. Complete the **Add target** form:
   * **Name**: Enter a meaningful identifier for your target.
   * **URL**: Enter the base URL for your API.
   * **API Type**: Select **API**, then select **Bruno Collection**.
   * **Upload method**: Choose one of the following:
     * **Folder**: Upload the Bruno collection folder from your local machine.
     * **ZIP file**: Upload a compressed ZIP file that contains the collection folder.
   * **Environment file**: Select the environment file. Bruno stores environment variables in the collection folder. Snyk API & Web can import these variables automatically. You can also configure them later in the **Scanner** tab.
3. Click **Add**.

## Configure Bruno environment variables

1. Navigate to the target **Settings**.
2. In the **Scanner** tab, locate **API Scanning Settings**.
3. Enter the required Bruno environment variables in the corresponding fields.

If your Bruno collection folder includes an `environments/` subfolder, Snyk API & Web automatically detects and imports the collection environment variables. If multiple environment files exist, select which environment to use.

You can view all imported variables in the interface. Enter values for secret variables before you start a scan. For better security, store secret variable values in the [credentials manager](https://help.probely.com/en/articles/13871348-how-to-manage-target-authentication-credentials-in-snyk-api-web).

## Add authentication

If you want to run authenticated scans for your API, you can use several methods:

* If your collection has static authentication defined, Snyk API & Web uses it. Supported methods are Basic Auth, Bearer Token, and API Key. Define authentication at the collection level. Snyk API & Web does not support folder-level or request-level authentication.
* If you dynamically generate your authentication token, configure authentication with scripts. See [Configure Bruno authentication](https://app.gitbook.com/s/QTiGglVLHTj5smrGKzeI/snyk-api-and-web-bruno-collection/configure-bruno-collection-authentication).

## Verify the configuration

After you complete the configuration, the target is ready to scan. [Test your configuration](https://help.probely.com/en/articles/14420138-how-to-test-target-configuration), then run a scan to verify that Snyk API & Web tests all requests in the collection.
