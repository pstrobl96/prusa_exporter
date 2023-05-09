### Old way of instalation

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
