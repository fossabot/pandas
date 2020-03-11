FROM ubuntu:18.04 

LABEL vendor=cloustone
LABEL product=pandas

ADD certs/* /etc/ssl/pandas
ADD bin/* /
RUN chmod 755 pandas-shiro
# RUN chmod 755 /dockerize

USER root

CMD ["/pandas-shiro"]
