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

public class BACnetContextTagUnknown extends BACnetContextTag implements Message {

  // Accessors for discriminator values.
  public BACnetDataType getDataType() {
    return BACnetDataType.UNKNOWN;
  }

  // Properties.
  protected final byte[] unknownData;

  // Arguments.
  protected final Short tagNumberArgument;
  protected final Long actualLength;

  public BACnetContextTagUnknown(
      BACnetTagHeader header, byte[] unknownData, Short tagNumberArgument, Long actualLength) {
    super(header, tagNumberArgument);
    this.unknownData = unknownData;
    this.tagNumberArgument = tagNumberArgument;
    this.actualLength = actualLength;
  }

  public byte[] getUnknownData() {
    return unknownData;
  }

  @Override
  protected void serializeBACnetContextTagChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetContextTagUnknown");

    // Array Field (unknownData)
    writeByteArrayField("unknownData", unknownData, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("BACnetContextTagUnknown");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetContextTagUnknown _value = this;

    // Array field
    if (unknownData != null) {
      lengthInBits += 8 * unknownData.length;
    }

    return lengthInBits;
  }

  public static BACnetContextTagUnknownBuilder staticParseBuilder(
      ReadBuffer readBuffer, Short tagNumberArgument, BACnetDataType dataType, Long actualLength)
      throws ParseException {
    readBuffer.pullContext("BACnetContextTagUnknown");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    byte[] unknownData = readBuffer.readByteArray("unknownData", Math.toIntExact(actualLength));

    readBuffer.closeContext("BACnetContextTagUnknown");
    // Create the instance
    return new BACnetContextTagUnknownBuilder(unknownData, tagNumberArgument, actualLength);
  }

  public static class BACnetContextTagUnknownBuilder
      implements BACnetContextTag.BACnetContextTagBuilder {
    private final byte[] unknownData;
    private final Short tagNumberArgument;
    private final Long actualLength;

    public BACnetContextTagUnknownBuilder(
        byte[] unknownData, Short tagNumberArgument, Long actualLength) {

      this.unknownData = unknownData;
      this.tagNumberArgument = tagNumberArgument;
      this.actualLength = actualLength;
    }

    public BACnetContextTagUnknown build(BACnetTagHeader header, Short tagNumberArgument) {
      BACnetContextTagUnknown bACnetContextTagUnknown =
          new BACnetContextTagUnknown(header, unknownData, tagNumberArgument, actualLength);
      return bACnetContextTagUnknown;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetContextTagUnknown)) {
      return false;
    }
    BACnetContextTagUnknown that = (BACnetContextTagUnknown) o;
    return (getUnknownData() == that.getUnknownData()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getUnknownData());
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
