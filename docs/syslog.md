# Enable SYSLOG in printer

Another issue with SYSLOG is configuration. You need munually enable sending metrics in printers GUI - step by step bellow. And you need to run specific gcode file, that specifies SYSLOG server. 

With M330 and M334 g-code you can configure your SYSLOG server and there are two ways how to get needed g-code. I've created ![config.gcode](examples/syslog/config.gcode) containing everything needed. **You just need to change IP address and you can change port. however while testing I had issues with numerous text editors and printer was very picky.** I was successful with `nano` and basic `echo` command in terminal. However you can as well activate additional metrics that can be found here ![config_full.gcode](examples/syslog/config_full.gcode). If you enable additional metrics then you need to run your configuration gcode after rebooting the printer. 

Second way is to add these two lines in to start of g-code with PrusaSlicer. Don't forget to change IP address and port! `10008` is default that I used but use whatever you want but you need also change that value in configuration.
```
M330 SYSLOG
M334 192.168.20.2 10008
```

For logs you can use `M340`. 

```
M340 192.168.20.2 10007
```

Or you can combine both and you'll get

```
M330 SYSLOG
M334 192.168.20.2 10008
M340 192.168.20.2 10007
```

After loading gcode on to flash drive you can enable the metrics in printer.

Open `Settings`  
![syslog0](readme/syslog/screenshot_0.jpg)

Navigate to `Network`  
![syslog1](readme/syslog/screenshot_1.jpg)

Find `Metrics & Log`  
![syslog2](readme/syslog/screenshot_2.jpg)

Now click on `Allow` - Confirm and change value to `Any Host`. Next switch on `Enabled Stored at Startup`  
![syslog3](readme/syslog/screenshot_3.jpg)
![syslog4](readme/syslog/screenshot_4.jpg)
![syslog5](readme/syslog/screenshot_5.jpg)

Now run your configuration gcode we created before
![syslog6](readme/syslog/screenshot_6.jpg)
![syslog7](readme/syslog/screenshot_7.jpg)

Navigate back to `Metrics & Log` and find `Current Configuration` - click on `Metrics Host` and store it as Host   
![syslog8](readme/syslog/screenshot_8.jpg)
![syslog9](readme/syslog/screenshot_9.jpg)

Click on `Metrics Port` and store it as Metrics Port  
![syslog10](readme/syslog/screenshot_10.jpg)
![syslog11](readme/syslog/screenshot_11.jpg)

After configuration it should look like this. Only IP address should be different. And if different port was choosen then also port.  
![syslog12](readme/syslog/screenshot_12.jpg)

