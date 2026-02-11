# Convert Project-scoped ignores to asset-scoped ignores

## Conversion setup

{% hint style="info" %}
If you are new to Snyk or using Snyk Code Projects, you can skip this step, as there are no existing ignores to convert.
{% endhint %}

Conversion provides you with control over which ignores Snyk converts. For example, if you monitor multiple branches for a given repository, you can decide what ignore metadata Snyk converts and uses as the single source of truth going forward.

The following scenario assumes that you have ignored issues in your Snyk Code Project prior to enabling Snyk Code Consistent Ignores.

If you have enabled the **Snyk Code Consistent Ignores** feature and have not performed a rescan since, in order for the **Ignore across repository** button to be active, it is possible that you need to [retest the Project](../../../../scan-with-snyk/snyk-code/manage-code-vulnerabilities/#retesting-code-repository). In most cases, the **Ignore across repository** button is active, and this step is not necessary.

## Convert individual ignore

In the Snyk Web UI, navigate to a Project and open an issue card that contains an issue that was ignored before enabling this feature.

A warning states that the ignore was created through the legacy system and is not consistent across the repository. All issues ignored before enabling the feature appear with this warning, so you can determine what ignore metadata must be converted and used as the source of truth.

To convert an issue from a Project-scoped to an asset-scoped ignore, click **Ignore across repository**.

## How Snyk handles metadata

Consider the following implications for conversion:

* **Ignored by** and **Timestamp** reflect the time and user who performed the conversion, not the original ignore creator or the date of the ignore.
* **Reason**, **Description**, and **Expiration date** are all retained.

## Bulk ignore conversion

### Bulk conversion from the Snyk Web UI

{% hint style="info" %}
Conversion using Snyk Web UI is only available for Projects with 200 or fewer legacy ignores. If you need to convert ignores for Projects with more than 200 ignores, see [Convert ignores using Snyk API](convert-project-scoped-ignores-to-asset-scoped-ignores.md#convert-ignores-using-snyk-api).
{% endhint %}

You can migrate pre-existing Project-scoped ignores to asset-scoped ignores manually on the Projects page.

To convert all legacy ignores from a Project page, navigate to your **Organization** > **Projects**, and then click **Convert (number of) ignores in this Project.**

### Convert ignores using Snyk API

For Snyk Code Projects with over 200 legacy Project-scoped ignores, or in scenarios where you want more control over the values applied for ignore attributes, consider using the Snyk API instead. This allows you to programmatically convert Project-scoped ignores to asset-scoped ignores.

The following sequence of API calls allows you to build a custom script tailored to your Organization's needs for this conversion. It does not provide a prescriptive approach, but provides information on the most likely endpoints to consider.

Steps to follow:

1. [List relevant Snyk Code Projects](convert-project-scoped-ignores-to-asset-scoped-ignores.md#list-relevant-snyk-code-projects)
2. [Retrieve existing Project-scoped ignores for a Project](convert-project-scoped-ignores-to-asset-scoped-ignores.md#retrieve-existing-project-scoped-ignores-for-a-project)
3. [Map Project-scoped findings to asset-scoped findings](convert-project-scoped-ignores-to-asset-scoped-ignores.md#map-project-scoped-findings-to-asset-scoped-findings)
4. [Create the new asset-scoped ignore policy](convert-project-scoped-ignores-to-asset-scoped-ignores.md#create-the-new-asset-scoped-ignore-policy)
5. [Delete the original Project-scoped ignore](convert-project-scoped-ignores-to-asset-scoped-ignores.md#delete-the-original-project-scoped-ignore)
6. [Verify the changes](convert-project-scoped-ignores-to-asset-scoped-ignores.md#verify-the-changes)

#### List relevant Snyk Code Projects

* Goal: Identify the project\_id for the Snyk Code Projects whose ignores you want to convert.
* [API Endpoint](https://apidocs.snyk.io/?version=2024-10-15#get-/orgs/-org_id-/projects): `GET /rest/orgs/{org_id}/projects`
* Filter: Use the query parameter `&types=sast` to list only Snyk Code Projects.

{% hint style="info" %}
The response contains details about each Project, including its ID (`project_id`), origin (e.g., GitHub, CLI), and target repository or branch information. This helps you select the correct Projects, especially if the same repository is imported multiple times, such as through different SCM integrations or using the custom branch feature.
{% endhint %}

#### Retrieve existing Project-scoped ignores for a Project

* Goal: For a specific Project identified in step 1, get the details of all its current project-scoped ignores.
* [API Endpoint](../../../../snyk-api/reference/ignores-v1.md): `GET /v1/org/{org_id}/project/{project_id}/ignores`
* Response format: The response is an object where each key is a project-scoped finding ID (`snyk/org/project/finding/v1`) that has an ignore applied. The value is an array containing details about the ignore(s) for that finding.

```javascript
{
  "1ddad474-39f1-4ac4-b9c6-f2f79a65fd88": [ // <-- This is the Project-scoped finding ID
	{
  	  "reason": "The reason for the ignore",
  	  "created": "2025-04-03T12:19:23.852Z",
  	  "ignoredBy": {
    	    "id": "a1fd39ee-8253-4dab-8df9-8b5ba5bdcbc9",
    	    "name": "John Doe",
    	    "email": "john.doe@snyk.io"
  	  },
  	  "reasonType": "not-vulnerable"
	}
  ]
  // ... potentially more entries for other ignored findings in this Project
}
```

{% hint style="info" %}
Note down the Project-scoped finding ID (e.g., `1ddad474-...`), the `reason`, and the `reasonType` for each ignore you intend to convert.
{% endhint %}

#### Map Project-scoped findings to asset-scoped findings

* Goal: Find the corresponding asset-scoped finding ID (`snyk/asset/finding/v1`) for each Project-scoped finding ID identified in step 2.
* [API Endpoint](https://apidocs.snyk.io/?version=2024-10-15#get-/orgs/-org_id-/issues): `GET /rest/orgs/{org_id}/issues`
* Filters:
  * Use `&scan_item.type=project&scan_item.id={project_id}` to filter issues for the specific project.
  * Optionally, use `&status=open` if you only need to map open findings. Note that it is possible that ignores findings can also exist for resolved findings.
* Response format: The response contains a list of issues within the Project.

```javascript
{
  "data": [
	{
    	    // ... other issue details
    	    "key_asset": "33235ab8-2535-4f12-9115-2d57f41c81f1", // <-- Asset-scoped ID
    	    "key": "1ddad474-39f1-4ac4-b9c6-f2f79a65fd88",   	// <-- Project-scoped ID
    	    // ... other issue details
	}
  ]
  // ... potentially more issues
}
```

{% hint style="info" %}
For each Project-scoped finding ID (`key`) from step 2, find the matching issue in this response and extract its corresponding asset-scoped finding ID (`key_asset`). You will need this `key_asset` value to create the new ignore policy.
{% endhint %}

#### Create the new asset-scoped ignore policy

* Goal: Create a new policy that replicates the original ignore but targets the asset-scoped finding ID obtained in step 3.
* [API Endpoint:](https://apidocs.snyk.io/?version=2024-10-15#post-/orgs/-org_id-/policies) `POST /rest/orgs/{org_id}/policies`
* Payload: Construct the payload using the information gathered. Map the `reasonType` from step 2 to the `ignore_type` field, use the reason from step 2, and use the asset-scoped ID (`key_asset`) from step 3 as the value in the condition.

```javascript
{
  "data": {
	"attributes": {
  	"action": {
    	"data": {
      		"ignore_type": "not-vulnerable", // Mapped from reasonType in Step 2
      		"reason": "The reason for the ignore" // From reason in Step 2
    	}
  	},
  	"action_type": "ignore",
  	"conditions_group": {
    	"conditions": [
      	{
        	"field": "snyk/asset/finding/v1",
        	"operator": "includes",
        	"value": "33235ab8-2535-4f12-9115-2d57f41c81f1" // The key_asset value from Step 3
      	}
    	],
    		"logical_operator": "and" // Typically 'and' for single condition
  	},
  	"name": "Consistent Ignore - Converted [Optional Identifier]" // Choose a meaningful name
	},
	"type": "policy"
  }
}
```

#### Delete the original Project-scoped ignore

* Goal: Remove the legacy Project-scoped ignore now that an equivalent asset-scoped policy exists.
* [API Endpoint:](../../../../snyk-api/reference/ignores-v1.md) `DELETE /v1/org/{org_id}/project/{project_id}/ignore/{project_scoped_id}`
* Key Information: Use the project\_id from step 1 and the `{project_scoped_id}` which is the key from step 2 response (e.g., `1ddad474-39f1-4ac4-b9c6-f2f79a65fd88`).

#### Verify the changes

* Goal: Ensure the conversion was successful and the finding remains ignored under the new policy.
* Action: Retest the relevant Project in Snyk (`snyk code test` or using the Snyk Web UI).
* Check: Confirm that the finding previously covered by the Project-scoped ignore is still ignored.
