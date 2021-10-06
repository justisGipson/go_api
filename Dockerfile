FROM golang:1.17-alpine AS builder

# set necessary env variables needed & build server
ENV CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64

# move to working directory
WORKDIR /build

# copy and download dependency w/ go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# copy code into container
COPY . .

RUN go build -ldflags="-s -w" -o main .

FROM scratch

# copy binary and config files from /build
COPY --from=builder ["/build/main", "/build/.env", "/"]

# expose necessary port
EXPOSE 5000

# command to run when starting container
ENTRYPOINT ["/main"]

