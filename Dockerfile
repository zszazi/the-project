# Multi Stage docker file 

# Stage 1

FROM golang:1.22 as base

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN env GOOS=linux GOARCH=amd64 go build -o app .

# Stage 2 - create a distro-less image - reduced size 

FROM gcr.io/distroless/base

COPY --from=base /app/app .

COPY --from=base /app/static ./static

EXPOSE 8080

CMD ["./app"]