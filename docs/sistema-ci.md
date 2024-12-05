# Elección del sistema de integración continua

Los criterios a tener en cuenta para elegir un sistema de integración continua son los siguientes:
- **Gratuito**: La idea es que el sistema sea gratuito, por lo menos de forma parcial, para poder usarlo sin coste en nuestro proyecto. También es importante que la versión gratuita ofrecida, más o menos completa, no deje de funcionar pasado un tiempo de prueba.

Siguiendo esos criterios, se nos presentan varias opciones interesantes. Después de probar muchas de las opciones, se analizará cada una de ellas:

- **[CircleCI](https://circleci.com/)**: Fácil de configurar y muy versátil, se ha podido configurar correctamente sin mayor problema.
- **[Semaphore](https://semaphoreci.com/)**: No da problemas para configurar sin hacer uso de la imagen de Docker, pero al intentar usarla, no se ha podido configurar correctamente.
- **[TravisCI](https://www.travis-ci.com/)**: Al descubrirse que es de pago, no se ha probado.
- **[Buildkite](https://buildkite.com/)**: Es un poco confuso a la hora de configurar los pipelines, no he podido llegar a configurar uno correctamente.
- **[CirrusCI](https://cirrus-ci.org/)**: No se ha probado por no poder hacer uso de un repositorio público, ya que es necesario para el plan gratuito. De todas formas, parece ser que cumpliría con los requisitos.
- **[Codefresh](https://codefresh.io/)**: Al intentar configurar un pipeline, solicitaba usar Kubernetes, lo cual no es necesario para este proyecto.
- **[Buddy](https://buddy.works/)**: Es bastante sencilla y flexible, pero al no ser el repositorio público, da muchos problemas para configurar el pipeline.

En resumen, varias de las opciones podrían ser válidas. Por ejemplo, si no fuera porque el repositorio debe estar en privado, Buddy y CirrusCI sería muy buenas opciones a considerar. Semaphore, aunque haya dado algún problema que otro, también está bastante bien. Sin embargo, destacando sobre estos, CircleCI es la opción elegida por ser la que menos problemas ha dado y la que más fácil ha sido de configurar al no dar los problemas especificados con las demás opciones, así como por ser la que más se ajusta a los criterios establecidos.

## Versiones probadas

En GitHub Actions, se ha decidido probar la versión más antigua de Go que soporta los módulos, la 1.11, ya que si
funciona en esta versión, es casi seguro que funcionará en las versiones posteriores. También se ha decidido probar con una de las versiones oficiales estables, la 1.22.10, para asegurarse de que el proyecto funciona en las versiones estables más recientes.

Por otro lado, en CircleCI, se ha decidido probar la versión de Go asociada a la imagen de Docker, que a su vez es la versión en la que se está desarrollando el proyecto, por lo que tiene sentido hacerlo.