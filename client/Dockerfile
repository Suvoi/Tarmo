# Stage 1: Build
FROM node:24.11.0-alpine AS build

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

# Build and export static HTML
RUN npm run build

# Stage 2: Serve with Nginx
FROM nginx:alpine

# Copy exported static files (Next.js outputs to /app/out by default)
COPY --from=build /app/out /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
