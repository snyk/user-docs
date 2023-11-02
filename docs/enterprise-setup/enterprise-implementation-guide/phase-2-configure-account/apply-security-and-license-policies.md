# Apply security and license policies

Snyk Policies define how Snyk behaves when identifying issues. Policies give you a quick and automated way to identify, prioritize and triage issues. This saves valuable development time, and allows developers to take more responsibility and ownership for security, reducing the “noise” level.

See [Policies](../../../manage-issues/policies/) for more details.

## Security policies

Group administrators can define security policies, providing an automated way to identify certain issues or types of issues, and apply actions like changing the severity or ignoring the issue based on your conditions.&#x20;

* Configure policies to increase priority or decrease as needed.&#x20;
* Create ignores where needed

See [Security policies](../../../manage-issues/policies/security-policies/) for more details.

## License policies

Group administrators can set license policies to define Snyk behavior for treating license issues. For example, you can allow or disallow packages with certain license types, to avoid using packages containing incompatible licenses.

By default we determine the severity of licenses in the following way:

* High severity - licenses that definitely present issues for commercial software.
* Medium severity - licenses that have clauses that may be of concern and should be reviewed.

Configure policies to match your requirements.

See [Getting Started with Snyk License Compliance Management](../../../scan-using-snyk/start-scanning-using-the-cli-web-ui-or-api/scan-open-source-libraries-and-licenses/snyk-license-compliance-management.md) for more details.







