# Snyk CLI for container security

## Test images with the Snyk Container CLI

The Snyk Container command line interface (CLI) helps you find and fix vulnerabilities in container images on your local machine.

## Install the Snyk CLI

Use any of the following:

* npm – `npm install -g snyk`
* Homebrew – `brew tap snyk/tap && brew install snyk`
* Scoop - `scoop bucket add snyk https://github.com/snyk/scoop-snyk`
* [A manual installer](https://github.com/snyk/snyk/releases) available from GitHub

For more detailed installation guidance and options, see [Install the Snyk CLI](https://support.snyk.io/hc/articles/360003812538#UUID-7ccc55c8-51f7-ff54-5acf-680dc62bc27e).

## Authentication

After installation, authenticate with Snyk to test your image, running snyk auth from the CLI:

```
snyk auth
```

For more details about authentication, see [Authenticate the CLI with your account](https://docs.snyk.io/snyk-cli/install-the-snyk-cli/authenticate-the-cli-with-your-account)

## Testing an image

To test an image run:

```
snyk container test debian
```

This:

1. Downloads the image if it’s not already available locally in your Docker daemon
2. Determines the software installed in the image
3. Sends that bill of materials to the Snyk Service
4. Returns a list of the vulnerabilities in your image

You can use Snyk to test any image that you can pull from a remote registry, or any image you have built locally and made available in your local Docker daemon.

```
snyk container test <repository>:<tag>
```

If you use a Dockerfile to build your image, you can provide that when running Snyk.

```
snyk container test <repository>:<tag> --file=Dockerfile
```

Specifying a Dockerfile provides more context, and allows Snyk to provide clear recommendations on how to fix discovered vulnerabilities.

## Monitoring an image

Snyk Container also has the concept of monitoring an image. This provides the following advantages:

* Snyk will alert you if new vulnerabilities are disclosed that affect your image, without you having to retest it locally
* Interactively filter the results and explore the list of vulnerabilities in your web browser
* Results on Snyk can be shared with other members of your team

You can also access aggregate reports of vulnerabilities across all of your projects.

{% hint style="info" %}
**Feature availability**\
This aggregate reports feature is available with all paid plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

To monitor an image run:

```
snyk container monitor <repository>:<tag>
```

monitor will:

1. Download the image if it is not already available locally in your Docker daemon
2. Determine the software installed in the image
3. Send that build of materials to the Snyk Service
4. Return a link to the Snyk service where you can see the results

![](../../../.gitbook/assets/monitor.png)

{% hint style="info" %}
**Note**\
It’s common to use both test and monitor with Snyk. test is great for quick checks, monitor can be used for ongoing assurance and easier sharing of results.
{% endhint %}

## More information

* [Understand Snyk Container CLI results](https://docs.snyk.io/snyk-container/snyk-cli-for-container-security/understanding-snyk-container-cli-results)
* [Advanced CLI usage](https://docs.snyk.io/snyk-container/snyk-cli-for-container-security/advanced-snyk-container-cli-usage)
* Learn more about [container security](https://snyk.io/learn/container-security/)
