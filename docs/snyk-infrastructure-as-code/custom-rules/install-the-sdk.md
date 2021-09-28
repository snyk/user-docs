# Install the SDK

​Install the SDK using one of these options:

* [​Install the SDK with npm​](install-the-sdk.md#install-the-snyk-cli-with-npm)
* [​Install the SDK using the prebuilt binaries​](install-the-sdk.md#install-the-snyk-cli-using-the-prebuilt-binaries)
* [​Install the SDK with Homebrew​](install-the-sdk.md#install-the-snyk-cli-with-homebrew)
* [​Install the SDK with the Windows Scoop package manager​](install-the-sdk.md#install-the-snyk-cli-with-the-windows-scoop-package-manager)

After installation, you can get started writing rules with our [Getting started](getting-started-with-the-sdk/) guide.

### Install the SDK with npm

Install our SDK using npm.

#### **Prerequisites**

* Ensure you’ve installed the latest version of npm on your local environment, using Node version 10 or later.

#### **Steps**

Run this command to install it for local use:

```text
npm install -g snyk-iac-rules
```

Once installed, you are ready to use the SDK. Run the following command to verify that it works:

```text
snyk-iac-rules --help
```

### Install the SDK using the prebuilt binaries

You can download and use the SDK's prebuilt binaries. To download the prebuilt binary, visit the [**Releases tab**](https://github.com/snyk/snyk-iac-rules/releases) in the SDK repository page in GitHub:

![](../../.gitbook/assets/screenshot-2021-09-24-at-13.44.36.png)

After you've downloaded the desired binary archive, open a terminal and run the following commands \(note that these commands assume you're running on an Intel-based macOS and downloading version `0.0.5` of the SDK\):

```text
$ tar xzf snyk-iac-rules_0.0.5_Darwin_x86_64.tar.gz 
$ sudo mv snyk-iac-rules /usr/local/bin
```

To verify the ability to use it, run:

```text
snyk-iac-rules --help
```

### Install the SDK with Homebrew

From macOS and Linux environments, you can use Homebrew to install our SDK. The repository for installation is stored in [our GitHub](https://github.com/snyk/homebrew-tap).

#### **Prerequisites**

* Supported for macOS and Linux environments only.
* Ensure [Homebrew](https://brew.sh/index_he) has already been installed:

  ```text
  brew tap snyk/tap
  ```

#### **Steps**

Install the SDK as follows:

```text
brew install snyk-iac-rules
```

### Install the SDK with the Windows Scoop package manager

From Windows environments, you can use Scoop to install our `snyk-iac-rules` SDK. The repository for installation is stored in [our GitHub](https://github.com/snyk/scoop-snyk).

#### **Prerequisites**

* Supported for Windows environments only.
* Ensure [Scoop](https://scoop.sh/) has already been installed:

  ```text
  scoop bucket add snyk https://github.com/snyk/scoop-snyk
  ```

#### **Steps**

Install the SDK as follows:

```text
scoop install snyk-iac-rules
```

### See also

* [​Getting started with the SDK​](getting-started-with-the-sdk/)
* ​[SDK reference​](sdk-reference.md)

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)​
{% endhint %}

