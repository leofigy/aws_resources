// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-05-14 14:12:55.399703429 -0500 CDT m=+0.000132171
package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync"
)

type CognitoSyncType struct {
	service      *cognitosync.CognitoSync
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

func CognitoSyncFactory(cfg aws.Config) Factory {
	i := new(CognitoSyncType)

	i.SetService(cfg)

	return i
}

func (i *CognitoSyncType) SetPartialName() {
	// "AWS::CognitoSync::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::CognitoSync::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *CognitoSyncType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsInput or ListVpcsInput" accordingly
	prefix := search[r{resource: "cognitosync", kind: partialName}]
	name := fmt.Sprintf("%s%sInput", prefix, partialName)

	i.inputName = name
}

func (i *CognitoSyncType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsOutput or ListVpcsOutput" accordingly
	prefix := search[r{resource: "cognitosync", kind: partialName}]
	name := fmt.Sprintf("%s%sOutput", prefix, partialName)

	i.outputName = name
}

func (i *CognitoSyncType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsRequest or ListVpcsRequest" accordingly
	prefix := search[r{resource: "cognitosync", kind: partialName}]
	name := fmt.Sprintf("%s%sRequest", prefix, partialName)

	i.methodName = name
}

func (i *CognitoSyncType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *CognitoSyncType) Configure(param interface{}) error {
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

func (i *CognitoSyncType) SetService(cfg aws.Config) {
	srv := cognitosync.New(cfg)

	i.service = srv
}

func (i *CognitoSyncType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("cognitosync", i.inputName)
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

func (i *CognitoSyncType) GetResources() {}

func (i *CognitoSyncType) GetResourcesDetail() {}
