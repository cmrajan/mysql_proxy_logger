mysql_proxy_logger
===================

Simple query logging mysql proxy written in Go 

updated code with many new features

has some sensible defaults, but is also configurable 

go build mysql_proxy_logger.go


./mysql_proxy_logger --help

  -P="3306": mysqlserverPort
  -h="127.0.0.1": mysqlserverIP
  -p="3307": localport


scenario 1
mysql, proxy and app are on same machine



mysql runs on 127.0.0.1 and connects at 3306
proxy runs at 3307 and connects to the proxy
app is configured to use the proxy at 127.0.0.1 with port 3307 with existing mysql username password


run
./mysql_proxy_logger

scenario 2
mysql is in separate box, proxy and app are on same box

mysql runs on seperate box, can be called from outside with  202.54.202.54 and accepts connection at 3306
run the proxy with host and port set to use the mysql server info
run the app with app hostname and port at 127.0.0.1 and 3307, mysql username and password remains same

./mysql_proxy_logger -h 202.54.202.54 -P 3307



scenario 3

mysql and proxy are on same box, app server is in a separate box

mysql server runs as usual
proxy runs as usual, without any change in config
app uses the mysql username, password and hostname as such, only the port number should be changed to 3307


run
./mysql_proxy_logger


scenario 4

need to run multiple proxies, 

for example one for tracking the application, one for tracking the running of the background task

./mysql_proxy_logger -p 3307  # first proxy
./mysql_proxy_logger -p 3308  # second proxy



logging to file

./myprox  2>&1 > sql.log



logging to file and showing it on stdout

./myprox  2>&1 | tee sql.log

myprox
======
original code which I forked from https://github.com/acsellers/myprox




# Install

install golang

## OSX

brew install golang

## Ubuntu

sudo apt-get install golang



git clone https://github.com/senthilnayagam/myprox

cd myprox


# Usage

runs the mysql proxy at 127.0.0.1 at port 3307

go run mysql_proxy_logger.go -p 3307

 configure it in your database.yml

# test it

 mysql -uroot -H 127.0.0.1 -P3307 -p


when you run your queries now it passes thorough the proxy and will show all queries in the proxy console

 ctrl+c will kill the proxy


log the queries to a file

go run mysql_proxy_logger.go -p 3307 2>&1 | tee rails_sql.log


# build it

go build mysql_proxy_logger.go

it will be compiled and a binary added to current folder

now you can use it like

./mysql_proxy_logger -p 3307 2>&1 | tee rails_sql.log



# usage

it is not designed for production usage, it is for debugging and testing purpose


if you need mysql/mariadb load balancer/failover/reverse proxy for production use Haproxy




# To Do

had a use case where the db was on a different server, so need to configure it

add timestamp to the logs to make the trace easier when paired with the web server log

