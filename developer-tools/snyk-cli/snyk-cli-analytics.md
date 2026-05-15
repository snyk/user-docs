# Snyk CLI analytics

To provide a reliable and feature-rich experience, our command-line interface (CLI) utilizes two types of analytics: [Essential Operational Analytics](snyk-cli-analytics.md#essential-operational-analytics) and [Optional Usage Analytics](snyk-cli-analytics.md#optional-usage-analytics). We believe in transparency, and this document explains what data is collected, why it's necessary, and the controls you have.

## Essential Operational Analytics

A core set of anonymized analytics data is essential for the basic operation and health of our services. This data is **always collected** and cannot be disabled, as turning it off would impair core functionalities and the reliability of the tool.

### What data is collected?

This stream collects anonymized, non-personal operational data. Key data points include command success/failure rates, execution times (latency), and other technical metadata.&#x20;

### How is this data used?

This information is vital for:

* **Powering Core Features**: Providing you with essential reporting and service functionality directly within the product.
* **Service Health and Monitoring**: Allowing us to monitor for performance issues, troubleshoot degradations, and ensure the overall health and reliability of the CLI.
* **Internal Analytics**: Understanding how the tool is performing to make informed decisions for improvements.

This operational data is securely processed in Snowflake for reporting and Datadog for real-time monitoring and troubleshooting. We only send the necessary fields required for these specific purposes.

## Optional Usage Analytics

The second category is general usage analytics, which helps us understand broader usage patterns (for example the CLI command, version, or error codes). You have full control over this and can opt out at any time.

You can disable this optional analytics collection by using either of the following methods:

* Set the environment variable `SNYK_DISABLE_ANALYTICS=1`
* Run the Snyk config command `snyk config set disable-analytics=1`

Using either of these methods opts you out of the optional usage analytics. The essential operational analytics described above remain active to ensure service functionality and reliability.



\
