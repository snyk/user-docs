# Examples for the Update existing integration endpoint

Examples follow for the [Update existing integration](../reference/integrations-v1.md#org-orgid-integrations-integrationid) endpoint in the [Integrations (v1)](../reference/integrations-v1.md) API.

## Set up Broker for an existing integration

### Command

```
curl --include \
     --request PUT \
     --header "Content-Type: application/json; charset=utf-8" \
     --header "Authorization: token API_KEY" \
     --data-binary "{
    \"type\": \"github\",
    \"broker\": { \"enabled\": true }
}" \
''
```

### Response

```
{
  "id": "9a3e5d90-b782-468a-a042-9a2073736f0b",
  "brokerToken": "4a18d42f-0706-4ad0-b127-24078731fbed"
}
```

## Update credentials for an existing non-brokered integration

### Command

```
curl --include \
     --request PUT \
     --header "Content-Type: application/json; charset=utf-8" \
     --header "Authorization: token API_KEY" \
     --data-binary "{
    \"type\": \"gitlab\",
    \"credentials\": { \"token\": \"GITLAB_TOKEN\" }
}" \
''
```

### Response

```
{
  "id": "9a3e5d90-b782-468a-a042-9a2073736f0b"
}
```

## Disable broker for an existing integration

### Command

```
curl --include \
     --request PUT \
     --header "Content-Type: application/json; charset=utf-8" \
     --header "Authorization: token API_KEY" \
     --data-binary "{
    \"type\": \"github\",
    \"broker\": { \"enabled\": false },
    \"credentials\": { \"token\": \"GITHUB_TOKEN\" }
}" \
''
```

### Response

```
{
  "id": "9a3e5d90-b782-468a-a042-9a2073736f0b"
}
```
