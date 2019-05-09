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
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

type ECRType struct {
	service      *ecr.ECR
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

type ECRTypeConfig struct {
	resourceType string
}

func ECRFactory(cfg aws.Config) Factory {
	i := new(ECRType)

	i.SetService(cfg)

	return i
}

func (i *ECRType) SetPartialName() {
	// "AWS::ECR::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::ECR::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *ECRType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *ECRType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *ECRType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *ECRType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *ECRType) Configure(param interface{}) error {
	config, ok := param.(ECRTypeConfig)
	if !ok {
		return errors.New("config is not a valid param (ECRTypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *ECRType) SetService(cfg aws.Config) {
	srv := ecr.New(cfg)

	i.service = srv
}

func (i *ECRType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("ecr", i.inputName)
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

func (i *ECRType) GetResources() {}

func (i *ECRType) GetResourcesDetail() {}
