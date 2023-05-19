# Configure integration for GCR

Configure integration from Snyk with your Google container registries to scan for vulnerabilities and fix security and license issues accordingly.

## **Prerequisites for configuring the GCR integration**

Set up a new service account key from your GCR account; see [Enable permissions to access GCR](enable-permissions-to-access-gcr.md).

## **Steps in configuring the GCR integration**

1. Navigate to your Snyk Organization, and go to **Integrations>GCR**.
2. From the GCR hostname, enter the [registry storage region](https://cloud.google.com/container-registry/docs/pushing-and-pulling) for the images you want to scan in the format region.gcr.io, for example, gcr.io or asia.gcr.io.
3. Paste the entire contents of the JSON key file you created from Google into the JSON key file field in your Snyk account, as shown in the following screenshot.
4. Click **Save**.

Snyk checks the credentials, and when the check is successful, the page reloads with a notification that the connection succeeded.

<figure><img src="../../../.gitbook/assets/uuid-47cf04cb-248e-5d0f-d35a-f36fbb624614-en.png" alt="GCR account credentials"><figcaption><p>GCR account credentials</p></figcaption></figure>

You can now [add images to Snyk from GCR](add-images-to-snyk-from-acr.md).
