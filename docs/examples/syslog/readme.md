### SYSLOG metrics and logs

Here you can find examples of data that printers produces via SYSLOG. In metrics.log you'll find data that are produced when only metrics endpoint is configured. In logs_metrics.log you'll find data that are produced when both logs and metrics endpoints are active.

Note:
logs_metrics.log - XL
metrics.log - MK3.9

## Sample of parsed data

Parsed data looks like this

```
splitter_5V_current: 0.485248
active_extruder: 0
xlbuddy5VCurrent: 0.479294
dwarf_board_temp: 40
bed_curr,n=0: 1.909
bed_curr,n=1: 0.385
cpu_usage: 47
loadcell_scale: 0.019200
dwarf_mcu_temp: 37
Sandwitch5VCurrent: 0.878210
loadcell_threshold: -125.000000
loadcell_hysteresis: 80.000000
5VVoltage: 5.043011
24VVoltage: 24.098385
fw_version: 5.1.1
media_prefetched: 7010
adj_z:0.000000
buddy_bom:14
buddy_revision:27
cpu_usage:17
cur_mmu_imp:-0.003726
curr_inp:9.405615
curr_nozz:0.396320
heap:63008,89636
heater_enabled:1
is_printing:0
loadcell_age:-3141
loadcell_hp:0.000000
loadcell_value:-5141.760254
loadcell_xy:0.000000
mac: XX:XX:XX:XX:XX:XX
oc_inp:0
oc_nozz:0
points_dropped:25
temp_brd:38.253677
temp_mcu:45
volt_bed:24.134409
volt_nozz:23.954302
```