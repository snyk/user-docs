## 2024-04-11


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects</span>
  - Updated
</div>

- added `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` to the `data/items/relationships/target` response property `oneOf` list for the response status `200`




<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects/{project_id}</span>
  - Updated
</div>

- added `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` to the `data/relationships/target` response property `oneOf` list for the response status `200`




<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects/{project_id}</span>
  - Updated
</div>

- added `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` to the `data/relationships/target` response property `oneOf` list for the response status `200`




<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/self</span>
  - Added
</div>

- Retrieves information about the the user making the request.

## 2024-02-21


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs</span>
  - Updated
</div>

- for the `query` request parameter `name`, the maxLength was set to `100`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- for the `query` request parameter `slug`, the maxLength was set to `100`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- added the pattern `^[\w.-]+$` to the `query` request parameter `slug`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- added the optional property `data/items/attributes/access_requests_enabled` to the response with the `200` status




<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/groups/{group_id}/orgs</span>
  - Added
</div>

- Get a paginated list of all the organizations belonging to the group.
  By default, this endpoint returns the organizations in alphabetical order of their name.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}</span>
  - Updated
</div>

- added the optional property `data/attributes/access_requests_enabled` to the response with the `200` status




<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}</span>
  - Updated
</div>

- added the optional property `data/attributes/access_requests_enabled` to the response with the `200` status


## 2024-01-23


<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects/{project_id}</span>
  - Updated
</div>

- removed `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` from the `data/relationships/target` response property `oneOf` list for the response status `200`




<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/targets</span>
  - Added
</div>

- Get a list of an organization`s targets.



<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/targets/{target_id}</span>
  - Added
</div>

- Delete the specified target.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/targets/{target_id}</span>
  - Added
</div>

- Get a specified target for an organization.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi</span>
  - Updated
</div>

- added the optional property `errors/items/links` to the response with the `400` status

- added the optional property `errors/items/links` to the response with the `401` status

- added the optional property `errors/items/links` to the response with the `404` status

- added the optional property `errors/items/links` to the response with the `500` status




<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi/{version}</span>
  - Updated
</div>

- added the optional property `errors/items/links` to the response with the `400` status

- added the optional property `errors/items/links` to the response with the `401` status

- added the optional property `errors/items/links` to the response with the `404` status

- added the optional property `errors/items/links` to the response with the `500` status




<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects</span>
  - Updated
</div>

- removed `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` from the `data/items/relationships/target` response property `oneOf` list for the response status `200`




<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects/{project_id}</span>
  - Updated
</div>

- removed `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` from the `data/relationships/target` response property `oneOf` list for the response status `200`


## 2024-01-04


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi/{version}</span>
  - Updated
</div>

- removed the optional property `errors/items/links` from the response with the `400` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `401` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `404` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `500` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/groups/{group_id}/issues</span>
  - Added
</div>

- Get a list of a group`s issues.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/groups/{group_id}/issues/{issue_id}</span>
  - Added
</div>

- Get an issue



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/issues</span>
  - Added
</div>

- Get a list of an organization`s issues.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/issues/{issue_id}</span>
  - Added
</div>

- Get an issue



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi</span>
  - Updated
</div>

- removed the optional property `errors/items/links` from the response with the `400` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `401` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `404` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `500` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)

## 2023-12-21


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/custom_base_images/{custombaseimage_id}</span>
  - Updated
</div>

- removed `#/components/schemas/VersioningSchemaDateType` from the `data/attributes/versioning_schema` response property `oneOf` list for the response status `200`




<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/custom_base_images</span>
  - Updated
</div>

- removed `#/components/schemas/VersioningSchemaDateType` from the `data/attributes/versioning_schema` request property `oneOf` list
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed `#/components/schemas/VersioningSchemaDateType` from the `data/attributes/versioning_schema` response property `oneOf` list for the response status `201`




<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/custom_base_images/{custombaseimage_id}</span>
  - Updated
</div>

- removed `#/components/schemas/VersioningSchemaDateType` from the `data/attributes/versioning_schema` request property `oneOf` list
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed `#/components/schemas/VersioningSchemaDateType` from the `data/attributes/versioning_schema` response property `oneOf` list for the response status `200`


## 2023-11-03


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi</span>
  - Updated
</div>

- added the optional property `errors/items/links` to the response with the `400` status

- added the optional property `errors/items/links` to the response with the `401` status

- added the optional property `errors/items/links` to the response with the `404` status

- added the optional property `errors/items/links` to the response with the `500` status




<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi/{version}</span>
  - Updated
