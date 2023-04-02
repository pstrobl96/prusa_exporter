## Prusa Prometheus Exporter

This is implementation of Prometheus exporter for Prusas printers running Prusa Link Web - Like Prusa Mini or SL1. You can check any number of printers if you want to as long it has accesible Prusa Link API.

### printers.yaml

Exporter loads printers.yaml (file with connections to printers) from environment variable called **PRUSA_EXPORTER_PRINTERS**. If you want to put this file in folder, where exporter is located then just set it to *printers.yaml*.

#### Format of printers.yaml
```
printers:
  apiKeys:
  - address: <your_printer_ip>
    apikey: <your_prusa_link_apikey>
    name: <familiar_name_of_printer>
    type: <mini>
password:
  - address: <your_printer_ip>
    username: <your_prusa_link_username>
    password: <your_prusa_link_password>
    name: <familiar_name_of_printer>
    type: <xl_or_mk4>

```

### Where to find exporter

Exporter runs at port 10009, selected specificaly, because all of the official ports from range 9000-9999 is occupied.

### How to install exporter

#### Docker

// TODO

#### Helm chart

// TODO
