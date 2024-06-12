# V1 API errors

The V1 API uses standard HTTP error codes for error responses.&#x20;

```json
{
    "code": 404,
    "message": "Org 39db46b1-ad57-47e6-a87d-e34f6968030b was not found or you may not have the correct permissions to access the org.",
    "error": "Org 39db46b1-ad57-47e6-a87d-e34f6968030b was not found or you may not have the correct permissions to access the org."
}
```

The error reference will also be supplied in the `x-error-reference` header in the server reply.

Example `500` response:

```
HTTP/1.1 500 Internal Server Error
x-error-reference: a45ec9c1-065b-4f7b-baf8-dbd1552ffc9f
Content-Type: application/json; charset=utf-8
Content-Length: 1848
Vary: Accept-Encoding
Date: Sun, 10 Sep 2017 06:48:40 GMT
```

