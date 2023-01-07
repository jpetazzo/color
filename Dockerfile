FROM golang:alpine
ENV CGO_ENABLED=0
COPY webcolor.go /go
RUN go build webcolor.go
FROM scratch
COPY --from=0 /go/webcolor /webcolor
CMD ["/webcolor"]
EXPOSE 80
