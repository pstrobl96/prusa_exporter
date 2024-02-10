### SYSLOG metrics and logs

Here you can find examples of data that printers produces via SYSLOG. In metrics.log you'll find data that are produced when only metrics endpoint is configured. In logs_metrics.log you'll find data that are produced when both logs and metrics endpoints are active.

Note:
logs_metrics.log - XL
metrics.log - MK3.9

## Sample of parsed data

Parsed data looks like this

```
dwarf_board_temp: 37
active_extruder: 1
buddy_bom: 7
loadcell_scale: 0.019200
loadcell_threshold_cont: -40.000000
media_prefetched: 7057
splitter_5V_current: 0.241136
bed_curr,n=0: 1.816
points_dropped: 0
loadcell_threshold: -125.000000
loadcell_hysteresis: 80.000000
cpu_usage: 37
dwarf_mcu_temp: 35
adj_z: 0.000000
buddy_revision: 7
mac: XX:XX:XX:XX:XX:XX
bed_curr,n=1: 0.306
is_printing: 1
curr_nozz: 1.424067
buddy_bom: 40
adj_z: 0.000000
loadcell_scale: 0.019200
oc_inp: 0
loadcell_hysteresis: 80.000000
volt_bed: 24.026344
volt_nozz: 23.882259
curr_inp: -3.329260
cpu_usage: 26
loadcell_threshold: -125.000000
oc_nozz: 0
loadcell_threshold_cont: -40.000000
heater_enabled: 1
media_prefetched: 7915
points_dropped: 38
fw_version: 5.1.1
cur_mmu_imp: -0.013000
is_printing: 1
buddy_revision: 37
```