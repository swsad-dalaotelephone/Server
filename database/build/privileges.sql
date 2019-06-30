use mysql;
-- new user
set password for root@localhost = password('SYSU_baobaozhuan2019');
-- important
grant all on *.* to root@'%' identified by 'SYSU_baobaozhuan2019' with grant option;
-- use privileges
flush privileges;