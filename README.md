[![ci](https://github.com/pstrobl96/prusa_exporter/actions/workflows/ci.yml/badge.svg)](https://github.com/pstrobl96/prusa_exporter/actions/workflows/ci.yml) ![issues](https://img.shields.io/github/issues/pstrobl96/prusa_exporter) ![go](https://img.shields.io/github/go-mod/go-version/pstrobl96/prusa_exporter) ![tag](https://img.shields.io/github/v/tag/pstrobl96/prusa_exporter) ![license](https://img.shields.io/github/license/pstrobl96/prusa_exporter)

# Prusa Exporter - formerly Buddy Link Prometheus Exporter

This is an implementation of Prometheus Exporter for Prusa printers running Buddy boards (Prusa MK4, XL, and Mini) or Einsy boards (Prusa MK3(S(+)) with Prusa Link installed). Multi-target is supported so you can check any number of printers as long it has accessible Prusa Link API. This works even for the old Prusa Connect Local.

For Mk3s with Einsy board you need to use at least version 0.7.0rc3 of Prusa Link or higher, because there are many more metrics to scrape than in the older versions. You can find the most up to date version in the [Prusa Link repository](https://github.com/prusa3d/Prusa-Link/releases).

- [Where to find prusa exporter](#where-to-find-prusa-exporter)
- [Roadmap](#roadmap)
- [How to install prusa exporter](#how-to-install-prusa-exporter)
   * [Git Clone](#git-clone)
   * [Docker Compose](#docker-compose)
   * [Raspberry Pi](#raspberry-pi)
      + [Downloading image ](#downloading-image)
      + [Raspberry Pi Imager](#raspberry-pi-imager)
         - [WiFi](#wifi)
         - [LAN](#lan)
      + [Flashing](#flashing)
      + [Config](#config)
   * [Config](#config-1)
      + [prusa.yml](#prusayml)
      + [agent.yml](#agentyml)
      + [prometheus.yml](#prometheusyml)
      + [promtail.yml](#promtailyml)
   * [Starting](#starting)
- [Grafana Dashboards](#grafana-dashboards)
   * [Buddy](#buddy)
   * [Legacy](#legacy)
   * [Einsy](#einsy)
   * [Overview - preview](#overview-preview)
- [Metrics example](#metrics-example)

## Where to find prusa exporter

Prusa exporter runs on port 10009, but you can choose different port in `prusa.yml`. Metrics are accessible at `/metrics` endpoint.

## Roadmap

This list contains current and future features along with completion status:

- [x] Scrape of metrics from [Prusa Link](https://github.com/prusa3d/Prusa-Link/tree/0.7.0rc3)
- [x] Use of Grafana Cloud
- [x] CI pipeline with Docker Hub publish
- [x] Local instance of Grafana / Prometheus / Loki
- [x] Raspberry Pi Image
- [ ] Odroid C4 Image
- [ ] Implementation with [exporter-toolkit](#22)
- [ ] Support for [connection](#21) to Einsy with username and password

## How to install prusa exporter

### Git Clone

First things first. You need to clone the repo and that which is very easy, right?

`git clone https://github.com/prusa3d/Prusa-Firmware-Buddy.git`

### Docker Compose

I've created docker-compose.yaml file, that can be used for deploy of exporter. You would need [Docker](https://docs.docker.com/engine/install/) and [docker-compose](https://docs.docker.com/compose/install/linux/) plugin installed. Right now it is possible to use `docker compose up` only with Linux because I do not build image for Linux.

### Raspberry Pi

I also created Raspberry Pi image that can be flashed to memory card. If you choose this path you'll need following.

- Raspberry Pi (*4 and 5 tested*) with 64 bit support
- At least *Class 10* and at least *16 gigs* Memory card 

Of course all other accessories like computer, card reader, power supply etc. are mandatory. 

#### Downloading image 

Download image from [releases page](https://github.com/pstrobl96/prusa_exporter/releases) or alternatively you can choose CI pipeline run and download particular artifact. Downloaded *img.xz* file is named image_{date_of_build}-prusa_exporter-image.img.xz and needs to be flashed to the memory card.

#### Raspberry Pi Imager

[Download](https://www.raspberrypi.com/software/) and install Raspberry Pi Imager. You can alternatively use different tool but rpi-imager is easiest in terms of settings. 

![rpiimager0](docs/readme/rpiimager0.png)

After installing open the Raspberry Pi Imager. **Don't** click `Choose Device` instead of that click on `Choose OS`. Scroll down, you'll find `Use Custom`. Select downloaded image.

![rpiimager1](docs/readme/rpiimager1.png)

![rpiimager2](docs/readme/rpiimager2.png)

Now connect memory card to your computer and click `Choose Storage`. **BEWARE** - you can mistakenly choose wrong storage media and flashing process includes formating your drive. Now select your Raspberry Pi memory card. Now click `Next`.

![rpiimager3](docs/readme/rpiimager3.png)

Now it depends if you want to connect via LAN or WiFi.

##### WiFi

If you want to use wireless ethernet, then click at `Edit Settings`. Click at `Configure wireless LAN` and write your WiFi name (SSID) and password. Don't forget to select correct Wireless LAN country. Next be sure that `Eject media when finished` is **unchecked** . Click `Save` and after that click on `Yes`. If you are sure that all content of your memory card would be erased, click `Yes`.

![rpiimager4](docs/readme/rpiimager4.png)

![rpiimager7](docs/readme/rpiimager7.png)

![rpiimager6](docs/readme/rpiimager6.png)

##### LAN

If you want to use wired ethernet, then click at `Edit Settings`, click `Options` and be sure that `Eject media when finished` is **unchecked** . Click `Save` and after that click on `Yes`. If you are sure that all content of your memory card would be erased, click `Yes`.

![rpiimager7](docs/readme/rpiimager7.png)

![rpiimager6](docs/readme/rpiimager6.png)

#### Flashing

Now wait for the Raspberry Pi Imager to complete the flash process.

#### Config

Now we need to configure *Grafana Agent* and *prusa_exporter*. After flashing you should see new partition connected to system, can be called `boot` or `bootfs`. In Windows you'll get also letter of partition, nowadays most probably `D:` - can varies. If you don't see new partition. Eject memory card from the system and reconnect it. 

In boot partition you'll find two files `agent.yaml` and `prusa.yml`. Configuration is mentioned in next part of README.

### Config

Please take a look at the [sample configuration examples](docs/examples/config) for prusa exporter, Prometheus, and Promtail. You will need to change few things to get it up and running. Of course you can change everything you want. If you are using Grafana Cloud, you can find your API key at [grafana.com](https://grafana.com/) -> My Account -> Grafana Cloud instance -> Send Metrics / Send Logs.

I also prepared a configuration for on-premise Prometheus and Loki if you do not want to use Cloud solution and you want to have your data somewhere local. You can find these [on-premise configs](docs/examples/config/on_premise) in the on_premise subfolder.  

#### prusa.yml

Prusa exporter loads [prusa.yml](docs/examples/config/prusa.yml) from an environment variable called `$PRUSA_EXPORTER_CONFIG`. If you put this file in the same folder where prusa exporter is located then simply set it to `prusa.yml`. Prusa exporter has implemented a config reloader that runs by default every 300 seconds (5 minutes).

You will find two sections in the config file, `exporter` and `printers`.

`exporter` is used for configuration of exporter itself:

```
exporter:
  metrics_port: 10009 # exporter port
  scrape_timeout: 1 # scrape timeout of Prusa Link
  reload_inteval: 300 # interval in seconds for config reloader
  log_level: info
```

`metrics_port`: you can set whatever you want. It is the port where Prometheus would scrape metrics endpoint. **Required**

`scrape_timeout`: Value in seconds that implies timeout of scraping Prusa Link devices. Not necessary needed for Einsy but needed for Buddy becuase printer sometimes do not return values. **Required**

`reload_inteval`: Because feature of config reloading is implemeneted, you need to specify interval of reloading. **Required**

`log_level`: log level of logger, default is info. **Optional**

`printers` is used for configuring your target printers. Please note that `type` is informational and optional; if you define it it will be part of your metric labelset.

Note: Currently, you can not log into Einsy (Raspberry Pi Zero) boards with username and passwort. You need to generate an API key in Prusa Link settings. This will be resolved in a future release.

```
printers:
  buddy:
  - address: <address_of_printer>
    name: <your_printer_name>
    type: mini
    apikey: <your_printer_apikey>
  - address: <address_of_printer>
    username: maker # I'm not aware that there is posibility to change user name in XL or MK4 printers - default is maker
    pass: <password>
    name: <your_printer_name>
    type: <mini/xl/mk4> **optional**
  einsy:
  - address: <address_of_printer>
    apiKey: <your_printer_apikey>
    name: <your_printer_name>
    type: <mk2.5 or mk3> **optional**
  legacy:
  - address: <address_of_printer>
    name: <your_printer_name>
    type: mini **optional**
```

#### agent.yml

Grafana Agent is used in Raspberry Pi image and currently works only with Grafana Cloud - if you don't configure it different way. You need to change `url`, `username` and `password`. You can get these values in configuration of your Grafana Cloud. How you can find in [Grafana Cloud documentation](https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-prometheus/).

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

#### prometheus.yml

In [prometheus.yml](docs/examples/config/prometheus.yml) you need to change the `remote_write` section. This section is responsible for writing data to Grafana Cloud instance. You can get all values in config of your Grafana instance. You can get more information in [Grafana Docs](https://grafana.com/docs/grafana-cloud/data-configuration/metrics/metrics-prometheus/).

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

#### promtail.yml

In [promtail.yml](docs/examples/config/promtail.yml) you need to change the `clients` section. Thanks to this block promtail will send logs to your Grafana Cloud Loki instance instead of local Loki. More details of log ingestion in [Grafana docs](https://grafana.com/docs/grafana-cloud/data-configuration/logs/collect-logs-with-promtail/).

| key      | value                                                 |
|----------|-------------------------------------------------------|
| url      | this is string that you can generate in Grafana Cloud |

```
clients:
  - url: https://<User Name>:<Your Grafana.com API Key>@logs-prod-eu-west-0.grafana.net/loki/api/v1/push
```

### Starting

Starting of exporter is simple. Just change directory to where docker-compose.yaml and configs are and run following command.

```
docker compose up

```

:tada: if everthing went alright your instance is up and running and you can find metrics at [/metrics](http://localhost:10009/metrics) endpoint.

## Grafana Dashboards

I also prepared one dashboard per board which you can find in the [docs/examples/grafana](docs/examples/grafana) folder.

### Buddy

![dashboard](docs/examples/grafana/buddy.png)

### Legacy

![dashboard](docs/examples/grafana/legacy.png)

### Einsy

![dashboard](docs/examples/grafana/einsy.png)

### Overview - preview

![dashboard](docs/examples/grafana/overview.png)

## Metrics example

Example how metrics looks can be found in ![this](docs/examples/metrics_example.md) file.