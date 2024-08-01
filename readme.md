for mac/linux:
ps aux | grep learning

bot.service:
```
[Unit]
Description=bot

[Service]
WorkingDirectory=/root/bot
ExecStart=/root/bot/bot
Restart=always
User=root
KillMode=control-group
KillSignal=SIGINT

StandardOutput=append:/root/bot/output.log
StandardError=append:/root/bot/error.log

[Install]
WantedBy=multi-user.target
```

start.sh:
```sh
#!/bin/bash
sudo systemctl start bot.service
sudo systemctl status bot.service
```

stop.sh:
```sh
#!/bin/bash
sudo systemctl stop bot.service
```