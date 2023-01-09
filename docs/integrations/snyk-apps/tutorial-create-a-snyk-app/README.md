# Tutorial: create a Snyk App

In this tutorial, we'll create a simple Snyk App using TypeScript to show users a list of their projects within Snyk.

## Prerequisites

* NodeJS
* A Snyk account

## Configure the basics

To begin, create a directory to house the project somewhere on your device. From within the newly created directory, we'll initialize a `package.json` manifest for our application to keep track of our dependencies and ensure our project is portable:

Run `npm init` and follow the prompts. You can add as much or as little information as you want here. For now, the default options suffice.

Now that we have a place to save dependency information, use `npm` to install TypeScript as a development dependency:

```
npm install typescript --save-dev
```

At this point, TypeScript has been installed, but we'll need a configuration file to provide compilation options to the `tsc` binary. Create a TypeScript configuration file i called `tsconfig.json` n the root of the project. Use the template that follows:

```json
{
  "compilerOptions": {
    "target": "es6",
    "module": "commonjs",
    "declaration": true,
    "sourceMap": true,
    "moduleResolution": "node",
    "outDir": "./dist",
    "rootDir": "./src",
    "strict": true,
    "skipLibCheck": true,
    "esModuleInterop": true
  },
  "exclude": [
    "./tests",
    "./dist"
  ]
}
```

The options we've provided tell TypeScript to emit ES6 JavaScript, which type of module code to generate, and whether or not to provide a corresponding source map for the compiled files, and specify a few other handy options. For a complete overview of the possible options, visit https://aka.ms/tsconfig.json

For the purposes of this tutorial, the most noteworthy options we've set are `rootDir` and `outDir`. These options describe where our source `.ts` files and our compiled `.js` files, respectively, live within our project. Create the directories referenced by the setting values:

```bash
mkdir ./dist
mkdir ./src
```

## Test it out

Now that we have the basic parameters in place, we'll create a simple Hello World by creating the file `./src/index.ts`.

```
const world = 'world';

export function hello(world: string = world): string {
  return `Hello ${world}! `;
}
```

Now we can confirm that everything is wired up correctly. Run `npx tsc` to compile the project.

If everything is successful, the project tree looks like this:

```
my-snyk-app/
 - dist/
   - index.d.ts
   - index.js
   - index.js.map
 - src/
   - index.ts
 - node_modules/
   - <lots of things here>
 - package-lock.json
 - tsconfig.json
```

The `tsc` program compiled our source TypeScript file into ES6 JavaScript and provided a source map for debugging.

{% hint style="info" %}
The compiler can also be put into watch mode to continually poll your `.ts` files for changes: `npx tsc -w`
{% endhint %}
