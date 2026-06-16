# Troubleshoot login failed with login form

In targets with authentication, Snyk API & Web must log in to reach areas reserved for authenticated users. If scans fail to log in using a login form, follow these troubleshooting steps.

## The problem

When running scans on a target with a login form, Snyk API & Web fails to log in.

## Troubleshoot the problem

Go through the following steps to identify possible causes and solutions.

### Step 1: Test the current credentials

Verify the credentials configured in target settings are still valid.

1. Navigate to the **Targets** page.
2. Click the **gear icon** for the target to access settings.
3. Open the **Authentication** tab and review the **LOGIN FORM** configuration.
4. Note the login URL and credentials.
5. Open a browser and navigate to the target's login URL.
6. Attempt to log in with the current credentials.

If the login fails, check the following causes and solutions:

| Cause                        | Solution                                                                                                                                                         |
| ---------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| The credentials are invalid. | Obtain valid login credentials and update them in target settings. Visit [Login form](../../configure-targets/configure-authentication/configure-login-form.md). |
| The credentials expired.     | Obtain new login credentials and update them in target settings. Visit [Login form](../../configure-targets/configure-authentication/configure-login-form.md).   |

### Step 2: Test the login flow

Verify the login flow is still a login form.

1. Open a browser and navigate to the target's login URL.
2. Check if the login page displays a standard login form.

If the login is not a login form, the target authentication fails.

| Cause                                                                                   | Solution                                                                                                                                                                                          |
| --------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| The login flow is not a login form but a complex login (for example, multi-step login). | Configure the target authentication to use a login sequence, which supports complex logins. Visit [Login sequence](../../configure-targets/configure-authentication/configure-login-sequence.md). |

### Step 3: Check the field names

Verify the field name values configured in target authentication.

1. Navigate to the **Targets** page.
2. Click the **gear icon** for the target to access settings.
3. Open the **Authentication** tab and review the **LOGIN FORM** configuration.
4. Note the **FIELD NAME** values for username and password fields.
5. Open a browser and navigate to the target's login URL.
6. Right-click and select **Inspect** to see the attributes of input fields on the login form.
7. Verify the **FIELD NAME** values contain a valid "id", "name", or CSS selector from the login form.

If the **FIELD NAME** values are not valid, Snyk API & Web cannot authenticate.

| Cause                                                                                                                               | Solution                                                                                                                                                                                                                          |
| ----------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| The **FIELD NAME** value does not contain a valid "id", "name", or CSS selector that identifies that input field in the login form. | Update the **FIELD NAME** value with the "id", "name", or CSS selector that uniquely identifies that input field in the login form. Visit [Login form](../../configure-targets/configure-authentication/configure-login-form.md). |

### Step 4: Test for a blocking WAF

Check if a Web Application Firewall (WAF) is blocking access to the authentication page.

1. Open a browser in incognito mode and navigate to the target's URL.
2. Navigate to the authentication page with the login form.
3. Check if a WAF blocks access.

If a WAF blocks access to the authentication page, Snyk API & Web cannot authenticate.

| Cause                                                                    | Solution                                                                                                                                                           |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| A WAF is blocking access to the authentication page with the login form. | Add Snyk API & Web IPs to the WAF's allowlist. Visit [Configure IPs in WAFs](../../start-scanning/overview-scan-access-and-connectivity/configure-ips-in-wafs.md). |

After following these steps and applying the solutions, scans should be able to log in to your target.

For more information, visit [Configure authentication](../../configure-targets/configure-authentication/).
