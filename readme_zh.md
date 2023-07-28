使用tidb和grafana监控记录生活

1. 使用chrome插件 导入 tidb
2. 使用github action 解析数据
3. 使用grafana 展示数据和弹出警告信息
4. TODO:: 使用大模型分析日志中的具体情感体现和相关度


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

启动web服务

```bash
go run app.go

```

建表导入数据等

```bash

go run migrations

```


模拟github action的离线任务

```bash

#win
go run  .\crontabs\cutrlog\ 
#linux
go run  crontab/cutrlog

```

修改chrome 插件的url

manifest.json content-script.js 
中的 http://127.0.0.1:8080/
默认是本机测试环境

新建grafana服务，然后导入grafana接入tidb cloud
导入grafana模板，查看自己的阅读习惯和设置告警。