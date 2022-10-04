# Example: Setting up custom mapping for Azure AD

The following information shows how to configure custom mapping of roles for Azure AD, using the [Custom Mapping Option](custom-mapping-option.md).

In this configuration:

* Azure AD Security groups will be mapped to Snyk organizations
* Azure AD Security will be mapped to Snyk organization membership roles
* **The user role in Snyk is pre-set in each Azure AD Security Group for all members of that group.**

## Configuration

Once you’re set up with groups and users:

1.  In your Snyk App in Azure AD, open the Single Sign On Settings.

    1. Dashboard -> Enterprise Applications -> Snyk -> Single Sign On

    <img src="../../.gitbook/assets/Screen Shot 2022-06-08 at 8.22.43 AM.png" alt="" data-size="original">
2.  Edit Attributes and Claims

    <img src="../../.gitbook/assets/Screen Shot 2022-06-08 at 8.26.20 AM.png" alt="" data-size="line">
3.  Add new Claim

    <img src="../../.gitbook/assets/Screen Shot 2022-06-08 at 8.28.37 AM.png" alt="" data-size="original">
4.  Configure Attribute

    <img src="../../.gitbook/assets/Screen Shot 2022-06-08 at 8.32.50 AM.png" alt="" data-size="original">

    1. **Name**: roles
    2. Expand the Claim conditions section:
       1. For each unique value (unique combinations of security groups--as a reminder, each security group reflects a unique combination of organization membership and user role), you will need to create a new condition. **Order is important**. If you have more than one condition with the same group/s included in 'scoped groups' the conditions will be evaluated top to bottom and the last value that includes the group/s will be utilized. It is suggested to have the conditions in increasing order of scoped groups for this reason.
       2. Set User type to Members\
          ![](<../../.gitbook/assets/Screen Shot 2022-06-08 at 9.19.38 AM.png>)
       3. The scoped groups should be the security groups that you're assigning one or more org membership/user role combination to.\
          ![](<../../.gitbook/assets/select groups.png>)
       4. Select Attribute as the Source
       5.  Set the Value to the Snyk org + user role slugs in the following format:\
           snyk-orgslug-role

           1. For more than one, separate by comma
           2. There should NOT be any spaces or capital letter(s) in the org + user role slugs.
           3. Do not include double quotes as Azure AD automatically adds them.

           <img src="../../.gitbook/assets/Screen Shot 2022-06-08 at 9.20.22 AM.png" alt="" data-size="original">
       6. Repeat Steps 2-5 for each condition
