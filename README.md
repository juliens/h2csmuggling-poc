# How to

`docker-compose up` to launch Traefik, nginx and an upgradable to h2c backend

`go run ./` to launch the client test


```
2020/09/12 16:49:11 Try Traefik HTTP
2020/09/12 16:49:11 Try to upgrade
2020/09/12 16:49:11 Switched success
2020/09/12 16:49:11 Try normal request
2020/09/12 16:49:11 Try to send another request in the upgraded connection
2020/09/12 16:49:11 same status code 401 and 401
2020/09/12 16:49:11 Seems OK (secure)
2020/09/12 16:49:11 ---------------
2020/09/12 16:49:11 Try nginx HTTP
2020/09/12 16:49:11 Try to upgrade
2020/09/12 16:49:11 Switched success
2020/09/12 16:49:11 Try normal request
2020/09/12 16:49:11 Try to send another request in the upgraded connection
2020/09/12 16:49:11 different status code 200 and 403
You got the flag!
2020/09/12 16:49:11 ---------------
2020/09/12 16:49:11 Try Traefik HTTPS
2020/09/12 16:49:11 Try to upgrade
2020/09/12 16:49:11 upgrade failed: Get "https://127.0.0.1:8443": http2: invalid Upgrade request header: ["h2c"]
2020/09/12 16:49:11 ---------------
2020/09/12 16:49:11 Try nginx HTTPS
2020/09/12 16:49:11 Try to upgrade
2020/09/12 16:49:11 Switched success
2020/09/12 16:49:11 Try normal request
2020/09/12 16:49:11 Try to send another request in the upgraded connection
2020/09/12 16:49:11 different status code 200 and 403
You got the flag!
```
