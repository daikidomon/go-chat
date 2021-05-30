FROM python:3.8
ENV PYTHONUNBUFFERED 1

RUN mkdir -p /src
WORKDIR /src
ADD requirements.txt /src/
ADD uwsgi.ini /src/
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

WORKDIR /src
# CMD ["uwsgi", "--ini", "/src/uwsgi.ini"]
