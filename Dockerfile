FROM registry.cn-shanghai.aliyuncs.com/pigeonligh/build:golang-1.16-alpine AS build

WORKDIR /build

ADD go.* ./

RUN go mod download

ADD cmd ./cmd
ADD pkg ./pkg
ADD Makefile ./

RUN make

FROM panubo/sshd:latest

COPY --from=build /build/_output/bin/ssh-board /bin/
COPY --from=build /build/_output/bin/add /bin/

ENV SSH_ENABLE_ROOT false
ENV SSH_ENABLE_PASSWORD_AUTH false
ENV DISABLE_SFTP true
ENV DISABLE_SCP true
ENV DISABLE_RSYNC true
ENV SSH_USERS noname:1000:1000:/bin/ssh-board

RUN echo "" >> /etc/authorized_keys/noname
RUN echo "PermitUserEnvironment yes" >> /etc/ssh/sshd_config 

ADD start.sh /bin/

CMD ["sh", "/bin/start.sh"]
