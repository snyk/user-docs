# How the Snyk Controller handles your data

After you install the Snyk Controller in your Kubernetes cluster, it pulls images from your container registries:

1. The Snyk Controller has read-only access to the Kubernetes workloads and container registries.
2. The Snyk Controller pulls images to disk, scans them, and then deletes the images once the scanning completes.
3. The Snyk Controller collects only the minimum relevant information in order to perform vulnerability analysis, specifically the list of dependencies in the image and the workload metadata.
4. The data that the Snyk Controller collects is stored in Snyk's backend, but it is deleted after eight days. Below is the data that Snyk Controller collects:
   * Workload metadata (Kubernetes configuration):
     * Labels, annotations
     * PodSpec
   * Image scan results and image metadata:
     * List of OS packages
     * Application dependencies
5. The Snyk Controller's scanning result is available for import in the Snyk Web UI after some time between a few minutes and a few hours. The length of the delay depends on multiple factors:
   * The Snyk Controller doing the first scan of the cluster
   * The size of the Kubernetes clusters and workloads
   * How the customer uses manual or auto-import of the Kubernetes workload
   * The size of the images
   * The size of scanning queues
   * The network speed.

{% hint style="warning" %}
Snyk highly recommends NOT storing sensitive data in plain text as an environment variable in the container, for example, password, authentication token, and SSH key. Alternatively, you can store the sensitive data in a Secret, mount it as a Volume, and access the information from there.
{% endhint %}
