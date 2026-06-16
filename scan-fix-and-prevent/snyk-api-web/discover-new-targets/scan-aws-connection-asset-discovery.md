# Scan an AWS connection for asset discovery

Organizations often lack visibility into all their assets (web pages and APIs), which can lead to overlooked vulnerabilities and inadvertent security exposure. With Snyk API & Web asset discovery, you can identify your organization's assets and protect them before they become a liability.

Scanning an AWS Route 53 connection for asset discovery involves two steps:

1. Obtain AWS access key and secret access key.
1. Add the AWS Route 53 connection.

## Obtain AWS access key and secret access key

To add an AWS Route 53 connection, you need an AWS access key and secret access key. To obtain them, follow these steps:

1. Sign in to the **AWS Management Console** with your credentials.
1. Navigate to the **Identity and Access Management** (IAM) service:
   1. Type **iam** in the search box.
   1. Select the **IAM** service.

   <figure><img src="../../../.gitbook/assets/scan-aws-connection-asset-discovery-iam-service.png" alt="AWS Management Console search for IAM service"></figure>

1. In the **Users** section, click **Create User**.

   <figure><img src="../../../.gitbook/assets/scan-aws-connection-asset-discovery-create-user.png" alt="IAM Users section with Create User button"></figure>

1. Type the **User name** (in this example, it is `test-user`) and click **Next**.

   <figure><img src="../../../.gitbook/assets/scan-aws-connection-asset-discovery-user-name.png" alt="Create user form with user name field"></figure>

1. Select **Attach policies directly**, choose the desired policies to apply to the user, and click **Next**.

   <figure><img src="../../../.gitbook/assets/scan-aws-connection-asset-discovery-attach-policies.png" alt="Permissions page with attach policies directly option"></figure>

   If you have not created the policy yet, click **Create policy** and use the example below to configure a policy to allow listing zones and resources from Route 53. After that, refresh the **Permissions policies** list.

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

1. Review the user settings and click **Create user**.

   <figure><img src="../../../.gitbook/assets/scan-aws-connection-asset-discovery-review-user.png" alt="Review and create user page"></figure>

1. With the user created, click the user name in the list to view its details.

   <figure><img src="../../../.gitbook/assets/scan-aws-connection-asset-discovery-user-list.png" alt="IAM users list showing the newly created user"></figure>

1. Select the **Security credentials** tab and click **Create access key**.

   <figure><img src="../../../.gitbook/assets/scan-aws-connection-asset-discovery-create-access-key.png" alt="Security credentials tab with Create access key button"></figure>

1. On the **Access key best practices and alternatives** screen, select **Other** and click **Next**.

   <figure><img src="../../../.gitbook/assets/scan-aws-connection-asset-discovery-access-key-use-case.png" alt="Access key best practices screen with Other option selected"></figure>

1. Type a tag description if needed and click **Create access key**.

   <figure><img src="../../../.gitbook/assets/scan-aws-connection-asset-discovery-tag-description.png" alt="Optional tag description field before creating access key"></figure>

1. On the next screen, you can retrieve the **Access key** and **Secret access key**.

   <figure><img src="../../../.gitbook/assets/scan-aws-connection-asset-discovery-retrieve-keys.png" alt="Access key and secret access key displayed for retrieval"></figure>

## Add the AWS connection

In Snyk API & Web, add the AWS connection for asset discovery:

1. Select the **Discovery** page.
1. Click **Add source** to open the configuration modal.
1. Select **Connect with AWS Route 53** and click **Next**.
1. On the next screen, enter the **AWS access key** and **AWS secret access key** with the values obtained in the previous step and click **Connect**.

After successfully connecting with AWS, Snyk API & Web starts running regular discovery scans automatically on your account. In Snyk API & Web, check the **Discovery** page. Once the asset discovery is finished, all newly found assets are added to the list. At the top of the page, information about the number of newly found assets is displayed. Click on it to filter the list.
