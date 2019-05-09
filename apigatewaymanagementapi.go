// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-05-09 13:34:28.060197985 -0500 CDT m=+0.000154309
package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
)

type ApiGatewayManagementApiType struct {
	service      *apigatewaymanagementapi.ApiGatewayManagementApi
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

func ApiGatewayManagementApiFactory(cfg aws.Config) Factory {
	i := new(ApiGatewayManagementApiType)

	i.SetService(cfg)

	return i
}

func (i *ApiGatewayManagementApiType) SetPartialName() {
	// "AWS::ApiGatewayManagementApi::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::ApiGatewayManagementApi::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *ApiGatewayManagementApiType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *ApiGatewayManagementApiType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *ApiGatewayManagementApiType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *ApiGatewayManagementApiType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *ApiGatewayManagementApiType) Configure(param interface{}) error {
	config, ok := param.(TypeConfig)
	if !ok {
		return errors.New("config is not a valid param (TypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *ApiGatewayManagementApiType) SetService(cfg aws.Config) {
	srv := apigatewaymanagementapi.New(cfg)

	i.service = srv
}

func (i *ApiGatewayManagementApiType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("apigatewaymanagementapi", i.inputName)
	if err != nil {
		// We can ignore this kind of errors because there is not resources by the
		// i.inputName
		//log.Println(err)
		return
	}

	method := reflect.ValueOf(i.service).MethodByName(i.methodName)
	called := method.Call([]reflect.Value{reflect.ValueOf(instance)})

	send := reflect.Indirect(called[0]).MethodByName("Send")
	calledSend := send.Call([]reflect.Value{})

	res := calledSend[0]

	fmt.Printf("%v\n", res)
}

func (i *ApiGatewayManagementApiType) GetResources() {}

func (i *ApiGatewayManagementApiType) GetResourcesDetail() {}
