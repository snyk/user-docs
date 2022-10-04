# AWS Lambda setup: create Lambda function to connect Snyk to Slack

AWS Lambda functions are used to connect Snyk to Slack because these functions are an inexpensive and efficient way of running code triggered by events, for example when there is a new Snyk vulnerability.

**Note:** The Lambda function and API Gateway must be configured in the same region. You can check this on the top right of the AWS Console.

To create an AWS Lambda function, follow these steps:

1. Go to the AWS Console.
2. Navigate to Lambda.
3. Click **Create function**.
4. Choose **Node.js 16.x** for the **Runtime**.
5. Choose **X86\_64** for the **Architecture**.
6. Add or create a role with the default policy **AmazonAPIGatewayInvokeFullAccess** to interact with the AWS API Gateway.
7.  Verify that the AWS Console screen shows these entries:\


    <figure><img src="https://lh6.googleusercontent.com/xzJzGjfuzj0U27-pxcaIcrU-wBj8DTuEiQpivJZAnqRAO3rEPccx48l8KSZ5AE01BYJDwjJwkiFMR-Oj3ozWyG-CI20bwFtK_yjY9HKEoY0-4V4pa8l351JqrYdkK29va1x7BdlPoQ7N12SROjDQy3CmUQsDTtQ5lYOw3QvwoG1c1nDms-EFiQSElA" alt="AWS Console with entries to create a Lambda function"><figcaption><p>AWS Console with entries to create a Lambda function</p></figcaption></figure>
8.  Navigate to the Code Source and paste the code provided here as an example.

    ```javascript
    let slackWebhookUrl= "slackwebhookurl";

    //customised messaging to Slack with issue information, modify as needed
    async function messageSlack(message,snykProjectUrl,snykProjectName,snykIssuePackage,snykIssueUrl,snykIssueId,severity,snykIssuePriority) {
        
        //strings modified to avoid Axios/Slack errors 
        snykProjectUrl = snykProjectUrl.replace(/['"]+/g, '')
        snykProjectName = snykProjectName.replace(/['"]+/g, '')
        snykIssueUrl = snykIssueUrl.replace(/['"]+/g, '')
        snykIssueId = snykIssueId.replace(/['"]+/g, '')
        snykIssuePackage = snykIssuePackage.replace(/['"]+/g, '')
        severity = severity.replace(/['"]+/g, '')
        
        //construct message
        let payload = { "text": `${message}`,
                        "blocks": [
    		                {
                    			"type": "header",
                    			"text": {
                    				"type": "plain_text",
                    				"text": `${message}`,
                    			}
                    		},{
                    			"type": "section",
                    			"text": {
                    				"type": "mrkdwn",
                    				"text": "Snyk has found a new vulnerability in the project:\n*<"+snykProjectUrl+"|"+snykProjectName+">*"
                    			}
                    		},
                    		{
                    			"type": "divider"
                    		},
                    		{
                    			"type": "section",
                    			"fields": [
                    				{
                    					"type": "mrkdwn",
                    					"text": "*Package name:*\n"+snykIssuePackage
                    				},
                    				{
                    					"type": "mrkdwn",
                    					"text": "*Vulnerability:*\n<"+snykIssueUrl+"|"+snykIssueId+">"
                    				},
                    				{
                    					"type": "mrkdwn",
                    					"text": "*Severity:*\n"+severity
                    				},
                    				{
                    					"type": "mrkdwn",
                    					"text": "*Priority Score:*\n"+snykIssuePriority
                    				}
                    			]
                    		},
                    		{
                    			"type": "actions",
                    			"elements": [
                    				{
                    					"type": "button",
                    					"text": {
                    						"type": "plain_text",
                    						"text": "View in Snyk"
                    					},
                    					"style": "primary",
                    					"url": snykProjectUrl,
                    					"value": "browseUrl"
                    				}
                    			]
                    		}
    	               ]};
        
        //send message 
        const res = await axios.post(slackWebhookUrl, payload);
        const data = res.data;
        console.log(data);
    }

    exports.handler = async (event) => {
        // Securing integrity of payload, this can be moved to another Lambda function and called seperately for authentication 
        let response;
            
        const { hmac_verification } = process.env;
        const hmac = crypto.createHmac('sha256', hmac_verification);
        const buffer = JSON.stringify(event.body);

        hmac.update(buffer, 'utf8');
        const stored_signature = `sha256=${hmac.digest('hex')}`;

        let sent_signature=event.headers['X-Hub-Signature']; //per documentation
        sent_signature=event.headers['x-hub-signature']; //actually sent

        if(stored_signature !== sent_signature) {
        console.log('Integrity of request compromised, aborting');
            response = {
                statusCode: 403,
                body: JSON.stringify("Bad request"),
            };
        return response;
    }

        let snykbody = JSON.stringify(event.body);

        // If integrity is ok, verify that the webhook actually contains the project object, iterate and filter
        if(snykbody.indexOf("project") !== -1 && snykbody.indexOf("newIssues") !== -1){
            
            // Iterate through new issues
            var len = event.body['newIssues'].length;
            
            for(let x=0;x<len;x++){    
                // Get Severity
                let severity = JSON.stringify(event.body['newIssues'][x]['issueData']['severity']);
                // Filter
                if(severity.includes("high") || severity.includes("critical")){
                    
                    let snykProjectName = JSON.stringify(event.body['project'].name);
                    let snykProjectUrl = JSON.stringify(event.body['project'].browseUrl);
                    let snykIssueUrl = JSON.stringify(event.body['newIssues'][x]['issueData'].url);
                    let snykIssueId = JSON.stringify(event.body['newIssues'][x].id);
                    let snykIssuePackage = JSON.stringify(event.body['newIssues'][x].pkgName);
                    let snykIssuePriority = JSON.stringify(event.body['newIssues'][x]['priority'].score);
                    let message = "New Snyk Vulnerability";
                    
                    // Send the result to Slack
                    const result = await messageSlack(
                        message,snykProjectUrl,snykProjectName,snykIssuePackage,snykIssueUrl,snykIssueId,severity,snykIssuePriority
                    );
                } 
            }
        }
        //do nothing, or modify for any preferable action
        else{
            console.log('Valid webhook, but project missing or empty');
        }
        
        //respond to Snyk
        console.log('Valid webhook');
        response = {
            statusCode: 200,
            body: JSON.stringify('Success'),
        };
        
        return response;
    };
    ```


9.  Verify that the code you entered is displayed in the Code Source.\


    <figure><img src="https://lh3.googleusercontent.com/97qnO6V9xBXaf6dyO0hg41Y2vmeB1-0aPK-qskqTI-L2WII3d75zb4XsK6Mg5ljJUEdS7AGYJ5sQ5IoDHvzofkfK_gPId9e-XuBqEGkuWNxlIyL4IHu7-S8hrbGKnuyOehU2fjScDi0jazvuhWkADyFDGkkdAdzQGSEfWO30YGPJ9x4ocfwFXS5LfQ" alt="AWS code source display"><figcaption><p>AWS code source display</p></figcaption></figure>
10. In the code, modify the `slackWebhookUrl` to match your Slack webhook URL.
11. For more information on the script you have pasted, go to [Configure the AWS Lambda Script](configure-the-aws-lambda-script.md).
