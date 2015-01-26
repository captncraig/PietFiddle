from debian:jessie

ADD website /website
WORKDIR /website

ENV PORT=3000

EXPOSE 3000

CMD /website/app
