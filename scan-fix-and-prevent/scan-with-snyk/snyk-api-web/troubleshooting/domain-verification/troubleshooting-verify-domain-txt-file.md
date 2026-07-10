---
description: How to troubleshoot domain verification with a TXT file in Snyk API and Web
---

# Troubleshoot domain verification with TXT file

To run full scans, you must verify domain ownership. Visit [Verify domain ownership](../../configure-targets/verify-domain-ownership/) to learn why verification is required.

## The problem

Domain verification using a TXT file fails with the error: `Token file not found.`

## Troubleshoot the problem

Work through the following steps to identify possible causes and solutions.

### Step 1: Test direct access

Test if the TXT file can be accessed directly.

**Using a browser:**

1.  Open a browser and navigate to:

    ```
    https://<domain_to_verify>/<your_file_name>.txt
    ```

    For example:

    ```
    https://example.com/00000000-0000-0000-0000-000000000000.txt
    ```

**Using curl:**

```bash
curl -i https://<domain_to_verify>/<your_file_name>.txt
```

For example:

```bash
curl -i https://example.com/00000000-0000-0000-0000-000000000000.txt
```

If the request fails with a **404 - Not Found** error, the file is not directly accessible, and Snyk API & Web cannot verify the domain.

Check the following causes and solutions:

| Cause                                                                                  | Solution                                                                                                                                               |
| -------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| The TXT file does not exist in the server's root folder.                               | Ensure the file is in the server's root folder. Visit [Verify with TXT file](../../configure-targets/verify-domain-ownership/verify-with-txt-file.md). |
| The TXT file is in the server's root folder but does not have the correct permissions. | Ensure the file has the correct permissions for Snyk API & Web to access it.                                                                           |

### Step 2: Test redirects

Test if there are redirects when accessing the TXT file.

**Using a browser:**

1.  Open a browser and navigate to:

    ```
    https://<domain_to_verify>/<your_file_name>.txt
    ```

**Using curl:**

```bash
curl -i https://<domain_to_verify>/<your_file_name>.txt
```

If there is a redirect, the TXT file is fetched from another URL, and Snyk API & Web cannot verify the domain.

For example, `https://example.com/00000000-0000-0000-0000-000000000000.txt` could redirect to `https://www.example.com/00000000-0000-0000-0000-000000000000.txt`.

Check the following cause and solution:

| Cause                                        | Solution                                                                                                                                     |
| -------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| The file's URL is redirected to another URL. | Create a target with the domain of the redirected URL. In the preceding example, create a target for `www.example.com` instead of `example.com`. |

### Step 3: Test blockers

Test if human checks (for example, a CAPTCHA) or Web Application Firewalls (WAFs) are blocking access to the TXT file.

1. Open a browser in incognito mode.
2.  Navigate to:

    ```
    https://<domain_to_verify>/<your_file_name>.txt
    ```

If a human check appears or a WAF blocks access to the TXT file, Snyk API & Web cannot verify the domain.

Check the following causes and solutions:

| Cause                                         | Solution                                                                                                                                                           |
| --------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| A human check is blocking access to the file. | Remove the human check when accessing the TXT file.                                                                                                                |
| A WAF is blocking access to the file.         | Add Snyk API & Web IPs to the WAF's allowlist. Visit [Configure IPs in WAFs](../../start-scanning/overview-scan-access-and-connectivity/configure-ips-in-wafs.md). |

After following these steps and applying the solutions, you can verify the domain using a TXT file.

For more information, visit [Verify with TXT file](../../configure-targets/verify-domain-ownership/verify-with-txt-file.md).
