# Configuration of Grafana Agent / Prometheus / Promtail / Loki

## agent.yml

Grafana Agent is used in Raspberry Pi image and currently works only with Grafana Cloud - if you don't configure it different way. You need to change `url`, `username` and `password`. You can get these values in configuration of your Grafana Cloud. How you can find in [Grafana Cloud documentation](https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-prometheus/). Example configuration can be found in [agent.yaml](examples/config/grafana_cloud/agent.yaml)

```
metrics:
  global:
    scrape_interval: 15s
    remote_write:
    - url: <YOUR CLOUD METRICS URL>
      basic_auth:
        username: "<YOUR CLOUD METRICS USERNAME>"
        password: "<YOUR CLOUD METRICS PASSWORD>"
```

If you want to parse logs from printer you can use this pipeline

```
logs:
  positions_directory: /var/lib/grafana-agent
  configs:
  - name: prusa
    clients:
    - url: <url>
    scrape_configs:
    - job_name: system
      pipeline_stages:
      static_configs:
      - labels:
          job: varlogs
          __path__: /var/log/*.log
    - job_name: prusa_syslog
      pipeline_stages:
        - json:
            expressions:
              message:
              stream:
              time:
              level:
              app_name:
              client:
              hostname:
        - labels:
            stream:
            time:
            level:
            app_name:
            client:
            hostname:
        - output:
            source: message
      static_configs:
      - labels:
          stream: 'stdout'
          job: prusa_printers_logs
          __path__: /var/log/prusa/exporter.log
```

## prometheus.yml

In [prometheus.yml](examples/config/on_premise/prometheus.yml) you need to change the `remote_write` section. This section is responsible for writing data to Grafana Cloud instance. You can get all values in config of your Grafana instance. You can get more information in [Grafana Docs](https://grafana.com/docs/grafana-cloud/data-configuration/metrics/metrics-prometheus/).

| key      | value                                  |
|----------|----------------------------------------|
| url      | this is where your instance is running |
| username | name that is used for login            |
| password | unique key used for login              |

```
remote_write:
- url: https://prometheus-prod-01-eu-west-0.grafana.net/api/prom/push
  basic_auth:
    username: userName
    password: apiKey
```

## promtail.yml

In [promtail.yml](examples/config/on_premise/promtail.yml) you need to change the `clients` section. Thanks to this block promtail will send logs to your Grafana Cloud Loki instance instead of local Loki. More details of log ingestion in [Grafana docs](https://grafana.com/docs/grafana-cloud/data-configuration/logs/collect-logs-with-promtail/).

| key      | value                                                 |
|----------|-------------------------------------------------------|
| url      | this is string that you can generate in Grafana Cloud |

```
clients:
  - url: https://<User Name>:<Your Grafana.com API Key>@logs-prod-eu-west-0.grafana.net/loki/api/v1/push
```
