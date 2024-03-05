# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 1.8081e-05
go_gc_duration_seconds{quantile="0.25"} 2.6053e-05
go_gc_duration_seconds{quantile="0.5"} 2.8187e-05
go_gc_duration_seconds{quantile="0.75"} 3.3634e-05
go_gc_duration_seconds{quantile="1"} 0.000229939
go_gc_duration_seconds_sum 0.256905881
go_gc_duration_seconds_count 7532
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 30
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.21.7"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 2.147872e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 1.7618376752e+10
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 4336
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 1.04114742e+08
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 4.024872e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 2.147872e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 6.938624e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 3.907584e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 7569
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 5.5296e+06
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 1.0846208e+07
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.7096653653245385e+09
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 1.04122311e+08
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 24000
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 31200
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 336000
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 407400
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.194304e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 2.322352e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 1.736704e+06
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 1.736704e+06
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 1.9373072e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 20
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 94.52
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 1.073741816e+09
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 21
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 1.7993728e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.7096650018e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 1.268477952e+09
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes 1.8446744073709552e+19
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 73
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
# HELP prusa_active_extruder Active extruder - used for XL
# TYPE prusa_active_extruder gauge
prusa_active_extruder{ip="192.168.20.130",mac="10:9c:70:25:17:3d"} 1
# HELP prusa_ambient_temp Status of the printer ambient temp
# TYPE prusa_ambient_temp gauge
prusa_ambient_temp{printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1"} 22.7
# HELP prusa_axis Returns information about position of axis.
# TYPE prusa_axis gauge
prusa_axis{printer_address="192.168.20.130",printer_axis="x",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl"} 0
prusa_axis{printer_address="192.168.20.130",printer_axis="y",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl"} 0
prusa_axis{printer_address="192.168.20.130",printer_axis="z",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl"} 0
prusa_axis{printer_address="192.168.20.139",printer_axis="x",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39"} 0
prusa_axis{printer_address="192.168.20.139",printer_axis="y",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39"} 0
prusa_axis{printer_address="192.168.20.139",printer_axis="z",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39"} 0
prusa_axis{printer_address="192.168.20.157",printer_axis="x",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3"} 0
prusa_axis{printer_address="192.168.20.157",printer_axis="y",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3"} 210
prusa_axis{printer_address="192.168.20.157",printer_axis="z",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3"} 51
prusa_axis{printer_address="192.168.20.173",printer_axis="x",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef"} 0
prusa_axis{printer_address="192.168.20.173",printer_axis="y",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef"} 0
prusa_axis{printer_address="192.168.20.173",printer_axis="z",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef"} 0
# HELP prusa_axis_z_adjustment Axis Z adjustment
# TYPE prusa_axis_z_adjustment gauge
prusa_axis_z_adjustment{ip="192.168.20.130",mac="10:9c:70:25:17:3d"} 0
prusa_axis_z_adjustment{ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 0
prusa_axis_z_adjustment{ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 0
prusa_axis_z_adjustment{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 0
prusa_axis_z_adjustment{ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
# HELP prusa_bed_temp Current temp of printer bed in Celsius
# TYPE prusa_bed_temp gauge
prusa_bed_temp{printer_address="192.168.20.130",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl"} 23.4
prusa_bed_temp{printer_address="192.168.20.139",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39"} 90.1
prusa_bed_temp{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3"} 22.8
prusa_bed_temp{printer_address="192.168.20.173",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef"} 24.4
prusa_bed_temp{printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1"} 54
# HELP prusa_bed_temp_offset Offset bed temp
# TYPE prusa_bed_temp_offset gauge
prusa_bed_temp_offset{printer_address="192.168.20.130",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl"} 0
prusa_bed_temp_offset{printer_address="192.168.20.139",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39"} 0
prusa_bed_temp_offset{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3"} 0
prusa_bed_temp_offset{printer_address="192.168.20.173",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef"} 0
prusa_bed_temp_offset{printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1"} 0
# HELP prusa_bed_temp_target Target bed temp
# TYPE prusa_bed_temp_target gauge
prusa_bed_temp_target{printer_address="192.168.20.130",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl"} 0
prusa_bed_temp_target{printer_address="192.168.20.139",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39"} 90
prusa_bed_temp_target{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3"} 0
prusa_bed_temp_target{printer_address="192.168.20.173",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef"} 0
prusa_bed_temp_target{printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1"} 0
# HELP prusa_buddy_bom Buddy bom
# TYPE prusa_buddy_bom gauge
prusa_buddy_bom{ip="192.168.20.130",mac="10:9c:70:25:17:3d"} 7
prusa_buddy_bom{ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 40
prusa_buddy_bom{ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 14
prusa_buddy_bom{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 14
prusa_buddy_bom{ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 38
# HELP prusa_buddy_fw Buddy firmware version
# TYPE prusa_buddy_fw gauge
prusa_buddy_fw{ip="192.168.20.176",mac="10:9c:70:2d:17:2e",version="5.1.1"} 1
# HELP prusa_buddy_revision Buddy revision
# TYPE prusa_buddy_revision gauge
prusa_buddy_revision{ip="192.168.20.130",mac="10:9c:70:25:17:3d"} 7
prusa_buddy_revision{ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 37
prusa_buddy_revision{ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 27
prusa_buddy_revision{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 27
prusa_buddy_revision{ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 37
# HELP prusa_chamber_temp Status of the printer chamber temp
# TYPE prusa_chamber_temp gauge
prusa_chamber_temp{printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1"} 22.7
# HELP prusa_chamber_temp_offset Offset chamber temp
# TYPE prusa_chamber_temp_offset gauge
prusa_chamber_temp_offset{printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1"} 0
# HELP prusa_chamber_temp_target Traget chamber temp
# TYPE prusa_chamber_temp_target gauge
prusa_chamber_temp_target{printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1"} 0
# HELP prusa_cover Status of the printer - 0 = open, 1 = closed
# TYPE prusa_cover gauge
prusa_cover{printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1"} 1
# HELP prusa_cpu_temp Status of the printer cpu temp
# TYPE prusa_cpu_temp gauge
prusa_cpu_temp{printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1"} 54
# HELP prusa_cpu_usage_ratio CPU usage from 0.0 to 1.0
# TYPE prusa_cpu_usage_ratio gauge
prusa_cpu_usage_ratio{ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 0.32
prusa_cpu_usage_ratio{ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 0.15
prusa_cpu_usage_ratio{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 0.37
prusa_cpu_usage_ratio{ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0.22
# HELP prusa_current Current of different devices in / on the printer in miliampers
# TYPE prusa_current gauge
prusa_current{device="bed_0",ip="192.168.20.130",mac="10:9c:70:25:17:3d",rail=""} 25
prusa_current{device="bed_1",ip="192.168.20.130",mac="10:9c:70:25:17:3d",rail=""} 27
prusa_current{device="inp",ip="192.168.20.139",mac="10:9c:70:2d:9c:fe",rail=""} -4202.509
prusa_current{device="inp",ip="192.168.20.173",mac="10:9c:70:2c:da:8",rail=""} 9951.395
prusa_current{device="inp",ip="192.168.20.174",mac="10:9c:70:2c:37:82",rail=""} 6967.796
prusa_current{device="inp",ip="192.168.20.176",mac="10:9c:70:2d:17:2e",rail=""} -309.27500000000003
prusa_current{device="mmu",ip="192.168.20.139",mac="10:9c:70:2d:9c:fe",rail=""} -13
prusa_current{device="mmu",ip="192.168.20.173",mac="10:9c:70:2c:da:8",rail=""} -3.726
prusa_current{device="mmu",ip="192.168.20.174",mac="10:9c:70:2c:37:82",rail=""} -8.363000000000001
prusa_current{device="mmu",ip="192.168.20.176",mac="10:9c:70:2d:17:2e",rail=""} -13
prusa_current{device="nozz",ip="192.168.20.139",mac="10:9c:70:2d:9c:fe",rail=""} 1192.32
prusa_current{device="nozz",ip="192.168.20.173",mac="10:9c:70:2c:da:8",rail=""} 16.793
prusa_current{device="nozz",ip="192.168.20.174",mac="10:9c:70:2c:37:82",rail=""} 1088.202
prusa_current{device="nozz",ip="192.168.20.176",mac="10:9c:70:2d:17:2e",rail=""} 0
prusa_current{device="sandwich",ip="192.168.20.130",mac="10:9c:70:25:17:3d",rail="5V"} 598.3739999999999
prusa_current{device="splitter",ip="192.168.20.130",mac="10:9c:70:25:17:3d",rail="5V"} 238.15900000000002
prusa_current{device="xlBuddy",ip="192.168.20.130",mac="10:9c:70:25:17:3d",rail="5V"} 494.179
# HELP prusa_current_raw Current of different devices in / on the printer in raw sensor value
# TYPE prusa_current_raw gauge
prusa_current_raw{device="inp",ip="192.168.20.174",mac="10:9c:70:2c:37:82",rail=""} 434
prusa_current_raw{device="inp",ip="192.168.20.176",mac="10:9c:70:2d:17:2e",rail=""} 503
prusa_current_raw{device="nozz",ip="192.168.20.174",mac="10:9c:70:2c:37:82",rail=""} 250
prusa_current_raw{device="nozz",ip="192.168.20.176",mac="10:9c:70:2d:17:2e",rail=""} 1
# HELP prusa_fan_active Fan active
# TYPE prusa_fan_active gauge
prusa_fan_active{fan="heatbreak",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 8026
prusa_fan_active{fan="heatbreak",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
prusa_fan_active{fan="print",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 3493
prusa_fan_active{fan="print",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
# HELP prusa_fan_speed Returns information about speed of hotend fan in rpm.
# TYPE prusa_fan_speed gauge
prusa_fan_speed{fan="blower",printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1"} 0
prusa_fan_speed{fan="hotend",printer_address="192.168.20.130",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl"} 3648
prusa_fan_speed{fan="hotend",printer_address="192.168.20.139",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39"} 7751
prusa_fan_speed{fan="hotend",printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3"} 0
prusa_fan_speed{fan="hotend",printer_address="192.168.20.173",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef"} 0
prusa_fan_speed{fan="print",printer_address="192.168.20.130",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl"} 0
prusa_fan_speed{fan="print",printer_address="192.168.20.139",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39"} 2285
prusa_fan_speed{fan="print",printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3"} 0
prusa_fan_speed{fan="print",printer_address="192.168.20.173",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef"} 0
prusa_fan_speed{fan="rear",printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1"} 0
prusa_fan_speed{fan="uv",printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1"} 0
# HELP prusa_fan_speed_ratio Fan
# TYPE prusa_fan_speed_ratio gauge
prusa_fan_speed_ratio{fan="heatbreak",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 1
prusa_fan_speed_ratio{fan="heatbreak",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
prusa_fan_speed_ratio{fan="print",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 0.34901960784313724
prusa_fan_speed_ratio{fan="print",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
# HELP prusa_farm_mode Return if printer is set to farm mode
# TYPE prusa_farm_mode gauge
prusa_farm_mode{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3"} 0
# HELP prusa_filament Name of printed (b)gcode
# TYPE prusa_filament gauge
prusa_filament{filament="---",ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 1
# HELP prusa_files Number of files in storage
# TYPE prusa_files gauge
prusa_files{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3",printer_storage="PrusaLink gcodes"} 0
prusa_files{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3",printer_storage="SD Card"} 80
# HELP prusa_fsensor_raw Filament Sensor - raw sensor value
# TYPE prusa_fsensor_raw gauge
prusa_fsensor_raw{ip="192.168.20.174",mac="10:9c:70:2c:37:82",sensor="0"} 1.308758e+06
prusa_fsensor_raw{ip="192.168.20.176",mac="10:9c:70:2d:17:2e",sensor="0"} 29630
# HELP prusa_gui_loop_duration Gui loop duration
# TYPE prusa_gui_loop_duration gauge
prusa_gui_loop_duration{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 10
prusa_gui_loop_duration{ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 10
# HELP prusa_heap_free Free heap
# TYPE prusa_heap_free gauge
prusa_heap_free{ip="192.168.20.130",mac="10:9c:70:25:17:3d"} 49032
prusa_heap_free{ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 60788
prusa_heap_free{ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 63968
prusa_heap_free{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 58860
prusa_heap_free{ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 67588
# HELP prusa_heap_total Total heap
# TYPE prusa_heap_total gauge
prusa_heap_total{ip="192.168.20.130",mac="10:9c:70:25:17:3d"} 77048
prusa_heap_total{ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 92496
prusa_heap_total{ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 89636
prusa_heap_total{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 89444
prusa_heap_total{ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 89636
# HELP prusa_heater_enabled Heater enabled
# TYPE prusa_heater_enabled gauge
prusa_heater_enabled{ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 1
prusa_heater_enabled{ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 0
prusa_heater_enabled{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 1
prusa_heater_enabled{ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
# HELP prusa_info Returns information about printer.
# TYPE prusa_info gauge
prusa_info{api_version="0.9.0-legacy",printer_address="192.168.20.157",printer_hostname="connect.prusa3d.com",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_location="Groove Street",printer_model="I3MK3S",printer_name="mk3smmu3",prusalink_name="MK3S with MMU3",serial_number="CZPX5222X004XK04220",server_version="0.7.2",version_text="PrusaLink 0.7.2"} 1
prusa_info{api_version="2.0.0",printer_address="192.168.20.130",printer_hostname="PrusaXL",printer_job_name="",printer_job_path="",printer_location="",printer_model="XL",printer_name="xl",prusalink_name="",serial_number="10101-07117692302090107",server_version="2.1.2",version_text="PrusaLink"} 1
prusa_info{api_version="2.0.0",printer_address="192.168.20.139",printer_hostname="PrusaMK4",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_location="",printer_model="MK4",printer_name="mk39",prusalink_name="",serial_number="10589-3742441631728135",server_version="2.1.2",version_text="PrusaLink"} 1
prusa_info{api_version="2.0.0",printer_address="192.168.20.173",printer_hostname="PrusaMK4",printer_job_name="",printer_job_path="",printer_location="",printer_model="MK4",printer_name="mk39q3wfef",prusalink_name="",serial_number="4914-27145608112152703",server_version="2.1.2",version_text="PrusaLink"} 1
# HELP prusa_loadcell Value from loadcell sensor
# TYPE prusa_loadcell gauge
prusa_loadcell{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 8175.053223
prusa_loadcell{ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 4835.84668
# HELP prusa_loadcell_age Loadcell age
# TYPE prusa_loadcell_age gauge
prusa_loadcell_age{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} -3142
prusa_loadcell_age{ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} -3143
# HELP prusa_loadcell_hp Loadcell filtered z load
# TYPE prusa_loadcell_hp gauge
prusa_loadcell_hp{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 0
prusa_loadcell_hp{ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
# HELP prusa_loadcell_hysteresis Loadcell hysteresis
# TYPE prusa_loadcell_hysteresis gauge
prusa_loadcell_hysteresis{ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 80
prusa_loadcell_hysteresis{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 80
# HELP prusa_loadcell_raw Value from loadcell sensor in raw sensor value
# TYPE prusa_loadcell_raw gauge
prusa_loadcell_raw{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} -531110
prusa_loadcell_raw{ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 251867
# HELP prusa_loadcell_scale Loadcell scale
# TYPE prusa_loadcell_scale gauge
prusa_loadcell_scale{ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 0.0192
prusa_loadcell_scale{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 0.0192
# HELP prusa_loadcell_threshold Loadcell threshold
# TYPE prusa_loadcell_threshold gauge
prusa_loadcell_threshold{ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} -125
prusa_loadcell_threshold{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} -125
# HELP prusa_loadcell_threshold_cont Loadcell threshold continuous
# TYPE prusa_loadcell_threshold_cont gauge
prusa_loadcell_threshold_cont{ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} -40
prusa_loadcell_threshold_cont{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} -40
# HELP prusa_loadcell_xy Loadcell XY
# TYPE prusa_loadcell_xy gauge
prusa_loadcell_xy{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 0
prusa_loadcell_xy{ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
# HELP prusa_material Returns information about loaded filament. Returns 0 if there is no loaded filament
# TYPE prusa_material gauge
prusa_material{printer_address="192.168.20.130",printer_filament="PETG",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl"} 1
prusa_material{printer_address="192.168.20.139",printer_filament="PETG",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39"} 1
prusa_material{printer_address="192.168.20.157",printer_filament=" - ",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3"} 0
prusa_material{printer_address="192.168.20.173",printer_filament="---",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef"} 0
# HELP prusa_media_prefetched_bytes Media prefetched in bytes
# TYPE prusa_media_prefetched_bytes gauge
prusa_media_prefetched_bytes{ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 7447
prusa_media_prefetched_bytes{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 6938
# HELP prusa_mmu Returns information if MMU is enabled.
# TYPE prusa_mmu gauge
prusa_mmu{printer_address="192.168.20.130",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl"} 0
prusa_mmu{printer_address="192.168.20.139",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39"} 0
prusa_mmu{printer_address="192.168.20.173",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef"} 0
# HELP prusa_network_in Network in
# TYPE prusa_network_in gauge
prusa_network_in{device="esp",ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 13181
prusa_network_in{device="eth",ip="192.168.20.130",mac="10:9c:70:25:17:3d"} 515681
prusa_network_in{device="eth",ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 9.930676e+06
prusa_network_in{device="eth",ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 394556
prusa_network_in{device="eth",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 1.12784159e+08
prusa_network_in{device="eth",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 2.36532321e+08
# HELP prusa_network_out Network out
# TYPE prusa_network_out gauge
prusa_network_out{device="esp",ip="192.168.20.130",mac="10:9c:70:25:17:3d"} 8370
prusa_network_out{device="esp",ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 4302
prusa_network_out{device="eth",ip="192.168.20.130",mac="10:9c:70:25:17:3d"} 2.698047e+06
prusa_network_out{device="eth",ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 2.9208e+07
prusa_network_out{device="eth",ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 1.153696e+06
prusa_network_out{device="eth",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 2.523787802e+09
# HELP prusa_nozzle_size Returns information about selected nozzle size.
# TYPE prusa_nozzle_size gauge
prusa_nozzle_size{printer_address="192.168.20.130",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl"} 0.4
prusa_nozzle_size{printer_address="192.168.20.139",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39"} 0.4
prusa_nozzle_size{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3"} 0.4
prusa_nozzle_size{printer_address="192.168.20.173",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef"} 0.4
# HELP prusa_overcurrent Overcurrent of different devices in / on the printer
# TYPE prusa_overcurrent gauge
prusa_overcurrent{device="inp",ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 0
prusa_overcurrent{device="inp",ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 0
prusa_overcurrent{device="inp",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 0
prusa_overcurrent{device="inp",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
prusa_overcurrent{device="nozz",ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 0
prusa_overcurrent{device="nozz",ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 0
prusa_overcurrent{device="nozz",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 0
prusa_overcurrent{device="nozz",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
# HELP prusa_points_dropped Points dropped
# TYPE prusa_points_dropped gauge
prusa_points_dropped{ip="192.168.20.130",mac="10:9c:70:25:17:3d"} 0
prusa_points_dropped{ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 27
prusa_points_dropped{ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 27
prusa_points_dropped{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 4647
prusa_points_dropped{ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 773
# HELP prusa_print_flow_ratio Returns information about of filament flow in ratio (0.0 - 1.0).
# TYPE prusa_print_flow_ratio gauge
prusa_print_flow_ratio{printer_address="192.168.20.130",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl"} 1
prusa_print_flow_ratio{printer_address="192.168.20.139",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39"} 1
prusa_print_flow_ratio{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3"} 1
prusa_print_flow_ratio{printer_address="192.168.20.173",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef"} 1
# HELP prusa_print_speed_ratio Current setting of printer speed in values from 0.0 - 1.0
# TYPE prusa_print_speed_ratio gauge
prusa_print_speed_ratio{printer_address="192.168.20.130",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl"} 1
prusa_print_speed_ratio{printer_address="192.168.20.139",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39"} 1
prusa_print_speed_ratio{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3"} 1
prusa_print_speed_ratio{printer_address="192.168.20.173",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef"} 1
# HELP prusa_print_time Returns information about current print time.
# TYPE prusa_print_time gauge
prusa_print_time{printer_address="192.168.20.130",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl"} 0
prusa_print_time{printer_address="192.168.20.139",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39"} 3972
prusa_print_time{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3"} 0
prusa_print_time{printer_address="192.168.20.173",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef"} 0
# HELP prusa_printing Printing printer
# TYPE prusa_printing gauge
prusa_printing{ip="192.168.20.130",mac="10:9c:70:25:17:3d"} 0
prusa_printing{ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 1
prusa_printing{ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 0
prusa_printing{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 1
prusa_printing{ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
# HELP prusa_printing_progress Returns information about completion of current print in percents
# TYPE prusa_printing_progress gauge
prusa_printing_progress{printer_address="192.168.20.130",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl"} 0
prusa_printing_progress{printer_address="192.168.20.139",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39"} 0.09
prusa_printing_progress{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3"} 0
prusa_printing_progress{printer_address="192.168.20.173",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef"} 0
# HELP prusa_printing_time_remaining Returns time that remains for completion of current print
# TYPE prusa_printing_time_remaining gauge
prusa_printing_time_remaining{printer_address="192.168.20.130",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl"} 0
prusa_printing_time_remaining{printer_address="192.168.20.139",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39"} 33480
prusa_printing_time_remaining{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3"} 0
prusa_printing_time_remaining{printer_address="192.168.20.173",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef"} 0
# HELP prusa_pwm PWM value of nozzle and bed mostly
# TYPE prusa_pwm gauge
prusa_pwm{device="bed",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 138
prusa_pwm{device="bed",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
prusa_pwm{device="nozzle",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 108
prusa_pwm{device="nozzle",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
# HELP prusa_status Returns information status of printer.
# TYPE prusa_status gauge
prusa_status{printer_address="192.168.20.130",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl",printer_state="Operational"} 1
prusa_status{printer_address="192.168.20.139",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39",printer_state="Printing"} 4
prusa_status{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3",printer_state="Operational"} 1
prusa_status{printer_address="192.168.20.173",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef",printer_state="Operational"} 1
prusa_status{printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1",printer_state="Ready"} 1
# HELP prusa_stepper_ipos Stepper possition from startup
# TYPE prusa_stepper_ipos gauge
prusa_stepper_ipos{axis="x",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} -19634
prusa_stepper_ipos{axis="x",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
prusa_stepper_ipos{axis="y",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} -31025
prusa_stepper_ipos{axis="y",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
prusa_stepper_ipos{axis="z",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 8271
prusa_stepper_ipos{axis="z",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
# HELP prusa_stepper_pos Stepper possition
# TYPE prusa_stepper_pos gauge
prusa_stepper_pos{axis="x",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 47.860001
prusa_stepper_pos{axis="x",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} -1
prusa_stepper_pos{axis="y",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 51.209999
prusa_stepper_pos{axis="y",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} -4
prusa_stepper_pos{axis="z",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 5.105
prusa_stepper_pos{axis="z",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
# HELP prusa_temp Temperature of different devices in / on the printer
# TYPE prusa_temp gauge
prusa_temp{device="bed",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 89.993752
prusa_temp{device="bed",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 22.261906
prusa_temp{device="brd",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 73.092949
prusa_temp{device="brd",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 35.14706
prusa_temp{device="dwarf_board",ip="192.168.20.130",mac="10:9c:70:25:17:3d"} 40
prusa_temp{device="hbr_0",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 48.28
prusa_temp{device="hbr_0",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 23.52
prusa_temp{device="mcu",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 74
prusa_temp{device="mcu",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 41
prusa_temp{device="noz_0",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 234.92
prusa_temp{device="noz_0",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 24.81
# HELP prusa_temp_target Target temperature of different devices in / on the printer
# TYPE prusa_temp_target gauge
prusa_temp_target{device="bed",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 90
prusa_temp_target{device="bed",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
prusa_temp_target{device="noz_0",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 235
prusa_temp_target{device="noz_0",ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 0
# HELP prusa_tmc_sg Trinamic SG
# TYPE prusa_tmc_sg gauge
prusa_tmc_sg{axis="x",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 0
prusa_tmc_sg{axis="y",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 71
prusa_tmc_sg{axis="z",ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 0
# HELP prusa_tool_temp Status of the printer tool temp
# TYPE prusa_tool_temp gauge
prusa_tool_temp{printer_address="192.168.20.130",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl",tool="0"} 32
prusa_tool_temp{printer_address="192.168.20.139",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39",tool="0"} 250.3
prusa_tool_temp{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3",tool="0"} 22.4
prusa_tool_temp{printer_address="192.168.20.173",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef",tool="0"} 25
prusa_tool_temp{printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1",tool="0"} 24.5
# HELP prusa_tool_temp_offset Offset tool temp
# TYPE prusa_tool_temp_offset gauge
prusa_tool_temp_offset{printer_address="192.168.20.130",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl",tool="0"} 0
prusa_tool_temp_offset{printer_address="192.168.20.139",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39",tool="0"} 0
prusa_tool_temp_offset{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3",tool="0"} 0
prusa_tool_temp_offset{printer_address="192.168.20.173",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef",tool="0"} 0
prusa_tool_temp_offset{printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1",tool="0"} 0
# HELP prusa_tool_temp_target Target tool temp
# TYPE prusa_tool_temp_target gauge
prusa_tool_temp_target{printer_address="192.168.20.130",printer_job_name="",printer_job_path="",printer_model="XL",printer_name="xl",tool="0"} 0
prusa_tool_temp_target{printer_address="192.168.20.139",printer_job_name="ruzovka_x-carriage-back(1)_0.4n_0.2mm_PETG_MK4_10h19m.bgcode",printer_job_path="/usb/RUZOVK~2.BGC",printer_model="MK4",printer_name="mk39",tool="0"} 250
prusa_tool_temp_target{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="I3MK3S",printer_name="mk3smmu3",tool="0"} 0
prusa_tool_temp_target{printer_address="192.168.20.173",printer_job_name="",printer_job_path="",printer_model="MK4",printer_name="mk39q3wfef",tool="0"} 0
prusa_tool_temp_target{printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1",tool="0"} 0
# HELP prusa_up Return information about online printers. If printer is registered as offline then returned value is 0.
# TYPE prusa_up gauge
prusa_up{printer_address="192.168.20.130",printer_model="XL",printer_name="xl"} 1
prusa_up{printer_address="192.168.20.139",printer_model="MK4",printer_name="mk39"} 1
prusa_up{printer_address="192.168.20.157",printer_model="I3MK3S",printer_name="mk3smmu3"} 1
prusa_up{printer_address="192.168.20.173",printer_model="MK4",printer_name="mk39q3wfef"} 1
prusa_up{printer_address="192.168.20.31",printer_model="SL1",printer_name="sl1"} 1
# HELP prusa_up_syslog Printer up - from syslog metric - ttl is by default 60 seconds but can be different and it depends on choosen interval. That means if printer wont sent any data for 60 seconds is considered down.
# TYPE prusa_up_syslog gauge
prusa_up_syslog{ip="192.168.20.130",mac="10:9c:70:25:17:3d"} 1
prusa_up_syslog{ip="192.168.20.139",mac="10:9c:70:2d:9c:fe"} 1
prusa_up_syslog{ip="192.168.20.173",mac="10:9c:70:2c:da:8"} 1
prusa_up_syslog{ip="192.168.20.174",mac="10:9c:70:2c:37:82"} 1
prusa_up_syslog{ip="192.168.20.176",mac="10:9c:70:2d:17:2e"} 1
# HELP prusa_uv_temp Status of the printer uv temp
# TYPE prusa_uv_temp gauge
prusa_uv_temp{printer_address="192.168.20.31",printer_job_name="",printer_job_path="",printer_model="SL1",printer_name="sl1"} 24.5
# HELP prusa_voltage Voltage of different devices in / on the printer
# TYPE prusa_voltage gauge
prusa_voltage{device="",ip="192.168.20.130",mac="10:9c:70:25:17:3d",rail="24V"} 24.242474
prusa_voltage{device="",ip="192.168.20.130",mac="10:9c:70:25:17:3d",rail="5V"} 5.043011
prusa_voltage{device="bed",ip="192.168.20.139",mac="10:9c:70:2d:9c:fe",rail=""} 23.954302
prusa_voltage{device="bed",ip="192.168.20.173",mac="10:9c:70:2c:da:8",rail=""} 24.134409
prusa_voltage{device="bed",ip="192.168.20.174",mac="10:9c:70:2c:37:82",rail=""} 24.098385
prusa_voltage{device="bed",ip="192.168.20.176",mac="10:9c:70:2d:17:2e",rail=""} 24.062366
prusa_voltage{device="nozz",ip="192.168.20.139",mac="10:9c:70:2d:9c:fe",rail=""} 23.954302
prusa_voltage{device="nozz",ip="192.168.20.173",mac="10:9c:70:2c:da:8",rail=""} 0.036022
prusa_voltage{device="nozz",ip="192.168.20.174",mac="10:9c:70:2c:37:82",rail=""} 23.846235
prusa_voltage{device="nozz",ip="192.168.20.176",mac="10:9c:70:2d:17:2e",rail=""} 0
# HELP prusa_voltage_raw Voltage of different devices in / on the printer in raw sensor value
# TYPE prusa_voltage_raw gauge
prusa_voltage_raw{device="bed",ip="192.168.20.174",mac="10:9c:70:2c:37:82",rail=""} 669
prusa_voltage_raw{device="bed",ip="192.168.20.176",mac="10:9c:70:2d:17:2e",rail=""} 668
prusa_voltage_raw{device="nozz",ip="192.168.20.174",mac="10:9c:70:2c:37:82",rail=""} 662
prusa_voltage_raw{device="nozz",ip="192.168.20.176",mac="10:9c:70:2d:17:2e",rail=""} 1