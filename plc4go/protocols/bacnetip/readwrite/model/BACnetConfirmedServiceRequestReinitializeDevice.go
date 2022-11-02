/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestReinitializeDevice is the corresponding interface of BACnetConfirmedServiceRequestReinitializeDevice
type BACnetConfirmedServiceRequestReinitializeDevice interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetReinitializedStateOfDevice returns ReinitializedStateOfDevice (property field)
	GetReinitializedStateOfDevice() BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
	// GetPassword returns Password (property field)
	GetPassword() BACnetContextTagCharacterString
}

// BACnetConfirmedServiceRequestReinitializeDeviceExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestReinitializeDevice.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestReinitializeDeviceExactly interface {
	BACnetConfirmedServiceRequestReinitializeDevice
	isBACnetConfirmedServiceRequestReinitializeDevice() bool
}

// _BACnetConfirmedServiceRequestReinitializeDevice is the data-structure of this message
type _BACnetConfirmedServiceRequestReinitializeDevice struct {
	*_BACnetConfirmedServiceRequest
	ReinitializedStateOfDevice BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
	Password                   BACnetContextTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetReinitializedStateOfDevice() BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged {
	return m.ReinitializedStateOfDevice
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetPassword() BACnetContextTagCharacterString {
	return m.Password
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestReinitializeDevice factory function for _BACnetConfirmedServiceRequestReinitializeDevice
func NewBACnetConfirmedServiceRequestReinitializeDevice(reinitializedStateOfDevice BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged, password BACnetContextTagCharacterString, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestReinitializeDevice {
	_result := &_BACnetConfirmedServiceRequestReinitializeDevice{
		ReinitializedStateOfDevice:     reinitializedStateOfDevice,
		Password:                       password,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReinitializeDevice(structType interface{}) BACnetConfirmedServiceRequestReinitializeDevice {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReinitializeDevice); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReinitializeDevice); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReinitializeDevice"
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (reinitializedStateOfDevice)
	lengthInBits += m.ReinitializedStateOfDevice.GetLengthInBits()

	// Optional Field (password)
	if m.Password != nil {
		lengthInBits += m.Password.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestReinitializeDeviceParse(theBytes []byte, serviceRequestLength uint32) (BACnetConfirmedServiceRequestReinitializeDevice, error) {
	return BACnetConfirmedServiceRequestReinitializeDeviceParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), serviceRequestLength) // TODO: get endianness from mspec
}

func BACnetConfirmedServiceRequestReinitializeDeviceParseWithBuffer(readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestReinitializeDevice, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReinitializeDevice"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReinitializeDevice")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (reinitializedStateOfDevice)
	if pullErr := readBuffer.PullContext("reinitializedStateOfDevice"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for reinitializedStateOfDevice")
	}
	_reinitializedStateOfDevice, _reinitializedStateOfDeviceErr := BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedParseWithBuffer(readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _reinitializedStateOfDeviceErr != nil {
		return nil, errors.Wrap(_reinitializedStateOfDeviceErr, "Error parsing 'reinitializedStateOfDevice' field of BACnetConfirmedServiceRequestReinitializeDevice")
	}
	reinitializedStateOfDevice := _reinitializedStateOfDevice.(BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged)
	if closeErr := readBuffer.CloseContext("reinitializedStateOfDevice"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for reinitializedStateOfDevice")
	}

	// Optional Field (password) (Can be skipped, if a given expression evaluates to false)
	var password BACnetContextTagCharacterString = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("password"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for password")
		}
		_val, _err := BACnetContextTagParseWithBuffer(readBuffer, uint8(1), BACnetDataType_CHARACTER_STRING)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'password' field of BACnetConfirmedServiceRequestReinitializeDevice")
		default:
			password = _val.(BACnetContextTagCharacterString)
			if closeErr := readBuffer.CloseContext("password"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for password")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReinitializeDevice"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReinitializeDevice")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestReinitializeDevice{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		ReinitializedStateOfDevice: reinitializedStateOfDevice,
		Password:                   password,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReinitializeDevice"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReinitializeDevice")
		}

		// Simple Field (reinitializedStateOfDevice)
		if pushErr := writeBuffer.PushContext("reinitializedStateOfDevice"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for reinitializedStateOfDevice")
		}
		_reinitializedStateOfDeviceErr := writeBuffer.WriteSerializable(m.GetReinitializedStateOfDevice())
		if popErr := writeBuffer.PopContext("reinitializedStateOfDevice"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for reinitializedStateOfDevice")
		}
		if _reinitializedStateOfDeviceErr != nil {
			return errors.Wrap(_reinitializedStateOfDeviceErr, "Error serializing 'reinitializedStateOfDevice' field")
		}

		// Optional Field (password) (Can be skipped, if the value is null)
		var password BACnetContextTagCharacterString = nil
		if m.GetPassword() != nil {
			if pushErr := writeBuffer.PushContext("password"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for password")
			}
			password = m.GetPassword()
			_passwordErr := writeBuffer.WriteSerializable(password)
			if popErr := writeBuffer.PopContext("password"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for password")
			}
			if _passwordErr != nil {
				return errors.Wrap(_passwordErr, "Error serializing 'password' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReinitializeDevice"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReinitializeDevice")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) isBACnetConfirmedServiceRequestReinitializeDevice() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
