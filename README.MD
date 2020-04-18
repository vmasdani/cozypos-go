# CozyPOS GO
Run this program:
- Install dependencies with `go get`
- Prepare database
- Run with `go build && ./cozypos-go`
- Generate secret with `http://localhost:8080/generate`
- Make .env file in root directory containing:
```
SECRET=<your_generated_secret>
DB_NAME=<your_db_name>
DB_USERNAME=<your_db_username>
DB_PASSWORD=<your_db_password>
```