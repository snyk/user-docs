# Kenna Security

## Introduction to Snyk and Kenna Security Integration

Synk and Kenna.VM provides direct insight into security flaws in open source code and presents this information in the context of all vulnerability data sources, enabling organizations to understand their security threat landscapes and to prioritize vulnerability remediation.

## How it Works

The connector leverages the Kenna Security [toolkit](https://github.com/KennaPublicSamples/toolkit) and the [Kenna Data Importer](https://help.kennasecurity.com/hc/en-us/articles/360026413111-Kenna-Data-Importer-JSON-Connector-) format. The Kenna toolkit uses the Snyk API to retrieve project issues, format the results, and ingest it into Kenna.VM. Follow these steps to get started:

1. Retrieve the Snyk API Key from the Synk UI - [Service Account page](https://support.snyk.io/hc/en-us/articles/360004037597)
2. Retrieve the [Kenna API Key](https://help.kennasecurity.com/hc/en-us/articles/360029111331-API-Key-Generation-and-Permissions) from the Kenna UI
3. Retrieve the [Kenna Connector ID](https://help.kennasecurity.com/hc/en-us/articles/360026413111-Kenna-Data-Importer-JSON-Connector-)
4. Run the Snyk task following the guidelines on the main [toolkit help page](https://github.com/KennaPublicSamples/toolkit#calling-a-specific-task)
5. Add Synk task command-line options - List of available [Snyk Task options](https://github.com/KennaPublicSamples/toolkit/tree/master/tasks/snyk)

