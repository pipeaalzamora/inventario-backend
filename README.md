# Smart Order

<p align="center">
  <img src="https://app.smartorders.inventario.dotsolutions.cl/assets/BoxLogo-Dar0SZ5Z.png" alt="Smart Order" width="200"/>
</p>



**Smart Order** es una aplicación en desarrollo creada por [dotsolution.io](https://dotsolutions.io/), diseñada para optimizar la gestión de inventario, usuarios, perfiles y notificaciones dentro de un ecosistema empresarial.  


## 📋 Requisitos previos

Antes de comenzar asegúrate de tener instalado:

- [Go](https://go.dev/dl/) >= 1.22
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- (Opcional, para desarrollo) [Air](https://github.com/air-verse/air)

## 🚀 Setup del Proyecto

1. Clonar el repositorio.  
2. Copiar el archivo de configuración y renombrarlo:

```bash
cp config/config.example.json config/config.json
```
Levantar la aplicación con:

```bash
docker compose up -d
```
Para reiniciar la base de datos y aplicar cambios en el schema:

```bash
docker compose down -v
docker compose up -d
```

## 🖥️ Ejecución de la aplicación
La aplicación está desarrollada en Golang y utiliza Air como herramienta de live-reloading.

### 🔹 Con Air (modo recomendado)
```bash
air
```
🔄 ¿Qué hace Air?
Air es una herramienta para desarrollo en Go que detecta cambios en el código y recompila/reinicia automáticamente la aplicación, lo que permite un flujo de trabajo mucho más rápido y eficiente sin tener que detener y volver a correr manualmente el proyecto.

### 🔹 Sin Air (modo normal)
```bash
go run main.go
```
