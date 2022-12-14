# Session length

If a user logs in and then is inactive for 30 consecutive days, the user is logged out automatically. The user must log in again. This protects any account from being exposed inadvertently because the user is inactive.

You can configure the session length expiration using either the Snyk Web UI or API. You can set the length from a minimum of 5 minutes to a maximum of 30 days.

When session length expiration has been configured, tracking of session length starts within 60 seconds or when a user logs in, whichever comes first.

A user who is a member of multiple Groups, where each of those Groups has a different session length, is always logged out automatically after the shortest time configured for any of those Groups.

## Configure session length for a Snyk Group

Group admins can change the default session length of the maximum, 30 days, to any value between 30 days and a minimum of 5 minutes.

### **Prerequisites for configuring session length**

You must be an administrator of the Group to update the session length.

This feature is available to plans that support Groups. See [pricing plans](https://snyk.io/plans/) for more details.

### **Steps to configure session length**

1. Log in to your Snyk account and navigate to the Group for which you want o configure session length.
2. Navigate to **Settings** to update the Group settings.
3. In the **Session expiration** area, enter values for the session length. Valid value ranges are as follows:
   1. Days - 0-30
   2. Hours - 0-23
   3. Minutes - 0-59

<figure><img src="../../../.gitbook/assets/uuid-21093b2a-7003-b47a-cb62-2e6dd147323e-en.png" alt="Group settings, change Session expiration"><figcaption><p>Group settings, change Session expiration</p></figcaption></figure>
