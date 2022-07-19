# Updating the Snyk CLI Version

When using the Snyk CLI, it is recommended to use the latest version, since new features and improvements are continuously added to it. Some new features will require you to use a certain Snyk CLI version or later versions. Since the Snyk CLI version is not updated automatically, you need to manually update it to the latest version. However, you do not need to update the Snyk CLI every time you use it.

**Note**: For more information on each Snyk CLI version, see the [Snyk CLI - Release Notes page](https://github.com/snyk/cli/releases). &#x20;

### **Checking the current Snyk CLI version**

**To check your current Snyk CLI version:**

* In the terminal, enter one of the following commands:

```
snyk version 
```

&#x20;– or –

```
snyk --version 
```

– or –

```
snyk -v  
```

### **Updating your Snyk CLI Version**

Updating the Snyk CLI version only impacts your CLI version. CLI updates do not upgrade your dependencies automatically.

**To update your Snyk CLI to the latest version:**

* In the terminal, enter one of the following commands:

```
npm update snyk -g
```

– or –

```
npm i -g snyk
```

**To update your Snyk CLI to a specific version:**

* In the terminal, enter:

```
npm i -g snyk@<version_no>
```

For example, to update the Snyk CLI version to version 1.835.0, we enter:

```
npm i -g snyk@1.835.0
```

&#x20;
