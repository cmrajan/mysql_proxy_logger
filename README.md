myprox
======

Simple query logging mysql proxy written in Go 



log to a file

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

