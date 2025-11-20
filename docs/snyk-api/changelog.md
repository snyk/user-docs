## 2025-11-05

### GET - `/orgs/{org_id}/packages/{purl}/issues` - Updated
- for the `query` request parameter `limit`, the type/format was changed from `number`/`` to `integer`/``
![Badge](https://img.shields.io/badge/Breaking-yellow)
- for the `query` request parameter `offset`, the type/format was changed from `number`/`` to `integer`/``
![Badge](https://img.shields.io/badge/Breaking-yellow)
- for the `query` request parameter `limit`, the max was set to `1000.00`
![Badge](https://img.shields.io/badge/Breaking-yellow)
- for the `query` request parameter `limit`, the min was set to `1.00`
![Badge](https://img.shields.io/badge/Breaking-yellow)
- for the `query` request parameter `offset`, the min was set to `0.00`
![Badge](https://img.shields.io/badge/Breaking-yellow)
- api operation id `fetchIssuesPerPurl` removed and replaced with `getIssuesPerPurl`

- added the optional property `meta/match` to the response with the `200` status



### POST - `/orgs/{org_id}/packages/issues` - Updated
- added the optional property `meta/packages` to the response with the `200` status


## 2025-09-28 - Updated 2025-11-13

### POST - `/orgs/{org_id}/export` - Updated
- added the new optional request property `data/attributes/filters/empty_project_tags`



### POST - `/groups/{group_id}/export` - Updated
- added the new optional request property `data/attributes/filters/empty_project_tags`


## 2025-09-28 - Updated 2025-11-13

### GET - `/orgs/{org_id}/issues` - Updated
- added the optional property `data/items/attributes/coordinates/items/created_at` to the response with the `200` status

- added the optional property `data/items/attributes/coordinates/items/last_introduced_at` to the response with the `200` status

- added the optional property `data/items/attributes/coordinates/items/last_resolved_at` to the response with the `200` status

- added the optional property `data/items/attributes/coordinates/items/last_resolved_details` to the response with the `200` status

- added the optional property `data/items/attributes/coordinates/items/state` to the response with the `200` status

- added the optional property `data/items/attributes/coordinates/items/updated_at` to the response with the `200` status



### GET - `/orgs/{org_id}/issues/{issue_id}` - Updated
- added the optional property `data/attributes/coordinates/items/created_at` to the response with the `200` status

- added the optional property `data/attributes/coordinates/items/last_introduced_at` to the response with the `200` status

- added the optional property `data/attributes/coordinates/items/last_resolved_at` to the response with the `200` status

- added the optional property `data/attributes/coordinates/items/last_resolved_details` to the response with the `200` status

- added the optional property `data/attributes/coordinates/items/state` to the response with the `200` status

- added the optional property `data/attributes/coordinates/items/updated_at` to the response with the `200` status



### GET - `/groups/{group_id}/issues` - Updated
- added the optional property `data/items/attributes/coordinates/items/created_at` to the response with the `200` status

- added the optional property `data/items/attributes/coordinates/items/last_introduced_at` to the response with the `200` status

- added the optional property `data/items/attributes/coordinates/items/last_resolved_at` to the response with the `200` status

- added the optional property `data/items/attributes/coordinates/items/last_resolved_details` to the response with the `200` status

- added the optional property `data/items/attributes/coordinates/items/state` to the response with the `200` status

- added the optional property `data/items/attributes/coordinates/items/updated_at` to the response with the `200` status



### GET - `/groups/{group_id}/issues/{issue_id}` - Updated
- added the optional property `data/attributes/coordinates/items/created_at` to the response with the `200` status

- added the optional property `data/attributes/coordinates/items/last_introduced_at` to the response with the `200` status

- added the optional property `data/attributes/coordinates/items/last_resolved_at` to the response with the `200` status

- added the optional property `data/attributes/coordinates/items/last_resolved_details` to the response with the `200` status

- added the optional property `data/attributes/coordinates/items/state` to the response with the `200` status

- added the optional property `data/attributes/coordinates/items/updated_at` to the response with the `200` status


## 2025-09-28 - Updated 2025-10-22

### POST - `/orgs/{org_id}/export` - Updated
- added the new optional request property `data/attributes/url_expiration_seconds`



### POST - `/groups/{group_id}/export` - Updated
- added the new optional request property `data/attributes/url_expiration_seconds`


# Changelog

### 2025-09-28 - Updated 2025-10-14

#### GET - `/orgs/{org_id}/issues` - Updated

* the `data/items/attributes/coordinates/items/representations/items/oneOf[subschema #3]/cloud_resource/resource/name` response property`s maxLength was unset from` 256`for the response status`200\` ![Badge](https://img.shields.io/badge/Breaking-yellow)

#### GET - `/orgs/{org_id}/issues/{issue_id}` - Updated

* the `data/attributes/coordinates/items/representations/items/oneOf[subschema #3]/cloud_resource/resource/name` response property`s maxLength was unset from` 256`for the response status`200\` ![Badge](https://img.shields.io/badge/Breaking-yellow)

#### GET - `/groups/{group_id}/issues` - Updated

* the `data/items/attributes/coordinates/items/representations/items/oneOf[subschema #3]/cloud_resource/resource/name` response property`s maxLength was unset from` 256`for the response status`200\` ![Badge](https://img.shields.io/badge/Breaking-yellow)

#### GET - `/groups/{group_id}/issues/{issue_id}` - Updated

* the `data/attributes/coordinates/items/representations/items/oneOf[subschema #3]/cloud_resource/resource/name` response property`s maxLength was unset from` 256`for the response status`200\` ![Badge](https://img.shields.io/badge/Breaking-yellow)

### 2025-09-28

#### POST - `/orgs/{org_id}/policies` - Updated

* added the new optional request property `data/attributes/source`

### 2024-10-15 - Updated 2025-09-29

#### POST - `/orgs/{org_id}/export` - Updated

* added the new optional request property `data/attributes/filters/product_name`
* added the new optional request property `data/attributes/filters/project_tags`
* added the new optional request property `data/attributes/filters/project_type`

#### POST - `/groups/{group_id}/export` - Updated

* added the new optional request property `data/attributes/filters/product_name`
* added the new optional request property `data/attributes/filters/project_tags`
* added the new optional request property `data/attributes/filters/project_type`

## Changelog

#### 2024-10-15 - Updated 2025-09-08

**POST - `/orgs/{org_id}/policies` - Updated**

* removed the `cancelled` enum value from the `data/attributes/review` response property for the response status `201`

**GET - `/orgs/{org_id}/policies` - Updated**

* removed the `cancelled` enum value from the `data/items/attributes/review` response property for the response status `200`

**PATCH - `/orgs/{org_id}/policies/{policy_id}` - Updated**

* removed the enum value `cancelled` of the request property `data/attributes/review` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the `cancelled` enum value from the `data/attributes/review` response property for the response status `200`

**GET - `/orgs/{org_id}/policies/{policy_id}` - Updated**

* removed the `cancelled` enum value from the `data/attributes/review` response property for the response status `200`

#### 2024-10-15 - Updated 2025-08-15

**GET - `/orgs/{org_id}/policies` - Updated**

* added the new enum value `ignore-type` to the `query` request parameter `order_by`
* added the new enum value `requested-by` to the `query` request parameter `order_by`

#### 2024-10-15 - Updated 2025-08-14

**PATCH - `/orgs/{org_id}/policies/{policy_id}` - Updated**

* added the new `cancelled` enum value to the request property `data/attributes/review`

#### 2024-10-15 - Updated 2025-08-06

**POST - `/orgs/{org_id}/export` - Updated**

* removed the request property `data/attributes/destination` ![Badge](https://img.shields.io/badge/Breaking-yellow)

**POST - `/groups/{group_id}/export` - Updated**

* removed the request property `data/attributes/destination` ![Badge](https://img.shields.io/badge/Breaking-yellow)

#### 2024-10-15 - Updated 2025-07-07

**POST - `/orgs/{org_id}/service_accounts` - Updated**

* added the new `access_token` enum value to the `data/attributes/auth_type` response property for the response status `201` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new optional request property `data/attributes/access_token_expires_at`
* added the new `access_token` enum value to the request property `data/attributes/auth_type`
* added the optional property `data/attributes/access_token` to the response with the `201` status
* added the optional property `data/attributes/access_token_expires_at` to the response with the `201` status
* added the optional property `data/attributes/created_at` to the response with the `201` status

**GET - `/orgs/{org_id}/service_accounts` - Updated**

* added the new `access_token` enum value to the `data/items/attributes/auth_type` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the optional property `data/items/attributes/access_token` to the response with the `200` status
* added the optional property `data/items/attributes/access_token_expires_at` to the response with the `200` status
* added the optional property `data/items/attributes/created_at` to the response with the `200` status

**PATCH - `/orgs/{org_id}/service_accounts/{serviceaccount_id}` - Updated**

* added the new `access_token` enum value to the `data/attributes/auth_type` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the optional property `data/attributes/access_token` to the response with the `200` status
* added the optional property `data/attributes/access_token_expires_at` to the response with the `200` status
* added the optional property `data/attributes/created_at` to the response with the `200` status

**GET - `/orgs/{org_id}/service_accounts/{serviceaccount_id}` - Updated**

* added the new `access_token` enum value to the `data/attributes/auth_type` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the optional property `data/attributes/access_token` to the response with the `200` status
* added the optional property `data/attributes/access_token_expires_at` to the response with the `200` status
* added the optional property `data/attributes/created_at` to the response with the `200` status

**POST - `/orgs/{org_id}/service_accounts/{serviceaccount_id}/secrets` - Updated**

* added the new `access_token` enum value to the `data/attributes/auth_type` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the optional property `data/attributes/access_token` to the response with the `200` status
* added the optional property `data/attributes/access_token_expires_at` to the response with the `200` status
* added the optional property `data/attributes/created_at` to the response with the `200` status

**GET - `/orgs/{org_id}/packages/{purl}/issues` - Updated**

* added `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` to the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**POST - `/orgs/{org_id}/packages/issues` - Updated**

* added `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` to the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**POST - `/groups/{group_id}/service_accounts` - Updated**

* added the new `access_token` enum value to the `data/attributes/auth_type` response property for the response status `201` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new optional request property `data/attributes/access_token_expires_at`
* added the new `access_token` enum value to the request property `data/attributes/auth_type`
* added the optional property `data/attributes/access_token` to the response with the `201` status
* added the optional property `data/attributes/access_token_expires_at` to the response with the `201` status
* added the optional property `data/attributes/created_at` to the response with the `201` status

**GET - `/groups/{group_id}/service_accounts` - Updated**

* added the new `access_token` enum value to the `data/items/attributes/auth_type` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the optional property `data/items/attributes/access_token` to the response with the `200` status
* added the optional property `data/items/attributes/access_token_expires_at` to the response with the `200` status
* added the optional property `data/items/attributes/created_at` to the response with the `200` status
* added the optional property `meta` to the response with the `200` status

**PATCH - `/groups/{group_id}/service_accounts/{serviceaccount_id}` - Updated**

* added the new `access_token` enum value to the `data/attributes/auth_type` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the optional property `data/attributes/access_token` to the response with the `200` status
* added the optional property `data/attributes/access_token_expires_at` to the response with the `200` status
* added the optional property `data/attributes/created_at` to the response with the `200` status

**GET - `/groups/{group_id}/service_accounts/{serviceaccount_id}` - Updated**

* added the new `access_token` enum value to the `data/attributes/auth_type` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the optional property `data/attributes/access_token` to the response with the `200` status
* added the optional property `data/attributes/access_token_expires_at` to the response with the `200` status
* added the optional property `data/attributes/created_at` to the response with the `200` status

**POST - `/groups/{group_id}/service_accounts/{serviceaccount_id}/secrets` - Updated**

* added the new `access_token` enum value to the `data/attributes/auth_type` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the optional property `data/attributes/access_token` to the response with the `200` status
* added the optional property `data/attributes/access_token_expires_at` to the response with the `200` status
* added the optional property `data/attributes/created_at` to the response with the `200` status

### Changelog

**2024-10-15 - Updated 2025-06-12**

**GET - `/tenants` - Added**

* Get a list of all Tenants which the calling user is a member of

**PATCH - `/tenants/{tenant_id}` - Added**

* Update the details of a tenant

**Required permissions**

* `Edit Tenant Details (tenant.edit)`

**GET - `/tenants/{tenant_id}` - Added**

* Get the full details of a Tenant.

**Required permissions**

* `View Tenant Details (tenant.read)`

**POST - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections/{connection_id}/bulk_migration` - Added**

* Performs bulk migration for integrations from legacy to universal broker

**Required permissions**

* `View Tenant Details (tenant.read)`

**GET - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections/{connection_id}/bulk_migration` - Added**

* Lists organization IDs associated with a connection type to be bulk migrated to universal broker

**Required permissions**

* `View Tenant Details (tenant.read)`

**GET - `/orgs/{org_id}` - Updated**

* added the new optional `query` request parameter `expand`
* added the optional property `data/relationships` to the response with the `200` status

**GET - `/orgs/{org_id}/projects` - Updated**

* added the optional property `data/items/attributes/settings/auto_dependency_upgrade/is_inherited` to the response with the `200` status

**PATCH - `/orgs/{org_id}/projects/{project_id}` - Updated**

* added the optional property `data/attributes/settings/auto_dependency_upgrade/is_inherited` to the response with the `200` status

**GET - `/orgs/{org_id}/projects/{project_id}` - Updated**

* added the optional property `data/attributes/settings/auto_dependency_upgrade/is_inherited` to the response with the `200` status

**GET - `/orgs/{org_id}/packages/{purl}/issues` - Updated**

* removed `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` from the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**POST - `/orgs/{org_id}/packages/issues` - Updated**

* removed `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` from the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**2024-10-15 - Updated 2025-06-04**

**GET - `/groups/{group_id}/assets/{asset_id}` - Added**

* Get an Asset by its ID

**Required permissions**

* `View Groups (group.read)`

**GET - `/groups/{group_id}/assets/{asset_id}/relationships/projects` - Added**

* List asset projects with pagination

**Required permissions**

* `View Groups (group.read)`

**GET - `/groups/{group_id}/assets/{asset_id}/relationships/assets` - Added**

* List related assets with pagination

**Required permissions**

* `View Groups (group.read)`

**POST - `/groups/{group_id}/assets/search` - Added**

* List Assets with filters

**Required permissions**

* `View Groups (group.read)`

**2024-10-15 - Updated 2025-05-27**

**GET - `/orgs/{org_id}` - Updated**

* the response property `data` became optional for the status `200`![Badge](https://img.shields.io/badge/Breaking-yellow)
* the response property `jsonapi` became optional for the status `200`![Badge](https://img.shields.io/badge/Breaking-yellow)
* the response property `links` became optional for the status `200`![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the optional property `links/first` from the response with the `200` status![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the optional property `links/last` from the response with the `200` status![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the optional property `links/next` from the response with the `200` status![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the optional property `links/prev` from the response with the `200` status![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the optional property `links/related` from the response with the `200` status![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the non-success response with the status `409`
* added the optional property `data/attributes/created_at` to the response with the `200` status
* added the optional property `data/attributes/updated_at` to the response with the `200` status
* the response property `data/attributes` became required for the status `200`
* the `data/type` response`s property pattern` ^\[a-z]\[a-z0-9]_(\_\[a-z]\[a-z0-9]_)\*$`was added for the status`200\`

**2024-10-15 - Updated 2025-05-20**

**GET - `/orgs/{org_id}/projects` - Updated**

* added the optional property `data/items/attributes/settings/pull_requests/is_enabled` to the response with the `200` status

**PATCH - `/orgs/{org_id}/projects/{project_id}` - Updated**

* added the optional property `data/attributes/settings/pull_requests/is_enabled` to the response with the `200` status

**GET - `/orgs/{org_id}/projects/{project_id}` - Updated**

* added the optional property `data/attributes/settings/pull_requests/is_enabled` to the response with the `200` status

**2024-10-15 - Updated 2025-05-15**

**GET - `/orgs/{org_id}/policies` - Updated**

* added the new optional `query` request parameter `order_by`
* added the new optional `query` request parameter `order_direction`

**2024-10-15 - Updated 2025-05-08**

**GET - `/orgs/{org_id}/packages/{purl}/issues` - Updated**

* added `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` to the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**POST - `/orgs/{org_id}/packages/issues` - Updated**

* added `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` to the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**2024-10-15 - Updated 2025-04-28**

**GET - `/orgs/{org_id}/packages/{purl}/issues` - Updated**

* removed `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` from the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**POST - `/orgs/{org_id}/packages/issues` - Updated**

* removed `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` from the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**GET - `/orgs/{org_id}/issues` - Updated**

* added the optional property `data/items/attributes/coordinates/items/representations/items/oneOf[subschema #4]/sourceLocation/commit_id` to the response with the `200` status
* added the optional property `data/items/attributes/key_asset` to the response with the `200` status

**GET - `/orgs/{org_id}/issues/{issue_id}` - Updated**

* added the optional property `data/attributes/coordinates/items/representations/items/oneOf[subschema #4]/sourceLocation/commit_id` to the response with the `200` status
* added the optional property `data/attributes/key_asset` to the response with the `200` status

**GET - `/groups/{group_id}/issues` - Updated**

* added the optional property `data/items/attributes/coordinates/items/representations/items/oneOf[subschema #4]/sourceLocation/commit_id` to the response with the `200` status
* added the optional property `data/items/attributes/key_asset` to the response with the `200` status

**GET - `/groups/{group_id}/issues/{issue_id}` - Updated**

* added the optional property `data/attributes/coordinates/items/representations/items/oneOf[subschema #4]/sourceLocation/commit_id` to the response with the `200` status
* added the optional property `data/attributes/key_asset` to the response with the `200` status

**2024-10-15 - Updated 2025-04-25**

**POST - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments` - Added**

* Creates a new Broker Deployment for an installation

**GET - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments` - Added**

* List Broker deployments for a given install ID

**PATCH - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}` - Added**

* Updates a Broker deployment for a given install ID

**DELETE - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}` - Added**

* Delete a Broker deployment for a given install ID

**POST - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/credentials` - Added**

* Creates a new Deployment credential

**GET - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/credentials` - Added**

* List Deployment credentials for a given deployment ID

**PATCH - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/credentials/{credential_id}` - Added**

* Updates a Deployment credential for an deployment

**GET - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/credentials/{credential_id}` - Added**

* Get all Deployment credential data for an deployment

**DELETE - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/credentials/{credential_id}` - Added**

* Deletes an existing Deployment credential for an deployment

**POST - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/contexts` - Added**

* Creates a new Broker Context

**GET - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/contexts` - Added**

* List Deployment contexts for a given deployment ID

**POST - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections` - Added**

* Creates a new Broker connection for an deployment

**GET - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections` - Added**

* List all Broker connections for a given deployment

**DELETE - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections` - Added**

* Deletes all existing Broker connections for an deployment

**PATCH - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections/{connection_id}` - Added**

* Updates a Broker connection for an deployment

**GET - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections/{connection_id}` - Added**

* Get all Broker connection data for an deployment

**DELETE - `/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections/{connection_id}` - Added**

* Deletes an existing Broker connection for an deployment

**PATCH - `/tenants/{tenant_id}/brokers/installs/{install_id}/contexts/{context_id}` - Added**

* Updates a Broker Context for an deployment

**GET - `/tenants/{tenant_id}/brokers/installs/{install_id}/contexts/{context_id}` - Added**

* List Broker context for a given broker context ID

**DELETE - `/tenants/{tenant_id}/brokers/installs/{install_id}/contexts/{context_id}` - Added**

* Deletes an existing broker context

**DELETE - `/tenants/{tenant_id}/brokers/installs/{install_id}/contexts/{context_id}/integrations/{integration_id}` - Added**

* Deletes an existing Broker context association for an integration

**PATCH - `/tenants/{tenant_id}/brokers/installs/{install_id}/contexts/{context_id}/integration` - Added**

* Updates an integration to be associated with a Broker context

**GET - `/tenants/{tenant_id}/brokers/installs/{install_id}/connections/{connection_id}/contexts` - Added**

* List Broker contexts for a given broker connection ID

**GET - `/tenants/{tenant_id}/brokers/deployments` - Added**

* List Broker deployments for the tenant

**DELETE - `/tenants/{tenant_id}/brokers/connections/{connection_id}/orgs/{org_id}/integrations/{integration_id}` - Added**

* Deletes an existing Broker connection for an deployment

**POST - `/tenants/{tenant_id}/brokers/connections/{connection_id}/orgs/{org_id}/integration` - Added**

* Configures integrations to use the Broker connection for an deployment

**GET - `/tenants/{tenant_id}/brokers/connections/{connection_id}/integrations` - Added**

* Get all integrations using the Broker connection

**POST - `/orgs/{org_id}/policies` - Updated**

* the `data/attributes/action/data/reason` request property`s maxLength was set to` 10000\`![Badge](https://img.shields.io/badge/Breaking-yellow)
* the `data/attributes/name` request property`s maxLength was set to` 255\`![Badge](https://img.shields.io/badge/Breaking-yellow)

**PATCH - `/orgs/{org_id}/policies/{policy_id}` - Updated**

* the `data/attributes/action/data/reason` request property`s maxLength was set to` 10000\`![Badge](https://img.shields.io/badge/Breaking-yellow)
* the `data/attributes/name` request property`s maxLength was set to` 255\`![Badge](https://img.shields.io/badge/Breaking-yellow)

**GET - `/orgs/{org_id}/brokers/connections` - Added**

* List all Broker connections integrated with a given org

**2024-10-15 - Updated 2025-04-22**

**GET - `/orgs/{org_id}/policies` - Updated**

* added the new optional `query` request parameter `search`

**2024-10-15 - Updated 2025-04-02**

**POST - `/orgs/{org_id}/policies` - Updated**

* added the new optional request property `data/meta`

**2024-10-15 - Updated 2025-04-01**

**POST - `/orgs/{org_id}/memberships` - Updated**

* the request property `data` became required![Badge](https://img.shields.io/badge/Breaking-yellow)
* the request property `data/relationships` became required![Badge](https://img.shields.io/badge/Breaking-yellow)
* the request property `data/relationships/org/data` became required![Badge](https://img.shields.io/badge/Breaking-yellow)
* the request property `data/relationships/org/data/id` became required![Badge](https://img.shields.io/badge/Breaking-yellow)
* the request property `data/relationships/org/data/type` became required![Badge](https://img.shields.io/badge/Breaking-yellow)
* the request property `data/relationships/role/data` became required![Badge](https://img.shields.io/badge/Breaking-yellow)
* the request property `data/relationships/role/data/id` became required![Badge](https://img.shields.io/badge/Breaking-yellow)
* the request property `data/relationships/role/data/type` became required![Badge](https://img.shields.io/badge/Breaking-yellow)
* the request property `data/relationships/user/data` became required![Badge](https://img.shields.io/badge/Breaking-yellow)
* the request property `data/relationships/user/data/id` became required![Badge](https://img.shields.io/badge/Breaking-yellow)
* the request property `data/relationships/user/data/type` became required![Badge](https://img.shields.io/badge/Breaking-yellow)
* the request property `data/type` became required![Badge](https://img.shields.io/badge/Breaking-yellow)

**2024-10-15 - Updated 2025-03-19**

**PATCH - `/orgs/{org_id}/memberships/{membership_id}` - Updated**

* the request property `data` became required![Badge](https://img.shields.io/badge/Breaking-yellow)

#### Changelog

**2024-10-15 - Updated 2025-03-04**

**GET - `/orgs/{org_id}/packages/{purl}/issues` - Updated**

* added `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` to the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**POST - `/orgs/{org_id}/packages/issues` - Updated**

* added `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` to the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**2024-10-15 - Updated 2025-02-11**

**GET - `/orgs/{org_id}/packages/{purl}/issues` - Updated**

* removed `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` from the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**POST - `/orgs/{org_id}/packages/issues` - Updated**

* removed `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` from the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**2024-10-15 - Updated 2025-02-04**

**POST - `/orgs/{org_id}/policies` - Added**

* Create a new org-level policy.

_Org level Policy APIs Access Notice:_ Access to our Org level Policy APIs is currently restricted via "snykCodeConsistentIgnores" feature flag and will result in a 403 Forbidden error without the flag enabled. Please contact your account representative for eligibility requirements.

**GET - `/orgs/{org_id}/policies` - Added**

* Get all policies for the requested organisation.

_Org level Policy APIs Access Notice:_ Access to our Org level Policy APIs is currently restricted via "snykCodeConsistentIgnores" feature flag and will result in a 403 Forbidden error without the flag enabled. Please contact your account representative for eligibility requirements.

**PATCH - `/orgs/{org_id}/policies/{policy_id}` - Added**

* Update the org-level policy.

_Org level Policy APIs Access Notice:_ Access to our Org level Policy APIs is currently restricted via "snykCodeConsistentIgnores" feature flag and will result in a 403 Forbidden error without the flag enabled. Please contact your account representative for eligibility requirements.

**GET - `/orgs/{org_id}/policies/{policy_id}` - Added**

* Get a specific org-level policy based on its ID.

_Org level Policy APIs Access Notice:_ Access to our Org level Policy APIs is currently restricted via "snykCodeConsistentIgnores" feature flag and will result in a 403 Forbidden error without the flag enabled. Please contact your account representative for eligibility requirements.

**DELETE - `/orgs/{org_id}/policies/{policy_id}` - Added**

* Delete an existing org-level policy.

_Org level Policy APIs Access Notice:_ Access to our Org level Policy APIs is currently restricted via "snykCodeConsistentIgnores" feature flag and will result in a 403 Forbidden error without the flag enabled. Please contact your account representative for eligibility requirements.

**2024-10-15 - Updated 2025-01-22**

**GET - `/orgs/{org_id}/packages/{purl}/issues` - Updated**

* added `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` to the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**POST - `/orgs/{org_id}/packages/issues` - Updated**

* added `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` to the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**2024-10-15 - Updated 2025-01-13**

**GET - `/orgs/{org_id}/issues` - Updated**

* added the optional property `data/items/attributes/exploit_details` to the response with the `200` status
* added the optional property `data/items/attributes/severities` to the response with the `200` status

**GET - `/orgs/{org_id}/issues/{issue_id}` - Updated**

* added the optional property `data/attributes/exploit_details` to the response with the `200` status
* added the optional property `data/attributes/severities` to the response with the `200` status

**GET - `/groups/{group_id}/issues` - Updated**

* added the optional property `data/items/attributes/exploit_details` to the response with the `200` status
* added the optional property `data/items/attributes/severities` to the response with the `200` status

**GET - `/groups/{group_id}/issues/{issue_id}` - Updated**

* added the optional property `data/attributes/exploit_details` to the response with the `200` status
* added the optional property `data/attributes/severities` to the response with the `200` status

**2024-10-15 - Updated 2025-01-07**

**GET - `/orgs/{org_id}/issues` - Updated**

* added `loaded_package` discriminator mapping keys to the `data/items/attributes/risk/factors/items/` response property for the response status `200`
* added `#/components/schemas/LoadedPackageRiskFactor` to the `data/items/attributes/risk/factors/items/` response property `oneOf` list for the response status `200`

**GET - `/orgs/{org_id}/issues/{issue_id}` - Updated**

* added `loaded_package` discriminator mapping keys to the `data/attributes/risk/factors/items/` response property for the response status `200`
* added `#/components/schemas/LoadedPackageRiskFactor` to the `data/attributes/risk/factors/items/` response property `oneOf` list for the response status `200`

**GET - `/groups/{group_id}/issues` - Updated**

* added `loaded_package` discriminator mapping keys to the `data/items/attributes/risk/factors/items/` response property for the response status `200`
* added `#/components/schemas/LoadedPackageRiskFactor` to the `data/items/attributes/risk/factors/items/` response property `oneOf` list for the response status `200`

**GET - `/groups/{group_id}/issues/{issue_id}` - Updated**

* added `loaded_package` discriminator mapping keys to the `data/attributes/risk/factors/items/` response property for the response status `200`
* added `#/components/schemas/LoadedPackageRiskFactor` to the `data/attributes/risk/factors/items/` response property `oneOf` list for the response status `200`

**2024-10-15 - Updated 2024-12-09**

**GET - `/orgs/{org_id}` - Updated**

* the response property `data` became optional for the status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the response property `jsonapi` became optional for the status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the response property `links` became optional for the status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the optional property `links/first` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the optional property `links/last` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the optional property `links/next` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the optional property `links/prev` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the optional property `links/related` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the non-success response with the status `409`
* added the optional property `data/attributes/created_at` to the response with the `200` status
* added the optional property `data/attributes/updated_at` to the response with the `200` status
* the response property `data/attributes` became required for the status `200`
* the `data/type` response`s property pattern` ^\[a-z]\[a-z0-9]_(\_\[a-z]\[a-z0-9]_)\*$`was added for the status`200\`

**2024-10-15 - Updated 2024-11-28**

**GET - `/orgs/{org_id}/projects/{project_id}/sbom` - Updated**

* added the new optional `query` request parameter `exclude`

**2024-10-15 - Updated 2024-11-06**

**GET - `/orgs/{org_id}/packages/{purl}/issues` - Updated**

* removed `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` from the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**POST - `/orgs/{org_id}/packages/issues` - Updated**

* removed `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` from the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**2024-10-15 - Updated 2024-10-31**

**GET - `/orgs/{org_id}/packages/{purl}/issues` - Updated**

* added `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` to the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**POST - `/orgs/{org_id}/packages/issues` - Updated**

* added `#/components/schemas/ResourcePathRepresentation, #/components/schemas/PackageRepresentation` to the `data/items/attributes/coordinates/items/representations/items/` response property `anyOf` list for the response status `200`

**2024-10-15 - Updated 2024-10-30**

**GET - `/orgs/{org_id}/issues` - Updated**

* added the new `function` enum value to the `data/items/attributes/coordinates/items/reachability` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new `no-info` enum value to the `data/items/attributes/coordinates/items/reachability` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new `not-applicable` enum value to the `data/items/attributes/coordinates/items/reachability` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new `package` enum value to the `data/items/attributes/coordinates/items/reachability` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)

**GET - `/orgs/{org_id}/issues/{issue_id}` - Updated**

* added the new `function` enum value to the `data/attributes/coordinates/items/reachability` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new `no-info` enum value to the `data/attributes/coordinates/items/reachability` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new `not-applicable` enum value to the `data/attributes/coordinates/items/reachability` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new `package` enum value to the `data/attributes/coordinates/items/reachability` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)

**GET - `/groups/{group_id}/issues` - Updated**

* added the new `function` enum value to the `data/items/attributes/coordinates/items/reachability` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new `no-info` enum value to the `data/items/attributes/coordinates/items/reachability` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new `not-applicable` enum value to the `data/items/attributes/coordinates/items/reachability` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new `package` enum value to the `data/items/attributes/coordinates/items/reachability` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)

**GET - `/groups/{group_id}/issues/{issue_id}` - Updated**

* added the new `function` enum value to the `data/attributes/coordinates/items/reachability` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new `no-info` enum value to the `data/attributes/coordinates/items/reachability` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new `not-applicable` enum value to the `data/attributes/coordinates/items/reachability` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new `package` enum value to the `data/attributes/coordinates/items/reachability` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)

**2024-10-15**

**Simplified API Versioning**

Going forward, Snyk will expose one API specification per version-date, rather than one for each stability. New versions of the Snyk API will only be published when necessitated by breaking changes. For newer versions, you should only specify the date for beta versions, i.e `2024-10-15` rather than `2024-10-15~beta`. It's important to note that existing versions won't be affected by these changes; this new approach only applies to upcoming new versions.

**2024-08-25 - Updated 2024-10-10**

**GET - `/self` - Updated**

* added `#/components/schemas/User20240422, #/components/schemas/ServiceAccount20240422` to the `data/attributes` response property `anyOf` list for the response status `200`
* removed `#/components/schemas/ServiceAccount` from the `data/attributes` response property `anyOf` list for the response status `200`

**GET - `/orgs/{org_id}/projects` - Updated**

* added `#/components/schemas/ProjectRelationshipsTarget20230215` to the `data/items/relationships/target` response property `oneOf` list for the response status `200`
* removed `#/components/schemas/ProjectRelationshipsTarget` from the `data/items/relationships/target` response property `oneOf` list for the response status `200`

**PATCH - `/orgs/{org_id}/projects/{project_id}` - Updated**

* added `#/components/schemas/ProjectRelationshipsTarget20230215` to the `data/relationships/target` response property `oneOf` list for the response status `200`
* removed `#/components/schemas/ProjectRelationshipsTarget` from the `data/relationships/target` response property `oneOf` list for the response status `200`

**GET - `/orgs/{org_id}/projects/{project_id}` - Updated**

* added `#/components/schemas/ProjectRelationshipsTarget20230215` to the `data/relationships/target` response property `oneOf` list for the response status `200`
* removed `#/components/schemas/ProjectRelationshipsTarget` from the `data/relationships/target` response property `oneOf` list for the response status `200`

**GET - `/orgs/{org_id}/packages/{purl}/issues` - Updated**

* removed the optional property `data/items/attributes/coordinates/items/representation` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the optional property `data/items/attributes/key` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the optional property `data/items/attributes/slots/exploit` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the optional property `data/items/attributes/severities/items/type` to the response with the `200` status
* added the optional property `data/items/attributes/severities/items/version` to the response with the `200` status
* added the optional property `data/items/attributes/slots/exploit_details` to the response with the `200` status
* added the required property `data/items/attributes/coordinates/items/representations` to the response with the `200` status

**POST - `/orgs/{org_id}/packages/issues` - Updated**

* removed the optional property `data/items/attributes/slots/exploit` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the optional property `data/items/attributes/severities/items/type` to the response with the `200` status
* added the optional property `data/items/attributes/severities/items/version` to the response with the `200` status
* added the optional property `data/items/attributes/slots/exploit_details` to the response with the `200` status

**GET - `/orgs/{org_id}/invites` - Updated**

* the `data/items/attributes/role` response`s property type/format changed from` string`/`uuid`to`string`/`` for status` 200\` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the `org_invitation` enum value from the `data/items/type` response property for the response status `200`

**2024-08-25 - Updated 2024-09-11**

**POST - `/orgs/{org_id}/apps` - Updated**

* added the new required request property `name` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new required request property `redirect_uris` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new required request property `scopes` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the request property `data` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new optional request property `access_token_ttl_seconds`
* added the new optional request property `context`

**GET - `/orgs/{org_id}/apps` - Updated**

* the `data/items/attributes/redirect_uris` response property`s minItems was decreased from` 1`to`0`for the response status`200\` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the response property `data/items/attributes/client_id` became required for the status `200`
* the response property `data/items/attributes/redirect_uris` became required for the status `200`

**PATCH - `/orgs/{org_id}/apps/{client_id}` - Updated**

* the `data/attributes/redirect_uris` response property`s minItems was decreased from` 1`to`0`for the response status`200\` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the request property `data` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new optional request property `access_token_ttl_seconds`
* added the new optional request property `name`
* added the new optional request property `redirect_uris`
* the response property `data/attributes/client_id` became required for the status `200`
* the response property `data/attributes/redirect_uris` became required for the status `200`

**GET - `/orgs/{org_id}/apps/{client_id}` - Updated**

* the `data/attributes/redirect_uris` response property`s minItems was decreased from` 1`to`0`for the response status`200\` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the response property `data/attributes/client_id` became required for the status `200`
* the response property `data/attributes/redirect_uris` became required for the status `200`

**2024-08-25 - Updated 2024-09-03**

**POST - `/groups/{group_id}/memberships` - Updated**

* the response property `data/relationships/group` became required for the status `201`
* the response property `data/relationships/group/data/attributes` became required for the status `201`
* the response property `data/relationships/group/data/attributes/name` became required for the status `201`
* the response property `data/relationships/group/data/id` became required for the status `201`
* the response property `data/relationships/role` became required for the status `201`
* the response property `data/relationships/role/data/attributes` became required for the status `201`
* the response property `data/relationships/role/data/attributes/name` became required for the status `201`
* the response property `data/relationships/role/data/id` became required for the status `201`
* the response property `data/relationships/user` became required for the status `201`
* the response property `data/relationships/user/data/attributes` became required for the status `201`
* the response property `data/relationships/user/data/attributes/email` became required for the status `201`
* the response property `data/relationships/user/data/attributes/name` became required for the status `201`
* the response property `data/relationships/user/data/attributes/username` became required for the status `201`
* the response property `data/relationships/user/data/id` became required for the status `201`

**GET - `/groups/{group_id}/memberships` - Updated**

* the response property `data/items/relationships/group` became required for the status `200`
* the response property `data/items/relationships/group/data/attributes` became required for the status `200`
* the response property `data/items/relationships/group/data/attributes/name` became required for the status `200`
* the response property `data/items/relationships/group/data/id` became required for the status `200`
* the response property `data/items/relationships/role` became required for the status `200`
* the response property `data/items/relationships/role/data/attributes` became required for the status `200`
* the response property `data/items/relationships/role/data/attributes/name` became required for the status `200`
* the response property `data/items/relationships/role/data/id` became required for the status `200`
* the response property `data/items/relationships/user` became required for the status `200`
* the response property `data/items/relationships/user/data/attributes` became required for the status `200`
* the response property `data/items/relationships/user/data/attributes/email` became required for the status `200`
* the response property `data/items/relationships/user/data/attributes/name` became required for the status `200`
* the response property `data/items/relationships/user/data/attributes/username` became required for the status `200`
* the response property `data/items/relationships/user/data/id` became required for the status `200`

**2024-08-25 - Updated 2024-08-30**

**POST - `/orgs/{org_id}/memberships` - Updated**

* the response property `data/relationships/org` became required for the status `201`
* the response property `data/relationships/org/data/attributes` became required for the status `201`
* the response property `data/relationships/org/data/attributes/name` became required for the status `201`
* the response property `data/relationships/org/data/id` became required for the status `201`
* the response property `data/relationships/role` became required for the status `201`
* the response property `data/relationships/role/data/attributes` became required for the status `201`
* the response property `data/relationships/role/data/attributes/name` became required for the status `201`
* the response property `data/relationships/role/data/id` became required for the status `201`
* the response property `data/relationships/user` became required for the status `201`
* the response property `data/relationships/user/data/attributes` became required for the status `201`
* the response property `data/relationships/user/data/attributes/email` became required for the status `201`
* the response property `data/relationships/user/data/attributes/name` became required for the status `201`
* the response property `data/relationships/user/data/attributes/username` became required for the status `201`
* the response property `data/relationships/user/data/id` became required for the status `201`

**GET - `/orgs/{org_id}/memberships` - Updated**

* the response property `data/items/relationships/org` became required for the status `200`
* the response property `data/items/relationships/org/data/attributes` became required for the status `200`
* the response property `data/items/relationships/org/data/attributes/name` became required for the status `200`
* the response property `data/items/relationships/org/data/id` became required for the status `200`
* the response property `data/items/relationships/role` became required for the status `200`
* the response property `data/items/relationships/role/data/attributes` became required for the status `200`
* the response property `data/items/relationships/role/data/attributes/name` became required for the status `200`
* the response property `data/items/relationships/role/data/id` became required for the status `200`
* the response property `data/items/relationships/user` became required for the status `200`
* the response property `data/items/relationships/user/data/attributes` became required for the status `200`
* the response property `data/items/relationships/user/data/attributes/email` became required for the status `200`
* the response property `data/items/relationships/user/data/attributes/name` became required for the status `200`
* the response property `data/items/relationships/user/data/attributes/username` became required for the status `200`
* the response property `data/items/relationships/user/data/id` became required for the status `200`

**2024-08-25**

**POST - `/orgs/{org_id}/memberships` - Added**

* Create a org membership for a user with role

**GET - `/orgs/{org_id}/memberships` - Added**

* Returns all memberships of the org

**PATCH - `/orgs/{org_id}/memberships/{membership_id}` - Added**

* Update a org membership for a user with role

**DELETE - `/orgs/{org_id}/memberships/{membership_id}` - Added**

* Remove a user\`s membership of the group.

**GET - `/groups/{group_id}/org_memberships` - Added**

* Get list of org memberships of a group user

**POST - `/groups/{group_id}/memberships` - Added**

* Create a group membership for a user with role

**GET - `/groups/{group_id}/memberships` - Added**

* Returns all memberships of the group

**PATCH - `/groups/{group_id}/memberships/{membership_id}` - Added**

* Update a role from a group membership

**DELETE - `/groups/{group_id}/memberships/{membership_id}` - Added**

* Deletes a membership from a group

**2024-08-22**

**GET - `/orgs/{org_id}/projects/{project_id}/sbom` - Updated**

* removed the required property `bomFormat` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the required property `components` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the required property `dependencies` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the required property `dependencies` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the required property `metadata` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the required property `metadata` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the required property `specVersion` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the required property `version` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the optional property `components` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new enum value `cyclonedx1.5+json` to the `query` request parameter `format`
* added the new enum value `cyclonedx1.5+xml` to the `query` request parameter `format`
* added the new enum value `cyclonedx1.6+json` to the `query` request parameter `format`
* added the new enum value `cyclonedx1.6+xml` to the `query` request parameter `format`

**2024-08-15**

**GET - `/orgs/{org_id}/audit_logs/search` - Updated**

* for the `query` request parameter `size`, default value `100.00` was added ![Badge](https://img.shields.io/badge/Breaking-yellow)

**GET - `/groups/{group_id}/audit_logs/search` - Updated**

* for the `query` request parameter `size`, default value `100.00` was added ![Badge](https://img.shields.io/badge/Breaking-yellow)

**2024-06-21 - Updated 2024-06-27**

**POST - `/orgs/{org_id}/collections` - Updated**

* the `data/attributes/name` response property's maxLength was unset from `255` for the response status `201` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the `data/attributes/name` response property's minLength was decreased from `1` to `0` for the response status `201` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the `data/attributes/name` response's property pattern `^([a-zA-Z0-9 _\-\/:.])+$` was removed for the status `201`

**GET - `/orgs/{org_id}/collections` - Updated**

* the `data/items/attributes/name` response property's maxLength was unset from `255` for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the `data/items/attributes/name` response property's minLength was decreased from `1` to `0` for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the `data/items/attributes/name` response's property pattern `^([a-zA-Z0-9 _\-\/:.])+$` was removed for the status `200`

**PATCH - `/orgs/{org_id}/collections/{collection_id}` - Updated**

* the `data/attributes/name` response property's maxLength was unset from `255` for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the `data/attributes/name` response property's minLength was decreased from `1` to `0` for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the `data/attributes/name` response's property pattern `^([a-zA-Z0-9 _\-\/:.])+$` was removed for the status `200`

**GET - `/orgs/{org_id}/collections/{collection_id}` - Updated**

* the `data/attributes/name` response property's maxLength was unset from `255` for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the `data/attributes/name` response property's minLength was decreased from `1` to `0` for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the `data/attributes/name` response's property pattern `^([a-zA-Z0-9 _\-\/:.])+$` was removed for the status `200`

**2024-06-21 - Updated 2024-06-25**

**PATCH - `/orgs/{org_id}` - Updated**

* request property `data/type` was restricted to a list of enum values ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the request property `data/attributes` became required ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the request property `data/id` became required ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the request property `data/type` became required ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new `org` enum value to the `data/type` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new `org` enum value to the request property `data/type`
* removed the pattern `^[a-z][a-z0-9]*(_[a-z][a-z0-9]*)*$` from the request property `data/type`
* the `data/type` response's property pattern `^[a-z][a-z0-9]*(_[a-z][a-z0-9]*)*$` was removed for the status `200`

**2024-06-21**

**POST - `/orgs/{org_id}/invites` - Updated**

* removed the request property `data/relationships` ![Badge](https://img.shields.io/badge/Breaking-yellow)

**2024-06-18**

**POST - `/groups/{group_id}/settings/pull_request_template` - Updated**

* removed the request property `data/attributes/branch_name` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the optional property `data/attributes/branch_name` from the response with the `201` status ![Badge](https://img.shields.io/badge/Breaking-yellow)

**GET - `/groups/{group_id}/settings/pull_request_template` - Updated**

* removed the optional property `data/attributes/branch_name` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)

**2024-06-06**

**GET - `/orgs/{org_id}/projects` - Updated**

* removed the optional property `data/items/attributes/settings/auto_dependency_upgrade/is_inherited` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)

**PATCH - `/orgs/{org_id}/projects/{project_id}` - Updated**

* removed the optional property `data/attributes/settings/auto_dependency_upgrade/is_inherited` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)

**GET - `/orgs/{org_id}/projects/{project_id}` - Updated**

* removed the optional property `data/attributes/settings/auto_dependency_upgrade/is_inherited` from the response with the `200` status ![Badge](https://img.shields.io/badge/Breaking-yellow)

**2024-05-23**

**DELETE - `/self/apps/installs/{install_id}` - Updated**

* api operation id `deleteUserAppInstallByID` removed and replaced with `deleteUserAppInstallById`

**DELETE - `/orgs/{org_id}/apps/installs/{install_id}` - Updated**

* api operation id `deleteAppOrgInstallByID` removed and replaced with `deleteAppOrgInstallById`

**DELETE - `/groups/{group_id}/apps/installs/{install_id}` - Updated**

* api operation id `deleteGroupAppInstallByID` removed and replaced with `deleteGroupAppInstallById`

**2024-05-08**

**POST - `/groups/{group_id}/settings/pull_request_template` - Added**

* Configures a group level pull request template that will be used on any org or project within that group

**GET - `/groups/{group_id}/settings/pull_request_template` - Added**

* Get your groups pull request template

**DELETE - `/groups/{group_id}/settings/pull_request_template` - Added**

* Delete your groups pull request template. This means Snyk pull requests will start to use the default template for this group.

**2024-04-29**

**GET - `/orgs/{org_id}/audit_logs/search` - Updated**

* deleted the `query` request parameter `event` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* deleted the `query` request parameter `exclude_event` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new optional `query` request parameter `events`
* added the new optional `query` request parameter `exclude_events`

**GET - `/groups/{group_id}/audit_logs/search` - Updated**

* deleted the `query` request parameter `event` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* deleted the `query` request parameter `exclude_event` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new optional `query` request parameter `events`
* added the new optional `query` request parameter `exclude_events`

**2024-04-22**

**GET - `/self` - Added**

* Retrieves information about the the user making the request.

**GET - `/orgs/{org_id}/projects` - Updated**

* added `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` to the `data/items/relationships/target` response property `oneOf` list for the response status `200`

**PATCH - `/orgs/{org_id}/projects/{project_id}` - Updated**

* added `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` to the `data/relationships/target` response property `oneOf` list for the response status `200`

**GET - `/orgs/{org_id}/projects/{project_id}` - Updated**

* added `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` to the `data/relationships/target` response property `oneOf` list for the response status `200`

**2024-02-28**

**GET - `/orgs` - Updated**

* for the `query` request parameter `name`, the maxLength was set to `100` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* for the `query` request parameter `slug`, the maxLength was set to `100` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the pattern `^[\w.-]+$` to the `query` request parameter `slug` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the optional property `data/items/attributes/access_requests_enabled` to the response with the `200` status

**PATCH - `/orgs/{org_id}` - Updated**

* added the optional property `data/attributes/access_requests_enabled` to the response with the `200` status

**GET - `/orgs/{org_id}` - Updated**

* added the optional property `data/attributes/access_requests_enabled` to the response with the `200` status

**GET - `/groups/{group_id}/orgs` - Added**

* Get a paginated list of all the organizations belonging to the group. By default, this endpoint returns the organizations in alphabetical order of their name.

**2024-02-21**

**GET - `/orgs/{org_id}/targets` - Added**

* Get a list of an organization\`s targets.

**GET - `/orgs/{org_id}/targets/{target_id}` - Added**

* Get a specified target for an organization.

**DELETE - `/orgs/{org_id}/targets/{target_id}` - Added**

* Delete the specified target.

**GET - `/orgs/{org_id}/projects` - Updated**

* removed `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` from the `data/items/relationships/target` response property `oneOf` list for the response status `200`

**PATCH - `/orgs/{org_id}/projects/{project_id}` - Updated**

* removed `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` from the `data/relationships/target` response property `oneOf` list for the response status `200`

**GET - `/orgs/{org_id}/projects/{project_id}` - Updated**

* removed `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` from the `data/relationships/target` response property `oneOf` list for the response status `200`

**2024-01-23**

**GET - `/orgs/{org_id}/issues` - Added**

* Get a list of an organization\`s issues.

**GET - `/orgs/{org_id}/issues/{issue_id}` - Added**

* Get an issue

**GET - `/groups/{group_id}/issues` - Added**

* Get a list of a group\`s issues.

**GET - `/groups/{group_id}/issues/{issue_id}` - Added**

* Get an issue

**2024-01-04**

**POST - `/custom_base_images` - Updated**

* removed `#/components/schemas/VersioningSchemaDateType` from the `data/attributes/versioning_schema` request property `oneOf` list ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed `#/components/schemas/VersioningSchemaDateType` from the `data/attributes/versioning_schema` response property `oneOf` list for the response status `201`

**PATCH - `/custom_base_images/{custombaseimage_id}` - Updated**

* removed `#/components/schemas/VersioningSchemaDateType` from the `data/attributes/versioning_schema` request property `oneOf` list ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed `#/components/schemas/VersioningSchemaDateType` from the `data/attributes/versioning_schema` response property `oneOf` list for the response status `200`

**GET - `/custom_base_images/{custombaseimage_id}` - Updated**

* removed `#/components/schemas/VersioningSchemaDateType` from the `data/attributes/versioning_schema` response property `oneOf` list for the response status `200`

**2023-11-06**

**DELETE - `/orgs/{org_id}/projects/{project_id}` - Added**

* Delete one project in the organization by project ID.

**2023-11-03**

**GET - `/self/apps/{app_id}/sessions` - Added**

* Get a list of active OAuth sessions for the app.

**DELETE - `/self/apps/{app_id}/sessions/{session_id}` - Added**

* Revoke an active user app session.

**GET - `/self/apps/installs` - Added**

* Get a list of apps installed for an user.

**DELETE - `/self/apps/installs/{install_id}` - Added**

* Revoke access for an app by install ID.

**POST - `/orgs/{org_id}/apps` - Updated**

* added the new required request property `data` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the request property `access_token_ttl_seconds` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the request property `context` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the request property `name` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the request property `redirect_uris` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the request property `scopes` ![Badge](https://img.shields.io/badge/Breaking-yellow)

**GET - `/orgs/{org_id}/apps` - Updated**

* the response property `data/items/attributes/client_id` became optional for the status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the response property `data/items/attributes/redirect_uris` became optional for the status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)

**PATCH - `/orgs/{org_id}/apps/{client_id}` - Updated**

* added the new required request property `data` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the response property `data/attributes/client_id` became optional for the status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the response property `data/attributes/redirect_uris` became optional for the status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the request property `access_token_ttl_seconds` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the request property `name` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the request property `redirect_uris` ![Badge](https://img.shields.io/badge/Breaking-yellow)

**GET - `/orgs/{org_id}/apps/{client_id}` - Updated**

* the response property `data/attributes/client_id` became optional for the status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the response property `data/attributes/redirect_uris` became optional for the status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)

**POST - `/orgs/{org_id}/apps/installs` - Added**

* Install a Snyk Apps to this organization, the Snyk App must use unattended authentication eg client credentials.

**GET - `/orgs/{org_id}/apps/installs` - Added**

* Get a list of apps installed for an organization.

**DELETE - `/orgs/{org_id}/apps/installs/{install_id}` - Added**

* Revoke app authorization for an Snyk Organization with install ID.

**POST - `/orgs/{org_id}/apps/installs/{install_id}/secrets` - Added**

* Manage client secret for non-interactive Snyk App installations.

**POST - `/orgs/{org_id}/apps/creations` - Added**

* Create a new Snyk App for an organization.

**GET - `/orgs/{org_id}/apps/creations` - Added**

* Get a list of apps created by an organization.

**PATCH - `/orgs/{org_id}/apps/creations/{app_id}` - Added**

* Update app creation attributes with App ID.

**GET - `/orgs/{org_id}/apps/creations/{app_id}` - Added**

* Get a Snyk App by its App ID.

**DELETE - `/orgs/{org_id}/apps/creations/{app_id}` - Added**

* Delete an app by its App ID.

**POST - `/orgs/{org_id}/apps/creations/{app_id}/secrets` - Added**

* Manage client secret for the Snyk App.

**POST - `/groups/{group_id}/apps/installs` - Added**

* Install a Snyk Apps to this group, the Snyk App must use unattended authentication eg client credentials.

**GET - `/groups/{group_id}/apps/installs` - Added**

* Get a list of apps installed for a group.

**DELETE - `/groups/{group_id}/apps/installs/{install_id}` - Added**

* Revoke app authorization for an Snyk Group with install ID.

**POST - `/groups/{group_id}/apps/installs/{install_id}/secrets` - Added**

* Manage client secret for non-interactive Snyk App installations.

**2023-11-02**

**GET - `/orgs/{org_id}/container_images` - Added**

* List instances of container image

**GET - `/orgs/{org_id}/container_images/{image_id}` - Added**

* Get instance of container image

**GET - `/orgs/{org_id}/container_images/{image_id}/relationships/image_target_refs` - Added**

* List instances of image target references for a container image

**2023-09-13**

**GET - `/orgs/{org_id}/projects` - Updated**

* added `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` to the `data/items/relationships/target` response property `oneOf` list for the response status `200`

**PATCH - `/orgs/{org_id}/projects/{project_id}` - Updated**

* added `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` to the `data/relationships/target` response property `oneOf` list for the response status `200`

**GET - `/orgs/{org_id}/projects/{project_id}` - Updated**

* added `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` to the `data/relationships/target` response property `oneOf` list for the response status `200`

**2023-09-12**

**GET - `/orgs/{org_id}/projects` - Updated**

* removed `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` from the `data/items/relationships/target` response property `oneOf` list for the response status `200`

**PATCH - `/orgs/{org_id}/projects/{project_id}` - Updated**

* removed `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` from the `data/relationships/target` response property `oneOf` list for the response status `200`

**GET - `/orgs/{org_id}/projects/{project_id}` - Updated**

* removed `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` from the `data/relationships/target` response property `oneOf` list for the response status `200`

**POST - `/orgs/{org_id}/collections` - Added**

* Create a collection

**GET - `/orgs/{org_id}/collections` - Added**

* Return a list of organization\`s collections with issues counts and projects count.

**PATCH - `/orgs/{org_id}/collections/{collection_id}` - Added**

* Edit a collection

**GET - `/orgs/{org_id}/collections/{collection_id}` - Added**

* Get a collection

**DELETE - `/orgs/{org_id}/collections/{collection_id}` - Added**

* Delete a collection

**POST - `/orgs/{org_id}/collections/{collection_id}/relationships/projects` - Added**

* Add projects to a collection by specifying an array of project ids

**GET - `/orgs/{org_id}/collections/{collection_id}/relationships/projects` - Added**

* Return a list of organization\`s projects that are from the specified collection.

**DELETE - `/orgs/{org_id}/collections/{collection_id}/relationships/projects` - Added**

* Remove projects from a collection by specifying an array of project ids

**2023-09-11**

**PATCH - `/orgs/{org_id}/settings/sast` - Added**

* Enable/Disable the Snyk Code settings for an org

**GET - `/orgs/{org_id}/audit_logs/search` - Added**

* Search audit logs for an Organization. Supported event types:
  * api.access
  * org.app\_bot.create
  * org.app.create
  * org.app.delete
  * org.app.edit
  * org.cloud\_config.settings.edit
  * org.collection.create
  * org.collection.delete
  * org.collection.edit
  * org.create
  * org.delete
  * org.edit
  * org.ignore\_policy.edit
  * org.integration.create
  * org.integration.delete
  * org.integration.edit
  * org.integration.settings.edit
  * org.language\_settings.edit
  * org.notification\_settings.edit
  * org.org\_source.create
  * org.org\_source.delete
  * org.org\_source.edit
  * org.policy.edit
  * org.project\_filter.create
  * org.project\_filter.delete
  * org.project.add
  * org.project.attributes.edit
  * org.project.delete
  * org.project.edit
  * org.project.fix\_pr.auto\_open
  * org.project.fix\_pr.manual\_open
  * org.project.ignore.create
  * org.project.ignore.delete
  * org.project.ignore.edit
  * org.project.monitor
  * org.project.pr\_check.edit
  * org.project.remove
  * org.project.settings.delete
  * org.project.settings.edit
  * org.project.stop\_monitor
  * org.project.tag.add
  * org.project.tag.remove
  * org.project.test
  * org.request\_access\_settings.edit
  * org.sast\_settings.edit
  * org.service\_account.create
  * org.service\_account.delete
  * org.service\_account.edit
  * org.settings.feature\_flag.edit
  * org.target.create
  * org.target.delete
  * org.user.add
  * org.user.invite
  * org.user.invite.accept
  * org.user.invite.revoke
  * org.user.invite\_link.accept
  * org.user.invite\_link.create
  * org.user.invite\_link.revoke
  * org.user.leave
  * org.user.provision.accept
  * org.user.provision.create
  * org.user.provision.delete
  * org.user.remove
  * org.user.role.create
  * org.user.role.delete
  * org.user.role.details.edit
  * org.user.role.edit
  * org.user.role.permissions.edit
  * org.webhook.add
  * org.webhook.delete
  * user.org.notification\_settings.edit

**GET - `/groups/{group_id}/audit_logs/search` - Added**

* Search audit logs for a Group. Some Organization level events are supported as well as the following Group level events:
  * api.access
  * group.cloud\_config.settings.edit
  * group.create
  * group.delete
  * group.edit
  * group.notification\_settings.edit
  * group.org.add
  * group.org.remove
  * group.policy.create
  * group.policy.delete
  * group.policy.edit
  * group.request\_access\_settings.edit
  * group.role.create
  * group.role.delete
  * group.role.edit
  * group.service\_account.create
  * group.service\_account.delete
  * group.service\_account.edit
  * group.settings.edit
  * group.settings.feature\_flag.edit
  * group.sso.add
  * group.sso.auth0\_connection.create
  * group.sso.auth0\_connection.edit
  * group.sso.create
  * group.sso.delete
  * group.sso.edit
  * group.sso.membership.sync
  * group.sso.remove
  * group.tag.create
  * group.tag.delete
  * group.user.add
  * group.user.remove
  * group.user.role.edit

**2023-09-07**

**POST - `/orgs/{org_id}/service_accounts` - Added**

* Create a service account for an organization. The service account can be used to access the Snyk API.

**GET - `/orgs/{org_id}/service_accounts` - Added**

* Get all service accounts for an organization.

**PATCH - `/orgs/{org_id}/service_accounts/{serviceaccount_id}` - Added**

* Update the name of an organization-level service account by its ID.

**GET - `/orgs/{org_id}/service_accounts/{serviceaccount_id}` - Added**

* Get an organization-level service account by its ID.

**DELETE - `/orgs/{org_id}/service_accounts/{serviceaccount_id}` - Added**

* Delete a service account in an organization.

**POST - `/orgs/{org_id}/service_accounts/{serviceaccount_id}/secrets` - Added**

* Manage the client secret of an organization service account by the service account ID.

**POST - `/groups/{group_id}/service_accounts` - Added**

* Create a service account for a group. The service account can be used to access the Snyk API.

**GET - `/groups/{group_id}/service_accounts` - Added**

* Get all service accounts for a group.

**PATCH - `/groups/{group_id}/service_accounts/{serviceaccount_id}` - Added**

* Update the name of a group\`s service account by its ID.

**GET - `/groups/{group_id}/service_accounts/{serviceaccount_id}` - Added**

* Get a group-level service account by its ID.

**DELETE - `/groups/{group_id}/service_accounts/{serviceaccount_id}` - Added**

* Permanently delete a group-level service account by its ID.

**POST - `/groups/{group_id}/service_accounts/{serviceaccount_id}/secrets` - Added**

* Manage the client secret of a group service account by the service account ID.

**2023-08-28**

**GET - `/orgs/{org_id}/projects` - Updated**

* added the new optional `query` request parameter `names_start_with`
* added the new optional `query` request parameter `target_file`
* added the new optional `query` request parameter `target_reference`
* added the new optional `query` request parameter `target_runtime`

**PATCH - `/orgs/{org_id}/projects/{project_id}` - Updated**

* deleted the `query` request parameter `user_id` ![Badge](https://img.shields.io/badge/Breaking-yellow)

**2023-08-21**

**POST - `/orgs/{org_id}/packages/issues` - Updated**

* added the optional property `meta` to the response with the `200` status

**POST - `/custom_base_images` - Added**

* In order to create a custom base image, you first need to import your base images into Snyk. You can do this through the CLI, UI, or API.

This endpoint marks an image as a custom base image. This means that the image will get added to the pool of images from which Snyk can recommend base image upgrades.

Note, after the first image in a repository gets added, a versioning schema cannot be passed in this endpoint. To update the versioning schema, the PATCH endpoint must be used.

**GET - `/custom_base_images` - Added**

* Get a list of custom base images with support for ordering and filtering. Either the org\_id or group\_id parameters must be set to authorize successfully.

**PATCH - `/custom_base_images/{custombaseimage_id}` - Added**

* Updates a custom base image\`s attributes

**GET - `/custom_base_images/{custombaseimage_id}` - Added**

* Get a custom base image

**DELETE - `/custom_base_images/{custombaseimage_id}` - Added**

* Delete a custom base image resource. (the related container project is unaffected)

**2023-06-22**

**GET - `/orgs/{org_id}/settings/sast` - Added**

* Retrieves the SAST settings for an org

**2023-05-29**

**GET - `/orgs` - Added**

* Get a paginated list of organizations you have access to.

**PATCH - `/orgs/{org_id}` - Added**

* Update the details of an organization

**GET - `/orgs/{org_id}` - Added**

* Get the full details of an organization.

**2023-04-28**

**POST - `/orgs/{org_id}/invites` - Updated**

* added the new required request property `data` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* the `data/attributes/role` response`s property type/format changed from` string`/`` to` string`/`uuid`for status`201\` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the request property `email` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* removed the request property `role` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new `org_invitation` enum value to the `data/type` response property for the response status `201` ![Badge](https://img.shields.io/badge/Breaking-yellow)

**GET - `/orgs/{org_id}/invites` - Updated**

* the `data/items/attributes/role` response`s property type/format changed from` string`/`` to` string`/`uuid`for status`200\` ![Badge](https://img.shields.io/badge/Breaking-yellow)
* added the new `org_invitation` enum value to the `data/items/type` response property for the response status `200` ![Badge](https://img.shields.io/badge/Breaking-yellow)

**2023-04-17**

**POST - `/orgs/{org_id}/packages/issues` - Added**

* This endpoint is not available to all customers. If you are interested please contact support. Query issues for a batch of packages identified by Package URL (purl). Only direct vulnerabilities are returned, transitive vulnerabilities (from dependencies) are not returned because they can vary depending on context.

**2023-03-20**

**GET - `/orgs/{org_id}/projects/{project_id}/sbom` - Added**

* This endpoint lets you retrieve the SBOM document of a software project. It supports the following formats:
* CycloneDX version 1.4 in JSON (set `format` to `cyclonedx1.4+json`).
* CycloneDX version 1.4 in XML (set `format` to `cyclonedx1.4+xml`).
* SPDX version 2.3 in JSON (set `format` to `spdx2.3+json`).

By default it will respond with an empty JSON:API response.

**2023-02-15**

**GET - `/orgs/{org_id}/projects` - Added**

* List all Projects for an Org.

**PATCH - `/orgs/{org_id}/projects/{project_id}` - Added**

* Updates one project of the organization by project ID.

**GET - `/orgs/{org_id}/projects/{project_id}` - Added**

* Get one project of the organization by project ID.

**2022-12-14**

**POST - `/orgs/{org_id}/slack_app/{bot_id}` - Added**

* Create new Slack notification default settings for a given tenant.

**GET - `/orgs/{org_id}/slack_app/{bot_id}` - Added**

* Get Slack integration default notification settings for the provided tenant ID and bot ID.

**DELETE - `/orgs/{org_id}/slack_app/{bot_id}` - Added**

* Remove the given Slack App integration

**GET - `/orgs/{org_id}/slack_app/{bot_id}/projects` - Added**

* Slack notification settings overrides for projects. These settings overrides the default settings configured for the tenant.

**POST - `/orgs/{org_id}/slack_app/{bot_id}/projects/{project_id}` - Added**

* Create Slack settings override for a project.

**PATCH - `/orgs/{org_id}/slack_app/{bot_id}/projects/{project_id}` - Added**

* Update Slack notification settings for a project.

**DELETE - `/orgs/{org_id}/slack_app/{bot_id}/projects/{project_id}` - Added**

* Remove Slack settings override for a project.

**2022-11-14**

**GET - `/orgs/{org_id}/invites` - Added**

* List pending user invitations to an organization.

**DELETE - `/orgs/{org_id}/invites/{invite_id}` - Added**

* Cancel a pending user invitations to an organization.

**2022-11-07**

**GET - `/orgs/{org_id}/slack_app/{tenant_id}/channels` - Added**

* Requires the Snyk Slack App to be set up for this org, will retrieve a list of channels the Snyk Slack App can access. Note that it is currently only possible to page forwards through this collection, no prev links will be generated and the ending\_before parameter will not function.

**GET - `/orgs/{org_id}/slack_app/{tenant_id}/channels/{channel_id}` - Added**

* Requires the Snyk Slack App to be set up for this org. It will return the Slack channel name for the provided Slack channel ID.

**2022-09-15**

**GET - `/orgs/{org_id}/packages/{purl}/issues` - Added**

* Query issues for a specific package version identified by Package URL (purl). Snyk returns only direct vulnerabilities. Transitive vulnerabilities (from dependencies) are not returned because they can vary depending on context.

**2022-06-01**

**POST - `/orgs/{org_id}/invites` - Added**

* Invite a user to an organization with a role.

**2022-03-11**

**GET - `/self/apps` - Added**

* Get a list of apps that can act on your behalf.

**DELETE - `/self/apps/{app_id}` - Added**

* Revoke access for an app by app id

**POST - `/orgs/{org_id}/apps` - Added**

* Create a new app for an organization. Deprecated, use /orgs/{org\_id}/apps/creations instead.

**GET - `/orgs/{org_id}/apps` - Added**

* Get a list of apps created by an organization. Deprecated, use /orgs/{org\_id}/apps/creations instead.

**PATCH - `/orgs/{org_id}/apps/{client_id}` - Added**

* Update app attributes. Deprecated, use /orgs/{org\_id}/apps/creations/{app\_id} instead.

**GET - `/orgs/{org_id}/apps/{client_id}` - Added**

* Get an App by client id. Deprecated, use /orgs/{org\_id}/apps/creations/{app\_id} instead.

**DELETE - `/orgs/{org_id}/apps/{client_id}` - Added**

* Delete an app by app id. Deprecated, use /orgs/{org\_id}/apps/creations/{app\_id} instead.

**POST - `/orgs/{org_id}/apps/{client_id}/secrets` - Added**

* Manage client secrets for an app. Deprecated, use /orgs/{org\_id}/apps/creations/{app\_id}/secrets instead.

**GET - `/orgs/{org_id}/app_bots` - Added**

* Get a list of app bots authorized to an organization. Deprecated, use /orgs/{org\_id}/apps/installs instead.

**DELETE - `/orgs/{org_id}/app_bots/{bot_id}` - Added**

* Revoke app bot authorization. Deprecated, use /orgs/{org\_id}/apps/installs/{install\_id} instead.

**2021-12-09**

**PATCH - `/orgs/{org_id}/settings/iac` - Added**

* Update the Infrastructure as Code Settings for an org.

**GET - `/orgs/{org_id}/settings/iac` - Added**

* Get the Infrastructure as Code Settings for an org.

**PATCH - `/groups/{group_id}/settings/iac` - Added**

* Update the Infrastructure as Code Settings for a group.

**GET - `/groups/{group_id}/settings/iac` - Added**

* Get the Infrastructure as Code Settings for a group.
