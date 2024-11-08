# Gestor de tareas

Para el lenguaje de programación de este proyecto, Go, existen varías opciones conocias para la gestión de tareas, de las cuales se destacan las siguientes:

- **Makefile**: Es un archivo que contiene un conjunto de reglas que permiten automatizar tareas de compilación, ejecución, limpieza, entre otras. Es muy común en proyectos de Go, y es muy útil para definir tareas comunes que se ejecutan frecuentemente.
  - **Ventajas**:
    - Ampliamente conocido y utilizado, por lo que es fácil encontrar ejemplos y documentación.
    - Muy potente y flexible, ya que permite definir tareas de forma muy detallada.
  - **Desventajas**:
    - Es bastante propenso a errores de sintaxis, lo que puede dificultar la creación de tareas complejas, además de que no es muy amigable para principiantes.
    - Algunas operaciones pueden requerir comandos Unix específicos, dificultando la portabilidad entre sistemas.

- **Taskfile**: Es una herramienta que permite definir tareas en un archivo `Taskfile.yml`, que se ejecutan mediante el comando `task`. Es una alternativa más moderna y amigable que el `Makefile`, y es muy útil para proyectos pequeños o medianos.
  - **Ventajas**:
    - Sintaxis más clara y legible, lo que facilita la creación de tareas y su mantenimiento.
    - Flexible y extensible, ya que permite definir tareas en cualquier lenguaje de scripting.
  - **Desventajas**:
    - Menos potente que el `Makefile`, lo cual lo convierte en una opción menos considerable para proyectos muy grandes o complejos.
    - Menos conocido y utilizado, por lo que puede ser más difícil encontrar ejemplos y documentación.

- **Magefile**: Es una herramienta que permite definir tareas en un archivo `magefile.go`, que se ejecutan mediante el comando `mage`. Es una alternativa más moderna y amigable que el `Makefile`, y es muy útil para proyectos pequeños o medianos.
  - **Ventajas**:
    - Utiliza Go como lenguaje de scripting, eliminando la necesidad de aprender una nueva sintaxis.
    - Altamente portable, ya que no depende de características específicas del sistema operativo.
  - **Desventajas**:
    - Solo es útil para proyectos en Go, lo que limita su uso en proyectos que utilicen otros lenguajes.
    - Al igual que el `Taskfile`, es menos conocido y utilizado que el `Makefile`, lo que puede dificultar la búsqueda de ejemplos y documentación.

## Elección del gestor de tareas

Para este proyecto, se ha decidido utilizar el `Taskfile` como gestor de tareas, debido a las siguientes razones:

- **Sintaxis clara y legible**: El uso de `YAML` para definir las tareas facilita su comprensión y mantenimiento, lo que es especialmente útil para proyectos pequeños o medianos. Adenás, siempre viene bien dominar este tipo de formato.
- **Flexibilidad y extensibilidad**: La posibilidad de definir tareas en cualquier lenguaje de scripting no nos cierra las puertas a futuras necesidades que puedan surgir en el proyecto de utilizar otros lenguajes.
- **Comunidad creciente**: Aunque es menos conocido que el `Makefile`, va ganando popularidad en la comunidad, lo que facilita la búsqueda de ejemplos y documentación.

En resumen, considero que el `Taskfile` es una opción moderna y versátil que se adapta bien a las necesidades de este proyecto, y que facilita la automatización de tareas comunes de forma sencilla y eficiente.