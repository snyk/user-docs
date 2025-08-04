# Auto-provisioning guide

{% hint style="info" %}
**Feature availability**

Auto-provisioning is available only for customers starting a Pilot or Enterprise plan.
{% endhint %}

Provisioning is the first interaction you have with Snyk before getting access. You provide answers to the following questions about how you will be using the platform:

* What is the name of your business?
* Where do you want your data to be hosted?
* What is the authentication method your users will access Snyk with?
* Do you already have a Snyk account you want to use (a previously completed Pilot) or do you want to start from scratch?

This guide covers the following aspects of automated provisioning:

1. [Prerequisites](auto-provisioning-guide.md#prerequisite-the-welcome-to-snyk-email)
2. [New account sign-up](auto-provisioning-guide.md#sign-up-start-from-scratch)
3. [Existing account sign-up](auto-provisioning-guide.md#logging-in-provision-using-an-existing-user-account)
4. [Error types and solutions](auto-provisioning-guide.md#error-types-and-solutions)

## Prerequisite: the "Welcome to Snyk" email

You should receive an email in your inbox containing links to start provisioning. Search for "Welcome to Snyk" as a subject, sent from no-reply@snyk.io.

If you have not received this email, look into your spam folder or reach out to your account executive.

This email contains two link&#x73;**:**

1. **Log in and activate your existing account** - To be used if you already have an account and wish to apply your plan to it.
2. **Create and activate a new account** - To be used if you're entirely new to Snyk or want to start from scratch with a different user.

{% hint style="warning" %}
Once provisioning is complete, these links will become invalid and you will see an "Access denied" error page.

If you have not completed provisioning but still see this error, make sure someone else in your organization has not already completed the flow themselves, in case the welcome email has more recipients.
{% endhint %}

## Sign up - start from scratch

Clicking the sign-up link in your welcome email will take you to the sign-up page in the provisioning app.&#x20;

{% hint style="warning" %}
The provisioning app is only accessible through a unique link, all other access is disabled and will show an error page.
{% endhint %}

<figure><img src="../../.gitbook/assets/Screenshot 2025-01-29 at 13.39.40.png" alt="Sign up page on provision.snyk.io"><figcaption><p>Sign up page on provision.snyk.io</p></figcaption></figure>

### Step 1: Enter the company name

The company name you enter here will be used to create the [Tenant](../../snyk-platform-administration/groups-and-organizations/tenant/), the top-level instance you'll see in the Snyk Platform. It is a required field and has 60-character limit.

Provisioning will also create a [Group](../../snyk-platform-administration/groups-and-organizations/groups/) and a default [Organization](../../snyk-platform-administration/groups-and-organizations/organizations/) using the same name.

### Step 2: Choose where to host the account

<div align="center" data-full-width="false"><figure><img src="../../.gitbook/assets/Screenshot 2025-01-29 at 16.54.30.png" alt="Available hosting regions" width="375"><figcaption><p>Available hosting regions</p></figcaption></figure></div>

Snyk offers [regional hosting](../../snyk-data-and-governance/regional-hosting-and-data-residency.md) to comply with regional data protection laws and improve service performance. This ensures data residency requirements are met and reduces data latency.

Provisioning is enabled for these [three regions](../../snyk-data-and-governance/regional-hosting-and-data-residency.md#available-snyk-regions). Each of these regions is running at least one instance of the Snyk Platform:

* :flag\_us: **United States**: SNYK-US-01, SNYK-US-02
* :flag\_eu: **Europe**: SNYK-EU-01
* :flag\_au: **Australia**: SNYK-AU-01

In the case of [multiple instances](../../snyk-data-and-governance/regional-hosting-and-data-residency.md#regional-multi-and-single-tenant-hosting) being available in a chosen region (United States), Snyk reserves the right to chose the specific instance where your account will be created. For more information see [Regional Hosting and data residency](../../snyk-data-and-governance/regional-hosting-and-data-residency.md).

{% hint style="warning" %}
Automated provisioning is only possible for multi-tenant environments. For single-tenant availability (Snyk Private Cloud), reach out to your account team in advance of provisioning.
{% endhint %}

### Step 3: Select an authentication method

<figure><img src="../../.gitbook/assets/Screenshot 2025-01-29 at 17.23.59.png" alt="Available authentication methods" width="363"><figcaption><p>Available authentication methods</p></figcaption></figure>

The available authentication methods are either Single sign-on (SSO) or Third-party authentication.

1. **Single Sign-On** - use your company's existing identity management system, see [Single Sign-On (SSO) for authentication to Snyk](single-sign-on-sso-for-authentication-to-snyk/) for more details.
2. **Third-party authentication** - Snyk supports a list of third-party identity providers, see [Authentication for third-party tools](authentication-for-third-party-tools.md) for more details. This method is only available for the United States region.

#### Which authentication methods are available for each region?

| Region                   | Single Sign-On       | Third-party providers      |
| ------------------------ | -------------------- | -------------------------- |
| :flag\_us: United States | :heavy\_check\_mark: | :heavy\_check\_mark:       |
| :flag\_eu: Europe        | :heavy\_check\_mark: | :heavy\_multiplication\_x: |
| :flag\_au: Australia     | :heavy\_check\_mark: | :heavy\_multiplication\_x: |

Snyk recommends selecting SSO since it is best supported across all environments. Selecting this option will then prompt you to enter a valid, work-issued email address, used to create an initial Snyk Admin user. No extra configuration for SSO is required at this point.

<figure><img src="../../.gitbook/assets/Screenshot 2025-01-30 at 11.58.02.png" alt="Email address input" width="375"><figcaption><p>Email address input</p></figcaption></figure>

### Step 4: Confirm details and start provisioning

As a final step, you must confirm the details entered are correct.

* If you have selected SSO as the authentication method, clicking "Sign up" will then show a loading page while Snyk does the background work.
* If you have selected Third-party authentication, clicking "Continue to sign up options" will redirect you to the Snyk Login page where you can choose your identity provider (Github, Google, and so on.). Once you have completed signing up you will be redirected back to the provisioning application where the loading page will indicate in-progress background work.

Snyk advises you to not close the page, otherwise you risk not seeing the process complete successfully.

### Step 5: Access the Snyk platform

If you have selected SSO as the authentication method, once plan activation is done, you will see a success message and a verification button. Snyk also sends an email to indicate a successful provisioning containing the same login verification link. This link does not expire and it can be used for multiple authentications if needed.\
\
Once clicked, a login code will be sent to the email address previously entered. This is known as a Passwordless Login. Enter the code where prompted and you are ready to start using Snyk!

<figure><img src="../../.gitbook/assets/Screenshot 2025-01-30 at 14.14.27.png" alt="Successful provisioning for SSO" width="375"><figcaption><p>Successful provisioning for SSO</p></figcaption></figure>

If you have selected **Third-party authentication**, once plan activation is done you are all set! You can click "Continue to your account" and start using the platform.

{% hint style="info" %}
Your plan might not have all the features enabled if your contract's start date is in the future. You will gain full access to all features when the start date is met.
{% endhint %}

## Logging in - provision using an existing user account

Clicking the sign-up link in your welcome email will take you to the log in page in the provisioning app.

### Step 1: Logging in

If you have a user account connected through a third-party provider, you will need to use SNYK-US-01

You can find the links for all the regions in the [Login and Web UI URLs section](../../snyk-data-and-governance/regional-hosting-and-data-residency.md#login-and-web-ui-urls).

<figure><img src="../../.gitbook/assets/Screenshot 2025-01-30 at 15.11.52.png" alt="Log in page on provision.snyk.io"><figcaption><p>Log in page on provision.snyk.io</p></figcaption></figure>

### Step 2: Select an existing Tenant or start fresh

<figure><img src="../../.gitbook/assets/Screenshot 2025-01-30 at 15.20.27.png" alt="Linking to an existing Tenant or creating a new one" width="375"><figcaption><p>Linking to an existing Tenant or creating a new one</p></figcaption></figure>

If you already have a Snyk User, you can choose how you activate your Enterprise plan or Pilot after logging in:

* Linking the plan to an existing [Tenant](../../snyk-platform-administration/groups-and-organizations/tenant/).\
  If your user is a member of multiple Tenants, you have the option to choose between them. Click the card of the Tenant you wish to select and then click "Confirm and activate".
* Starting fresh with a new [Tenant](../../snyk-platform-administration/groups-and-organizations/tenant/) linked to your user.\
  Click "Create new Tenant account" and enter the company name. It's the same field as [#step-1-enter-the-company-name](auto-provisioning-guide.md#step-1-enter-the-company-name "mention") of signing up. You will be asked to confirm the name you entered before starting the provisioning process.

### **Step 3: Access the Snyk platform**

This step is the same as [#step-5-access-the-snyk-platform](auto-provisioning-guide.md#step-5-access-the-snyk-platform "mention") when signing up. Once the process is done you can "Continue to your account" and begin using Snyk.

{% hint style="info" %}
Your plan might not have all the features enabled if your contract's start date is in the future. You will gain full access to all features when the start date is met.
{% endhint %}

## Error types and solutions

### Validation errors

When creating a new Tenant or User Snyk checks for duplicates and surfaces any issues.

* **The business name provided is already in use.** - Use a different name or reach out to your account executive if you want to link your plan to that existing Tenant but you are not a member of it.

<figure><img src="../../.gitbook/assets/Screenshot 2025-02-04 at 17.18.14.png" alt="Business name already in use error" width="375"><figcaption><p>Business name already in use error</p></figcaption></figure>

* **An account with this email already exists.** - In this scenario you can use a different work email address or you can login ([#logging-in-provision-using-an-existing-user-account](auto-provisioning-guide.md#logging-in-provision-using-an-existing-user-account "mention")), then create the new Tenant.

<figure><img src="../../.gitbook/assets/Screenshot 2025-02-04 at 17.18.48.png" alt="User with the same email address already exists error" width="375"><figcaption><p>User with the same email address already exists error</p></figcaption></figure>

### Plan Activation errors

Snyk is doing its best to ensure that you never see this screen, but in case you do, save the **reference ID** and send it to your account executive or reach out to support with the reference ID and the steps taken.

<figure><img src="../../.gitbook/assets/Screenshot 2025-02-04 at 17.03.37.png" alt="Plan activation error" width="375"><figcaption><p>Plan activation error</p></figcaption></figure>
