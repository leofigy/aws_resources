// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-06-02 13:56:03.549666 -0500 CDT m=+0.005104325
package aws_resources

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
)

type CodePipelineType struct {
	service      *codepipeline.CodePipeline
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

func CodePipelineFactory(cfg aws.Config) Factory {
	i := new(CodePipelineType)

	i.SetService(cfg)

	return i
}

func (i *CodePipelineType) SetPartialName() {
	// "AWS::CodePipeline::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::CodePipeline::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *CodePipelineType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsInput or ListVpcsInput" accordingly
	prefix := search[r{resource: "codepipeline", kind: partialName}]
	name := fmt.Sprintf("%s%sInput", prefix, partialName)

	i.inputName = name
}

func (i *CodePipelineType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsOutput or ListVpcsOutput" accordingly
	prefix := search[r{resource: "codepipeline", kind: partialName}]
	name := fmt.Sprintf("%s%sOutput", prefix, partialName)

	i.outputName = name
}

func (i *CodePipelineType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsRequest or ListVpcsRequest" accordingly
	prefix := search[r{resource: "codepipeline", kind: partialName}]
	name := fmt.Sprintf("%s%sRequest", prefix, partialName)

	i.methodName = name
}

func (i *CodePipelineType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *CodePipelineType) Configure(param interface{}) error {
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

func (i *CodePipelineType) SetService(cfg aws.Config) {
	srv := codepipeline.New(cfg)

	i.service = srv
}

func (i *CodePipelineType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("codepipeline", i.inputName)
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

func (i *CodePipelineType) GetResources() {}

func (i *CodePipelineType) GetResourcesDetail() {}
