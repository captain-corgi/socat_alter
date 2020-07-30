# Pull and build go app
FROM golang:alpine as builder
RUN mkdir /src
RUN cd /src/
COPY * /src/
RUN ls /src
WORKDIR /src
RUN go build

# Copy and run binary file in light-weight image
FROM alpine
RUN mkdir /app
WORKDIR /app
COPY --from=builder /src/socat_alter /app/
COPY --from=builder /src/routes.json /app/
CMD ["./socat_alter"]