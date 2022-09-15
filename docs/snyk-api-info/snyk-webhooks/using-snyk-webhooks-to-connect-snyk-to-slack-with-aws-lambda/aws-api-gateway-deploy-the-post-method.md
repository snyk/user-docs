# AWS API Gateway: deploy the POST method

Deploy with configured POST method so the AWS Lambda function can start receiving the information.

Follow these steps to deploy the POST method:

1. Go to the **Resources** tab.
2. Click **POST**.
3.  On the **Actions** tab, click **Deploy API**.\


    <figure><img src="https://lh3.googleusercontent.com/MVnbbBF4_quh1tD-sWln5t7RdNn6kui43IRi_KHshS-jKEDVkFmsf9IpAI7Ly1Eo6ZPLnVx3WTZJ13qJdTKPWm9vr2FU1ERamyAo7N1-687QeGswSozAvB9eo8oyafqCdoDxt4nlGSDZBoyh2zf6ONWZDnN656UyodXV07glWvxCfBlkfPf7Sz8HLg" alt="AWS API Gateway POST method Resources, Action tab with Deploy API selected"><figcaption><p>AWS API Gateway POST method Resources, Action tab with Deploy API selected</p></figcaption></figure>
4.  Select the **Deployment stage** to which you want to deploy the new API, in this case, the **default** stage.\


    <figure><img src="https://lh6.googleusercontent.com/xiLxfQ4yO5vb39TKW84JQe8X05sZ01stYMXtY9H8w-V2vad54nEtBI94mYQBUnGGMrmp0aEiMrn5OA9xtDnqH3BjS1UyrE0Bxsx6-Oui3XW5vxi15x0AN-rMZCWHgi2NEhNxOc-PkYbpFCJLn6n88wfDetGwi19ka0ZojM2cNLyEjeGPugScFtAcww" alt="AWS Gateway Deploy API to default stage"><figcaption><p>AWS Gateway Deploy API to default stage</p></figcaption></figure>
5. Navigate back to your Lambda function and In the Lambda trigger configuration, verify you see a new API endpoint.
6.  Copy the API endpoint from the API Gateway boxes (obscured) for use in setting up the Snyk webhook.

    <figure><img src="https://lh4.googleusercontent.com/EOoL3PCnKMj0HI6jkRdVsE44DwAcnFN8M8jM3Obp_FA5AXTryIHTMtGn66LlSTquVfH__0wVfjKV5bUTCxwgJzClgcdPqFTrtaq57NCd-eKBoSgFFHN49Fdqw8OsBLQai5pFsGQwGhcNpqIeto4fmXozicUeJ2A25wkh81HVmxrQH53IS-oEZZTlmQ" alt="AWS Lambda function trigger configuration showing new endpoint"><figcaption><p>AWS Lambda function trigger configuration showing new endpoint</p></figcaption></figure>
7. Now that the API endpoint is saved, set up the Snyk Webhook.\
