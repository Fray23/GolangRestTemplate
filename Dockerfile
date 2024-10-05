FROM golang:1.23

WORKDIR /code
COPY src/ /code
#RUN go build -o hello hello.go
# CMD [". /hello"]
