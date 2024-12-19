FROM node:22-alpine3.20

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn install --frozen-lockfile && yarn cache clean --force

COPY . .

RUN yarn build

ENV HOST=0.0.0.0

ENV NODE_ENV=production

EXPOSE 3000

CMD ["yarn", "start"]