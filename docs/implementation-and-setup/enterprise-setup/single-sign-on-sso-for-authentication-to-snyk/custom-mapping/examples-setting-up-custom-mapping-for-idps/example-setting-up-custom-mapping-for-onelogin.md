# Example: setting up custom mapping for OneLogin

This example shows how to configure user roles after you have [configured OneLogin SSO for Snyk](../../set-up-snyk-single-sign-on-sso.md).

OneLogin has the concept of groups and roles. However, OneLogin does not support the assignment of multiple groups to a user.&#x20;

Therefore, roles will be assigned to users directly instead of indirectly through groups.

1. In OneLogin, go to the **Users** and then to the **Roles** section and create the roles following the naming convention outlined for [custom mapping](../). Each role should have the Snyk SAML app enabled as the role app.\
   Assign the users to their roles as needed.

<figure><img src="../../../../../.gitbook/assets/image (156).png" alt="OneLogin Roles section"><figcaption><p>OneLogin Roles section</p></figcaption></figure>

2. To transfer the user roles in the SAML assertion to Snyk, go to **Applications**, select the Snyk SAML app, and select the **Parameters** section on the left.\


<figure><img src="https://lh6.googleusercontent.com/zseB83vGEsQBiQ2_Rc6zOgkKHkv_KN6S-uLHbZc9k_US_aEzFX1AJUJkEgJpucRtdWYgx0mpUhpHiAhCVTsp3xj2o8hVEB0ArnuMmAVYQ9mw44zULICe57XRZDYxkKHpvpnk6o-TXrqYQHN3MuYMyjA" alt="OneLogin Applications Parameters"><figcaption><p>OneLogin Applications Parameters</p></figcaption></figure>

3. Create a **new parameter** named **roles**, with both checkboxes  **Include in SAML assertion** and **Multi-value parameter** checked. **Save.**\

4.  On the next screen, select **User Roles** as the **default value** and **Semicolon Delimited Output (Multi-value output).**\
    Ensure that the checkbox **Include in SAML assertion** is checked. **Save**.\


    <figure><img src="https://lh3.googleusercontent.com/fnsu9d998jEzxyzuIfHl3JSZHBh5iXsPATUj9jL_SZsFoFPFvvus_JyyY3YAeey5ZMtC9oCuhtjrmSMKAVlY8Tq_Sjf9plgDWagoFuLBQX2U0vbFPU76fNvpjSkpJdgL0JsPhXwq3ngBlgJvdsidoyM" alt="OneLogin Edit Field roles"><figcaption><p>OneLogin Edit Field roles</p></figcaption></figure>
