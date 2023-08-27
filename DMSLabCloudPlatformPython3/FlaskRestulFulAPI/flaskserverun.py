from FlaskRestulFulAPI import app,RestFulAPIServer

rfsapiserver = RestFulAPIServer("Documents/ConfigurationOfPlatform.yaml")

if __name__== "__main__":
    app.run(rfsapiserver.ipAddr, rfsapiserver.port, debug= True)
