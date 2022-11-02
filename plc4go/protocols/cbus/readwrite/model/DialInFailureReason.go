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

// DialInFailureReason is an enum
type DialInFailureReason uint8

type IDialInFailureReason interface {
	utils.Serializable
}

const (
	DialInFailureReason_PHONE_STOPPED_RINGING DialInFailureReason = 0x01
)

var DialInFailureReasonValues []DialInFailureReason

func init() {
	_ = errors.New
	DialInFailureReasonValues = []DialInFailureReason{
		DialInFailureReason_PHONE_STOPPED_RINGING,
	}
}

func DialInFailureReasonByValue(value uint8) (enum DialInFailureReason, ok bool) {
	switch value {
	case 0x01:
		return DialInFailureReason_PHONE_STOPPED_RINGING, true
	}
	return 0, false
}

func DialInFailureReasonByName(value string) (enum DialInFailureReason, ok bool) {
	switch value {
	case "PHONE_STOPPED_RINGING":
		return DialInFailureReason_PHONE_STOPPED_RINGING, true
	}
	return 0, false
}

func DialInFailureReasonKnows(value uint8) bool {
	for _, typeValue := range DialInFailureReasonValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastDialInFailureReason(structType interface{}) DialInFailureReason {
	castFunc := func(typ interface{}) DialInFailureReason {
		if sDialInFailureReason, ok := typ.(DialInFailureReason); ok {
			return sDialInFailureReason
		}
		return 0
	}
	return castFunc(structType)
}

func (m DialInFailureReason) GetLengthInBits() uint16 {
	return 8
}

func (m DialInFailureReason) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DialInFailureReasonParse(theBytes []byte) (DialInFailureReason, error) {
	return DialInFailureReasonParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func DialInFailureReasonParseWithBuffer(readBuffer utils.ReadBuffer) (DialInFailureReason, error) {
	val, err := readBuffer.ReadUint8("DialInFailureReason", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading DialInFailureReason")
	}
	if enum, ok := DialInFailureReasonByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return DialInFailureReason(val), nil
	} else {
		return enum, nil
	}
}

func (e DialInFailureReason) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e DialInFailureReason) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("DialInFailureReason", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e DialInFailureReason) PLC4XEnumName() string {
	switch e {
	case DialInFailureReason_PHONE_STOPPED_RINGING:
		return "PHONE_STOPPED_RINGING"
	}
	return ""
}

func (e DialInFailureReason) String() string {
	return e.PLC4XEnumName()
}
