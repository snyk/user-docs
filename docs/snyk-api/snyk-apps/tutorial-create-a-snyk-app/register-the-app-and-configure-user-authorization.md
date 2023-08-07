# Register the App and configure user authorization

In the previous sections of this tutorial, we set up our TypeScript project, added an Express server, and configured some basic routing. We'll be building on top of the project we created in the previous sections. It is highly recommended that you complete the previous portions of this tutorial before continuing if you have not done so already.

## Creating and registering a Snyk App

We've made some good progress with our TypeScript application so far, but at the moment, that's all it is - a TypeScript application. To turn it into a bonafide Snyk App, we'll need to register our project as a new App using Snyk's API

### Prerequisites

* A Snyk account with API privileges
* A [Snyk API Token](https://docs.snyk.io/features/snyk-api-info/authentication-for-api)
* The `orgid` of the Snyk Organization that will be registered as the App owner

### Obtaining an `orgid`

There are two methods for retrieving an `orgid`. The first is to log in to your Snyk account and visit the organization settings page of the organization for which you wish to retrieve the ID. The path to the organization settings page is:

```
https://snyk.io/org/{your-org-name}/manage/settings
```

Alternatively, you may retrieve an organization's `orgid` via the `https://snyk.io/api/v1/orgs` API endpoint, using your API token in the authorization header. For details about this endpoint, view its [documentation](https://snyk.docs.apiary.io/#reference/organizations/the-snyk-organization-for-a-request/list-all-the-organizations-a-user-belongs-to).

### About Snyk Apps and the Snyk API

Snyk Apps have first class access to the API, regardless of whether users installing the App have paid for access or not. To take advantage of this feature, Apps must use API endpoints with the domain https://api.snyk.io/rather than the conventional https://snyk.io/api/, when accessing the API within the App.

### Registering our app with Snyk

Registration of a new Snyk App is a performed via a simple POST request to Snyk's API. While we could certainly configure the App we've been building throughout this tutorial to perform the request, we'll instead make the request directly using `curl` to avoid creating a function that can only be run a single time.

The body of the request requires the following details:

* `name`: The name of the Snyk App
* `redirectUris`: The accepted callback location(s) during end-user authentication
* `scopes`: The account permissions the Snyk App will ask a user to grant

A note on scopes: Once registered, a Snyk Apps scopes cannot currently be changed. The only recourse is deleting the Snyk App using the [Delete App](https://snykv3.docs.apiary.io/#reference/apps/single-app-management/delete-app) API endpoint and registering the app again as a new Snyk App.

At the time of this writing, **Snyk Apps is still in beta**. At the moment, t**here is only one available scope: `apps:beta`**. This scope **allows** the App to test and monitor existing projects, as well as read information about Snyk organizations, existing projects, issues, and reports.

One of the **limitations of the Snyk Apps beta** is that **a Snyk App may only be authorized by users who have administrator access to the organization to which the Snyk App is registered**.

With your API token and `orgid` in hand, perform the following command in your terminal, substituting the values as necessary. For this tutorial, use `http://localhost:3000/callback` for the `redirectUris` value.

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
     'https://api.snyk.io/rest/orgs/<ORG_ID>/apps?version='
```

The response from Snyk contains two important values necessary to complete our Snyk App's integration: `clientId` and `clientSecret`. Store these values somewhere safe. This is the only time you will see your `clientSecret` from Snyk. As a warning, **never share your `clientSecret` publicly**. This is used to authenticate your App with Snyk.

Now that we've registered the app as a Snyk App, we can start adjusting our TypeScript project to allow users to authorize it.

## User authorization with a Snyk App

User authentication for Snyk Apps is done by way of a webpage URL containing query parameters that match up with our Snyk App's data. We'll need to replace the query parameter values in this URL and send users to the final link in a web browser. From there they can grant account access to the Snyk App.

Once access has been provisioned, the user will be kicked back to our app's registered `callbackURL`, which we defined as `http://localhost:3000/callback`.

Essentially, our app needs to generate a link like the following and then send the user to it when it's time to authorize:

```
https://app.snyk.io/oauth2/authorize?response_type=code&client_id={clientId}&redirect_uri={redirectURI}&scope={scopes}&nonce={nonce}&state={state}&version={version}
```

Though some of the query parameters may be somewhat obvious, we will go over them. We're going to modify our Snyk App to to generate this URL for our users.

* `version`: The current version can be found in [Snyk's OAuth API documentation](https://snykoauth2.docs.apiary.io/#reference/apps/app-authorization/authorize-an-app).
* `scopes` and `redirect_uri`: These values must match what was sent with our registration command from [Registering our app with Snyk](register-the-app-and-configure-user-authorization.md#registering-our-app-with-snyk).
* `state`: This is used to carry any App-specific state from this `/authorize` call to the callback on the `redirect_uri` (such as a user's ID). It must be verified in our callback to [prevent CSRF attacks](https://datatracker.ietf.org/doc/html/rfc6749#section-10.12).
* `nonce`: A highly randomized string stored alongside a timestamp on the app side before calling `/authorize`, then verified on the returned access token. This is intended for the Snyk App to have confidence in the security of the response.

Once a connection is complete, the user is redirected to the provided redirect URI (our `/callback` route in this case) with query string parameters `code` and `state` added on, which are necessary for the next steps of authorization.

That next step involves taking the authorization code received as query parameter in the response of the previous step, and turning it into an _access token_. To do this, a Snyk App makes a POST request to the token endpoint: `https://api.snyk.io/oauth2/token`. That POST request needs some specific data in its request body, including the _authorization code_, _client id_, _client secret_, and so on.

When successful, that POST request's response contains everything a Snyk App needs to communicate with Snyk on behalf of the authorizing user, namely, a **refresh\_token** and an **access token**.

The access token gets used for future API calls and has a much shorter expiry than the refresh token. The refresh token can be used only one time to get a new access token when it expires. In other words, **the refresh\_token will no longer be usable if its own expiry time passes or if it is used to refresh the access\_token.** Ultimately, this means that our Snyk App will need to make frequent API calls to perform **refresh\_token** exchanges.

{% hint style="info" %}
Both of these tokens should be encrypted before storing them.
{% endhint %}

## Updating our Snyk App to handle user authorization

Based on the preceding information, our Snyk App has some new requirements. We will need to do a few things within our TypeScript app to successfully authorize a Snyk user account with our Snyk app:

1. Send API requests and process responses.
2. Keep track of data, like token expiry.
3. Encrypt and decrypt secret data.
4. Turn authorization _codes_ into authorization _tokens._
5. Refresh those authorization tokens.
6. Handle errors and inform users of authorization success or failure.

From here on, we'll be doing a lot of refactoring in our Snyk App and we'll be jumping into a number of different files. To help make the process easier to follow, this tutorial uses the convention of adding a commented filepath to the first line of code snippets, describing where they belong. In your own code; these comments aren't necessary.

We'll also add a number of new packages to help address our new requirements. For convenience, go ahead and run the following in the root of your project:

```bash
npm install --save passport \
    passport-oauth2 \
    @snyk/passport-snyk-oauth2 \
    @types/passport \
    @types/uuid \
    express-session \
    axios \
    uuid \
    ejs \
    jwt-decode \
    cryptr \
    "github:dankreiger/lowdb#chore/esm-cjs-hybrid-WITH-LIB" # This allows lowdb to be used with commonjs modules
    luxon;

npm install --save-dev @types/cryptr \
    @types/ejs \
    @types/express-session \
    @types/luxon \
    @types/passport \
    @types/passport-oauth2 \
    @types/uuid
```

### Store data and config in the Snyk App

#### Application configuration

Application config (for example, client secrets, api tokens, other config and so on) should generally be stored securely and kept outside of the App itself. However, for brevity this tutorial simply adds the configuration info as exportable constants in the `App.ts` file and leaves the actual implementation details to you. These are values that the Snyk App references in many different places.

```typescript
// ./src/app.ts

...

import { v4 as uuid4 } from "uuid";

...

export const APP_BASE = "https://app.snyk.io";
export const API_BASE = "https://api.snyk.io";
export const CLIENT_ID = "[replace with your client id]";
export const CLIENT_SECRET = "[replace with your client secret]";
export const ENCRYPTION_SECRET = uuid4();
export const REDIRECT_URI = "https://localhost:3000/callback";
export const TOKEN_URL = "/oauth2/token";
export const AUTHORIZATION_URL = "/oauth2/authorize";
export const SCOPE = "apps:beta";
export const NONCE = uuid4();
export const STATE = true;
...
```

#### Storing data

One of the things we'll want to do is capture some information about the users that authorize our Snyk App. Again, a true implementation is left up to you. For the purposes of this tutorial, we'll use the excellent `lowdb`, a small local JSON database with low overhead.

We'll first create a new middleware function in `app.ts` to initialize a `lowdb` database file at `./db/` and tell the `App` constructor to call it.

```typescript
// ./src/app.ts

...
import * as path from "path";
import * as fs from "fs";

...

constructor(controllers: Controller[], port: number) {
  this.app = express();
  this.initDatabaseFile();
  this.initRoutes(controllers);
  this.server = this.listen(3000);
}

...

private initDatabaseFile() {
  try {
    const dbFolder = path.join(__dirname, "../db");
    dbPath = path.join(dbFolder, "db.json");
    console.log(`Using db: ${dbPath}`);
    if (!fs.existsSync(dbPath)) {
      if (!fs.existsSync(dbFolder)) {
        fs.mkdirSync(dbFolder);
      }
    }
  } catch (error) {
    console.error(error);
  }
}

...

export let dbPath: string;
export default App;
```

With the database initialization handled, we'll create some new helper methods to make reading, writing, and updating database entries simpler. Because this is a TypeScript project, we'll be creating interfaces or types around the data structures we'll be storing, so we'll create two files: `./src/interfaces/DB.ts` and `./src/util/DB.ts`:

```bash
touch ./src/interfaces/DB.ts
mkdir -p ./src/util
touch ./src/util/DB.ts
```

Populate the interface file with an interface describing each piece of the authorization data we'll be storing, and a wrapping interface we can apply to the entire database:

```typescript
// ./src/interfaces/DB.ts

export interface DB {
  installs: AuthData[];
}

export interface AuthData {
  date: Date;
  userId: string;
  orgId: string;
  access_token: string;
  expires_in: 3600;
  scope: string;
  token_type: string;
  nonce: string;
  refresh_token: string;
}
```

In this tutorial, we'll only need to perform three basic interactions with our database: read, write, and update.

Within the file we created in `./src/util`, create a function for each. Our read function will return a Promise with the database contents; the write function will take an object that matches the `AuthData` interface we just described; and the update function will attempt to rewrite an entry, returning a boolean denoting success or failure.

```typescript
// ./src/util/DB.ts

import { dbPath } from "../app";
import { AuthData, DB } from "../interfaces/DB";
import { Low, JSONFile } from "lowdb";

export async function readFromDb(): Promise<DB> {
  const adapter = new JSONFile<DB>(dbPath);
  const db = new Low<DB>(adapter);
  await db.read();
  // Return existing data or create a new DB
  return db.data ?? buildNewDb();
}

function buildNewDb(): DB {
  return { installs: [] };
}

export async function writeToDb(data: AuthData): Promise<void> {
  const existingData = await readFromDb();
  existingData.installs.push(data);
  // Creates a new DB if one doesn't already exists
  const adapter = new JSONFile(dbPath);
  const db = new Low(adapter);
  db.data = existingData;
  return db.write();
}

export async function updateDb(
  oldData: AuthData,
  newData: AuthData
): Promise<boolean> {
  const adapter = new JSONFile<DB>(dbPath);
  const db = new Low<DB>(adapter);
  await db.read();
  if (db.data == null) {
    return false;
  }
  // After reading check if data exists in the database
  const installs = db.data?.installs || [];

  const index = installs.findIndex((install) => install.date === oldData.date);
  if (index === -1) return false;
  installs[index] = newData;
  // Replace the existing install with new one
  db.data.installs = installs;
  await db.write();
  return true;
}
```

### Prepare for API calls

Earlier, we installed the popular `axios` package to handle API calls. We know that we'll need to make some repetitive calls to the same API, so we abstract some helper functions to make our code easily re-usable across the project. Create an `APIHelpers.ts` file in the `util` directory.

Before we fill that out, take note that while we are consistently hitting Snyk's API, we'll likely need to make requests against multiple versions of the API, depending on the endpoint's status in the migration from Snyk API v1 to Snyk REST API. One way we can handle this is by defining a TypeScript Enum and within our functions, swapping any necessary query parameters by comparing an argument to the enum's possible values.

Add the following content to a new file, or to `APIHelpers.ts` if you prefer; just make sure to export it for later use.

```typescript
// ./interfaces/API.ts
export const enum APIVersion {
  V1 = "v1",
  REST = "rest",
}
```

We start by adding a single function to simplify our Apps' calls to the Snyk API. The function takes a `tokenType` (either _bearer_ or _token_), the `token` itself, and an `APIVersion` (conveniently corresponding to the enum we just defined).

```typescript
// ./src/util/APIHelpers.ts

import qs from "qs";
import axios, { AxiosInstance } from "axios";
import { API_BASE, CLIENT_ID, CLIENT_SECRET, TOKEN_URL } from "../app";
import { AuthData } from "../interfaces/DB";

export function callSnykApi(tokenType: string, token: string, version: APIVersion): AxiosInstance {
  const contentType = version === APIVersion.V1 ? "application/json": "application/vnd.api+json";

  const axiosInstance = axios.create({
    baseURL: `${API_BASE}/${version}`,
    headers: {
      "Content-Type": contentType,
      Authorization: `${tokenType} ${token}`,
    },
  });

  return axiosInstance;
}
```

Because this function is an `AxiosInstance`, we can easily talk to the API's different endpoints by calling `.get()`, `.post()`, or any other methods usually available to such an object.

We will see it in action by defining a second async function to retrieve our Snyk Apps' Organization ID:

```typescript
// ./src/util/APIHelpers.ts

...

export async function getAppOrgID(tokenType: string, accessToken: string): Promise<{ orgId: string }> {
  try {
    const clientId = CLIENT_ID;
    const result = await callSnykApi(tokenType, accessToken, APIVersion.REST)({
      method: "GET",
      url: `/apps/${clientId}/orgs?version=2021-08-11~experimental`,
    });
    // Fetch the first org
    const org = result.data.data[0];
    return {
      orgId: org.id,
    };
  } catch (error) {
    console.error("Error fetching org info: " + error);
    throw error;
  }
}
```

### Make encrypt / decrypt easy

It's a good idea to encrypt the data we'll be pulling out of the API. We will define a small class for doing so. The class has two members:

1. `secret`: The key used to encrypt data
2. `cryptr`: Instance of the Cryptr library

```typescript
// ./src/util/encrypt-decrypt.ts

import Cryptr from "cryptr";

export class EncryptDecrypt {
  private secret: string;
  private cryptr: Cryptr;

  constructor(secret: string) {
    // Uses the passed secret
    this.secret = secret as string;
    // Initialize the Cryptr instance with the secret
    this.cryptr = new Cryptr(this.secret);
  }

  /**
   * Function used to encrypt data
   * @param {String} message to be encrypted
   * @returns {String} encrypter message
   */
  public encryptString(message: string): string {
    return this.cryptr.encrypt(message);
  }
  /**
   * Function used to decrypt data
   * @param  {String} encryptedString to be decrypted
   * @returns {String} decrypted string
   */
  public decryptString(encryptedString: string): string {
    return this.cryptr.decrypt(encryptedString);
  }
}
```

### Configure Passport.js and the Snyk-OAuth2 strategy

We've laid the groundwork; now it's time to start actually doing things.

As discussed in a previous section, our app needs to send would-be authorizers to a specific token URL. We'll add an `/auth` route in our Snyk App and add some authentication middleware to Express. For this, we'll use the excellent [passportjs](https://www.passportjs.org), the [passport-oauth2](https://https/www.passportjs.org/packages/passport-oauth2) authentication strategy, along with Snyk's [@snyk/passport-snyk-oauth2](https://www.npmjs.com/package/@snyk/passport-snyk-oauth2). `passport` and its friends handle a large portion of what would otherwise be a lengthy and complicated authentication process.

Because `passport` takes its encapsulation philosophy seriously, we'll need to handle everything else about the auth process. We need to set up an instance of the passport strategy we'll be using. We'll put our database helpers from earlier to use here as well, adding an entry into our database when we receive successful authorization.

It's worth taking the time to go over this file and make sure you've understood everything its doing.

```typescript
// ./util/OAuth2Strategy.ts

import type { Request } from "express";
import axios, { AxiosResponse, AxiosInstance } from "axios";
import OAuth2Strategy, { VerifyCallback } from "passport-oauth2";
import SnykOAuth2Strategy, { ProfileFunc } from "@snyk/passport-snyk-oauth2";
import { v4 as uuid4 } from "uuid";
import jwt_decode from "jwt-decode";
import { EncryptDecrypt } from "../util/encrypt-decrypt";
import { writeToDb } from "../util/db";
import { AuthData } from "../interfaces/db";
import { getAppOrgID } from "../util/APIHelpers";

// This just wraps up the tutorial's app config to avoid writing
// each config variable.
// You'd likely want to parse environment variables or something.
import * as config from "../app";

// Set up a new type definition for the parameters we'll be sending with our auth.
type Params = {
  expires_in: number;
  scope: string;
  token_type: string;
};

// There is more data here but we only care about the nonce
type JWT = { nonce: string };

export function getOAuth2(): SnykOAuth2Strategy {
  const nonce = uuid4();

  // User can pass their own implementation of fetching the profile
  // by providing the profileFunc implementation. Snyk OAuth2 strategy
  // will call this function to fetch the profile associated with request
  const profileFunc: ProfileFunc = function(accessToken: string) {
    return axios.get("https://api.snyk.io/v1/user/me", {
      headers: {
        "Content-Type": "application/json; charset=utf-8",
        Authorization: `bearer ${accessToken}`,
      },
    });
  };

  // Note*: the value of version being manually added
  return new SnykOAuth2Strategy(
    {
      authorizationURL: `${config.APP_BASE}${config.AUTHORIZATION_URL}?version=2021-08-11~experimental`,
      tokenURL: `${config.API_BASE}${config.TOKEN_URL}`,
      clientID: config.CLIENT_ID,
      clientSecret: config.CLIENT_SECRET,
      callbackURL: "http://localhost:3000/callback",
      scope: config.SCOPE,
      scopeSeparator: " ",
      state: true,
      passReqToCallback: true,
      nonce,
      profileFunc,
    },
    async function (
      req: Request,
      access_token: string,
      refresh_token: string,
      params: Params,
      profile: AxiosResponse,
      done: VerifyCallback
    ) {
      try {
        // Notify passport that all work, like the storing
        // of data in the DB, has been completed
        const userId = profile.data.id;
        const decoded: JWT = jwt_decode(access_token);
        if (nonce !== decoded.nonce)
          throw new Error("Nonce values do not match");
        const { expires_in, scope, token_type } = params;

        const { orgId } = await getAppOrgID(token_type, access_token);
        const ed = new EncryptDecrypt(config.ENCRYPTION_SECRET as string);

        await writeToDb({
          date: new Date(),
          userId,
          orgId,
          access_token: ed.encryptString(access_token),
          expires_in,
          scope,
          token_type,
          refresh_token: ed.encryptString(refresh_token),
          nonce,
        } as AuthData);
      } catch (error) {
        return done(error as Error, false);
      }
      return done(null, { nonce });
    }
  );
}
```

### Update Express middleware

With our passport strategy implemented, modify `app.ts` to set up the `passport` middleware as shown in the next code block. Rather than calling it directly, we'll create a function called `initGlobalMiddlewares()` allowing us to set up a few other middlewares at the same time:

* `express.json()`: Express middleware for handling JSON requests
* `express.urlencoded()`: Express middleware to handle URL encoded calls
* `expressSession`: The `express-session` middleware package which is extended by `passport`
* `setupPassport`: To nitialize `passport` setup

```typescript
// ./src/app.ts

...

import passport from "passport";
import { getOAuth2 } from "./util/OAuth2Strategy";

...

constructor(controllers: Controller[], port: number) {
  ...
  this.initDatabaseFile();
  this.initGlobalMiddlewares();
  this.initRoutes(controllers);
  ...
}

...

private setupPassport() {
  passport.use(getOAuth2());
  this.app.use(passport.initialize());
  this.app.use(passport.session());
  passport.serializeUser((user: Express.User, done) => {
    done(null, user);
  });
  passport.deserializeUser((user: Express.User, done) => {
    done(null, user);
  });
}

private initGlobalMiddlewares() {
  this.app.use(express.json());
  this.app.use(express.urlencoded({ extended: true }));
  this.app.use(
    expressSession({
      secret: uuid4(),
      resave: false,
      saveUninitialized: true,
    })
  );
  this.setupPassport();
}

...
```

### Handle the authorization and callback routes

The authorization and callback controllers are comparatively simple. Create two new controller files:

```bash
mkdir -p ./src/routes/auth;
mkdir -p ./src/routes/callback;
touch ./src/routes/auth/authController.ts
touch ./src/routes/callback/callbackController.ts
```

The `AuthController` handles authentication of the App via the previously described authorization flow. This is the third step of `passport` setup. Every controller class implements the controller interface which has two members: the path and the router.

This controller handles the `/auth` route, which is what we'll use to send users (via `passport`) to the Snyk website for authorization approval.

```typescript
// ./src/routes/auth/authController.ts

import type { Controller } from "../../interfaces/Controller";
import { Router } from "express";
import passport from "passport";

class AuthController implements Controller {
  // The base URL path for this controller
  public path = "/auth";
  // Express router for this controller
  public router = Router();

  constructor() {
    this.initRoutes();
  }

  /**
   * The /auth route is called to authenticate the App
   * via Snyk using passportjs authenticate method
   */
  private initRoutes() {
    this.router.get(`${this.path}`, passport.authenticate("snyk-oauth2"));
  }
}

export default AuthController;
```

Once a user approves authorization to our Snyk App via the Snyk website, the user is kicked back to our callback URI, `/callback`. We'll handle this route similarly, invoking passport again. This is the final step of user authorization.

The `CallbackController` accepts requests on `/callback`, but also creates two sub-routes, `/callback/success` and `/callback/failure`, to handle the different possible outcomes it might receive from Snyk.

```typescript
// ./src/routes/callback/callbackController.ts

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
    // Path to handle the result of authentication flow or the callback/redirect_uri
    this.router.get(`${this.path}`, this.passportAuthenticate());
    // Path to handle success, same as what we pass to passport
    this.router.get(`${this.path}/success`, this.success);
    // Path to handle failure, same as what we pass to passport
    this.router.get(`${this.path}/failure`, this.failure);
  }
  
  private passportAuthenticate() {
    return passport.authenticate('snyk-oauth2', {
      successRedirect: '/callback/success',
      failureRedirect: '/callback/failure',
    });
  }

  private success(req: Request, res: Response, next: NextFunction) {
    // return res.render('callback');
    return res.send('SUCCESS!');
  }

  private failure(req: Request, res: Response, next: NextFunction) {
    // return next(new HttpException(401, 'Authentication failed'));
  }
}

export default CallbackController;
```

Before we're done, we need to make sure we add a reference to our new controllers in our `index.ts`.

```typescript
// ./src/index.ts

import IndexController from "./routes/index/indexController";
import AuthController from "./routes/auth/authController";
import CallbackController from "./routes/callback/callbackController";
import App from "./app";

new App([
   new IndexController(),
   new AuthController(),
   new CallbackController()],
  3000
);
```

### Refresh token management

If we build and run our Snyk App at this point, hitting the `/auth` route will successfully jump us out to the Snyk authorization portal and, provided we confirm the authorization, we'll get kicked back to our local app's callback route at `/callback`. If we had a very simplistic, one-off use case, we could end things here. But there's one more piece of the puzzle we should figure out if we're going to keep our user's authorization fresh; that is token expiry.

If you ran the app to test things, take a look at database entries. If you've been following along, you should see something like this:

```json
{
  "installs": [
    {
      "date": "2021-12-28T19:15:02.043Z",
      "userId": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
      "orgId": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
      "access_token": "xxxxxxxxxxxxxxxxxxxxxxxxxx",
      "expires_in": 3599,
      "scope": "apps:beta",
      "token_type": "bearer",
      "refresh_token": "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy",
      "nonce": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    },
  ]
}
```

That `expires_in` value will continue to count down until 0. If it does, the user will need to re-authorize.

To keep our access token from going stale, we need to make a POST request using our `refresh_token` to get an updated `access_token`. See [Set up the refresh token exchange](https://docs.snyk.io/features/integrations/snyk-apps/getting-started-with-snyk-apps/set-up-the-refresh-token-exchange).

We can automate this process in our Snyk App by utilizing [Axios interceptors](https://axios-http.com/docs/interceptors) to _intercept_ the requests we make and ensure we have an up-to-date `access_token`.

Create the file `./src/util/interceptors.ts`, importing all the packages, classes, and so on that we'll need at the top:

```typescript
// ./src/util/interceptors.ts

import type { AxiosRequestConfig } from "axios";
import { AuthData } from "../interfaces/DB";
import { DateTime } from "luxon";
import { readFromDb, updateDb } from "./DB";
import { mostRecent } from "../routes/projects/projectsController";
import { EncryptDecrypt } from "./encrypt-decrypt";
import { refreshAuthToken } from "../util/APIHelpers";
import { ENCRYPTION_SECRET } from "../app";
import axios from "axios";
```

We'll add a total of three interceptors.

The first, `refreshTokenReqInterceptor` will refresh the `auth_token` using the `refresh_token` when the `auth_token` expires. It takes an _AxiosRequestConfig_ request as an argument that can be used in the interceptor for additional checks.

```typescript
// ./src/util/interceptors.ts

...

export async function refreshTokenReqInterceptor(request: AxiosRequestConfig): Promise<AxiosRequestConfig> {
  // Read the latest data(auth token, refresh token and expiry)
  const db = await readFromDb();
  const data = mostRecent(db.installs);
  // If no data then continue with the request
  if (!data) return request;
  // Data used to calculate the expiry
  const expiresIn = data.expires_in;
  const createdDate = data.date;
  // Used npm library luxon to parse the date and calculate expiry
  const parsedCreateDate = DateTime.fromISO(createdDate.toString());
  const expirationDate = parsedCreateDate.plus({ seconds: expiresIn });
  // Check if expired
  if (expirationDate < DateTime.now()) {
    await refreshAndUpdateDb(data);
  }
  return request;
}
```

`refreshTokenRespInterceptor` will be used during request responses . Refresh/retry the token only when the response being received is a 401 Unauthorized; this is what Passport returns when things go awry.

```typescript
// ./src/util/interceptors.ts

...

export async function refreshTokenRespInterceptor(error: AxiosError): Promise<AxiosError> {
  const status = error.response ? error.response.status: null;

  // Only refresh & retry the token on 401 Unauthorized, in case the access-token is
  //  invalidated before it expires, such as the signing key being rotated in an emergency.
  if (status === 401) {
    // Read the latest data(auth token, refresh token and expiry)
    const db = await readFromDb();
    const data = mostRecent(db.installs);
    // If no data then fail the retry
    if (!data) return Promise.reject(error);

    const newAccessToken = await refreshAndUpdateDb(data);

    // Use the new access token to retry the failed request
    error.config.headers['Authorization'] = `${data.token_type} ${newAccessToken}`;
    return axios.request(error.config);
  }

  return Promise.reject(error);
}
```

Lastly, `refreshAndUpdateDb` refreshes the access token for a given database record and updates the database again before returning the newly refreshed token.

```typescript
// ./src/util/interceptors.ts

...

async function refreshAndUpdateDb(data: AuthData): Promise<string> {
  // Create a instance for encryption and decryption
  const eD = new EncryptDecrypt(process.env[Envars.EncryptionSecret] as string);
  // Make request to refresh token
  const { access_token, expires_in, refresh_token, scope, token_type } = await refreshAuthToken(
    eD.decryptString(data.refresh_token),
  );
  // Update the access and refresh token with the newly fetched access and refresh token
  // along with the expiry and other required info
  await updateDb(data, {
    ...data,
    access_token: eD.encryptString(access_token),
    expires_in,
    refresh_token: eD.encryptString(refresh_token),
    token_type,
    scope,
    date: new Date(),
  });

  return access_token;
}
```

With our interceptors defined, the only thing we need to do is update our `callSnykApi` function to utilize them. Interceptors are methods of the `axiosInstance` object, so we'll add them after the `axios.create()` call and before the function's `return`.

```typescript
// ./src/util/APIHelpers.ts
...

import {
  refreshTokenReqInterceptor,
  refreshTokenRespInterceptor,
} from "./interceptors";

...

export function callSnykApi(tokenType: string, token: string, version: APIVersion): AxiosInstance {

  ...

  axiosInstance.interceptors.request.use(
    refreshTokenReqInterceptor,
    Promise.reject
  );
  axiosInstance.interceptors.response.use(
    (response) => response,
    refreshTokenRespInterceptor
  );

  ...

}

...
```

## Wrap-up

If you've made it this far, congratulations! You've learned how to register a Snyk App with Snyk, configure the authorization flow, keep the `auth_token` from getting stale, and set up a great starting point using TypeScript.

In the next module of this tutorial, we'll add a template system and configure our app to show users all of their projects from Snyk in our App.
