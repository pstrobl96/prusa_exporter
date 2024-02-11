### SYSLOG metrics and logs

Here you can find examples of data that printers produces via SYSLOG. In metrics.log you'll find data that are produced when only metrics endpoint is configured. In logs_metrics.log you'll find data that are produced when both logs and metrics endpoints are active.

Note:
logs_metrics.log - XL
metrics.log - MK3.9

## Sample of parsed data

Parsed data looks like this

```
5VVoltage: 5.043011
heap: 50212,76536
buddy_revision: 7
buddy_bom: 7
adj_z: 0.000000
points_dropped: 0
dwarf_board_temp: 24
bed_curr,n=0: 0.062
Sandwitch5VCurrent: 0.491202
is_printing: 0
splitter_5V_current: 0.235182
mac: XX:XX:XX:XX:XX:XX
24VVoltage: 24.350536
dwarf_mcu_temp: 25
xlbuddy5VCurrent: 0.449525
bed_curr,n=1: 0.060
cpu_usage: 25
active_extruder: 5
oc_nozz: 0
curr_inp: -0.709515
oc_inp: 0
cur_mmu_imp: -0.013000
heater_enabled: 0
fw_version: 5.1.1
volt_bed: 24.098385
volt_nozz: 0.000000
curr_nozz: 0.000000
```