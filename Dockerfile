FROM golang:1.17 AS golang
FROM node:16

WORKDIR /usr/src/app

COPY --from=golang /usr/bin/go /usr/bin/go
COPY . .

RUN yarn

ENTRYPOINT [ "yarn" ]