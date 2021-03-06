// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	versioned3 "k8s.io/metrics/pkg/client/clientset/versioned"
	"reflect"
)

func AnyPtrToVersioned3Clientset() *versioned3.Clientset {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*versioned3.Clientset))(nil)).Elem()))
	var nullValue *versioned3.Clientset
	return nullValue
}

func EqPtrToVersioned3Clientset(value *versioned3.Clientset) *versioned3.Clientset {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *versioned3.Clientset
	return nullValue
}
