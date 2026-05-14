# Install or upgrade to version of Node.js required for Snyk CLI

Node.js v12 or higher is required for Snyk CLI version 1.853.0 and higher. Snyk recommends running as recent a version of Node.js and npm as possible to ensure full compatibility.

### How to upgrade Node.js

Node.js comes with npm pre-installed, but the manager is updated more frequently than Node.js.

To update Node.js, you need need the npm [n module](https://www.npmjs.com/package/n).

Run this code to clear the npm cache, install `n`, and install the latest stable version of Node.js:

```
sudo npm cache clean -f
sudo npm install -g n
sudo n stable
```

To install the latest release, use `n latest`. Alternatively, you can run `n #.#.#` to get a specific Node.js version.

Follow these steps to upgrade Node.js:

1. Run `npm-v` to see which version you have.
2. Then run`npm install npm@latest -g` to install the newest npm update.
3. Run `npm -v` again to make sure npm updated correctly.

### How to upgrade npm

To upgrade npm, run the following:

```
npm install -g npm
```
