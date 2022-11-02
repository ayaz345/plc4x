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

// KnxLayer is an enum
type KnxLayer uint8

type IKnxLayer interface {
	utils.Serializable
}

const (
	KnxLayer_TUNNEL_LINK_LAYER KnxLayer = 0x02
	KnxLayer_TUNNEL_RAW        KnxLayer = 0x04
	KnxLayer_TUNNEL_BUSMONITOR KnxLayer = 0x80
)

var KnxLayerValues []KnxLayer

func init() {
	_ = errors.New
	KnxLayerValues = []KnxLayer{
		KnxLayer_TUNNEL_LINK_LAYER,
		KnxLayer_TUNNEL_RAW,
		KnxLayer_TUNNEL_BUSMONITOR,
	}
}

func KnxLayerByValue(value uint8) (enum KnxLayer, ok bool) {
	switch value {
	case 0x02:
		return KnxLayer_TUNNEL_LINK_LAYER, true
	case 0x04:
		return KnxLayer_TUNNEL_RAW, true
	case 0x80:
		return KnxLayer_TUNNEL_BUSMONITOR, true
	}
	return 0, false
}

func KnxLayerByName(value string) (enum KnxLayer, ok bool) {
	switch value {
	case "TUNNEL_LINK_LAYER":
		return KnxLayer_TUNNEL_LINK_LAYER, true
	case "TUNNEL_RAW":
		return KnxLayer_TUNNEL_RAW, true
	case "TUNNEL_BUSMONITOR":
		return KnxLayer_TUNNEL_BUSMONITOR, true
	}
	return 0, false
}

func KnxLayerKnows(value uint8) bool {
	for _, typeValue := range KnxLayerValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastKnxLayer(structType interface{}) KnxLayer {
	castFunc := func(typ interface{}) KnxLayer {
		if sKnxLayer, ok := typ.(KnxLayer); ok {
			return sKnxLayer
		}
		return 0
	}
	return castFunc(structType)
}

func (m KnxLayer) GetLengthInBits() uint16 {
	return 8
}

func (m KnxLayer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxLayerParse(theBytes []byte) (KnxLayer, error) {
	return KnxLayerParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func KnxLayerParseWithBuffer(readBuffer utils.ReadBuffer) (KnxLayer, error) {
	val, err := readBuffer.ReadUint8("KnxLayer", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading KnxLayer")
	}
	if enum, ok := KnxLayerByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return KnxLayer(val), nil
	} else {
		return enum, nil
	}
}

func (e KnxLayer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e KnxLayer) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("KnxLayer", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e KnxLayer) PLC4XEnumName() string {
	switch e {
	case KnxLayer_TUNNEL_LINK_LAYER:
		return "TUNNEL_LINK_LAYER"
	case KnxLayer_TUNNEL_RAW:
		return "TUNNEL_RAW"
	case KnxLayer_TUNNEL_BUSMONITOR:
		return "TUNNEL_BUSMONITOR"
	}
	return ""
}

func (e KnxLayer) String() string {
	return e.PLC4XEnumName()
}
