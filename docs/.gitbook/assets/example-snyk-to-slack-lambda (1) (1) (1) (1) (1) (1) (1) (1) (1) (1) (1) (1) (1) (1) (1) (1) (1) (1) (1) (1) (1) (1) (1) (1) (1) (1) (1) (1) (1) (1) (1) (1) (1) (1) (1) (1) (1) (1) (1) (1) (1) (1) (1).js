const crypto = require('crypto');
const axios = require('axios');

//update with your webhook url
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