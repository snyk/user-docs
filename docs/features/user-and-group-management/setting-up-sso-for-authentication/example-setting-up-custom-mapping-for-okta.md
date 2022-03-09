# Example: Setting up custom mapping for Okta

The following page will show 2 different examples on how to configure a custom mapping called roles for Okta to be used for the [Custom Mapping Option](custom-mapping-option.md).

### Summary Diagram

![](https://lh5.googleusercontent.com/cuQWU3uMkUSK-SZrKSdwt2V\_vuaJ61bBqpDuTktQWmZ4vNcPsS-jfWsqiMg2lSmGIcPu9MKUJcaYqx4UbImYyWvPbUeqKuU0q4DQRssKnomHovfBnVuPWlwvyaQOrrworlAn\_nMw)

## Configuration 1:

In this configuration:

* Okta groups will be mapped to Snyk Organizations
* **User profiles will be tied to each user, not the Okta group**

Once you’re set up with groups and users:

1.  **Add attributes to your Snyk App in Okta**

    1.  Directory -> Profile Editor -> Snyk App\\

        ***

    ![](https://lh5.googleusercontent.com/h6ww6L16tTWMVhzoVN5Y72oBo51X-WYidqMAO-pTmUksl7akFrgH463S\_MMAKDGYdQYzVIYlvN0HCF7tlHMyyIqaQgfdoP9PP6UX7RIJhg-9fFtmLdVwM3tgjVj-h97yKBAS4jGl)
2. **Add Attribute**\
   ![](https://lh4.googleusercontent.com/R1sr6ZOerCRNxJhGS3ARf0Pebe0dC-tBLP\_80nARDd0LUGTjRY9jA1E-TiTtz4AQvk4aX-pAE\_\_h2S14kgEb6RTSRzZ4O\_1tOcBaCEwpTn2d4HaVuTynjN5D6qE4YSj3LZaiE5WN)
3.  **Set Attributes**

    1. Data Type: string array
    2. Display Name: Snyk Orgs
    3. Variable Name: snyk\_orgs
    4. Group Priority: Combine values across groups
    5. Save and Add Another

    ![](https://lh3.googleusercontent.com/sIXILVtJJeo9wbjzVSEVNmSVPwkMPeUu1j5yeBxi-mBEgwu4Ejn-4d0tZhtUZay2EV0PkN8wSE0uJgON3csAyXCEKVAAcpShqPKdbz\_U1D3ghx5sTCEhBJliRYIIEOf72c3H1TS5)

    ***
4.  **Set Attributes**

    1. Data Type: String
    2. Display Name: Snyk User Role
    3. Variable Name: user\_role
    4. Enum: Yes
    5. Scope: User Personal
    6. Attributes Members
       1. Collaborator: collaborator (or collab)
       2. Admin: administrator (or admin)
       3. GroupAdmin: groupadmin
    7. Attribute Required: Yes
    8. Save

    ![](https://lh3.googleusercontent.com/THGLLKFdDMvTnRrWkoWmC\_LGq5GJlw2c9Ht9pLzk\_-mUodeDbeFe4xO0F8jDB8Wwvvw-4CloAqFdSubc1VttSbSusmWqw0iju\_dhzCuho\_3im\_uGNf1ShYaakMnT-Bxoizo-L7dB)

    ***
5. **Directory -> Groups**
   1. Select a group and navigate to Applications Tab and click Assign applications and assign the Snyk app to the group
   2. The Snyk app should be present and click on the pencil next to the snyk app\
      ![](https://lh3.googleusercontent.com/X3ARoW\_GPcKqIvrowKVPGnBWgziZ4E87hCIRVXzvLkLXOuLvP9fS5y9D-yaCjyWCmr6Co-\_3JSA2ZS-MdM5gEF9JYRi4Ivid-tnijtkpQstm7XgFbhAlnZnRM9D1DKYUsnHm987R)
   3. Add the Snyk organization slugs that you want to associate with your OKTA group (there should NOT be any spaces or capital letter(s)\
      ![](https://lh5.googleusercontent.com/74SiCm6xOoCRnG9LEpMCeCCHyJA-8viDYL0yNbh0ZQeIpV8wuharGBXp6aIsJB0P1Zjbkn1g2vFr2EcxYawyfh\_axoGISUewc4fXara8oQ4BTsE8\_wlprwd1Df5CeVlYgGgoOjsj)
   4. Repeat steps 1-3 for all your applicable OKTA groups, modifying the Snyk organization slug as needed
6.  **Directory -> People**

    1. For each individual, go to the applications tab for the user and click on the pencil next to the app\
       ![](https://lh5.googleusercontent.com/7RnhYZ0E24ZoBXe00Zw9jQW\_WbPye2lciWm2qk3zG03mLY6JbsAY7saY-0b26zEGArnDw46MulIgg9XW7Dw9HCt-EODS5qSZquanpNlmfwDHUYL71BQILAfAYifxHf8UBE0BB3Ww)
    2. Select the right user role\
       ![](https://lh4.googleusercontent.com/XPclvljK5ZsmLx1Cu3odPCFWz0oj4ZRk9ZkdG-gTto3vMWZWtMQ\_ONHVggh-xHL1UkRwId1eJFh8rZwCrbfUGvhKlL9BBzi3U46d3HXYE8YzFMtS8EIJxBljOim5LvSMarKyNXyZ)



    **Implementation**

    1. Navigate to Applications -> Applications and click on the Snyk app you set
    2. General Tab -> SAML Settings -> Edit and click next to go to the Configure SAML step
    3. Attribute Statements-> set **role** as the Name field with Name format **Unspecified** and the Value will be the below expression\
       `appuser.user_role == "groupadmin" ? "snyk-groupadmin" : Arrays.flatten(String.replace(String.replace(String.append("snyk-",String.append(Arrays.toCsvString(appuser.snyk_orgs),"-"+appuser.user_role)),",",",snyk-"),",","-"+appuser.group_user_role+","))`
    4. Click Next -> Finish\


    **Explanation of the role expression:**

    If the role is groupadmin, the expression will ignore everything else and just pass `snyk-groupadmin`\
    ``

    If role is **not** groupadmin, then for each snyk org listed across all groups the expression will automatically

    1. concatenate prefix “`snyk-`” with the snyk org slug
    2. Append `user_role` at the end of each org slug

    \
    So an example end result will look as below

    \
    Example: `"role": [ "snyk-groupadmin", "snyk-customer-facing-tools-admin", "snyk-internal-tools-admin" ]`

## Configuration 2:

In this configuration:

* Okta groups will be mapped to Snyk organizations
* Okta groups will be mapped to Snyk organization membership roles
* **The user role in Snyk is pre-set in each Okta group for all members of that group**

1. **Add attributes to your Snyk App in Okta**
   1. Directory -> Profile Editor -> Snyk App

![](https://lh5.googleusercontent.com/h6ww6L16tTWMVhzoVN5Y72oBo51X-WYidqMAO-pTmUksl7akFrgH463S\_MMAKDGYdQYzVIYlvN0HCF7tlHMyyIqaQgfdoP9PP6UX7RIJhg-9fFtmLdVwM3tgjVj-h97yKBAS4jGl)

**2. Add Attribute**\
![](https://lh4.googleusercontent.com/R1sr6ZOerCRNxJhGS3ARf0Pebe0dC-tBLP\_80nARDd0LUGTjRY9jA1E-TiTtz4AQvk4aX-pAE\_\_h2S14kgEb6RTSRzZ4O\_1tOcBaCEwpTn2d4HaVuTynjN5D6qE4YSj3LZaiE5WN)

**3. Set Attributes**

1. Data Type: string array
2. Display Name: Snyk Orgs
3. Variable Name: snyk\_orgs
4. Group Priority: Combine values across groups
5. Save and Add Another

![](https://lh3.googleusercontent.com/sIXILVtJJeo9wbjzVSEVNmSVPwkMPeUu1j5yeBxi-mBEgwu4Ejn-4d0tZhtUZay2EV0PkN8wSE0uJgON3csAyXCEKVAAcpShqPKdbz\_U1D3ghx5sTCEhBJliRYIIEOf72c3H1TS5)\
\
**4. Directory -> Groups**

1. Select a group and navigate to Applications Tab and click Assign applications and assign the Snyk app to the group
2. The Snyk app should be present and click on the pencil next to the snyk app\
   ![](https://lh3.googleusercontent.com/X3ARoW\_GPcKqIvrowKVPGnBWgziZ4E87hCIRVXzvLkLXOuLvP9fS5y9D-yaCjyWCmr6Co-\_3JSA2ZS-MdM5gEF9JYRi4Ivid-tnijtkpQstm7XgFbhAlnZnRM9D1DKYUsnHm987R)
3. Add the Snyk organization slugs that you want to associate with your OKTA group (there should NOT be any spaces or capital letter(s) **NOTE**: In this configuration, you add the role within this Snyk Org Slugs assignment\
   \
   ![](https://lh4.googleusercontent.com/qUN0SI64WQqAGCs2YPrvIW0lyZAyZDnGgpYe\_mXyGIPa2XqgBJJa3DBpg\_qGdoHxXql7kNrzrBkzY7T660es0qGcSH5wSbBw1DANk9f1\_q6SHDQXjxNFKRaVVCuZICVkFbnGYUz6)
4. Repeat steps 1-3 for all your applicable OKTA groups, as this will assign both the same orgslug and role to each user within the configured group.

### Implementation

1. Navigate to Applications -> Applications and click on the Snyk app you set
2. General Tab -> Edit SAML Settings and click next to go to the Configure SAML step
3. Set an Attribute Statement named “role” of an unspecified type
4.  Attribute Statements-> set **role** as the Name field with Name format **Unspecified** and the Value will be the below expression.

    `Arrays.flatten(String.replace(String.append("snyk-",Arrays.toCsvString(appuser.snyk_orgs)),",",",snyk-"))`

### **Explanation of the value expression**

1.  The end result should look like the below. The expression in the **role** attribute automatically adds the 'snyk' in front of the snyk\_orgs value (which contains \<orgslug>-\<role> **OR** \<groupadmin>) such that the end result should look similar to the below example.\


    Example result: `[“snyk-org1-admin”,”snyk-org2-admin”, “snyk-org3-collaborator”,”snyk-org4-collaborator”,”snyk-groupadmin”]`
