# test-bia :zap:

Microservicio construido en Go para almancenar informacion de los satelites, mostrar sus mensajes y ubicación implementando una arquitectura limpia para separar las entidades y servicios de dominio de la infraestructura.

El servicio se expone a través de una API.

## Requisitos

- Docker

## Inicialización

1. **Construir el Contenedor**: Usa el siguiente comando para construir la imagen del microservicio:

   ```shell
   docker build -t satelite-service .
   ```

2. **Ejecutar el Contenedor**: Para ejecutar el contenedor, utiliza el siguiente comando:

   ```shell
   docker run -p 8080:8080 --network host satelite-service .
   ```

## Estructura del Proyecto

La estructura del proyecto sigue una orientación a capas de la arquitectura: domain - application - infrastructure. Aunque se consideró el Go standard project layout, que es más go-idiomático, se optó por la opción orientada a capas debido a su mayor semántica en relación con la arquitectura utilizada.

## Tecnologías y Librerías

- **Framework**: [Echo](https://github.com/labstack/echo)
- **Librerías**:
  - validatorv10
  - testify

## Guías y Convenciones

Se siguieron algunas recomendaciones del documento [Effective Go](https://golang.org/doc/effective_go) para mantener un código idiomático, a la vez que intento complementarlo con Clean Code.

## Uso de la API

Este microservicio ofrece diferentes endpoint para obtener datos. A continuación, encontrarás detalles sobre cómo utilizarla.

###

**Endpoint:** `Post /topsecret/`

**Ejemplo de Petición:**

```http
POST /topsecret_split/kenobi
```

**Respuesta:**

```json
{
  "location": {
    "x": -482.33545,
    "y": 1544.0127
  },
  "message": "este es un mensaje "
}
```

Guardar datos de un satélite

**Endpoint:** `POST /topsecret_split/{satelite_name}`

**Ejemplo de Petición:**

```http
POST /topsecret_split/kenobi
```

- Body

```json
{
  "name": "kenobi",
  "distance": 100.0,
  "message": ["este", "", "", "mensaje", ""]
}
```

**Respuesta:**

```json
{
  "distance": 100.0,
  "message": ["este", "", "", "mensaje", ""]
}
```

**Endpoint:** `GET /topsecret_split/{satelite_name}`

**Ejemplo de Petición:**

```http
GET /topsecret_split/kenobi
```

**Respuesta:**

```json
{
  "location": {
    "x": -482.33545,
    "y": 1544.0127
  },
  "message": "este es un mensaje "
}
```
