# Elección del sistema de integración continua

Los criterios a tener en cuenta para elegir un sistema de integración continua son los siguientes:
- **Facilidad de implementación**: el sistema debe ser fácil de configurar e implementar para pasar los tests de integración.
- **Fácil integración con GitHub**: el sistema debe ser fácil de integrar con GitHub para que se ejecute automáticamente al hacer un push.
- **Gratuito**: Al ser un proyecto pequeño, no es necesario gastar dinero por un sistema de pago con las múltiples opciones gratuitas que existen.
- **SaaS**: Se prefiere un sistema de integración continua como servicio (SaaS) para no tener que preocuparse de la infraestructura y tener que montarlo en local o en un servidor propio.

Siguiendo esos criterios, se nos presentan varias opciones interesantes:

- **CircleCI**: Es una plataforma de integración y entrega continua basada en la nube que se enfoca en la velocidad y eficiencia. Su arquitectura está basada en contenedores y permite ejecutar builds y tests en paralelo. Además, cuenta con una interfaz visual para la configuración de los pipelines. Es una buena opción a considerar.

- **Jenkins**: Es una herramienta de integración continua de código abierto que es fácil de configurar y personalizar. Está pensado para implementaciones locales e incluye una amplia variedad de complementos para una mayor compatibilidad y funcionalidad. Utiliza pipelines para automatizar el proceso de integración continua. Es una buena opción a considerar, pero peca de no ser SaaS.

- **Semaphore**: Es una plataforma de integración y entrega continua alojada en la nube que se utiliza para probar y desplegar proyectos de software alojados en GitHub. Cuenta con soporte nativo para Docker y, al igual que CircleCI, trae una interfaz visual para la configuración de los pipelines. 

- **TravisCI**: Es una herramienta de integración continua que funciona como un SaaS. Aunque es interesante y ofrece un periodo de prueba gratuito, este solo dura 14 días, por lo que no nos vale para este proyecto.

- **Buildkite**: Esta herramienta de integración y entrega continua se destaca por su capacidad de escalar y su flexibilidad. Tiene una interfaz intuitiva y permite configurar pipelines dinámicos, además de poder integrarse con Docker. Es un poco confusa a la hora de configurar los pipelines.

- **CirrusCI**: Se integra de manera fluida con GitHub y, al igual que CircleCI y Semaphore, permite la configuración como código. Al igual que las demás, se pueden hacer ejecuciones en contenedores. Este no se probará, ya que el plan gratuito requiere un repositorio público, lo cual no es posible en este caso.

- **Codefresh**: Es una plataforma de integración y entrega continua diseñada específicamente para equipos de desarrollo que construyen y despliegan aplicaciones nativas de la nube utilizando tecnologías como Docker, Kubernetes y Serverless. Los pipelines se definen en un archivo YAML y están basados en contenedores. 

- **Buddy**: Es una plataforma de automatización de DevOps y CI/CD basada en la nube diseñada para desarrolladores web. Usa una interfaz visual para la configuración de los pipelines y permite la integración con múltiples servicios y herramientas de desarrollo. Es bastante sencilla y flexible.


## Elección

Después de probar muchas de las opciones, se analizará cada una de ellas:

- **CircleCI**: Fácil de configurar y muy versátil, se ha podido configurar correctamente sin mayor problema.
- **Jenkins**: Como se comenta, al no ser SaaS, no se ha probado ya que lo suyo es tener una infraestructura propia.
- **Semaphore**: No da problemas para configurar sin hacer uso de la imagen de Docker, pero al intentar usarla, no se ha podido configurar correctamente.
- **TravisCI**: Al descubrirse que es de pago, no se ha probado.
- **Buildkite**: Es un poco confuso a la hora de configurar los pipelines, no he podido llegar a configurar uno correctamente.
- **CirrusCI**: No se ha probado por no poder hacer uso de un repositorio público, ya que es necesario para el plan gratuito. De todas formas, parece ser que cumpliría con los requisitos.
- **Codefresh**: Al intentar configurar un pipeline, solicitaba usar Kubernetes, lo cual no es necesario para este proyecto.
- **Buddy**: Es bastante sencilla y flexible, pero al no ser el repositorio público, da muchos problemas para configurar el pipeline.

En resumen, varias de las opciones podrían ser válidas. Por ejemplo, si no fuera porque el repositorio debe estar en privado, Buddy y CirrusCI sería muy buenas opciones a considerar. Semaphore, aunque haya dado algún problema que otro, también está bastante bien. Sin embargo, destacando sobre estos, CircleCI es la opción elegida por ser la que menos problemas ha dado y la que más fácil ha sido de configurar al no dar los problemas especificados con las demás opciones, así como por ser la que más se ajusta a los criterios establecidos.