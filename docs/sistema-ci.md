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

- **Buildkite**: 

- **CirrusCI**: 

- **Codefresh**

- **Buddy**