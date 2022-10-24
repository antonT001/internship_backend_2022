#!/bin/bash

if [ ! "$(ls -A /var/lib/mysql)" ]; then
    /usr/bin/mysql_install_db
fi

/etc/init.d/mysql start
/usr/bin/mysqladmin -u root -h `hostname` password 'htDjPlfaOyD3x1wh'

mysql -u root -phtDjPlfaOyD3x1wh -e "CREATE DATABASE user_balance CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"
mysql -u root -phtDjPlfaOyD3x1wh -e "CREATE USER 'user_balance'@'%' IDENTIFIED BY '1QcD6VOqlbJntYDe'"
mysql -u root -phtDjPlfaOyD3x1wh -e "GRANT ALL PRIVILEGES ON user_balance . * TO 'user_balance'@'%'"
mysql -u root -phtDjPlfaOyD3x1wh -e "FLUSH PRIVILEGES"

while true ; do sleep 5; done;