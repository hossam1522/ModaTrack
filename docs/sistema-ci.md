# Elección del sistema de integración continua

Los criterios a tener en cuenta para elegir un sistema de integración continua son los siguientes:
- **Facilidad de implementación**: el sistema debe ser fácil de configurar e implementar para pasar los tests de integración.
- **Fácil integración con GitHub**: el sistema debe ser fácil de integrar con GitHub para que se ejecute automáticamente al hacer un push.
- **Gratuito**: Al ser un proyecto pequeño, no es necesario gastar dinero por un sistema de pago con las múltiples opciones gratuitas que existen.

Siguiendo esos criterios, se nos presentan varias opciones interesantes:

- **CircleCI**: Es una de las herramientas más utilizadas para la integración continua. Es fácil de configurar, además de tratarse de un SaaS con un plan gratuito para proyectos de código abierto.

- **Jenkins**: Es una herramienta de integración continua de código abierto que es fácil de configurar y personalizar. Está pensado para implementaciones locales e incluye una amplia variedad de complementos para una mayor compatibilidad y funcionalidad. Utiliza pipelines para automatizar el proceso de integración continua.

- **Semaphore**: Parecido a CircleCI, es una herramienta de integración continua que funciona como un SaaS. Ofrece un plan gratuito y hace uso de pipelines y jobs para automatizar el proceso de integración continua.

- **Travis CI**: Es una herramienta de integración continua que funciona como un SaaS. Aunque es interesante y ofrece un periodo de prueba gratuito, este solo dura 14 días, por lo que no nos vale para este proyecto.

- **Buildkite**: Al igual que Jenkins, esta ofrece una gran cantidad de funcionalidades en forma de plugins. Esta está centrada en ser simple y fácil de usar, así que es una buena opcón a considerar.

- **CirrusCI**: Parecido a Semaphore y a CircleCI, es otra herramienta de integración continua que funciona como un SaaS y que ofrece opciones gratuitas para proyectos de código abierto.