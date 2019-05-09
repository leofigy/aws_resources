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
	"github.com/aws/aws-sdk-go-v2/service/opsworks"
)

type OpsWorksType struct {
	service      *opsworks.OpsWorks
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

func OpsWorksFactory(cfg aws.Config) Factory {
	i := new(OpsWorksType)

	i.SetService(cfg)

	return i
}

func (i *OpsWorksType) SetPartialName() {
	// "AWS::OpsWorks::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::OpsWorks::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *OpsWorksType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *OpsWorksType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *OpsWorksType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *OpsWorksType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *OpsWorksType) Configure(param interface{}) error {
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

func (i *OpsWorksType) SetService(cfg aws.Config) {
	srv := opsworks.New(cfg)

	i.service = srv
}

func (i *OpsWorksType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("opsworks", i.inputName)
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

func (i *OpsWorksType) GetResources() {}

func (i *OpsWorksType) GetResourcesDetail() {}
