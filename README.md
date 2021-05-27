#API service

Run:
```bash
go run main.go
```
Get latest block
```bash
curl -X GET "http://localhost:8080/api/v1/eth/latest-block"
```

Get recent N block
```bash
curl -X GET "http://localhost:8080/api/v1/eth/get-recent-block?limit=5"
```

Get block by id
```bash
curl -X GET "http://localhost:8080/api/v1/eth/get-block-by-number?blockNum=9170682"
```

Get transaction by tx_hash and save to db
```bash
curl -X GET "http://localhost:8080/api/v1/eth/get-tx?hash=0x241d8239aaf78d4d78d1057b0e9f03d60bbcaada0d1f530fc3fbc1b828d5b177"
```