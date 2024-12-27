# megatude

REST API for earth quake in indonesia base on bmkg data

# Tech Stack

- Golang 1.23
- Echo v4 (backend service)
- GORM (migration)
- PostgreSQL (dbms)
- Cron (schedule tasker from golang) `https://github.com/robfig/cron`

# Instalation

1. Make config.yaml file for app configuration
2. install docker and start docker service
3. run docker with command `docker compose watch`
