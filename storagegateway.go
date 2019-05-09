// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-05-09 13:01:20.611969172 -0500 CDT m=+0.000147082
package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
)

type StorageGatewayType struct {
	service      *storagegateway.StorageGateway
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

type StorageGatewayTypeConfig struct {
	resourceType string
}

func StorageGatewayFactory(cfg aws.Config) Factory {
	i := new(StorageGatewayType)

	i.SetService(cfg)

	return i
}

func (i *StorageGatewayType) SetPartialName() {
	// "AWS::StorageGateway::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::StorageGateway::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *StorageGatewayType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *StorageGatewayType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *StorageGatewayType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *StorageGatewayType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *StorageGatewayType) Configure(param interface{}) error {
	config, ok := param.(StorageGatewayTypeConfig)
	if !ok {
		return errors.New("config is not a valid param (StorageGatewayTypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *StorageGatewayType) SetService(cfg aws.Config) {
	srv := storagegateway.New(cfg)

	i.service = srv
}

func (i *StorageGatewayType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("storagegateway", i.inputName)
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

func (i *StorageGatewayType) GetResources() {}

func (i *StorageGatewayType) GetResourcesDetail() {}
