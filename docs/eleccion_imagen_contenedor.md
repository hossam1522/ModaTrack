# Elección de la imagen base del contenedor de pruebas

Para elegir correctamente la imagen base del contenedor de pruebas, tendremos en cuenta ciertos criterios, tales como:

- **Tamaño de la imagen**: Es importante que la imagen base sea lo más pequeña posible, ya que esto reducirá el espacio en disco necesario para almacenarla, además de que realmente no necesitamos todas las funcionalidades que trae una imagen completa.
- **Estabilidad**: La imagen base debe ser estable y segura, es decir, debe ser "oficial" y estar actualizada. De esta forma, nos aseguramos de que no haya vulnerabilidades en la imagen que puedan ser exploradas por un atacante.
- **Simplicidad**: La imagen base debe ser simple, es decir, debe ser fácil de usar y configurar para nuestros propósitos, con la menor cantidad de pasos posibles.

En base a estos criterios, podemos elegir una imagen base que cumpla con nuestras necesidades. Algunas de las imágenes más destacables para este proyecto son:

- **golang:alpine**: Es una imagen de Go en Alpine Linux, que es una distribución de Linux muy ligera. Es mantenida por los desarrolladores de Go y es ideal para aplicaciones que no requieren de muchas dependencias, aunque debido a ello puede ser una desventaja en algunos casos por la falta de algunas librerías.

- **golang:bookworm**: Es una imagen de Go en Debian Bookworm, que es una distribución de Linux más completa que Alpine, además de ser extremadamente estable. Un inconveniente es que es más pesada que Alpine, pero puede ser una buena opción si necesitamos más dependencias.

- **scratch**: Es una imagen vacía, es decir, no contiene nada. Es ideal para aplicaciones que no requieren de un sistema operativo completo, ya que podemos añadir solo lo que necesitamos. Es la imagen más pequeña y segura, pero también la más difícil de usar, ya que requiere de más configuración, no como las anteriores que ya vienen con Go instalado.

- **bitnami/golang**: Es una imagen de Go mantenida por Bitnami, que es una empresa conocida por sus imágenes de Docker de alta calidad, ya que traen los últimos fixes de seguridad y características lo más pronto posible. Es una buena opción, pero tiene el problema de que las imágenes siguen siendo bastante pesadas si las comparamos con golang:alpine, por ejemplo.

- **debian:stable-slim**: Es una imagen de Debian en su versión estable, pero con prácticamente ningún paquete instalado. De hecho, es una imagen ultraligera que solo contiene lo esencial para ejecutar aplicaciones. Es una buena opción si queremos reducir el tamaño de la imagen, pero pasa lo mismo que con scratch, que requiere de más configuración al no tener Go instalado, pero la configuración es bastante sencilla igualmente.


## Elección de la imagen base

Como se ha visto, existen varias imágenes a considerar para nuestro contenedor de pruebas, además de otras muchas que también son opciones interesantes pero no se han mencionado. En base a los criterios establecidos, todas las imágenes mencionadas cumplen bien el requisito de estabilidad. En cambio, para un proyecto que, de momento, no requiere apenas dependencias, y teniendo en cuenta la búsqueda de simplicidad, las imágenes **golang:alpine** y **debian:stable-slim** son las más adecuadas. Entre ambas, se ha decidido elegir **debian:stable-slim** para simplemente instalar lo mínimo necesario para ejecutar Go con una imagen que, de base, es hasta más ligera que **golang:alpine**.

Además, mencionar que la versión de Go a utilizar en el contenedor de pruebas será la última versión estable publicada, ya que es importante mantener el proyecto actualizado y con las últimas características y correcciones de seguridad. Esta versión, a la fecha de redacción de este documento, es la 1.23.3, según la [documentación oficial](https://go.dev/dl/).