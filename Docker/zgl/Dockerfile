# ARG python_image_v="python:3.11-slim"
# FROM ${python_image_v}
# USER root

# ARG work_dir="/src/"

# WORKDIR ${work_dir}

# ENV POETRY_HOME=/opt/poetry
# RUN apt-get update
# RUN apt-get install --no-install-recommends -y curl
# RUN curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python -
# RUN cd /usr/local/bin
# RUN ln -s /opt/poetry/bin/poetry
# ENV PATH /root/.poetry/bin:$PATH

FROM python:3.11-slim

WORKDIR /src

RUN apt update; apt -y upgrade
RUN pip install --upgrade pip

RUN pip install poetry
