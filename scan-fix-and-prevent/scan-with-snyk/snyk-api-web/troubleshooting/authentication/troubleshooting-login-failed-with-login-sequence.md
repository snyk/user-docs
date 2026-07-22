---
description: How to troubleshoot a failed login with a login sequence in Snyk API and Web
nav_context: classic
---

# Troubleshoot login failed with login sequence

In targets with authentication, Snyk API & Web must log in to reach areas reserved for authenticated users. If scans fail to log in using a login sequence for complex authentication flows, follow these troubleshooting steps.

## The problem

When running scans on a target that uses a login sequence to complete a complex authentication flow, Snyk API & Web fails to log in.

## Troubleshoot the problem

Work through the following step to identify possible causes and solutions.

### Check the login sequence

Verify the JSON of the login sequence is still valid.

1. Navigate to the **Targets** page.
2. Click the **gear icon** for the target to access settings.
3. Open the **Authentication** tab and review the **LOGIN SEQUENCE** configuration.
4. Download the JSON of the login sequence.
5. Open a browser, navigate to the target's URL, and follow the login flow.
6. On each step of the login flow:
   1. Right-click and select **Inspect** to obtain the attributes or CSS selectors that identify the input fields.
   2. Note those attributes or CSS selectors and the values used in the input fields.
7. In the JSON of the login sequence, validate the following:
   1. The sequence follows the login flow.
   2. The attributes or CSS selectors defined to identify the input fields are correct.
   3. The values defined for use in the input fields are correct.

If the JSON does not reflect the login flow, Snyk API & Web cannot authenticate.

Check the following causes and solutions:

| Cause                                                                          | Solution                                                                                                                                                                                                 |
| ------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| The login sequence does not follow the login flow.                             | Record a new login sequence and update it in target authentication. Visit [Login sequence](../../configure-targets/configure-authentication/configure-login-sequence.md).                                |
| Input fields in the login sequence have incorrect attributes or CSS selectors. | Record a new login sequence and update it in target authentication. Visit [Login sequence](../../configure-targets/configure-authentication/configure-login-sequence.md).                                |
| Input fields in the login sequence have incorrect values.                      | Record a new login sequence with correct credential values and update it in target authentication. Visit [Login sequence](../../configure-targets/configure-authentication/configure-login-sequence.md). |

After following these steps and applying the solutions, Snyk can log in to your target.

For more information, visit [Configure authentication](../../configure-targets/configure-authentication/).
