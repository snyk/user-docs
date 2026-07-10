---
description: How to scan an AWS connection for asset discovery in Snyk API and Web
---

# Scan an AWS connection for asset discovery

Organizations often lack visibility into all their assets (web pages and APIs), which can lead to overlooked vulnerabilities and inadvertent security exposure. With Snyk API & Web asset discovery, you can identify your organization's assets and protect them before they become a liability.

Scanning an AWS Route 53 connection for asset discovery involves two steps:

1. Obtain AWS access key and secret access key.
2. Add the AWS Route 53 connection.

## Obtain AWS access key and secret access key

To add an AWS Route 53 connection, you need an AWS access key and secret access key. To obtain them, follow these steps:

1. Log in to the **AWS Management Console** with your credentials.
2. Navigate to the **Identity and Access Management** (IAM) service:
   1. Type **iam** in the search box.
   2. Select the **IAM** service.
3. In the **Users** section, click **Create User**.
4. Type the **User name** (in this example, it is `test-user`) and click **Next**.
5.  Select **Attach policies directly**, select the desired policies to apply to the user, and click **Next**.

    If you have not created the policy yet, click **Create policy** and use the following example to configure a policy to allow listing zones and resources from Route 53. After that, refresh the **Permissions policies** list.

    ```json
    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Action": [
                    "route53:ListHostedZones",
                    "route53:ListResourceRecordSets"
                ],
                "Resource": "*"
            }
        ]
    }
    ```
6. Review the user settings and click **Create user**.
7. With the user created, click the user name in the list to view its details.
8. Select the **Security credentials** tab and click **Create access key**.
9. On the **Access key best practices and alternatives** screen, select **Other** and click **Next**.
10. Type a tag description if needed and click **Create access key**.
11. On the next screen, you can retrieve the **Access key** and **Secret access key**.

## Add the AWS connection

In Snyk API & Web, add the AWS connection for asset discovery:

1. Select the **Discovery** page.
2. Click **Add source** to open the configuration modal.
3. Select **Connect with AWS Route 53** and click **Next**.
4. On the next screen, enter the **AWS access key** and **AWS secret access key** with the values obtained in the previous step and click **Connect**.

After successfully connecting with AWS, Snyk starts running regular discovery scans automatically on your account. In Snyk, check the **Discovery** page. After the asset discovery finishes, Snyk adds all newly found assets to the list. At the top of the page, Snyk displays information about the number of newly found assets. Click it to filter the list.
