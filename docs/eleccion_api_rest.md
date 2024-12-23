# Elección de framework para API REST

Nos basaremos en los siguientes criterios para elegir el framework para la API REST:
- **Soporte**: El framework debe estar activamente mantenido, para evitar problemas de seguridad y compatibilidad.
- **Pocas dependencias**: Para evitar deuda técnica y problemas de compatibilidad, es preferible que cuantas menos dependencias tenga el framework, mejor.


Compararemos distintos frameworks para elegir el que mejor se adapte a nuestros criterios:

- **[Gin](https://github.com/gin-gonic/gin)**: Este framework es bastante popular en la comunidad de Go, y destaca por su rapidez. Es activamente mantenido, pero tiene bastantes dependencias. La versión mínima de Go requerida es la 1.21, por lo que queda descartado.
- **[Echo](https://github.com/labstack/echo)**: Con este framework pasa algo similar que con Gin, es bastante popular y rápido, pero tiene bastantes dependencias, aunque bastante menos que Gin. La versión mínima de Go requerida es la 1.20, por lo que tampoco es válido.
- **[Mux](https://github.com/gorilla/mux)**: Esta opción es muy interesante ya que no tiene dependencias, pero no está activamente mantenido. Además, la versión mínima de Go requerida es la 1.20, así que lo mismo que con los anteriores, queda descartado.
- **[Chi](https://github.com/go-chi/chi)**: Al igual que el anterior, no tiene dependencias, pero se diferencia en que sí está activamente mantenido. Además, cumple con la versión mínima de Go requerida, por lo que es una buena opción.
- **[Fiber](https://github.com/gofiber/fiber)**: No tiene demasiadas dependencias pero sigue teniendo bastantes. La versión mínima de Go requerida es la 1.23, y está activamente mantenido.
- **[Swag](https://github.com/swaggo/swag)**: Esta es otra opción muy interesante, ya que cumple con la versión mínima de Go requerida y está activamente mantenido. Su único problema es que tiene bastantes dependencias.
- **[Go-swagger](https://github.com/go-swagger/go-swagger)**: Al igual que la mayoría, cumple el requisito de soporte, pero acaba quedando descartado por tener bastantes dependencias y requerir una versión superior a la 1.18.
- **[Goyave](https://github.com/go-goyave/goyave)**: Finalmente, lo mismo que con el anterior.


Basándonos en los criterios mencionados anteriormente, las opciones más adecuadas serían Chi y Swag, ya que ambos son sistemas activamente mantenidos y cumplen con el criterio de versión mínima. Sin embargo, dado que Chi no tiene dependencias, es la opción más adecuada para nuestra aplicación. Por lo tanto, elegiremos Chi como framework para nuestra API REST.
