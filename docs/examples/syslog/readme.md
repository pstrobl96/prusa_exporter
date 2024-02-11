### SYSLOG metrics and logs

Here you can find examples of data that printers produces via SYSLOG. In metrics.log you'll find data that are produced when only metrics endpoint is configured. In logs_metrics.log you'll find data that are produced when both logs and metrics endpoints are active.

Note:
logs_metrics.log - XL
metrics.log - MK3.9

## Sample of parsed data

Parsed data looks like this

```
points_dropped: 0
splitter_5V_current: 0.485248
active_extruder: 0
xlbuddy5VCurrent: 0.479294
dwarf_board_temp: 40
bed_curr,n=0: 1.909
bed_curr,n=1: 0.385
cpu_usage: 47
is_printing: 1
adj_z: 0.000000
loadcell_scale: 0.019200
dwarf_mcu_temp: 37
heap: 44408,76536
Sandwitch5VCurrent: 0.878210
loadcell_threshold: -125.000000
loadcell_hysteresis: 80.000000
5VVoltage: 5.043011
24VVoltage: 24.098385
volt_bed: 24.206451
volt_nozz: 0.072043
cur_mmu_imp: -0.003726
heater_enabled: 0
oc_nozz: 0
oc_inp: 0
buddy_revision: 105
fw_version: 5.1.1
curr_nozz: 0.493721
media_prefetched: 7010
buddy_bom: 40
curr_inp: 8.932606
mac: XX:XX:XX:XX:XX:XX
```