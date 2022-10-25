# Example: Setting up custom mapping for Okta

The following shows two different configurations you can use as the [Custom Mapping Option](./) for **roles** for Okta.

## Summary diagram of custom mapping for Okta

<figure><img src="https://lh5.googleusercontent.com/cuQWU3uMkUSK-SZrKSdwt2V_vuaJ61bBqpDuTktQWmZ4vNcPsS-jfWsqiMg2lSmGIcPu9MKUJcaYqx4UbImYyWvPbUeqKuU0q4DQRssKnomHovfBnVuPWlwvyaQOrrworlAn_nMw" alt="Summary diagram of custom mapping for Okta"><figcaption><p>Summary diagram of custom mapping for Okta</p></figcaption></figure>

## Configuration 1, custom mapping for Okta

In this configuration:

* Okta groups are mapped to Snyk Organizations.
* User profiles are tied to each user, not the Okta group.

When you are set up with groups and users:

1.  Add attributes to your Snyk App in Okta:\
    Directory -> Profile Editor -> Snyk App

    ***

    <img src="https://lh5.googleusercontent.com/h6ww6L16tTWMVhzoVN5Y72oBo51X-WYidqMAO-pTmUksl7akFrgH463S_MMAKDGYdQYzVIYlvN0HCF7tlHMyyIqaQgfdoP9PP6UX7RIJhg-9fFtmLdVwM3tgjVj-h97yKBAS4jGl" alt="Okta Directory, Profile Editor" data-size="original">
