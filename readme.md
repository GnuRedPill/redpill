Using tidb and grafana to monitor and record life

1. using the chrome plugin to import tidb
2. use github action to parse the data
3. using grafana to display data and pop up warning messages
4. TODO:: Using big models to analyse specific sentiment and correlation in logs



powershell 

```powershell

$Env:tidb_host=""

$Env:tidb_dsn=""


```

bash


```bash

export tidb_host=""

export tidb_dsn=""

```

start web api 

```bash
go run app.go

```

init table and import dict 

```bash

go run migrations

```


dev env run task

```bash

#win
go run  .\crontabs\cutrlog\ 
#linux
go run  crontab/cutrlog

```

edit chrome Extension url

manifest.json content-script.js 
in  http://127.0.0.1:8080/
The default is the local test environment

download https://dl.grafana.com/oss/release/grafana-10.0.3.windows-amd64.zip 
or https://dl.grafana.com/oss/release/grafana-10.0.3.linux-amd64.tar.gz
Create a new grafana service and then import (doc/grafana) grafana to access tidb cloud
Import grafana templates to see your reading habits and set up alerts.
