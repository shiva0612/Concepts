FROM golang:1.19 as builder
ENV GOOS linux
ENV CGO_ENABLED 0
RUN mkdir appbin
WORKDIR /app
# COPY <src> <des>
COPY go.mod go.sum ./ 
RUN go mod download
COPY . .
RUN go build -o /appbin/server .
WORKDIR /


FROM alpine:3.14 as production
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /appbin .
EXPOSE 8080
CMD ./server

