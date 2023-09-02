# Fase de compilación
FROM golang:1.20 AS builder

WORKDIR /app

# Copiar el código fuente y los archivos necesarios
COPY . .

# Descargar las dependencias
RUN go mod download

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o satelite-service ./cmd/main

# Fase de producción
FROM scratch

# Exponer el puerto en el que la aplicación escuchará
EXPOSE 8080

# Copiar el binario compilado de la fase de compilación
COPY --from=builder /app/satelite-service /satelite-service

# Establecer el punto de entrada
ENTRYPOINT ["/satelite-service"]
