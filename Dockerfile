from debian:jessie

ADD . /editor

EXPOSE 3000

ENTRYPOINT /editor/app