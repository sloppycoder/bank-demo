FROM node:12-alpine 

ADD package*.json *.js *.proto bin/grpc_health_probe ./

RUN npm install

EXPOSE 50051

ENTRYPOINT [ "node", "casa-account-server.js" ]
