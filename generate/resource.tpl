// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
package aws_resources

import (
	"errors"
	"fmt"
  "log"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/{{ ToLower .Resource }}"
)

type {{ .Resource }}Type struct {
	service      *{{ ToLower .Resource }}.{{ .Resource }}
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

func {{ .Resource }}Factory(cfg aws.Config) Factory {
	i := new({{ .Resource }}Type)

	i.SetService(cfg)

	return i
}

func (i *{{ .Resource }}Type) SetPartialName() {
	// "AWS::{{ .Resource }}::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::{{ .Resource }}::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *{{ .Resource }}Type) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsInput or ListVpcsInput" accordingly
	prefix := search[r{resource: "{{ ToLower .Resource }}", kind: partialName}]
	name := fmt.Sprintf("%s%sInput", prefix, partialName)

	i.inputName = name
}

func (i *{{ .Resource }}Type) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsOutput or ListVpcsOutput" accordingly
	prefix := search[r{resource: "{{ ToLower .Resource }}", kind: partialName}]
	name := fmt.Sprintf("%s%sOutput", prefix, partialName)

	i.outputName = name
}

func (i *{{ .Resource }}Type) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsRequest or ListVpcsRequest" accordingly
	prefix := search[r{resource: "{{ ToLower .Resource }}", kind: partialName}]
	name := fmt.Sprintf("%s%sRequest", prefix, partialName)

	i.methodName = name
}

func (i *{{ .Resource }}Type) SetResourceType(t string) {
	i.resourceType = t
}

func (i *{{ .Resource }}Type) Configure(param interface{}) error {
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

func (i *{{ .Resource }}Type) SetService(cfg aws.Config) {
	srv := {{ ToLower .Resource }}.New(cfg)

	i.service = srv
}

func (i *{{ .Resource }}Type) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("{{ ToLower .Resource }}", i.inputName)
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

func (i *{{ .Resource }}Type) GetResources() {}

func (i *{{ .Resource }}Type) GetResourcesDetail() {}
