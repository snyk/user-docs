# Versioning Schema options

A versioning schema is a system for identifying and organizing different versions of a Project. It is used to track changes and updates to the Project over time and to help users identify which version they are using.

A versioning schema typically consists of a series of numbers or labels that are incremented to reflect the progression of versions. For example, a versioning schema might use a series of numbers, such as "1.0", "1.1", "2.0", and so on, to indicate major and minor releases of a product.

A consistent and well-defined versioning schema helps users and tools understand and track the development of a Project.

## **SemVer**&#x20;

When Snyk scans the following images and marks them as custom base images in the order shown, the SemVer logic is:

1. `developer-java/oracle-jre-rhel7/8e32:1.8.0`
2. `developer-java/oracle-jre-rhel7/8e32:1.9.2`
3. `developer-java/oracle-jre-rhel7/8e32:1.7.0`

Snyk recommends the second image, as it is the newest image **based on the semantic versioning of the tag**.

If Snyk cannot find a [standard semantic versioning schema](https://semver.org/) in the tag, the recommendation is the last image that was marked as a custom base image. In this example, the last image marked is the third image. The last image marked is determined in the **timestamp logic**.

Another example of the **timestamp logic** follows. When you scan the following images and mark them in the following order as custom base images using the Project settings in the Snyk interface, the timestamp logic is:

1. `developer-java/oracle-jre-rhel7/8e32:1.8.0_2021021008`
2. `developer-java/oracle-jre-rhel7/8e32:1.8.0_2021022508`
3. `developer-java/oracle-jre-rhel7/8e32:1.8.0_2021031708`

Snyk recommends the third image, as it was the last image marked as a custom base image.

## Single Selection&#x20;

This versioning schema allows only a single image to be given as a recommendation.

The specific image that will be recommended to the entire repository is configurable through the API by passing a boolean that will indicate whether this image should be the current recommendation.

Since only a single image can be the recommended image, any other image previously selected will be unselected automatically.

## Additional versioning schema options

For additional options, see [Custom Versioning Schema for Custom Base Images](custom-versioning-schema-for-custom-base-images.md).
