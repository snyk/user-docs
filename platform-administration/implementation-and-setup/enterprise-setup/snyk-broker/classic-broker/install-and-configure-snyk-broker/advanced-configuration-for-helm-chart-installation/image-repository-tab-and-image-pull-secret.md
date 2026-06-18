# Image repository, tab, and Image Pull Secret

You can choose to use your own container registry and tag instead of the public images by customizing the `values.yaml` file to specify your container registry uri and tag.

If your container registry requires an image pull secret, you can specify an image secret.

Note that the chart does not create the Image Pull Secret. The Image Pull Secret must already be present on your cluster.
