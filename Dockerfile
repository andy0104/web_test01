#this is the build stage
FROM golang:1.23.4 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .

# the run stage
FROM scratch
WORKDIR /app
# copy ssl certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY --from=builder /app/api .
EXPOSE 8081
CMD ["./api"]