# todo-app

## curl

```sh
curl -X POST \
"http://localhost:1323/todos" \
-H 'Content-Type: application/json' \
-d '{
    "title": "title2",
    "body": "body2"
}'
```

```sh
curl -X GET \
"http://localhost:1323/todos"
```
