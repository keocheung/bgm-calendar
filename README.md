# bgm-calendar

### A custom iCalendar generator for bgm.tv

## Install
### Docker Compose
```yaml
version: '3'
services:
  bgm-calendar:
    image: ghcr.io/keocheung/bgm-calendar
    container_name: bgm-calendar
    volumes:
      - /etc/localtime:/etc/localtime:ro
    network_mode: bridge
    ports:
      - 8080:8080
    environments:
      - BGM_CALENDAR_PORT=8080            # Optional
      - BGM_CALENDAR_API_HOST=api.bgm.tv  # Optional
    restart: unless-stopped
```

## Usage
```shell
curl http://localhost:8080/users/{username}/games.ics
```