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
	"github.com/aws/aws-sdk-go-v2/service/codestar"
)

type CodeStarType struct {
	service      *codestar.CodeStar
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

type CodeStarTypeConfig struct {
	resourceType string
}

func CodeStarFactory(cfg aws.Config) Factory {
	i := new(CodeStarType)

	i.SetService(cfg)

	return i
}

func (i *CodeStarType) SetPartialName() {
	// "AWS::CodeStar::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::CodeStar::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *CodeStarType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *CodeStarType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *CodeStarType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *CodeStarType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *CodeStarType) Configure(param interface{}) error {
	config, ok := param.(CodeStarTypeConfig)
	if !ok {
		return errors.New("config is not a valid param (CodeStarTypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *CodeStarType) SetService(cfg aws.Config) {
	srv := codestar.New(cfg)

	i.service = srv
}

func (i *CodeStarType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("codestar", i.inputName)
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

func (i *CodeStarType) GetResources() {}

func (i *CodeStarType) GetResourcesDetail() {}
