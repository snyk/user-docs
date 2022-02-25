# Snyk CLI for container security

The Snyk Container command line interface ([CLI](../../../features/snyk-cli/)) helps you find and fix vulnerabilities in container images on your local machine.

To use the CLI you must first [install](../../../features/snyk-cli/install-the-snyk-cli/) it and then [authenticate](../../../features/snyk-cli/commands/auth.md).

## Testing an image

To test an image run:

```
snyk container test debian
```

This command does the following:

1. Downloads the image if it is not already available locally in your Docker daemon
2. Determines the software installed in the image
3. Sends that bill of materials to the Snyk service
4. Returns a list of the vulnerabilities in your image

You can use Snyk to test any image that you can pull from a remote registry, or any image you have built locally and made available in your local Docker daemon.

```
snyk container test <repository>:<tag>
```

If you use a Dockerfile to build your image, you can specify that when running `snyk container test`.

```
snyk container test <repository>:<tag> --file=Dockerfile
```

Specifying a Dockerfile provides more context, and allows Snyk to provide clear recommendations on how to fix discovered vulnerabilities.

## Monitoring an image

Snyk Container also allows you to monitor an image. This provides the following advantages:

* Snyk alerts you if new vulnerabilities are disclosed that affect your image, without your having to retest your image locally.
* Snyk interactively filters the results and explores the list of vulnerabilities in your web browser.
* You can share results on Snyk with other members of your team.

You can also access aggregate reports of vulnerabilities across all of your projects.

{% hint style="info" %}
**Feature availability**\
This aggregate reports feature is available with all paid plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

To monitor an image run:

```
snyk container monitor <repository>:<tag>
```

This command does the following

1. Downloads the image if it is not already available locally in your Docker daemon
2. Determines the software installed in the image
3. Sends that bill of materials to the Snyk service
4. Returns a link to the Snyk service where you can see the results

![](../../../.gitbook/assets/monitor.png)

{% hint style="info" %}
**Note**\
It’s common to use both `test` and `monitor` with Snyk Container. The `test` command is great for quick checks. The `monitor` command can be used for ongoing assurance and easier sharing of results.
{% endhint %}

## More information

* [Understand Snyk Container CLI results](https://docs.snyk.io/snyk-container/snyk-cli-for-container-security/understanding-snyk-container-cli-results)
* [Advanced CLI usage](https://docs.snyk.io/snyk-container/snyk-cli-for-container-security/advanced-snyk-container-cli-usage)
* Learn more about [container security](https://snyk.io/learn/container-security/)