</div>

- added the optional property `errors/items/links` to the response with the `400` status

- added the optional property `errors/items/links` to the response with the `401` status

- added the optional property `errors/items/links` to the response with the `404` status

- added the optional property `errors/items/links` to the response with the `500` status




<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects/{project_id}</span>
  - Added
</div>

- Delete one project in the organization by project ID.

## 2023-11-02


<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/groups/{group_id}/apps/installs/{install_id}/secrets</span>
  - Added
</div>

- Manage client secret for non-interactive Snyk App installations.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/self/apps/{app_id}/sessions</span>
  - Added
</div>

- Get a list of active OAuth sessions for the app.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps</span>
  - Updated
</div>

- the response property `data/items/attributes/client_id` became optional for the status `200`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- the response property `data/items/attributes/redirect_uris` became optional for the status `200`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)



<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps/{client_id}</span>
  - Updated
</div>

- added the new required request property `data`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- the response property `data/attributes/client_id` became optional for the status `200`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- the response property `data/attributes/redirect_uris` became optional for the status `200`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property `access_token_ttl_seconds`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property `name`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property `redirect_uris`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/groups/{group_id}/apps/installs</span>
  - Added
</div>

- Get a list of apps installed for a group.



<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/groups/{group_id}/apps/installs</span>
  - Added
</div>

- Install a Snyk Apps to this group, the Snyk App must use unattended authentication eg client credentials.



<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/groups/{group_id}/apps/installs/{install_id}</span>
  - Added
</div>

- Revoke app authorization for an Snyk Group with install ID.



<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps/creations/{app_id}</span>
  - Added
</div>

- Update app creation attributes with App ID.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps/installs</span>
  - Added
</div>

- Get a list of apps installed for an organization.



<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps/installs/{install_id}</span>
  - Added
</div>

- Revoke app authorization for an Snyk Organization with install ID.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps/{client_id}</span>
  - Updated
</div>

- the response property `data/attributes/client_id` became optional for the status `200`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- the response property `data/attributes/redirect_uris` became optional for the status `200`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi/{version}</span>
  - Updated
</div>

- removed the optional property `errors/items/links` from the response with the `400` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `401` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `404` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `500` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps/creations</span>
  - Added
</div>

- Get a list of apps created by an organization.



<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps/creations</span>
  - Added
</div>

- Create a new Snyk App for an organization.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps/creations/{app_id}</span>
  - Added
</div>

- Get a Snyk App by its App ID.



<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps/installs/{install_id}/secrets</span>
  - Added
</div>

- Manage client secret for non-interactive Snyk App installations.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/self/apps/installs</span>
  - Added
</div>

- Get a list of apps installed for an user.



<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps/creations/{app_id}/secrets</span>
  - Added
</div>

- Manage client secret for the Snyk App.



<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/self/apps/installs/{install_id}</span>
  - Added
</div>

- Revoke access for an app by install ID.



<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/self/apps/{app_id}/sessions/{session_id}</span>
  - Added
</div>

- Revoke an active user app session.



<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps</span>
  - Updated
</div>

- added the new required request property `data`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property `access_token_ttl_seconds`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property `context`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property `name`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property `redirect_uris`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property `scopes`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi</span>
  - Updated
</div>

- removed the optional property `errors/items/links` from the response with the `400` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `401` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `404` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `500` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)



<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps/creations/{app_id}</span>
  - Added
</div>

- Delete an app by its App ID.



<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps/installs</span>
  - Added
</div>

- Install a Snyk Apps to this organization, the Snyk App must use unattended authentication eg client credentials.

## 2023-10-24


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/container_images</span>
  - Added
</div>

- List instances of container image



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/container_images/{image_id}</span>
  - Added
</div>

- Get instance of container image



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/container_images/{image_id}/relationships/image_target_refs</span>
  - Added
</div>

- List instances of image target references for a container image

## 2023-09-12


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects/{project_id}</span>
  - Updated
</div>

- added `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` to the `data/relationships/target` response property `oneOf` list for the response status `200`




<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects/{project_id}</span>
  - Updated
</div>

- added `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` to the `data/relationships/target` response property `oneOf` list for the response status `200`




<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects</span>
  - Updated
</div>

- added `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` to the `data/items/relationships/target` response property `oneOf` list for the response status `200`


## 2023-09-11


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects/{project_id}</span>
  - Updated
</div>

- removed `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` from the `data/relationships/target` response property `oneOf` list for the response status `200`




<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects/{project_id}</span>
  - Updated
