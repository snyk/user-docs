# Auth

## SYNOPSIS

`snyk` `auth` \[\<API\_TOKEN>] \[]

## DESCRIPTION

Authenticate Snyk CLI with a Snyk account. Running `$ snyk auth` without an \<API\_TOKEN> will open a browser window and asks you to login with Snyk account and authorize. When inputting an \<API\_TOKEN>, it will be validated with Snyk API.

When running in a CI environment \<API\_TOKEN> is required.

## OPTIONS

* \[\<API\_TOKEN>]: Your Snyk token. May be an user token or a service account.\
  [How to get your account token](https://snyk.co/ucT6J)\
  [How to use Service Accounts](https://snyk.co/ucT6L)\
