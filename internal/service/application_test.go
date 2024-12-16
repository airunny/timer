package service

import (
	"fmt"
	"testing"

	"github.com/airunny/wiki-go-tools/objectid"
)

func TestService_GenApplicationSecret(t *testing.T) {
	s := new(Service)
	fmt.Println(s.generateApplicationSecret(objectid.ObjectID()))
}
