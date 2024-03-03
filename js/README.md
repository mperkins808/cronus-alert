# cronus-alert

The easiest way to get push notifications for events that you define

### 1. Install

```bash
npm i cronus-alert
```

### 2. Obtain an API key from [cronusmonitoring.com](https://cronusmonitoring.com)

### 3. Example usage

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
