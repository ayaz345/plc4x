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

public class BACnetLogRecordLogDatumUnsignedValue extends BACnetLogRecordLogDatum
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetContextTagUnsignedInteger unsignedValue;

  // Arguments.
  protected final Short tagNumber;

  public BACnetLogRecordLogDatumUnsignedValue(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetContextTagUnsignedInteger unsignedValue,
      Short tagNumber) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber);
    this.unsignedValue = unsignedValue;
    this.tagNumber = tagNumber;
  }

  public BACnetContextTagUnsignedInteger getUnsignedValue() {
    return unsignedValue;
  }

  @Override
  protected void serializeBACnetLogRecordLogDatumChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetLogRecordLogDatumUnsignedValue");

    // Simple Field (unsignedValue)
    writeSimpleField("unsignedValue", unsignedValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetLogRecordLogDatumUnsignedValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetLogRecordLogDatumUnsignedValue _value = this;

    // Simple field (unsignedValue)
    lengthInBits += unsignedValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetLogRecordLogDatumUnsignedValueBuilder staticParseBuilder(
      ReadBuffer readBuffer, Short tagNumber) throws ParseException {
    readBuffer.pullContext("BACnetLogRecordLogDatumUnsignedValue");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetContextTagUnsignedInteger unsignedValue =
        readSimpleField(
            "unsignedValue",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (4),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    readBuffer.closeContext("BACnetLogRecordLogDatumUnsignedValue");
    // Create the instance
    return new BACnetLogRecordLogDatumUnsignedValueBuilder(unsignedValue, tagNumber);
  }

  public static class BACnetLogRecordLogDatumUnsignedValueBuilder
      implements BACnetLogRecordLogDatum.BACnetLogRecordLogDatumBuilder {
    private final BACnetContextTagUnsignedInteger unsignedValue;
    private final Short tagNumber;

    public BACnetLogRecordLogDatumUnsignedValueBuilder(
        BACnetContextTagUnsignedInteger unsignedValue, Short tagNumber) {

      this.unsignedValue = unsignedValue;
      this.tagNumber = tagNumber;
    }

    public BACnetLogRecordLogDatumUnsignedValue build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber) {
      BACnetLogRecordLogDatumUnsignedValue bACnetLogRecordLogDatumUnsignedValue =
          new BACnetLogRecordLogDatumUnsignedValue(
              openingTag, peekedTagHeader, closingTag, unsignedValue, tagNumber);
      return bACnetLogRecordLogDatumUnsignedValue;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetLogRecordLogDatumUnsignedValue)) {
      return false;
    }
    BACnetLogRecordLogDatumUnsignedValue that = (BACnetLogRecordLogDatumUnsignedValue) o;
    return (getUnsignedValue() == that.getUnsignedValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getUnsignedValue());
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
