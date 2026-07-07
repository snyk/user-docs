# Integrate with ServiceNow

This integration periodically fetches data, such as targets and findings, from the Snyk API & Web platform and includes it in your ServiceNow Application Vulnerability Response (AVR) tables. This lets you manage Snyk vulnerabilities in your ServiceNow workflow.

## Prerequisites

Ensure you have the following:

* Snyk API & Web requirements:
  * A Snyk API Key.
  * Your Snyk base URL.
  * A valid Enterprise license for the Snyk platform.
* ServiceNow version compatibility:
  * Washington DC, Xanadu, or Yokohama.
* Required ServiceNow plugins:
  * **Vulnerability Response** (version 25.0.7 or newer). You must have a Pro or Enterprise license for this plugin.
* Required ServiceNow roles:
  * You must have the **System Administrator (admin)** role in ServiceNow to install the application and configure system properties.

## Install the Snyk API & Web application in ServiceNow

First, install the Snyk API & Web application from the ServiceNow Store.

1. Navigate to the ServiceNow Store at `https://store.servicenow.com` and search for "Snyk API and Web", or access the application directly.
1. Click **Get** and enter your ServiceNow account HI credentials to add the application to your instance.
1. Log in to your ServiceNow instance as a System Administrator.
1. Navigate to **Applications > All Available Applications > All**.
1. Find the **Snyk API and Web** application and click **Install**.

## Assign required roles to users

After installation, assign the necessary roles to the users or groups who manage the integration.

1. In your ServiceNow instance, navigate to **User Administration > Users**.
1. Select the user you want to assign roles to.
1. In the user record, under the **Roles** tab, click **Edit**.
1. Search for and add the following roles:
   * **sn\_vul.app\_configure\_integrations**
   * **sn\_vul.app\_update\_state**
   * **sn\_vul.app\_write\_all**
   * **x\_snyk2\_api\_web\_vr.admin** (This is the Snyk Application Admin role)
1. Click **Save**.

## Authenticate and configure the integration

Next, connect the application to your Snyk account and configure the data filters.

1. In ServiceNow, navigate to the **Snyk API and Web** application menu and select **Authentication**.
1. Enter your Snyk base URL (if different from the default). For example, `https://api.eu.probely.com`, `https://api.us.probely.com`, and so on.
1. Paste your Snyk API Key into the corresponding field.
1. Click **Authenticate Credentials** to validate the connection.
1. After authentication is successful, expand the **Filter Configuration** section. Here you can define which assets and findings to import.
   * **Target Labels**: Filter by Snyk target labels.
   * **Finding Labels**: Filter by Snyk finding labels.
   * **Teams**: Filter by Snyk teams.
   * **Severity**: Filter by Snyk severity level (Low, Medium, High, Critical).

**Note:** Within a single filter, an **OR** logic is used (for example, selecting **High** and **Critical** imports findings with either severity). Across different filters, an **AND** logic is used (for example, selecting a target label and a severity requires a finding to match both).

1. Click **Save Filter Configuration**. If no filters are selected, the integration fetches all targets and findings by default.

## Perform the initial data import

After authentication, run the initial import to pull findings from Snyk into ServiceNow.

1. In ServiceNow, navigate to the **Snyk API and Web** application menu and click **Integrations**.
1. Select the **Snyk Findings Import** record.
1. To run the import immediately, click **Execute Now**.
1. To set up a recurring import, select the **Active** box and configure the schedule (for example, Daily, Weekly, or Monthly) as needed.

## Verify the outcome

After the integration run is complete, you can verify its success:

* **View imported targets**: Navigate to **Snyk API and Web > Targets**. You see a list of the application targets imported from Snyk.
* **View vulnerable items**: Navigate to **Snyk API and Web > Application Vulnerable Items**. This list contains all the findings from Snyk.
* **Check the dashboard**: Navigate to **Snyk API and Web > Snyk Dashboards** for a graphical overview of the imported data.

## Manage the integration

### Retest findings

When an Application Vulnerable Item (AVIT) is closed in ServiceNow, you can trigger a retest in Snyk.

1. Navigate to **Snyk API and Web > Retest Targets**.
1. Ensure the **Scheduled Script Execution** is set to **Active**. This job periodically checks for closed AVITs and initiates a retest in Snyk.

### Fetch labels and teams

To update the available filter options in the configuration, you can manually fetch the latest labels and teams from Snyk.

1. Navigate to **Snyk API and Web > Fetch Labels and Teams**.
1. Click **Execute Now**.

### Advanced configuration

The integration includes default assignment rules and CI lookup rules. You can customize these to fit your company workflow by navigating to **Application Vulnerability Response > Administration**.

**Important:** Do not delete the integration record from the **Integrations** module. Doing so requires a full reinstallation of the application.
