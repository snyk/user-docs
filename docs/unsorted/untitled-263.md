# Alternate way to search Snyk vulnerabilities database

Searching the Snyk vulnerability database may not always the expected results as package names have subtle nuances that can cause the difference between its name in your project and its name in our database.

Another way to look for vulnerability details about a package in our database is through the URL. You can formulate a URL by combining the portions like the example below:

`Vuln DB URL/` + `Package Manager:` + `Package Name`

`https://snyk.io/vuln/` + `maven:` + `com.datastax.cassandra:cassandra-driver-core` which gives you: [https://snyk.io/vuln/maven:com.datastax.cassandra:cassandra-driver-core](https://snyk.io/vuln/maven:com.datastax.cassandra:cassandra-driver-core)

