FROM golang:1.20

WORKDIR /watcher
COPY . /watcher

WORKDIR /watcher/cmd/api
RUN go build -o /bin/watcher -ldflags '-w -s' -tags netgo -a -installsuffix cgo -v .

EXPOSE 8080
CMD ["/bin/watcher"]