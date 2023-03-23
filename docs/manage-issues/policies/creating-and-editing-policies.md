# Creating and editing policies

## **Create a policy**

1. Click **Add new policy** from the Policy Manager, and a modal will appear where you can create your policy
2. Set a policy name and a description to help you quickly identify a policy\
   **Note**: Policies within the same category cannot have the same name. Policies **cannot** be saved without a policy name applied
3. Select whether you’d like to apply your policy to organizations or project attributes
4. Select the desired [Organizations](assign-a-policy-to-organizations.md) or [attributes](assign-a-policy-to-projects.md)
5. Add rules to the policy. [Click here for information on adding license rules](../../scan-application-code/snyk-open-source/license-policies/setting-a-license-policy.md) or [click here for information on adding security rules](../security-policies/how-to-create-a-security-policy-and-set-rules.md)
6. Click **Submit** in the top right-hand corner.drop-down

<figure><img src="../../.gitbook/assets/screenshot_2020-05-26_at_9.47.26_am.png" alt="Submit a policy"><figcaption><p>Submit a policy</p></figcaption></figure>

## Edit a policy

1. Click on the policy name of an existing policy in the Policy Manager tab to open the modal and then make your changes
2. Adjust [Organizations](assign-a-policy-to-organizations.md), [attributes](assign-a-policy-to-projects.md) and rules as desired
3. Click **Submit** to save your changes

## **Other Actions:**

Clicking on the **...** on the right-hand side of the policies table will provide a drop down that allows you to delete a policy.

Deleting a policy cannot be undone. If you delete a policy that has Organizations assigned to it, those organizations will return to the default policy.

Duplicating a policy will copy over the rules of a policy, but not the assigned organizations or attributes. The new policy will automatically be called ‘Copy of (Policy Name)…” and can be edited.

<figure><img src="../../.gitbook/assets/screen_shot_2021-08-11_at_2.11.06_pm.png" alt="Copy a policy"><figcaption><p>Copy a policy</p></figcaption></figure>
