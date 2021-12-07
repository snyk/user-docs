---
description: >-
  https://docs.snyk.io/products/snyk-infrastructure-as-code/detecting-infrastructure-as-code-files-using-a-broker
---

# Broker Example - How to import IaC files

This document will show and example of how to configure broker to allow specific IaC file types. Currently, Snyk supports Terraform, CloudFormation, ARM, Kubernetes, and Helm Charts. By default Snyk does not scan for these file types with an SCM integration using Broker, so it is necessary to read from an accept.json to allow the desired filetypes for Snyk to scan.&#x20;

#### Pre Requisites: Docker

### To configure a Broker to allow for IaC files to be scanned

1. Follow the steps on setting up the Broker configuration for you desired SCM
2. Prior to starting the broker container, within the broker project, go to the client-templates folder, select the desired SCM integration, and download the accept.json.sample and rename it to accept.json
3. Go to [https://docs.snyk.io/products/snyk-infrastructure-as-code/detecting-infrastructure-as-code-files-using-a-broker](https://docs.snyk.io/products/snyk-infrastructure-as-code/detecting-infrastructure-as-code-files-using-a-broker) and copy the code snippet for the desired SCM integration and add it to the accept.json. (Note: if you would like to scan Terraform files, you will need to add the .tf filetype to the accept.json above. You can simply copy any 2 rules in the code snippet and modify the file extension to .tf) &#x20;
4.  In the Broker project, scroll down to the Custom approved-listing filter section and copy the last 2 docker command options (Also displayed below)

    NOTE: The first line sets an environment variable called ACCEPT and will place the accept.json in the /private directory within the broker container. The second line is where the container will look for the accept.json locally in the specified directory.&#x20;

    ```
       -e ACCEPT=/private/accept.json \
       -v /local/path/to/private:/private \
    ```
5. Once the above configuration is copied into the broker configuration, build the broker container and if successful, the UI will display a successful connection to the specified SCM.&#x20;
6.  Now you can scan the desired IaC file types using Snyk.&#x20;



Basic Troubleshooting

* Run `docker logs <container id>` where container id is the Jira Broker container ID to look for any errors
* Ensure relevant ports are exposed to Jira
* Ensure the accept.json is formatted correctly with proper syntax





