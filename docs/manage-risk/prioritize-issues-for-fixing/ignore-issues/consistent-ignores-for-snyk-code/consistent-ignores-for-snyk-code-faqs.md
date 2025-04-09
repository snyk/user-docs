# Consistent Ignores for Snyk Code FAQs

{% hint style="info" %}
**Release status**

Snyk Code Consistent Ignores is in Early Access and available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).

To make sure Snyk Code Consistent Ignores Early Access meets your needs and requirements, review [Known limitations](known-limitations.md) and [FAQ](consistent-ignores-for-snyk-code-faqs.md) sections.
{% endhint %}

This FAQ section addresses common concerns about the Snyk Code Consistent Identity Early Access program. You can share feedback with your Snyk account team about these items. Snyk will consider all suggestions but implementation or addressing these issues is not guaranteed for the upcoming GA version.

## Ignore limitations outside my Snyk Organization

Ignores only work within the Organization where they're defined. You need to run tests in the Organization where you stored an ignore for it to be taken into account. This is also valid for Snyk IDE and CLI environments where developers work in repositories that span multiple Snyk organizations.

Depending on feedback during the Early Access period, we may offer a broader scope for ignores beyond an individual Organization.

## Existing DeepCode inline ignores (legacy) are not supported or migrated

Deepcode inline ignores are a legacy feature that is only available for certain customers. If you have pre-existing Deepcode inline ignores, Snyk will remove them from test results. They will not be marked as `Open` or `Ignored`.

### Recommendation

Recreate these ignores using the new Snyk Code Consistent Ignore process, either using Snyk API or through the Projects page. Reach out to your Snyk account team or Snyk product management to share feedback on what will make this process easier for you.

## Repository renames may result in ignores being lost

Snyk may fail to complete testing after you rename a repository, depending on whether the underlying SCM supports redirects. If Snyk successfully runs subsequent tests (e.g., GitHub), ignores may not be applied.

### Recommendation

**Before renaming**

[Convert all ignores](./#convert-project-scoped-ignores-to-asset-scoped-ignores) to apply to the entire repository (consistent ignores).

**After renaming**

1. Delete all targets associated with that repository.
2. Reimport the newly renamed repository.

Previous Consistent Ignores will apply to the newly named repository. New clones in IDEs/CLI that reference the new name will also take into account ignores, even with the old git URL, in case some developers haven't updated their remote repositories.

## Granular ignores

Using this feature, you can no longer apply an ignore to a single Snyk Project. This means that ignores cannot be applied to a specific branch within a single Organization.

If you have specific use cases that require this functionality, reach out to your Snyk account team or Snyk product management for feedback.

## Project attribute policies

Policies defined against Project attributes will continue to work within Snyk Projects where the attributes match. However, they will not apply across the repository to other Projects or in Snyk IDE, CLI, or PR checks flows. To apply policies across Projects and branches for the same repository, define them against Organizations.

## CI/CD support for snyk test --code

Most native Snyk CI/CD plugins (e.g., Jenkins, AWS Pipelines) do not support Snyk Code. As a workaround, some users have been applying the `--code` flag to the standard `snyk test` command to invoke a Snyk Code scan instead of a Snyk Open Source scan. Snyk Code Consistent Ignores does not support this workflow.
