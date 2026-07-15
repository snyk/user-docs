# Impact testing

Impact testing lets you understand how a proposed Rule Extension would change findings for a Snyk Code Project before you publish it. It runs your extension against the Project and returns the findings that would be added or removed. Impact testing is available both in the Snyk Web UI and through the API; this page covers the API workflow.

{% hint style="info" %}
* Impact testing is for proposed changes only. You cannot run an impact test against a Rule Extension that already exists.
* Select the smallest relevant Project for your first test so you can evaluate changes quickly. Impact test requests are rate-limited to 5 per second, 50 per minute, and 500 per hour.
{% endhint %}

For request and response schemas, see the [Impact Testing API reference](https://docs.snyk.io/developer-tools/snyk-api/reference/sastruleextensions#post-groups-group_id-sast-rule_extensions-tests).

{% stepper %}
{% step %}
**Prepare and trigger an impact test**

Before you start, gather the `project_id`, `org_id`, `fully_qualified_name`, `rule_keys`, and `extension_type` values for the Rule Extension you want to test.

Send a `POST` request to `/rest/groups/{group_id}/sast/rule_extensions/tests` with the following required attributes:

| Attribute              | Type      | Required | Description                                                                                                                      |
| ---------------------- | --------- | :------: | -------------------------------------------------------------------------------------------------------------------------------- |
| `project_id`           | UUID      |    Yes   | The Snyk Code Project to test against.                                                                                           |
| `org_id`               | UUID      |    Yes   | The Organization that owns the Project.                                                                                          |
| `fully_qualified_name` | string    |    Yes   | The FQN of the sanitizer you want to test (for example, `org.example.MySqlSanitizer`).                                           |
| `rule_keys`            | string\[] |    Yes   | One or more rule keys the extension applies to (for example, `["Sqli"]`). See [Supported rules](supported-rules.md).             |
| `extension_type`       | enum      |    Yes   | How the sanitizer operates. One of: `flows_through_sanitizer`, `if_true_sanitizer`, `if_false_sanitizer`, `any_usage_sanitizer`. |

Example request body:

{% code fullWidth="false" %}
```json
{
  "data": {
    "type": "tests",
    "attributes": {
      "project_id": "a3b8d4e1-7f1c-4a09-8b7e-2c6f5d1a0b3f",
      "org_id": "123e4567-e89b-12d3-a456-426614174000",
      "fully_qualified_name": "org.example.MySqlSanitizer",
      "rule_keys": ["Sqli"],
      "extension_type": "if_true_sanitizer"
    }
  }
}
```
{% endcode %}

A successful request returns `202 Accepted` with a `test_id` and a status of `STARTED`.
{% endstep %}

{% step %}
**Poll for completion**

Send a `GET` request to `/rest/groups/{group_id}/sast/rule_extensions/tests/{test_id}` to check the test status.

Possible status values:

* `STARTED` — the test is still running (`200 OK`)
* `COMPLETED` — results are ready; the response returns a `303` redirect to the results endpoint
* `ERRORED` — the test failed (`200 OK`); the response includes an error code and description
{% endstep %}

{% step %}
**Retrieve and review results**

After the status reaches `COMPLETED`, send a `GET` request to `/rest/groups/{group_id}/sast/rule_extensions/tests/{test_id}/results`.

The response contains:

* **`removed_findings`** — findings that no longer appear because the sanitizer addresses them.
* **`added_findings`** — findings introduced as a side effect of the extension.
* **`findings` entries** — each entry includes `count` and `changes`. Each change includes `rule_id`, `rule_name`, `location`, `severity`, `file_path`, and `priority_score`.
* **`total_count`** — the combined count of added and removed findings. A `total_count` of `0` means the extension has no effect on the tested Project.

Review the results to confirm that the removed findings match your expectations and that no unexpected findings were added. If the results are acceptable, create and publish the Rule Extension.
{% endstep %}
{% endstepper %}