</div>

- removed `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` from the `data/relationships/target` response property `oneOf` list for the response status `200`




<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/collections</span>
  - Added
</div>

- Return a list of organization`s collections with issues counts  and projects count.



<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/collections/{collection_id}</span>
  - Added
</div>

- Delete a collection



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/collections/{collection_id}/relationships/projects</span>
  - Added
</div>

- Return a list of organization`s projects that are from the specified collection.



<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/collections/{collection_id}/relationships/projects</span>
  - Added
</div>

- Remove projects from a collection by specifying an array of project ids



<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/collections/{collection_id}/relationships/projects</span>
  - Added
</div>

- Add projects to a collection by specifying an array of project ids



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects</span>
  - Updated
</div>

- removed `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` from the `data/items/relationships/target` response property `oneOf` list for the response status `200`




<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/collections</span>
  - Added
</div>

- Create a collection



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/collections/{collection_id}</span>
  - Added
</div>

- Get a collection



<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/collections/{collection_id}</span>
  - Added
</div>

- Edit a collection

## 2023-09-07


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/groups/{group_id}/audit_logs/search</span>
  - Added
</div>

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




<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/audit_logs/search</span>
  - Added
</div>

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




<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/settings/sast</span>
  - Added
</div>

- Enable/Disable the Snyk Code settings for an org

## 2023-08-29


<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/service_accounts/{serviceaccount_id}/secrets</span>
  - Added
</div>

- Manage the client secret of an organization service account by the service account ID.



<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/groups/{group_id}/service_accounts/{serviceaccount_id}</span>
  - Added
</div>

- Update the name of a group`s service account by its ID.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/service_accounts</span>
  - Added
</div>

- Get all service accounts for an organization.



<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/service_accounts</span>
  - Added
</div>

- Create a service account for an organization. The service account can be used to access the Snyk API.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/service_accounts/{serviceaccount_id}</span>
  - Added
</div>

- Get an organization-level service account by its ID.



<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/service_accounts/{serviceaccount_id}</span>
  - Added
</div>

- Update the name of an organization-level service account by its ID.



<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/service_accounts/{serviceaccount_id}</span>
  - Added
</div>

- Delete a service account in an organization.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/groups/{group_id}/service_accounts</span>
  - Added
</div>

- Get all service accounts for a group.



<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/groups/{group_id}/service_accounts</span>
  - Added
</div>

- Create a service account for a group. The service account can be used to access the Snyk API.



<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/groups/{group_id}/service_accounts/{serviceaccount_id}</span>
  - Added
</div>

- Permanently delete a group-level service account by its ID.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/groups/{group_id}/service_accounts/{serviceaccount_id}</span>
  - Added
</div>

- Get a group-level service account by its ID.



<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/groups/{group_id}/service_accounts/{serviceaccount_id}/secrets</span>
  - Added
</div>

- Manage the client secret of a group service account by the service account ID.

## 2023-08-24


<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects/{project_id}</span>
  - Updated
</div>

- deleted the `query` request parameter `user_id`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects</span>
  - Updated
</div>

- added the new optional `query` request parameter `names_start_with`

- added the new optional `query` request parameter `target_file`

- added the new optional `query` request parameter `target_reference`

- added the new optional `query` request parameter `target_runtime`


## 2023-08-04


<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/custom_base_images/{custombaseimage_id}</span>
  - Added
</div>

- Updates a custom base image`s attributes



<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/packages/issues</span>
  - Updated
</div>

- added the optional property `meta` to the response with the `200` status




<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/custom_base_images</span>
  - Added
</div>

- Get a list of custom base images with support for ordering and filtering.
  Either the org_id or group_id parameters must be set to authorize successfully.




<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/custom_base_images</span>
  - Added
</div>

- In order to create a custom base image, you first need to import your base images into Snyk.
  You can do this through the CLI, UI, or API.

This endpoint marks an image as a custom base image. This means that the image will get
added to the pool of images from which Snyk can recommend base image upgrades.

Note, after the first image in a repository gets added, a versioning schema cannot be passed in this endpoint.
To update the versioning schema, the PATCH endpoint must be used.




<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/custom_base_images/{custombaseimage_id}</span>
  - Added
</div>

- Delete a custom base image resource. (the related container project is unaffected)



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/custom_base_images/{custombaseimage_id}</span>
  - Added
</div>

- Get a custom base image

## 2023-06-19


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi</span>
  - Updated
</div>

- added the optional property `errors/items/links` to the response with the `400` status

- added the optional property `errors/items/links` to the response with the `401` status

