# First stage: Get Golang image from DockerHub.
FROM golang:1.21.10 AS backend-builder

# Set our working directory for this stage.
WORKDIR /app

# Copy all of our files.
COPY . .

# Get and install all dependencies.
RUN CGO_ENABLED=0 go build -o server ./cmd/api/main.go

FROM node:20 AS base-frontend
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

RUN corepack enable
COPY . /frontend
WORKDIR /frontend

FROM base-frontend AS frontend-prod-deps
WORKDIR /frontend/ui
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --prod --frozen-lockfile

FROM base-frontend AS frontend-build

WORKDIR /frontend/ui
ARG SENTRY_AUTH_TOKEN

RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm run build

# Last stage: discard everything except our executables.
FROM alpine:latest AS prod

# Set our next working directory.
WORKDIR /app

# Copy our executable and our built React application.
COPY --from=backend-builder /app/server .
COPY --from=frontend-build /frontend/public ./public

ENV APP_ENV=production

# Declare entrypoints and activation commands.
EXPOSE 8000
ENTRYPOINT ["./server"]
