FROM golang:lastest as builders
RUN mkdir /run
WORKDIR /run
COPY ../* /app
RUN go build -o apps /app/src/main.go
EXPOSE 9088


FROM alpine:3.13.5
COPY --from=builders /run/apps /run/apps
ENTRYPOINT [ "/run/apps" ]
CMD[]