- added the optional property `errors/items/links` to the response with the `404` status

- added the optional property `errors/items/links` to the response with the `500` status




<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi/{version}</span>
  - Updated
</div>

- added the optional property `errors/items/links` to the response with the `400` status

- added the optional property `errors/items/links` to the response with the `401` status

- added the optional property `errors/items/links` to the response with the `404` status

- added the optional property `errors/items/links` to the response with the `500` status




<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/settings/sast</span>
  - Added
</div>

- Retrieves the SAST settings for an org

## 2023-05-29


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi/{version}</span>
  - Updated
</div>

- removed the optional property `errors/items/links` from the response with the `400` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `401` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `404` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `500` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi</span>
  - Updated
</div>

- removed the optional property `errors/items/links` from the response with the `400` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `401` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `404` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `500` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)

## 2023-04-28


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs</span>
  - Added
</div>

- Get a paginated list of organizations you have access to.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}</span>
  - Added
</div>

- Get the full details of an organization.



<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}</span>
  - Added
</div>

- Update the details of an organization

## 2023-04-27


<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/invites</span>
  - Updated
</div>

- added the new required request property `data`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- the `data/attributes/role` response`s property type/format changed from `string`/`` to `string`/`uuid` for status `201`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property `email`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the request property `role`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- added the new `org_invitation` enum value to the `data/type` response property for the response status `201`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/invites</span>
  - Updated
</div>

- the `data/items/attributes/role` response`s property type/format changed from `string`/`` to `string`/`uuid` for status `200`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- added the new `org_invitation` enum value to the `data/items/type` response property for the response status `200`
  ![Badge](https://img.shields.io/badge/Breaking-yellow)

## 2023-03-30


<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/packages/issues</span>
  - Added
</div>

- This endpoint is not available to all customers. If you are interested please contact support. Query issues for a batch of packages identified by Package URL (purl). Only direct vulnerabilities are returned, transitive vulnerabilities (from dependencies) are not returned because they can vary depending on context.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi</span>
  - Updated
</div>

- added the optional property `errors/items/links` to the response with the `400` status

- added the optional property `errors/items/links` to the response with the `401` status

- added the optional property `errors/items/links` to the response with the `404` status

- added the optional property `errors/items/links` to the response with the `500` status




<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi/{version}</span>
  - Updated
</div>

- added the optional property `errors/items/links` to the response with the `400` status

- added the optional property `errors/items/links` to the response with the `401` status

- added the optional property `errors/items/links` to the response with the `404` status

- added the optional property `errors/items/links` to the response with the `500` status


## 2023-03-29


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi</span>
  - Updated
</div>

- removed the optional property `errors/items/links` from the response with the `400` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `401` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `404` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `500` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi/{version}</span>
  - Updated
</div>

- removed the optional property `errors/items/links` from the response with the `400` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `401` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `404` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `500` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)

## 2023-03-08


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects/{project_id}/sbom</span>
  - Added
</div>

- This endpoint lets you retrieve the SBOM document of a software project.
  It supports the following formats:
* CycloneDX version 1.4 in JSON (set `format` to `cyclonedx1.4+json`).
* CycloneDX version 1.4 in XML (set `format` to `cyclonedx1.4+xml`).
* SPDX version 2.3 in JSON (set `format` to `spdx2.3+json`).

By default it will respond with an empty JSON:API response.

## 2023-01-30


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects</span>
  - Added
</div>

- List all Projects for an Org.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects/{project_id}</span>
  - Added
</div>

- Get one project of the organization by project ID.



<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/projects/{project_id}</span>
  - Added
</div>

- Updates one project of the organization by project ID.

## 2022-11-14


<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/slack_app/{bot_id}/projects/{project_id}</span>
  - Added
</div>

- Create Slack settings override for a project.



<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/slack_app/{bot_id}</span>
  - Added
</div>

- Remove the given Slack App integration



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/slack_app/{bot_id}</span>
  - Added
</div>

- Get Slack integration default notification settings for the provided tenant ID and bot ID.



<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/slack_app/{bot_id}</span>
  - Added
</div>

- Create new Slack notification default settings for a given tenant.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/slack_app/{bot_id}/projects</span>
  - Added
</div>

- Slack notification settings overrides for projects. These settings overrides the default settings configured for the tenant.



<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/slack_app/{bot_id}/projects/{project_id}</span>
  - Added
</div>

- Remove Slack settings override for a project.



<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/slack_app/{bot_id}/projects/{project_id}</span>
  - Added
</div>

- Update Slack notification settings for a project.

