FROM --platform=linux/amd64 golang:1.18.9-alpine3.17 as builder

WORKDIR /build

COPY . ./

RUN go mod download && go build -o godevops cmd/main.go

WORKDIR /dist

RUN cp /build/godevops .

FROM scratch

COPY --from=builder /dist/godevops .

ENTRYPOINT ["./godevops"]