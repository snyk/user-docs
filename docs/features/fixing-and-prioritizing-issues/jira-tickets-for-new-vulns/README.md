---
description: https://github.com/snyk-tech-services/jira-tickets-for-new-vulns
---

# Jira-tickets-for-new-vulns

Jira-tickets-for-new-vulns is a tool that allows users to automate Jira ticket creation and assignment for new vulnerabilities.&#x20;

### Use Cases:

#### Automation of Jira Ticket creation&#x20;

* Security teams can automate the creation and assignment of new Jira tickets for new vulnerabilities
* The additional filters provide granularity and allow for flexibility in the creation and assignment of new Jira tickets

(NOTE: Currently, this is not bidirectional and tickets can be created from Snyk to Jira, but tickets cannot be removed from Snyk, nor does closing tickets in JIra automatically remove Jira tickets from Snyk)

### Example of using jira-tickets-for-new-vulns (Using Binary):

#### Goal:&#x20;

Use Jira-tickets-for-new-vulns to automate ticket creation for new vulnerabilities to Jira by running the script as a cronjob or scheduled task&#x20;

#### Pre-requisites:&#x20;

* Snyk Token (found [here](https://app.snyk.io/account))
* Docker

#### Steps:&#x20;

1.  Clone the github project for snyk-bulk

    `git clone`[`https://github.com/snyk-tech-services/jira-tickets-for-new-vulns/`](https://github.com/snyk-tech-services/jira-tickets-for-new-vulns/tree/develop)``
2.  Access the jira-tickets-for-new-vulns folder

    `cd jira-tickets-for-new-vulns`
3. Download the appropriate binary [here](https://github.com/snyk-tech-services/jira-tickets-for-new-vulns/releases) and place it in the jira-tickets-for-new-vulns directory
4.  Look at the Github project and add the desired parameters

    Mandatory Parameters: (Additional Optional Parameters available in Github)

    ```
    ./snyk-jira-sync-<yourplatform> 
        --orgID <SNYK_ORG_ID>                    // Snyk -> Settings
        --projectID <Snyk Project ID>            // Can find it in the URL of the project <UUID value> or via API in Snyk
        --token <API Token>                      // Snyk API Token
        --jiraProjectID <12345>                  // Jira project ID (number) found at https://your-domain.atlassian.net/rest/api/3/project and under the ID key, you will see the ProjectID value 
        --jiraTicketType <Ticket Type in jira>   // Jira ticket type  found in Projects -> Project Settings -> Issue types
    ```
5.  Run the script&#x20;

    Example: `./snyk-jira-sync-<binary> --orgID <orgID> --projectID <projectID> --token <Snyk API Token> --jiraProjectID <jiraProjectID> --jiraTicketType <Jira Ticket Type>`
6. For the use of extended options, please refer to the Github project.&#x20;
