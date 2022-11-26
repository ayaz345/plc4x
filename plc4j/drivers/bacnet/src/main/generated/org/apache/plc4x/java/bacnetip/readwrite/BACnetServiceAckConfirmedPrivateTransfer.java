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
package org.apache.plc4x.java.bacnetip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class BACnetServiceAckConfirmedPrivateTransfer extends BACnetServiceAck implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.CONFIRMED_PRIVATE_TRANSFER;
  }

  // Properties.
  protected final BACnetVendorIdTagged vendorId;
  protected final BACnetContextTagUnsignedInteger serviceNumber;
  protected final BACnetConstructedData resultBlock;

  // Arguments.
  protected final Long serviceAckLength;

  public BACnetServiceAckConfirmedPrivateTransfer(
      BACnetVendorIdTagged vendorId,
      BACnetContextTagUnsignedInteger serviceNumber,
      BACnetConstructedData resultBlock,
      Long serviceAckLength) {
    super(serviceAckLength);
    this.vendorId = vendorId;
    this.serviceNumber = serviceNumber;
    this.resultBlock = resultBlock;
    this.serviceAckLength = serviceAckLength;
  }

  public BACnetVendorIdTagged getVendorId() {
    return vendorId;
  }

  public BACnetContextTagUnsignedInteger getServiceNumber() {
    return serviceNumber;
  }

  public BACnetConstructedData getResultBlock() {
    return resultBlock;
  }

  @Override
  protected void serializeBACnetServiceAckChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetServiceAckConfirmedPrivateTransfer");

    // Simple Field (vendorId)
    writeSimpleField("vendorId", vendorId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (serviceNumber)
    writeSimpleField("serviceNumber", serviceNumber, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (resultBlock) (Can be skipped, if the value is null)
    writeOptionalField("resultBlock", resultBlock, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetServiceAckConfirmedPrivateTransfer");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetServiceAckConfirmedPrivateTransfer _value = this;

    // Simple field (vendorId)
    lengthInBits += vendorId.getLengthInBits();

    // Simple field (serviceNumber)
    lengthInBits += serviceNumber.getLengthInBits();

    // Optional Field (resultBlock)
    if (resultBlock != null) {
      lengthInBits += resultBlock.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetServiceAckConfirmedPrivateTransferBuilder staticParseBuilder(
      ReadBuffer readBuffer, Long serviceAckLength) throws ParseException {
    readBuffer.pullContext("BACnetServiceAckConfirmedPrivateTransfer");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetVendorIdTagged vendorId =
        readSimpleField(
            "vendorId",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetVendorIdTagged.staticParse(
                        readBuffer, (short) (0), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagUnsignedInteger serviceNumber =
        readSimpleField(
            "serviceNumber",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetConstructedData resultBlock =
        readOptionalField(
            "resultBlock",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetConstructedData.staticParse(
                        readBuffer,
                        (short) (2),
                        (BACnetObjectType) (BACnetObjectType.VENDOR_PROPRIETARY_VALUE),
                        (BACnetPropertyIdentifier)
                            (BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE),
                        (BACnetTagPayloadUnsignedInteger) (null)),
                readBuffer));

    readBuffer.closeContext("BACnetServiceAckConfirmedPrivateTransfer");
    // Create the instance
    return new BACnetServiceAckConfirmedPrivateTransferBuilder(
        vendorId, serviceNumber, resultBlock, serviceAckLength);
  }

  public static class BACnetServiceAckConfirmedPrivateTransferBuilder
      implements BACnetServiceAck.BACnetServiceAckBuilder {
    private final BACnetVendorIdTagged vendorId;
    private final BACnetContextTagUnsignedInteger serviceNumber;
    private final BACnetConstructedData resultBlock;
    private final Long serviceAckLength;

    public BACnetServiceAckConfirmedPrivateTransferBuilder(
        BACnetVendorIdTagged vendorId,
        BACnetContextTagUnsignedInteger serviceNumber,
        BACnetConstructedData resultBlock,
        Long serviceAckLength) {

      this.vendorId = vendorId;
      this.serviceNumber = serviceNumber;
      this.resultBlock = resultBlock;
      this.serviceAckLength = serviceAckLength;
    }

    public BACnetServiceAckConfirmedPrivateTransfer build(Long serviceAckLength) {

      BACnetServiceAckConfirmedPrivateTransfer bACnetServiceAckConfirmedPrivateTransfer =
          new BACnetServiceAckConfirmedPrivateTransfer(
              vendorId, serviceNumber, resultBlock, serviceAckLength);
      return bACnetServiceAckConfirmedPrivateTransfer;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetServiceAckConfirmedPrivateTransfer)) {
      return false;
    }
    BACnetServiceAckConfirmedPrivateTransfer that = (BACnetServiceAckConfirmedPrivateTransfer) o;
    return (getVendorId() == that.getVendorId())
        && (getServiceNumber() == that.getServiceNumber())
        && (getResultBlock() == that.getResultBlock())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getVendorId(), getServiceNumber(), getResultBlock());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
