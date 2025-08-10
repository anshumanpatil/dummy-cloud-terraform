FROM golang:1.24.5

WORKDIR /go/src/app
COPY . .
RUN go mod tidy
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]


# docker build . -t go-app
# docker run -p 8090:8090 go-app
# cd .. && make install && cd example
# terraform init && terraform apply --auto-approve
