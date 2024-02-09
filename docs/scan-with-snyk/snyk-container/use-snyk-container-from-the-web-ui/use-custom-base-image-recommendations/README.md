# Use Custom Base Image Recommendations

{% hint style="warning" %}
**Release status**&#x20;

Custom Base Image Recommendations is available only for Enterprise plans.

See [Pricing plans](https://snyk.io/plans).
{% endhint %}

## **Overview of Custom Base Image Recommendations (CBIR)**

When scanning a container image, Snyk provides recommendations based on the base image it detects. Because Snyk precomputes recommendations only for Docker Official Images, these recommendations apply only to a subset of images.

Customers often maintain their own internal base images, built on top of Docker Official Images or other upstream images. These are provided as a service to a wider set of development teams. For example: `somecompany/base-python:3.12.1`.

The Custom Base Image Recommendation feature allows Snyk to recommend an image upgrade from a pool of your internal images. This allows teams to be notified of newer and more secure versions of their internal base images.

## How CBIR works

To use the Custom Base Image Recommendations feature, the base images, as well as the application images built on top of a base image, must be imported to Snyk as Projects. These Projects can be in different organizations, but they must belong to the same Group.

To receive custom base image recommendations for an application image, you must specify a Dockerfile in the Project for the application image.

{% hint style="info" %}
As opposed to public Docker Official Images, Snyk can only detect a custom base image when a Dockerfile is linked to the application image Project.
{% endhint %}

All custom base image recommendations are considered minor upgrades, regardless of the image tag.

To determine the latest version of a base image across Projects imported for the same repository, Snyk allows configuring a [versioning schema](versioning-schema-for-custom-base-images.md).

Custom base image recommendation supports Automatic fix PRs. If you are not using the latest version of the base image, then immediately after image import, Snyk automatically issues a fix pull request against your Dockerfile to upgrade to the latest available custom base image version.

## **Enable CBIR**

The following steps describe how to configure custom base images.

### Configure an image as a custom base image

The team that is responsible for creating and maintaining custom base images for the Organization performs this step.

#### Build and import a base image

1. Build a custom base image.
2. Import the image to a Snyk Project using one of the following options:
   * The Snyk Web UI: Import an image into Snyk using a container registry
   * The CLI: Use `snyk container monitor`  command.
     * Set the `--project-name` flag (mandatory) to give the Project a unique name. Snyk recommends using the image name and tag without the repository. For example `base-python:3.9.2_2021110408`.
     * Example of a Snyk CLI command: \
       `snyk container monitor somecompany/base-python:3.9.2_2021110408 --project-name=base-python:3.9.2_2021110408 --org=ORGANIZATION_ID/ORGANIZATION_NAME`

#### Mark the created Project as a custom base image

{% hint style="info" %}
You can also do this through the [Snyk REST API](https://apidocs.snyk.io/#tag--Custom-Base-Images).
{% endhint %}

1. From the Web UI, open the Project of the imported base image.
2. On the Project page, navigate to **Settings**.
3. Enable **Custom Base Image**. This allows Snyk to recognize this image as a base image in other Projects.
4. If you want Snyk to use this image as a source to determine the best upgrade path, enable **Include in recommendations**.

<figure><img src="../../../../.gitbook/assets/enable_CBIR.png" alt="Enable Custom Base Image recommendation"><figcaption><p>Enable Custom Base Image recommendation</p></figcaption></figure>

5. Click **Save changes**.

If this is the first Project you marked as a custom base image in the image's repository, you must set a versioning schema. For more information, see [Versioning schemas for custom base images](versioning-schema-for-custom-base-images.md). If a schema already exists for the image's repository, it is displayed after saving the changes.

{% hint style="warning" %}
Any changes you make to the versioning schema apply to all custom base images in the repository.
{% endhint %}

You can also edit the **Custom base Image** settings for Projects that you have already marked as custom base images.

### Receive custom base image recommendations

The applications team responsible for using pre-built custom base images and adding additional layers on top of the pre-built images for their applications performs this step.

First, import an image to a new Snyk Project. Ensure that the Project is in the same Group as the custom images. You can import an image either from the CLI or the Web UI.

{% hint style="info" %}
If the same image is scanned from both the CLI and Web UI, Snyk creates two Projects and monitors both.
{% endhint %}

#### Import a new image through the CLI and set the Dockerfile

Here is an example command for importing an image through the CLI:

`snyk container monitor somecompany/app-python:2021110408 --file=path/to/Dockerfile`

Use `--file` (mandatory) to specify the path to the Dockerfile.

#### Use a previously imported image and set the Dockerfile

You can modify previously imported container Projects in order to attach a Dockerfile.

On the Project page of the application image, navigate to **Settings** and configure the Dockerfile by clicking **Configure Dockerfile** and selecting your source control system from the dropdown.

<figure><img src="../../../../.gitbook/assets/configure_project_dockerfile (1).png" alt=""><figcaption><p>Configure the Project Dockerfile</p></figcaption></figure>

Choose the Dockerfile repository and add the path to your Dockerfile. Click **Update Dockerfile**.

<figure><img src="https://lh5.googleusercontent.com/4cyspvfpv1ZA-4rmhU7DzngLigf8c6rgEu5d7wHiiy7QMbIHy8Qw6qqS0VLEAEYpAfBADISvvQAyCkGqeoBgKxexDxzVPBJvNzB44MSvBzGlPd0NNuWrZyv_73NggOYlSjZCER0z" alt="Configure the path to your Dockerfile"><figcaption><p>Configure the path to your Dockerfile</p></figcaption></figure>

#### See Custom Base Image Recommendations

Next, navigate to the **Project** page to see the recommendations for the image.

{% hint style="info" %}
You may need to retest a project when setting the Dockerfile for an existing project.
{% endhint %}

<figure><img src="https://lh5.googleusercontent.com/G--7GkeQ6i0bwTWE1tdC_Gg5d727JdQQfclEQ1n2opt5vtRDjT2FBChFpSZBD9V1TleoLigSzhtEERg4tfVI6yIua5Q5nGeNycmR93BYCG1DsiREvhNWKtFdZ4imJZvC1ypmDKOI" alt="Custom Base Image Recommendations example"><figcaption><p>Example of Custom Base Image Recommendations</p></figcaption></figure>

## Known limitations of CBIR

* When you scan an application image, custom base image recommendations do not appear unless you attach the Dockerfile to the scanned Project.
* The image's registry is ignored when recommendations are given for custom base images. When showing recommendations and fix PRs, images with the same repository but different registries are treated as coming from the same registry (the current base image's registry).
