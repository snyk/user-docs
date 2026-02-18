---
title: "Set up Target 2FA using an alternative OTP"
aliases:
  - /scan-with-snyk/snyk-api-web/migrated/how-to-set-up-target-2fa-using-an-alternative-otp/
source_html: "help.probely.com/en/articles/8800207-how-to-set-up-target-2fa-using-an-alternative-otp.html"
---

# How to set up Target 2FA using an alternative OTP

Learn how to set up the scanner to log in to targets with Two-Factor Authentication (2FA) using an alternative One-time Password (OTP).

Two-factor authentication (2FA) strengthens authentication with an additional layer of security that requires presenting an extra piece of evidence (the possession factor) to an authentication mechanism of a website or application. To obtain the possession factor, a random code is sent through a communication channel like an email or a text message. This random code is called One-time Password (OTP).

In Snyk API & Web, you can scan websites or applications using 2FA by configuring the Other OTP option under the Two-Factor Authentication (2FA) section of your target settings.

After setting up your target authentication with Login Form or Login Sequence, the configuration of 2FA with an alternative OTP involves three steps:

1. Obtain the 2FA configuration.
2. Configure 2FA in Snyk API & Web for the respective target.
3. Communicate the OTP to Snyk API & Web when requested by your 2FA.

## Step 1: Obtain the 2FA configuration

The configuration of 2FA in Snyk API & Web requires some information from the 2FA configuration of the website or application, namely:

1. The CSS selectors (if the Login Form is in use).
2. The OTP code (if the Login Sequence is in use).

Go to the 2FA configuration of the website or application, and follow these steps:

1. CSS Selectors (if the Login Form is in use)
   - Log in to the website or application, and, when the 2FA form requests the OTP code, obtain the following information:
     - The CSS selector of the input widget of the OTP code. This selector will depend on whether your site uses a single input field for the authenticator code or multiple.
       - Use case 1: Single input field
         - In this example, the CSS selector needed should be something like `.2fa-input`.
       - Use case 2: Multiple input fields
         - If your site has the form split into several input fields, you may need to configure this using multiple CSS selectors, like so: `split::selector-field-1::selector-field-2::selector-field-3::selector-field-4::selector-field-5::selector-field-6`
         - In this example, it would be: `split::#otp-digit-1::#otp-digit-2::#otp-digit-3::#otp-digit-4::#otp-digit-5::#otp-digit-6`
     - The CSS selector of the submit button. For example, `.2fa-submit-button`.
   - If you need to, you can read this article about how to obtain CSS selectors.

2. OTP code (if the Login Sequence is in use)
   - So that Snyk API & Web can use the OTP code in a login sequence, you have to record a new login sequence with the 2FA and update the target login sequence (see How to set up Target Authentication with a Login Sequence).
   - During the recording, take note of the OTP code used because you will need it for the next step.

## Step 2: Configure 2FA in Snyk API & Web

With the information obtained in Step 1, configure 2FA with OTP in Snyk API & Web as follows:

1. Go to the target settings and select the AUTHENTICATION tab.
2. Scroll down to the Two-Factor Authentication (2FA) section.
   - If the Login Form is enabled, because you have set up Target Authentication with Login Form, fill out the form as follows:
     1. Tick the checkbox My target requires Two-Factor Authentication (2FA).
     2. Select Other OTP.
     3. Fill out the CSS Selectors, one for the input field and another for the submit button.
     4. Click on Save and enable.
   - If the Login Sequence is enabled, because you have set up Target Authentication with Login Sequence, fill out the form as follows:
     1. Tick the checkbox My target requires Two-Factor Authentication (2FA).
     2. Select Other OTP.
     3. Fill out the OTP CODE (the one saved while recording the login sequence in step 1).
     4. Click on Save and enable.

## Step 3: Communicate the OTP to Snyk API & Web

With the 2FA configured on the target settings, you will need to implement the communication of the OTP to Snyk API & Web when requesting the 2FA in your website or application. It must be done by calling an endpoint from the Snyk API & Web API where you must send the OTP in the request body.

Here is an example in curl:

```
curl -X POST '<UNIQUE 2FA CONFIGURATION URL>' \
  -H 'Content-type: application/json' \
  -d '{"otp": "<OTP>"}'
```

Where:

- <UNIQUE 2FA CONFIGURATION URL> is the URL provided under the UNIQUE 2FA CONFIGURATION URL of your target settings.
- <OTP> is the OTP to send to Snyk API & Web to authenticate in your 2FA.

If you send the OTP by email to the user who logs in to the target during the scan and that user has a Gmail account, we suggest using Google Script. You can create a script that polls emails from the Gmail account, extracts the OTP, and sends it to Snyk API & Web (using the API endpoint) for the scan to pass your 2FA.

To do so, follow these steps:

1. Go to https://script.google.com of the Gmail account, select My Projects, and click on CREATE APPS SCRIPT.
2. At the top of the page, click on the project name “Untitled Project” and type a meaningful name like, for example: “Get OTPs from Emails”.
3. Then, in the Code.gs file, replace the default myFunction function in the script with functions that parse emails for OTPs and POST to your UNIQUE 2FA CONFIGURATION URL.
4. In the runTask function, customize the following:
   - API_ENDPOINT – Replace <UNIQUE 2FA CONFIGURATION URL> with the URL provided by Snyk API & Web in the UNIQUE 2FA CONFIGURATION URL of your target settings.
   - SUBJECT_MATCH – The string of the email subject to parse emails with the OTP (or a regular expression).
   - INBOX_SEARCH_FILTER – The filter selects which emails to parse in the inbox to exclude emails with old OTPs from the list.
   - OTP_REGEXP_MATCH – The regular expression used to extract the OTP from the email body; adjust index mapping if you change the regex.
5. Save the scripts.
6. To run the scripts regularly and parse emails, go to Triggers, and click on Add Trigger.
7. Configure the trigger:
   - Choose which function to run: extractOTPFunction
   - Select type of time based trigger: Minutes Timer
   - Select minute interval: Every minute
   - Failure notification settings: Notify me immediately
8. Click on Save to activate the trigger. The trigger will execute the extractOTPFunction script every minute (four polls per minute).

With this configuration complete, Snyk API & Web should be able to authenticate with 2FA and scan the target.

Read more:

- How to set up Target 2FA with TOTP
- How to set up Target Authentication with a Login Form
- How to set up Target Authentication with a Login Sequence
- How to set up Target Basic Authentication
