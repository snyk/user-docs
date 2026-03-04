# Image repository, tab, and Image Pull Secret

You can choose to use your own container registry and tag instead of the public images by customizing the `values.yaml` file to specify your container registry uri and tag.

If your container registry requires an image pull secret, you can specify an image secret.

Note that the Image Pull Secret is NOT created by the chart; rather, the image Pull Secret is expected to be present on your cluster.
