# API Change Log

## 2024-04-11

### GET - /orgs/{org_id}/projects/{project_id} - Updated
- added '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' to the 'data/relationships/target' response property 'oneOf' list for the response status '200'



### PATCH - /orgs/{org_id}/projects/{project_id} - Updated
- added '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' to the 'data/relationships/target' response property 'oneOf' list for the response status '200'



### GET - /self - Added
- Retrieves information about the the user making the request.


### GET - /orgs/{org_id}/projects - Updated
- added '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' to the 'data/items/relationships/target' response property 'oneOf' list for the response status '200'


## 2024-02-21

### GET - /orgs - Updated
- for the 'query' request parameter 'name', the maxLength was set to '100'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- for the 'query' request parameter 'slug', the maxLength was set to '100'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- added the pattern '^[\w.-]+$' to the 'query' request parameter 'slug'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- added the optional property 'data/items/attributes/access_requests_enabled' to the response with the '200' status



### GET - /groups/{group_id}/orgs - Added
- Get a paginated list of all the organizations belonging to the group.
By default, this endpoint returns the organizations in alphabetical order of their name.


### GET - /orgs/{org_id} - Updated
- added the optional property 'data/attributes/access_requests_enabled' to the response with the '200' status



### PATCH - /orgs/{org_id} - Updated
- added the optional property 'data/attributes/access_requests_enabled' to the response with the '200' status


## 2024-01-23

### GET - /orgs/{org_id}/projects/{project_id} - Updated
- removed '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' from the 'data/relationships/target' response property 'oneOf' list for the response status '200'



### PATCH - /orgs/{org_id}/projects/{project_id} - Updated
- removed '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' from the 'data/relationships/target' response property 'oneOf' list for the response status '200'



### GET - /orgs/{org_id}/targets - Added
- Get a list of an organization's targets.


### DELETE - /orgs/{org_id}/targets/{target_id} - Added
- Delete the specified target.


### GET - /orgs/{org_id}/targets/{target_id} - Added
- Get a specified target for an organization.


### GET - /openapi - Updated
- added the optional property 'errors/items/links' to the response with the '400' status

- added the optional property 'errors/items/links' to the response with the '401' status

- added the optional property 'errors/items/links' to the response with the '404' status

- added the optional property 'errors/items/links' to the response with the '500' status



### GET - /openapi/{version} - Updated
- added the optional property 'errors/items/links' to the response with the '400' status

- added the optional property 'errors/items/links' to the response with the '401' status

- added the optional property 'errors/items/links' to the response with the '404' status

- added the optional property 'errors/items/links' to the response with the '500' status



### GET - /orgs/{org_id}/projects - Updated
- removed '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' from the 'data/items/relationships/target' response property 'oneOf' list for the response status '200'


## 2024-01-04

