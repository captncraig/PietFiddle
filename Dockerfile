from debian:jessie

ADD editor /editor
WORKDIR /editor

ENV PORT=3000

EXPOSE 3000

CMD /editor/app
