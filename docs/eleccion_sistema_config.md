# Elección del sistema de configuración

Nos basaremos en los siguientes criterios para elegir el sistema de configuración:
- **Soporte**: El sistema de configuración debe estar activamente mantenido, para evitar problemas de seguridad y compatibilidad.
- **Seguridad**: Al estar almacenando información sensible, es importante que el sistema de configuración sea seguro y no tenga vulnerabilidades conocidas.
- **Pocas dependencias**: Para evitar deuda técnica y problemas de compatibilidad, es preferible que cuantas menos dependencias tenga el sistema de configuración, mejor.

Compararemos distintos sistemas de configuración para elegir el más adecuado:

- **[GoDotEnv](https://github.com/joho/godotenv)**: Este es un sistema muy simple que lee variables de entorno desde un archivo `.env`. Aunque tiene algún commit reciente, no parece estar muy activamente mantenido, aunque tiene una comunidad activa. Según [Synk](https://snyk.io/advisor/golang/github.com/joho/godotenv), no tiene problemas de seguridad conocidos.
- **[Viper](https://github.com/spf13/viper)**: Este es más complejo que GoDotEnv, pero tiene más características, por lo que es más completo. Es un proyecto activamente mantenido y tiene una comunidad grande. Según [Synk](https://snyk.io/advisor/golang/github.com/spf13/viper), no tiene problemas de seguridad conocidos, con una mejor puntuación de seguridad que GoDotEnv. El problema es que requiere bastantes dependencias.
- **[Cleanenv](https://github.com/ilyakaznacheev/cleanenv)**: Al igual que GoDotEnv, es un sistema que busca ser simple y fácil de usar, pero a su vez apenas tiene actividad en el proyecto, aunque no tiene problemas de seguridad conocidos según [Synk](https://snyk.io/advisor/golang/github.com/ilyakaznacheev/cleanenv).
- **[Koanf](https://github.com/knadh/koanf)**: Alternativa muy parecida a Viper, pero que según ellos mismos busca ser más limpia y ligera, haciendo uso de muchas menos dependencias. Es un proyecto activamente mantenido y no tiene problemas de seguridad conocidos según [Synk](https://snyk.io/advisor/golang/github.com/knadh/koanf/v2).
- **[Env](https://github.com/caarlos0/env)**: Al igual que GoDotEnv, es un sistema muy simple que lee variables de entorno desde un archivo `.env`, aunque a diferencia de este, este proyecto es muy activo y tiene una comunidad grande. Al igual que los demás, no tiene problemas de seguridad conocidos según [Synk](https://snyk.io/advisor/golang/github.com/caarlos0/env/v11), pero destaca sobre todo no usar ninguna dependencia.


Si nos aferramos a los criterios mencionados anteriormente, las opciones más adecuadas serían Koanf y Env, ya que ambos son sistemas activamente mantenidos, seguros y con pocas dependencias. Sin embargo, dado que Env no tiene dependencias, es la opción más adecuada para nuestro proyecto. Por lo tanto, elegiremos Env como sistema de configuración para nuestro proyecto.