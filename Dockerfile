FROM alpine:edge AS build
RUN apk add --no-cache ca-certificates && \
    update-ca-certificates
RUN apk update && \
        apk upgrade --available && \
        apk add --no-cache \
        openssl \
        curl
RUN apk add --update go gcc g++
WORKDIR /demowebapp
COPY go.mod go.sum *.go ./
RUN CGO_ENABLED=1 go build -o ./webapp

FROM alpine:edge
WORKDIR /demowebapp
COPY assets/ ./assets/
COPY db/ ./db/
COPY *.json ./
COPY --from=build /demowebapp/webapp ./
ENV PORT 8080
EXPOSE ${PORT}
#CMD tail -f /dev/null
CMD ["./webapp"]

# If building behind a proxy
# docker build --tag demowebapp \
#    --build-arg http_proxy=$http_proxy \
#    --build-arg https_proxy=$https_proxy \
#    .

# Mount host volume to ./db path for persitence
# docker run -v "$(pwd)/db":/demowebapp/db \
#    -d -p 8080:8080 \
#    --name=demowebapp \
#    demowebapp

# Use env-vars to change default Port etc
# deocker run -v "$(pwd)":/demowebapp/db \
#    -e PORT=8081 \
#    -d -p 8082:8081 \
#    --name=demowebapp \
#    demowebapp