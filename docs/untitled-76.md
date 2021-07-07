# Configure integration for GCR

* [ Container security with GCR](/hc/en-us/articles/360003916098-Container-security-with-GCR)
* [ Enable permissions to access GCR](/hc/en-us/articles/360004191777-Enable-permissions-to-access-GCR)
* [ Configure integration for GCR](/hc/en-us/articles/360003916118-Configure-integration-for-GCR)

##  Configure integration for GCR

Configure integration from Snyk with your Google container registries to scan for vulnerabilities and fix security and license issues accordingly.

**Prerequisites:**

[Set up a new service account key](/hc/articles/360004191777#UUID-53c3d159-a436-9605-ec76-6bdc016fd824) from your GCR account.

**Steps:**

1. Navigate to your Snyk organization, and go to Integrations=&gt;GCR.
2. From the GCR hostname, enter the [registry storage region](https://cloud.google.com/container-registry/docs/pushing-and-pulling) for the images you want to scan in the format region.gcr.io. For example: gcr.io or asia.gcr.io.
3. Paste the entire contents of the JSON key file you created from Google into the JSON key file field in your Snyk account, like the following:
4. Click Save.

   Snyk checks the credentials and when successful, the page reloads with a notification that the connection succeeded.

