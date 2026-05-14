# Snyk Images and EOL image policy

This page outlines the Snyk end-of-life (EOL) policy for images provided by the Snyk Images build toolchain, which is available in the [`snyk-images` GitHub repository](https://github.com/snyk/snyk-images).

## Summary of EOL image policy

As part of Snyk's commitment to helping customers ship secure code, Snyk will provide images that contain software supported by its upstream vendor as a convenience.

### Image maintenance

As a general practice, Snyk does not remove images once they are published. However, Snyk will not build, maintain, or ship images based on EOL software.

### User responsibility

To ensure you are using the most current and secure images, monitor the software you use. When a product announces an image version as EOL, switch to a Snyk image that contains supported software where appropriate.

### Example of an EOL image

For example, Snyk stopped building images based on Python 3.7, which is not actively supported and does not receive security updates from Python. Thus users should migrate to Python 3.8 or higher.

## Snyk Images lifecycle

### Annual review

Snyk conducts an annual review of images. Images containing software that has reached its end-of-life (EOL) status will not be rebuilt or updated. This means the CLI and base image will no longer receive updates, bug fixes, or security patches.

### Image availability

Snyk does not remove images from Docker Hub once the images are published. Users should check the [endoflife.date](https://endoflife.date/) website to ensure their build pipelines are not using EOL software.  Snyk will not issue warnings about an image becoming EOL.

Snyk follows a phased approach to building, maintaining, and discontinuing maintenance of an image. The phases are [Creation](snyk-images-and-eol-image-policy.md#creation-phase), [Active support and maintenance](snyk-images-and-eol-image-policy.md#active-support-and-maintenance-phase), and [End of Life (EOL)](snyk-images-and-eol-image-policy.md#end-of-life-eol-phase).

### Creation phase

Based on demand or request, a base image can be added to `snyk-images` if the base image is vendor-supported.

Note that Snyk accepts base images only from trusted vendors.

{% hint style="info" %}
Snyk is not responsible for securing the vendor-provided base image itself.
{% endhint %}

The following software is bundled with the Snyk CLI in a Docker image.  For a complete list of available versions, see [the current images](https://github.com/snyk/snyk-images/tree/master?tab=readme-ov-file#current-images) list in the `snyk-images` repository.

* Alpine
* Swift
* Clojure
* Composer
* Docker
* Golang
* Gradle
* Gradle dk
* Maven
* Maven jdk
* Microsoft .NET
* Node
* Python
* Ruby
* Ubuntu
* Scala/sbt

### Active support and maintenance phase

As long as the vendor supports the base image, the Snyk image will include the latest stable Snyk CLI version and a regularly updated version of the base image. &#x20;

### End of Life (EOL) phase

After the vendor software reaches EOL, that is, stops receiving active support and security updates, Snyk will also stop building and maintaining the impacted image.

Snyk will update the `snyk-images` repository and user documentation to remove the EOL images from the list of supported images.

For more information, see the lists of [currently supported](https://github.com/snyk/snyk-images?tab=readme-ov-file#current-images) and [unsupported ](https://github.com/snyk/snyk-images?tab=readme-ov-file#vendor-unsupported-base-images)images in the [`snyk-images` repository](https://github.com/snyk/snyk-images).
