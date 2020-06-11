// very bad nodejs code below...sorry...

const fs = require("fs");
const parseArgs = require("minimist");
const path = require("path");
const _ = require("lodash");
const grpc = require("grpc");
const protoLoader = require("@grpc/proto-loader");
const health = require("grpc-health-check");

let PROTO_PATH = __dirname + "/demo-bank.proto";
let packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});
let api = grpc.loadPackageDefinition(packageDefinition).demobank.api;

function dummyCasaAccount(call, callback) {
  let lastUpdate = {
    seconds: Math.floor(new Date(2020, 4, 31, 18, 19, 20, 0) / 1000 + 3600 * 8),
    nanos: 0,
  };
  let dummy = Object.assign(
    {
      account_id: "00000000",
      nickname: "dummy-v2",
      prod_code: "1111",
      prod_name: "Bottomless CASA",
      currency: "SGD",
      status: "DORMANT",
      status_last_updated: lastUpdate,
      balances: [
        {
          amount: 10.0,
          type: "CURRENT",
          credit_flag: true,
          last_updated: lastUpdate,
        },
        {
          amount: 10.0,
          type: "AVAILABLE",
          credit_flag: true,
          last_updated: lastUpdate,
        },
      ],
    },
    api.CasaAccount
  );

  console.log("returning dummy account data");
  
  callback(null, dummy);
}

function getServer() {
  let server = new grpc.Server();

  server.addService(
    health.service,
    new health.Implementation({
      "": proto.grpc.health.v1.HealthCheckResponse.ServingStatus.SERVING,
    })
  );

  server.addService(api.CasaAccountService.service, {
    GetAccount: dummyCasaAccount,
  });

  return server;
}

if (require.main === module) {
  let argv = parseArgs(process.argv, {
    string: "p",
  });
  port = argv.p !== undefined ? argv.p : "50051";

  let server = getServer();
  server.bind("0.0.0.0:" + port, grpc.ServerCredentials.createInsecure());
  server.start();
}

exports.getServer = getServer;
