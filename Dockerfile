FROM golang:alpine as builders
RUN go env -w GOPROXY=https://goproxy.cn,direct
COPY . /app/
WORKDIR /app/
RUN go build -o apps src/main.go
EXPOSE 9088


FROM alpine
COPY --from=builders /app/apps /ohmyfans/apps
ENTRYPOINT ["/ohmyfans/apps"]
CMD []


