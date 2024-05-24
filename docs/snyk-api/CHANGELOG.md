## 2024-05-23

### POST - /groups/{group_id}/apps/installs/{install_id}/secrets - Updated
- added the optional property `data/attributes/installed_at` to the response with the `200` status



### GET - /groups/{group_id}/settings/pull_request_template - Added
- Get your groups pull request template


### POST - /groups/{group_id}/settings/pull_request_template - Added
- Configures a group level pull request template that will be used on any org or project within that group


### GET - /orgs/{org_id}/apps/installs - Updated
- added the optional property `data/items/attributes/installed_at` to the response with the `200` status



### POST - /orgs/{org_id}/apps/installs/{install_id}/secrets - Updated
- added the optional property `data/attributes/installed_at` to the response with the `200` status



### GET - /openapi/{version} - Updated
- removed the optional property `errors/items/links` from the response with the `400` status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `401` status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `404` status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `500` status
![Badge](https://img.shields.io/badge/Breaking-yellow)


### GET - /groups/{group_id}/apps/installs - Updated
- added the optional property `data/items/attributes/installed_at` to the response with the `200` status



### DELETE - /groups/{group_id}/apps/installs/{install_id} - Updated
- api operation id `deleteGroupAppInstallByID` removed and replaced with `deleteGroupAppInstallById`



### DELETE - /groups/{group_id}/settings/pull_request_template - Added
- Delete your groups pull request template. This means Snyk pull requests will start to use the default template for this group.


### DELETE - /orgs/{org_id}/apps/installs/{install_id} - Updated
- api operation id `deleteAppOrgInstallByID` removed and replaced with `deleteAppOrgInstallById`



### GET - /self/apps/installs - Updated
- added the optional property `data/items/attributes/installed_at` to the response with the `200` status



### DELETE - /self/apps/installs/{install_id} - Updated
- api operation id `deleteUserAppInstallByID` removed and replaced with `deleteUserAppInstallById`



### GET - /openapi - Updated
- removed the optional property `errors/items/links` from the response with the `400` status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `401` status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `404` status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `500` status
![Badge](https://img.shields.io/badge/Breaking-yellow)

# Change Log

## 2024-04-22
## 2024-04-11

### GET - /self - Updated
- endpoint added

### GET - /orgs/{org_id}/projects - Updated
- added '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' to the 'data/items/relationships/target' response property 'oneOf' list for the response status '200'

### GET - /orgs/{org_id}/projects/{project_id} - Updated
- added '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' to the 'data/relationships/target' response property 'oneOf' list for the response status '200'

### PATCH - /orgs/{org_id}/projects/{project_id} - Updated
- added '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' to the 'data/relationships/target' response property 'oneOf' list for the response status '200'
## 2024-03-15
## 2024-03-12
## 2024-02-28
## 2024-02-21

### GET - /orgs - Updated
- for the 'query' request parameter 'name', the maxLength was set to '100'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- for the 'query' request parameter 'slug', the maxLength was set to '100'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- added the pattern '^[\w.-]+$' to the 'query' request parameter 'slug'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- added the optional property 'data/items/attributes/access_requests_enabled' to the response with the '200' status



### GET - /groups/{group_id}/orgs - Updated
- endpoint added

### GET - /orgs/{org_id} - Updated
- added the optional property 'data/attributes/access_requests_enabled' to the response with the '200' status

### PATCH - /orgs/{org_id} - Updated
- added the optional property 'data/attributes/access_requests_enabled' to the response with the '200' status
## 2024-01-23

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

### GET - /orgs/{org_id}/projects/{project_id} - Updated
- removed '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' from the 'data/relationships/target' response property 'oneOf' list for the response status '200'

### PATCH - /orgs/{org_id}/projects/{project_id} - Updated
- removed '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' from the 'data/relationships/target' response property 'oneOf' list for the response status '200'

### GET - /orgs/{org_id}/targets - Updated
- endpoint added

### DELETE - /orgs/{org_id}/targets/{target_id} - Updated
- endpoint added

### GET - /orgs/{org_id}/targets/{target_id} - Updated
- endpoint added
## 2024-01-04

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
### GET - /groups/{group_id}/issues - Updated
- endpoint added

### GET - /groups/{group_id}/issues/{issue_id} - Updated
- endpoint added

### GET - /orgs/{org_id}/issues - Updated
- endpoint added

### GET - /orgs/{org_id}/issues/{issue_id} - Updated
- endpoint added
## 2023-12-21

### POST - /custom_base_images - Updated
- removed '#/components/schemas/VersioningSchemaDateType' from the 'data/attributes/versioning_schema' request property 'oneOf' list
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed '#/components/schemas/VersioningSchemaDateType' from the 'data/attributes/versioning_schema' response property 'oneOf' list for the response status '201'

### PATCH - /custom_base_images/{custombaseimage_id} - Updated
- removed '#/components/schemas/VersioningSchemaDateType' from the 'data/attributes/versioning_schema' request property 'oneOf' list
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed '#/components/schemas/VersioningSchemaDateType' from the 'data/attributes/versioning_schema' response property 'oneOf' list for the response status '200'

### GET - /custom_base_images/{custombaseimage_id} - Updated
- removed '#/components/schemas/VersioningSchemaDateType' from the 'data/attributes/versioning_schema' response property 'oneOf' list for the response status '200'
## 2023-12-19
## 2023-12-14
## 2023-12-01
## 2023-11-27
## 2023-11-06
## 2023-11-03

### GET - /openapi/{version} - Updated
- added the optional property 'errors/items/links' to the response with the '400' status

- added the optional property 'errors/items/links' to the response with the '401' status

- added the optional property 'errors/items/links' to the response with the '404' status



- added the optional property 'errors/items/links' to the response with the '500' status

### DELETE - /orgs/{org_id}/projects/{project_id} - Updated
- endpoint added

### GET - /openapi - Updated
- added the optional property 'errors/items/links' to the response with the '400' status

- added the optional property 'errors/items/links' to the response with the '401' status

- added the optional property 'errors/items/links' to the response with the '404' status



- added the optional property 'errors/items/links' to the response with the '500' status
## 2023-11-02

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
### DELETE - /groups/{group_id}/apps/installs/{install_id} - Updated
- endpoint added

### GET - /orgs/{org_id}/apps/creations - Updated
- endpoint added

### DELETE - /self/apps/{app_id}/sessions/{session_id} - Updated
- endpoint added

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
### DELETE - /orgs/{org_id}/apps/creations/{app_id} - Updated
- endpoint added

### POST - /orgs/{org_id}/apps/creations/{app_id}/secrets - Updated
- endpoint added

### POST - /orgs/{org_id}/apps/installs/{install_id}/secrets - Updated
- endpoint added

### POST - /groups/{group_id}/apps/installs/{install_id}/secrets - Updated
- endpoint added

### GET - /orgs/{org_id}/apps/creations/{app_id} - Updated
- endpoint added

### GET - /orgs/{org_id}/apps/installs - Updated
- endpoint added

### POST - /orgs/{org_id}/apps/installs - Updated
- endpoint added

### GET - /orgs/{org_id}/apps - Updated
- the response property 'data/items/attributes/client_id' became optional for the status '200'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- the response property 'data/items/attributes/redirect_uris' became optional for the status '200'
![Badge](https://img.shields.io/badge/Breaking-yellow)
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
### GET - /orgs/{org_id}/apps/{client_id} - Updated
- the response property 'data/attributes/client_id' became optional for the status '200'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- the response property 'data/attributes/redirect_uris' became optional for the status '200'
![Badge](https://img.shields.io/badge/Breaking-yellow)
### POST - /groups/{group_id}/apps/installs - Updated
- endpoint added

### DELETE - /self/apps/installs/{install_id} - Updated
- endpoint added

### GET - /self/apps/{app_id}/sessions - Updated
- endpoint added

### GET - /self/apps/installs - Updated
- endpoint added

### GET - /groups/{group_id}/apps/installs - Updated
- endpoint added

### POST - /orgs/{org_id}/apps/creations - Updated
- endpoint added

### PATCH - /orgs/{org_id}/apps/creations/{app_id} - Updated
- endpoint added

### DELETE - /orgs/{org_id}/apps/installs/{install_id} - Updated
- endpoint added
## 2023-10-24

### GET - /orgs/{org_id}/container_images - Updated
- endpoint added

### GET - /orgs/{org_id}/container_images/{image_id} - Updated
- endpoint added

### GET - /orgs/{org_id}/container_images/{image_id}/relationships/image_target_refs - Updated
- endpoint added
## 2023-09-20
## 2023-09-14
## 2023-09-13
## 2023-09-12

### GET - /orgs/{org_id}/projects - Updated
- added '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' to the 'data/items/relationships/target' response property 'oneOf' list for the response status '200'

### GET - /orgs/{org_id}/projects/{project_id} - Updated
- added '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' to the 'data/relationships/target' response property 'oneOf' list for the response status '200'

### PATCH - /orgs/{org_id}/projects/{project_id} - Updated
- added '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' to the 'data/relationships/target' response property 'oneOf' list for the response status '200'
## 2023-09-11

### GET - /orgs/{org_id}/collections - Updated
- endpoint added

### PATCH - /orgs/{org_id}/collections/{collection_id} - Updated
- endpoint added

### GET - /orgs/{org_id}/projects/{project_id} - Updated
- removed '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' from the 'data/relationships/target' response property 'oneOf' list for the response status '200'

### GET - /orgs/{org_id}/collections/{collection_id}/relationships/projects - Updated
- endpoint added

### POST - /orgs/{org_id}/collections/{collection_id}/relationships/projects - Updated
- endpoint added

### GET - /orgs/{org_id}/projects - Updated
- removed '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' from the 'data/items/relationships/target' response property 'oneOf' list for the response status '200'

### PATCH - /orgs/{org_id}/projects/{project_id} - Updated
- removed '#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget' from the 'data/relationships/target' response property 'oneOf' list for the response status '200'

### POST - /orgs/{org_id}/collections - Updated
- endpoint added

### DELETE - /orgs/{org_id}/collections/{collection_id} - Updated
- endpoint added

### GET - /orgs/{org_id}/collections/{collection_id} - Updated
- endpoint added

### DELETE - /orgs/{org_id}/collections/{collection_id}/relationships/projects - Updated
- endpoint added
## 2023-09-07

### GET - /orgs/{org_id}/audit_logs/search - Updated
- endpoint added

### PATCH - /orgs/{org_id}/settings/sast - Updated
- endpoint added

### GET - /groups/{group_id}/audit_logs/search - Updated
- endpoint added
## 2023-08-29

### POST - /orgs/{org_id}/service_accounts - Updated
- endpoint added

### PATCH - /orgs/{org_id}/service_accounts/{serviceaccount_id} - Updated
- endpoint added

### GET - /groups/{group_id}/service_accounts - Updated
- endpoint added

### DELETE - /groups/{group_id}/service_accounts/{serviceaccount_id} - Updated
- endpoint added

### GET - /groups/{group_id}/service_accounts/{serviceaccount_id} - Updated
- endpoint added

### PATCH - /groups/{group_id}/service_accounts/{serviceaccount_id} - Updated
- endpoint added

### POST - /groups/{group_id}/service_accounts/{serviceaccount_id}/secrets - Updated
- endpoint added

### GET - /orgs/{org_id}/service_accounts - Updated
- endpoint added

### DELETE - /orgs/{org_id}/service_accounts/{serviceaccount_id} - Updated
- endpoint added

### GET - /orgs/{org_id}/service_accounts/{serviceaccount_id} - Updated
- endpoint added

### POST - /groups/{group_id}/service_accounts - Updated
- endpoint added

### POST - /orgs/{org_id}/service_accounts/{serviceaccount_id}/secrets - Updated
- endpoint added
## 2023-08-28
## 2023-08-24

### GET - /orgs/{org_id}/projects - Updated
- added the new optional 'query' request parameter 'names_start_with'

- added the new optional 'query' request parameter 'target_file'



- added the new optional 'query' request parameter 'target_reference'

- added the new optional 'query' request parameter 'target_runtime'

### PATCH - /orgs/{org_id}/projects/{project_id} - Updated
- deleted the 'query' request parameter 'user_id'
![Badge](https://img.shields.io/badge/Breaking-yellow)## 2023-08-21
## 2023-08-04

### GET - /custom_base_images - Updated
- endpoint added

### POST - /custom_base_images - Updated
- endpoint added

### DELETE - /custom_base_images/{custombaseimage_id} - Updated
- endpoint added

### GET - /custom_base_images/{custombaseimage_id} - Updated
- endpoint added

### PATCH - /custom_base_images/{custombaseimage_id} - Updated
- endpoint added

### POST - /orgs/{org_id}/packages/issues - Updated
- added the optional property 'meta' to the response with the '200' status
## 2023-06-23
## 2023-06-22
## 2023-06-19

### GET - /orgs/{org_id}/settings/sast - Updated
- endpoint added

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
![Badge](https://img.shields.io/badge/Breaking-yellow)## 2023-04-28

### GET - /orgs - Updated
- endpoint added

### GET - /orgs/{org_id} - Updated
- endpoint added

### PATCH - /orgs/{org_id} - Updated
- endpoint added
## 2023-04-27

### GET - /orgs/{org_id}/invites - Updated
- the 'data/items/attributes/role' response's property type/format changed from 'string'/'' to 'string'/'uuid' for status '200'
![Badge](https://img.shields.io/badge/Breaking-yellow)
- added the new 'org_invitation' enum value to the 'data/items/type' response property for the response status '200'
![Badge](https://img.shields.io/badge/Breaking-yellow)


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
![Badge](https://img.shields.io/badge/Breaking-yellow)## 2023-04-17
## 2023-03-30

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

### POST - /orgs/{org_id}/packages/issues - Updated
- endpoint added
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
![Badge](https://img.shields.io/badge/Breaking-yellow)## 2023-03-20
## 2023-03-08

### GET - /orgs/{org_id}/projects/{project_id}/sbom - Updated
- endpoint added


## 2023-03-03
## 2023-02-15
## 2023-01-30

### GET - /orgs/{org_id}/projects - Updated
- endpoint added

### GET - /orgs/{org_id}/projects/{project_id} - Updated
- endpoint added

### PATCH - /orgs/{org_id}/projects/{project_id} - Updated
- endpoint added
## 2023-01-04
## 2022-12-21
## 2022-12-15
## 2022-12-14
## 2022-11-14

### DELETE - /orgs/{org_id}/slack_app/{bot_id}/projects/{project_id} - Updated
- endpoint added

### PATCH - /orgs/{org_id}/slack_app/{bot_id}/projects/{project_id} - Updated
- endpoint added

### POST - /orgs/{org_id}/slack_app/{bot_id}/projects/{project_id} - Updated
- endpoint added

### DELETE - /orgs/{org_id}/slack_app/{bot_id} - Updated
- endpoint added

### GET - /orgs/{org_id}/slack_app/{bot_id} - Updated
- endpoint added

### POST - /orgs/{org_id}/slack_app/{bot_id} - Updated
- endpoint added

### GET - /orgs/{org_id}/slack_app/{bot_id}/projects - Updated
- endpoint added
## 2022-11-07

### GET - /orgs/{org_id}/invites - Updated
- endpoint added

### DELETE - /orgs/{org_id}/invites/{invite_id} - Updated
- endpoint added
## 2022-10-06

### GET - /orgs/{org_id}/slack_app/{tenant_id}/channels - Updated
- endpoint added

### GET - /orgs/{org_id}/slack_app/{tenant_id}/channels/{channel_id} - Updated
- endpoint added
## 2022-09-15
## 2022-09-14

### GET - /orgs/{org_id}/packages/{purl}/issues - Updated
- endpoint added


## 2022-08-12
## 2022-06-08
## 2022-06-01
## 2022-04-06

### POST - /orgs/{org_id}/invites - Updated
- endpoint added


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

### GET - /openapi - Updated
- removed the optional property 'errors/items/links' from the response with the '400' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '401' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '404' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '500' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
### POST - /orgs/{org_id}/apps - Updated
- endpoint added

### PATCH - /orgs/{org_id}/apps/{client_id} - Updated
- endpoint added

### GET - /self/apps - Updated
- endpoint added

### GET - /openapi/{version} - Updated
- removed the optional property 'errors/items/links' from the response with the '400' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '401' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '404' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property 'errors/items/links' from the response with the '500' status
![Badge](https://img.shields.io/badge/Breaking-yellow)
### GET - /orgs/{org_id}/app_bots - Updated
- endpoint added

### DELETE - /orgs/{org_id}/app_bots/{bot_id} - Updated
- endpoint added

### GET - /orgs/{org_id}/apps - Updated
- endpoint added

### DELETE - /orgs/{org_id}/apps/{client_id} - Updated
- endpoint added

### GET - /orgs/{org_id}/apps/{client_id} - Updated
- endpoint added

### POST - /orgs/{org_id}/apps/{client_id}/secrets - Updated
- endpoint added

### DELETE - /self/apps/{app_id} - Updated
- endpoint added
## 2022-02-16
## 2022-02-01
## 2022-01-31
## 2021-12-09
## 2021-09-29

### GET - /groups/{group_id}/settings/iac - Updated
- endpoint added

### PATCH - /groups/{group_id}/settings/iac - Updated
- endpoint added

### GET - /orgs/{org_id}/settings/iac - Updated
- endpoint added

### PATCH - /orgs/{org_id}/settings/iac - Updated
- endpoint added
## 2021-09-13
## 2021-08-20
## 2021-08-13
## 2021-06-13
## 2021-06-07
## 2021-06-04
## 2021-06-01
