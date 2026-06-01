FROM python:3.12-slim

# Instalamos Nginx y Supervisor
RUN apt-get update && apt-get install -y --no-install-recommends \
    nginx supervisor \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Dependencias de Python
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copiamos el código de la API y el código del Worker
COPY api_principal.py .
COPY worker_procesamiento.py .

# Copiamos las configuraciones
COPY nginx.conf /etc/nginx/nginx.conf
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

EXPOSE 80 5000

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]
