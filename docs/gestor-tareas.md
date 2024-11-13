# Gestor de tareas

La elección de un gestor de tareas es una decisión importante en cualquier proyecto de software, ya que puede tener un gran impacto en la eficiencia y mantenibilidad del proyecto. Un gestor de tareas es una herramienta que permite automatizar tareas comunes en el desarrollo de software, como la instalación de dependencias, la ejecución de tests, la compilación del código, entre otras.

Para elegir correctamente un gestor de tareas para este proyecto, los siguientes aspectos deben ser considerados:

- Lenguaje de programación del proyecto. En este caso, Go es el lenguaje de programación usado, así que habrá que tener en cuenta las opciones disponibles para este lenguaje.
- Las mejores prácticas y herramientas recomendadas para Go.
- Reducción de la deuda técnica a la hora de mantener y escalar el proyecto.
- Facilidad de uso y mantenimiento.
- Posibilidad de extender a otros lenguajes de programación en el futuro si fuera necesario.

Teniendo en cuenta estos criterios establecidos, se nos presentan, principalmente, dos opciones para la elección del gestor de tareas: `Make` y `Task`. 

Hay otros gestores de tareas que podrían ser considerados, como `Mage`, pero este se ha descartado por el último criterio mencionado, ya que solo es útil para proyectos en Go, lo que limita su uso en proyectos que utilicen otros lenguajes.

Compararemos `Make` y `Task` para determinar cuál es la mejor opción para este proyecto en cada uno de los criterios mencionados:

- **Lenguaje de programación del proyecto**: Ambos gestores de tareas son compatibles con Go, por lo que no hay diferencias significativas en este aspecto.
- **Mejores prácticas y herramientas recomendadas para Go**: Ambas herramientas son ampliamente utilizadas en proyectos de Go, por lo que cumplen con este criterio.
- **Reducción de la deuda técnica**: Aunque generalmente se considera que `Make` es más propenso a errores de sintaxis y menos flexible que `Task`, realmente si se siguen las buenas prácticas y se mantienen los archivos de configuración limpios y organizados, no debería haber una gran diferencia en este aspecto, menos todavía en un proyecto como este.
- **Facilidad de uso y mantenimiento**: En este punto, tal y como indica su propia documentación oficial, `Task` busca ser una alternativa más simple y fácil de usar que otros gestores de tareas, tales como `Make`. Aunque el proyecto no sea lo suficimiente complejo como para que la sintaxis de `Make` sea un problema, `Task` es el ganador en este aspecto.
- **Posibilidad de extender a otros lenguajes de programación en el futuro**: En este aspecto, ambos gestores de tareas son igualmente válidos, ya que ambos son compatibles con múltiples lenguajes de programación.

## Elección del gestor de tareas

Como se ha podido comprobar, ambos gestores de tareas son más que válidos para este proyecto, pudiendo usarse cualquiera de ellos sin problemas, ya que siguiendo los criterios establecidos, ambos cumplen con las necesidades del proyecto. Sin embargo, debido al criterio de facilidad de uso y mantenimiento, se ha decidido utilizar `Task` como gestor de tareas para este proyecto.