# AWS API Gateway: add the POST method to connect Snyk to Slack

The payload Slack will receive will have a message, so create a POST method that will receive the message, verify it is a valid message, and then send on to the AWS Lambda function.

Follow these steps to add the POST Method:

1. Navigate to the AWS API Gateway you have created
2. Click **Resources**.
3. To create the method, go to **Actions** -> **Create Method** -> **Post**.
4.  Configure the AWS API Gateway to work with the Lambda function you created by adding the Gateway in the adjacent Lambda function box:\
    Choose the **Lambda Function Integration type**.\
    Select **Default Timeout**.\


    <figure><img src="https://lh3.googleusercontent.com/3WjrkRdG1_TnfQ5w-9Ivg6J0xjic4znbfN3_76HX6quIGo5sydsEub8aMXrv9_MQsfAorYc4gUOwgIGK9JOpu0ysmI_dXFFtwlRk6LarMYu5xEgOHsJ2_9qHgKdw4Kf3MTFKX2v2EkBD5e80zC9tEZXUnFJnCfPLbaGCGv2h4omcpK10ntHdYvaVBA" alt="AWS Lambda function box"><figcaption><p>AWS Lambda function box</p></figcaption></figure>
5. In the **Resources**, lick the new **POST** method.
6.  Click **Integration Request** (top right on the AWS Gateway POST method execution screen).\


    <figure><img src="https://lh5.googleusercontent.com/_Prq2fJ7F-NE4jEiw1tqYIn0Bq-HTG0_wahTwkrod8zisAkjtKmL3O1Y0c8XEh2iYeibdkh1jWYR3V_jGvdWCbUEfE5LXd7I7cTovohFD81-NFGTvesu1jIFGKjRIWm88dAG_qcgKBQVMO7YrHvVcnERYFvr91I18K36137u2z4suVA_3P_xj8aCpQ" alt="AWS Gateway POST method execution"><figcaption><p>AWS Gateway POST method execution</p></figcaption></figure>
7. Scroll to the bottom and add a **Mapping Template** with **application/json Content-Type**. To the template add the following code:\
   `{`\
   &#x20;  `"method": "$context.httpMethod",` \
   &#x20;  `"body" : $input.json('$'),` \
   &#x20;  `"headers": {` \
   &#x20;        `#foreach($param in $input.params().header.keySet())` \
   &#x20;       `"$param": "$util.escapeJavaScript($input.params().header.get($param))"`\
   &#x20;       `#if($foreach.hasNext),#end #end` \
   &#x20;    `}` \
   `}`
8.  Verify the resulting display reflects these entries.\


    <figure><img src="https://lh6.googleusercontent.com/d0AynUJWVROWc0ff2EnY_NAT7kqkrvBThXMw8d9D0StX1KKoig7ol2uDqLoMOCt7UBP7C3RYiSUrcZlg9yglP08fVXf8WBxOvGHtV25hw5PsfQC_lWfoDJkl38kIaqBNxIdg_k7W4Qg5jvQ-faPp4ySOF5j15vWRxCEjxzvAIhsHpl3s3dE2lolJdg" alt="AWS API Gateway POST request Mapping Template with code"><figcaption><p>AWS API Gateway POST request Mapping Template with code</p></figcaption></figure>
