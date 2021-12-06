# Getting & storing data from Snyk

In the previous sections of this tutorial, we set up our TypeScript
project, added an Express server, and configured some basic
routing. We'll be building on top of the project we created in the
previous sections, so if you haven't, it is highly recommended that
you first complete the previous portions of this tutorial before
continuing.

## Creating and registering a Snyk App

We've made some good progress with our TypeScript application so far,
but at the moment, that's all it is - a TypeScript application. To
turn it into a bonified Snyk App, we'll need to register our project
as a new App using Snyk's API

### Prerequisites

- A Snyk account with API privileges
- A [ Snyk API Token
  ](https://docs.snyk.io/features/snyk-api-info/authentication-for-api)
- The `orgId` of the Snyk Organization that will be registered as the
  App owner

#### Obtaining an `orgId`

There are two methods for retrieving an `orgId`. The first is to
simply login to your Snyk account and visit the organization settings
page for the organization you wish to retrieve the ID of. The path to
the organization settings page is:

```
https://snyk.io/org/{your-org-name}/manage/settings
```

Alternatively, you may also retrieve an organization's `orgId` via the
`https://snyk.io/api/v1/orgs` API endpoint, using your API token in
the authorization header. For details about this endpoint, view it's
documentation
[here](https://snyk.docs.apiary.io/#reference/organizations/the-snyk-organization-for-a-request/list-all-the-organizations-a-user-belongs-to).

### About Snyk Apps and the Snyk API

Snyk Apps have first class access to the API, regardless of whether
users installing the App have paid for access or not. To take
advantage of this feature, Apps must use API endpoints with the
domainhttps://api.snyk.io/rather than the conventional
https://snyk.io/api/, when accessing the API within the App.

### Registering our app with Snyk

Registration of a new Snyk App is a performed via a simple POST
request to Snyk's API. While we could certainly configure the App
we've been building throughout this tutorial to perform the request,
we'll instead make the request directly using `curl` to avoid creating
a function that can only be run a single time.

The body of the request requires the following details:

 - `name`: The name of the Snyk App
 - `redirectUris`: The accepted callback location(s) during end-user authentication
 - `scopes`: The account permissions the Snyk App will ask a user to grant

> A note on `scopes`: Once registered, a Snyk Apps `scopes` cannot currently be changed. The only recourse is deleting the Snyk App using the [Delete App](https://snykv3.docs.apiary.io/#reference/apps/single-app-management/delete-app) API endpoint and registering it again as a new Snyk App.

**At the time of this writing, Snyk Apps is still in beta. At the
moment, there is only one available scope: `apps:beta`. This scope
allows the App to test and monitor existing projects, as well as
read information about Snyk organizations, existing projects,
issues, and reports.**

With your API token and `orgId` in hand, perform the following command
in your terminal, substituting the values as necessary. For this
tutorial, use `http://localhost:3000/callback` for the `redirectUris`
value.

> Tip: You can avoid inputting your API Token and other secrets directly into your shell by adding them as export statements in a file and sourcing the file to set them as environment variables.

```bash
curl --include \
     --request POST \
     --header "Content-Type: application/vnd.api+json" \
     --header "Authorization: token <API_TOKEN>" \
     --data-binary "{
       \"name\": \"My Awesome App\",
       \"redirectUris\": [ \"http://localhost:3000/callback\" ],
       \"scopes\": [ \"apps:beta\" ]
       }" \
     'https://api.snyk.io/v3/orgs/<ORG_ID>/apps?version='
```

The response from Snyk contains two important values necessary to
complete our Snyk App's integration: `clientId` and
`clientSecret`. Store these values somewhere safe for now, we'll store
these as environment variables later, but this is the only time you
will ever see your `clientSecret` from Snyk. As a warning, **never
share your `clientSecret` publicly**. This is used to authenticate
your App with Snyk.

Now that we've registered as a Snyk App, we can start adjusting our
TypeScript project.

### User Authorization with our Snyk App

Since we're developing this Snyk App from scratch, we'll act as our
own test subjects and authorize ourselves to our app.

#### User authorization basics

User authentication for Snyk Apps is done by way of a webpage URL
containing query parameters that match up with our Snyk App's
data. We'll need to replace the query parameter values in this URL and
visit the final link in a web browser to grant account access to the
Snyk App.

```
https://app.snyk.io/oauth2/authorize?response_type=code&client_id={clientId}&redirect_uri={redirectURI}&scope={scopes}&nonce={nonce}&state={state}&version={version}
```

While some of the query parameters are somewhat obvious, let's go over
them. We're going to modify our Snyk App to to generate this URL for
our users.

 - `version`: The current version can be found in [Snyk's API documentation](https://snykoauth2.docs.apiary.io/#reference/apps/app-authorization/authorize-an-app).
 - `scopes` and `redirect_uri`: These values must match what was sent with our registration command from earlier.
 - `state`: This is used to carry any App specific state from this `/authorize` call to the callback on the `redirect_uri` (such as a user's ID). It must be verified in our callback to [prevent CSRF attacks](https://datatracker.ietf.org/doc/html/rfc6749#section-10.12).
 - `nonce`: A highly randomized string stored alongside a timestamp on the app side before calling `/authorize`, then verified on the returned access token.

Once a connection is complete, the user is redirected to the provided
redirect URI with query string parameters `code` and `state` added on,
which are necessary for the next steps of authentication.

Before we move on, let's get back into our Snyk App and address the
above requirements.

##### Updating the Snyk App to handle user authorization

Based on the above information, our Snyk App has a few new requirements:

 - A new `/callback` route
 - Nonce generation
 - A way to utilize externally stored secrets like `clientSecret`

Let's handle the `/callback` route first by creating a new controller.

```bash
mkdir -p ./src/routes/callback
touch ./src/routes/callback/callbackController.ts
```

We'll import our Controller interface definition and a few objects from Express.

```typescript
import type { Controller } from '../../interfaces/Controller';
import type { NextFunction, Request, Response } from 'express';
import { Router } from 'express';

export class CallbackController implements Controller {
  public path = '/callback';
  public router: Router = Router();

  constructor() {
    this.initRoutes();
  }

  private initRoutes() {
    this.router.get(`${this.path}`, this.success);
  }

  private success(req: Request, res: Response, next: NextFunction) {
    // return res.render('callback');
    return res.send('SUCCESS!');
  }

  private failure(req: Request, res: Response, next: NextFunction) {
    // return next(new HttpException(401, 'Authentication failed'));
  }
}
```

---

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
