# AWS Lambda setup: add security through an environment variable

For security reasons the script that you created uses an environment variable: `hmac_verification` with a shared secret to validate the payload is coming from Snyk and has not been tampered with.

Follow these steps to add security through an environment variable:

1. Go to the **Configuration** tab in the AWS Lambda function.
2. Click **Environment variables**.
3. Add the new environment variable.
4. Use `hmac_verification` as your key.
5.  Enter your preferred secret as the key value. Store this secret somewhere safe for use again later.\


    <figure><img src="https://lh4.googleusercontent.com/eXXBAsVL2kDNpr9fDt_PErj9x0z7nBa-KywuWXJ0nGpuwwEiBiu8p0wFJLMacewmkRnYfrWSMzXqzhHAhRjifx-uEJF_BZm5Y0SazSMw60zKq8JOsLiGpqb7Risfr5zVBoBI7uiOJyMp_7G_HCajTB_vpIEVJotV4u1cJ4yO_t2wEi1jEARxk2sLjQ" alt="AWS Lambda function add environment variable interface"><figcaption><p>AWS Lambda function add environment variable interface</p></figcaption></figure>
6. For added security, consider using Secrets Manager or Parameter Store for the shared secret.
