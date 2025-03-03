# Roadmap - Golang Projects: task tracker

El proyecto nace siguiendo las indicaciones de [roadmap.sh/projects/task-tracker](https://roadmap.sh/projects/task-tracker).

El requerimiento fue implementar un sistema de seguimiento de tareas en la consola, utilizando únicamente la biblioteca estándar de Go.

## Bitácora

Durante el desarrollo del proyecto aprendí y experimenté con varias herramientas de la biblioteca estándar de Go. Empecé utilizando `fmt` para estructurar los mensajes en consola e incorporé `os`  de a poco para gestionar la creación de directorios y archivos JSON, lo cual era requerimiento para la persistencia de datos.

Dentro de este proceso, `strings.Builder` se convirtió en un recurso nuevo que me permitió construir el *output* de forma flexible y con el formato que necesitaba.

Avanzado el proyecto, decidí utilizar `log` para el manejo de errores, modificando las partes que tenían un manejo incompleto y reduciendo la deuda técnica que tuve al comienzo para iterar rápido en el funcionamiento.

Cada una de estas herramientas (y otras no mencionadas) me ayudaron a ir puliendo el código y a entender mejor cómo integrarlas.

## Documentación

### Instalación

1. Clona el repositorio:

```bash
git clone https://github.com/neecosanudo/roadmap-go-task-tracker.git
```

2. Accede al directorio del proyecto:

```bash
cd roadmap-go-task-tracker
```

3. Compila el proyecto para generar el binario:

```bash
go build -o task-cli main.go
```

### Uso del Proyecto

El proyecto se ejecuta desde la consola utilizando el binario generado (por ejemplo, task-cli).

Algunos comandos disponibles son:

```bash
# Agregar una nueva tarea
task-cli add "Buy groceries"

# Actualizar y eliminar una tarea
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marcar una tarea como "en progreso" o "hecha"
task-cli mark-in-progress 1
task-cli mark-done 1

# Listar todas las tareas
task-cli list

# Listar todas las tareas según un solo estado
task-cli list done
task-cli list todo
task-cli list in-progress
```

### Notas Adicionales

El proyecto utiliza un archivo JSON ubicado en la carpeta `./data` para almacenar la lista de tareas.

Si la carpeta o el archivo no existen, el programa los creará automáticamente al iniciarse.

Asegúrate de tener instalado Go y configurado el entorno de desarrollo para compilar y ejecutar aplicaciones en Go.
