### curlコマンド
- curl localhost:8000/snippets
- curl localhost:8000/users
- curl -X POST -H "Content-Type: application/json" -d '{"code": "print(123)"}' localhost:8000/snippets/
- curl -X POST -H "Content-Type: application/json" -d '{"code": "print(123)"}' -u s12100040:password localhost:8000/snippets/
