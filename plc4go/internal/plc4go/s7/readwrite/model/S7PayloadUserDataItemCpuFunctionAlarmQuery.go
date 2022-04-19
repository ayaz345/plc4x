/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const S7PayloadUserDataItemCpuFunctionAlarmQuery_FUNCTIONID uint8 = 0x00
const S7PayloadUserDataItemCpuFunctionAlarmQuery_NUMBERMESSAGEOBJ uint8 = 0x01
const S7PayloadUserDataItemCpuFunctionAlarmQuery_VARIABLESPEC uint8 = 0x12
const S7PayloadUserDataItemCpuFunctionAlarmQuery_LENGTH uint8 = 0x08

// S7PayloadUserDataItemCpuFunctionAlarmQuery is the data-structure of this message
type S7PayloadUserDataItemCpuFunctionAlarmQuery struct {
	*S7PayloadUserDataItem
	SyntaxId  SyntaxIdType
	QueryType QueryType
	AlarmType AlarmType
}

// IS7PayloadUserDataItemCpuFunctionAlarmQuery is the corresponding interface of S7PayloadUserDataItemCpuFunctionAlarmQuery
type IS7PayloadUserDataItemCpuFunctionAlarmQuery interface {
	IS7PayloadUserDataItem
	// GetSyntaxId returns SyntaxId (property field)
	GetSyntaxId() SyntaxIdType
	// GetQueryType returns QueryType (property field)
	GetQueryType() QueryType
	// GetAlarmType returns AlarmType (property field)
	GetAlarmType() AlarmType
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) GetCpuFunctionType() uint8 {
	return 0x04
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) GetCpuSubfunction() uint8 {
	return 0x13
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) GetDataLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.S7PayloadUserDataItem.ReturnCode = returnCode
	m.S7PayloadUserDataItem.TransportSize = transportSize
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) GetParent() *S7PayloadUserDataItem {
	return m.S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) GetSyntaxId() SyntaxIdType {
	return m.SyntaxId
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) GetQueryType() QueryType {
	return m.QueryType
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) GetAlarmType() AlarmType {
	return m.AlarmType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) GetFunctionId() uint8 {
	return S7PayloadUserDataItemCpuFunctionAlarmQuery_FUNCTIONID
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) GetNumberMessageObj() uint8 {
	return S7PayloadUserDataItemCpuFunctionAlarmQuery_NUMBERMESSAGEOBJ
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) GetVariableSpec() uint8 {
	return S7PayloadUserDataItemCpuFunctionAlarmQuery_VARIABLESPEC
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) GetLength() uint8 {
	return S7PayloadUserDataItemCpuFunctionAlarmQuery_LENGTH
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemCpuFunctionAlarmQuery factory function for S7PayloadUserDataItemCpuFunctionAlarmQuery
func NewS7PayloadUserDataItemCpuFunctionAlarmQuery(syntaxId SyntaxIdType, queryType QueryType, alarmType AlarmType, returnCode DataTransportErrorCode, transportSize DataTransportSize) *S7PayloadUserDataItemCpuFunctionAlarmQuery {
	_result := &S7PayloadUserDataItemCpuFunctionAlarmQuery{
		SyntaxId:              syntaxId,
		QueryType:             queryType,
		AlarmType:             alarmType,
		S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	_result.Child = _result
	return _result
}

func CastS7PayloadUserDataItemCpuFunctionAlarmQuery(structType interface{}) *S7PayloadUserDataItemCpuFunctionAlarmQuery {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionAlarmQuery); ok {
		return &casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionAlarmQuery); ok {
		return casted
	}
	if casted, ok := structType.(S7PayloadUserDataItem); ok {
		return CastS7PayloadUserDataItemCpuFunctionAlarmQuery(casted.Child)
	}
	if casted, ok := structType.(*S7PayloadUserDataItem); ok {
		return CastS7PayloadUserDataItemCpuFunctionAlarmQuery(casted.Child)
	}
	return nil
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionAlarmQuery"
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Const Field (functionId)
	lengthInBits += 8

	// Const Field (numberMessageObj)
	lengthInBits += 8

	// Const Field (variableSpec)
	lengthInBits += 8

	// Const Field (length)
	lengthInBits += 8

	// Simple field (syntaxId)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (queryType)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (alarmType)
	lengthInBits += 8

	return lengthInBits
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadUserDataItemCpuFunctionAlarmQueryParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (*S7PayloadUserDataItemCpuFunctionAlarmQuery, error) {
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionAlarmQuery"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Const Field (functionId)
	functionId, _functionIdErr := readBuffer.ReadUint8("functionId", 8)
	if _functionIdErr != nil {
		return nil, errors.Wrap(_functionIdErr, "Error parsing 'functionId' field")
	}
	if functionId != S7PayloadUserDataItemCpuFunctionAlarmQuery_FUNCTIONID {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7PayloadUserDataItemCpuFunctionAlarmQuery_FUNCTIONID) + " but got " + fmt.Sprintf("%d", functionId))
	}

	// Const Field (numberMessageObj)
	numberMessageObj, _numberMessageObjErr := readBuffer.ReadUint8("numberMessageObj", 8)
	if _numberMessageObjErr != nil {
		return nil, errors.Wrap(_numberMessageObjErr, "Error parsing 'numberMessageObj' field")
	}
	if numberMessageObj != S7PayloadUserDataItemCpuFunctionAlarmQuery_NUMBERMESSAGEOBJ {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7PayloadUserDataItemCpuFunctionAlarmQuery_NUMBERMESSAGEOBJ) + " but got " + fmt.Sprintf("%d", numberMessageObj))
	}

	// Const Field (variableSpec)
	variableSpec, _variableSpecErr := readBuffer.ReadUint8("variableSpec", 8)
	if _variableSpecErr != nil {
		return nil, errors.Wrap(_variableSpecErr, "Error parsing 'variableSpec' field")
	}
	if variableSpec != S7PayloadUserDataItemCpuFunctionAlarmQuery_VARIABLESPEC {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7PayloadUserDataItemCpuFunctionAlarmQuery_VARIABLESPEC) + " but got " + fmt.Sprintf("%d", variableSpec))
	}

	// Const Field (length)
	length, _lengthErr := readBuffer.ReadUint8("length", 8)
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}
	if length != S7PayloadUserDataItemCpuFunctionAlarmQuery_LENGTH {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7PayloadUserDataItemCpuFunctionAlarmQuery_LENGTH) + " but got " + fmt.Sprintf("%d", length))
	}

	// Simple Field (syntaxId)
	if pullErr := readBuffer.PullContext("syntaxId"); pullErr != nil {
		return nil, pullErr
	}
	_syntaxId, _syntaxIdErr := SyntaxIdTypeParse(readBuffer)
	if _syntaxIdErr != nil {
		return nil, errors.Wrap(_syntaxIdErr, "Error parsing 'syntaxId' field")
	}
	syntaxId := _syntaxId
	if closeErr := readBuffer.CloseContext("syntaxId"); closeErr != nil {
		return nil, closeErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (queryType)
	if pullErr := readBuffer.PullContext("queryType"); pullErr != nil {
		return nil, pullErr
	}
	_queryType, _queryTypeErr := QueryTypeParse(readBuffer)
	if _queryTypeErr != nil {
		return nil, errors.Wrap(_queryTypeErr, "Error parsing 'queryType' field")
	}
	queryType := _queryType
	if closeErr := readBuffer.CloseContext("queryType"); closeErr != nil {
		return nil, closeErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x34) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x34),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (alarmType)
	if pullErr := readBuffer.PullContext("alarmType"); pullErr != nil {
		return nil, pullErr
	}
	_alarmType, _alarmTypeErr := AlarmTypeParse(readBuffer)
	if _alarmTypeErr != nil {
		return nil, errors.Wrap(_alarmTypeErr, "Error parsing 'alarmType' field")
	}
	alarmType := _alarmType
	if closeErr := readBuffer.CloseContext("alarmType"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionAlarmQuery"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadUserDataItemCpuFunctionAlarmQuery{
		SyntaxId:              syntaxId,
		QueryType:             queryType,
		AlarmType:             alarmType,
		S7PayloadUserDataItem: &S7PayloadUserDataItem{},
	}
	_child.S7PayloadUserDataItem.Child = _child
	return _child, nil
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionAlarmQuery"); pushErr != nil {
			return pushErr
		}

		// Const Field (functionId)
		_functionIdErr := writeBuffer.WriteUint8("functionId", 8, 0x00)
		if _functionIdErr != nil {
			return errors.Wrap(_functionIdErr, "Error serializing 'functionId' field")
		}

		// Const Field (numberMessageObj)
		_numberMessageObjErr := writeBuffer.WriteUint8("numberMessageObj", 8, 0x01)
		if _numberMessageObjErr != nil {
			return errors.Wrap(_numberMessageObjErr, "Error serializing 'numberMessageObj' field")
		}

		// Const Field (variableSpec)
		_variableSpecErr := writeBuffer.WriteUint8("variableSpec", 8, 0x12)
		if _variableSpecErr != nil {
			return errors.Wrap(_variableSpecErr, "Error serializing 'variableSpec' field")
		}

		// Const Field (length)
		_lengthErr := writeBuffer.WriteUint8("length", 8, 0x08)
		if _lengthErr != nil {
			return errors.Wrap(_lengthErr, "Error serializing 'length' field")
		}

		// Simple Field (syntaxId)
		if pushErr := writeBuffer.PushContext("syntaxId"); pushErr != nil {
			return pushErr
		}
		_syntaxIdErr := m.SyntaxId.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("syntaxId"); popErr != nil {
			return popErr
		}
		if _syntaxIdErr != nil {
			return errors.Wrap(_syntaxIdErr, "Error serializing 'syntaxId' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (queryType)
		if pushErr := writeBuffer.PushContext("queryType"); pushErr != nil {
			return pushErr
		}
		_queryTypeErr := m.QueryType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("queryType"); popErr != nil {
			return popErr
		}
		if _queryTypeErr != nil {
			return errors.Wrap(_queryTypeErr, "Error serializing 'queryType' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x34))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (alarmType)
		if pushErr := writeBuffer.PushContext("alarmType"); pushErr != nil {
			return pushErr
		}
		_alarmTypeErr := m.AlarmType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("alarmType"); popErr != nil {
			return popErr
		}
		if _alarmTypeErr != nil {
			return errors.Wrap(_alarmTypeErr, "Error serializing 'alarmType' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionAlarmQuery"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmQuery) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
