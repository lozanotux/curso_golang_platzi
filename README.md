# Curso de GOLANG

Existen dos formas de ejecutar los ejemplos de este repositorios, para fines prácticos, utilice el siguiente comando:
```bash
$ go run src/1_vars_prints.go
```

## Paquetes

Para inicializar un paquete, debe crearse un archivo `go.mod` (de modulo) que contenga el nombre del paquete y la version de golang que se va a utilizar.

Una forma rapida de generarlo es utilizando el comando:
```bash
$ go mod init
```

## Gestor de paquetes

Para instalar un paquete o dependencia en nuestro proyecto, debemos usar el siguiente comando:
```bash
$ go get -v -u golang.org/x/tour
```

Ya viene incluido en una instalación de `golang` y para encontrar paquetes externos, podemos utilizar [awesome-go.com](awesome-go.com).

Para verificar que todas nuestras dependencias se encuentran sanas, podemos ejecutar el siguiente comando:
```bash
$ go mod verify
```