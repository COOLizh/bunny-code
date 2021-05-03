FROM node:12.22.1-alpine3.12 as builder
WORKDIR /app
COPY package.json .
COPY default.conf .
RUN npm install
COPY . .
RUN npm run build

FROM nginx:1.18-alpine
RUN rm /etc/nginx/conf.d/default.conf
COPY --from=builder /app/default.conf /etc/nginx/conf.d/
COPY --from=builder /app/build/ /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]
