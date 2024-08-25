FROM golang:1.23

WORKDIR /code
COPY . /code
#RUN go build -o hello hello.go
# CMD [". /hello"]
