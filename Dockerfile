FROM node:16

WORKDIR /usr/src/app
COPY . .

RUN yarn

ENTRYPOINT [ "yarn" ]