# User Documentation

This repository is a read-only copy of the [Snyk user documentation](https://docs.snyk.io) source files. It is updated automatically from the repository that publishes docs.snyk.io.

Do not open pull requests in this repository. Changes made here are overwritten by the next update and never reach docs.snyk.io.

## Contributing

### Internal contributor guidelines

If you are an internal Snyk contributor, draft your content in [`snyk/user-docs-internal`](https://github.com/snyk/user-docs-internal), the source repository for docs.snyk.io, and then submit it to the Docs team with the `/ship-it` Slack workflow. `/ship-it` is the single intake point for all documentation and release-notification requests: it creates the Jira ticket for you and sends you a confirmation DM with tracking links.

Before you run `/ship-it`, complete these steps so the submission does not fail partway through:

- Confirm you have write access to `snyk/user-docs-internal`.
- Connect Atlassian to Slack, which is required for Jira ticket creation.
- Draft your content in GitHub and open a pull request in `snyk/user-docs-internal`. Use a branch rather than a fork, to allow GitBook previews.
- Place pages under `docs/<section>/`, and update that section's `SUMMARY.md` if you are adding new pages.
- Follow the `snyk-docs-writing-rules` skill while drafting.

Then run `/ship-it` from anywhere in Slack, or from the Ship It app in your sidebar. Have your pull request URL ready and complete all the fields in one sitting, because the multi-step form does not save partial progress.

For details on the process, go to the internal [writing user documentation](https://snyksec.atlassian.net/wiki/spaces/DRC/pages/1819541615/Writing+user+documentation) page.

To ask questions or to report issues to the user docs team, use `/ship-it`. The legacy submission forms remain available only as a fallback under Workflows → "Contact the Docs team."

### External contributor guidelines

If you are an external contributor and you find an issue such as a typo or a broken link, [contact support](https://support.snyk.io) and include the page URL. Pull requests in this repository are not merged.

If you want to add an integration to Snyk, [apply to become a Snyk partner](https://partners.snyk.io/English/register_email.aspx).

## Security

For any security issues or concerns, go to [SECURITY.md](SECURITY.md).
