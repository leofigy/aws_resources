// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-05-09 11:54:14.958521953 -0500 CDT m=+0.000134534
package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mturk"
)

type MTurkType struct {
	service      *mturk.MTurk
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

type MTurkTypeConfig struct {
	resourceType string
}

func MTurkFactory(cfg aws.Config) Factory {
	i := new(MTurkType)

	i.SetService(cfg)

	return i
}

func (i *MTurkType) SetPartialName() {
	// "AWS::MTurk::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::MTurk::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *MTurkType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *MTurkType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *MTurkType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *MTurkType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *MTurkType) Configure(param interface{}) error {
	config, ok := param.(MTurkTypeConfig)
	if !ok {
		return errors.New("config is not a valid param (MTurkTypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *MTurkType) SetService(cfg aws.Config) {
	srv := mturk.New(cfg)

	i.service = srv
}

func (i *MTurkType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("mturk", i.inputName)
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

func (i *MTurkType) GetResources() {}

func (i *MTurkType) GetResourcesDetail() {}