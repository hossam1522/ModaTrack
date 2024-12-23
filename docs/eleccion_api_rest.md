# Elección de framework para API REST

Nos basaremos en los siguientes criterios para elegir el framework para la API REST:
- **Soporte**: El framework debe estar activamente mantenido, para evitar problemas de seguridad y compatibilidad.
- **Pocas dependencias**: Para evitar deuda técnica y problemas de compatibilidad, es preferible que cuantas menos dependencias tenga el framework, mejor.


Compararemos distintos frameworks para elegir el que mejor se adapte a nuestros criterios:

- **[Gin](https://github.com/gin-gonic/gin)**: Este framework es bastante popular en la comunidad de Go, y destaca por su rapidez. Es activamente mantenido, pero si miramos sus [dependencias](https://github.com/gin-gonic/gin/blob/master/go.mod), vemos que requiere de más de 38 módulos distintos.

- **[Echo](https://github.com/labstack/echo)**: Con este framework pasa algo similar que con Gin, es bastante popular y rápido, además de destacar por su simplicidad. Es activamente mantenido, y requiere de bastantes menos [dependencias](https://github.com/labstack/echo/blob/master/go.mod) que Gin, siendo necesarios 14 módulos distintos.

- **[Mux](https://github.com/gorilla/mux)**: Esta opción es muy interesante ya que no requiere de ninguna [dependencia](https://github.com/gorilla/mux/blob/main/go.mod), pero se puede ver que, en el momento de escribir esto, el último commit es de hace más de 6 meses, lo que podría ser un problema en cuanto a soporte se refiere.

- **[Chi](https://github.com/go-chi/chi)**: Framework que se desteca por su ligereza y rapidez. Al igual que el anterior, no requiere de ninguna [dependencia](https://github.com/go-chi/chi/blob/master/go.mod), pero se diferencia en que sí está activamente mantenido. Es una buena opción a considerar.

- **[Fiber](https://github.com/gofiber/fiber)**: Otro framework que destaca por su rendimiento. Activamente mantenido como la mayoría, pero requiere de más de 20 [dependencias](https://github.com/gofiber/fiber/blob/main/go.mod), lo cual lo hace menos atractivo.

- **[Goyave](https://github.com/go-goyave/goyave)**: Este framework parece estar más enfocado a proyectos medianos o grandes, ya que incluye muchas funcionalidades y lo mencionan ellos mismos. Es activamente mantenido, pero requiere de más de 45 [dependencias](https://github.com/go-goyave/goyave/blob/master/go.mod). No es la mejor opción para nuestra aplicación.


Basándonos en los criterios mencionados anteriormente, las opciones más interesantes sin Chi, Mux y Echo. En cuanto a Mux, aunque no requiere de dependencias, el hecho de que no esté activamente mantenido nos hace descartarlo. Por otro lado, Echo es activamente mantenido y no requiere de tantas dependencias en comparación con los demás frameworks, pero estando Chi que directamente ni requiere, Chi se convierte en la mejor opción para nuestra aplicación.
