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
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

type KMSType struct {
	service      *kms.KMS
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

func KMSFactory(cfg aws.Config) Factory {
	i := new(KMSType)

	i.SetService(cfg)

	return i
}

func (i *KMSType) SetPartialName() {
	// "AWS::KMS::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::KMS::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *KMSType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *KMSType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *KMSType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *KMSType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *KMSType) Configure(param interface{}) error {
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

func (i *KMSType) SetService(cfg aws.Config) {
	srv := kms.New(cfg)

	i.service = srv
}

func (i *KMSType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("kms", i.inputName)
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

func (i *KMSType) GetResources() {}

func (i *KMSType) GetResourcesDetail() {}
