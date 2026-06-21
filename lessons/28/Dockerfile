FROM golang:1.25
WORKDIR /src
COPY . /src
RUN go build -o /bin/myapp ./main.go
CMD [ "/bin/myapp" ]