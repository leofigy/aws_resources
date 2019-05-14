// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-05-13 18:14:08.716266173 -0500 CDT m=+0.000161905
package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

type APIGatewayType struct {
	service      *apigateway.APIGateway
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

func APIGatewayFactory(cfg aws.Config) Factory {
	i := new(APIGatewayType)

	i.SetService(cfg)

	return i
}

func (i *APIGatewayType) SetPartialName() {
	// "AWS::APIGateway::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::APIGateway::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *APIGatewayType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *APIGatewayType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *APIGatewayType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *APIGatewayType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *APIGatewayType) Configure(param interface{}) error {
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

func (i *APIGatewayType) SetService(cfg aws.Config) {
	srv := apigateway.New(cfg)

	i.service = srv
}

func (i *APIGatewayType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("apigateway", i.inputName)
	if err != nil {
		// We can ignore this kind of errors because there is not resources by the
		// i.inputName
		log.Println(err)
		return
	}

	method := reflect.ValueOf(i.service).MethodByName(i.methodName)
	called := method.Call([]reflect.Value{reflect.ValueOf(instance)})

	send := reflect.Indirect(called[0]).MethodByName("Send")
	calledSend := send.Call([]reflect.Value{})

	res := calledSend[0]

	fmt.Printf("%v\n", res)
}

func (i *APIGatewayType) GetResources() {}

func (i *APIGatewayType) GetResourcesDetail() {}
