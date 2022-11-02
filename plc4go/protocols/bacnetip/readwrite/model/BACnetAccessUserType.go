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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetAccessUserType is an enum
type BACnetAccessUserType uint16

type IBACnetAccessUserType interface {
	utils.Serializable
}

const (
	BACnetAccessUserType_ASSET                    BACnetAccessUserType = 0
	BACnetAccessUserType_GROUP                    BACnetAccessUserType = 1
	BACnetAccessUserType_PERSON                   BACnetAccessUserType = 2
	BACnetAccessUserType_VENDOR_PROPRIETARY_VALUE BACnetAccessUserType = 0xFFFF
)

var BACnetAccessUserTypeValues []BACnetAccessUserType

func init() {
	_ = errors.New
	BACnetAccessUserTypeValues = []BACnetAccessUserType{
		BACnetAccessUserType_ASSET,
		BACnetAccessUserType_GROUP,
		BACnetAccessUserType_PERSON,
		BACnetAccessUserType_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetAccessUserTypeByValue(value uint16) (enum BACnetAccessUserType, ok bool) {
	switch value {
	case 0:
		return BACnetAccessUserType_ASSET, true
	case 0xFFFF:
		return BACnetAccessUserType_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetAccessUserType_GROUP, true
	case 2:
		return BACnetAccessUserType_PERSON, true
	}
	return 0, false
}

func BACnetAccessUserTypeByName(value string) (enum BACnetAccessUserType, ok bool) {
	switch value {
	case "ASSET":
		return BACnetAccessUserType_ASSET, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetAccessUserType_VENDOR_PROPRIETARY_VALUE, true
	case "GROUP":
		return BACnetAccessUserType_GROUP, true
	case "PERSON":
		return BACnetAccessUserType_PERSON, true
	}
	return 0, false
}

func BACnetAccessUserTypeKnows(value uint16) bool {
	for _, typeValue := range BACnetAccessUserTypeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetAccessUserType(structType interface{}) BACnetAccessUserType {
	castFunc := func(typ interface{}) BACnetAccessUserType {
		if sBACnetAccessUserType, ok := typ.(BACnetAccessUserType); ok {
			return sBACnetAccessUserType
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAccessUserType) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetAccessUserType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAccessUserTypeParse(theBytes []byte) (BACnetAccessUserType, error) {
	return BACnetAccessUserTypeParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func BACnetAccessUserTypeParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetAccessUserType, error) {
	val, err := readBuffer.ReadUint16("BACnetAccessUserType", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetAccessUserType")
	}
	if enum, ok := BACnetAccessUserTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetAccessUserType(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetAccessUserType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetAccessUserType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetAccessUserType", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAccessUserType) PLC4XEnumName() string {
	switch e {
	case BACnetAccessUserType_ASSET:
		return "ASSET"
	case BACnetAccessUserType_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetAccessUserType_GROUP:
		return "GROUP"
	case BACnetAccessUserType_PERSON:
		return "PERSON"
	}
	return ""
}

func (e BACnetAccessUserType) String() string {
	return e.PLC4XEnumName()
}
