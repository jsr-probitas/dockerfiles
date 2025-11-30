FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder
ARG TARGETOS TARGETARCH
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o echo-http .

FROM scratch
COPY --from=builder /app/echo-http /echo-http
EXPOSE 80
ENTRYPOINT ["/echo-http"]
