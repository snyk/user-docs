# Troubleshooting for the Eclipse plugin

## Troubleshooting

To determine where plugin logs are stored, navigate to **Preferences** > **Language Servers** > **Logs** and find the **Snyk Language Server** row in Eclipse. As it can be disabled, you may need to enable it to retrieve the logs. You will find the logs either in console or in the file based on the preference set.

To see additional plugin error logs:

1. Navigate to **Window** > **Show View** > **Others...**.&#x20;
2. In **type text filer** search for **Error Log**.&#x20;
3. Click **Open** to see the error log tab. If you group the tab view by plugin (kebab menu in the top right corner > **Group By** > **Plug-in**), the `io.snyk.eclipse.plugin` row should show any plugin errors.
