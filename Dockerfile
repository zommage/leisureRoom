FROM centos:7
ARG config 
ADD ./logs /home/leisureRoom/logs
ADD ./conf /home/leisureRoom/conf
ADD ./leisureRoom /home/leisureRoom/leisureRoom
RUN ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
WORKDIR /home/leisureRoom

RUN mv ${config} ./conf/active.conf

VOLUME ["/home/leisureRoom/logs"]

CMD ["/home/leisureRoom/leisureRoom", "--config=./conf/active.conf"]
