# Error executing binary because of corporate policy (Windows)

## **Problem**

Snyk is trying to run files from AppData but that is not allowed under the customer's corporate policy.

## Solution

As part of operations of the Snyk CLI, it extracts different files and executables to a temporary folder, which for Windows is located in C:\Users\\\<userprofile>\AppData\\`.`

If running files from this location is not allowed by group policy or other security restrictions, set the environment variable `SNYK_CACHE_PATH` on each machine. This will override the directory Snyk uses for the cache.
