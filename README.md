# User Documentation

This repository is a read-only copy of the [Snyk user documentation](https://docs.snyk.io) source files. It is updated automatically from the repository that publishes docs.snyk.io.

Do not open pull requests in this repository. Changes made here are overwritten by the next update and never reach docs.snyk.io.

## Contributing

### Internal contributor guidelines

If you are an internal Snyk contributor, such as a PM or SME, request docs with the `/ship-it` Slack workflow. `/ship-it` is the single intake point for all documentation and release-notification requests. You don't need to write the draft yourself.

1. Run `/ship-it` from anywhere in Slack, or from the Ship It app in your sidebar, and add a link to the PRD or technical spec. Complete all the fields in one sitting, because the multi-step form does not save partial progress. Connect Atlassian to Slack first, because the workflow creates a Jira ticket.
2. The ship-it automation drafts the docs as a draft pull request in [`snyk/user-docs-internal`](https://github.com/snyk/user-docs-internal), the source repository for docs.snyk.io, and replies in the Slack thread with the link.
3. Review the draft for technical accuracy, and resolve every `[ACTION REQUIRED]` item. A technical writer does the final review and merges it.

For a minor update, or if you prefer to write the draft by hand, open a pull request in `snyk/user-docs-internal` instead. Use a branch rather than a fork, to allow GitBook previews, and place pages under `docs/<section>/`. Then run `/ship-it` and add the pull request URL in the **Docs Draft** field.

For details on the process, go to the internal [writing user documentation](https://snyksec.atlassian.net/wiki/spaces/Docs/pages/1819541615/Writing+user+documentation) page.

To ask questions or to report issues to the user docs team, use `/ship-it`. The legacy submission forms remain available only as a fallback under Workflows → "Contact the Docs team."

### External contributor guidelines

If you are an external contributor and you find an issue such as a typo or a broken link, [contact support](https://support.snyk.io) and include the page URL. Pull requests in this repository are not merged.

If you want to add an integration to Snyk, [apply to become a Snyk partner](https://partners.snyk.io/English/register_email.aspx).

## Security

For any security issues or concerns, go to [SECURITY.md](SECURITY.md).
