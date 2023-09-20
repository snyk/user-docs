# How Snyk Controller handles your data

Once Snyk Controller is installed in your Kubernetes cluster, it will pull images from your container registries.

1. Snyk Controller has read-only access to the Kubernetes workloads and container registries.
2. Snyk Controller pulls images to disk, scans them, and then deletes the images once scanning is done.
3. Snyk Controller collects only the minimum relevant information in order to perform vulnerability analysis, that is, the list of dependencies in the image and the workload metadata.
4. The data that the Snyk Controller collects is stored in Snyk's backend, but it is deleted if it becomes more than eight days old. These are the data that Snyk Controller collects:
   1. Workload metadata (Kubernetes configuration)
      * Labels, annotations
      * PodSpec
   2. Policy for the workloads, specifically the policy for [Automatic import/deletion of Kubernetes workloads projects](../../../../scan-containers/kubernetes-workload-and-image-scanning/kubernetes-integration-features/automatic-import-deletion-of-kubernetes-workloads-projects/)
   3. Image scan results and image metadata
      * list of OS packages
      * application dependencies
5. Snyk Controller's scanning result will be updated in the Snyk UI after some time between a couple of minutes to a few hours because the result depends on the following factors:
   1. Snyk Controller doing the first scan of the cluster
   2. Size of the Kubernetes clusters and workloads
   3. Use by the customer of manual or auto import of the Kubernetes workload
   4. Size of the images
   5. The size of scanning queues
   6. The network speed

{% hint style="danger" %}
Snyk highly recommends **NOT** storing sensitive data in plain text as an environment variable in the container, for example, password, authentication token, and SSH key. Alternatively, you can store the sensitive data in a **Secret**, mount it as a **Volume**, and access the information from there.
{% endhint %}
