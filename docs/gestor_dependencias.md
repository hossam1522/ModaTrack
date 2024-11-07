# Gestor de dependencias

Como indica la propia documentación del lenguaje, Go maneja dependencias mediante módulos. Un módulo es un conjunto de paquetes que se publican, versionan y distribuyen juntos, proporcionando todo lo necesario para poder gestionar las dependencias del proyecto.

Esto se introdujo en Go 1.11, y se ha ido mejorando en versiones posteriores. Se conoce como **Go Modules**, y tiene la ventaja de que no es necesario depender de un sistema de gestión de dependencias externo, como **Dep** o **Glide**, gestores de dependencias que han ido en desuso con la introducción de los módulos. 

Es lógico pensar que la opción más recomendable es utilizar el propio gestor de dependencias oficial, el cual las gestiona utilizando un par de archivos: `go.mod` y `go.sum`, y una serie de comandos que permiten interactuar con estas dependencias.
