# Enable Snyk Code

When you create a new Organization, Snyk Code (SAST) scanning is disabled by default. Snyk recommends enabling Snyk Code before you import your first Projects to Snyk. If Snyk Code is enabled after a Project is imported, Snyk will not detect Snyk Code files.

1. Select the **Settings > Snyk Code** option.
2. Click the toggle to enable Snyk Code, and then click **Save changes.**

You can use the API endpoint [Enable/Disable the Snyk Code settings for an org](../../../../snyk-api/reference/sastsettings.md#orgs-org_id-settings-sast) to perform this operation at scale for a large number of organizations.
