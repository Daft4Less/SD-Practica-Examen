FROM maven:3.9-eclipse-temurin-21 AS builder
WORKDIR /build
COPY pom.xml .
RUN mvn dependency:go-offline -B
COPY src src
RUN mvn package -DskipTest
RUN find /root/.m2 -name "h2-*.jar" | head -1 | xargs -I{} cp {} /tmp/h2-server.jar
FROM eclipse-temurin:21-jre-jammy 
RUN apt-get update && apt-get install -y --no-install-recommends \
    nginx supervisor && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /build/target/inventario-1.0.0.jar app.jar
COPY --from=builder /tmp/h2-server.jar h2-server.jar
COPY nginx.conf /etc/nginx/nginx.conf
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

VOLUME ["/app/data"]
EXPOSE 80 5000 9092
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]