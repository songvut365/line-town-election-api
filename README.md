# LINE Town Election API

## Why am I used this

- `Golang and Fiber` Good structure, Good performance.
- `SQLite3` Lightweight and No installation just a one file.
- `GORM` Easy CRUD, Auto migration and Connecting to database.

## Setup

### .env file

- `PORT=8080` port for run server.
- `DATABASE=./database/election.db` path of database file.
- `EASY_TOKEN=xxxx` token for authorization bearer.

```cmd
go mod tidy
go run main.go
```

## More API Details

- `Create Candidate` all input are required.
- `Update Candidate` all input are required, except voted count.
- `Delete Candidate` delete candidate will also delete votes.
- `Check Vote Status` national id is required, string type.
- `Vote` all input are required.
- `Toggle Election` enable is required, boolean type.
- `Get Election Export Result (download)` file may not update if downloaded in the browser.

## Documents

- [Fiber](http://gofiber.io/)
- [Sqlite3](https://www.sqlite.org/index.html)
- [Gorm](https://gorm.io/)
