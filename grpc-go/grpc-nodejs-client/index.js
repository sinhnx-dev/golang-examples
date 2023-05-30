var PROTO_PATH = __dirname + '/valid.proto';
var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');

var packageDefinition = protoLoader.loadSync(
  PROTO_PATH,
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
  });
var validService = grpc.loadPackageDefinition(packageDefinition).valid.ValidService;
var validClient= new validService("localhost:2011", grpc.credentials.createInsecure())

validClient.validEmail({data: "contact@sinhnx.dev"}, (err, recipe) => {
  if (err) {
    // process error
    console.error(err)
  } else {
    // process valid
    console.log(recipe)
  }
})