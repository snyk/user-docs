# CLI for Snyk Code

{% hint style="info" %}
This feature is currently in beta.
{% endhint %}

The Snyk Command Line Interface \([CLI](../../snyk-cli/)\) for Snyk Code helps you find and fix security flaws in your code on your local machine.

### **Install the Snyk CLI**

Use any of the following:

* npm: `npm install -g snyk`
* Homebrew: `brew tap snyk/tap && brew install snyk`
* Scoop: `scoop bucket add snyk <https://github.com/snyk/scoop-snyk>`
* [**A manual installer**](https://github.com/snyk/snyk/releases/) available from GitHub

For more detailed installation guidance and options, see [Install the Snyk CLI](../../snyk-cli/install-the-snyk-cli/).

### **Authentication**

After installation, authenticate with Snyk to test your image, running snyk auth from the CLI:

```text
snyk auth
```

For more details about authentication, see [Authenticate the CLI with your account](../../snyk-cli/install-the-snyk-cli/authenticate-the-cli-with-your-account/)

### **Testing a project or folder**

To test the current folder, run `snyk code test` with no parameters.

To test another context, run `snyk code test <my-folder-path>` with a path to a file or folder as the parameter.

All sub-folders inside the provided folder are also scanned.

