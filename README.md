# Agenda_Personal_API
API en lenguaje Go con ayuda del framework Beego y base de datos relacional PostgreSQL sobre una Agenda Personal.
## Tecnologías Implementadas
- Golang
- Beego
- PostgreSQL
## Instalación
> Dirigirse a la carpeta de go en la que quiere clonar el proyecto 

`cd go/src`
> Clonar el repositorio

`git clone https://github.com/CrisEs2506/agenda_personal_crud.git`
> Entrar en la carpeta del proyecto clonado

`cd agenda_personal_crud`
> Instalar dependencias del proyecto

`go get`
> Crear archivo para las variables de entorno .env
```
export AGENDA_PERSONAL_PGUSER = [nombre del usuario de postgres]
export AGENDA_PERSONAL_PGPASS = [contraseña del usuario de postgres]
export AGENDA_PERSONAL_PGHOST = [nombre o número de host]
export AGENDA_PERSONAL_PGPORT = [puerto en que corre postgres]
export AGENDA_PERSONAL_PGDB = [nombre de la base de datos en donde se encuentra el schema]
export AGENDA_PERSONAL_PGSCHEMA = [nombre del schema]
```
## Ejecución del Proyecto
`bee run`
