# Start by building the application.
FROM golang:1.20-alpine as builder

RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
#RUN apt update && apt install -y ca-certificates openssl
#ARG cert_location=/etc/ssl/certs
# Get certificate from "github.com"
#RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null| openssl x509 -outform PEM > ${cert_location}/github.crt
# Get certificate from "proxy.golang.org"
#RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null| openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
# Update certificates
#RUN update-ca-certificates

WORKDIR /go/src/app
COPY go.mod go.sum main.go .

RUN go mod download && go mod verify
#RUN go mod tidy
#RUN go vet -v
#RUN go test -v
RUN CGO_ENABLED=0 go build -o /go/bin/app

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian11
COPY --from=builder /go/bin/app /
CMD ["/app"]
