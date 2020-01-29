FROM alpine 

LABEL vendor=cloudstone
LABEL product=pandas

ADD certs/* /etc/ssl/pandas
ADD bin/* /
RUN chmod 755 /main
# RUN chmod 755 /dockerize

USER 6342627

CMD ["/main"]
