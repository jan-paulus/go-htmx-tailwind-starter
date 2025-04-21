# Build go binary
FROM golang:1.24 AS go-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o ./main .

# Build   
FROM node:22 AS node-builder

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install 

COPY static/ ./static

COPY templates/ ./templates

RUN npm run build

FROM golang:1.24

WORKDIR /app

COPY --from=go-builder app/main ./main
COPY --from=node-builder app/static/ ./static
COPY --from=node-builder app/templates/ ./templates

EXPOSE 8080

ENTRYPOINT ["/app/main"]

