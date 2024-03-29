# cronus-alert

The easiest way to get push notifications for events that you define

This repo contains the cronus-alert package in

1. Javascript
2. Go
3. Python

All packages require an API key from [cronusmonitoring.com](https://cronusmonitoring.com)

## Javascript

### 1. Install

```bash
npm i cronus-alert
```

### 2. Example usage

```javascript
function main() {
  // Set API Key to an environment variable
  const cronusAlert = new CronusAlert(process.env.CRONUS_API_TOKEN);

  cronusAlert
    .FireAlert(
      "The system is down", // Summary
      "Someone forgot to lock the gate", // Description
      Status.Firing // Optional status
    )
    .then((resp) => {
      console.log(resp.data); // API response
    })
    .catch((err) => {
      console.error(err); // network error
    });
}

main();
```

## Go

### 1. Install

```bash
go get github.com/mperkins808/cronus-alert/go/pkg/cronusalert
```

### 2. Example usage

```go
func main() {

  client := cronusalert.NewCronusAlertClient(os.Getenv("CRONUS_API_KEY"))
  alert := cronusalert.Alert{
      Summary:     "Issue detected in server",
      Description: "Failing to read database transactions",
      Status:      cronusalert.FIRING,
    }

  err := client.Fire(alert)
  if err != nil {
      // handle error
  }

}

```

## Python

### 1. Install

```bash
pip install cronusalert
```

### 2. Example usage

```python
from cronusalert import client

token = os.environ.get("CRONUS_API_KEY")
alertClient = client.CronusAlertClient(token)
alert = client.Alert("Issue detected in server", "Python server is burning down", client.AlertType.FIRING)
alertClient.fire(alert)

```
