# Session length

To ensure your account is safe from being inadvertently exposed through inactive users, after a user logs into their account and is then inactive for 30 consecutive days, they are automatically logged out, and must re-enter their credentials to log back in.

You can configure the session length expiration value, through the Snyk Web UI and from our API. The length value can be set from a minimum of 5 minutes, up to a maximum of 30 days.

Once configured, recording and tracking of session length and expiration initiates either when the user logs into the site or within 60 seconds of making the change - whichever comes first. Additionally, a user who is a member of multiple groups, each of which has a different session length configured, always receives the most restrictive session length - they are always automatically logged out based on the group configured with the shortest configured session length.



## Configure session length for a Snyk group

The default session length that is applied for members of a group is 30 days. Group admins can change this value to any value between 30 days and a minimum of 5 minutes.

**Prerequisites**

You must be an administrator of the group in order to update the session length.

{% hint style="info" %}
**Feature availability**\
This feature is available to plans that support groups. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

**Steps**

1. Log in to your Snyk account and navigate to the group for which you’d like to configure session length.
2. Navigate to **Settings** to update the Group settings.
3. From the **Session expiration** area, enter values for the session length. valid value ranges are as follows:
   1. Days - 0-30
   2. Hours - 0-23
   3. Minutes - 0-59

![](../../../.gitbook/assets/uuid-21093b2a-7003-b47a-cb62-2e6dd147323e-en.png)
