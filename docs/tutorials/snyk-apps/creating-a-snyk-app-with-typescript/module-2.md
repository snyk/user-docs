# Getting & storing data from Snyk

## Outline

1. Creating an App via the API
 - Getting your auth token and orgID
 - Requesting scopes

 <!-- One time only, single run script -->
 <!-- create-app.ts -->

2. Preparing our app to authorize users
 - authorize access with their chosen organization
 - approve requested scopes
 - link out to constructed authorization link

3. Updating our app to handle sessions and store data
<!-- passport, oauth, nonce, etc... -->
 <!-- OAuth2Strategy.ts -->
 <!-- authController.ts -->
 <!-- callbackController.ts  -->
 <!-- utils/db/*.ts -->

4. Setting up auth code and refresh token exchanges for our app
 - `access_token` & `refresh_token` info
 - mention expiry

 <!-- refreshAuthToken.ts -->
 <!-- interceptors.ts -->

5. Retrieve app org IDs

 <!-- getAppOrg.ts -->

6. Finally getting data
 - Helpers for cleaner code

 <!-- callSnykApi.ts -->
