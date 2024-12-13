# Elección del sistema de logs

Nos basaremos en los siguientes criterios para elegir el sistema de logs:
- **Soporte**: El sistema de logs debe estar activamente mantenido, para evitar problemas de seguridad y compatibilidad.
- **Pocas dependencias**: Para evitar deuda técnica y problemas de compatibilidad, es preferible que cuantas menos dependencias tenga el sistema de logs, mejor.
- **Versión mínima de Go**: La versión mínima de Go que soporta el sistema de logs debe ser la 1.18, evitando así problemas de compatibilidad y tener que modificar los pipelines de CI/CD, ya que esa es la versión mínima actualmente utilizada en los tests de integración de la aplicación, establecida en la [documentación](sistema-ci.md).
- **Logs estructurados**: Los logs deben ser estructurados, para facilitar su análisis y búsqueda.

Compararemos distintos sistemas de logs para elegir el que mejor se adapte a nuestros criterios:

- **[Zerolog](https://github.com/rs/zerolog)**: Zerolog es una librería de logs estructurados para Go que se caracteriza por su rapidez y su facilidad de uso. Tiene pocas dependencias y es activamente mantenido. La versión mínima requerida de Go es la 1.15, así que es totalmente válida en ese aspecto.
- **[Zap](https://github.com/uber-go/zap)**: Zap es una librería de logs que se destaca principalmente por su rendimiento, además de ser también estructurados. También es activamente mantenido, pero tiene más dependencias que Zerolog, además de que la versión mínima de Go requerida es la 1.19, por lo que esta librería queda descartada.
- **[Slog](https://pkg.go.dev/log/slog)**: Es una librería estándar de Go que proporciona logs estructurados. Al ser estándar, es práctimante seguro que va a seguir siendo mantenido, además de que no tiene dependencias. El problema es que se introdujo en la versión 1.21 de Go, por lo que queda descartado.
- **[Phuslu/log](https://github.com/phuslu/log)**: Inspirada en Zerolog, es una librería de logs estructurados que se caracteriza por su rapidez, siendo incluso más rápida que Zap. La versión mínima de Go requerida es la 1.18, la necesaria según los criterios, además de que no tiene dependencias. Es activamente mantenido, por lo que es una buena opción.
- **[Logrus](https://github.com/sirupsen/logrus)**: Es una librería de logs estructurados que es muy popular en la comunidad de Go. Cumple también con los criterios, al no tener apenas dependencias y ser activamente mantenido., además de requerir una versión mínima de la 1.13.

Si nos aferramos a los criterios establecidos, las opciones que cumplen con todos ellos son Zerolog, Phuslu/log y Logrus. Como, entre todas ellas, la que menos dependencias tiene es Phuslu/log (ninguna), destaca sobre el resto y es la que elegiremos para nuestro sistema de logs.