# EOL policy: Snyk Images

This page outlines the Snyk end-of-life policy for images provided by the [Snyk Images build tool chain](https://github.com/snyk/snyk-images). Snyk is introducing this policy, and it takes effect immediately as explained in the [product announcement](https://updates.snyk.io/important-update-on-snyk-images-obsolete-software-packages-294548).

## Purpose of policy

As a continued commitment to helping customers ship secure code, Snyk will provide images that contain software supported by its upstream vendor. Snyk will not build, maintain, or ship images based on EOL software. For example, Snyk has stopped building images based on Python 3.7, which is not actively supported does not receive security updates from Python.&#x20;

## Snyk Images life cycle

Snyk is updating this page annually for planned removal of images containing software nearing end of life.&#x20;

Users are encouraged to check this page from time to time, and also to determine whether their build pipelines are using EOL software by checking the [endoflife.date](https://endoflife.date/) website.&#x20;

Snyk follows a phased approach to building, maintaining, and sunsetting an image.&#x20;

### Creation phase

Based on demand and requests, a base image can be added to `snyk-images` if the base image is vendor-supported.

Snyk accepts only base images from trusted vendors. The table that follows lists software bundled with the Snyk CLI in a Docker image.  For a complete list of versions that are currently available, see [the current images](https://github.com/snyk/snyk-images/tree/master?tab=readme-ov-file#current-images) repository.&#x20;

{% hint style="info" %}
Snyk is not responsible for securing the vendor-provided base image itself.
{% endhint %}

| Software bundled with Snyk |
| -------------------------- |
| Alpine                     |
| Swift                      |
| Clojure                    |
| Composer                   |
| Docker                     |
| Golang                     |
| Gradle                     |
| Gradle dk                  |
| Maven                      |
| Maven jdk                  |
| Microsoft .NET             |
| Node                       |
| Python                     |
| Ruby                       |
| Ubuntu                     |
| Scala/sbt                  |

### Active support and maintenance phase

As long as the vendor supports the base image, the Snyk image will include the latest stable Snyk CLI version and a regularly updated version of the base image. &#x20;

### End of Life (EOL) phase

After the vendor software reaches EOL, that is, stops receiving active support and security updates, Snyk will also stop building, maintaining, and shipping the impacted image. Snyk will take this action during a four-month planned removal process.&#x20;

Snyk will update the current images repository and user documentation to reflect this change followed by complete removal of images containing EOL software.

Here is how Snyk implements this phased approach during the four-month annual cycle:

| Annual Timeline | Planned Changes                                                                                   | Customer Impact                                                                          |
| --------------- | ------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| May             | Update current images repository and user documentation with the list of Snyk Images nearing EOL. | No changes to existing images at this point,                                             |
| June            | Stop building impacted images,                                                                    | Impacted images will stop receiving latest stable Snyk CLI release and security updates, |
| July            | Create and publish new Snyk images where applicable,                                              | Snyk will introduce new images as applicable,                                            |
| July            | Create and publish migration user guides on docs.snyk.io where applicable,                        | User documentation will be updated with migration guides,                                |
| August          | Remove impacted images from Docker Hub,                                                           | Impacted images will be removed,                                                         |

For more details, see [Images planned to be removed in 2024](../../../scm-ide-and-ci-cd-workflow-and-integrations/snyk-ci-cd-integrations/eol-policy-snyk-images/images-planned-to-be-removed-in-2024.md) and [Images no longer supported by Snyk](../../../scm-ide-and-ci-cd-workflow-and-integrations/snyk-ci-cd-integrations/eol-policy-snyk-images/images-no-longer-supported-by-snyk.md).
