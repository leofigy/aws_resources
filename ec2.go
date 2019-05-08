package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type EC2Type struct {
	service      *ec2.EC2
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

type EC2TypeConfig struct {
	resourceType string
}

func EC2Factory(cfg aws.Config) Factory {
	i := new(EC2Type)

	i.SetService(cfg)

	return i
}

func (i *EC2Type) SetPartialName() {
	// "AWS::EC2::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::EC2::", "")

	if name == "VPC" {
		// "VPC" to "vpc"
		name = strings.ToLower(name)

		// "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *EC2Type) SetInputName() {
	if i.partialName == "" {
		return
	}

	// "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *EC2Type) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *EC2Type) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *EC2Type) SetResourceType(t string) {
	i.resourceType = t
}

func (i *EC2Type) Configure(param interface{}) error {
	config, ok := param.(EC2TypeConfig)
	if !ok {
		return errors.New("config is not a valid param (EC2TypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *EC2Type) SetService(cfg aws.Config) {
	srv := ec2.New(cfg)

	i.service = srv
}

func (i *EC2Type) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("ec2", i.inputName)
	if err != nil {
		// We can ignore this kind of errors because there is not resources bu the
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

func (i *EC2Type) GetResources() {}

func (i *EC2Type) GetResourcesDetail() {}
