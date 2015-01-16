from debian:jessie

ADD editor /editor

EXPOSE 3000

ENTRYPOINT /editor/app