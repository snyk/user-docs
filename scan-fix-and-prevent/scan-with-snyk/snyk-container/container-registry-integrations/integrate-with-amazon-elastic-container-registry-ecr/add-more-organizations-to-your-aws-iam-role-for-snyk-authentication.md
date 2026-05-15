# Add more Organizations to your AWS IAM role for Snyk authentication

After creating an AWS IAM role for Snyk, you can add more Organizations to the same role for repeated use.

1. In [Snyk](https://app.snyk.io/), retrieve and copy the IDs for any additional Snyk Organizations that you want to integrate and save the IDs so you can paste them into a script in the subsequent steps.
2. In AWS, navigate to the **Trust relationships** tab for the role you want to update with additional Organizations.
3.  Click **Edit trust relationship**.

    Ensure the value of "sts:ExternalId" is enclosed with square brackets and insert the additional Organization ID inside those brackets. Use a comma ( , ) to separate between Organization ID values. For example:

    ```
    "sts:ExternalId": [
    "11111111-1111-1111-1111-111111111111",
    "22222222-2222-2222-2222-222222222222",
    "c2fa1651-601d-41gc-abe9-03691f5287d8"
    ]
    ```

    Where:

    `"11111111-1111-1111-1111-111111111111"` = a unique Org ID

    `"22222222-2222-2222-2222-222222222222"` = another unique Org ID

    `"c2fa1651-601d-41gc-abe9-03691f5287d8"` = the ID for the Organization from which you are currently setting up the integration
