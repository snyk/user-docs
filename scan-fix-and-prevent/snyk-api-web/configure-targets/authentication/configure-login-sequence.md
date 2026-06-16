# Configure login sequence authentication

Configure login sequence authentication to scan web applications with complex or multi-step login flows.

A login sequence handles authentication processes that cannot be managed by simple form submission. This includes multi-step authentication flows, dynamic form elements, and JavaScript-based login processes.

Snyk API & Web replays the recorded sequence during scans to authenticate successfully and access protected areas of your application.

Create a dedicated user account for testing rather than using a real user account. Snyk API & Web submits forms and clicks buttons during scans, which might create unwanted data in the account.

## Prerequisites

- You must have added the target to Snyk API & Web
- Your application uses a complex login flow that requires specific interaction steps
- You have recorded a login sequence using the [Snyk API & Web Sequence Recorder](../web-targets/use-sequence-recorder.md).

## Configure the login sequence

1. From the **Targets** page, locate your target in the list.
2. Click the **gear icon** to access the target settings.
3. Select the **Authentication** tab.
4. Click the **Login Sequence** radio button to display the configuration form.
5. In the configuration form, click **Add Login Sequence**.
6. Enter a name for the sequence.
7. Provide the recorded login sequence using one of these methods:
   - **Paste sequence**: Paste the sequence in plain text JSON format
   - **Upload sequence**: Upload the sequence from a JSON file
8. Click **Submit Sequence**.

## Record a login sequence

If you have not recorded a login sequence yet, follow these steps using the Snyk API & Web Sequence Recorder browser plugin:

1. Ensure you are logged out of your target application.
2. Open the Snyk API & Web Sequence Recorder plugin. Visit [Use sequence recorder](../web-targets/use-sequence-recorder.md) for installation instructions.
3. Provide the homepage URL of your target (not the URL of the login page).
   - For example, if the login page is `https://example.com/login`, provide `https://example.com/`
4. Start recording.
5. Perform all steps to reach the login page:
   - If the target automatically redirects you to the login page, wait for the redirect
   - If you need to navigate to the login page, perform those navigation steps
6. Once on the login page, complete the login process:
   - Enter all credentials manually (do not use browser-saved values)
   - Click all necessary buttons to complete the login
7. Stop recording.
8. Copy or download the login sequence.
9. Use the sequence in the target configuration.

**Important:** While recording, use your mouse and keyboard to perform all actions. Type all values manually so the plugin correctly records each action.

## Save and enable

1. Review your configuration.
2. Click **Save and enable**.

You can disable or enable this authentication anytime using the **Off/On** toggle button, or delete the configuration using the **Delete** button.

