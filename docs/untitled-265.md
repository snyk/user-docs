# Broker support: Manifest file size

##  Broker support: Manifest file size

While testing and importing files from the Broker make sure that file size is less than 1 MB. With the broker we cannot process files that are larger than 1 MB because it basically stops customers from controlling what files we can access - we would be able to access every file if we wanted through the big files API.

To address this issue, an additional Blob API endpoint should be whitelisted in `accept.json`. This should be in a private array

{  
 "//": "used to get given manifest file",  
 "method": "GET",  
 "path": "/repos/:owner/:repo/git/blobs/:sha",  
 "origin": "https://${GITHUB\_TOKEN}@${GITHUB\_API}"  
}

Reference: [Broker troubleshooting](https://github.com/snyk/broker#troubleshooting)

