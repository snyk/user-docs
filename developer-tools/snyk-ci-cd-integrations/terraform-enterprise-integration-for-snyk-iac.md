# Terraform Enterprise integration for Snyk IaC

## Terraform Enterprise overview

[Terraform Enterprise](https://www.terraform.io/enterprise) (TFE) by HashiCorp is an offering that provides a private instance of the [Terraform Cloud](https://cloud.hashicorp.com/products/terraform) application, with no resource limits and with additional enterprise-grade architectural features like audit logging and SAML single-sign-on.

## **Snyk integration with Terraform Enterprise overview**

Snyk integration with Terraform Enterprise works the same way as the Snyk integration for Terraform Cloud. See the [Snyk integration for Terraform Cloud](terraform-cloud-integration-for-snyk-iac-using-run-tasks/) page for details about how to set up the integration.

## Network requirements for Terraform Enterprise integration

Snyk integration for Terraform Enterprise relies on network connectivity between your Terraform Enterprise instance and the Snyk platform. If you have tried to set up an integration as explained for [Snyk integration for Terraform Cloud](terraform-cloud-integration-for-snyk-iac-using-run-tasks/) and have been unable to get the integration to work, the following steps can help you identify the problem:

* To check whether a connection from your Terraform Enterprise instance to the Snyk API can be established:
  * Log on to your Terraform Enterprise server.
  * Make an HTTP request to a Snyk API endpoint; for example, you could use curl to initiate an HTTP request.
  * Even if you receive a 401/Unauthorized response from a Snyk API endpoint, that is an acceptable response; you are checking only basic network connectivity.
* You must also verify that Snyk servers can reach your Terraform Enterprise instance. Ensure that your network configuration (firewalls and so on) allows receiving network traffic from the Snyk platform.
