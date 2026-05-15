---
hidden: true
noIndex: true
---

# Consistent Ignores for Snyk Code FAQs

This FAQ section addresses common concerns about the Snyk Code Consistent Identity Early Access program. You can share feedback about these items with your Snyk account team.&#x20;

## Ignore limitations outside my Snyk Organization

Ignores only work within the Organization where they're defined. You need to run tests in the Organization where you stored an ignore for it to be taken into account. This is also valid for Snyk IDE and CLI environments where developers work in repositories that span multiple Snyk organizations.

Depending on feedback during the Early Access period, we may offer a broader scope for ignores beyond an individual Organization.

## Existing DeepCode inline ignores (legacy) are not supported or migrated

DeepCode inline ignores are a legacy feature that is only available to certain customers. If you have pre-existing DeepCode inline ignores, Snyk removes them from the test results. The results are not marked as `Open` or `Ignored`.

### Recommendation

Recreate these ignores using the new Snyk Code Consistent Ignore process, either using Snyk API or through the Projects page. Reach out to your Snyk account team or Snyk product management to share feedback on what will make this process easier for you.

## Repository renames may result in ignores being lost

Snyk may fail to complete testing after you rename a repository, depending on whether the underlying SCM supports redirects. If Snyk successfully runs subsequent tests (e.g., GitHub), ignores may not be applied.

### Recommendation

**Before renaming**

[Convert all ignores](convert-project-scoped-ignores-to-asset-scoped-ignores.md) to apply to the entire repository (consistent ignores).

**After renaming**

1. Delete all targets associated with that repository.
2. Reimport the newly renamed repository.

Previous Consistent Ignores are applied to the newly named repository. New clones in IDEs/CLI that reference the new name take into account the ignores, even with the old git URL, in case some developers haven't updated their remote repositories.

## Granular ignores

Using this feature, you can no longer apply an ignore to a single Snyk Project. This means that ignores cannot be applied to a specific branch within a single Organization.

If you have specific use cases that require this functionality, reach out to your Snyk account team or Snyk product management for feedback.

## Project attribute policies

Policies defined against Project attributes will continue to work within Snyk Projects where the attributes match. The policies are not applied across the repository to other Projects or in Snyk IDE, CLI, or PR checks flows. To apply policies across Projects and branches for the same repository, define them against Organizations.

## CI/CD support for snyk test --code

Most native Snyk CI/CD plugins (e.g., Jenkins, AWS Pipelines) do not support Snyk Code. As a workaround, some users have been applying the `--code` flag to the standard `snyk test` command to invoke a Snyk Code scan instead of a Snyk Open Source scan. Snyk Code Consistent Ignores does not support this workflow.
