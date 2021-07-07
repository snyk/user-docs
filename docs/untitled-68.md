# Configure integration for ACR

* [ Container security with ACR: integrate and test](/hc/en-us/articles/360003946957-Container-security-with-ACR-integrate-and-test)
* [ Configure integration for ACR](/hc/en-us/articles/360003915958-Configure-integration-for-ACR)
* [ Add images to Snyk from ACR](/hc/en-us/articles/360003915978-Add-images-to-Snyk-from-ACR)

##  Configure integration for ACR

Enable integration between an ACR registry and a Snyk organization, and start managing your vulnerabilities. To integrate with multiple registries, create a unique organization for each one.

**Steps:**

1. Access your ACR account and retrieve unique service principal credentials for use by Snyk with the AcrPull role. For help doing this, see the [ACR documentation](https://docs.microsoft.com/en-us/azure/container-registry/container-registry-auth-service-principal).
2. Log in to your Snyk account.
3. Navigate to Integrations from the menu bar at the top, find and click the ACR option:
4. The ACR configuration page in the Settings area loads:
5. Enter the user name, password and container registry name \(myregistry.azurecr.io\) that you received when you generated a service principal for this integration.
6. Click Save.

   Snyk tests the connection values and the page reloads, now displaying ACR integration information. A confirmation message that the details were saved also appears in green at the top of the screen. In addition, if the connection to Azure failed, notification appears under the Connected to ACR section.

