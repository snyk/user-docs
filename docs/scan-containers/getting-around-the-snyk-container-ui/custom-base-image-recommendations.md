# Custom Base Image recommendations

{% hint style="info" %}
This feature is currently in Beta. Contact your Snyk account team if you are interested in participating.
{% endhint %}

## **Overview**

When scanning a container image, Snyk provides recommendations based on the base image detected. These recommendations apply only to a subset of images, as Snyk precomputes recommendations only for Docker Official Images.

Customers often maintain their own internal, customized base images, built on top of Docker Official Images or other upstream images. These are provided as a service to a wider set of development teams, for example, `somecompany/java:v1.2.4`.

Through the Custom Base Image Recommendation feature, Snyk can recommend an image upgrade from a pool of the customer’s internal images. This allows teams to be notified of newer and more secure versions of their internal base images.

## How Custom Base Image Recommendations work

* _**For a customer to use the Custom Base Image Recommendations feature**_, Snyk needs to enable it for each **Organization** that wants to be able to mark images as custom base images (for example, the platform team).
  * This means that every user in the Organization (platform team in this example) will be able to mark images as custom base images in the project settings.
  * Later, projects in the same **Group** as the organization (platform team in this example) will be able to receive custom base image recommendations.
* The current _**logic**_ is: for the same image family (same repo and name), Snyk recommends the newest image based on the **semantic versioning of the image tag**. If Snyk is unable to find a [standard semantic versioning schema](https://semver.org/) in the tag, the recommendation is the last image that was marked as a custom base image based on the timestamp of marking. Refer to the [steps in enabling the feature](custom-base-image-recommendations.md#how-to-enable-custom-base-image-recommendations) for more information.
* The user must specify a Dockerfile in the project in order to receive custom base image recommendations. Refer to the [steps in enabling the feature](custom-base-image-recommendations.md#how-to-enable-custom-base-image-recommendations) for more information.
* All custom base image recommendations are considered minor upgrades, regardless of the image tag.
* Automatic fix PRs are supported for custom base image recommendations. If the user is not using the latest version of the base image, then immediately after image import Snyk automatically issues a fix pull request against your Dockerfile, to upgrade to the latest custom base image version available.
* In order for Snyk to identify whether a project is using a custom base image, the same custom base image must be imported and marked as such in the project's settings.

## **How to enable Custom Base Image Recommendations**

### Configure an image as a custom base image

This step is done by the team that is responsible for creating and maintaining custom base images for the Organization, the platform team in these instructions.

1. Ask Snyk to enable the feature for the Organization being used by the platform team.
2. Create a custom base image.
3. Import the image to a Snyk’s project either:
   1. Through the Web UI: Import an image into Snyk using a container registry.
   2.  Or through the CLI: Use `--file` (optional) to specify the path to the Dockerfile, and `--project-name` (mandatory) to give the project a **unique** name. Snyk recommends using the image **name and tag, without the repo**.\
       Example image name: `oracle-jre-rhel7/8e32:1.8.0_2021022508`)

       Example Snyk CLI command: `snyk container monitor snykgoof/custom-base-python:3.9.2_2021110408 --file=path/to/Dockerfile.3.9.2 --project-name=custom-base-python:3.9.2_2021110408 --org=ORGANIZATION_ID/ORGANIZATION_NAME`
4. Mark the project as a custom base image.
   1.  Go to the **Settings** page for the project.

       <figure><img src="../../.gitbook/assets/image (188) (1) (1) (1) (1) (1).png" alt="Navigate to Settings page for the project"><figcaption><p>Navigate to Settings page for the project</p></figcaption></figure>
   2. Under **Custom Base Image Recommendation**, select **Treat as custom base image**.
   3. Click **Update image status**.
5. Mark whether the image should be eligible for recommendations.
   1. In **Custom Base Image Recommendation,** select the **Use in recommendations** checkbox.
   2. Click **Update image status**.
6. To test the feature, go through the preceding steps for at least two different images from the same repository in order to get recommendations.

### Receiving custom base image recommendations

This step is done by the applications team responsible for using pre-built custom base images and adding additional layers on top of the pre-built images for their applications.

First, import an image to a new Snyk project. Check that the project is in the same **Group** as the custom images. You can import an image from the CLI or the Web UI.

{% hint style="info" %}
If the same image is scanned from both the CLI and UI, Snyk creates two projects and monitor both.
{% endhint %}

#### Import an image through the CLI

The following is an example command: `snyk container monitor snykgoof/custom-base-python:3.9.2_2021110408 --file=path/to/Dockerfile.3.9.2`

Use --file (**mandatory**) to specify the path to the Dockerfile.

Use the -`-exclude-base-image-vulns` flag (**optional**) for the `snyk test` command to not show the base image vulnerabilities.

#### Import an image through the Web UI

Configure the Dockerfile through the project’s settings (**mandatory**):

<figure><img src="https://lh5.googleusercontent.com/tPfU1mB9wZ-eSLTXHh5lRG58zh5xsnoTggeQ1xA7s7yShWoIZm4rfy4_qoE-aFGr4wYucMJrUebsmwri4Ba8B4bHZ5Nd4ax_qvv5vxdIJZbNAdH3JGI_uwhALj7U99bOS57s3xPI" alt="Configure the project&#x27;s Dockerfile"><figcaption><p>Configure the project's Dockerfile</p></figcaption></figure>

<figure><img src="https://lh5.googleusercontent.com/4cyspvfpv1ZA-4rmhU7DzngLigf8c6rgEu5d7wHiiy7QMbIHy8Qw6qqS0VLEAEYpAfBADISvvQAyCkGqeoBgKxexDxzVPBJvNzB44MSvBzGlPd0NNuWrZyv_73NggOYlSjZCER0z" alt="Configure the path to your Dockerfile"><figcaption><p>Configure the path to your Dockerfile</p></figcaption></figure>

#### Get Custom Base Image Recommendations

Next, get recommendations for the image.

<figure><img src="https://lh5.googleusercontent.com/G--7GkeQ6i0bwTWE1tdC_Gg5d727JdQQfclEQ1n2opt5vtRDjT2FBChFpSZBD9V1TleoLigSzhtEERg4tfVI6yIua5Q5nGeNycmR93BYCG1DsiREvhNWKtFdZ4imJZvC1ypmDKOI" alt="Custom Base Image Recommendations example"><figcaption><p>Custom Base Image Recommendations example</p></figcaption></figure>

## Known limitations

* Marking an image as a custom base image is supported only through the UI, and not through the API and CLI.
* Custom base image recommendations will not appear when scanning an image unless the user attaches the Dockerfile to the project.
* The image’s registry is ignored when recommendations are given for custom base images. Images with the same repository but different registries are treated as coming from the same registry (the current base image’s registry) in showing recommendations and fix PRs.
* Once imported and marked, a custom base image project should not be moved to a different Organization or Group within Snyk.

## **Case study for Custom Base Image Recommendations**

1. Your company's platform team, responsible for creating and maintaining custom base images for the organization, scans and marks images in Snyk as custom base images.
2. Your company's application teams, using those pre-built custom base images and adding additional layers on top of the pre-built images for their applications, can get recommendations for upgrading to a newer internal version.

### **SemVer recommendation logic**

As an example, when scanning the following images and marking them (in the following order) as custom base images, the SemVer logic is:

1. `developer-java/oracle-jre-rhel7/8e32:1.8.0`
2. `developer-java/oracle-jre-rhel7/8e32:1.9.2`
3. `developer-java/oracle-jre-rhel7/8e32:1.7.0`

Snyk recommends the second image, as it is the newest image **based on the semantic versioning of the tag**.

If Snyk cannot find a [standard semantic versioning schema](https://semver.org/) in the tag, the recommendation is the last image that was marked as a custom base image (in this example, the third image), as determined in the Timestamp recommendation logix.

### **Timestamp recommendation logic**

As an example, when scanning the following images and marking them (in the following order) as custom base images through the Snyk interface, under project settings, the timestamp logic is:

1. `developer-java/oracle-jre-rhel7/8e32:1.8.0_2021021008`
2. `developer-java/oracle-jre-rhel7/8e32:1.8.0_2021022508`
3. `developer-java/oracle-jre-rhel7/8e32:1.8.0_2021031708`

Snyk recommends the third image, as it was last marked as a custom base image.
