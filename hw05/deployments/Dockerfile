FROM golang:1.16 as builder
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/app

FROM scratch
COPY --from=builder /app/app /bin/app
EXPOSE 8000
CMD ["/bin/app"]