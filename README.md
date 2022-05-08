# LINE Town Election API

## Why am I used this

- `Golang and Fiber` Good structure, Good performance
- `SQLite3` Lightweight, No installation just a one file and Easy for run on local
- `MySQL` Flexiblem, Popular and Good for database server
- `GORM` Easy CRUD, Auto migration, Connecting and Switching to database

## Setup

### .env file

- `PORT=8080` port for run server
- `EASY_TOKEN=xxxx` token for authorization bearer
- `BEARER=Bearer` bearer for authorization bearer
- `DATABASE_TYPE=SQLite` type of database
- `SQLITE_URI=./database/election.db` path of sqlite file
- `MYSQL_URI=root:1234@tcp(localhost:3306)/line-town-election?charset=utf8mb4&parseTime=True&loc=Local` uri of mysql server
- `CSV_FILE=./public/export/result.csv` path of csv file for read
- `CSV_FILE_SEND=./public/export/result.csv` path of csv file for send

## How to run

### Option 1 - Run with SQLite

```
$ go mod download
$ go run main.go
```

### Option 2 - Run by docker-compose with MySQL

```
$ docker-compose up -d
```

### Example with cURL: (Get Candidates)

```
$ curl --location --request GET 'http://localhost:8080/api/candidates' \
    --header 'Authorization: Bearer xxxx'
```


## API Details

### Candidate

- `GET /api/candidate` Get Candidates - can query by name and limit
- `GET /api/candidate/:candidateId` Get a Candidate
- `POST /api/candidate` Create Candidate - all input are required
- `PUT /api/candidate/:candidateId` Update Candidate - all input are required, except voted count
- `DELETE /api/candidate/:candidateId` Delete Candidate - votes also delete

### Vote

- `POST /api/vote/status` Check Vote Status - national id is required, string type
- `POST /api/vote` Vote - all input are required

### Election

- `POST /api/election/toggle` Toggle Election - enable is required, boolean type
- `GET /api/election/count` Get Election Count
- `GET /api/election/result` Get Election Result
- `GET /api/election/export` Get Election Export Result - file may not update if downloaded in the browser

## Web Socket

- `/ws/candidate/` Stream Vote All Candidates
- `/ws/candidate/:candidateId` Stream Vote Candidate By Id

## Add-on

- `/chart` Show Chart Top 10 Candidate
- `/api/candidates?name=Songvut Nakrong&limit=5` Candidate query by name and limit
- `/ws/candidate` Stream all candidates vote