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

// NPDUNetworkPriority is an enum
type NPDUNetworkPriority uint8

type INPDUNetworkPriority interface {
	utils.Serializable
}

const (
	NPDUNetworkPriority_LIFE_SAVETY_MESSAGE        NPDUNetworkPriority = 3
	NPDUNetworkPriority_CRITICAL_EQUIPMENT_MESSAGE NPDUNetworkPriority = 2
	NPDUNetworkPriority_URGENT_MESSAGE             NPDUNetworkPriority = 1
	NPDUNetworkPriority_NORMAL_MESSAGE             NPDUNetworkPriority = 0
)

var NPDUNetworkPriorityValues []NPDUNetworkPriority

func init() {
	_ = errors.New
	NPDUNetworkPriorityValues = []NPDUNetworkPriority{
		NPDUNetworkPriority_LIFE_SAVETY_MESSAGE,
		NPDUNetworkPriority_CRITICAL_EQUIPMENT_MESSAGE,
		NPDUNetworkPriority_URGENT_MESSAGE,
		NPDUNetworkPriority_NORMAL_MESSAGE,
	}
}

func NPDUNetworkPriorityByValue(value uint8) (enum NPDUNetworkPriority, ok bool) {
	switch value {
	case 0:
		return NPDUNetworkPriority_NORMAL_MESSAGE, true
	case 1:
		return NPDUNetworkPriority_URGENT_MESSAGE, true
	case 2:
		return NPDUNetworkPriority_CRITICAL_EQUIPMENT_MESSAGE, true
	case 3:
		return NPDUNetworkPriority_LIFE_SAVETY_MESSAGE, true
	}
	return 0, false
}

func NPDUNetworkPriorityByName(value string) (enum NPDUNetworkPriority, ok bool) {
	switch value {
	case "NORMAL_MESSAGE":
		return NPDUNetworkPriority_NORMAL_MESSAGE, true
	case "URGENT_MESSAGE":
		return NPDUNetworkPriority_URGENT_MESSAGE, true
	case "CRITICAL_EQUIPMENT_MESSAGE":
		return NPDUNetworkPriority_CRITICAL_EQUIPMENT_MESSAGE, true
	case "LIFE_SAVETY_MESSAGE":
		return NPDUNetworkPriority_LIFE_SAVETY_MESSAGE, true
	}
	return 0, false
}

func NPDUNetworkPriorityKnows(value uint8) bool {
	for _, typeValue := range NPDUNetworkPriorityValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastNPDUNetworkPriority(structType interface{}) NPDUNetworkPriority {
	castFunc := func(typ interface{}) NPDUNetworkPriority {
		if sNPDUNetworkPriority, ok := typ.(NPDUNetworkPriority); ok {
			return sNPDUNetworkPriority
		}
		return 0
	}
	return castFunc(structType)
}

func (m NPDUNetworkPriority) GetLengthInBits() uint16 {
	return 2
}

func (m NPDUNetworkPriority) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NPDUNetworkPriorityParse(theBytes []byte) (NPDUNetworkPriority, error) {
	return NPDUNetworkPriorityParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func NPDUNetworkPriorityParseWithBuffer(readBuffer utils.ReadBuffer) (NPDUNetworkPriority, error) {
	val, err := readBuffer.ReadUint8("NPDUNetworkPriority", 2)
	if err != nil {
		return 0, errors.Wrap(err, "error reading NPDUNetworkPriority")
	}
	if enum, ok := NPDUNetworkPriorityByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return NPDUNetworkPriority(val), nil
	} else {
		return enum, nil
	}
}

func (e NPDUNetworkPriority) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e NPDUNetworkPriority) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("NPDUNetworkPriority", 2, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e NPDUNetworkPriority) PLC4XEnumName() string {
	switch e {
	case NPDUNetworkPriority_NORMAL_MESSAGE:
		return "NORMAL_MESSAGE"
	case NPDUNetworkPriority_URGENT_MESSAGE:
		return "URGENT_MESSAGE"
	case NPDUNetworkPriority_CRITICAL_EQUIPMENT_MESSAGE:
		return "CRITICAL_EQUIPMENT_MESSAGE"
	case NPDUNetworkPriority_LIFE_SAVETY_MESSAGE:
		return "LIFE_SAVETY_MESSAGE"
	}
	return ""
}

func (e NPDUNetworkPriority) String() string {
	return e.PLC4XEnumName()
}
