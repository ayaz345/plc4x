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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const BVLC_BACNETTYPE uint8 = 0x81

// BVLC is the corresponding interface of BVLC
type BVLC interface {
	utils.LengthAware
	utils.Serializable
	// GetBvlcFunction returns BvlcFunction (discriminator field)
	GetBvlcFunction() uint8
	// GetBvlcPayloadLength returns BvlcPayloadLength (virtual field)
	GetBvlcPayloadLength() uint16
}

// BVLCExactly can be used when we want exactly this type and not a type which fulfills BVLC.
// This is useful for switch cases.
type BVLCExactly interface {
	BVLC
	isBVLC() bool
}

// _BVLC is the data-structure of this message
type _BVLC struct {
	_BVLCChildRequirements
}

type _BVLCChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetBvlcFunction() uint8
}

type BVLCParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child BVLC, serializeChildFunction func() error) error
	GetTypeName() string
}

type BVLCChild interface {
	utils.Serializable
	InitializeParent(parent BVLC)
	GetParent() *BVLC

	GetTypeName() string
	BVLC
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BVLC) GetBvlcPayloadLength() uint16 {
	return uint16(uint16(uint16(m.GetLengthInBytes())) - uint16(uint16(4)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_BVLC) GetBacnetType() uint8 {
	return BVLC_BACNETTYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLC factory function for _BVLC
func NewBVLC() *_BVLC {
	return &_BVLC{}
}

// Deprecated: use the interface for direct cast
func CastBVLC(structType interface{}) BVLC {
	if casted, ok := structType.(BVLC); ok {
		return casted
	}
	if casted, ok := structType.(*BVLC); ok {
		return *casted
	}
	return nil
}

func (m *_BVLC) GetTypeName() string {
	return "BVLC"
}

func (m *_BVLC) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Const Field (bacnetType)
	lengthInBits += 8
	// Discriminator Field (bvlcFunction)
	lengthInBits += 8

	// Implicit Field (bvlcLength)
	lengthInBits += 16

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BVLC) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCParse(theBytes []byte) (BVLC, error) {
	return BVLCParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func BVLCParseWithBuffer(readBuffer utils.ReadBuffer) (BVLC, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLC"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLC")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (bacnetType)
	bacnetType, _bacnetTypeErr := readBuffer.ReadUint8("bacnetType", 8)
	if _bacnetTypeErr != nil {
		return nil, errors.Wrap(_bacnetTypeErr, "Error parsing 'bacnetType' field of BVLC")
	}
	if bacnetType != BVLC_BACNETTYPE {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BVLC_BACNETTYPE) + " but got " + fmt.Sprintf("%d", bacnetType))
	}

	// Discriminator Field (bvlcFunction) (Used as input to a switch field)
	bvlcFunction, _bvlcFunctionErr := readBuffer.ReadUint8("bvlcFunction", 8)
	if _bvlcFunctionErr != nil {
		return nil, errors.Wrap(_bvlcFunctionErr, "Error parsing 'bvlcFunction' field of BVLC")
	}

	// Implicit Field (bvlcLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	bvlcLength, _bvlcLengthErr := readBuffer.ReadUint16("bvlcLength", 16)
	_ = bvlcLength
	if _bvlcLengthErr != nil {
		return nil, errors.Wrap(_bvlcLengthErr, "Error parsing 'bvlcLength' field of BVLC")
	}

	// Virtual field
	_bvlcPayloadLength := uint16(bvlcLength) - uint16(uint16(4))
	bvlcPayloadLength := uint16(_bvlcPayloadLength)
	_ = bvlcPayloadLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BVLCChildSerializeRequirement interface {
		BVLC
		InitializeParent(BVLC)
		GetParent() BVLC
	}
	var _childTemp interface{}
	var _child BVLCChildSerializeRequirement
	var typeSwitchError error
	switch {
	case bvlcFunction == 0x00: // BVLCResult
		_childTemp, typeSwitchError = BVLCResultParseWithBuffer(readBuffer)
	case bvlcFunction == 0x01: // BVLCWriteBroadcastDistributionTable
		_childTemp, typeSwitchError = BVLCWriteBroadcastDistributionTableParseWithBuffer(readBuffer, bvlcPayloadLength)
	case bvlcFunction == 0x02: // BVLCReadBroadcastDistributionTable
		_childTemp, typeSwitchError = BVLCReadBroadcastDistributionTableParseWithBuffer(readBuffer)
	case bvlcFunction == 0x03: // BVLCReadBroadcastDistributionTableAck
		_childTemp, typeSwitchError = BVLCReadBroadcastDistributionTableAckParseWithBuffer(readBuffer, bvlcPayloadLength)
	case bvlcFunction == 0x04: // BVLCForwardedNPDU
		_childTemp, typeSwitchError = BVLCForwardedNPDUParseWithBuffer(readBuffer, bvlcPayloadLength)
	case bvlcFunction == 0x05: // BVLCRegisterForeignDevice
		_childTemp, typeSwitchError = BVLCRegisterForeignDeviceParseWithBuffer(readBuffer)
	case bvlcFunction == 0x06: // BVLCReadForeignDeviceTable
		_childTemp, typeSwitchError = BVLCReadForeignDeviceTableParseWithBuffer(readBuffer)
	case bvlcFunction == 0x07: // BVLCReadForeignDeviceTableAck
		_childTemp, typeSwitchError = BVLCReadForeignDeviceTableAckParseWithBuffer(readBuffer, bvlcPayloadLength)
	case bvlcFunction == 0x08: // BVLCDeleteForeignDeviceTableEntry
		_childTemp, typeSwitchError = BVLCDeleteForeignDeviceTableEntryParseWithBuffer(readBuffer)
	case bvlcFunction == 0x09: // BVLCDistributeBroadcastToNetwork
		_childTemp, typeSwitchError = BVLCDistributeBroadcastToNetworkParseWithBuffer(readBuffer, bvlcPayloadLength)
	case bvlcFunction == 0x0A: // BVLCOriginalUnicastNPDU
		_childTemp, typeSwitchError = BVLCOriginalUnicastNPDUParseWithBuffer(readBuffer, bvlcPayloadLength)
	case bvlcFunction == 0x0B: // BVLCOriginalBroadcastNPDU
		_childTemp, typeSwitchError = BVLCOriginalBroadcastNPDUParseWithBuffer(readBuffer, bvlcPayloadLength)
	case bvlcFunction == 0x0C: // BVLCSecureBVLL
		_childTemp, typeSwitchError = BVLCSecureBVLLParseWithBuffer(readBuffer, bvlcPayloadLength)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [bvlcFunction=%v]", bvlcFunction)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BVLC")
	}
	_child = _childTemp.(BVLCChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BVLC"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLC")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_BVLC) SerializeParent(writeBuffer utils.WriteBuffer, child BVLC, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BVLC"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BVLC")
	}

	// Const Field (bacnetType)
	_bacnetTypeErr := writeBuffer.WriteUint8("bacnetType", 8, 0x81)
	if _bacnetTypeErr != nil {
		return errors.Wrap(_bacnetTypeErr, "Error serializing 'bacnetType' field")
	}

	// Discriminator Field (bvlcFunction) (Used as input to a switch field)
	bvlcFunction := uint8(child.GetBvlcFunction())
	_bvlcFunctionErr := writeBuffer.WriteUint8("bvlcFunction", 8, (bvlcFunction))

	if _bvlcFunctionErr != nil {
		return errors.Wrap(_bvlcFunctionErr, "Error serializing 'bvlcFunction' field")
	}

	// Implicit Field (bvlcLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	bvlcLength := uint16(uint16(m.GetLengthInBytes()))
	_bvlcLengthErr := writeBuffer.WriteUint16("bvlcLength", 16, (bvlcLength))
	if _bvlcLengthErr != nil {
		return errors.Wrap(_bvlcLengthErr, "Error serializing 'bvlcLength' field")
	}
	// Virtual field
	if _bvlcPayloadLengthErr := writeBuffer.WriteVirtual("bvlcPayloadLength", m.GetBvlcPayloadLength()); _bvlcPayloadLengthErr != nil {
		return errors.Wrap(_bvlcPayloadLengthErr, "Error serializing 'bvlcPayloadLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BVLC"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BVLC")
	}
	return nil
}

func (m *_BVLC) isBVLC() bool {
	return true
}

func (m *_BVLC) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
