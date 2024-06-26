# Versioning schemas for custom base images

You must set a versioning schema for the first Project you marked as a custom base image in the image's repository. You can edit the Custom base Image settings for Projects that you have already marked as custom base images. For details. see [Mark the created Project as a custom base image](./#mark-the-created-project-as-a-custom-base-image).

Snyk recommends updating a custom base image to the latest available version from the pool of images imported as Projects, marked as custom base images and enabled for recommendations. For more information, see [Enable CBIR](./#enable-cbir).

To determine which is the latest version, Snyk uses a versioning schema that allows comparing image tags.

You can manage the versioning schema for images imported to Projects for a repository in the Web UI or using the Snyk REST API.

Snyk supports the following versioning schemas.

## **Semantic Versioning (SemVer)**

This schema adheres to the [Semantic Versioning Specification](https://semver.org/) (SemVer).

In the below example, when Snyk scans the images and marks them as custom base images in the order shown, the SemVer logic is:

1. `somecompany/alpine-base:3.18.3`
2. `somecompany/alpine-base:3.19.0`
3. `somecompany/alpine-base:3.18.5`

Snyk recommends the second image, as it is the newest image based on the semantic versioning of the tag.

The following commonly used examples are not supported SemVer tags:

* `v1.2.3`
* `1.2`
* `latest`

## Single Selection

This versioning schema allows you to manually set a single image as a recommendation.

When selecting an image, any other previously selected images are automatically unselected.

## Custom versioning schema

If none of the above schemas match your requirements, you may create a custom versioning schema. For more information, see [Custom Versioning Schema for custom base images](custom-versioning-schema-for-custom-base-images.md).
