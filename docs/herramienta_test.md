# Herramientas para testear

Go ya incluye un paquete de testing en su librería estándar. Este paquete se llama `testing`, y es bastante completo e incluye prácticamente todas las funcionalidades necesarias para realizar tests básicos en Go. De hecho, al ser la herramienta estándar, es la más recomendada por la comunidad de Go para hacer tests.

Además, otra cosa bastante buena a considerar de `testing` es que no añade dependencias adicionales al proyecto, lo cual puede ser una ventaja a la hora de mantenerlo, ya que hace que la deuda técnica no exista en comparación con otras herramientas de testeo, las cuales sí o sí añaden dependencias.

Además, es un paquete que está bien documentado y no tiene apenas complejidad, por lo que es fácil de usar y mantener, convirtiéndolo en una buena opción para cualquier tipo de proyecto.

Basándonos en criterios de facilidad de uso y mantenimiento, evitación de la deuda técnica y recomendaciones de la comunidad, `testing` parece ser, en principio, la mejor opción para testear en Go. Aún así, existen otras librerías o paquetes que pueden ser útiles para testear, como pueden ser `testify`, `ginkgo` o `goconvey`. A continuación, se hará un pequeño análisis de ellas:

### Testify

`testify` es un conjunto de paquetes que proporcionan una variedad de funciones útiles para escribir tests en Go. Algunas de las funcionalidades que proporciona son assertions, mocks, suites, etc. Es una herramienta bastante completa y que puede ser útil en proyectos más complejos que necesiten de estas funcionalidades. Es una herramienta bastante popular y recomendada por la comunidad de Go. Más información en su [repositorio](https://github.com/stretchr/testify).

### Ginkgo

`ginkgo` es un framework de testing BDD (Behavior-Driven Development) para Go. Es una herramienta bastante completa que utiliza un lenguaje específico de dominio (DSL) que facilita la organización de las pruebas en bloques anidados, además de más funcionalidades que podemos investigar en su [repositorio](https://github.com/onsi/ginkgo). Es una herramienta bastante potente y completa, así que es una opción a tener en cuenta.

### Goconvey

`goconvey` es una herramienta de testing la cual que proporciona una interfaz web para visualizar los resultados de los tests. Es una herramienta bastante completa y fácil de usar, ya que no requiere de conocimientos avanzados para poder utilizarla además de que se integra con el propio `go test`. Aunque es una herramienta bastante buena, no ha sido actualizada en los últimos 9 meses a la hora de la realización de este documento, como se puede ver en su [repositorio](https://github.com/smartystreets/goconvey), por lo que puede que no sea la mejor opción en cuanto a deuda técnica y mantenimiento se refiere. 


## Elección de la herramienta de testeo

Todas las opciones mencionadas tienen sus ventajas y desventajas. De todas formas, como se ha dicho al principio, si seguimos criterios de facilidad de uso y mantenimiento, evitación de la deuda técnica y recomendaciones de la comunidad, `testing` parece ser la mejor opción para testear en Go, ya que aunque las demás opciones añadan funcionalidades interesantes, para este caso no se va a necesitar de ellas y solo estarían añadiendo complejidad y deuda técnica al proyecto. Por lo tanto, la elección final es `testing`.