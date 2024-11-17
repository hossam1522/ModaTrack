# Herramientas para testear

Para elegir correctamente las herramientas para poder realizar pruebas para comprobar el correcto funcionamiento del proyecto, los siguientes aspectos deben ser considerados:

- Lenguaje de programación del proyecto. En este caso, Go es el lenguaje de programación usado, así que habrá que tener en cuenta las opciones disponibles para este lenguaje.
- Las mejores prácticas y herramientas recomendadas para Go.
- Reducción de la deuda técnica a la hora de mantener y escalar el proyecto.
- Facilidad de uso y mantenimiento.
- Facilidad para creación de mocks (implementaciones simuladas de objetos que simulan el comportamiento de los objetos reales) para testear el código.

Se nos presentan varias opciones para la elección de las herramientas de testeo, entre las que se encuentran `testing`, `testify`, `ginkgo`, `goconvey`, entre otros. Analizaremos principalmente estos cuatro mencionados, ya que se encuentran entre los más recomendados por la comunidad de Go.

Teniendo en cuenta los criterios mencionados anteriormente, se procederá a hacer un análisis punto por punto para ver cuál de las herramientas de testeo es la más adecuada para este proyecto:

- **Lenguaje de programación del proyecto**: Todas las herramientas mencionadas son compatibles con Go, por lo que no hay problemas en este aspecto.

- **Mejores prácticas y herramientas recomendadas para Go**: La comunidad de Go recomienda principalmente `testing`, ya que es la herramienta estándar de Go para realizar tests. `testify` también es una buena opción, probablemente la más famosa después de `testing`, ya que añade funcionalidades adicionales a `testing` que pueden ser útiles. `ginkgo` y `goconvey` son menos recomendadas, pero siguen siendo opciones válidas.

- **Reducción de la deuda técnica**: La herramienta que gana en este aspecto es `testing`, ya que es la herramienta estándar de Go y no añade dependencias adicionales al proyecto, lo cual supone una ventaja a la hora de mantenerlo. Las demás herramientas añaden dependencias adicionales, lo cual puede aumentar la deuda técnica. Otra cuestión a considerar es la activad de los proyectos. En este apartado, todos los proyectos están activos excepto `goconvey`, que no ha sido actualizado en los últimos 9 meses a la hora de la realización de este documento.

- **Facilidad de uso y mantenimiento**: En un proyecto como el actual, este apartado no debería ser un problema ya que el proyecto no es lo suficientemente complejo como para que la facilidad de uso y mantenimiento de las herramientas de testeo sea un problema. De todas formas, al no tener dependencias externas, se puede considerar que `testing` es la herramienta más fácil de usar y mantener. 

- **Facilidad para creación de mocks**: `testing` no tiene una forma nativa de crear mocks, lo cual puede ser un problema si se necesita crear mocks para testear el código, aunque existen maneras de hacerlo sin necesidad de una librería específica, como viene explicado en el siguiente [artículo](https://medium.com/safetycultureengineering/flexible-mocking-for-testing-in-go-f952869e34f5). Las demás herramientas mencionadas facilitan la creación de mocks, lo cual puede ser una ventaja en este aspecto.


## Elección de la herramienta de testeo

Siguiendo los criterios establecidos y después de haber realizado un análisis punto por punto, exceptuando `goconvey` por el tema del aumento de la deuda técnica y la falta de actualizaciones, todas las opciones pueden ser perfectamente válidas para este proyecto. Sin embargo, en la mayoría de aspectos, siempre ha habido una herramienta que ha destacado sobre las demás, y esa ha sido `testing`. Por lo tanto, se ha decidido usar `testing` como herramienta de testeo para este proyecto por todas las razones explicadas anteriormente.