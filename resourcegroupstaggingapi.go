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
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
)

type ResourceGroupsTaggingAPIType struct {
	service      *resourcegroupstaggingapi.ResourceGroupsTaggingAPI
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

func ResourceGroupsTaggingAPIFactory(cfg aws.Config) Factory {
	i := new(ResourceGroupsTaggingAPIType)

	i.SetService(cfg)

	return i
}

func (i *ResourceGroupsTaggingAPIType) SetPartialName() {
	// "AWS::ResourceGroupsTaggingAPI::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::ResourceGroupsTaggingAPI::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *ResourceGroupsTaggingAPIType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *ResourceGroupsTaggingAPIType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *ResourceGroupsTaggingAPIType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *ResourceGroupsTaggingAPIType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *ResourceGroupsTaggingAPIType) Configure(param interface{}) error {
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

func (i *ResourceGroupsTaggingAPIType) SetService(cfg aws.Config) {
	srv := resourcegroupstaggingapi.New(cfg)

	i.service = srv
}

func (i *ResourceGroupsTaggingAPIType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("resourcegroupstaggingapi", i.inputName)
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

func (i *ResourceGroupsTaggingAPIType) GetResources() {}

func (i *ResourceGroupsTaggingAPIType) GetResourcesDetail() {}
