FROM golang:alpine AS build-container
ADD . /work
WORKDIR /work
ENV GO111MODULE=on
RUN apk --update add --no-cache git mercurial
RUN go build -o=migrate ./db/migrate.go
RUN go build

FROM alpine:3.9
COPY --from=build-container /work/PersonalDataRepository /usr/local/bin/PersonalDataRepository
COPY --from=build-container /work/migrate /usr/local/bin/migrate
CMD ["usr/local/bin/PersonalDataRepository"]
