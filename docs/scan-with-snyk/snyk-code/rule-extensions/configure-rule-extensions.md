# Configure Rule Extensions

## Create rule extension

Add a custom sanitizer to an existing Snyk Code rule.

1. Log in to the Snyk Web UI and select your [Group](https://docs.snyk.io/snyk-admin/groups-and-organizations).
2. Navigate to **Rule extensions**.
3.  Select **Create sanitizer** and define functional attributes:

    1. **Sanitizer**: Provide the fully qualified name (FQN). Need help identifying your FQN? See the [FQN identification guide](identify-your-sanitizers-fqn.md).
    2. **Rule name**: Select one or more rules you want this extension to apply to. This is a Rule key as specified in the [Supported rules](supported-rules.md) page.
    3. **Sanitization type**: Select one of the four ways your function sanitizes data. For more details, see [Custom sanitizers](custom-sanitizers.md).

    <figure><img src="../../../.gitbook/assets/rule_extensions_add_sanitizer.png" alt="Add sanitizer as a rule extension in Snyk Code" width="378"><figcaption><p>Add sanitizer as a rule extension in Snyk Code</p></figcaption></figure>
4. Define the rule extension outline attributes:
   1. **Scope**: Choose if you want to apply the rule extension to the entire Group, or a subset of Organizations within the Group.
   2. **Description** (Optional): Add a short description of what the rule extension does.
   3. **Status**: Specifies how the rule extension is treated after you save it:
      * **Published**: The rule extension is live and picked up by the next test.
      * **Draft**: The extension is not included in any tests after it has been saved. Use this when collaborating with others for review purposes. After the draft stage is done, you can edit the status and publish the rule extension.
5. Select **Save** to save your Rule Extension.

## Update rule extension

You can update attributes such as sanitization type, scope, description, or status. Sanitizing functions and the rule extension name are not editable.

1. Log in to the Snyk Web UI and select your [Group](https://docs.snyk.io/snyk-admin/groups-and-organizations).
2. Navigate to **Rule extensions**.
3. Select the horizontal ellipsis and then select **Edit**.
4. Select **Save** to save your Rule Extension.

## Delete rule extension

After deletion, the rule extension is deleted and tests ‌are no longer picking it up.

1. Log in to the Snyk Web UI and select your [Group](https://docs.snyk.io/snyk-admin/groups-and-organizations).
2. Navigate to **Rule extensions**.
3. Select the horizontal ellipsis and then select **Delete**.
4. Confirm your action to delete the Rule Extension.
