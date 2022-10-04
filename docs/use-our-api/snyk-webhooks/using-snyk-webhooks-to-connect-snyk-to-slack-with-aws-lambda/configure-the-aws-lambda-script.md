# Configure the AWS Lambda script

You can configure the example code provided in multiple ways to get the information you want into Slack.

The key areas are the payload (where the Slack message is configured) and the filtering (where the Snyk information is processed).

Apart from these areas, you also configure the encryption and secret verification. Configuring these is beyond the scope of these instructions.

## Filter the Snyk payload

The following code filters the Snyk payload.

Add as code:

```javascript
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
```

The code does the following:

* Checks for a valid project and issue
* Iterates through the issues, checking if the severity is high or critical
* Passes all the information you specify to the Slack message. Specify the information that you want.

You can modify the filter to cover only a specific CWE for example, or to allow all vulnerabilities.

## Slack message format

The Slack payload is formatted in the messageSlack function as follows.

```javascript
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
```

Snyk used Slack’s built-in block builder to design the payload in the desired format. With the block builder you can configure the JSON of the payload to display more information like the CWE, add more interactions, and even use the ignore API from Snyk to create an ignore button.

You can find more information about the block builder ([Block Kit](https://api.slack.com/block-kit)) on the Slack website.



