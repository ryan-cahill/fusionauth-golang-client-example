FROM postgres:9.6

RUN  apt-get update \
  && apt-get install -y wget unzip

# https://stackoverflow.com/questions/34751814/build-postgres-docker-container-with-initial-schema
# https://fusionauth.io/docs/v1/tech/installation-guide/upgrade
RUN wget https://files.fusionauth.io/products/fusionauth/1.18.4/fusionauth-database-schema-1.18.4.zip && \
  mkdir fusionauth_migrations && unzip fusionauth-database-schema-1.18.4.zip -d fusionauth_migrations/ && \
  cp fusionauth_migrations/postgresql.sql /docker-entrypoint-initdb.d/0-base-migration.sql && \
  cp fusionauth_migrations/migrations/postgresql/1.18.2.sql /docker-entrypoint-initdb.d/1-1.18.2.sql
