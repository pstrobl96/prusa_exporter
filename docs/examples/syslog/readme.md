### SYSLOG metrics and logs

Here you can find examples of data that printers produces via SYSLOG. In metrics.log you'll find data that are produced when only metrics endpoint is configured. In logs_metrics.log you'll find data that are produced when both logs and metrics endpoints are active.

Note:
logs_metrics.log - XL
metrics.log - MK3.9

## Sample of parsed data

Parsed data looks like this

```
loadcell_threshold: -125.000000
mac: XX:XX:XX:XX:XX:XX
bed_curr,n=1: 0.338
is_printing: 1
splitter_5V_current: 0.512041
active_extruder: 1
dwarf_mcu_temp: 35
loadcell_scale: 0.019200
bed_curr,n=0: 1.836
media_prefetched: 7136
cpu_usage: 38
loadcell_threshold_cont: -40.000000
loadcell_hysteresis: 80.000000
buddy_bom: 7
dwarf_board_temp: 37
points_dropped: 0
adj_z: 0.000000
buddy_revision: 7
adj_z: 0.000000
curr_nozz: 1.635662
oc_inp: 0
loadcell_scale: 0.019200
buddy_bom: 40
media_prefetched: 6417
loadcell_threshold_cont: -40.000000
points_dropped: 38
fw_version: 5.1.1
cpu_usage: 20
heater_enabled: 1
cur_mmu_imp: -0.013000
oc_nozz: 0
is_printing: 1
buddy_revision: 37
volt_bed: 23.882259
volt_nozz: 23.810213
curr_inp: -8.750678
loadcell_threshold: -125.000000
loadcell_hysteresis: 80.000000
```