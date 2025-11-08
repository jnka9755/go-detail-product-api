# go-detail-product-api

API simple en Go para devolver detalles de productos a partir de archivos JSON locales.

Este repositorio expone dos endpoints HTTP:

- GET /products -> lista todos los productos (enriquecidos con su categoría y vendedor)
- GET /product/{id} -> devuelve un producto por su id (por ejemplo `MLA1`)

Características principales
- Lee datos desde los ficheros JSON en `./data/` (products, categories, sellers, reviews).
- Al inicializar el repositorio carga categorías y vendedores en memoria y enriquece los productos con esas estructuras.
- Usa un error sentinel `ErrNotFound` para distinguir "producto no encontrado" de otros fallos.

Contenido del repo

- `main.go` - arranca el servidor HTTP en `localhost:8080` y registra las rutas.
- `Makefile` - tareas comunes: build, run, test, fmt, vet, tidy, deps, lint, clean.
- `data/` - JSONs de ejemplo (products.json, categories.json, sellers.json, reviews.json).
- `src/products/` - implementación: controller, service, repository y modelos.

Requisitos

- Go (recomendado 1.18+). Verifica con `go version`.
- make (opcional, para usar los targets del `Makefile`).

Instalación y ejecución (máquina nueva)

1. Clonar el repositorio

```bash
git clone https://github.com/jnka9755/go-detail-product-api
cd go-detail-product-api
```

2. Descargar dependencias

```bash
make deps
# o alternativamente
go mod download
```

3. Verificar/limpiar módulo

```bash
make tidy
```

4. Ejecutar la aplicación

```bash
make run
# o
go run main.go
```

El servidor escuchará en `localhost:8080`.

Comandos útiles (Makefile)

- `make run` -> ejecuta `go run main.go`
- `make run-dev` -> ejecuta con detector de race `go run -race main.go`
- `make test` -> ejecuta `go test ./...`
- `make tidy` -> `go mod tidy`
- `make deps` -> `go mod download`

Uso / ejemplos

Listar productos:

```bash
curl -s http://localhost:8080/products | jq
```

Obtener un producto por id (por ejemplo `MLA1`):

```bash
curl -s http://localhost:8080/product/MLA1 | jq
```

Notas importantes

- Rutas y ficheros de datos: los JSON se cargan con rutas relativas (`data/*.json`). Ejecuta el servidor desde la raíz del proyecto para que las rutas funcionen correctamente.
- Si falta algún fichero JSON, el repositorio lo registrará en stdout/stderr y la respuesta puede ser una lista vacía o un error interno.
- El controlador responde con HTTP 404 cuando se devuelve el sentinel `ErrNotFound` (producto no encontrado). Para otros errores responde 500.

Depuración rápida

- Para ver los logs del servidor ejecuta `make run` y observa la salida en la consola.
- Si quieres ver fallos de parsing de JSON, revisa los mensajes que se loguean durante `NewRepository` y la carga de ficheros.

Extensiones sugeridas

- Hacer que las funciones de carga (`loadCategories`, `loadSellers`) retornen errores y exponerlos para retornar 500 desde el inicio si fallan.
- Añadir tests unitarios para el repositorio (cargar ficheros de prueba) y para el servicio.
- Añadir paginación y filtros a `GET /products`.

Contacto

Si necesitas ayuda adicional o quieres que ejecute la aplicación y verifique la configuración en este entorno, dime y lo hago.
