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
	"github.com/aws/aws-sdk-go-v2/service/health"
)

type HealthType struct {
	service      *health.Health
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

func HealthFactory(cfg aws.Config) Factory {
	i := new(HealthType)

	i.SetService(cfg)

	return i
}

func (i *HealthType) SetPartialName() {
	// "AWS::Health::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::Health::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *HealthType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *HealthType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *HealthType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *HealthType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *HealthType) Configure(param interface{}) error {
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

func (i *HealthType) SetService(cfg aws.Config) {
	srv := health.New(cfg)

	i.service = srv
}

func (i *HealthType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("health", i.inputName)
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

func (i *HealthType) GetResources() {}

func (i *HealthType) GetResourcesDetail() {}
