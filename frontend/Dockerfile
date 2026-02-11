# Build stage
FROM oven/bun:1-alpine AS build

WORKDIR /app

COPY package.json bun.lockb ./
RUN bun install --frozen-lockfile

COPY . .
RUN bunx ng build --configuration production

# Serve stage
FROM nginx:alpine

COPY --from=build /app/dist/rzeczy-znalezione/browser /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
