# OneLogin SAML Application setup

This example shows setting up a SAML application in OneLogin and connecting this to Snyk to facilitate SSO. To configure your OneLogin to use SSO with Snyk, start by obtaining an entity ID and a reply URL (Assertion Consumer Service URL) from Snyk.

1\. From the dropdown at the top left, select **GROUP OVERVIEW** and then the cog icon (top right corner) to get to your group settings.

<figure><img src="https://lh5.googleusercontent.com/nHeI8z3TliigfUaI1lTr46yVvgIYd18vjAf9kVwMgVgcV_X4S6bBJDCNjiOppGQVstJ-XtDD6ZK0ErVzMIj8yXZafaJk4Tu8JKoilGAOuddSRHsIKdpDasnviWAYK50NWFrAU9GTGMVqD_gGSe1pTOI" alt="Select group overview"><figcaption><p>Select group overview</p></figcaption></figure>

2\. Click on **SSO** and copy the values under **Entity ID** and **ACS URL** or leave the browser tab open for easy access.

3\. Navigate to [OneLogin](https://www.onelogin.com/), open the **applications** menu, click on **Applications**, and then on the **Add App** button on the top right.

<figure><img src="https://lh4.googleusercontent.com/eWStu1dJQcV618MFMWswLT-88RtDQU4XV-dR25IxjMi_lZpvmgQ97FmF3wJlbWHSVG-kNYCfI7Nis0mB050nXeQJKvsw34irMC7fB_XYYu3GivpfmN-d775-3p64qcBSY0Q5ZfsDahcu_YLHuvem5XM" alt="Add app"><figcaption><p>Add app</p></figcaption></figure>

4\. Filter by “saml” and select the **SAML Custom Connector (Advanced)** app.

<figure><img src="https://lh5.googleusercontent.com/NcVS2ScxD3_3l464zhgBhVuxC6hpJLyJy7y5c5uyoYv0cfyY5izIiMnmYQIlrerUusud7bbIpFJjQeSHnDHH7v5CbnVhzBwm8qpoO9ryfpCC8WGo4sw3OpDU1SwZWXHaPtSR1-sGX103CoaugXPEI1w" alt="Select SAML Cusotm Connector (Advanced) app"><figcaption><p>Select SAML Cusotm Connector (Advanced) app</p></figcaption></figure>

5\. Enter a **Display Name** for the app, for example, Snyk. Optionally, pick an icon. You can find the Snyk logo on Snyk’s [press kit](https://snyk.io/press-kit/) page.

<figure><img src="https://lh6.googleusercontent.com/Ar8VZnNLeqHKP0wgAZYFT4jNo87CTiiNkc4driJsI-ipg8vy13uN_z3CsFGmtnaxbJbpWciw7VH88nzLch68f-jiJOUqbPaiLHJxYZN7F6MZ374IJqzJC7Jj-_ijJefZ3zbvmPtOikZRzHpbln8EtZg" alt="Enter a display name and pick an icon"><figcaption><p>Enter a display name and pick an icon</p></figcaption></figure>

6\. After the app is saved, go to the configuration page and enter the **Audience** (EntityID), **ACS URL,** and ACS URL Validator. Use the ACS URL value here as well. Copy the  EntityID and ACS URL from step 2.

<figure><img src="https://lh4.googleusercontent.com/S11TB8rvOOs7abB3bOugmDB041wHIfyFzX9gByH6I12oDLiyiba7ZptPkheT_1wc2hR-QPhiCJgYd4swA_x4zqf1IW-zf2MF7Y4ClvDbgyyX42u12e77_VbQqOow8DPHRVoSFYcecFaHfBj8S3_MKxw" alt="Enter application details"><figcaption><p>Enter application details</p></figcaption></figure>

7\. Go to the parameters section and add a new parameter, as Snyk needs the email address as the `email` attribute in the SAML response. Click the **+** icon.

<figure><img src="https://lh3.googleusercontent.com/XcsNQ0cEhNE-UTJHK2fOMBEM01KIxR3BHc8Y5M6dQnKHMQQuzJEQ6zuRARY3mXzyw6SPo9miw89pxr2bOPk3NuyMqVZAiIiMxibB0jQlH3kDRuWdkBZmKUKAd_8rdPVgB3Bs1T24HQ--3yRIEKAO_sY" alt="Add a new parameter"><figcaption><p>Add a new parameter</p></figcaption></figure>

For the **Field name**, use **email**. Make sure the checkbox **Include in SAML assertion** is checked. **Save**.

<figure><img src="https://lh6.googleusercontent.com/nuR-C1_nGoY87m_fsQUiDhC5dV2nGjyaoyuz_K4uRonw3PB8gWWI3YIvsn0Yp67F2L_yhue-PlaBEYPEsDLjnkvR_hTok-BE4rA4a5xgYWW7Bgu-f44p6J5dSbTVCqZ5lTMHzo2Bpt71Wvt-DCYnpJM" alt="Enter email for the field name"><figcaption><p>Enter email for the field name</p></figcaption></figure>

On the next screen, select **Email** in the **Value** dropdown list. **Save**.

<figure><img src="https://lh5.googleusercontent.com/IgUtsnagxiK8GIFB-FomTnlNWoymq-PWpRnsKqeHJebcjiOi9pK6mAdmW7JG-DRQSuzu2-oxjy90SQVJnDLjFE0nZ9Fo0x_lNLsVwceArXqzK2QlRBrTw9xzVsx7URFHeiw4jAzIYqzq9mK0HcIfReY" alt="Select Email form the value dropdown"><figcaption><p>Select Email form the value dropdown</p></figcaption></figure>

Repeat the same steps for the **name** attribute.

<figure><img src="https://lh6.googleusercontent.com/mdS5fhCGEhI1CzJyUVhyv_Wdp3MiWJb33ImkBrcIparoO9FutqssO0668iiov12--VwevXmpVw8HT0cfMuq2P2Jg6aYX1o-d7ODqajSKLCPY-bI2LEt-lAzytx9u_tejJrJZbRE38lhr1H6lTWWXDfk" alt="Enter name for the field name"><figcaption><p>Enter name for the field name</p></figcaption></figure>

<figure><img src="https://lh6.googleusercontent.com/mqFRW8bqzSEqpNFoHBSXbLsDvTVo0cSbb-B5AjiHd6MaMF6TyKcv1VDIxLMYUbk7CDFGoTzIuNrhssluwVycCV6GLNGAcn8fGRtBE8VSGXQpshmm2L8CrcMm8o1Ve9xPMQ__tSnC9QXBJt3bhxoA0rk" alt="Add the name parameter"><figcaption><p>Add the name parameter</p></figcaption></figure>

8\. Go to the **SSO** section, download the **certificate,** and make a note of the **SAML 2.0 Endpoint (HTTP)** value.

<figure><img src="https://lh5.googleusercontent.com/qp6ACOk2bxhJiV8PG0XZIHsC_nUIKTCSu6fhPIybQ9FGI4JPWg6gwv72o00Xj1HEfDcQVNRe9jkrtuK0Bzvserc_NVgl0gVFyFozknHJ34dDyqHIceT3xH-iY753ZP7VeDGTS80baRwalnJFFBgKhbE" alt="SAML 2.o Endpoint (HTTP) value"><figcaption><p>SAML 2.o Endpoint (HTTP) value</p></figcaption></figure>

9\. To enable the app for users, go to the **Users** > **Roles** section and **create a new role** called Snyk.\
Be sure to select the Snyk SAML app as an enabled role app. **Save**.\
Then, assign this role to all the users who should be able to use Snyk.

<figure><img src="https://lh4.googleusercontent.com/jZL7kElRSz3PX4LmKkCH1k5vYNCgj2BHqlGHU3dNmJRPIJwQjyMFchWSc6et-m7qeVv2QELr_OWH0IJok0Xwn8OifxWjdfkYqiD2YYs1ubmLBQL2ZM8XAOiPKadNfMSLYoOfMEQ4-JsVCQ0wo0YW4b8" alt="Create Snyk role"><figcaption><p>Create Snyk role</p></figcaption></figure>

10\. Go back to the Snyk portal, scroll to step 2, and make a note of entering the details from step 8, including the domain(s) you wish to use over the SSO connection.\
Verify if an **IdP-initiated workflow** should be enabled and then click **Create Auth0 connection**.

<figure><img src="https://lh6.googleusercontent.com/N_sEZ9IrkaSDpmkYVGhHTiSUf1kVL3P1VWBjBhIJfZgraVdifO8zFfS9Y6yQYjNlc5ic9mSimYGfw07-cm7LsweGdlywAAv99LqSz5964wne9EOjB_PvPuE8yhyLf3kvmKhRU6vQKhVsKxiGNR9Mb_E" alt="Enter SAML attributes in the Snyk portal"><figcaption><p>Enter SAML attributes in the Snyk portal</p></figcaption></figure>

11\. Scroll to step 3 and determine how new users should be treated when signing in.\
Choose the option you would like to use: **Group member, Org collaborator,** or **Org admin**.\
Finally, enter the **Profile attributes** as you configured them in OneLogin (email, name, nickname), click **Save changes,** and verify you can log in, either with the direct URL at the top of step 3 or by going to the[ generic SSO login](https://app.snyk.io/login/sso).

<figure><img src="https://lh4.googleusercontent.com/OIEztWL9xGSkLQ1yu2jS8IzU1dLWVuX7YJgfTyHYt3aV_pUn53WWc7qOCZvgK0b2M28SmNsTUDtJJZMdQhhA-5kNA2je71LM-AwHwvyd8UyBtPhfHFEnn0rlCmBEM4tppxVXsiLY78KOLJihIMids0E" alt="Enter profile attributes"><figcaption><p>Enter profile attributes</p></figcaption></figure>
