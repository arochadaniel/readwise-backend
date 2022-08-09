FROM golang

COPY . /app

WORKDIR /app

RUN go get ./...
RUN go install github.com/cosmtrek/air@latest

EXPOSE 4000
CMD ["air"]
