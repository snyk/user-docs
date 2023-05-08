# Broker rules for Snyk Code

Specific Broker rules are required to use Broker with Snyk Code.

The Broker allows easy injection of these rules via the `ACCEPT_CODE=true` environment variable.

This Helm Chart sets `ACCEPT_CODE=true` by default so no further action is needed for the rules pertaining to Snyk Code.
