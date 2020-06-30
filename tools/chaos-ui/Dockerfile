FROM node:12.16.1-buster-slim AS builder
LABEL stage=builder
RUN mkdir -p /app
WORKDIR /app
COPY package.json /app
RUN yarn install
COPY . /app
RUN yarn build

FROM nginx:alpine
RUN mkdir -p /var/app
WORKDIR /var/app
COPY --from=builder /app/build .
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80