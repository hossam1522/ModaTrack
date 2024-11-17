# Metodología para creación de tests

Para poder realizar tests de forma efectiva, es necesario seguir una metodología que nos permita asegurarnos de que estamos probando el código de la forma más eficiente posible. En este proyecto se ha decidido seguir la metodología de testeo TDD (Test-Driven Development), que consiste en escribir los tests antes de escribir el código que se quiere probar. De esta forma, se asegura que el código que se escribe cumple con los requisitos establecidos en los tests.

Los pasos a seguir para la creación de tests siguiendo la metodología TDD son los siguientes:

1. **Escoger un requisito**: Seleccionar un requisito del proyecto que se quiera implementar.
2. **Escribir un test que falle**: Escribir un test que verifique que el requisito seleccionado no está implementado.
3. **Escribir el código mínimo para que el test pase**: Implementar el código mínimo necesario para que el test pase.
4. **Ejecutar los tests**: Ejecutar todos los tests para asegurarse de que el nuevo test no ha roto nada.
5. **Refactorizar el código**: Si es necesario, refactorizar el código para mejorar su calidad.
6. **Actualizar la lista de requisitos**: Marcar el requisito seleccionado como implementado y seleccionar otro requisito para seguir con el ciclo, además de añadir nuevos requisitos si es necesario.

Siguiendo esta metodología, se asegura que el código que se escribe está correctamente probado y cumple con los requisitos establecidos en los tests. Además, se fomenta la creación de código de calidad y se reduce la probabilidad de introducir errores en el código.

Más información sobre la metodología TDD se puede encontrar en el siguiente [enlace](https://softwarecrafters.io/javascript/tdd-test-driven-development).