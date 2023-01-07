FROM golang:1.18

WORKDIR /go/src/app
COPY . .

RUN go build -o main main.go
EXPOSE 8090
CMD ["./main"]


# docker build . -t go-app
# docker run -p 8090:8090 go-app
# cd .. && make install && cd example
# terraform init && terraform apply --auto-approve
