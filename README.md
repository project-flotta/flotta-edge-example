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