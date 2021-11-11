FROM golang:alpine
COPY webcolor.go /go
RUN go build webcolor.go
FROM alpine
COPY --from=0 /go/webcolor /webcolor
CMD ["/webcolor"]
EXPOSE 80
