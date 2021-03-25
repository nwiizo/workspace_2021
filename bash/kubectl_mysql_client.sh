kubectl run -it --rm --image=mysql:5.6 --restart=Never mysql-client -n ${DB_NAMESPACE} -- mysqldump -u ${DB_USER} -p${DB_PASSWORD} -h ${DB_HOSTNAME} ${DB_NAME} --single-transaction
