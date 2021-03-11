package reflection

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Query struct {
	Service  string
	Method   string
	Request  string
	Response string
}

func (q Query) String() string {
	return fmt.Sprintf("service: %s method: %s request: %s response: %s", q.Service, q.Method, q.Request, q.Response)
}

type Deliverable struct {
	MsgName string
}

func (d Deliverable) String() string {
	return fmt.Sprintf("deliverable: %s", d.MsgName)
}

type ServiceMsg struct {
	Method     string
	Descriptor protoreflect.MessageDescriptor
}

type QueryDescriptor struct {
	Request     protoreflect.MessageDescriptor
	Response    protoreflect.MessageDescriptor
	ServiceName string
}

type methodsMap map[string]QueryDescriptor

func (m methodsMap) insert(methodName string, method protoreflect.MethodDescriptor) error {

	if _, exists := m[methodName]; exists {
		return fmt.Errorf("method already exists: %s", methodName)
	}
	if method.IsStreamingClient() || method.IsStreamingServer() {
		return fmt.Errorf("streaming rpcs are not supported: %s", methodName)
	}
	m[methodName] = QueryDescriptor{
		Request:     method.Input(),
		Response:    method.Output(),
		ServiceName: "todo", // TODO
	}

	return nil
}

func (m methodsMap) merge(m2 methodsMap) error {
	for k, v := range m2 {
		if _, exists := m[k]; exists {
			return fmt.Errorf("method already exists: %s", k)
		}
		m[k] = v
	}

	return nil
}