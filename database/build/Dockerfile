FROM mysql:5.7.25

# allow no password
ENV MYSQL_ALLOW_EMPTY_PASSWORD yes

# # set root password
# ENV MYSQL_ROOT_PASSWORD SYSU_baobaozhuan2019

# copy file into container
COPY setup.sh /mysql/setup.sh
COPY schema.sql /mysql/schema.sql
COPY privileges.sql /mysql/privileges.sql

# exec these command when container start up
CMD ["sh", "/mysql/setup.sh"]