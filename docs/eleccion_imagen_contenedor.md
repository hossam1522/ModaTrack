# Elección de la imagen base del contenedor de pruebas

Para elegir correctamente la imagen base del contenedor de pruebas, tendremos en cuenta ciertos criterios, tales como:

- **Tamaño de la imagen**: Es importante que la imagen base sea lo más pequeña posible, ya que esto reducirá el espacio en disco necesario para almacenarla, además de que realmente no necesitamos todas las funcionalidades que trae una imagen completa.
- **Seguridad**: La imagen base debe ser segura, es decir, debe ser oficial y estar actualizada. De esta forma, nos aseguramos de que no haya vulnerabilidades en la imagen que puedan ser exploradas por un atacante.
- **Rendimiento**: La imagen base debe tener un buen rendimiento, es decir, debe ser rápida y eficiente. 
- **Simplicidad**: La imagen base debe ser simple, es decir, debe ser fácil de usar y configurar para nuestros propósitos, con la menor cantidad de pasos posibles.

En base a estos criterios, podemos elegir una imagen base que cumpla con nuestras necesidades. Algunas de las imágenes más destacables para este proyecto son:

- **golang:alpine**: Es una imagen de Go en Alpine Linux, que es una distribución de Linux muy ligera. Es mantenida por los desarrolladores de Go y es ideal para aplicaciones que no requieren de muchas dependencias, aunque debido a ello puede ser una desventaja en algunos casos por la falta de algunas librerías.

- **golang:bookworm**: Es una imagen de Go en Debian Bookworm, que es una distribución de Linux más completa que Alpine, además de ser extremadamente estable. Un inconveniente es que es más pesada que Alpine, pero puede ser una buena opción si necesitamos más dependencias.

- **scratch**: Es una imagen vacía, es decir, no contiene nada. Es ideal para aplicaciones que no requieren de un sistema operativo completo, ya que podemos añadir solo lo que necesitamos. Es la imagen más pequeña y segura, pero también la más difícil de usar, ya que requiere de más configuración, no como las anteriores que ya vienen con Go instalado.


## Elección de la imagen base

Como se ha visto, existen varias imágenes a considerar para nuestro contenedor de pruebas, además de otras muchas que también son opciones interesantes pero no se han mencionado. En base a los criterios establecidos, todas las imágenes mencionadas cumplen bien los requisitos de seguridad y rendimiento. En cambio, para un proyecto que, de momento, no requiere apenas dependencias, y teniendo en cuenta la búsqueda de simplicidad, se ha decidido elegir la imagen **golang:alpine** como base para el contenedor de pruebas, por ser ligera (más que **golang:bookworm**) y por tener Go ya instalado, lo que nos ahorra tiempo y configuración (a diferencia de **scratch**).

En específico, se usará la última versión hasta el momento publicada en [Golang Docker Official Image](https://hub.docker.com/_/golang), la cual en este caso es **golang:1.23-alpine3.20**.