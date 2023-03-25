FROM node:18-alpine3.16

RUN mkdir -p /opt
COPY . /opt/
WORKDIR /opt

RUN npm i -g npm
RUN npm install 
RUN npm run build

CMD npm run start:dev