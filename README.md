# Go API client for swagger

No descripton provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Build date: 2017-09-11T15:52:17.214-04:00
- Build package: class io.swagger.codegen.languages.GoClientCodegen

## Installation

We are using following version of swagger:
```
version: 0.12.0
commit: 8135eb6728e43b73489e80f94426e6d387809502
```

Put the package under your project folder and add the following in import:
```
    "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://localhost/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApiApi* | [**AboutGet**](docs/ApiApi.md#aboutget) | **Get** /about | Get about
*UpdateSwitchApi* | [**UpdateSwitch**](docs/UpdateSwitchApi.md#updateswitch) | **Post** /updateSwitch | Update switch firmware


## Documentation For Models

 - [About](docs/About.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [ModelError](docs/ModelError.md)


## Documentation For Authorization
After generating the code we edited the file configure_on_network.go to disable authentication.
In configure_on_network.go - Generated Code:
	api.APIKeyHeaderAuth = func(token string) (interface{}, error) {
		return nil, errors.NotImplemented("api key auth (APIKeyHeader) authorization from header param [authorization] has not yet been implemented")
	}

In configure_on_network.go - Modified Code:
```
   api.APIKeyHeaderAuth = func(token string) (interface{}, error) {
		return token, nil
   }
```


## Documentation for Handlers

The handler is added in the controller folder. For example, the code  for the about api is in on-network/controllers/about/about.go

After generating the code, configure_on_nertwork.go had to be changed as well to call the handler code.
Here is an example to what had to be done for `about` api

In  on_network_api.go - Generated Code:
```
    api.AboutAboutGetHandler = about.AboutGetHandlerFunc(func(params about.AboutGetParams, principal interface{}) middleware.Responder {
    		return middleware.NotImplemented("operation about.AboutGet has not yet been implemented")
    	})
```

In  on_network_api.go - Modified Code:
```
    api.AboutAboutGetHandler = about.AboutGetHandlerFunc(func(params about.AboutGetParams, principal interface{}) middleware.Responder {
		return aboutctrl.MiddleWare(params.HTTPRequest)
	})
```


## APIKeyHeader

- **Type**: API key 
- **API key parameter name**: authorization
- **Location**: HTTP header


## Author



