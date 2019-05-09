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
	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
)

type DeviceFarmType struct {
	service      *devicefarm.DeviceFarm
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

type DeviceFarmTypeConfig struct {
	resourceType string
}

func DeviceFarmFactory(cfg aws.Config) Factory {
	i := new(DeviceFarmType)

	i.SetService(cfg)

	return i
}

func (i *DeviceFarmType) SetPartialName() {
	// "AWS::DeviceFarm::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::DeviceFarm::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *DeviceFarmType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *DeviceFarmType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *DeviceFarmType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *DeviceFarmType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *DeviceFarmType) Configure(param interface{}) error {
	config, ok := param.(DeviceFarmTypeConfig)
	if !ok {
		return errors.New("config is not a valid param (DeviceFarmTypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *DeviceFarmType) SetService(cfg aws.Config) {
	srv := devicefarm.New(cfg)

	i.service = srv
}

func (i *DeviceFarmType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("devicefarm", i.inputName)
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

func (i *DeviceFarmType) GetResources() {}

func (i *DeviceFarmType) GetResourcesDetail() {}
