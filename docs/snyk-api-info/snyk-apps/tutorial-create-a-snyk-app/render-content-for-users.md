# Render content for users

In the previous module, we covered registering our Snyk App, setting up the authorization flow, and handling user authorization within our App. All of those topics are integral to the functionality of every Snyk App, but they're all what you might call "behind the scenes" topics.

In this module, we'll switch gears to focus on displaying content to the users who have authorized with our Snyk App. Specifically, we want to show unauthorized users a big button they can click to authorize and authorized users a list of their projects from Snyk.

## Add a template engine to the Snyk App

While Express is perfectly capable of printing content to the screen and even rendering HTML server-side, life is much easier when using a template engine. For this tutorial we are using [EJS](https://ejs.co).

First, install the node packages needed in this part of the tutorial:

```bash
npm install --save ejs
```

Next, modify the `initGlobalMiddlewares()` function we created in our last module to tell express that we want to use a _view engine_, EJS in this case, and let it know where we'll be storing our view templates. We'll be storing our EJS templates in `./src/views` and keeping any common files such as images and CSS in `/.src/public`.

Create the new directories first.

```bash
mkdir -p ./src/views/partials
mkdir -p ./src/public
```

Now we can update `./src/app.ts`:

```typescript
// ./src/app.ts
...

class App {

  ...

  private initGlobalMiddlewares() {

    ...

    this.app.set("views", path.join(__dirname, "/views"));
    this.app.set("view engine", "ejs");
    this.app.use('/public', express.static(path.join(__dirname, '/public')));

    ...

  }

  ...

}
```

For each route that we'll provide a template for, we'll need to modify the corresponding controller and ensure that we're using `res.render("<template name>")` rather than something more simplistic like `res.send()`.

Example:

```typescript
...

private initRoutes() {
  this.router.get(`${this.path}`, this.indexPage);
}
private indexPage(req: Request, res: Response, next: NextFunction) {
  // THIS right here is what causes Express to render the EJS template.
  return res.render("index");
}

...
```

That's all there is to it.

EJS templates support the concept of partial inclusion. While not strictly necessary, it makes sense to add a subdirectory to our `./src/views` to differentiate partial templates like headers and footers from route templates. For the tutorial, we'll use `./src/views/partials` to store such templates.

## Basic EJS templates

The first template we'll create is a partial one, which we'll include in the other templates. This `header.ejs` will be the place we link stylesheets and other information that belong in the `<head>` of an HTML document.

```ejs
// ./views/partials/header.ejs

<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
    <link href="https://fonts.googleapis.com/css2?family=Roboto:wght@400;500&display=swap" rel="stylesheet" />
    <link href="https://raw.githubusercontent.com/snyk/snyk-apps-demo/main/src/public/css/styles.css" />
    <link rel="shortcut icon" href="https://raw.githubusercontent.com/snyk/snyk-apps-demo/main/src/public/images/snyk_dog.svg" type="image/x-icon" />

    <title>Snyk Apps Tutorial</title>
  </head>
</html>
```

This `index.ejs` template will cover our basic `/` route.

```ejs
// ./views/index.ejs

<%- include('./partials/header.ejs') %>

<body>
  <div class="index-page">
    <img class="index-page__snyk-logo" src="https://github.com/snyk/snyk-apps-demo/raw/main/src/public/images/snykLogoWithDog.svg" alt="snyk-logo" />
    <div class="index-page__card">
      <h1 class="index-page__title">Add Demo App</h1>
      <p class="index-page__description">Authorize this App to connect to your Snyk account.</p>
      <button class="button" onclick="location.href='/auth';">Install App</button>
    </div>
  </div>
</body>
```

`callback.ejs` will render for successful user authorizations.

```ejs
// ./views/callback.ejs

<%- include('./partials/header.ejs') %>
<body>
  <div>
    <h2 class="main__heading">Snyk Apps Tutorial</h2>
  </div>
  <div class="card__30">
    <div class="callback-page__success-box">
      <img class="snyk-con-img" src="https://github.com/snyk/snyk-apps-demo/raw/main/src/public/images/success_check.svg" alt="success" />
        <div>
          <h2 class="callback-page__success-text">Successfully connected to Snyk!</h2>
        </div>
        <button class="button" onclick="location.href='/projects';">List Projects</button>
      </div>
    </div>
  </div>
</body>
```

These templates should be enough to get you started adding your own templates to any new routes you create. If you intend to continue using EJS, refer to the documentation for information about the features offered.

Rendering content for your Snyk App can be as simple or complex as you'd like it to be. Because we're dealing with JavaScript, the options are very flexible!

## Showing users a list of projects

Now that we have some basic templates, take a look at how we can add some functionality to our Snyk App using a User's Snyk data. For this tutorial, we set up our app to allow users to view all of their projects within Snyk from within our app.

This is a basic and easily extendable feature.

We need to create:

* A new route controller
* A function (or functions) to pull the project data
* An EJS template for showing the projects

We start with the API work, using the `callSnykApi()` function we created in the previous module. Since this directly relates to a particular route, we'll store this file with its controller. Following the pattern we've used throughout these tutorial modules, we'll create both files at `./src/routes/projects/`.

```typescript
// ./src/routes/projects/projectsHandler.ts

import { readFromDb } from "../../util/DB";
import { callSnykApi } from "../../util/apiHelpers";
import { EncryptDecrypt } from "../../util/encrypt-decrypt";
import { AuthData } from "../../interfaces/DB";
import { APIVersion } from "../../interfaces/API";
import { ENCRYPTION_SECRET } from "../../app";

/**
 * Get projects handler that fetches all user projects
 * from the Snyk API using user access token. This for
 * example purposes. In production it will depend on your
 * token scopes on what you can and can not access
 * @returns List of user project or an empty array
 */
export async function getProjectsFromApi(): Promise<unknown[]> {
  // Read data from DB
  const db = await readFromDb();
  const data = mostRecent(db.installs);
  // If no data return empty array
  if (!data) return [];

  // Decrypt data(access token)
  const eD = new EncryptDecrypt(ENCRYPTION_SECRET as string);
  const access_token = eD.decryptString(data?.access_token);
  const token_type = data?.token_type;
  // Call the axios instance configured for Snyk API v1
  const result = await callSnykApi(
    token_type,
    access_token,
    APIVersion.V1
  ).post(`/org/${data?.orgId}/projects`);

  return result.data.projects || [];
}

/**
 *
 * @param {AuthData[]} installs get most recent install from list of installs
 * @returns the latest install or void
 */
export function mostRecent(installs: AuthData[]): AuthData | void {
  if (installs) {
    return installs[installs.length - 1];
  }
  return;
}
```

Next we'll write the route controller. Follow the pattern: `./src/routes/projects/projectsController.ts`.

```typescript
// ./src/routes/projects/projectsController.ts

import type { Controller } from "../../interfaces/Controller";
import type { NextFunction, Request, Response } from "express";
import { Router } from "express";
import { getProjectsFromApi } from "./projectsHandler";

export class ProjectsController implements Controller {
  // The base URL path for this controller
  public path = "/projects";
  // Express router for this controller
  public router = Router();

  constructor() {
    this.initRoutes();
  }

  private initRoutes() {
    // The route to render all user projects lists
    this.router.get(`${this.path}`, this.getProjects);
  }

  private async getProjects(req: Request, res: Response, next: NextFunction) {
    try {
      const projects = await getProjectsFromApi();
      return res.render("projects", {
        projects,
      });
    } catch (error) {
      return next(error);
    }
  }
}
```

Whenever we add a new route controller, we need to update `./index.ts` to include it.

```typescript
// ./src/index.ts

import IndexController from "./routes/index/indexController";
import AuthController from "./routes/auth/authController";
import CallbackController from "./routes/callback/callbackController";
import ProjectsController from "./routes/projects/projectsController";
import App from "./app";

new App([
   new IndexController(),
   new AuthController(),
   new CallbackController()
   new ProjectsController()],
  3000
);
```

## Wrap-up

Using the projects API handler and controller we created in this module, you should have all you need to create your own custom code and make your Snyk App do whatever you'd like it to do.

We used the v1 API but keep an eye on Snyk's REST API. As additional features are added, you may find new or more efficient endpoints to use in your Snyk App.
