# LINE Town Election API

## Why am I used this

- `Golang and Fiber` Good structure, Good performance
- `SQLite3` Lightweight, No installation just a one file and Easy for run on local
- `GORM` Easy CRUD, Auto migration and Connecting to database

## Setup

### .env file

- `PORT=8080` port for run server
- `EASY_TOKEN=xxxx` token for authorization bearer
- `BEARER=Bearer` bearer for authorization bearer
- `DATABASE=./database/election.db` path of database file
- `CSV_FILE=./public/export/result.csv` path of csv file for read
- `CSV_FILE_SEND=/public/export/result.csv` path of csv file for send

## How to run

### Local

```
$ go mod download
$ go run main.go
```

### Docker

```
$ docker build -t election-api .
$ docker run -d -p 8080:8080 --name election-api-01 --env-file .\.env election-api
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

## Documents

- [Fiber](http://gofiber.io/)
- [Sqlite3](https://www.sqlite.org/index.html)
- [Gorm](https://gorm.io/)
- [Testify](https://github.com/stretchr/testify)
- [Chart.js](https://www.chartjs.org/)
