FROM golang:1.19 AS builder

# PHASE 1: BUILD

# set working directory for the container
WORKDIR /build

# copy go project files into wd
COPY go.mod go.sum ./
# download all dependencies
RUN go mod download

# copy the entirety of the project into the working directory (technically unideal)
COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o /api cmd/main.go

# PHASE 2: TEST (unimplemented)

# PHASE 3: DEPLOY

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /api /api

EXPOSE 1323

CMD ["./api"]
