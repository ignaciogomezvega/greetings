# Saludos en Go

Este paquete proporciona una forma simple de obtener saludos personalizados en Go.

## Instalación

Ejecuta el siguiente comando para instalar el paquete:

```bash
go get .u github.com/ignaciogomezvega/greetings
```

## Uso
Aqui tienes un ejemplo de como utilizar el paquete en tu codigo:

```go
package main

import (
    "fmt"
    "github.com/ignaciogomezvega/greetings"
)

func main() {
    message, err := greetings.Hello("Alex")

    if err != nil {
        fmt.Println("Ocurrió un error:", err)
        return
    }

    fmt.Println(message)
}

```
Este ejemplo importa el paquete github.com/ignaciogomezvega/greetings y llama a la función Hello() saludo personalizado. Si ocurre un error, se imprime un mensaje de error.