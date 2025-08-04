# Organizations

Organizations represent business areas such as teams, products, or environments. An Organization contains [Snyk Projects](../../snyk-projects/) that contain the code team members are scanning.

The settings and policies of an Organization influence scan results, depending on which Organization is used when you import a Project, the default Organization, a personal Organization, or another Organization you specify.

When you sign up with Snyk, you have a default Organization in your Enterprise Group, named like your GitHub username. Any Projects you import using the CLI or GitHub integration appear in this Organization by default.

You can create Organizations in your Enterprise Group, and also create personal Organizations in your personal space. This allows you to monitor your own personal Projects outside of your Enterprise Group or in a sandbox space.

You can create an Organization or join an Organization by invitation. If you have more than one Organization, you can switch between Organizations in the Snyk WebUI or using a CLI option See [**Manage Organizations**](create-and-delete-organizations.md) for details.

Snyk sends notifications about newly disclosed vulnerabilities by Organization, You can turn notifications on or off for each Organization.
