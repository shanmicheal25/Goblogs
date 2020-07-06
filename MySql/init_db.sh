
#!/bin/bash
/usr/bin/mysqld_safe --skip-grant-tables &
  sleep 5
  mysql -u root -e "CREATE DATABASE blogsdb"
  mysql -u root blogsdb < /docker-entrypoint-initdb.d/test.sql