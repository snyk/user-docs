# Custom Base Image Recommendations

{% hint style="info" %}
This feature is currently in Beta. Contact your Snyk account team if you are interested in participating.
{% endhint %}

## **Overview of Custom Base Image Recommendations (CBIR)**

When scanning a container image, Snyk provides recommendations based on the base image detected. Because Snyk precomputes recommendations only for Docker Official Images, these recommendations apply only to a subset of images.

Customers often maintain their own internal, customized base images, built on top of Docker Official Images or other upstream images. These are provided as a service to a wider set of development teams, for example, `somecompany/java:v1.2.4`.

Through the Custom Base Image Recommendation feature, Snyk can recommend an image upgrade from a pool of the customer’s internal images. This allows teams to be notified of newer and more secure versions of their internal base images.

## How CBIR works

* _**For a customer to use the Custom Base Image Recommendations feature**_, Snyk needs to enable it for each **Organization** that wants to be able to mark images as custom base images, for example, the platform team.
  * This means that every user in the Organization (platform team in this example) will be able to mark images as custom base images in the Project settings.
  * Later, Projects in the same **Group** as the Organization (platform team in this example) will be able to receive custom base image recommendations.
* The current _**logic**_ is: for the same image repository (same repo and name), Snyk recommends the newest image based on the **semantic versioning of the image tag**. If Snyk is unable to find a [standard semantic versioning schema](https://semver.org/) in the tag, the recommendation is the last image that was marked as a custom base image based on the timestamp of marking. Refer to the [steps in enabling the feature](./#how-to-enable-custom-base-image-recommendations) for more information.
* The user must specify a Dockerfile in the Project in order to receive custom base image recommendations. Refer to [How to enable Custom Base Image Recommendations](./#how-to-enable-custom-base-image-recommendations) for more information.
* All custom base image recommendations are considered minor upgrades, regardless of the image tag.
* Automatic fix PRs are supported for custom base image recommendations. If you are not using the latest version of the base image, then immediately after image import Snyk automatically issues a fix pull request against your Dockerfile to upgrade to the latest custom base image version available.
* In order for Snyk to identify whether a Project is using a custom base image, the same custom base image must be imported and marked as such in the Project settings.

## **How to enable CBIR**

Custom base images can be configured by following the steps in this section and through the [Snyk API](https://apidocs.snyk.io/#tag--Custom-Base-Images), which is the recommended method.

### Configure an image as a custom base image

The team that is responsible for creating and maintaining custom base images for the Organization performs this step. This is the platform team in this example.

1. Ask Snyk to enable the feature for the Organization being used by the platform team.
2. Create a custom base image.
3. Import the image to a Snyk Project either:
   1. Through the Web UI: Import an image into Snyk using a container registry.
   2.  Or through the CLI: Use `--file` (optional) to specify the path to the Dockerfile, and `--project-name` (mandatory) to give the Project a **unique** name. Snyk recommends using the image **name and tag, without the repo**.\
       Example image name: `oracle-jre-rhel7/8e32:1.8.0_2021022508`)

       Example Snyk CLI command: `snyk container monitor snykgoof/custom-base-python:3.9.2_2021110408 --file=path/to/Dockerfile.3.9.2 --project-name=custom-base-python:3.9.2_2021110408 --org=ORGANIZATION_ID/ORGANIZATION_NAME`
4. Mark the project as a custom base image.
   1.  Go to the **Settings** page for the Project.

       <figure><img src="../../../../.gitbook/assets/image (188) (1) (1) (1).png" alt="Navigate to Settings page for the project"><figcaption><p>Navigate to Settings page for the project</p></figcaption></figure>
   2. Under **Custom Base Image Recommendation**, select **Treat as custom base image**.
   3. Click **Update image status**.
5. Mark whether the image should be eligible for recommendations.
   1. In **Custom Base Image Recommendation,** select the **Use in recommendations** checkbox.
   2. Click **Update image status**.
6. To test the feature, go through the preceding steps for at least two different images from the same repository in order to get recommendations.

### Receiving custom base image recommendations

The applications team responsible for using pre-built custom base images and adding additional layers on top of the pre-built images for their applications performs this step.

First, import an image to a new Snyk Project. Check that the Project is in the same **Group** as the custom images. You can import an image from the CLI or the Web UI.

{% hint style="info" %}
If the same image is scanned from both the CLI and UI, Snyk creates two Projects and monitors both.
{% endhint %}

#### Import an image through the CLI

The following is an example command: `snyk container monitor snykgoof/custom-base-python:3.9.2_2021110408 --file=path/to/Dockerfile.3.9.2`

Use --file (**mandatory**) to specify the path to the Dockerfile.

Use the -`-exclude-base-image-vulns` flag (**optional**) for the `snyk container test` command to _not_ show the base image vulnerabilities.

#### Import an image through the Web UI

Configure the Dockerfile through the Project settings (**mandatory**):

<figure><img src="https://lh5.googleusercontent.com/tPfU1mB9wZ-eSLTXHh5lRG58zh5xsnoTggeQ1xA7s7yShWoIZm4rfy4_qoE-aFGr4wYucMJrUebsmwri4Ba8B4bHZ5Nd4ax_qvv5vxdIJZbNAdH3JGI_uwhALj7U99bOS57s3xPI" alt="Configure the project&#x27;s Dockerfile"><figcaption><p>Configure the project's Dockerfile</p></figcaption></figure>

<figure><img src="https://lh5.googleusercontent.com/4cyspvfpv1ZA-4rmhU7DzngLigf8c6rgEu5d7wHiiy7QMbIHy8Qw6qqS0VLEAEYpAfBADISvvQAyCkGqeoBgKxexDxzVPBJvNzB44MSvBzGlPd0NNuWrZyv_73NggOYlSjZCER0z" alt="Configure the path to your Dockerfile"><figcaption><p>Configure the path to your Dockerfile</p></figcaption></figure>

#### Get Custom Base Image Recommendations

Next, navigate to the recommendations for the image.

<figure><img src="https://lh5.googleusercontent.com/G--7GkeQ6i0bwTWE1tdC_Gg5d727JdQQfclEQ1n2opt5vtRDjT2FBChFpSZBD9V1TleoLigSzhtEERg4tfVI6yIua5Q5nGeNycmR93BYCG1DsiREvhNWKtFdZ4imJZvC1ypmDKOI" alt="Custom Base Image Recommendations example"><figcaption><p>Custom Base Image Recommendations example</p></figcaption></figure>

## Known limitations of CBIR

* Marking an image as a custom base image is supported only through the UI and not through the API and CLI.
* Custom base image recommendations will not appear when you scan an image unless you attach the Dockerfile to the Project.
* The image's registry is ignored when recommendations are given for custom base images. Images with the same repository but different registries are treated as coming from the same registry (the current base image's registry) in showing recommendations and fix PRs.
* After a custom base image has been imported and marked, that custom base image Project should not be moved to a different Organization or Group within Snyk.

##
