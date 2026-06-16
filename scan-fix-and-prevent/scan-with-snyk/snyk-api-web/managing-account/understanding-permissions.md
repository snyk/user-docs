# Understanding permissions

Learn what actions each permission grants to better configure your custom roles.

This article provides a detailed breakdown of the high-level permissions within Snyk API & Web, explaining what actions each permission grants. Permissions are then grouped into roles (either built-in or custom) and, along with a scope, dictate the actions a user can perform.

## Role and scope structure

Roles can be applied at three levels, dictating the scope of the actions a user can perform:

* **Account:** actions apply across the entire account.
* **Team:** actions apply only to the selected team.
* **Target:** actions apply only to the selected target.

To learn more about roles, visit [Roles and permissions](roles-and-permissions.md).

## Detailed permission breakdown

The following table lists the Name, ID, and a detailed description of the actions allowed for each high-level permission:

| Permission | Allowed Actions |
| --- | --- |
| **Name:** Account Settings<br>**ID:** account\_settings | **Manage Integrations:**<br>* Create, view, change, and delete third-party account integrations (for example, Akamai, Azure DevOps, Jira Cloud, and so on).<br><br>**Manage Labels:**<br>* Create, view, change, delete, and list Finding Labels, Target Labels, and User Labels.<br><br>**Manage Automation:**<br>* Create, view, change, delete, and list Webhooks. |
| **Name:** Audit Log<br>**ID:** audit\_log | **Review History:**<br>* Obtain the Audit Log entries. |
| **Name:** Billing<br>**ID:** billing | **Manage and View Billing:**<br>* Manage billing and payment information.<br>* List and download invoices. |
| **Name:** Change Finding<br>**ID:** change\_finding | **Modify Findings:**<br>* List and view findings.<br>* Perform bulk operations on findings.<br>* List and assign users to a finding.<br>* Add notes to findings.<br>* Manually synchronize findings with integrations already configured (for example, Azure DevOps, Jira Cloud, Shortcut).<br>* View Target Labels. |
| **Name:** Change Finding State<br>**ID:** change\_finding\_state | **Modify Findings State:**<br>* Change finding state (Accept risk, Mark as invalid, Reset).<br>* Change finding review status (Approved, Rejected). |
| **Name:** Change risk<br>**ID:** change\_risk | **Modify Risk Level:**<br>* Change the risk rating associated with a finding. |
| **Name:** Change Target Settings<br>**ID:** change\_target\_settings | **Manage Target Configuration:**<br>* Manage (add, change, delete, and view) Login or Logout configuration, Navigation Sequences, Partial Scans, Scanning Agents, Extra hosts, Seeds or Reject Lists, Custom Headers or Cookies, Blackout Period, Report types, Coverage details, and Technologies.<br>* Manage (add, change, delete, view) Scan Profiles, and assign them to the target.<br>* Configure API Schema files.<br>* Manage target-specific integrations (Azure DevOps, DefectDojo, Jira Cloud or Server, Shortcut, Slack).<br><br>**Manage Domains:**<br>* Manage (add, change, list, view, and verify) Domains.<br><br>**Manage Labels:**<br>* Create, change, delete, and list Finding Labels and Target Labels.<br><br>**Manage Webhooks:**<br>* Configure Scope Webhooks. |
| **Name:** Correlation Admin<br>**ID:** correlation\_admin | **Manage SAST and DAST Integration:**<br>* View and manage Snyk Code integration.<br>* View and assign or remove Snyk Code projects to or from targets.<br>* View correlation data in the finding details.<br>* Provide correlation feedback.<br>* Register correlation matches (thumbs up or thumbs down). |
| **Name:** Correlation Viewer<br>**ID:** correlation\_viewer | **View SAST and DAST Integration:**<br>* View Snyk Code projects assigned to targets.<br>* View correlation data in the finding details.<br>* Provide correlation feedback.<br>* Register correlation matches (thumbs up or thumbs down).<br>* Disable and delete Snyk Code integration. |
| **Name:** Create Target<br>**ID:** create\_target | **Add Targets:**<br>* Add, change, list, view, and verify Domains.<br>* Add Targets (including uploading and downloading target files).<br>* View and list available Scanning Agents. |
| **Name:** Delete Target<br>**ID:** delete\_target | **Remove Targets:**<br>* Delete Targets.<br>* Delete Domains.<br>* Delete Webhooks. |
| **Name:** Discovery<br>**ID:** discovery | **Manage Discovery:**<br>* List, and view Discovery Assets.<br>* View Discovery Scans.<br>* View Discovery Asset Logs.<br>* Change Discovery Assets (Mark as new or not new, Hide or Show, Rename, Manage Target Labels, Add Notes). |
| **Name:** Discovery Read-Only<br>**ID:** discovery\_read\_only | **View Discovery Data:**<br>* List and view Discovery Assets.<br>* View Discovery Asset Logs.<br>* View Discovery Scans. |
| **Name:** Manage Credentials<br>**ID:** manage\_credentials | This permission allows you to create, view, update, and delete **credentials created by other users**, depending on your assigned scope:<br>* **Team Scope:** You can only manage credentials within the **specific teams** where you have been granted "Manage Credentials" permissions.<br>* **Account Scope:** Additionally, you can assign a credential to **any team** within the entire account. |
| **Name:** Password Login Override<br>**ID:** password\_login\_override | **Authentication:**<br>* Override SSO configuration (that is, log in with username and password when SSO is configured). |
| **Name:** Role Assignment<br>**ID:** role\_assignment | **Manage User Roles:**<br>* List high-level permissions.<br>* List built-in roles.<br>* Add, change, delete, list, and view custom-roles.<br>* View User details.<br>* Add, change, delete, list and view User Labels. |
| **Name:** Scanning Agent Management<br>**ID:** scanning\_agent\_management | **Manage Scanning Agents:**<br>* Add, change, delete, list, and view. |
| **Name:** Start re-test<br>**ID:** start\_retest | **Trigger Re-tests:**<br>* Initiate a re-test for a finding. |
| **Name:** Start Scan<br>**ID:** start\_scan | **Manage Scans:**<br>* Initiate a manual Scan on a Target.<br>* Cancel, pause, and resume ongoing Scans.<br>* Add, change, delete, list, and view Scheduled Scans. |
| **Name:** Team Management<br>**ID:** team\_management | **Manage Teams:**<br>* Add, change, delete, and view teams.<br>* Move targets between teams. |
| **Name:** User Management<br>**ID:** user\_management | **Manage Users and Roles (Account Level):**<br>* View User Roles, High-Level Permissions, and available Roles.<br>* Add, change, list, enable or disable Users and API Keys.<br>* Manage Users and API Keys' roles (list, view, assign, and remove role from User or API Key).<br>* Manage User Labels. |
| **Name:** View Target<br>**ID:** view\_target | **View Target Data and Reports:**<br>* List and view Target, Target Settings (including integration-specific settings), Scans, Scheduled Scans, and Findings.<br>* View integration-specific links.<br>* View Target Labels.<br>* View User details.<br><br>**Manage Reports:**<br>* Download Scan Reports.<br>* Add, change, delete, list, and view Stored or Managed Reports. |
