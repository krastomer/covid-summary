# Simple JSON API to summarize COVID-19 stats

API created by Follow-up Question from **LINE MAN Wongnai**.
they can show summary covid stats from Province and AgeGroup who covid infected.

## How to use

Run this command in terminal

```
$ go run cmd/server/main.go
```

Then http://localhost:8080/covid/summary you can get data like

```json
{
  "Province": {
    "Amnat Charoen": 17,
   ...
  },
  "AgeGroup": {
    "0-30": 602,
    "31-60": 607,
    "61+": 769,
    "N/A": 22
  }
}
```

This Project use Clean Architecture for development. and Gin is a webframework.

Created by Kasama Thongsawang
