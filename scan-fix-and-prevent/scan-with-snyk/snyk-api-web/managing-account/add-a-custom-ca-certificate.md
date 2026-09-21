---
description: How to configure a custom CA certificate
---

# Add a custom CA certificate

## Trust a custom CA certificate

If your organization issues web certificates from an internal Certificate Authority (CA) rather than a publicly trusted one, you can upload your internal CA once at the account level so Snyk validates certificates against it, just as it validates against public CAs.

Snyk still reports certificates that do not chain to either the public trust store or your uploaded CA as untrusted.

### Configure the trusted CA

Uploading, replacing, or removing the trusted CA requires the **Account Settings** permission. Any user who can open the Scan Settings page can view the configured certificate details, even without this permission.

1. From the side menu in your Snyk API & Web account, navigate to **Settings > Scan Settings**.
2. Locate the **TRUSTED CA CERTIFICATE** module.
3. Click **Upload** and select a Privacy-Enhanced Mail (PEM) file containing your CA certificate.
   * If your organization uses a multi-tier public key infrastructure (PKI) with a Root CA and one or more Intermediate CAs, you can include any number of these certificates in the PEM file. If your Root CA is internal and not publicly trusted, make sure it is included in the uploaded PEM file, so Snyk can validate trust.
4. Verify that the certificate details Snyk extracts and displays match your CA, then confirm the upload:
   * **Common Name**
   * **Issuer**
   * **SHA-256 fingerprint**
   * **Expiry date**

After you upload the certificate, every scan across every target in the account validates server certificates against both your CA and the public trust store.

### Replace or remove the trusted CA

* To replace it: click **Upload** and select a new PEM file. This replaces the existing configuration entirely.
* To remove it: click **Remove**. This returns the account to the default public trust store.&#x20;

The [account audit log](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-api-web/managing-account/generate-and-use-audit-log-reports) records both actions, including the user and timestamp, providing an audit trail for untrusted-certificate findings.

### Certificate expiry

* If the file you upload contains a certificate that has already expired, Snyk rejects the upload and explains why. An expired CA cannot establish trust, so Snyk does not store it.
* If a certificate you have already uploaded expires while in use, Snyk flags it in the module. Upload a renewed certificate to restore trust.

### Troubleshooting

| Scenario                                                       | What you will see                                                                                                       | What to do                                                                                |
| -------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| The uploaded file is not valid PEM, or contains no certificate | "This file does not contain a valid certificate. Upload a PEM file containing your CA certificate."                     | Re-export your CA certificate in PEM format and upload again.                             |
| The uploaded certificate has already expired                   | "This certificate expired on \[date] and cannot be used to establish trust."                                            | Upload your current CA certificate.                                                       |
| A stored certificate has expired                               | The certificate row is flagged as expired, with a message that affected findings will return.                           | Upload the renewed CA certificate.                                                        |
| You do not have permission to make changes                     | "You do not have permission to change these settings. Please contact the account owner if you need to change anything." | Ask an account owner, or a user with the Account Settings permission, to make the change. |

