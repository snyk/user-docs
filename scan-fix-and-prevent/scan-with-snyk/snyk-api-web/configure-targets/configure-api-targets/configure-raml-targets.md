# Configure RAML targets

Convert RESTful API Modeling Language (RAML) definitions to work with Snyk API & Web for API scanning.

If you have an API defined with RAML, you can configure an API target from that RAML definition for Snyk API & Web to scan. The configuration involves two steps:

1. Change the file extension
2. Add the target as OpenAPI

## Change the file extension

Change the RAML file extension from `.raml` to `.yaml`.

## Add the target as OpenAPI

After you have the file with the new extension, create the target as an OpenAPI target:

1. Navigate to **Targets** and click **Add**.
2. Complete the **Add target** form:
   - **Name**: Enter a meaningful identifier for your target
   - **URL**: Enter the base URL for your API
   - **API Type**: Select **API**, then select **OpenAPI**
   - **OpenAPI schema upload**: Select this option
   - **File**: Choose the RAML file with the `.yaml` extension
3. Click **Add**.

Snyk performs all necessary conversions and creates the target, and you can then scan your RAML API.
