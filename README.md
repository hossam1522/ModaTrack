# ModaTrack

## Descripción del problema

Soy un vendedor de ropa y actualmente enfrento grandes dificultades para gestionar el inventario de mi tienda, lo que impacta negativamente en mis ventas y en la eficiencia de mi negocio. Esto es un gran problema debido a lo siguiente:

- Falta de visibilidad del inventario: No tengo una forma clara y rápida de saber cuántas prendas tengo en stock de cada producto. Esto me lleva a dos problemas principales:
  - Sobrestock: A menudo compro más prendas de las necesarias, lo que inmoviliza capital en productos que no se venden fácilmente.
  - Faltantes de stock: En ocasiones me quedo sin productos que tienen una alta demanda, lo que genera pérdidas en ventas potenciales.

- No tengo una forma clara de planificar qué ropa debo priorizar en mis pedidos y cuáles no vale la pena seguir vendiendo. Esto me lleva a invertir dinero en productos poco rentables, y a perder oportunidades de ventas en productos más populares o rentables.

- No tengo una forma eficiente de gestionar el cambio de ropa según la temporada. Esto provoca que mantenga prendas fuera de temporada en stock, lo que ocupa espacio y puede resultar en pérdidas si no se venden a tiempo.

## Planificación del proyecto

- [Milestones](docs/milestones.md)
- [User Stories](docs/user_stories.md)
- [User Journeys](docs/user_journeys.md)
- [Gestor de dependencias](docs/gestor_dependencias.md)
- [Gestor de tareas](docs/gestor_tareas.md)
- [Herramienta para testear](docs/herramienta_test.md)
- [Metodología para creación de tests](docs/metodologia_tests.md)
- [Elección de la imagen base del contenedor de pruebas](docs/eleccion_imagen_contenedor.md)
- [Sistema de configuración](docs/eleccion_sistema_config.md)
- [Sistema de logs](docs/eleccion_sistema_logs.md)

## Herramientas utilizadas

- Lenguaje de programación: Go

- Gestor de dependencias: Go Modules. Para más información, ver [Gestor de dependencias](docs/gestor_dependencias.md).

- Gestor de tareas: Task. Para más información, ver [Gestor de tareas](docs/gestor_tareas.md).
  - `task install`: Instala las dependencias del proyecto.
  - `task update`: Actualiza las dependencias del proyecto.
  - `task test`: Ejecuta los tests del proyecto.
  - `task default`: Muestra la lista de tareas disponibles.
  - `task check`: Verifica la sintaxis de los archivos del proyecto.
  
  Para poder utilizar Task, es necesario tenerlo instalado. Basta con ejecutar el siguiente comando:
  ```bash
  go install github.com/go-task/task/v3/cmd/task@latest
  ```

## Desarrollo del proyecto

Para el desarrollo del proyecto, se ha seguido la metodología TDD, que consiste en escribir los tests antes de escribir el código. Esto permite asegurarse de que el código se comporta como se espera y que no se introduce ningún bug. Se ha seguido el principio FIRST, para que los tests no
supongan ser un problema en vez de algo que ayude a desarrollar el proyecto correctamente. Para
más información, ver [Metodología para creación de tests](docs/metodologia_tests)

Las principales estructuras testeadas han sido aquellas relacionadas con las ventas
y el stock, ya que son las que más impacto tienen en el negocio y realmente la base de 
todo el sistema. Como se ha mencionado antes, para ejecutar los tests, se utiliza la herramienta Task, mediante el comando `task test`.

## Contenedor de pruebas

Para construir el contenedor de pruebas y ejecutar los tests en un entorno aislado, se ha elegido la imagen base **golang:alpine**. Para más información, ver [Elección de la imagen base del contenedor de pruebas](docs/eleccion_imagen_contenedor.md). Para probar la implementación, basta con ejecutar el siguiente comando:

```bash
docker build -t hossam1522/modatrack . && docker run -u 1001 -t -v "$(pwd):/app/test" hossam1522/modatrack
```

En caso de no querer usar el Dockerfile local, se puede utilizar la imagen ya construida en [Docker Hub](https://hub.docker.com/repository/docker/hossam1522/modatrack), mediante el siguiente comando:

```bash
docker run -u 1001 -t -v "$(pwd):/app/test" hossam1522/modatrack
```

### Documentación adicional

- [Licencia](LICENSE)
- [Documentación adicional](documentacion_adicional)