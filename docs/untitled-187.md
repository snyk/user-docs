# Creating and editing policies

* [ Shared Policies Overview](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360007476397-Shared-Policies-Overview/README.md)
* [ Creating and editing policies](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360007547397-Creating-and-editing-policies/README.md)
* [ Assign a policy to organizations](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360007590198-Assign-a-policy-to-organizations/README.md)
* [ Assign a policy to project attributes](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360018220857-Assign-a-policy-to-project-attributes/README.md)
* [ The .snyk file](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360007487097-The-snyk-file/README.md)

## Creating and editing policies

### Create a policy

1. Click from the Policy Manager, and a modal will appear where you can create your policy 
2. Set a policy name and a description to help you quickly identify a policy

   **Note:** Policies within the same category cannot have the same name.  
   Policies **cannot** be saved without a policy name applied

3. Select whether you’d like to apply your policy to organizations or project attributes
4. Select the desired [organizations](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360007590198/README.md) or [attributes](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360018220857/README.md)
5. Add rules to the policy. [Click here for information on adding license rules](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360007590258/README.md) or [click here for information on adding security rules](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360014473957/README.md)
6. Once you've added your rules, make sure to click **Submit** in the top right hand corner

### Edit a policy

1. Click on the policy name of an existing policy in the Policy Manager tab to open the modal and then make your changes
2. Adjust [organizations](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360007590198/README.md), [attributes](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360018220857/README.md) and/or rules as desired
3. Once you've made your changes in the Policy modal, make sure to click **Submit** in the top right hand corner to save your changes

### **Other Actions:**

Clicking on the![](https://lh5.googleusercontent.com/7j0CepqMc6Wv3grM8U17uR3IznA-Vammhg15rHgqSxGfe15MxpDtOJyP0qiOaJJylJxxV-r2aU7mVXJDOjNwwf2cKEZl9OdldZ0HkBdvKT2LV5NLGHRKD77VNos49oNxc1Had723)on the right-hand side of the policies table will provide a drop down that allows you to delete a policy.

Deleting a policy cannot be undone. If you delete a policy that has organizations assigned to it, those organizations will return to the [default policy](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360007476397/README.md).

Duplicating a policy will copy over the rules of a policy, but not the assigned organizations or attributes. The new policy will automatically be called ‘Copy of \(Policy Name\)…” and can be edited.

