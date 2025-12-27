# build stage
FROM node:18-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# serve stage
FROM nginx:alpine
<<<<<<< Updated upstream
<<<<<<< Updated upstream
=======
COPY nginx.conf /etc/nginx/conf.d/default.conf
>>>>>>> Stashed changes
=======
COPY nginx.conf /etc/nginx/conf.d/default.conf
>>>>>>> Stashed changes
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 80
