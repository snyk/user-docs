# Does snyk support multi factor authentication \(MFA\) when customers set up SSO \(Single sign on\)

##  Does snyk support multi factor authentication \(MFA\) when customers set up SSO \(Single sign on\)

Yes, although the MFA \(Multi-factor authentication\) part is handled by the customer's identity provider and not something that Snyk actually has to do. We get the go, no go, from identity provider after it handles the second part of the authentication flow. Only once we have the 'Go' do we sign the user up/log them in.

