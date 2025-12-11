# Comandos básicos para correr programas en Go 🐹

### Compilar paquetes y dependencias de nuestro programa en GO

Con esto generas el binario del programa
```bash
  go build hello_world.go
```
### Compilar y correr nuestro programa en GO
```bash
  go run hello_world.go
```
### Compilar nuestro programa en GO en cualquier SO
Se genera el .exe del SO especificado
```bash
  GOOS=windows GOARCH=arm64 go build hello_world.go
```
