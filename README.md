# GO Signal

Simple POC to provide a signal (start/stop) to continuous process controlled by endpoint requests

## Run
Execute on terminal
```sh
go run cmd/api/main.go
```

## Endpoints
### Start
```curlrc
curl -X POST \
  'http://localhost:8000/start'
```

### Stop
```curlrc
curl -X POST \
  'http://localhost:8000/stop'
```