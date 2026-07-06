# Verify domain with TXT file

Verify domain ownership by adding a `.txt` file to your website's root directory.

## Add the verification file

1. In Snyk API & Web, navigate to the **Targets** page.
1. Click the **Domains** tab.
1. Locate your domain and click **Verify**.
1. Select **File** as the verification method.
1. Note the **FILENAME** and **CONTENT** values provided.

## Create and upload the file

1. Create a `.txt` file with the specified **FILENAME**.
1. Set the file content to the specified **CONTENT**.
1. Upload the file to your website's root directory.

The file must be accessible at: `https://yourdomain.com/[filename].txt`

## Complete verification

1. Return to the verification dialog in Snyk API & Web.
1. Click **Verify**.

After successful verification, you can remove the `.txt` file from your website.

## Troubleshooting

### File not accessible

Snyk cannot access the file if it is:
- In the wrong location (must be in root directory)
- Protected by authentication or access restrictions
- Blocked by firewall rules

Test accessibility by manually accessing the file URL while logged out of your website.

### Hidden characters

Ensure the file content matches exactly with no additional hidden characters or formatting.

### Verification timing

If verification fails immediately after uploading the file, wait a few moments for the file to propagate across your web infrastructure and try again.
