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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// AdsReadDeviceInfoResponse is the data-structure of this message
type AdsReadDeviceInfoResponse struct {
	*AdsData
	Result       ReturnCode
	MajorVersion uint8
	MinorVersion uint8
	Version      uint16
	Device       []byte
}

// IAdsReadDeviceInfoResponse is the corresponding interface of AdsReadDeviceInfoResponse
type IAdsReadDeviceInfoResponse interface {
	IAdsData
	// GetResult returns Result (property field)
	GetResult() ReturnCode
	// GetMajorVersion returns MajorVersion (property field)
	GetMajorVersion() uint8
	// GetMinorVersion returns MinorVersion (property field)
	GetMinorVersion() uint8
	// GetVersion returns Version (property field)
	GetVersion() uint16
	// GetDevice returns Device (property field)
	GetDevice() []byte
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

func (m *AdsReadDeviceInfoResponse) GetCommandId() CommandId {
	return CommandId_ADS_READ_DEVICE_INFO
}

func (m *AdsReadDeviceInfoResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *AdsReadDeviceInfoResponse) InitializeParent(parent *AdsData) {}

func (m *AdsReadDeviceInfoResponse) GetParent() *AdsData {
	return m.AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *AdsReadDeviceInfoResponse) GetResult() ReturnCode {
	return m.Result
}

func (m *AdsReadDeviceInfoResponse) GetMajorVersion() uint8 {
	return m.MajorVersion
}

func (m *AdsReadDeviceInfoResponse) GetMinorVersion() uint8 {
	return m.MinorVersion
}

func (m *AdsReadDeviceInfoResponse) GetVersion() uint16 {
	return m.Version
}

func (m *AdsReadDeviceInfoResponse) GetDevice() []byte {
	return m.Device
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsReadDeviceInfoResponse factory function for AdsReadDeviceInfoResponse
func NewAdsReadDeviceInfoResponse(result ReturnCode, majorVersion uint8, minorVersion uint8, version uint16, device []byte) *AdsReadDeviceInfoResponse {
	_result := &AdsReadDeviceInfoResponse{
		Result:       result,
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		Version:      version,
		Device:       device,
		AdsData:      NewAdsData(),
	}
	_result.Child = _result
	return _result
}

func CastAdsReadDeviceInfoResponse(structType interface{}) *AdsReadDeviceInfoResponse {
	if casted, ok := structType.(AdsReadDeviceInfoResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*AdsReadDeviceInfoResponse); ok {
		return casted
	}
	if casted, ok := structType.(AdsData); ok {
		return CastAdsReadDeviceInfoResponse(casted.Child)
	}
	if casted, ok := structType.(*AdsData); ok {
		return CastAdsReadDeviceInfoResponse(casted.Child)
	}
	return nil
}

func (m *AdsReadDeviceInfoResponse) GetTypeName() string {
	return "AdsReadDeviceInfoResponse"
}

func (m *AdsReadDeviceInfoResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AdsReadDeviceInfoResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (result)
	lengthInBits += 32

	// Simple field (majorVersion)
	lengthInBits += 8

	// Simple field (minorVersion)
	lengthInBits += 8

	// Simple field (version)
	lengthInBits += 16

	// Array field
	if len(m.Device) > 0 {
		lengthInBits += 8 * uint16(len(m.Device))
	}

	return lengthInBits
}

func (m *AdsReadDeviceInfoResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsReadDeviceInfoResponseParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (*AdsReadDeviceInfoResponse, error) {
	if pullErr := readBuffer.PullContext("AdsReadDeviceInfoResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (result)
	if pullErr := readBuffer.PullContext("result"); pullErr != nil {
		return nil, pullErr
	}
	_result, _resultErr := ReturnCodeParse(readBuffer)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field")
	}
	result := _result
	if closeErr := readBuffer.CloseContext("result"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (majorVersion)
	_majorVersion, _majorVersionErr := readBuffer.ReadUint8("majorVersion", 8)
	if _majorVersionErr != nil {
		return nil, errors.Wrap(_majorVersionErr, "Error parsing 'majorVersion' field")
	}
	majorVersion := _majorVersion

	// Simple Field (minorVersion)
	_minorVersion, _minorVersionErr := readBuffer.ReadUint8("minorVersion", 8)
	if _minorVersionErr != nil {
		return nil, errors.Wrap(_minorVersionErr, "Error parsing 'minorVersion' field")
	}
	minorVersion := _minorVersion

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadUint16("version", 16)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field")
	}
	version := _version
	// Byte Array field (device)
	numberOfBytesdevice := int(uint16(16))
	device, _readArrayErr := readBuffer.ReadByteArray("device", numberOfBytesdevice)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'device' field")
	}

	if closeErr := readBuffer.CloseContext("AdsReadDeviceInfoResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsReadDeviceInfoResponse{
		Result:       result,
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		Version:      version,
		Device:       device,
		AdsData:      &AdsData{},
	}
	_child.AdsData.Child = _child
	return _child, nil
}

func (m *AdsReadDeviceInfoResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadDeviceInfoResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (result)
		if pushErr := writeBuffer.PushContext("result"); pushErr != nil {
			return pushErr
		}
		_resultErr := m.Result.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("result"); popErr != nil {
			return popErr
		}
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		// Simple Field (majorVersion)
		majorVersion := uint8(m.MajorVersion)
		_majorVersionErr := writeBuffer.WriteUint8("majorVersion", 8, (majorVersion))
		if _majorVersionErr != nil {
			return errors.Wrap(_majorVersionErr, "Error serializing 'majorVersion' field")
		}

		// Simple Field (minorVersion)
		minorVersion := uint8(m.MinorVersion)
		_minorVersionErr := writeBuffer.WriteUint8("minorVersion", 8, (minorVersion))
		if _minorVersionErr != nil {
			return errors.Wrap(_minorVersionErr, "Error serializing 'minorVersion' field")
		}

		// Simple Field (version)
		version := uint16(m.Version)
		_versionErr := writeBuffer.WriteUint16("version", 16, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		// Array Field (device)
		if m.Device != nil {
			// Byte Array field (device)
			_writeArrayErr := writeBuffer.WriteByteArray("device", m.Device)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'device' field")
			}
		}

		if popErr := writeBuffer.PopContext("AdsReadDeviceInfoResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsReadDeviceInfoResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
