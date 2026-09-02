# ==========================================
# STAGE 1: Build Frontend (Nuxt 4 / PWA)
# ==========================================
FROM node:20-alpine AS frontend-builder
WORKDIR /app/frontend

COPY fun-football-admin/package*.json ./
RUN npm ci --legacy-peer-deps

COPY fun-football-admin/ ./
ENV NODE_ENV=production
RUN npm run build

# ==========================================
# STAGE 2: Build Backend (Go API & WebSocket)
# ==========================================
FROM golang:alpine AS backend-builder
WORKDIR /app/backend

# Install build dependencies
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server main.go

# ==========================================
# STAGE 3: Final Lightweight Production Image
# ==========================================
FROM alpine:3.19
WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

# Copy backend binary
COPY --from=backend-builder /app/backend/server /app/server

# Copy frontend static build & prerendered routes
COPY --from=frontend-builder /app/frontend/.output/public /app/fun-football-admin/.output/public

# Create uploads directory
RUN mkdir -p /app/uploads

ENV PORT=8080
ENV GIN_MODE=release
EXPOSE 8080

CMD ["/app/server"]
