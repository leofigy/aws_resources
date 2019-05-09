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
	"github.com/aws/aws-sdk-go-v2/service/amplify"
)

type AmplifyType struct {
	service      *amplify.Amplify
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

type AmplifyTypeConfig struct {
	resourceType string
}

func AmplifyFactory(cfg aws.Config) Factory {
	i := new(AmplifyType)

	i.SetService(cfg)

	return i
}

func (i *AmplifyType) SetPartialName() {
	// "AWS::Amplify::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::Amplify::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *AmplifyType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *AmplifyType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *AmplifyType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *AmplifyType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *AmplifyType) Configure(param interface{}) error {
	config, ok := param.(AmplifyTypeConfig)
	if !ok {
		return errors.New("config is not a valid param (AmplifyTypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *AmplifyType) SetService(cfg aws.Config) {
	srv := amplify.New(cfg)

	i.service = srv
}

func (i *AmplifyType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("amplify", i.inputName)
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

func (i *AmplifyType) GetResources() {}

func (i *AmplifyType) GetResourcesDetail() {}
