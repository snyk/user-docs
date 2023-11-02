# Custom base images

If you build custom base images and store them in a registry for your development teams to select from, then we recommend you introduce Snyk container tests as part of this workflow.

1. The team responsible for creating the custom base images can use container tests to select the most secure base images and identify layer-based fixes. This ensures that the base images selected by your development teams start with a secure base.
2. As development teams build their custom tools and packages into these images, they can run additional container tests to ensure their changes are secure before being pushed to production use.&#x20;

Also see [Use Custom Base Image Recommendations](../../../scan-with-snyk/snyk-container/use-snyk-container-from-the-web-ui/use-custom-base-image-recommendations/).
