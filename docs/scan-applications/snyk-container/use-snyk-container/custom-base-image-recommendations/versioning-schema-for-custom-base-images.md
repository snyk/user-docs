# Versioning schema for custom base images

A versioning schema is a system used to identify and organize different versions of a Project. It is used to track changes and updates to the Project over time and to help users identify which version they are using.

A versioning schema typically consists of a series of numbers or labels that are incremented to reflect the progression of versions. For example, a versioning schema can use a series of numbers, such as "1.0", "1.1", "2.0", and so on, to indicate major and minor releases of a product.

A consistent and well-defined versioning schema helps users and tools understand and track the development of a Project.

## **Semantic Versioning (SemVer)**

In the below example, when Snyk scans the images and marks them as custom base images in the order shown, the SemVer logic is:

1. `developer-java/oracle-jre-rhel7/8e32:1.8.0`
2. `developer-java/oracle-jre-rhel7/8e32:1.9.2`
3. `developer-java/oracle-jre-rhel7/8e32:1.7.0`

Snyk recommends the second image, as it is the newest image based on the semantic versioning of the tag.

If Snyk cannot find a [standard semantic versioning schema](https://semver.org/) in the tag, it recommends the last image that was marked as a custom base image. In this example, the last image marked is the third image. The last image marked is determined in the timestamp logic.

Below is another example of the timestamp logic. When you scan the images in the example and mark them in the below order as custom base images using the Project settings in the Snyk interface, the timestamp logic is:

1. `developer-java/oracle-jre-rhel7/8e32:1.8.0_2021021008`
2. `developer-java/oracle-jre-rhel7/8e32:1.8.0_2021022508`
3. `developer-java/oracle-jre-rhel7/8e32:1.8.0_2021031708`

Snyk recommends the third image, as it was the last image marked as a custom base image.

## Single Selection&#x20;

This versioning schema allows only a single image to be given as a recommendation.

The specific image that is recommended to the entire repository is configurable through the API by passing a boolean that will indicate whether this image should be the current recommendation.

Since only a single image can be the recommended image, any other image previously selected will be automatically unselected.
