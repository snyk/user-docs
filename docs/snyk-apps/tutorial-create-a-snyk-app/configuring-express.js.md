# Configuring Express.js

## Set up Express to serve the application

So far, we've configured our project directories, set the TypeScript compiler options, and created a very simple TypeScript application to print out `Hello World`.

In this section, we'll expand our application by adding `ExpressJS` to serve our application and configure middleware to handle different kinds of requests.

## Install new dependencies

Run the following to install the new packages we'll need for this section. We'll cover each in turn.

```
npm install --save \
  express \
  express-rate-limit \
  express-session \
  http
```

Install TypeScript type and interface definitions for our packages as dev dependencies:

```
npm install --save-dev \
  @types/express \
  @types/express-rate-limit \
  @types/express-session \
  @types/node
```

## Containing the App

> This tutorial makes use of Object-oriented programming (OOP) concepts. If you aren't familiar with OOP, refer to [this MDN page](https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Objects/Object-oriented\_JS) for a brief primer.

To begin, create a new file to house the application code we'll be working on in this portion of the tutorial. If you've just come from the previous module where we configured a simple 'Hello world' app, you can leave your `index.ts` file alone; we'll come back to it.

```bash
touch ./src/app.ts
```

Add an import statement for Express and its TypeScript type / interface definitions at the top of the new file (`./src/app.ts`), then create an `App` class to house all the related functions and configuration required to run our application.

The class constructor will initialize Express on port 3000 when the class is instantiated. We'll create a private function `listen()` for the constructor to call which will handle the actual start of the Express server.

Don't forget to export the class at the bottom, we'll need it later.

```typescript
import express from 'express';
import type { Application } from 'express';
import type { Server } from 'http';

class App {
  public app: Application;
  private server: Server;

  constructor() {
    this.app = express();
    this.server = this.listen(3000);
  }

  private listen(port: number) {
    this.server = this.app.listen(port, () => {
      console.log(`App listening on port: ${port}`);
    });

    return this.server;
  }

}

export default App;
```

Test the new file using `npx tsc` and ensure an `app.js` file is successfully created in the `./dist` directory. If there are no errors, great! You're ready to start adding routes.

## Basic routes

Express will listen to any request on our given port, but it doesn't know _what to do_ when it receives a request. We need to tell Express _how_ to handle the various types of requests it might receive. How you configure this depends heavily on the architecture of your application, and which routes / pages you expect your app to respond to when hit.

For this tutorial, we'll keep things simple. Since our end goal is to provide users with a list of their projects within Snyk, assume we'll want at least two routes: an index route (`/`) to handle any initial connections and a projects route (`/projects`) to present the data from Snyk.

To keep our project organized, we'll create a `routes` directory to house our collection of routes and we'll give each route its own subdirectory which will contain the main controller file for the route and any extras we might need.

We'll focus on the index route for this module:

```bash
mkdir -p ./src/routes/index
touch ./src/routes/index/indexController.ts
```

Before we start working on the index controller, we should also create a [TypeScript interface](https://www.typescriptlang.org/docs/handbook/interfaces.html) definition to describe a common shape for controller data. This will help keep our controllers consistent and allow TypeScript to warn us early if a controller is missing something important.

Let's use a similar separation pattern as the one we decided on for routes. Since interface definitions in TypeScript are typically self-contained and descriptive, we'll skip creating a directory for each definition and store any interface files we create as siblings:

```bash
mkdir -p ./src/interfaces
touch ./src/interfaces/Controller.ts
```

Any route controller we create should at the very least take a path value and a router from Express. Edit `./src/interfaces/Controller.ts` and populate it with the content below.

```typescript
import type { Router } from 'express';

export interface Controller {
  path: string;
  router: Router;
}
```

With that done, open the route controller file (`./src/routes/index/indexController`) in your editor and add the following:

```typescript
import type { Controller } from '../../interfaces/Controller';
import type { Request, Response, NextFunction } from 'express';
import { Router } from 'express';

class IndexController implements Controller {
  public path = '/';
  public router = Router();

  constructor() {
    this.initRoutes();
  }

  private initRoutes() {
    this.router.get(`${this.path}`, this.indexPage);
  }

  private indexPage(req: Request, res: Response, next: NextFunction) {
    return res.send('index route!');
  }
}

export default IndexController;
```

Here we've imported a handful of TypeScript type/interface definitions as well as the Router object from the Express package.

We've set our `IndexController` to respond to `/` requests with the `indexPage` function, which renders 'index', but our App class doesn't yet have a way to register the route when it instantiates. Let's go back to our `./src/app.ts` file and add an `initRoutes()` function to establish that.

We'll also modify the App's constructor function to take an array of Controllers and a port value as arguments.

Edit `./src/app.ts` and update the contents to match the following:

```typescript
import express from 'express';
import type { Application } from 'express';
import type { Server } from 'http';
import type { Controller } from './interfaces/Controller';

class App {
  public app: Application;
  private server: Server;

  constructor(controllers: Controller[], port: number) {
    this.app = express();
    this.initRoutes(controllers);
    this.server = this.listen(port);
  }

  private initRoutes(controllers: Controller[]) {
    controllers.forEach((controller: Controller) => {
      this.app.use('/', controller.router);
    });
  }

  private listen(port: number) {
    this.server = this.app.listen(port, () => {
      console.log(`App listening on port: ${port}`);
    });

    return this.server;
  }

}

export default App;
```

## Running the Express server

So far, we've created an App class which, when instantiated, will run an Express server which responds to requests at `/` and we've added an interface definition to ensure all our controllers follow the same data pattern.

Let's run our project and try it out!

If you run the compiled `./dist/app.js` at this point, nothing interesting happens. This is because we've yet to actually instantiate our exported class. Do you remember where, in the previous module, we initialized our `package.json`? During that command, we were presented with questions pertaining to our app. One of the questions asked what our project's entrypoint was. If you were following along and selected the default value, this should be set to `index.js` in your `package.json`. This entrypoint is where we'll instantiate the App class.

Edit `./src/index.ts` (where we had that hello world code) and clear it out if there's any content left over, then add import statements for our App and other required objects.

Our entrypoint will only do one thing, instantiate the App class with the `new` keyword:

```typescript
import IndexController from './routes/indexController';
import App from './app';

new App(
  [
    new IndexController()
  ],
  3000
);
```

That's it!

Build the project using `npx tsc` then run the compiled entrypoint with `node`:

```bash
npx tsc && node ./dist/index.js
```

Provided everything was successful, you'll see the message `App listening on port: 3000` in the console.

Now visit http://localhost:3000 in your browser or use curl to check the response. If you see `index route!`, then congratulations! Express is serving the application and responding to routes!

In the next module of this tutorial, we'll dive deeper into routes and authentication and start working with the Snyk API.
