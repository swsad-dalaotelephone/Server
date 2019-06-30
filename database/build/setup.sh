#!/bin/bash
set -e

echo '1. starting mysql...'
service mysql start

echo '2. creating database...'
mysql < /mysql/schema.sql

# set password
echo '3. setting password....'
mysql < /mysql/privileges.sql

# check mysql status
echo `service mysql status`
echo '4. mysql for baobaozhuan is ready...'

tail -f /dev/null