# Url Shortner

Initial steps for learning go by building a simple url shortner. The database used is Redis for quick implementation

## Endpoints

- Create Url POST /create

```
curl --location 'http://localhost:3000/create' \
--header 'Content-Type: application/json' \
--data '{
    "url": "http://www.google.com",
    "user_id": "1234"
}'
```

- Get Url: GET /:shortUrl
