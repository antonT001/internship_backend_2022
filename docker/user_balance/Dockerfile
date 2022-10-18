FROM ubuntu:20.04

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get -y update
RUN apt-get install -y tzdata git vim wget mc

RUN rm /etc/localtime
RUN ln -s /usr/share/zoneinfo/Europe/Moscow /etc/localtime

RUN mkdir -p /var/www/user_balance/

RUN wget https://golang.org/dl/go1.18.6.linux-amd64.tar.gz -P /tmp

COPY ./docker/user_balance/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT /usr/local/bin/entrypoint.sh
