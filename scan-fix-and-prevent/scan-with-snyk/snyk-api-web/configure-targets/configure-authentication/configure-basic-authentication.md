---
description: How to configure basic authentication for Snyk API and Web targets
---

# Configure basic authentication

Configure basic authentication to scan targets protected by HTTP Basic Access Authentication.

Basic authentication is an authentication scheme built into the HTTP protocol. When you access a protected resource, the browser prompts you with a dialog to log in.

This authentication process differs from your application's own authentication system and from form-based authentication methods. Basic authentication sends credentials in the HTTP header rather than through form submission.

## Set up basic authentication

1. From the **Targets** page, locate your target and click the **gear icon** to access the target settings.
2. Select the **Authentication** tab.
3. Scroll down to the **Basic Auth** section.
4. Enter your credentials (username and password).
5. Click **Save and enable**.

Snyk API & Web saves the credentials and enables basic authentication. Snyk uses these credentials to access and scan your protected application.

You can disable or enable basic authentication anytime using the **Off/On** toggle button, or delete the configuration using the **Delete** button.
