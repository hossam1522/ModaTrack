# Gestor de tareas

Para elegir correctamente un gestor de tareas para este proyecto, los siguientes aspectos deben ser considerados:

- Lenguaje de programación del proyecto. En este caso, Go es el lenguaje de programación usado, así que habrá que tener en cuenta las opciones disponibles para este lenguaje.
- Las mejores prácticas y herramientas recomendadas para Go.
- Reducción de la deuda técnica a la hora de mantener y escalar el proyecto.
- Facilidad de uso y mantenimiento.

Se nos presentan varias opciones para la elección del gestor de tareas, entre las que se encuentran `Make`, `Task`, `Mage`, `Just`, entre otros. Analizaremos principalmente estos cuatro mencionados, ya que se encuentran entre los más recomendados por la comunidad de Go.

Teniendo en cuenta los criterios mencionados anteriormente, se procederá a hacer un análisis punto por punto para ver cuál de los gestores de tareas es el más adecuado para este proyecto:

- **Lenguaje de programación del proyecto**: Todos los gestores mencionados son compatibles con Go, por lo que no hay problemas en este aspecto.

- **Mejores prácticas y herramientas recomendadas para Go**: Como se ha mencionado anteriormente, se ha decidido comparar estos gestores de tareas debido a que son de los más recomendados, por no decir los que más, por la comunidad de Go. De todas formas, entre los mencionados, hay algunos más recomendados que otros, siendo el más recomendado `Make`, y el menos recomendado `Just`.

- **Reducción de la deuda técnica**: En este aspecto, los gestores de tareas que menos deuda técnica acumularán son aquellos que solo usen el lenguaje de programación en el que se está trabajando, ya que esto facilita el mantenimiento del proyecto, además de aquellos que se actualizan con frecuencia y siguen activos. Si se realiza una investigación, se puede llegar a ver que:
  - **Mage**: Aunque solo use Go, cosa que reduce la deuda técnica, en el momento de la realización de este documento  el repositorio no se ha actualizado desde hace ya un año, lo que puede ser un problema a la hora de mantener el proyecto. Si no fuera por esto, probablemente sería el ganador en este aspecto.
  - **Make**: Como en todos los siguientes casos, al contrario que Mage, Make no es un gestor de tareas específico para Go, lo que puede hacer que la deuda técnica aumente. El punto a favor es que el proyecto sigue siendo activo y se sigue actualizando.
  - **Just**: Al igual que Make, Just no es un gestor de tareas específico para Go, lo que puede aumentar la deuda técnica respecto a Mage, por ejemplo. Sin embargo, el proyecto es activo y se sigue actualizando.
  - **Task**: Aunque no sea específico de Go, al igual que los dos anteriores, este proyecto es bastante activo, así que no debería haber deuda técnica en este aspecto.

- **Facilidad de uso y mantenimiento**: Teniendo en cuenta la complejidad del proyecto, ninguno de los gestores de tareas mencionados debería ser un problema en este aspecto. Sin embargo, objetivamente hay algunos más fáciles de usar que otros. `Mage` es el más sencillo de usar ya que, como se menciona en el apartado anterior, usa Go, eliminando la necesidad de aprender una nueva sintaxis. `Task` es el segundo más sencillo de usar, ya que su sintaxis es bastante simple y fácil de entender, además de ser una herramienta hecha para ser más simple y fácil de usar que otros gestores, tales como `Make` y `Just`, que son los más complicados de usar, ya que su sintaxis es más compleja y menos intuitiva.

## Elección del gestor de tareas

Teniendo en cuenta todos los criterios establecidos y explicados, probablemnte la opción más lógica sería el uso de `Mage`, si no fuera por el hecho de que el proyecto lleva sin actualizarse un año, lo cual es mala señal ya que puede ser que el proyecto haya sido abandonado, incurriendo así en deuda técnica en caso de elegirlo.

Por lo tanto, debido a los criterios establecidos, se ha decidido usar `Task` como gestor de tareas para este proyecto, ya que aunque no sea tan recomendado como `Make`, sigue siendo muy recomendado por la comunidad, además de ser fácil de usar y, finalmente, ser un proyecto activo que se sigue actualizando a día de hoy.