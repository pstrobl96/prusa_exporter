# Config

Please take a look at the [sample configuration examples](docs/examples/config) for prusa exporter, Prometheus, and Promtail. You will need to change few things to get it up and running. Of course you can change everything you want. If you are using Grafana Cloud, you can find your API key at [grafana.com](https://grafana.com/) -> My Account -> Grafana Cloud instance -> Send Metrics / Send Logs.

I also prepared a configuration for on-premise Prometheus and Loki if you do not want to use Cloud solution and you want to have your data somewhere local. You can find these [on-premise configs](docs/examples/config/on_premise) in the on_premise subfolder.  

## prusa.yml

Prusa exporter loads [prusa.yml](docs/examples/config/prusa.yml) from an command flag `--config.file=<path>`. This flag can be empty and if so exporter will just try to load `prusa.yml` file located in the executable folder. Prusa exporter has implemented a config reloader that runs by default every 300 seconds (5 minutes).

You will find two sections in the config file, `exporter` and `printers`.

`exporter` is used for configuration of exporter itself:

```
exporter:
  scrape_timeout: 1000 # scrape timeout of Prusa Link in ms
  log_level: info
  prusalink:
    enabled: true
  syslog:
    metrics:
      enabled: true
      listen_address: 0.0.0.0:10008
    logs:
      enabled: true
      listen_address: 0.0.0.0:10007
      directory: /var/log/prusa
      filename: exporter.log
      max_size: 10 # in MB
      max_age: 7 # in days
      max_backups: 10
```

`scrape_timeout`: value in seconds that implies timeout of scraping Prusa Link devices in miliseconds. **Required**

`log_level`: log level of logger, default is info. **Optional**

`prusalink.enabled`: you can enable or disable prusalink metrics **Required**

`syslog`: **EXPERIMENTAL** 

`syslog.metrics.enabled`: **EXPERIMENTAL** activates or deactivates printer syslog metrics handling. **Required**

`syslog.metrics.listen_address`: **EXPERIMENTAL** address where should syslog metrics server run. **Required if enabled**

`syslog.logs.enabled`: **EXPERIMENTAL** activates or deactivates printer logs handling. **Required**

`syslog.logs.listen_address`: **EXPERIMENTAL** address where should syslog log server run. **Required if enabled**

`syslog.logs.directory`: **EXPERIMENTAL** path where logs from printers should be stored. **Required if enabled**

`syslog.logs.filename`: **EXPERIMENTAL** name of file for logs. **Required if enabled**

`syslog.logs.max_size`: **EXPERIMENTAL** max size of log file. **Required if enabled**

`syslog.logs.max_age`: **EXPERIMENTAL** max age in Days of log file. **Required if enabled**

`syslog.logs.max_backups`: **EXPERIMENTAL** max number of backups left. **Required if enabled**

`printers` is used for configuring your target printers. 

Note: Currently, you can not log into Einsy (Raspberry Pi Zero) boards with username and password. You need to generate an API key in Prusa Link settings. This will be resolved in a future release.

It is recommended to also fill `field` type in configuration. Exporter can detect type automatically but it does not work with **Prusa Connect** and it would not detect the printer model correctly.

Allowed types are following

| Printer model      | type    |
|--------------------|---------|
| Prusa XL           | XL      |
| Prusa MK4          | MK4     |
| Prusa MK3.9        | MK3     |
| Prusa MK3.5        | MK35    |
| Prusa Mini         | MINI    |
| Prusa i3 MK3S(+)   | I3MK3S  |
| Prusa i3 MK3       | I3MK3   |
| Prusa i3 MK2.5S    | I3MK25S |
| Pursa i3 MK2.5     | I3MK25  |
| Prusa SL1          | SL1     |
| Prusa SL1S (Speed) | SL1S    |
| Prusa M1           | SL1S    |
| Prusa iX (AFS)     | IX      |

```
printers:
  - address: <address_of_printer>
    username: maker
    password: <password>
    name: <your_printer_name> # optional
    type: MINI # or MK35 / MK39 / MK4 / XL / IX
  - address: <address_of_printer>
    apikey: <apikey>
    name: <your_printer_name> # optional
    type: I3MK25 # or I3MK25S / I3MK3 / I3MK3S
```
