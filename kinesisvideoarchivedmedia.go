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
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia"
)

type KinesisVideoArchivedMediaType struct {
	service      *kinesisvideoarchivedmedia.KinesisVideoArchivedMedia
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

func KinesisVideoArchivedMediaFactory(cfg aws.Config) Factory {
	i := new(KinesisVideoArchivedMediaType)

	i.SetService(cfg)

	return i
}

func (i *KinesisVideoArchivedMediaType) SetPartialName() {
	// "AWS::KinesisVideoArchivedMedia::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::KinesisVideoArchivedMedia::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *KinesisVideoArchivedMediaType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsInput or ListVpcsInput" accordingly
	prefix := search[r{resource: "kinesisvideoarchivedmedia", kind: partialName}]
	name := fmt.Sprintf("%s%sInput", prefix, partialName)

	i.inputName = name
}

func (i *KinesisVideoArchivedMediaType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsOutput or ListVpcsOutput" accordingly
	prefix := search[r{resource: "kinesisvideoarchivedmedia", kind: partialName}]
	name := fmt.Sprintf("%s%sOutput", prefix, partialName)

	i.outputName = name
}

func (i *KinesisVideoArchivedMediaType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsRequest or ListVpcsRequest" accordingly
	prefix := search[r{resource: "kinesisvideoarchivedmedia", kind: partialName}]
	name := fmt.Sprintf("%s%sRequest", prefix, partialName)

	i.methodName = name
}

func (i *KinesisVideoArchivedMediaType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *KinesisVideoArchivedMediaType) Configure(param interface{}) error {
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

func (i *KinesisVideoArchivedMediaType) SetService(cfg aws.Config) {
	srv := kinesisvideoarchivedmedia.New(cfg)

	i.service = srv
}

func (i *KinesisVideoArchivedMediaType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("kinesisvideoarchivedmedia", i.inputName)
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

func (i *KinesisVideoArchivedMediaType) GetResources() {}

func (i *KinesisVideoArchivedMediaType) GetResourcesDetail() {}
