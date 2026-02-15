# Примеры использования API

## Image

### Create Image (save)

request:
```bash
curl -X POST http://localhost:5555/api/v2/image/create 
  -F "file=@./zoshua-colah-gGrGRDEhtP4-unsplash.jpg"
```

response:
```bash
{"url_slug":"ChHBQuDSmPkW"}
```

---

### Get Image

request:
```bash
curl http://localhost:5555/api/v2/image/ChHBQuDSmPkW -o image_test_api_6.png
```

response: ```image_test_api_6.png```


___

## Text

### Create Text (save)

request:
```bash
curl -X POST http://localhost:5555/api/v2/text/create -H "Content-Type: application/json" -d 'hello'
```

response:
```bash
{"url_slug":"NoGDQT6E"}
```

---

### Get Text

request:
```bash
curl http://localhost:5555/api/v2/text/aBUK9WnjHu1T
```

response: 
```{"id":1,"url_slug":"aBUK9WnjHu1T","created_at":"2026-02-13T23:54:30.859137Z","content":"hello"}```