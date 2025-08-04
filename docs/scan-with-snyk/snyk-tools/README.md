# Snyk Tools

## Scope of Snyk Tools

Snyk Tools help with specific pain points that may not be addressed by Snyk product functionality, regardless of whether you use Snyk through the Web UI, CLI, API, or an integration. Snyk Tools extend the functions of the Snyk API and the Snyk CLI.

{% hint style="info" %}
You must have a [Snyk Account](https://snyk.io/login?cta=sign-up\&loc=nav\&page=support_docs_page) with populated Projects to use Snyk Tools.
{% endhint %}

## Key Snyk Tools

Snyk provides full documentation for the following key Snyk Tools:

* [snyk-api-import (docs)](tool-snyk-api-import/): Bulk import Projects into Snyk in a robust, paced way. Repo: [snyk-api-import](https://github.com/snyk/snyk-api-import)
* [snyk-delta (docs)](../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-delta.md): Get the delta between two Snyk snapshots. Repo: [snyk-delta](https://github.com/snyk-tech-services/snyk-delta)
* [snyk-filter (docs)](../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md): Takes the JSON output from the Snyk CLI and applies custom filtering of the results. Repo: [snyk-filter](https://github.com/snyk-tech-services/snyk-filter)
* [snyk-to-html](../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md): Snyk JSON to HTML Mapper takes the JSON output from `snyk test --json` and creates a local HTML file displaying the vulnerabilities discovered. Repo: [snyk-to-html](https://github.com/snyk/snyk-to-html)
* [jira-tickets-for-new-vulns (docs)](tool-jira-tickets-for-new-vulns.md): Sync your Snyk-monitored Projects and auto-open JIRA tickets for issues. Repo: [jira-tickets-for-new-vulns](https://github.com/snyk-tech-services/jira-tickets-for-new-vulns)
* [snyk-scm-contributors-count (docs)](../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-scm-contributors-count/): Count contributors for your SCM repos with commits in the last 90 days. Repo: [snyk-scm-contributors-count](https://github.com/snyk-tech-services/snyk-scm-contributors-count).

## Additional Snyk Tools

Refer to the repositories for instructions on how to use the following additional Snyk Tools.

* [snyk-disallow](https://github.com/snyk-tech-services/snyk-disallow): Get a viewer token for the Snyk Group to get a read|test-only token for CI or similar systems.
* [snyk-prevent-gh-commit-status](https://github.com/snyk-tech-services/snyk-prevent-gh-commit-status): POST commit status of a PR the result of [snyk-delta](https://github.com/snyk-tech-services/snyk-delta) executed in the CI.
* [snyk-cr-monitor](https://github.com/snyk-tech-services/snyk-cr-monitor): Gather Docker repos to test, then iterate through results to run multiple jobs simultaneously.
* [backstage-plugin-snyk](https://github.com/snyk-tech-services/backstage-plugin-snyk): Plugin to display security details from Snyk.
* [snyk-api-ts-client](https://github.com/snyk-tech-services/snyk-api-ts-client): Snyk API Typescript client.
* [snyk-transitive-ignore](https://github.com/snyk-tech-services/snyk-transitive-ignore): Generate the Snyk ignore policy dynamically based on a provided list of packages.
* [snyk-user-sync-tool](https://github.com/snyk-tech-services/snyk-user-sync-tool): Add, remove, and sync user memberships.
* [snyk-licenses-texts](https://github.com/snyk-tech-services/snyk-licenses-texts): Provides Organization-level licenses used, copyrights (until January 8, 2024), and dependencies data.
* [snyk-request-manager](https://github.com/snyk-tech-services/snyk-request-manager): Rate-controlled and retry-enabled request manager to interact with Snyk APIs.
* [snyk-repo-issue-tracker](https://github.com/snyk-tech-services/snyk-repo-issue-tracker): A python script/module that allows for generating a changeset of issues between runs against the Snyk Project issues API.
*   [snyk-repo-diff:](https://github.com/snyk-tech-services/snyk-repo-diff) Helps determine which repositories are not monitored by Snyk.

    This works by retrieving a list of all Projects in a given Snyk Group, that is, all Projects in all Organizations belonging to the same Snyk Group, and associating them with a list of repositories found in a given GitHub Organization.
* [snyk-issues-to-csv](https://github.com/snyk-tech-services/snyk-issues-to-csv): A python script that uses the PySnyk module along with the Pandas modules to collect all issues from the report API and combine them into a single CSV for an entire Group.
* [snyk-bulk](https://github.com/snyk-tech-services/snyk-bulk): Recursively scan source repositories for open-source vulnerabilities with the Snyk CLI outside of a build environment.
* [snyk-bulk-action-scripts](https://github.com/snyk-tech-services/snyk-bulk-action-scripts): A collection of scripts to edit integration settings for every Organization in a Group in Snyk.
* [snyk-deps-to-csv](https://github.com/snyk-tech-services/snyk-deps-to-csv): Collect all dependencies from all Organizations in a Group and output to a CSV file.
* [add-ignore-reason-to-csv-report](https://github.com/snyk-labs/add-ignore-reason-to-csv-report)**:** Add ignore reason and user data to CSV ignore report from the UI.

## Tool ideas

Do you have an idea for a tool? If so, check out [Snyk Apps](../../snyk-api/using-specific-snyk-apis/snyk-apps-apis/), which provides an opportunity to mold your Snyk experience to suit your specific needs. You can also contact [Snyk Support](https://support.snyk.io) with questions.
