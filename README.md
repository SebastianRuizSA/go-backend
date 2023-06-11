
### Basics
Uno de los conceptos claves es dividir funcionabilidades por paquetes(packages).
main() es el punto de entrada de cualquier programa.

### Packages
strings: 

In Go, constants must have a value that can be determined at compile time. They cannot be assigned a value that depends on a runtime calculation or 
the result of a function call like len(result).

The len() function calculates the length of a map at runtime, which means the value it returns cannot be used to assign a constant. Constants in Go must be assigned a value that is known and can be evaluated by the compiler without executing the program.

The reason the order of the keys appears to change with every execution in your code is because maps in Go do not preserve the order of their key-value pairs. The iteration order of a map is not guaranteed to be the same across different runs or even within the same run.

