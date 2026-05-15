# Scan and monitor images

It is common to use both `test` and `monitor` commands with Snyk Container. You can use the `snyk container test` command for quick checks. You can use the `snyk container monitor` command for ongoing assurance and to easily share results.

## Scan an image

To scan an image, run the `container test` command. For example:

```
snyk container test debian
```

The command:

1. Downloads the image, if it is not already available locally in your Docker daemon.
2. Determines the software installed in the image.
3. Sends the bill of materials to the Snyk service.
4. Returns a list of the vulnerabilities in your image.

You can use Snyk to test any image you can pull from a remote registry or any image you have built locally and made available in your local Docker daemon:

```
snyk container test <repository>:<tag>
```

If you use a Dockerfile to build your image, you can specify this when running `snyk container test`:

```
snyk container test <repository>:<tag> --file=Dockerfile
```

Specifying a Dockerfile provides more context and allows Snyk to provide clear recommendations on how to fix discovered vulnerabilities.

Snyk currently detects application vulnerabilities in your image by default.

## Monitor an image

Snyk Container also allows you to monitor images. This provides the following advantages:

* Snyk alerts you if new vulnerabilities that affect your image are disclosed without having to retest your image locally.
* Snyk interactively filters the results and explores the list of vulnerabilities in your web browser.
* You can share results on Snyk with other members of your team.

To learn more about container security, see [The importance of Container Monitoring](https://snyk.io/learn/container-security/container-monitoring/).

Users on [paid plans](https://snyk.io/plans) can also access aggregate reports of vulnerabilities across all of their Projects.

To monitor an image, run the `container monitor` command:

```
snyk container monitor <repository>:<tag>
```

This command:

1. Downloads the image if it is not already available locally in your Docker daemon.
2. Determines the software installed in the image.
3. Sends the bill of materials to the Snyk service.
4. Returns a link to the Snyk service, where you can see the results.

<figure><img src="../../../../.gitbook/assets/monitor (1).png" alt="Recommendatios for upgrading the base image"><figcaption><p>Recommendatios for upgrading the base image</p></figcaption></figure>
