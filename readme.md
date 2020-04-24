# Taxky - Aprende Jugando

### Una Nueva Manera De Aprender
Con Taxky tus hijos pueden aprende el significado de la responsabilidad, todo esto mediante tareas diarias con puntuación, estas tareas le darán puntos los cuales podrá intercambiar por lo que ellos deseen.

Servidor encargado de proveer información mediante servicios Rest

Tecnologias implementadas:

- Go lang
- Gorm
- Mysql
- Echo

# Instrucciones

Primero tendrás que clonar el repositorio a tu computadora con el siguiente comando dentro del directorio `$GOPATH/src/github.com/`.

`git clone  https://github.com/neto-dev/taxky_server.git`

Una vez descargado ingresar a la carpeta.

`cd taxky_server`

Para instalar las diferentes dependencias que se utilizan en el proyecto ejecutar el siguiente comando.

`go mod init`

En caso de que arrojé el siguiente error.

`go: modules disabled inside GOPATH/src by GO111MODULE=auto; see 'go help modules'`

Ejecutar el siguiente comando.  

`export GO111MODULE=on`

Seguido de:

`go mod init`

Despues ejecutamos:

`go mod vendor`

Listo en este punto ya contaremos con las dependencias descargadas en el directorio vendor y con esto ya podremos ejecutar nuestro proyecto.


Seguido crearemos el archivo variables.env en el cual crearemos las siguientes variables de entorno.

```
export DATABASE_USER='root'
export DATABASE_PASS='pass_db'
export DATABASE_HOST=''
export ENVIRONMENT='Development'
```

Estas ultimas son las que el sistema utilizara para conectar con la base de datos y para setear el entorno de producción.

En seguida ejecutamos el comando:

`source ./variables.env`

Esto para setear nuestras variables de entorno que usaremos en el proyecto.

Teniendo una vez esto ya podremos ejecutar el proyecto con el comando.

`npm run main.go`

Y listo las migraciones y toda la configuración se ejecutara automáticamente al correr el servicio.


### Autor
@neto-dev Ernesto Valenzuela, apasionado a la programación así como de las tecnologías open source, amante de los lenguajes Ruby, PHP, Python y Go, puedes contactarme directamente a correo hello@netov.dev o bien seguirme en twitter [@neto_dev](https://twitter.com/neto_dev) | Más: [Web](https://netov.dev)