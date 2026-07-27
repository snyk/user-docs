---
description: How to configure authentication for Snyk API and Web targets
nav_context: agnostic
---

# Configure authentication

Configure authentication to scan protected areas of your web application or API.

Websites and applications can have restricted areas meant for authenticated users only. Configuring authentication lets Snyk API & Web access these protected areas and identify vulnerabilities within the full scope of your target.

Authentication options differ between Web targets and API targets. Configure authentication in the **Authentication** section of your target settings.

## Web target authentication

Web targets provide authentication methods designed for web application login flows.

### Login form

Login form authentication works with standard username and password forms. The scanner identifies the login form fields, fills in the credentials, and submits the form to authenticate during scans.

Use this method for straightforward login forms with username and password fields that submit via standard HTTP POST requests.

Visit [Login form](configure-login-form.md) for configuration instructions.

### Login sequence

Login sequence authentication handles complex or multi-step login flows. Record the login steps using the Snyk API & Web Sequence Recorder browser plugin. The scanner replays the recorded sequence during scans to authenticate.

Use this method for multi-step authentication flows, dynamic form elements, or JavaScript-based login processes that cannot be handled by simple form submission.

Visit [Login sequence](configure-login-sequence.md) for configuration instructions.

### Two-factor authentication

Two-factor authentication (2FA) adds an additional security layer to either login form or login sequence authentication. Configure time-based one-time passwords (TOTP) or other one-time password mechanisms.

The scanner generates or retrieves the authentication code and enters it during the login process to complete authentication.

Visit [Two-factor authentication](configure-two-factor-authentication.md) and [Automate OTP extraction](automate-otp-extraction.md) for configuration instructions.

## API target authentication

API targets use authentication methods tailored for API security testing. Configuration options depend on your API target (OpenAPI, Postman Collection, or GraphQL).

API target authentication methods include:
- API keys
- Bearer tokens
- OAuth authentication flows
- Login endpoints that return authentication tokens
- Custom scripts

Visit the following guides for detailed setup steps:
- [OpenAPI authentication](configure-openapi-authentication.md)
- [Postman authentication](configure-postman-authentication.md)
- [GraphQL authentication](configure-graphql-authentication.md)

## Authentication for Web and API targets

Some authentication features apply to both Web targets and API targets.

### Basic authentication

Basic authentication uses HTTP Basic Access Authentication, where the scanner sends credentials in the HTTP header. Configure the username and password for the scanner to include in HTTP requests.

Use this method for applications or APIs that implement the HTTP Basic Auth protocol rather than form-based or token-based authentication.

Visit [Basic authentication](configure-basic-authentication.md) for configuration instructions.

### Logout detection

Logout detection helps the scanner maintain authenticated sessions throughout the scan. Configure indicators that show when the session ends, such as logout URLs, redirects to login pages, or specific page elements that appear only when logged out.

The scanner monitors these indicators and re-authenticates if it loses the session during scanning.

Visit [Logout detection](configure-logout-detection.md) for configuration instructions.

### Mutual TLS (mTLS) authentication

Mutual TLS authentication provides enhanced security by requiring both the client and server to authenticate using digital certificates. Unlike standard TLS, which authenticates only the server, mTLS ensures bidirectional authentication.

Upload a client authentication certificate (.p12 or .pfx format) and provide the certificate password. The scanner uses the configured certificate during scans to establish secure mTLS connections with your target.

Use this method for applications or APIs that require client-side certificates for authentication.

Visit [Mutual TLS](configure-mutual-tls.md) for configuration instructions.

### Managing credentials

Credential management lets you create reusable authentication credentials and apply them across multiple targets. This simplifies configuration when you have several targets that share the same authentication credentials.

Visit [Manage credentials](manage-credentials.md) for instructions on creating and managing shared credentials.