### GET - /openapi/{version} - Updated
- removed the optional property 'errors/items/links' from the response with the '400' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '401' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '404' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '500' status
![Badge](https://img.shields.io/badge/Breaking-yellow)


### GET - /groups/{group_id}/issues - Added
- Get a list of a group's issues.


### GET - /groups/{group_id}/issues/{issue_id} - Added
- Get an issue


### GET - /orgs/{org_id}/issues - Added
- Get a list of an organization's issues.


### GET - /orgs/{org_id}/issues/{issue_id} - Added
- Get an issue


### GET - /openapi - Updated
- removed the optional property 'errors/items/links' from the response with the '400' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '401' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '404' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '500' status
![Badge](https://img.shields.io/badge/Breaking-yellow)

## 2023-12-21

### GET - /custom_base_images/{custombaseimage_id} - Updated
- removed '#/components/schemas/VersioningSchemaDateType' from the 'data/attributes/versioning_schema' response property 'oneOf' list for the response status '200'



### POST - /custom_base_images - Updated
- removed '#/components/schemas/VersioningSchemaDateType' from the 'data/attributes/versioning_schema' request property 'oneOf' list
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed '#/components/schemas/VersioningSchemaDateType' from the 'data/attributes/versioning_schema' response property 'oneOf' list for the response status '201'



### PATCH - /custom_base_images/{custombaseimage_id} - Updated
- removed '#/components/schemas/VersioningSchemaDateType' from the 'data/attributes/versioning_schema' request property 'oneOf' list
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed '#/components/schemas/VersioningSchemaDateType' from the 'data/attributes/versioning_schema' response property 'oneOf' list for the response status '200'


## 2023-11-03

### GET - /openapi - Updated
- added the optional property 'errors/items/links' to the response with the '400' status

- added the optional property 'errors/items/links' to the response with the '401' status

- added the optional property 'errors/items/links' to the response with the '404' status

- added the optional property 'errors/items/links' to the response with the '500' status



### GET - /openapi/{version} - Updated
- added the optional property 'errors/items/links' to the response with the '400' status

- added the optional property 'errors/items/links' to the response with the '401' status

- added the optional property 'errors/items/links' to the response with the '404' status

- added the optional property 'errors/items/links' to the response with the '500' status



### DELETE - /orgs/{org_id}/projects/{project_id} - Added
- Delete one project in the organization by project ID.

## 2023-11-02

### GET - /orgs/{org_id}/apps/creations - Added
- Get a list of apps created by an organization.


### POST - /orgs/{org_id}/apps/creations - Added
- Create a new Snyk App for an organization.


### POST - /orgs/{org_id}/apps/installs - Added
- Install a Snyk Apps to this organization, the Snyk App must use unattended authentication eg client credentials.


### DELETE - /orgs/{org_id}/apps/installs/{install_id} - Added
- Revoke app authorization for an Snyk Organization with install ID.


### GET - /self/apps/{app_id}/sessions - Added
- Get a list of active OAuth sessions for the app.


### GET - /orgs/{org_id}/apps - Updated
- the response property 'data/items/attributes/client_id' became optional for the status '200'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- the response property 'data/items/attributes/redirect_uris' became optional for the status '200'
![Badge](https://img.shields.io/badge/Breaking-yellow)


### GET - /groups/{group_id}/apps/installs - Added
- Get a list of apps installed for a group.


### DELETE - /groups/{group_id}/apps/installs/{install_id} - Added
- Revoke app authorization for an Snyk Group with install ID.


### GET - /orgs/{org_id}/apps/{client_id} - Updated
- the response property 'data/attributes/client_id' became optional for the status '200'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- the response property 'data/attributes/redirect_uris' became optional for the status '200'
![Badge](https://img.shields.io/badge/Breaking-yellow)


### PATCH - /orgs/{org_id}/apps/{client_id} - Updated
- added the new required request property 'data'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- the response property 'data/attributes/client_id' became optional for the status '200'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- the response property 'data/attributes/redirect_uris' became optional for the status '200'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property 'access_token_ttl_seconds'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property 'name'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property 'redirect_uris'
![Badge](https://img.shields.io/badge/Breaking-yellow)


### GET - /orgs/{org_id}/apps/creations/{app_id} - Added
- Get a Snyk App by its App ID.


### POST - /orgs/{org_id}/apps/creations/{app_id}/secrets - Added
- Manage client secret for the Snyk App.


### DELETE - /self/apps/installs/{install_id} - Added
- Revoke access for an app by install ID.


### DELETE - /self/apps/{app_id}/sessions/{session_id} - Added
- Revoke an active user app session.


### GET - /openapi/{version} - Updated
- removed the optional property 'errors/items/links' from the response with the '400' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '401' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '404' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '500' status
![Badge](https://img.shields.io/badge/Breaking-yellow)


### POST - /groups/{group_id}/apps/installs - Added
- Install a Snyk Apps to this group, the Snyk App must use unattended authentication eg client credentials.


### DELETE - /orgs/{org_id}/apps/creations/{app_id} - Added
- Delete an app by its App ID.


### PATCH - /orgs/{org_id}/apps/creations/{app_id} - Added
- Update app creation attributes with App ID.


### GET - /orgs/{org_id}/apps/installs - Added
- Get a list of apps installed for an organization.


### POST - /orgs/{org_id}/apps/installs/{install_id}/secrets - Added
- Manage client secret for non-interactive Snyk App installations.


### GET - /self/apps/installs - Added
- Get a list of apps installed for an user.


### POST - /orgs/{org_id}/apps - Updated
- added the new required request property 'data'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property 'access_token_ttl_seconds'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property 'context'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property 'name'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property 'redirect_uris'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property 'scopes'
![Badge](https://img.shields.io/badge/Breaking-yellow)


### GET - /openapi - Updated
- removed the optional property 'errors/items/links' from the response with the '400' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '401' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '404' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '500' status
![Badge](https://img.shields.io/badge/Breaking-yellow)


### POST - /groups/{group_id}/apps/installs/{install_id}/secrets - Added
- Manage client secret for non-interactive Snyk App installations.

## 2023-10-24

### GET - /orgs/{org_id}/container_images - Added
- List instances of container image


### GET - /orgs/{org_id}/container_images/{image_id} - Added
- Get instance of container image


### GET - /orgs/{org_id}/container_images/{image_id}/relationships/image_target_refs - Added
- List instances of image target references for a container image

## 2023-09-12

### GET - /orgs/{org_id}/projects - Updated
- added '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' to the 'data/items/relationships/target' response property 'oneOf' list for the response status '200'



### GET - /orgs/{org_id}/projects/{project_id} - Updated
- added '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' to the 'data/relationships/target' response property 'oneOf' list for the response status '200'



### PATCH - /orgs/{org_id}/projects/{project_id} - Updated
- added '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' to the 'data/relationships/target' response property 'oneOf' list for the response status '200'


## 2023-09-11

### POST - /orgs/{org_id}/collections - Added
- Create a collection


### DELETE - /orgs/{org_id}/collections/{collection_id} - Added
- Delete a collection


### GET - /orgs/{org_id}/projects - Updated
- removed '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' from the 'data/items/relationships/target' response property 'oneOf' list for the response status '200'



### GET - /orgs/{org_id}/collections/{collection_id}/relationships/projects - Added
- Return a list of organization's projects that are from the specified collection.


### POST - /orgs/{org_id}/collections/{collection_id}/relationships/projects - Added
- Add projects to a collection by specifying an array of project ids


### GET - /orgs/{org_id}/projects/{project_id} - Updated
- removed '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' from the 'data/relationships/target' response property 'oneOf' list for the response status '200'



### PATCH - /orgs/{org_id}/projects/{project_id} - Updated
- removed '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' from the 'data/relationships/target' response property 'oneOf' list for the response status '200'



### GET - /orgs/{org_id}/collections - Added
- Return a list of organization's collections with issues counts  and projects count.


### GET - /orgs/{org_id}/collections/{collection_id} - Added
- Get a collection


### PATCH - /orgs/{org_id}/collections/{collection_id} - Added
- Edit a collection


### DELETE - /orgs/{org_id}/collections/{collection_id}/relationships/projects - Added
- Remove projects from a collection by specifying an array of project ids

## 2023-09-07

### GET - /groups/{group_id}/audit_logs/search - Added
- Search audit logs for a Group. Some Organization level events are supported as well as the following
Group level events:
  - api.access
  - group.cloud_config.settings.edit
  - group.create
  - group.delete
  - group.edit
  - group.notification_settings.edit
  - group.org.add
  - group.org.remove
  - group.policy.create
  - group.policy.delete
  - group.policy.edit
  - group.request_access_settings.edit
  - group.role.create
  - group.role.delete
  - group.role.edit
  - group.service_account.create
  - group.service_account.delete
  - group.service_account.edit
  - group.settings.edit
  - group.settings.feature_flag.edit
  - group.sso.add
  - group.sso.auth0_connection.create
  - group.sso.auth0_connection.edit
  - group.sso.create
  - group.sso.delete
  - group.sso.edit
  - group.sso.membership.sync
  - group.sso.remove
  - group.tag.create
  - group.tag.delete
  - group.user.add
  - group.user.remove
  - group.user.role.edit



### GET - /orgs/{org_id}/audit_logs/search - Added
- Search audit logs for an Organization. Supported event types:
  - api.access
  - org.app_bot.create
  - org.app.create
  - org.app.delete
  - org.app.edit
  - org.cloud_config.settings.edit
  - org.collection.create
  - org.collection.delete
  - org.collection.edit
  - org.create
  - org.delete
  - org.edit
  - org.ignore_policy.edit
  - org.integration.create
  - org.integration.delete
  - org.integration.edit
  - org.integration.settings.edit
  - org.language_settings.edit
  - org.notification_settings.edit
  - org.org_source.create
  - org.org_source.delete
  - org.org_source.edit
  - org.policy.edit
  - org.project_filter.create
  - org.project_filter.delete
  - org.project.add
  - org.project.attributes.edit
  - org.project.delete
  - org.project.edit
  - org.project.fix_pr.auto_open
  - org.project.fix_pr.manual_open
  - org.project.ignore.create
  - org.project.ignore.delete
  - org.project.ignore.edit
  - org.project.monitor
  - org.project.pr_check.edit
  - org.project.remove
  - org.project.settings.delete
  - org.project.settings.edit
  - org.project.stop_monitor
  - org.project.tag.add
  - org.project.tag.remove
  - org.project.test
  - org.request_access_settings.edit
  - org.sast_settings.edit
  - org.service_account.create
  - org.service_account.delete
  - org.service_account.edit
  - org.settings.feature_flag.edit
  - org.target.create
  - org.target.delete
  - org.user.add
  - org.user.invite
  - org.user.invite.accept
  - org.user.invite.revoke
  - org.user.invite_link.accept
  - org.user.invite_link.create
  - org.user.invite_link.revoke
  - org.user.leave
  - org.user.provision.accept
  - org.user.provision.create
  - org.user.provision.delete
  - org.user.remove
  - org.user.role.create
  - org.user.role.delete
  - org.user.role.details.edit
  - org.user.role.edit
  - org.user.role.permissions.edit
  - org.webhook.add
  - org.webhook.delete
  - user.org.notification_settings.edit



### PATCH - /orgs/{org_id}/settings/sast - Added
- Enable/Disable the Snyk Code settings for an org

## 2023-08-29

### POST - /groups/{group_id}/service_accounts/{serviceaccount_id}/secrets - Added
- Manage the client secret of a group service account by the service account ID.


### GET - /orgs/{org_id}/service_accounts - Added
- Get all service accounts for an organization.


### POST - /orgs/{org_id}/service_accounts - Added
- Create a service account for an organization. The service account can be used to access the Snyk API.


### PATCH - /orgs/{org_id}/service_accounts/{serviceaccount_id} - Added
- Update the name of an organization-level service account by its ID.


### POST - /groups/{group_id}/service_accounts - Added
- Create a service account for a group. The service account can be used to access the Snyk API.


### DELETE - /groups/{group_id}/service_accounts/{serviceaccount_id} - Added
- Permanently delete a group-level service account by its ID.


### GET - /groups/{group_id}/service_accounts/{serviceaccount_id} - Added
- Get a group-level service account by its ID.


### GET - /orgs/{org_id}/service_accounts/{serviceaccount_id} - Added
- Get an organization-level service account by its ID.


### POST - /orgs/{org_id}/service_accounts/{serviceaccount_id}/secrets - Added
- Manage the client secret of an organization service account by the service account ID.


### GET - /groups/{group_id}/service_accounts - Added
- Get all service accounts for a group.


### PATCH - /groups/{group_id}/service_accounts/{serviceaccount_id} - Added
- Update the name of a group's service account by its ID.


### DELETE - /orgs/{org_id}/service_accounts/{serviceaccount_id} - Added
- Delete a service account in an organization.

## 2023-08-24

### PATCH - /orgs/{org_id}/projects/{project_id} - Updated
- deleted the 'query' request parameter 'user_id'
![Badge](https://img.shields.io/badge/Breaking-yellow)


### GET - /orgs/{org_id}/projects - Updated
- added the new optional 'query' request parameter 'names_start_with'

- added the new optional 'query' request parameter 'target_file'

- added the new optional 'query' request parameter 'target_reference'

- added the new optional 'query' request parameter 'target_runtime'


## 2023-08-04

### GET - /custom_base_images - Added
- Get a list of custom base images with support for ordering and filtering.
Either the org_id or group_id parameters must be set to authorize successfully.



### POST - /custom_base_images - Added
- In order to create a custom base image, you first need to import your base images into Snyk.
You can do this through the CLI, UI, or API.

This endpoint marks an image as a custom base image. This means that the image will get
added to the pool of images from which Snyk can recommend base image upgrades.

Note, after the first image in a repository gets added, a versioning schema cannot be passed in this endpoint.
To update the versioning schema, the PATCH endpoint must be used.



### DELETE - /custom_base_images/{custombaseimage_id} - Added
- Delete a custom base image resource. (the related container project is unaffected)


### GET - /custom_base_images/{custombaseimage_id} - Added
- Get a custom base image


### PATCH - /custom_base_images/{custombaseimage_id} - Added
- Updates a custom base image's attributes


### POST - /orgs/{org_id}/packages/issues - Updated
- added the optional property 'meta' to the response with the '200' status


## 2023-06-19

### GET - /orgs/{org_id}/settings/sast - Added
- Retrieves the SAST settings for an org


### GET - /openapi - Updated
- added the optional property 'errors/items/links' to the response with the '400' status

- added the optional property 'errors/items/links' to the response with the '401' status

- added the optional property 'errors/items/links' to the response with the '404' status

- added the optional property 'errors/items/links' to the response with the '500' status



### GET - /openapi/{version} - Updated
- added the optional property 'errors/items/links' to the response with the '400' status

- added the optional property 'errors/items/links' to the response with the '401' status

- added the optional property 'errors/items/links' to the response with the '404' status

- added the optional property 'errors/items/links' to the response with the '500' status


## 2023-05-29

### GET - /openapi - Updated
- removed the optional property 'errors/items/links' from the response with the '400' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '401' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '404' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '500' status
![Badge](https://img.shields.io/badge/Breaking-yellow)


### GET - /openapi/{version} - Updated
- removed the optional property 'errors/items/links' from the response with the '400' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '401' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '404' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '500' status
![Badge](https://img.shields.io/badge/Breaking-yellow)

## 2023-04-28

### GET - /orgs - Added
- Get a paginated list of organizations you have access to.


### GET - /orgs/{org_id} - Added
- Get the full details of an organization.


### PATCH - /orgs/{org_id} - Added
- Update the details of an organization

## 2023-04-27

### POST - /orgs/{org_id}/invites - Updated
- added the new required request property 'data'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- the 'data/attributes/role' response's property type/format changed from 'string'/'' to 'string'/'uuid' for status '201'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property 'email'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property 'role'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- added the new 'org_invitation' enum value to the 'data/type' response property for the response status '201'
![Badge](https://img.shields.io/badge/Breaking-yellow)


### GET - /orgs/{org_id}/invites - Updated
- the 'data/items/attributes/role' response's property type/format changed from 'string'/'' to 'string'/'uuid' for status '200'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- added the new 'org_invitation' enum value to the 'data/items/type' response property for the response status '200'
![Badge](https://img.shields.io/badge/Breaking-yellow)

## 2023-03-30

### GET - /openapi/{version} - Updated
- added the optional property 'errors/items/links' to the response with the '400' status

- added the optional property 'errors/items/links' to the response with the '401' status

- added the optional property 'errors/items/links' to the response with the '404' status

- added the optional property 'errors/items/links' to the response with the '500' status



### POST - /orgs/{org_id}/packages/issues - Added
- This endpoint is not available to all customers. If you are interested please contact support. Query issues for a batch of packages identified by Package URL (purl). Only direct vulnerabilities are returned, transitive vulnerabilities (from dependencies) are not returned because they can vary depending on context.


### GET - /openapi - Updated
- added the optional property 'errors/items/links' to the response with the '400' status

- added the optional property 'errors/items/links' to the response with the '401' status

- added the optional property 'errors/items/links' to the response with the '404' status

- added the optional property 'errors/items/links' to the response with the '500' status


## 2023-03-29

### GET - /openapi - Updated
- removed the optional property 'errors/items/links' from the response with the '400' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '401' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '404' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '500' status
![Badge](https://img.shields.io/badge/Breaking-yellow)


### GET - /openapi/{version} - Updated
- removed the optional property 'errors/items/links' from the response with the '400' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '401' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '404' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '500' status
![Badge](https://img.shields.io/badge/Breaking-yellow)

## 2023-03-08

### GET - /orgs/{org_id}/projects/{project_id}/sbom - Added
- This endpoint lets you retrieve the SBOM document of a software project.
It supports the following formats:
* CycloneDX version 1.4 in JSON (set `format` to `cyclonedx1.4+json`).
* CycloneDX version 1.4 in XML (set `format` to `cyclonedx1.4+xml`).
* SPDX version 2.3 in JSON (set `format` to `spdx2.3+json`).

By default it will respond with an empty JSON:API response.

## 2023-01-30

### GET - /orgs/{org_id}/projects - Added
- List all Projects for an Org.


### GET - /orgs/{org_id}/projects/{project_id} - Added
- Get one project of the organization by project ID.


### PATCH - /orgs/{org_id}/projects/{project_id} - Added
- Updates one project of the organization by project ID.

## 2022-11-14

### POST - /orgs/{org_id}/slack_app/{bot_id}/projects/{project_id} - Added
- Create Slack settings override for a project.


### DELETE - /orgs/{org_id}/slack_app/{bot_id} - Added
- Remove the given Slack App integration


### GET - /orgs/{org_id}/slack_app/{bot_id} - Added
- Get Slack integration default notification settings for the provided tenant ID and bot ID.


### POST - /orgs/{org_id}/slack_app/{bot_id} - Added
- Create new Slack notification default settings for a given tenant.


### GET - /orgs/{org_id}/slack_app/{bot_id}/projects - Added
- Slack notification settings overrides for projects. These settings overrides the default settings configured for the tenant.


### DELETE - /orgs/{org_id}/slack_app/{bot_id}/projects/{project_id} - Added
- Remove Slack settings override for a project.


### PATCH - /orgs/{org_id}/slack_app/{bot_id}/projects/{project_id} - Added
- Update Slack notification settings for a project.

## 2022-11-07

### GET - /orgs/{org_id}/invites - Added
- List pending user invitations to an organization.


### DELETE - /orgs/{org_id}/invites/{invite_id} - Added
- Cancel a pending user invitations to an organization.

## 2022-10-06

### GET - /orgs/{org_id}/slack_app/{tenant_id}/channels - Added
- Requires the Snyk Slack App to be set up for this org, will retrieve a list of channels the Snyk Slack App can access. Note that it is currently only possible to page forwards through this collection, no prev links will be generated and the ending_before parameter will not function.


### GET - /orgs/{org_id}/slack_app/{tenant_id}/channels/{channel_id} - Added
- Requires the Snyk Slack App to be set up for this org. It will return the Slack channel name for the provided Slack channel ID.

## 2022-09-14

### GET - /orgs/{org_id}/packages/{purl}/issues - Added
- Query issues for a specific package version identified by Package URL (purl). Snyk returns only direct vulnerabilities. Transitive vulnerabilities (from dependencies) are not returned because they can vary depending on context.

## 2022-04-06

### POST - /orgs/{org_id}/invites - Added
- Invite a user to an organization with a role.

## 2022-03-11

### GET - /openapi - Updated
- added the optional property 'errors/items/links' to the response with the '400' status

- added the optional property 'errors/items/links' to the response with the '401' status

- added the optional property 'errors/items/links' to the response with the '404' status

- added the optional property 'errors/items/links' to the response with the '500' status



### GET - /openapi/{version} - Updated
- added the optional property 'errors/items/links' to the response with the '400' status

- added the optional property 'errors/items/links' to the response with the '401' status

- added the optional property 'errors/items/links' to the response with the '404' status

- added the optional property 'errors/items/links' to the response with the '500' status


## 2022-03-01

### DELETE - /orgs/{org_id}/app_bots/{bot_id} - Added
- Revoke app bot authorization. Deprecated, use /orgs/{org_id}/apps/installs/{install_id} instead.


### GET - /orgs/{org_id}/apps - Added
- Get a list of apps created by an organization. Deprecated, use /orgs/{org_id}/apps/creations instead.


### POST - /orgs/{org_id}/apps - Added
- Create a new app for an organization. Deprecated, use /orgs/{org_id}/apps/creations instead.


### GET - /orgs/{org_id}/apps/{client_id} - Added
- Get an App by client id. Deprecated, use /orgs/{org_id}/apps/creations/{app_id} instead.


### PATCH - /orgs/{org_id}/apps/{client_id} - Added
- Update app attributes. Deprecated, use /orgs/{org_id}/apps/creations/{app_id} instead.


### DELETE - /self/apps/{app_id} - Added
- Revoke access for an app by app id


### GET - /openapi - Updated
- removed the optional property 'errors/items/links' from the response with the '400' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '401' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '404' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '500' status
![Badge](https://img.shields.io/badge/Breaking-yellow)


### GET - /openapi/{version} - Updated
- removed the optional property 'errors/items/links' from the response with the '400' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '401' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '404' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '500' status
![Badge](https://img.shields.io/badge/Breaking-yellow)


### GET - /orgs/{org_id}/app_bots - Added
- Get a list of app bots authorized to an organization. Deprecated, use /orgs/{org_id}/apps/installs instead.


### DELETE - /orgs/{org_id}/apps/{client_id} - Added
- Delete an app by app id. Deprecated, use /orgs/{org_id}/apps/creations/{app_id} instead.


### POST - /orgs/{org_id}/apps/{client_id}/secrets - Added
- Manage client secrets for an app. Deprecated, use /orgs/{org_id}/apps/creations/{app_id}/secrets instead.


### GET - /self/apps - Added
- Get a list of apps that can act on your behalf.

## 2021-09-29

### GET - /groups/{group_id}/settings/iac - Added
- Get the Infrastructure as Code Settings for a group.


### PATCH - /groups/{group_id}/settings/iac - Added
- Update the Infrastructure as Code Settings for a group.


### GET - /orgs/{org_id}/settings/iac - Added
- Get the Infrastructure as Code Settings for an org.


### PATCH - /orgs/{org_id}/settings/iac - Added
- Update the Infrastructure as Code Settings for an org.

