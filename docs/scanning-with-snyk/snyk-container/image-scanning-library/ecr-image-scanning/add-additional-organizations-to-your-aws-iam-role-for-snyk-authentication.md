# Add additional organizations to your AWS IAM role for Snyk authentication

Once you've created an AWS IAM role for Snyk, you can add additional organizations to the same role for repeated use.

1. In [Snyk](https://app.snyk.io/), retrieve, copy the IDs for any additional Snyk organizations that you want to integrate and save them on the side. You'll need to paste them into a script in the coming steps.
2. In AWS, navigate to the Trust relationships tab for the role you would like to update with additional organizations.
3. Click **Edit trust relationship**.

   Make sure the value of "sts:ExternalId" is enclosed with square brackets and insert the additional organization ID inside those brackets. Use a comma \( , \) to separate between organization ID values. For example:

   ```text
   "sts:ExternalId": [
   "11111111-1111-1111-1111-111111111111",
   "22222222-2222-2222-2222-222222222222",
   "c2fa1651-601d-41gc-abe9-03691f5287d8"
   ]
   ```

   Where:

   `"11111111-1111-1111-1111-111111111111"` = a unique Org ID

   `"22222222-2222-2222-2222-222222222222"` = another unique Org ID

   `"c2fa1651-601d-41gc-abe9-03691f5287d8"` = the ID for the Org from which you are currently setting up the integration

