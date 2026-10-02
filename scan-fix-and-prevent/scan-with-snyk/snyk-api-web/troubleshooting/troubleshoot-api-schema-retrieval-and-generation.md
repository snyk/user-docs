---
description: >-
  Resolve common issues when retrieving or generating API schemas from
  repositories using Snyk API & Web.
---

# Troubleshoot API schema retrieval and generation

When you connect an OpenAPI target to a code repository, Snyk API & Web automatically fetches your repository content, searches for existing OpenAPI or Swagger specifications, and, if your account has AI usage enabled, generates prospective schemas directly from your source code. This background process runs asynchronously and notifies you by email when it completes.&#x20;

{% hint style="info" %}
Repository-based schema generation supports only OpenAPI formats.
{% endhint %}

## The Scan button is disabled

Newly created repository targets cannot be scanned until a schema is attached.

To resolve this, navigate to the **Scanner** tab in your target settings and choose an available schema from the list. If only a single schema is surfaced, Snyk API & Web links it automatically and enables scanning.

## You received "API schemas not obtained"

This message indicates that Snyk API & Web identified API code in your repository, but could not produce a usable schema.

Common causes include:

* Missing specs and disabled AI: your repository lacks OpenAPI or Swagger files, and your account does not have AI usage enabled.
* Failed validation: Snyk generated or found candidate schemas, but they failed structural validation checks.

To fix this issue, do one of the following:

* In the **Scanner** tab, upload a valid schema file or specify a schema URL directly.
* Contact Support to enable AI usage for your account.
* Click **Update schema** in the **Scanner** tab to restart retrieval against your latest code.

## A schema in your repository is not listed

Only schemas that pass validation are surfaced in the UI. If an existing OpenAPI specification file in your repository is missing from the **Scanner** tab, it did not pass validation.

Validate your file against OpenAPI specifications using tools like the [Swagger Editor](https://editor.swagger.io/) to identify and fix syntax errors. After resolved, commit the updated file and trigger schema retrieval again from the **Scanner** tab, or manually upload the corrected file.

If the specification validates successfully but still fails to display, contact Support with your target ID and file path.

## You received "APIs not found in repository"

Snyk API & Web cannot locate the API code within the selected repository. This typically occurs when the API is located in a different repository.

Use the **Change repository** option in the **Scanner** tab to select the correct repository.

## Your repository is not in the list

The repository selection list comes from your Snyk account. To make the repository available in Snyk API & Web, import it into Snyk, and then refresh the page.

## Your schemas are still generating

The automated schema discovery and generation process takes less than 15 minutes, depending on repository size.

You can safely navigate away from the platform. Snyk emails you when your schemas are ready. If processing takes more than 30 minutes, contact Support with your repository details and target ID.

## You cannot change the repository or request a new schema

These actions are temporarily disabled for the following reasons:

* A schema generation task is running for this target. Snyk unlocks these actions after the job completes.
* You hit a rate limit due to another recently submitted schema request. Wait before requesting additional updates.

## Your generated schema is incomplete or wrong

AI-generated schemas serve as initial draft representations inferred from source code. They can omit complex structures like dynamic routing or custom middleware.

Download the generated schema from the **Scanner** tab, modify the schema locally, and re-upload the edited specification.

If generated output consistently omits endpoints, contact Support with your framework specifications and target ID.

## Your generated schema shows “no servers defined”

This warning means the schema does not say which host or base URL the API lives on, which is common for schemas built from code, since the routes are in the repository, but the deployed URL usually is not.&#x20;

This is not an error. The generated schema is valid, and you can use it to scan your target.

## Related information

* [API targets](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-api-web/configure-targets/configure-api-targets)
* [What is a target?](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-api-web/configure-targets/what-is-a-target)
* [Network timeout errors](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-api-web/troubleshooting/network-timeout-errors)
