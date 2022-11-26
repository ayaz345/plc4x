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

import java.math.BigInteger;
import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class BACnetContextTagSignedInteger extends BACnetContextTag implements Message {

  // Accessors for discriminator values.
  public BACnetDataType getDataType() {
    return BACnetDataType.SIGNED_INTEGER;
  }

  // Properties.
  protected final BACnetTagPayloadSignedInteger payload;

  // Arguments.
  protected final Short tagNumberArgument;

  public BACnetContextTagSignedInteger(
      BACnetTagHeader header, BACnetTagPayloadSignedInteger payload, Short tagNumberArgument) {
    super(header, tagNumberArgument);
    this.payload = payload;
    this.tagNumberArgument = tagNumberArgument;
  }

  public BACnetTagPayloadSignedInteger getPayload() {
    return payload;
  }

  public BigInteger getActualValue() {
    Object o = getPayload().getActualValue();
    if (o instanceof BigInteger) return (BigInteger) o;
    return BigInteger.valueOf(((Number) o).longValue());
  }

  @Override
  protected void serializeBACnetContextTagChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetContextTagSignedInteger");

    // Simple Field (payload)
    writeSimpleField("payload", payload, new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BigInteger actualValue = getActualValue();
    writeBuffer.writeVirtual("actualValue", actualValue);

    writeBuffer.popContext("BACnetContextTagSignedInteger");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetContextTagSignedInteger _value = this;

    // Simple field (payload)
    lengthInBits += payload.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetContextTagSignedIntegerBuilder staticParseBuilder(
      ReadBuffer readBuffer,
      Short tagNumberArgument,
      BACnetDataType dataType,
      BACnetTagHeader header)
      throws ParseException {
    readBuffer.pullContext("BACnetContextTagSignedInteger");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetTagPayloadSignedInteger payload =
        readSimpleField(
            "payload",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetTagPayloadSignedInteger.staticParse(
                        readBuffer, (long) (header.getActualLength())),
                readBuffer));
    BigInteger actualValue =
        readVirtualField("actualValue", BigInteger.class, payload.getActualValue());

    readBuffer.closeContext("BACnetContextTagSignedInteger");
    // Create the instance
    return new BACnetContextTagSignedIntegerBuilder(payload, tagNumberArgument);
  }

  public static class BACnetContextTagSignedIntegerBuilder
      implements BACnetContextTag.BACnetContextTagBuilder {
    private final BACnetTagPayloadSignedInteger payload;
    private final Short tagNumberArgument;

    public BACnetContextTagSignedIntegerBuilder(
        BACnetTagPayloadSignedInteger payload, Short tagNumberArgument) {

      this.payload = payload;
      this.tagNumberArgument = tagNumberArgument;
    }

    public BACnetContextTagSignedInteger build(BACnetTagHeader header, Short tagNumberArgument) {
      BACnetContextTagSignedInteger bACnetContextTagSignedInteger =
          new BACnetContextTagSignedInteger(header, payload, tagNumberArgument);
      return bACnetContextTagSignedInteger;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetContextTagSignedInteger)) {
      return false;
    }
    BACnetContextTagSignedInteger that = (BACnetContextTagSignedInteger) o;
    return (getPayload() == that.getPayload()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getPayload());
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
