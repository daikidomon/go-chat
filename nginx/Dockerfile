FROM nginx:1.19
ADD conf /etc/nginx/conf.d
ADD uwsgi_params /etc/nginx/uwsgi_params

ARG TZ
ENV TZ ${TZ}
RUN echo ${TZ}
RUN echo ${TZ} > /etc/timezone \
   && dpkg-reconfigure -f noninteractive tzdata

RUN ln -sf /dev/stdout /var/log/nginx/access.log \
	&& ln -sf /dev/stdout /var/log/nginx/error.log