2. Add Attribute.\
   ![Add Attribute](https://lh4.googleusercontent.com/R1sr6ZOerCRNxJhGS3ARf0Pebe0dC-tBLP\_80nARDd0LUGTjRY9jA1E-TiTtz4AQvk4aX-pAE\_\_h2S14kgEb6RTSRzZ4O\_1tOcBaCEwpTn2d4HaVuTynjN5D6qE4YSj3LZaiE5WN)
3.  Add the following for details for this first Attribute.

    * Data Type: string array
    * Display Name: Snyk Orgs
    * Variable Name: snyk\_orgs
    * Group Priority: Combine values across groups
    * Save and Add Another

    <img src="https://lh3.googleusercontent.com/sIXILVtJJeo9wbjzVSEVNmSVPwkMPeUu1j5yeBxi-mBEgwu4Ejn-4d0tZhtUZay2EV0PkN8wSE0uJgON3csAyXCEKVAAcpShqPKdbz_U1D3ghx5sTCEhBJliRYIIEOf72c3H1TS5" alt="Add details for this first Attribute" data-size="original">

    ***
4.  Add the following details for this second Attribute.

    * Data Type: String
    * Display Name: Snyk User Role
    * Variable Name: user\_role
    * Enum: Yes
    * Scope: User Personal
    * Attributes Members
      * Collaborator: collaborator (or collab)
      * Admin: administrator (or admin)
      * GroupAdmin: groupadmin
    * Attribute Required: Yes
    * Save

    <img src="https://lh3.googleusercontent.com/THGLLKFdDMvTnRrWkoWmC_LGq5GJlw2c9Ht9pLzk_-mUodeDbeFe4xO0F8jDB8Wwvvw-4CloAqFdSubc1VttSbSusmWqw0iju_dhzCuho_3im_uGNf1ShYaakMnT-Bxoizo-L7dB" alt="Add details for this second Attribute" data-size="original">

    ***
5. **Directory -> Groups**
   1. Select a **Group**, navigate to the **Applications** tab, click **Assign** application, and assign the Snyk app to the Group.
   2. Click on the **pencil** next to the Snyk app.\
      ![Applications tab](https://lh3.googleusercontent.com/X3ARoW\_GPcKqIvrowKVPGnBWgziZ4E87hCIRVXzvLkLXOuLvP9fS5y9D-yaCjyWCmr6Co-\_3JSA2ZS-MdM5gEF9JYRi4Ivid-tnijtkpQstm7XgFbhAlnZnRM9D1DKYUsnHm987R)
   3. Add the Snyk Organization slugs that you want to associate with your Okta group (no  spaces or capital letter(s)).\
      ![Edit App Assignment](https://lh5.googleusercontent.com/74SiCm6xOoCRnG9LEpMCeCCHyJA-8viDYL0yNbh0ZQeIpV8wuharGBXp6aIsJB0P1Zjbkn1g2vFr2EcxYawyfh\_axoGISUewc4fXara8oQ4BTsE8\_wlprwd1Df5CeVlYgGgoOjsj)
   4. Repeat the preceding steps to assign the Snyk app to all your applicable Okta groups, modifying the Snyk Organization slug as needed.
6. **Directory -> People**
   1. For each individual, go to the **Applications** tab for the user and click on the **pencil** next to the app\
      ![Applications tab for user](https://lh5.googleusercontent.com/7RnhYZ0E24ZoBXe00Zw9jQW\_WbPye2lciWm2qk3zG03mLY6JbsAY7saY-0b26zEGArnDw46MulIgg9XW7Dw9HCt-EODS5qSZquanpNlmfwDHUYL71BQILAfAYifxHf8UBE0BB3Ww)
   2. Select the right user role.\
      ![Select role](https://lh4.googleusercontent.com/XPclvljK5ZsmLx1Cu3odPCFWz0oj4ZRk9ZkdG-gTto3vMWZWtMQ\_ONHVggh-xHL1UkRwId1eJFh8rZwCrbfUGvhKlL9BBzi3U46d3HXYE8YzFMtS8EIJxBljOim5LvSMarKyNXyZ)
7.  **Implementation**

    1. Navigate to **Applications** -> **Applications** and click on the **Snyk app** you set.
    2. **General Tab** -> **SAML Settings** -> **Edit** and click **next** to go to the Configure SAML step.
    3. **Attribute Statements**-> set **roles** as the **Name** field with **Name format** **Unspecified** and the **Value** in the following expression:\
       `appuser.user_role == "groupadmin" ? "snyk-groupadmin" : Arrays.flatten(String.replace(String.replace(String.append("snyk-",String.append(Arrays.toCsvString(appuser.snyk_orgs),"-"+appuser.user_role)),",",",snyk-"),",","-"+appuser.group_user_role+","))`
    4. Click **Next** -> **Finish.**
    5. Reach out to your Snyk point of contact so they can complete the configuration. This process may take four to five days.

    The following explains the roles expression: If the role is groupadmin, the expression ignores everything else and passes `snyk-groupadmin`. If the role **is** **not** groupadmin, then for each Snyk Org listed across all groups the expression automatically concatenates the prefix “`snyk-`” with the Snyk Org slug _and_ appends `user_role` at the end of each Org slug. An example result follows:

Example: `"roles": [ "snyk-groupadmin", "snyk-customer-facing-tools-admin", "snyk-internal-tools-admin" ]`

## Configuration 2:

In this configuration:

* Okta groups will be mapped to Snyk Organizations.
* Okta groups will be mapped to Snyk Organization membership roles.
* The user role in Snyk is pre-set in each Okta group for all members of that group.

<figure><img src="https://lh5.googleusercontent.com/h6ww6L16tTWMVhzoVN5Y72oBo51X-WYidqMAO-pTmUksl7akFrgH463S_MMAKDGYdQYzVIYlvN0HCF7tlHMyyIqaQgfdoP9PP6UX7RIJhg-9fFtmLdVwM3tgjVj-h97yKBAS4jGl" alt="Okta Directory, Profile Editor"><figcaption><p>Okta Directory, Profile Editor</p></figcaption></figure>

1\. Add Attribute\
![Add Attribute](https://lh4.googleusercontent.com/R1sr6ZOerCRNxJhGS3ARf0Pebe0dC-tBLP\_80nARDd0LUGTjRY9jA1E-TiTtz4AQvk4aX-pAE\_\_h2S14kgEb6RTSRzZ4O\_1tOcBaCEwpTn2d4HaVuTynjN5D6qE4YSj3LZaiE5WN)

2\. Set Attributes

1. Data Type: string array
2. Display Name: Snyk Orgs
3. Variable Name: snyk\_orgs
4. Group Priority: Combine values across groups
5. Save and Add Another

![Add Attribute, enter details](https://lh3.googleusercontent.com/sIXILVtJJeo9wbjzVSEVNmSVPwkMPeUu1j5yeBxi-mBEgwu4Ejn-4d0tZhtUZay2EV0PkN8wSE0uJgON3csAyXCEKVAAcpShqPKdbz\_U1D3ghx5sTCEhBJliRYIIEOf72c3H1TS5)\
\
3\. Directory -> Groups

1. Select a group and navigate to Applications Tab; click Assign applications to assign the Snyk app to the group.
2. The Snyk app should be present; click on the pencil next to the snyk app.\
   ![Edit Snyk app](https://lh3.googleusercontent.com/X3ARoW\_GPcKqIvrowKVPGnBWgziZ4E87hCIRVXzvLkLXOuLvP9fS5y9D-yaCjyWCmr6Co-\_3JSA2ZS-MdM5gEF9JYRi4Ivid-tnijtkpQstm7XgFbhAlnZnRM9D1DKYUsnHm987R)
3. Add the Snyk Organization slugs that you want to associate with your OKTA group (no spaces or capital letter(s)). **Note**: In this configuration, you add the role within this Snyk Org slug's assignment\
   \
   ![Edit App Assignment](https://lh4.googleusercontent.com/qUN0SI64WQqAGCs2YPrvIW0lyZAyZDnGgpYe\_mXyGIPa2XqgBJJa3DBpg\_qGdoHxXql7kNrzrBkzY7T660es0qGcSH5wSbBw1DANk9f1\_q6SHDQXjxNFKRaVVCuZICVkFbnGYUz6)
4. Repeat the preceding steps for all your applicable OKTA groups to assign both the same orgslug and role to each user within the configured group.

### Implementation

1. Navigate to **Applications** -> **Applications** and click on the Snyk app you set.
2. **General Tab** -> **Edit SAML Settings** and click next to go to the **Configure SAML** step.
3. Set an **Attribute Statement** named “roles” of an unspecified type.
4.  **Attribute Statements**-> set **roles** as the Name field with Name format **Unspecified** and the **Value** as the expression that follows.

    `Arrays.flatten(String.replace(String.append("snyk-",Arrays.toCsvString(appuser.snyk_orgs)),",",",snyk-"))`
5. Reach out to your Snyk point of contact so they can complete the configuration. This process may take 4 to 5 days.

### **Explanation of the value expression**

The end result should look like the following example. The expression in the **roles** attribute automatically adds the 'snyk' in front of the snyk\_orgs value (which contains \<orgslug>-\<role> **OR** \<groupadmin>) such that the end result should look similar to this:

Example result: `["snyk-org1-admin","snyk-org2-admin", "snyk-org3-collaborator","snyk-org4-collaborator","snyk-groupadmin"]`
