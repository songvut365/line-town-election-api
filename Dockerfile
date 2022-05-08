FROM golang:1.17-alpine

RUN apk add git
RUN apk add build-base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /out/dist

EXPOSE 8080

CMD [ "/out/dist" ]