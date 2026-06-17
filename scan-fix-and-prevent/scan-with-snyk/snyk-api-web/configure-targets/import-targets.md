# Import targets

Import multiple targets to your Snyk API & Web account using JSON, CSV, or YAML files instead of adding targets individually through the interface.

## Prerequisites

Prepare a file containing target information in one of the supported formats:
- JSON
- CSV
- YAML

## Required fields

Each target must include:
- **Name**: Target name
- **URL**: Target URL

## Optional fields

Configure targets during import by including optional settings:

- Seeds list (allowed paths)
- Reject list (blocked URLs)
- Custom headers
- Custom cookies
- Scan profile
- Report type
- Authentication settings
- Labels

Availability of these settings varies depending on the target type. For complete field specifications, visit the [Snyk API & Web API documentation](https://developers.probely.com/api/reference/targets).

## Import targets

1. Log in to Snyk.
2. On the **Targets** list, click the vertical ellipsis next to **Add**.
3. Select **Import** from the menu.
4. Upload your file containing the target information.

## Import targets with scheduled scans

Schedule future scans when importing targets by including a date and recurrence in your import file:

- **Hourly**: `h`
- **Daily**: `d`
- **Weekly**: `w`
- **Monthly**: `m`
- **Quarterly**: `q`

This approach sets up scheduled scans without using the interface.

### JSON file with one target and monthly scheduled scans

```json
[
  {
    "name": "example",
    "site": {
      "name": "example",
      "url": "https://example.com"
    },
    "scan_profile": "normal",
    "scheduledscan_datetime": "2025-11-01T10:20:05.224739",
    "scheduledscan_recurrence": "m"
  }
]
```

### CSV file with two targets and different scheduled scans

```csv
"example","https://example.com","normal","2025-06-01T07:00:00.224739","m","['label1', 'label2']"
"example","https://example.com","safe","2025-09-05T19:45:00.224739","d","['label3', 'label4']"
```

### YAML file with one target and daily scheduled scans

```yaml
- name: example
  site:
    name: example
    desc: this is an example input
    url: https://example.com
  type: single
  scheduled_scan:
    date_time: "2025-03-30T16:29:08"
    recurrence: "d"
```