## 2022-11-07


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/invites</span>
  - Added
</div>

- List pending user invitations to an organization.



<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/invites/{invite_id}</span>
  - Added
</div>

- Cancel a pending user invitations to an organization.

## 2022-10-06


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/slack_app/{tenant_id}/channels</span>
  - Added
</div>

- Requires the Snyk Slack App to be set up for this org, will retrieve a list of channels the Snyk Slack App can access. Note that it is currently only possible to page forwards through this collection, no prev links will be generated and the ending_before parameter will not function.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/slack_app/{tenant_id}/channels/{channel_id}</span>
  - Added
</div>

- Requires the Snyk Slack App to be set up for this org. It will return the Slack channel name for the provided Slack channel ID.

## 2022-09-14


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/packages/{purl}/issues</span>
  - Added
</div>

- Query issues for a specific package version identified by Package URL (purl). Snyk returns only direct vulnerabilities. Transitive vulnerabilities (from dependencies) are not returned because they can vary depending on context.

## 2022-04-06


<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/invites</span>
  - Added
</div>

- Invite a user to an organization with a role.

## 2022-03-11


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi</span>
  - Updated
</div>

- added the optional property `errors/items/links` to the response with the `400` status

- added the optional property `errors/items/links` to the response with the `401` status

- added the optional property `errors/items/links` to the response with the `404` status

- added the optional property `errors/items/links` to the response with the `500` status




<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi/{version}</span>
  - Updated
</div>

- added the optional property `errors/items/links` to the response with the `400` status

- added the optional property `errors/items/links` to the response with the `401` status

- added the optional property `errors/items/links` to the response with the `404` status

- added the optional property `errors/items/links` to the response with the `500` status


## 2022-03-01


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/self/apps</span>
  - Added
</div>

- Get a list of apps that can act on your behalf.



<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/self/apps/{app_id}</span>
  - Added
</div>

- Revoke access for an app by app id



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi</span>
  - Updated
</div>

- removed the optional property `errors/items/links` from the response with the `400` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `401` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `404` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `500` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/app_bots</span>
  - Added
</div>

- Get a list of app bots authorized to an organization. Deprecated, use /orgs/{org_id}/apps/installs instead.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps</span>
  - Added
</div>

- Get a list of apps created by an organization. Deprecated, use /orgs/{org_id}/apps/creations instead.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps/{client_id}</span>
  - Added
</div>

- Get an App by client id. Deprecated, use /orgs/{org_id}/apps/creations/{app_id} instead.



<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps/{client_id}/secrets</span>
  - Added
</div>

- Manage client secrets for an app. Deprecated, use /orgs/{org_id}/apps/creations/{app_id}/secrets instead.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/openapi/{version}</span>
  - Updated
</div>

- removed the optional property `errors/items/links` from the response with the `400` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `401` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `404` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
- removed the optional property `errors/items/links` from the response with the `500` status
  ![Badge](https://img.shields.io/badge/Breaking-yellow)



<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/app_bots/{bot_id}</span>
  - Added
</div>

- Revoke app bot authorization. Deprecated, use /orgs/{org_id}/apps/installs/{install_id} instead.



<div class="openapi-target">
  <span class="openapi-method openapi-method-post">POST</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps</span>
  - Added
</div>

- Create a new app for an organization. Deprecated, use /orgs/{org_id}/apps/creations instead.



<div class="openapi-target">
  <span class="openapi-method openapi-method-delete">DELETE</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps/{client_id}</span>
  - Added
</div>

- Delete an app by app id. Deprecated, use /orgs/{org_id}/apps/creations/{app_id} instead.



<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/apps/{client_id}</span>
  - Added
</div>

- Update app attributes. Deprecated, use /orgs/{org_id}/apps/creations/{app_id} instead.

## 2021-09-29


<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/groups/{group_id}/settings/iac</span>
  - Added
</div>

- Get the Infrastructure as Code Settings for a group.



<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/groups/{group_id}/settings/iac</span>
  - Added
</div>

- Update the Infrastructure as Code Settings for a group.



<div class="openapi-target">
  <span class="openapi-method openapi-method-get">GET</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/settings/iac</span>
  - Added
</div>

- Get the Infrastructure as Code Settings for an org.



<div class="openapi-target">
  <span class="openapi-method openapi-method-patch">PATCH</span>
  <span class="openapi-url"><span><span>https://api.snyk.io/rest</span></span>/orgs/{org_id}/settings/iac</span>
  - Added
</div>

- Update the Infrastructure as Code Settings for an org.

