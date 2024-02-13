```
# HELP prusa_buddy_active_extruder Active extruder - used for XL
# TYPE prusa_buddy_active_extruder gauge
prusa_buddy_active_extruder{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 1
# HELP prusa_buddy_axis_x Returns information about position of axis X.
# TYPE prusa_buddy_axis_x gauge
prusa_buddy_axis_x{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 0
prusa_buddy_axis_x{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} -1
prusa_buddy_axis_x{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 0
prusa_buddy_axis_x{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 180.4
prusa_buddy_axis_x{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} -1
prusa_buddy_axis_x{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} -1
# HELP prusa_buddy_axis_y Returns information about position of axis Y.
# TYPE prusa_buddy_axis_y gauge
prusa_buddy_axis_y{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 0
prusa_buddy_axis_y{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} -4
prusa_buddy_axis_y{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 0
prusa_buddy_axis_y{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} -3
prusa_buddy_axis_y{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} -4
prusa_buddy_axis_y{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} -4
# HELP prusa_buddy_axis_z Returns information about position of axis Z.
# TYPE prusa_buddy_axis_z gauge
prusa_buddy_axis_z{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 19.7
prusa_buddy_axis_z{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0
prusa_buddy_axis_z{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 5
prusa_buddy_axis_z{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 0
prusa_buddy_axis_z{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0
prusa_buddy_axis_z{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_buddy_axis_z_adjustment Axis Z adjustment
# TYPE prusa_buddy_axis_z_adjustment gauge
prusa_buddy_axis_z_adjustment{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 0
prusa_buddy_axis_z_adjustment{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0
prusa_buddy_axis_z_adjustment{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 0
prusa_buddy_axis_z_adjustment{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} -1.365007996559143
prusa_buddy_axis_z_adjustment{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0
prusa_buddy_axis_z_adjustment{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_buddy_bed_target_temperature Target temperature of printer bed in Celsius
# TYPE prusa_buddy_bed_target_temperature gauge
prusa_buddy_bed_target_temperature{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 60
prusa_buddy_bed_target_temperature{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0
prusa_buddy_bed_target_temperature{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 60
prusa_buddy_bed_target_temperature{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 0
prusa_buddy_bed_target_temperature{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0
prusa_buddy_bed_target_temperature{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_buddy_bed_temperature Current temperature of printer bed in Celsius
# TYPE prusa_buddy_bed_temperature gauge
prusa_buddy_bed_temperature{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 60
prusa_buddy_bed_temperature{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 23.3
prusa_buddy_bed_temperature{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 60
prusa_buddy_bed_temperature{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 22.5
prusa_buddy_bed_temperature{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 23.1
prusa_buddy_bed_temperature{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 22.5
# HELP prusa_buddy_cpu_usage CPU usage
# TYPE prusa_buddy_cpu_usage gauge
prusa_buddy_cpu_usage{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 56
prusa_buddy_cpu_usage{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 16
prusa_buddy_cpu_usage{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 21
prusa_buddy_cpu_usage{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 10
prusa_buddy_cpu_usage{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 15
prusa_buddy_cpu_usage{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 13
# HELP prusa_buddy_current_bed Current of bed
# TYPE prusa_buddy_current_bed gauge
prusa_buddy_current_bed{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest",rail="0"} 1.8769999742507935
prusa_buddy_current_bed{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest",rail="1"} 0.27399998903274536
# HELP prusa_buddy_current_input Input current
# TYPE prusa_buddy_current_input gauge
prusa_buddy_current_input{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 9.041762351989746
prusa_buddy_current_input{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} -3.1473329067230225
prusa_buddy_current_input{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 9.73308277130127
prusa_buddy_current_input{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} -0.6731299757957458
# HELP prusa_buddy_current_mmu Current of MMU
# TYPE prusa_buddy_current_mmu gauge
prusa_buddy_current_mmu{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} -0.013000000268220901
prusa_buddy_current_mmu{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} -0.008363000117242336
prusa_buddy_current_mmu{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} -0.003725999966263771
prusa_buddy_current_mmu{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} -0.013000000268220901
# HELP prusa_buddy_current_nozzle Current of nozzle
# TYPE prusa_buddy_current_nozzle gauge
prusa_buddy_current_nozzle{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0.020152000710368156
prusa_buddy_current_nozzle{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 1.5449780225753784
prusa_buddy_current_nozzle{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0.01679299958050251
prusa_buddy_current_nozzle{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_buddy_current_xlbuddy_5volts Current of xlBuddy 5V rail
# TYPE prusa_buddy_current_xlbuddy_5volts gauge
prusa_buddy_current_xlbuddy_5volts{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 0.4912019968032837
# HELP prusa_buddy_dwarf_board_temp Dwarf board temperature - used for XL
# TYPE prusa_buddy_dwarf_board_temp gauge
prusa_buddy_dwarf_board_temp{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 35
# HELP prusa_buddy_dwarf_mcu_temp Dwarf MCU temperature - used for XL
# TYPE prusa_buddy_dwarf_mcu_temp gauge
prusa_buddy_dwarf_mcu_temp{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 36
# HELP prusa_buddy_fan_hotend Returns information about speed of hotend fan in rpm.
# TYPE prusa_buddy_fan_hotend gauge
prusa_buddy_fan_hotend{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 7976
prusa_buddy_fan_hotend{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0
prusa_buddy_fan_hotend{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 7681
prusa_buddy_fan_hotend{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 0
prusa_buddy_fan_hotend{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0
prusa_buddy_fan_hotend{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_buddy_fan_print Returns information about speed of print fan in rpm.
# TYPE prusa_buddy_fan_print gauge
prusa_buddy_fan_print{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 1597
prusa_buddy_fan_print{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0
prusa_buddy_fan_print{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 6261
prusa_buddy_fan_print{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 0
prusa_buddy_fan_print{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0
prusa_buddy_fan_print{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_buddy_files Number of files in storage
# TYPE prusa_buddy_files gauge
prusa_buddy_files{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest",printer_storage="USB"} 23
prusa_buddy_files{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft",printer_storage="USB"} 4
prusa_buddy_files{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing",printer_storage="USB"} 57
prusa_buddy_files{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest",printer_storage="USB"} 0
prusa_buddy_files{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright",printer_storage="USB"} 0
prusa_buddy_files{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025",printer_storage="USB"} 7
# HELP prusa_buddy_heap_free Free heap
# TYPE prusa_buddy_heap_free gauge
prusa_buddy_heap_free{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 42432
prusa_buddy_heap_free{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 66628
prusa_buddy_heap_free{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 56024
prusa_buddy_heap_free{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 24560
prusa_buddy_heap_free{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 62832
prusa_buddy_heap_free{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 67236
# HELP prusa_buddy_heap_total Total heap
# TYPE prusa_buddy_heap_total gauge
prusa_buddy_heap_total{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 76536
prusa_buddy_heap_total{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 89636
prusa_buddy_heap_total{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 89636
prusa_buddy_heap_total{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 42332
prusa_buddy_heap_total{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 89636
prusa_buddy_heap_total{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 89636
# HELP prusa_buddy_heater_enabled Heater enabled
# TYPE prusa_buddy_heater_enabled gauge
prusa_buddy_heater_enabled{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0
prusa_buddy_heater_enabled{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 1
prusa_buddy_heater_enabled{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0
prusa_buddy_heater_enabled{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_buddy_info Returns information about printer.
# TYPE prusa_buddy_info gauge
prusa_buddy_info{printer_address="192.168.20.130",printer_hostname="PrusaXL",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest",printer_serial="10101-07117692302090107"} 1
prusa_buddy_info{printer_address="192.168.20.131",printer_hostname="PrusaMK4",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft",printer_serial="4914-27145608112160886"} 1
prusa_buddy_info{printer_address="192.168.20.139",printer_hostname="PrusaMK4",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing",printer_serial="10589-3742441631728135"} 1
prusa_buddy_info{printer_address="192.168.20.192",printer_hostname="PrusaMINI",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest",printer_serial="CZPX2421X017XC51551"} 1
prusa_buddy_info{printer_address="192.168.20.209",printer_hostname="PrusaMK4",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright",printer_serial="4914-27145608112152703"} 1
prusa_buddy_info{printer_address="192.168.20.211",printer_hostname="PrusaMK4",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025",printer_serial="10589-3742441631701959"} 1
# HELP prusa_buddy_loadcell_hysteresis Loadcell hysteresis
# TYPE prusa_buddy_loadcell_hysteresis gauge
prusa_buddy_loadcell_hysteresis{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 80
prusa_buddy_loadcell_hysteresis{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 80
# HELP prusa_buddy_loadcell_scale Loadcell scale
# TYPE prusa_buddy_loadcell_scale gauge
prusa_buddy_loadcell_scale{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 0.019200000911951065
prusa_buddy_loadcell_scale{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 0.019200000911951065
# HELP prusa_buddy_loadcell_threshold Loadcell threshold
# TYPE prusa_buddy_loadcell_threshold gauge
prusa_buddy_loadcell_threshold{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} -125
prusa_buddy_loadcell_threshold{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} -125
# HELP prusa_buddy_material Returns information about loaded filament. Returns 0 if there is no loaded filament
# TYPE prusa_buddy_material gauge
prusa_buddy_material{printer_address="192.168.20.130",printer_filament="PLA",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 1
prusa_buddy_material{printer_address="192.168.20.131",printer_filament="PETG",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 1
prusa_buddy_material{printer_address="192.168.20.139",printer_filament="PLA",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 1
prusa_buddy_material{printer_address="192.168.20.192",printer_filament="PLA",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 1
prusa_buddy_material{printer_address="192.168.20.209",printer_filament="PETG",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 1
prusa_buddy_material{printer_address="192.168.20.211",printer_filament="PETG",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 1
# HELP prusa_buddy_media_prefetched Media prefetched
# TYPE prusa_buddy_media_prefetched gauge
prusa_buddy_media_prefetched{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 8397
prusa_buddy_media_prefetched{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 5827
# HELP prusa_buddy_mmu Returns information if MMU is enabled.
# TYPE prusa_buddy_mmu gauge
prusa_buddy_mmu{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 0
prusa_buddy_mmu{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0
prusa_buddy_mmu{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 0
prusa_buddy_mmu{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 0
prusa_buddy_mmu{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0
prusa_buddy_mmu{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_buddy_nozzle_size Returns information about selected nozzle size.
# TYPE prusa_buddy_nozzle_size gauge
prusa_buddy_nozzle_size{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 0.4
prusa_buddy_nozzle_size{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0.4
prusa_buddy_nozzle_size{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 0.4
prusa_buddy_nozzle_size{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 0.6
prusa_buddy_nozzle_size{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0.4
prusa_buddy_nozzle_size{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0.25
# HELP prusa_buddy_nozzle_target_temperature Target temperature of printer nozzle in Celsius
# TYPE prusa_buddy_nozzle_target_temperature gauge
prusa_buddy_nozzle_target_temperature{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 225
prusa_buddy_nozzle_target_temperature{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0
prusa_buddy_nozzle_target_temperature{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 225
prusa_buddy_nozzle_target_temperature{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 0
prusa_buddy_nozzle_target_temperature{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0
prusa_buddy_nozzle_target_temperature{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_buddy_nozzle_temperature Current temperature of printer nozzle in Celsius
# TYPE prusa_buddy_nozzle_temperature gauge
prusa_buddy_nozzle_temperature{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 225
prusa_buddy_nozzle_temperature{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 25
prusa_buddy_nozzle_temperature{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 224.6
prusa_buddy_nozzle_temperature{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 21
prusa_buddy_nozzle_temperature{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 25
prusa_buddy_nozzle_temperature{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 25
# HELP prusa_buddy_overcurrent_input Overcurrent of input
# TYPE prusa_buddy_overcurrent_input gauge
prusa_buddy_overcurrent_input{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0
prusa_buddy_overcurrent_input{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 0
prusa_buddy_overcurrent_input{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0
prusa_buddy_overcurrent_input{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_buddy_overcurrent_nozzle Overcurrent of nozzle
# TYPE prusa_buddy_overcurrent_nozzle gauge
prusa_buddy_overcurrent_nozzle{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0
prusa_buddy_overcurrent_nozzle{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 0
prusa_buddy_overcurrent_nozzle{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0
prusa_buddy_overcurrent_nozzle{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_buddy_points_dropped Points dropped
# TYPE prusa_buddy_points_dropped gauge
prusa_buddy_points_dropped{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 0
prusa_buddy_points_dropped{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 27
prusa_buddy_points_dropped{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 38
prusa_buddy_points_dropped{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 0
prusa_buddy_points_dropped{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 27
prusa_buddy_points_dropped{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 25
# HELP prusa_buddy_print_flow_ratio Returns information about of filament flow in ratio (0.0 - 1.0).
# TYPE prusa_buddy_print_flow_ratio gauge
prusa_buddy_print_flow_ratio{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 1
prusa_buddy_print_flow_ratio{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 1
prusa_buddy_print_flow_ratio{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 1
prusa_buddy_print_flow_ratio{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 1
prusa_buddy_print_flow_ratio{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 1
prusa_buddy_print_flow_ratio{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 1
# HELP prusa_buddy_print_speed_ratio Current setting of printer speed in ratio (0.0-1.0)
# TYPE prusa_buddy_print_speed_ratio gauge
prusa_buddy_print_speed_ratio{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 1
prusa_buddy_print_speed_ratio{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 1
prusa_buddy_print_speed_ratio{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 1
prusa_buddy_print_speed_ratio{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 1
prusa_buddy_print_speed_ratio{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 1
prusa_buddy_print_speed_ratio{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 1
# HELP prusa_buddy_print_time Returns information about current print time.
# TYPE prusa_buddy_print_time gauge
prusa_buddy_print_time{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 44
prusa_buddy_print_time{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0
prusa_buddy_print_time{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 11768
prusa_buddy_print_time{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 0
prusa_buddy_print_time{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0
prusa_buddy_print_time{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_buddy_printing Return information about printing
# TYPE prusa_buddy_printing gauge
prusa_buddy_printing{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 1
prusa_buddy_printing{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0
prusa_buddy_printing{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 1
prusa_buddy_printing{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 0
prusa_buddy_printing{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0
prusa_buddy_printing{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_buddy_printing_progress Returns information about completion of current print in percents
# TYPE prusa_buddy_printing_progress gauge
prusa_buddy_printing_progress{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 0.6
prusa_buddy_printing_progress{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0
prusa_buddy_printing_progress{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 0.94
prusa_buddy_printing_progress{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 0
prusa_buddy_printing_progress{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0
prusa_buddy_printing_progress{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_buddy_printing_time_remaining Returns time that remains for completion of current print
# TYPE prusa_buddy_printing_time_remaining gauge
prusa_buddy_printing_time_remaining{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 7980
prusa_buddy_printing_time_remaining{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0
prusa_buddy_printing_time_remaining{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 720
prusa_buddy_printing_time_remaining{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 0
prusa_buddy_printing_time_remaining{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0
prusa_buddy_printing_time_remaining{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_buddy_status Returns information status of printer.
# TYPE prusa_buddy_status gauge
prusa_buddy_status{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest",printer_state="PRINTING"} 4
prusa_buddy_status{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft",printer_state="IDLE"} 1
prusa_buddy_status{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing",printer_state="PRINTING"} 4
prusa_buddy_status{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest",printer_state="IDLE"} 1
prusa_buddy_status{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright",printer_state="IDLE"} 1
prusa_buddy_status{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025",printer_state="IDLE"} 1
# HELP prusa_buddy_syslog_info Buddy info
# TYPE prusa_buddy_syslog_info gauge
prusa_buddy_syslog_info{buddy_bom="0",buddy_revision="105",printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 1
prusa_buddy_syslog_info{buddy_bom="14",buddy_revision="27",printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 1
prusa_buddy_syslog_info{buddy_bom="14",buddy_revision="27",printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 1
prusa_buddy_syslog_info{buddy_bom="38",buddy_revision="37",printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 1
prusa_buddy_syslog_info{buddy_bom="40",buddy_revision="37",printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 1
prusa_buddy_syslog_info{buddy_bom="7",buddy_revision="7",printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 1
# HELP prusa_buddy_up Return information about online printers. If printer is registered as offline then returned value is 0.
# TYPE prusa_buddy_up gauge
prusa_buddy_up{printer_address="192.168.20.130",printer_model="PrusaXL",printer_name="xltest"} 1
prusa_buddy_up{printer_address="192.168.20.131",printer_model="PrusaMK4",printer_name="encleft"} 1
prusa_buddy_up{printer_address="192.168.20.139",printer_model="PrusaMK4",printer_name="testing"} 1
prusa_buddy_up{printer_address="192.168.20.192",printer_model="PrusaMINI",printer_name="minitest"} 1
prusa_buddy_up{printer_address="192.168.20.209",printer_model="PrusaMK4",printer_name="encright"} 1
prusa_buddy_up{printer_address="192.168.20.211",printer_model="PrusaMK4",printer_name="mk4025"} 1
# HELP prusa_buddy_version Return information about printer. This metric contains information mostly about Prusa Link
# TYPE prusa_buddy_version gauge
prusa_buddy_version{printer_address="192.168.20.130",printer_api="2.0.0",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest",printer_server="2.1.2",printer_text="PrusaLink"} 1
prusa_buddy_version{printer_address="192.168.20.131",printer_api="2.0.0",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft",printer_server="2.1.2",printer_text="PrusaLink"} 1
prusa_buddy_version{printer_address="192.168.20.139",printer_api="2.0.0",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing",printer_server="2.1.2",printer_text="PrusaLink"} 1
prusa_buddy_version{printer_address="192.168.20.192",printer_api="2.0.0",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest",printer_server="2.1.2",printer_text="PrusaLink"} 1
prusa_buddy_version{printer_address="192.168.20.209",printer_api="2.0.0",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright",printer_server="2.1.2",printer_text="PrusaLink"} 1
prusa_buddy_version{printer_address="192.168.20.211",printer_api="2.0.0",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025",printer_server="2.1.2",printer_text="PrusaLink"} 1
# HELP prusa_buddy_voltage_24volts Voltage of 24V rail
# TYPE prusa_buddy_voltage_24volts gauge
prusa_buddy_voltage_24volts{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 24.026344299316406
# HELP prusa_buddy_voltage_5volts Voltage of 5V rail
# TYPE prusa_buddy_voltage_5volts gauge
prusa_buddy_voltage_5volts{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 5.04301118850708
# HELP prusa_buddy_voltage_bed Voltage of bed
# TYPE prusa_buddy_voltage_bed gauge
prusa_buddy_voltage_bed{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 24.350536346435547
prusa_buddy_voltage_bed{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 23.954301834106445
prusa_buddy_voltage_bed{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 24.062366485595703
prusa_buddy_voltage_bed{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 24.026344299316406
# HELP prusa_buddy_voltage_nozzle Voltage of nozzle
# TYPE prusa_buddy_voltage_nozzle gauge
prusa_buddy_voltage_nozzle{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0
prusa_buddy_voltage_nozzle{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 23.846235275268555
prusa_buddy_voltage_nozzle{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0.03602200001478195
prusa_buddy_voltage_nozzle{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_buddy_voltage_sandwich_5volts Voltage of sandwich 5V rail
# TYPE prusa_buddy_voltage_sandwich_5volts gauge
prusa_buddy_voltage_sandwich_5volts{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 1.0032429695129395
# HELP prusa_buddy_voltage_splitter_5volts Voltage of splitter 5V rail
# TYPE prusa_buddy_voltage_splitter_5volts gauge
prusa_buddy_voltage_splitter_5volts{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 0.43166300654411316
# HELP prusa_buddy_z_height Current height of Z
# TYPE prusa_buddy_z_height gauge
prusa_buddy_z_height{printer_address="192.168.20.130",printer_job_name="multiple_grots_0.4n_0.15mm_PLA,PLA,PLA,PLA_XLIS_5h36m.bgcode",printer_job_path="/usb/MULTIP~1.BGC",printer_model="PrusaXL",printer_name="xltest"} 19.7
prusa_buddy_z_height{printer_address="192.168.20.131",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encleft"} 0
prusa_buddy_z_height{printer_address="192.168.20.139",printer_job_name="PromCoin-bezel_1mm_0.4n_0.2mm_PLA_MK4IS_3h25m.bgcode",printer_job_path="/usb/PROMCO~1.BGC",printer_model="PrusaMK4",printer_name="testing"} 5
prusa_buddy_z_height{printer_address="192.168.20.192",printer_job_name="",printer_job_path="",printer_model="PrusaMINI",printer_name="minitest"} 0
prusa_buddy_z_height{printer_address="192.168.20.209",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="encright"} 0
prusa_buddy_z_height{printer_address="192.168.20.211",printer_job_name="",printer_job_path="",printer_model="PrusaMK4",printer_name="mk4025"} 0
# HELP prusa_einsy_axis_x Return coordinates - x axis of printer
# TYPE prusa_einsy_axis_x gauge
prusa_einsy_axis_x{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 50
# HELP prusa_einsy_axis_y Return coordinates - y axis of printer
# TYPE prusa_einsy_axis_y gauge
prusa_einsy_axis_y{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 190
# HELP prusa_einsy_axis_z Return coordinates - z axis of printer
# TYPE prusa_einsy_axis_z gauge
prusa_einsy_axis_z{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 10.2
# HELP prusa_einsy_bed_target_temperature Target temperature of printer bed in Celsius
# TYPE prusa_einsy_bed_target_temperature gauge
prusa_einsy_bed_target_temperature{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 0
# HELP prusa_einsy_bed_temperature Current temperature of printer bed in Celsius
# TYPE prusa_einsy_bed_temperature gauge
prusa_einsy_bed_temperature{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 22.5
# HELP prusa_einsy_farm_mode Return if printer is set to farm mode
# TYPE prusa_einsy_farm_mode gauge
prusa_einsy_farm_mode{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 0
# HELP prusa_einsy_files Number of files in storage
# TYPE prusa_einsy_files gauge
prusa_einsy_files{printer_address="192.168.20.157",printer_model="PrusaLink I3MK3S",printer_name="donkey",printer_storage="PrusaLink gcodes"} 0
prusa_einsy_files{printer_address="192.168.20.157",printer_model="PrusaLink I3MK3S",printer_name="donkey",printer_storage="SD Card"} 80
# HELP prusa_einsy_info Return info about printer
# TYPE prusa_einsy_info gauge
prusa_einsy_info{printer_address="192.168.20.157",printer_api="0.9.0-legacy",printer_hostname="connect.prusa3d.com",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_link_name="MK3S with MMU3",printer_location="Elf on a shelf",printer_model="PrusaLink I3MK3S",printer_name="donkey",printer_server="0.7.2",printer_sn="CZPX5222X004XK04220",printer_text="PrusaLink 0.7.2",printer_type="Prusa I3MK3S - FW: 3.13.1-6876"} 1
# HELP prusa_einsy_material Returns information about loaded filament. Returns 0 if there is no loaded filament
# TYPE prusa_einsy_material gauge
prusa_einsy_material{printer_address="192.168.20.157",printer_filament=" - ",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 0
# HELP prusa_einsy_nozzle_size Return size of nozzle
# TYPE prusa_einsy_nozzle_size gauge
prusa_einsy_nozzle_size{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 0.4
# HELP prusa_einsy_nozzle_target_temperature Target temperature of printer nozzle in Celsius
# TYPE prusa_einsy_nozzle_target_temperature gauge
prusa_einsy_nozzle_target_temperature{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 0
# HELP prusa_einsy_nozzle_temperature Current temperature of printer nozzle in Celsius
# TYPE prusa_einsy_nozzle_temperature gauge
prusa_einsy_nozzle_temperature{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 21.7
# HELP prusa_einsy_print_speed_ratio Current setting of printer speed in values from 0.0 - 1.0
# TYPE prusa_einsy_print_speed_ratio gauge
prusa_einsy_print_speed_ratio{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 0.8
# HELP prusa_einsy_print_time Returns actual printing time of current print
# TYPE prusa_einsy_print_time gauge
prusa_einsy_print_time{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 0
# HELP prusa_einsy_printing DEPRECATED - Return information about printing
# TYPE prusa_einsy_printing gauge
prusa_einsy_printing{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 0
# HELP prusa_einsy_printing_progress Returns information about completion of current print in percents
# TYPE prusa_einsy_printing_progress gauge
prusa_einsy_printing_progress{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 0
# HELP prusa_einsy_printing_time_remaining Returns time that remains for completion of current print
# TYPE prusa_einsy_printing_time_remaining gauge
prusa_einsy_printing_time_remaining{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 0
# HELP prusa_einsy_state Return state of printer
# TYPE prusa_einsy_state gauge
prusa_einsy_state{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey",printer_state="Operational"} 0
# HELP prusa_einsy_up Return if printer is up
# TYPE prusa_einsy_up gauge
prusa_einsy_up{printer_address="192.168.20.157",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 1
# HELP prusa_einsy_version DEPRECATED - Return information about printer. This metric contains information mostly about Prusa Link
# TYPE prusa_einsy_version gauge
prusa_einsy_version{printer_address="192.168.20.157",printer_api="0.9.0-legacy",printer_model="PrusaLink I3MK3S",printer_name="donkey",printer_server="0.7.2",printer_text="PrusaLink 0.7.2"} 1
# HELP prusa_einsy_z_height DEPRECATED - Current height of Z
# TYPE prusa_einsy_z_height gauge
prusa_einsy_z_height{printer_address="192.168.20.157",printer_job_name="fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_job_path="/SD Card/fosdem_0.2mm_PLA,PLA_MK3SMMU3_7h16m.gcode",printer_model="PrusaLink I3MK3S",printer_name="donkey"} 10.2
```