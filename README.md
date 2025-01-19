[![docker](https://img.shields.io/github/actions/workflow/status/pstrobl96/prusa_exporter/docker.yml)](https://github.com/pstrobl96/prusa_exporter/actions/workflows/docker.yml) 
[![rpi](https://img.shields.io/github/actions/workflow/status/pstrobl96/prusa_exporter/rpi.yml)](https://github.com/pstrobl96/prusa_exporter/actions/workflows/rpi.yml) 
![issues](https://img.shields.io/github/issues/pstrobl96/prusa_exporter) 
![go](https://img.shields.io/github/go-mod/go-version/pstrobl96/prusa_exporter) 
![tag](https://img.shields.io/github/v/tag/pstrobl96/prusa_exporter) 
![license](https://img.shields.io/github/license/pstrobl96/prusa_exporter)

# THE REPOSITORY IS UNDER RECONSTRUCTION

# prusa_exporter

If you've seen this repository before, you've probably noticed some minor changes. Basically I removed most of the features because `feature-creep` was getting worse and worse and I'm aiming for a simpler setup and much higher quality code, so version 1.0.0 is skipped and `Vistaized` - the first final version will be 2.0.0.

- [ ] [prusa_metric_handler](https://github.com/pstrobl96/prusa_metrics_handler) integration for getting syslog metrics
- [ ] [prusa_log_processor](https://github.com/pstrobl96/prusa_log_processor) integration for log processing
- [ ] [prusa_exporter](https://github.com/pstrobl96/prusa_exporter) to process metrics from Prusa Link in addition to logs and syslog metrics. It's like a package for all three components.

---

# OLD README.md

This is an implementation of Prometheus Exporter for Prusa printers running Buddy boards (Prusa MK4, XL, and Mini), Einsy boards (Prusa MK3(S(+)) with Prusa Link installed) or resin printers (SL1). Multi-target is supported out of the box so you can check any number of printers as long it has accessible Prusa Link API and you have enough computing power.

For MK3S with Einsy board you need to use at least version 0.7.0 of Prusa Link or higher, because there are many more metrics to scrape than in the older versions. You can find the most up to date version in the [Prusa Link repository](https://github.com/prusa3d/Prusa-Link/releases).

- [THE REPOSITORY IS UNDER RECONSTRUCTION](#the-repository-is-under-reconstruction)
- [prusa\_exporter](#prusa_exporter)
- [OLD README.md](#old-readmemd)
  - [Where to find prusa exporter](#where-to-find-prusa-exporter)
  - [Roadmap](#roadmap)
  - [How to install prusa exporter](#how-to-install-prusa-exporter)
    - [Git Clone](#git-clone)
    - [Docker Compose](#docker-compose)
    - [Logs](#logs)
    - [Metrics](#metrics)
    - [Raspberry Pi](#raspberry-pi)
    - [Starting](#starting)
  - [Grafana Dashboards](#grafana-dashboards)
    - [Prusa Link](#prusa-link)
    - [Syslog](#syslog)
    - [Overview](#overview)

## Where to find prusa exporter

Prusa exporter runs on port 10009, but you can choose different port in `prusa.yml`. Metrics are accessible at `/metrics` endpoint.

## Roadmap

This list contains current and future features along with completion status:

- [x] Scrape of metrics from [Prusa Link](https://github.com/prusa3d/Prusa-Link/tree/0.7.0rc3)
- [x] Use of Grafana Cloud
- [x] CI pipeline with Docker Hub publish
- [x] Local instance of Grafana / Prometheus / Loki
- [x] Raspberry Pi Image
- [ ] Support for [connection](#21) to Einsy with username and password
- [x] Support for MK3 - it was implemented before but I want overhaul it and make it work
- [x] Dashboard update
- [x] Configuration update
- [x] Send logs to Grafana Cloud
- [x] Enable node_exporter for Grafana Cloud
- [x] Optimize and get more syslog metrics
- [ ] Automatically send syslog config gcode to buddy boards 
- [ ] exporter toolkit implemenation
- [ ] Create endpoint for configuration update
- [ ] Unit tests
- [ ] Create systemd service for exporter and install script
- [ ] Properly provision on premise setup
- [ ] CI for binaries release
- [x] Enable log collection to Loki
- [x] SL1 support

## How to install prusa exporter

### Git Clone

First things first. You need to clone the repo and that which is very easy, right?

`git clone https://github.com/pstrobl96/prusa_exporter.git`

### Docker Compose

I've created docker-compose.yaml file, that can be used for deploy of exporter. You would need [Docker](https://docs.docker.com/engine/install/) and [docker-compose](https://docs.docker.com/compose/install/linux/) plugin installed. Right now it is possible to use `docker compose up` or if you want to try newest codebase, then just run `sudo docker compose -f docker-compose.testing.yaml up --build  --force-recreate` that will build new docker image of exporter with Grafana Agent.

### Logs

Printer logs and detailed metrics are sent via Syslog which is not best option but it is what it is. However to get data to Loki I need to process logs in exporter. I wanted to use Promtail for forwarding the logs to Loki however I was not successful with this approach because Promtail was throwing EOL errors so I forward logs into file in exporter and I configured Promtail to scrape and parse them. You can find how to configure logs in [config.md](docs/config.md) and [exporter.md](docs/exporter.md) 

### Metrics

Metrics that you can find in this exporter are "scraped" from two sources. First is Prusa Link, it is pretty usual REST API that returns all data in JSON. There is a lot of useful metrics but there are few that are missing. Like data from most of sensors and for example current or voltage. However this is not applicable to **Einsy printers like MK3, these supports only Prusa Link API. As well as resin printers like SL1.** You can find how to configure metrics in [config.md](docs/config.md) and [exporter.md](docs/exporter.md) 

For Buddy - SYSLOG exists. [Syslog](https://en.wikipedia.org/wiki/Syslog) is standard for logging for a quite while however printer used it for sending metrics. Trough UDP. So what I just did is that I created experimental Syslog UDP server within this exporter and I'm catching these "metrics". **Be aware that these metrics can be send only via wired ethernet. You are out of luck over the air.**

**The issue is that if you have more printers you'll create a lot of UDP traffic in the network.** If you have more printers this number multiplies. I choose flag this feature experimental because you cannot be sure you'll get the metrics, it's UDP and printers are sending data as much as they can but it is not consistent. Between printers there are differences - obvisouly. 

Example how metrics looks can be found in ![this](docs/examples/metrics_example.md) file. This file also includes Einsy and Buddy syslog metrics.

How to config Syslog metrics you can find in ![documentation](docs/syslog.md).

### Raspberry Pi

I also created Raspberry Pi image that can be flashed to memory card. If you choose this path you'll need following.

- Raspberry Pi (*4 and 5 tested*) with 64 bit support
- At least *Class 10* and at least *16 gigs* Memory card - preferably some kind of durable one

Of course all other accessories like computer, card reader, power supply etc. are mandatory. How to flash Raspberry Pi image you can find in ![documentation](docs/rpi_image.md)

### Starting

Starting of exporter is simple. Just change directory to where docker-compose.yaml and configs are and run following command.

```
docker compose up

```

:tada: if everthing went alright your instance is up and running and you can find metrics at [/metrics](http://localhost:10009/metrics) endpoint.

## Grafana Dashboards

I also prepared one dashboard per board which you can find in the [docs/examples/grafana](docs/examples/grafana) folder.

### Prusa Link

Download this dashboard straight from [Grafana.net](https://grafana.com/grafana/dashboards/20393)! Just use ID `20393` when importing.  

![dashboard](docs/examples/grafana/prusalink.png)

### Syslog

Download this dashboard straight from [Grafana.net](https://grafana.com/grafana/dashboards/20618)! Just use ID `20618` when importing.  

![dashboard](docs/examples/grafana/syslog.png)

### Overview

This dashboard is used for monitoring all of your printers. Basically - green means printing, blue means ready, yellow means warning and red is error. You need [polystat panel](https://github.com/grafana/grafana-polystat-panel) for this dashboard.

Download this dashboard straight from [Grafana.net](https://grafana.com/grafana/dashboards/20449)! Just use ID `20449` when importing.  

![dashboard](docs/examples/grafana/overview.png)