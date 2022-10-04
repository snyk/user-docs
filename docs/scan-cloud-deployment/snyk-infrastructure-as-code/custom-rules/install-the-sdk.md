# Install the SDK

​Install the SDK using one of these options:

* [Install the SDK with npm](install-the-sdk.md#install-the-sdk-with-npm)
* ​[Install the SDK using the prebuilt binaries​](install-the-sdk.md#install-the-sdk-using-the-prebuilt-binaries)
* [Install the SDK with Homebrew](install-the-sdk.md#install-the-sdk-with-homebrew)
* [​Install the SDK with the Windows Scoop package manager](install-the-sdk.md#install-the-sdk-with-the-windows-scoop-package-manager)
* [Install the SDK with Docker](install-the-sdk.md#install-the-sdk-with-docker)

After installation, you can get started writing rules with our [Getting started](getting-started-with-the-sdk/) guide.

### Install the SDK with npm

Install our SDK using npm.

#### **Prerequisites**

* Ensure you’ve installed the latest version of npm on your local environment, using Node version 10 or later.

#### **Steps**

Run this command to install it for local use:

```
npm install -g snyk-iac-rules
```

Once installed, you are ready to use the SDK. Run the following command to verify that it works:

```
snyk-iac-rules --help
```

### Install the SDK using the prebuilt binaries

You can download and use the SDK's prebuilt binaries. To download the prebuilt binary, visit the [**Releases tab**](https://github.com/snyk/snyk-iac-rules/releases) in the SDK repository page in GitHub:

![](../../../.gitbook/assets/screenshot-2021-09-24-at-13.44.36.png)

After you've downloaded the desired binary archive, open a terminal and run the following commands (note that these commands assume you're running on an Intel-based macOS and downloading version `0.0.5` of the SDK):

```
$ tar xzf snyk-iac-rules_0.0.5_Darwin_x86_64.tar.gz 
$ sudo mv snyk-iac-rules /usr/local/bin
```

To verify the ability to use it, run:

```
snyk-iac-rules --help
```

### Install the SDK with Homebrew

From macOS and Linux environments, you can use Homebrew to install our SDK. The repository for installation is stored in [our GitHub](https://github.com/snyk/homebrew-tap).

#### **Prerequisites**

* Supported for macOS and Linux environments only.
*   Ensure [Homebrew](https://brew.sh/index\_he) has already been installed:

    ```
    brew tap snyk/tap
    ```

#### **Steps**

Install the SDK as follows:

```
brew install snyk-iac-rules
```

### Install the SDK with the Windows Scoop package manager

From Windows environments, you can use Scoop to install our `snyk-iac-rules` SDK. The repository for installation is stored in [our GitHub](https://github.com/snyk/scoop-snyk).

#### **Prerequisites**

* Supported for Windows environments only.
*   Ensure [Scoop](https://scoop.sh) has already been installed:

    ```
    scoop bucket add snyk https://github.com/snyk/scoop-snyk
    ```

#### **Steps**

Install the SDK as follows:

```
scoop install snyk-iac-rules
```

### Install the SDK with Docker

You can use Docker to install and run our `snyk-iac-rules` SDK while writing your custom rules in your local directory. The image is stored [in our Docker Hub repo](https://hub.docker.com/r/snyk/snyk-iac-rules).

#### **Prerequisites**

* Ensure [Docker](https://docs.docker.com/get-docker/) has already been installed.
* Supported for Linux containers only.

#### **Steps**

Pull the Docker image as follows:

```
docker pull snyk/snyk-iac-rules
```

Run the SDK by using the following command:

```
docker run --rm -v $(pwd):/app snyk/snyk-iac-rules {SDK command}
```

For example, to generate a custom rules template you can run:

```
docker run --rm -v $(pwd):/app snyk/snyk-iac-rules template -r {rule_name}
```

### See also

* [​Getting started with the SDK​](getting-started-with-the-sdk/)
* ​[SDK reference​](sdk-reference.md)
