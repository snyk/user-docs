# Custom base images

If you build custom base images and store them in a registry for your development teams to select from, Snyk recommends that you introduce Snyk container tests as part of this workflow.

The team responsible for creating the custom base images can use container tests to select the most secure base images and identify layer-based fixes. This ensures that the base images selected by your development teams start with a secure base.

As development teams build their custom tools and packages into these images, they can run additional container tests to ensure their changes are secure before being pushed to production use.&#x20;

For more information, see [Use Custom Base Image Recommendations](../../../scan-with-snyk/snyk-container/use-snyk-container/use-custom-base-image-recommendations/).
