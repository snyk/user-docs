# Set up your target for testing BOLA vulnerabilities

Learn how to set up your target to be tested against Broken Object Level Authorization vulnerabilities.

Broken Object Level Authorization (BOLA) is a critical security vulnerability that occurs when unauthorized access to other users' resources is possible by changing part of a request, for example, changing a bank account number in a URL to access another user's data or resources.

You can use Snyk API & Web to test your APIs against this type of vulnerability by configuring the API target authentication and setting up two different users.

{% hint style="info" %}
BOLA testing is only available for OpenAPI and GraphQL targets.
{% endhint %}

Setting up your target to be tested against BOLA vulnerabilities involves two steps:

1. Configure additional user for authorization testing
1. Select the appropriate scan profile

This article describes these steps in detail.

## Step 1: Configure additional user for authorization testing

To configure the API Target authentication for BOLA testing:

1. Visit [Configure OpenAPI authentication](../../configure-targets/authentication/configure-openapi-authentication.md) or [Configure GraphQL authentication](../../configure-targets/authentication/configure-graphql-authentication.md) and follow the instructions for your authentication scenario.
1. When configuring authentication, check the **Add additional user for authorization testing** checkbox.
1. Configure the authentication for the second user using the same method (authentication payload or static headers/cookies). To reduce false positives, the second user should have the same level of privileges as the first user, or lower.
1. Complete the authentication configuration and ensure the toggle is set to **On**.

## Step 2: Select the appropriate scan profile

After configuring the API target authentication, choose the appropriate scan profile:

1. Access the Profile tab of that target's settings.
1. Choose either the **API normal** or **API full** scan profiles.

Once both sets of users are configured for API targets and the appropriate scan profile is selected, Snyk API & Web will test against BOLA vulnerabilities.

## Does the privilege level of the users matter?

Yes. The second user (the attacker) should not have access to the first user's resources.

The two users can have different privilege levels (for example, admin versus regular user), or they can have the same level, as long as the first user owns resources that should not be accessible to the second. That is what allows the scanner to test for BOLA vulnerabilities.

Ideally, the first user has multiple and varied private resources, so attempts by the second user to access them can effectively reveal access control flaws.

## Related content

- [Configure OpenAPI authentication](configure-openapi-authentication.md)
- [Configure GraphQL authentication](../../configure-targets/authentication/configure-graphql-authentication.md)