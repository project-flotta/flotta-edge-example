# Edge Example: Workload Application

## Before you begin
you need to have:
- Docker installed
- Docker-compose installed

## Stats
as this app will be running in edge device resource consumption will be important.
the following stats came from `docker stats` command:
```
docker stats --no-stream --format \
    "{\"container\":\"{{ .Container }}\",\"memory\":{\"raw\":\"{{ .MemUsage }}\",\"percent\":\"{{ .MemPerc }}\"},\"cpu\":\"{{ .CPUPerc }}\"}" | jq
```
```json
{
  "container": "957ccb2068d0",
  "memory": {
    "raw": "4.441MiB / 7.553GiB",
    "percent": "0.06%"
  },
  "cpu": "0.00%"
}
```

## For Developers
run the following commands to get started:
```
make run
```

to make a production build run:
```
make build-prod-image
```
```
make push-prod-image
```

## Future Works
- Add more sensors to the app, like CPU usage, memory usage, etc.
- Upload the data to Prometheus or Thanos as it provides more features for data visualization and analysis.