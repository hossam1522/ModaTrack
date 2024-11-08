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

## Herramientas utilizadas

- Lenguaje de programación: Go

- Gestor de dependencias: Go Modules. Para más información, ver [Gestor de dependencias](docs/gestor_dependencias.md).

- Gestor de tareas: Task. Para más información, ver [Gestor de tareas](docs/gestor_tareas.md).
  - `task install`: Instala las dependencias del proyecto.
  - `task update`: Actualiza las dependencias del proyecto.
  - `task test`: Ejecuta los tests del proyecto.
  - `task default`: Muestra la lista de tareas disponibles.
  
  Para poder utilizar Task, es necesario tenerlo instalado. Basta con ejecutar el siguiente comando:
  ```bash
  go install github.com/go-task/task/v3/cmd/task@latest
  ```

### Documentación adicional

- [Licencia](LICENSE)
- [Documentación adicional](documentacion_adicional)