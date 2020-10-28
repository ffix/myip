FROM golang:latest as build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myip .

FROM scratch
COPY --from=build /app/myip /myip
ENTRYPOINT ["/myip"]
EXPOSE 8000
