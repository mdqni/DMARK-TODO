FROM node:20 AS frontend-builder
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ .
RUN npm run build

FROM golang:1.21 AS backend-builder
WORKDIR /app/backend
COPY backend/go.mod backend/go.sum ./
RUN go mod tidy
COPY backend/ .
COPY --from=frontend-builder /app/frontend/dist ./frontend/dist
RUN go build -o app ./cmd/main.go

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=backend-builder /app/backend/app ./
COPY --from=frontend-builder /app/frontend/dist ./frontend/dist
EXPOSE 34115
CMD ["./app"]
