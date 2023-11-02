#!/bin/bash

apt update
apt install git wget tar

# https://gist.github.com/ookangzheng/8c0c7eba6d7b12d5da6a8bdeec1da9b2

rm -rf /usr/local/go
GOVERSION="1.20.3"
wget "https://golang.org/dl/go${GOVERSION}.linux-arm64.tar.gz" -4
tar -C /usr/local -xvf "go${GOVERSION}.linux-arm64.tar.gz"
rm "go${GOVERSION}.linux-arm64.tar.gz"

cat >> ~/.bashrc << 'EOF'
export GOPATH=$HOME/go
export PATH=/usr/local/go/bin:$PATH:$GOPATH/bin
EOF

echo "Go installed"

source ~/.bashrc

rm -rf /etc/prusa_exporter
cd /etc
git clone https://github.com/pstrobl96/prusa_exporter.git

cd /etc/prusa_exporter
go build

touch /etc/systemd/system/prusa.service
rm /etc/prusa_exporter/prusa.yaml
touch /etc/prusa_exporter/prusa.yaml

cat <<EOT >> /etc/prusa_exporter/prusa.yml
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
EOT

cat <<EOT >> /etc/systemd/system/prusa.service
[Unit]
Description=Prusa exporter service
After=network.target
StartLimitIntervalSec=0

[Service]
Type=simple
Restart=always
RestartSec=1
User=root
Environment=PRUSA_EXPORTER_CONFIG=/etc/prusa_exporter/prusa.yaml
Environment=PRUSA_EXPORTER_PORT=10009
ExecStart=/etc/prusa_exporter/prusa_exporter

[Install]
WantedBy=multi-user.target
EOT

systemctl daemon-reload
systemctl enable prusa
systemctl start prusa

echo "Done"
