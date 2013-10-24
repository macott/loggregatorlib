package testhelpers

import (
	"github.com/cloudfoundry/loggregatorlib/logmessage"
	"time"
	"testing"
	"code.google.com/p/gogoprotobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/cloudfoundry/loggregatorlib/signature"
)

func MarshalledErrorLogMessage(t *testing.T, messageString string, appId string) []byte {
	currentTime := time.Now()

	messageType := logmessage.LogMessage_ERR
	sourceType := logmessage.LogMessage_DEA
	protoMessage := &logmessage.LogMessage{
		Message:     []byte(messageString),
		AppId:       proto.String(appId),
		MessageType: &messageType,
		SourceType:  &sourceType,
		Timestamp:   proto.Int64(currentTime.UnixNano()),
	}

	message, err := proto.Marshal(protoMessage)
	assert.NoError(t, err)

	return message
}

func MarshalledLogMessage(t *testing.T, messageString string, appId string) []byte {
	currentTime := time.Now()

	messageType := logmessage.LogMessage_OUT
	sourceType := logmessage.LogMessage_DEA
	protoMessage := &logmessage.LogMessage{
		Message:     []byte(messageString),
		AppId:       proto.String(appId),
		MessageType: &messageType,
		SourceType:  &sourceType,
		Timestamp:   proto.Int64(currentTime.UnixNano()),
	}

	message, err := proto.Marshal(protoMessage)
	assert.NoError(t, err)

	return message
}

func MarshalledDrainedLogMessage(t *testing.T, messageString string, appId string, drainUrls ...string) []byte {
	currentTime := time.Now()

	messageType := logmessage.LogMessage_OUT
	sourceType := logmessage.LogMessage_WARDEN_CONTAINER
	protoMessage := &logmessage.LogMessage{
		Message:     []byte(messageString),
		AppId:       proto.String(appId),
		MessageType: &messageType,
		SourceType:  &sourceType,
		DrainUrls:   drainUrls,
		Timestamp:   proto.Int64(currentTime.UnixNano()),
	}

	message, err := proto.Marshal(protoMessage)
	assert.NoError(t, err)

	return message
}

func MarshalledDrainedNonWardenLogMessage(t *testing.T, messageString string, appId string, drainUrls ...string) []byte {
	currentTime := time.Now()

	messageType := logmessage.LogMessage_OUT
	sourceType := logmessage.LogMessage_DEA
	protoMessage := &logmessage.LogMessage{
		Message:     []byte(messageString),
		AppId:       proto.String(appId),
		MessageType: &messageType,
		SourceType:  &sourceType,
		DrainUrls:   drainUrls,
		Timestamp:   proto.Int64(currentTime.UnixNano()),
	}

	message, err := proto.Marshal(protoMessage)
	assert.NoError(t, err)

	return message
}


func NewLogMessage(messageString, appId string) *logmessage.LogMessage {
	currentTime := time.Now()

	messageType := logmessage.LogMessage_OUT
	sourceType := logmessage.LogMessage_WARDEN_CONTAINER
	protoMessage := &logmessage.LogMessage{
		Message:     []byte(messageString),
		AppId:       proto.String(appId),
		MessageType: &messageType,
		SourceType:  &sourceType,
		Timestamp:   proto.Int64(currentTime.UnixNano()),
	}
	return protoMessage
}


func MarshalledLogEnvelope(t *testing.T, unmarshalledMessage *logmessage.LogMessage, secret string) []byte {
	signatureOfMessage, err := signature.Encrypt(secret, signature.Digest(unmarshalledMessage.String()))
	assert.NoError(t, err)

	envelope := &logmessage.LogEnvelope{
		LogMessage: unmarshalledMessage,
		RoutingKey: proto.String(*unmarshalledMessage.AppId),
		Signature:  signatureOfMessage,
	}

	marshalledEnvelope, err := proto.Marshal(envelope)
	assert.NoError(t, err)

	return marshalledEnvelope
}

func AssertProtoBufferMessageEquals(t *testing.T, expectedMessage string, actual []byte) {
	receivedMessage := &logmessage.LogMessage{}

	err := proto.Unmarshal(actual, receivedMessage)
	assert.NoError(t, err)
	assert.Equal(t, expectedMessage, string(receivedMessage.GetMessage()))
}
