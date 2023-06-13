FROM golang:1.20

WORKDIR /app

COPY ./LetiSleti/LSbackend/go.mod ./LetiSleti/LSbackend/go.sum ./
RUN go mod download

COPY ./LetiSleti/LSbackend/ ./

RUN go build -o /LetiSletiApp

EXPOSE 8080

CMD [ "/LetiSletiApp" ]
