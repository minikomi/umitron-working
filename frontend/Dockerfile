FROM node:16-alpine3.14

WORKDIR /frontend
COPY package.json yarn.lock ./

RUN yarn install

CMD ["yarn", "start"]
