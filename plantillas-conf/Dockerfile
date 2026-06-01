# ==========================================
# ETAPA 1: BUILDER (La cocina sucia y pesada)
# -> Aquí compilamos y descargamos dependencias.
# -> Elige SOLO UN bloque (A, B o C) y borra el resto.
# ==========================================

# --- OPCIÓN A: JAVA / SPRING BOOT ---
FROM maven:3.9-eclipse-temurin-21 AS builder
WORKDIR /build
COPY pom.xml .
RUN mvn dependency:go-offline -B
COPY src src
RUN mvn package -DskipTests

# --- OPCIÓN B: FRONTEND (Angular / React / Node) ---
# FROM node:20-alpine AS builder
# WORKDIR /app
# COPY package*.json ./
# RUN npm install
# COPY . .
# RUN npm run build

# --- OPCIÓN C: PYTHON (FastAPI / Flask) ---
# FROM python:3.12-slim AS builder
# WORKDIR /app
# COPY requirements.txt .
# RUN pip install --user --no-cache-dir -r requirements.txt
# COPY . .


# ==========================================
# ETAPA 2: RUNTIME (El contenedor final, limpio y ligero)
# -> Esta imagen será la que vaya a producción.
# ==========================================

# Cambia la imagen base según tu lenguaje:
# Para Java: eclipse-temurin:21-jre-jammy
# Para Python: python:3.12-slim
# Para Frontend: nginx:alpine
FROM eclipse-temurin:21-jre-jammy 

# 1. (Opcional) Instalar herramientas extra de infraestructura si usas el monolito
RUN apt-get update && apt-get install -y --no-install-recommends \
    nginx supervisor \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# 2. Configurar la Seguridad (¡Súper importante!)
# Creamos un usuario sin permisos de administrador (non-root) para correr la app.
# Esto evita que si alguien hackea el contenedor, pueda tomar control de tu PC.
RUN groupadd -r appgroup && useradd -r -g appgroup appuser

# 3. Traer SOLO los archivos útiles de la Etapa 1
# Descomenta el que aplique a tu proyecto:

# -> Para JAVA:
COPY --from=builder /build/target/*.jar app.jar
# -> Para NODE/FRONTEND (alojando en Nginx):
# COPY --from=builder /app/dist/tu-proyecto /var/www/html
# -> Para PYTHON:
# COPY --from=builder /root/.local /root/.local
# COPY --from=builder /app /app

# 4. Copiar archivos de configuración de infraestructura (si los usas)
# COPY nginx.conf /etc/nginx/nginx.conf
# COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# 5. Darle permisos al usuario seguro sobre nuestra carpeta
RUN chown -R appuser:appgroup /app

# 6. Cambiamos del usuario "root" al usuario seguro
USER appuser

# 7. Exponer el puerto que usa tu app (Ej: 80 para Nginx, 5000 para Flask, 8080 para Spring)
EXPOSE 8080

# 8. El comando de encendido
# Si usas Supervisord:
# CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]
# Si es directo (Ej: Java):
CMD ["java", "-jar", "app.jar"]

# ... (todo lo de arriba de la Etapa 2 queda igual) ...

# 6. Cambiamos del usuario "root" al usuario seguro
USER appuser

# ==========================================
# ¡EL VOLUMEN! (Para bases de datos o archivos subidos)
# -> Borra esta línea si tu app no necesita guardar datos permanentes.
# ==========================================
VOLUME ["/app/data"]

# 7. Exponer el puerto que usa tu app
EXPOSE 8080

# 8. El comando de encendido
CMD ["java", "-jar", "app.jar"]
