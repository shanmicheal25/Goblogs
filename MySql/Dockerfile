FROM mysql:5.7
RUN chown -R mysql:root /var/lib/mysql/
ADD test.sql /etc/mysql/test.sql
RUN cp /etc/mysql/test.sql /docker-entrypoint-initdb.d
