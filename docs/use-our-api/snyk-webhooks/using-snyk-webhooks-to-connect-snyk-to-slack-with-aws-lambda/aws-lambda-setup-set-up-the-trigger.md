# AWS Lambda setup: set up the trigger

The goal in this setup to have the Lambda function triggered by the Snyk webhook.

To do this , use the AWS API Gateway to trigger the Lambda function every time a new event is received.

Follow these steps to add the AWS API Gateway to the function:

1. Navigate to the function and click **Add trigger**.
2. Select **API Gateway** as the trigger.
3. Configure the API Gateway as follows.\
   Select **Create a new API**.\
   Select **REST API**.\
   Select **API key** as **Security**.\
   In the **Additional settings**, enter the **API name** and select **default** as the **Deployment stage**.
4.  Verify that the screen looks like this.\


    <figure><img src="https://lh4.googleusercontent.com/MsJp0QG0nnTA4xr4uwH1L3cpSGMqEuDY-LnSpJUqIZKzpq8ZUL332aBGEYJOa2pjZ466lqJGxDYluJo0-XRzd7AHSkPdEFKEbG-AoCmQVFL79DBWdJNXI3mqBVOmRX242Xu9jEl8OBb2ovj2xLPliixrD3xrQVukrYyH2VSWKvk-RkZ_LYuEaCV-CA" alt="AWS Lambda function Add trigger interface"><figcaption><p>AWS Lambda function Add trigger interface</p></figcaption></figure>
5. Continue by following the steps to set up the AWS API Gateway.
