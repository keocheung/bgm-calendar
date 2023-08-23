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
      - BGM_CALENDAR_PORT=8080               # Optional. Default is 8080
      - BANGUMI_API_HOST=api.bgm.tv          # Optional. Default is api.bgm.tv
      - BANGUMI_ACCESS_TOKEN=asd123          # Optional. Default is empty. Create here: https://next.bgm.tv/demo/access-token
      - HTTP_PROXY=http://127.0.0.1:9090     # Optional. Default is empty
      - HTTPS_PROXY=http://127.0.0.1:9090    # Optional. Default is empty
      - NO_PROXY=example.com,192.168.0.0/16  # Optional. Default is empty
    restart: unless-stopped
```

## Usage
### 生成「想玩的游戏」的 ics 日历：
```shell
curl http://localhost:8080/users/{username}/games.ics
```