[![Go](https://github.com/pstrobl96/buddy-prometheus-exporter/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/pstrobl96/buddy-prometheus-exporter/actions/workflows/go.yml)

## Buddy Link Prometheus Exporter

This is an implementation of Prometheus Exporter for Prusa printers running boards named Buddy or Einsy (with Prusa Link installed) like Prusa MK4, XL, Mini or MK3S. Multi-target is supported so you can check any number of printers as long it has accessible Prusa Link API (Even the old Prusa Connect Local).

### buddy.yaml

Exporter loads buddy.yaml (file with connections to printers) from environment variable called **BUDDY_EXPORTER_CONFIG**. If you want to put this file in folder, where exporter is located then just set it to *buddy.yaml*. This file is loaded only at start of exporter so be sure restart it after change.

### Grafana Dashboard

I also prepared one dashboard per board that you can find in grafana folder.

#### Buddy

![dashboard](./grafana/buddy.png)

#### Legacy

![dashboard](./grafana/legacy.png)

#### Einsy

![dashboard](./grafana/einsy.png)

#### Format of buddy.yaml

In code block bellow you can see template for buddy.yaml config file. Type value is not that important, you can set anything you want. However this value would be written to labels in metrics, so be aware of that.

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
    type: <mini or xl>
  einsy:
  - address: <address_of_printer>
    apiKey: <your_printer_apikey>
    name: <your_printer_name>
    type: <mk2.5 or mk3>
  legacy:
  - address: <address_of_printer>
    name: <your_printer_name>
    type: mini

```

### Where to find exporter

Exporter runs at port 10009, but you can choose different port with `BUDDY_EXPORTER_PORT` environment variable.

### How to install exporter

#### Docker

I've made dockerfile. Enjoy

#### Old way

I've created shell script named [install_service.sh](install_service.sh). Copy its content to machine where you want to run exporter. Edit your buddy.yaml and you are good to go. You can also change `BUDDY_EXPORTER_PORT` variable to change where exporter should run.

```
printers:
  buddy:
  - address: 192.168.0.2
    name: printer1
    type: mini
    apikey: APIKEY
  - address: 192.168.0.3
    username: maker
    pass: PASSWORD
    name: printer2
    type: mk4
  einsy:
  - address: 192.168.0.4
    apikey: APIKEY
    name: printer3
    type: mk3
  - address: 192.168.0.5
    apikey: APIKEY
    name: printer4
    type: mk3
  legacy:
  - address: 192.168.0.6
    name: ol_but_reliable
    type: mini
```
