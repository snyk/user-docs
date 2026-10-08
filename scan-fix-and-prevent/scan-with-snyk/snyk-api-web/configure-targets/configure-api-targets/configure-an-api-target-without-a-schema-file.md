---
description: >-
  Add an OpenAPI target directly from a source code repository, without
  providing a schema file, using Snyk API & Web.
---

# Configure an API target without a schema file

{% hint style="info" %}
This feature is in Early Access. Contact your Snyk account team if you want to enable it.
{% endhint %}

You can now add OpenAPI targets by connecting them to a source code repository instead of providing a schema URL or file. Snyk API & Web automatically searches your repository for existing OpenAPI or Swagger files and, if your account has AI usage enabled, generates candidate schemas from your application code.

This eliminates the need to manually create or maintain API specifications. You can start scanning your APIs faster, even when documentation is missing or outdated.

## Prerequisites

To add an API target without a schema file:

* Connect your Snyk account at **Settings > Integrations > Snyk**. This integration gives Snyk access to your repositories.
* Make the repository containing your API code accessible through the Snyk integration (GitHub, GitLab, or Bitbucket).
* You must have permissions to add and configure targets in Snyk API & Web.
* To generate schemas from your code, your account must have AI usage enabled. If AI usage is not enabled, contact Support.

## Add and configure an OpenAPI target

Add an OpenAPI target by connecting it to a repository, then review and link a schema to enable scanning. Schema retrieval and generation run in the background, so there is a wait time between adding the target and being able to scan it.

### Start adding an OpenAPI target

1. Navigate to **Targets** in Snyk API & Web.
2. Click **Add** target.
3. Select **API** as the target type.
4. Select **OpenAPI** as the API format.
5. Choose **Repository** as the schema supply method and click **Select repository**.
6. On the repository selection screen, search or filter by Group and/or Organization to find the repository containing your API code.
7. Select one repository from the list and click **Select**.
8. Click **Add** or **Add & Configure**.

{% hint style="info" %}
Adding an OpenAPI target following these steps automatically enables the SAST/DAST integration for the selected repository.
{% endhint %}

### Wait for schema retrieval and generation

After you add the target, Snyk API & Web starts an asynchronous process:

1. Search for existing schemas: Snyk searches for existing OpenAPI or Swagger files (YAML or JSON) in the repository.
2. Schema generation: if AI usage is enabled, an LLM analyzes your source code to generate candidate OpenAPI schemas.

After retrieval and generation, Snyk validates all schemas before presenting them for review.

### Review and link a schema

When schemas are ready, you'll receive an email with a link to the target settings. Alternatively, navigate to the target manually:

1. Navigate to **Targets** and identify your target.
2. Open the **Scanner** tab in target settings.
3. Review the list of available schemas. Download, edit, and re-upload the schema as needed.
4. If Snyk API & Web finds or generates only one schema, it links that schema automatically. If multiple schemas are available, select the correct one and click **Link** to enable scanning.

### Configure authentication and scan

After linking a schema, you can configure authentication and other target settings:

1. In the **Authentication** tab, configure **API Target Authentication** as needed for your API.
2. (Optional) Click **Test configuration** to verify your settings.
3. Start a **Scan** to begin testing your API for vulnerabilities.

## Change the repository or regenerate schemas

### Change the linked repository

If you selected the wrong repository or your code has moved:

1. Navigate to the **Scanner** tab of your target settings.
2. Click **Change repository**.
3. Select a different repository and confirm.

Snyk re-runs schema retrieval and generation for the new repository. Your existing schema remains linked until a new one is available.

### Regenerate schemas from the current repository

If your code has changed and you want updated schemas:

1. Navigate to the **Scanner** tab of your target settings.
2. Click **Update schema**.

Snyk searches for existing schemas and (if AI usage is enabled) regenerates schemas from the updated code.

## Next steps

* Configure scheduled scans: set up automatic scanning on a schedule (daily, weekly, etc.).
* Review scan results: once your first scan completes, review detected vulnerabilities and prioritize remediation.
* Add more API targets: repeat this process to add other APIs from your repositories.
* Troubleshooting: if you encounter issues with schema retrieval or generation, visit [Troubleshoot API schema retrieval and generation](../../troubleshooting/troubleshoot-api-schema-retrieval-and-generation.md).
