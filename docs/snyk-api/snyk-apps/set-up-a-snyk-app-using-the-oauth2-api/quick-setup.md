# Quick Setup

Snyk Apps uses the Authorization code with Proof Key for Code Exchange (PKCE) OAuth2 flow. The key endpoints are:

* Authorization: `/oauth2/authorize`
* Token: `/oauth2/token`

The Snyk Apps authorization flow should work with any popular OAuth2 compatible client library, providing it supports PKCE.

A basic overview:

1. Generate a code verifier and code challenge
2. Request authorization code by directing the installing user to `/oauth2/authorize` with the generated code challenge, as well as all other required OAuth2 parameters
3. The user is asked to approve the request, selecting the org/group they are installing the Snyk App to
4. The user is redirected back to the redirect URI defined by the Snyk App with a authorization code, you have to extract the authorization code from the user request
5. Exchange the authorization code for an access token and refresh token pair, using the code verifier generated in the first step.

The following pages go into more detail about each step.
