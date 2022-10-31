# Example: Setting up OIDC Okta app with custom claims

## Configuration - OIDC Okta

1.  Create Okta OIDC app.

    ![Create new app integration](../../../../.gitbook/assets/Pasted\_Image\_6\_30\_22\_\_5\_01\_PM.png)


2.  Fill in Redirect URI.\
    ![Redirect URI](../../../../.gitbook/assets/Pasted\_Image\_6\_30\_22\_\_5\_10\_PM.png)\
    Find the [OIDC information to provide to Snyk](https://docs.snyk.io/features/user-and-group-management/setting-up-sso-for-authentication/set-up-snyk-single-sign-on-sso#oidc-information-to-provide-to-snyk). The information follows.


3.  Add Roles in Okta.\
    a. Add OIDC claim at the user level.

    ![](../../../../.gitbook/assets/Pasted\_Image\_7\_14\_22\_\_12\_16\_PM.png)

    ![OIDC claim at user level](../../../../.gitbook/assets/Pasted\_Image\_7\_14\_22\_\_12\_26\_PM.png)

    ![](../../../../.gitbook/assets/Pasted\_Image\_7\_14\_22\_\_12\_27\_PM.png)



    b. Add OIDC claim at the group level.

    ![OICD claim at group level steps 1 and 2](../../../../.gitbook/assets/Pasted\_Image\_7\_11\_22\_\_6\_12\_PM.png)

    ![OICD claim at group level step 3](../../../../.gitbook/assets/Pasted\_Image\_7\_15\_22\_\_5\_23\_PM.png)

    ![OIDC claim at group level step 4](../../../../.gitbook/assets/Pasted\_Image\_7\_15\_22\_\_5\_24\_PM.png)

    ![OIDC claim at group level step 5](../../../../.gitbook/assets/Pasted\_Image\_7\_15\_22\_\_5\_26\_PM.png)

    ![OIDC claim at group level step 6](../../../../.gitbook/assets/Pasted\_Image\_7\_15\_22\_\_5\_28\_PM.png)

    ![OIDC claim at group level step 7](../../../../.gitbook/assets/Pasted\_Image\_7\_15\_22\_\_5\_31\_PM.png)

    ![OIDC claim at group level step 8](../../../../.gitbook/assets/Pasted\_Image\_7\_15\_22\_\_5\_32\_PM.png)

    ![OIDC claim at group level step 9](../../../../.gitbook/assets/Pasted\_Image\_7\_15\_22\_\_5\_35\_PM.png)

    ![OIDC claim at group level step 10](../../../../.gitbook/assets/Pasted\_Image\_7\_15\_22\_\_5\_36\_PM.png)

    ![OIDC claim at group level steps 11 and 12](<../../../../.gitbook/assets/Pasted\_Image\_7\_15\_22\_\_5\_36\_PM (1).png>)
4.  Create a new claim.\
    ![Create new claim](../../../../.gitbook/assets/Pasted\_Image\_7\_11\_22\_\_6\_19\_PM.png)


5. Set attributes.
   1. Name: roles
   2. Include in token type: ID Token, Always
   3. Value type: Expression
   4. Value: Add your attribute
   5.  Include in: Set the scope that will be passed to Snyk\


       ![Edit claim to set attributes](<../../../../.gitbook/assets/Untitled (2).png>)



