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

// BACnetProgramRequest is an enum
type BACnetProgramRequest uint8

type IBACnetProgramRequest interface {
	utils.Serializable
}

const (
	BACnetProgramRequest_READY   BACnetProgramRequest = 0
	BACnetProgramRequest_LOAD    BACnetProgramRequest = 1
	BACnetProgramRequest_RUN     BACnetProgramRequest = 2
	BACnetProgramRequest_HALT    BACnetProgramRequest = 3
	BACnetProgramRequest_RESTART BACnetProgramRequest = 4
	BACnetProgramRequest_UNLOAD  BACnetProgramRequest = 5
)

var BACnetProgramRequestValues []BACnetProgramRequest

func init() {
	_ = errors.New
	BACnetProgramRequestValues = []BACnetProgramRequest{
		BACnetProgramRequest_READY,
		BACnetProgramRequest_LOAD,
		BACnetProgramRequest_RUN,
		BACnetProgramRequest_HALT,
		BACnetProgramRequest_RESTART,
		BACnetProgramRequest_UNLOAD,
	}
}

func BACnetProgramRequestByValue(value uint8) (enum BACnetProgramRequest, ok bool) {
	switch value {
	case 0:
		return BACnetProgramRequest_READY, true
	case 1:
		return BACnetProgramRequest_LOAD, true
	case 2:
		return BACnetProgramRequest_RUN, true
	case 3:
		return BACnetProgramRequest_HALT, true
	case 4:
		return BACnetProgramRequest_RESTART, true
	case 5:
		return BACnetProgramRequest_UNLOAD, true
	}
	return 0, false
}

func BACnetProgramRequestByName(value string) (enum BACnetProgramRequest, ok bool) {
	switch value {
	case "READY":
		return BACnetProgramRequest_READY, true
	case "LOAD":
		return BACnetProgramRequest_LOAD, true
	case "RUN":
		return BACnetProgramRequest_RUN, true
	case "HALT":
		return BACnetProgramRequest_HALT, true
	case "RESTART":
		return BACnetProgramRequest_RESTART, true
	case "UNLOAD":
		return BACnetProgramRequest_UNLOAD, true
	}
	return 0, false
}

func BACnetProgramRequestKnows(value uint8) bool {
	for _, typeValue := range BACnetProgramRequestValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetProgramRequest(structType interface{}) BACnetProgramRequest {
	castFunc := func(typ interface{}) BACnetProgramRequest {
		if sBACnetProgramRequest, ok := typ.(BACnetProgramRequest); ok {
			return sBACnetProgramRequest
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetProgramRequest) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetProgramRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetProgramRequestParse(theBytes []byte) (BACnetProgramRequest, error) {
	return BACnetProgramRequestParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func BACnetProgramRequestParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetProgramRequest, error) {
	val, err := readBuffer.ReadUint8("BACnetProgramRequest", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetProgramRequest")
	}
	if enum, ok := BACnetProgramRequestByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetProgramRequest(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetProgramRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetProgramRequest) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetProgramRequest", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetProgramRequest) PLC4XEnumName() string {
	switch e {
	case BACnetProgramRequest_READY:
		return "READY"
	case BACnetProgramRequest_LOAD:
		return "LOAD"
	case BACnetProgramRequest_RUN:
		return "RUN"
	case BACnetProgramRequest_HALT:
		return "HALT"
	case BACnetProgramRequest_RESTART:
		return "RESTART"
	case BACnetProgramRequest_UNLOAD:
		return "UNLOAD"
	}
	return ""
}

func (e BACnetProgramRequest) String() string {
	return e.PLC4XEnumName()
}
