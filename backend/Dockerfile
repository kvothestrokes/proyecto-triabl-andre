FROM golang:1.19-alpine AS build
WORKDIR /go/src/proyecto
COPY . .

RUN go build -o server .

FROM alpine:3.12
EXPOSE 8000
COPY --from=build /go/src/proyecto/server /server
CMD ["/server"]
