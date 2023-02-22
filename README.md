### Step1 run mongodb
```sh
docker compose up
```
### Step2 Register LINE Bot
https://manager.line.biz/account/@714kxhuc/setting


### Step3 run ngrok
```sh
./ngrok http 8080
```

### Step4 run linerobot
```sh
go run main.go
```

```sh
go build -ldflags "-s -w" -o main *.go

./main
```

![image](./vedio/demo.mov)
