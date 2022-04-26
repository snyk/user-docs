# How Snyk Controller handles your data

Once Snyk Controller is installed in your Kubernetes cluster, it will pull images from your container registries.

1. Snyk Controller has read-only access to the Kubernetes workloads and container registries.&#x20;
2. Snyk Controller pulls images to disk, scans them, and then deletes the images once scanning is done.
3. Snyk Controller collects only the minimum relevant information in order to perform vulnerability analysis, which is the list of dependencies in the image and the workload metadata.
4. The data that the Snyk Controller collects is stored in Snyk's backend but it is deleted if it becomes older than 8 days. These are the data that Snyk Controller collects:
   1. Workload metadata (Kubernetes configuration)
      * Labels, annotations
      * PodSpec
   2. Policy for the workloads, specifically the policy for [Automatic import/deletion of Kubernetes workloads projects](../kubernetes-integration-features/automatic-import-deletion-of-kubernetes-workloads-projects/)
   3. Image scan results and image metadata
      * list of OS packages
      * application dependencies

{% hint style="danger" %}
We highly recommend **NOT** to store sensitive data in plain text as an environment variable in the container, for example, password, authentication token and SSH key. Alternatively, you can store the sensitive data in a **Secret**, mount it as a **Volume**, and access the information from there.
{% endhint %}
