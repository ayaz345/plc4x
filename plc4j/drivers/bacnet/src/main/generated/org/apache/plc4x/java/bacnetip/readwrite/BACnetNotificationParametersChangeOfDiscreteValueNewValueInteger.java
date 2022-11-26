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

public class BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger
    extends BACnetNotificationParametersChangeOfDiscreteValueNewValue implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagSignedInteger integerValue;

  // Arguments.
  protected final Short tagNumber;

  public BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetApplicationTagSignedInteger integerValue,
      Short tagNumber) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber);
    this.integerValue = integerValue;
    this.tagNumber = tagNumber;
  }

  public BACnetApplicationTagSignedInteger getIntegerValue() {
    return integerValue;
  }

  @Override
  protected void serializeBACnetNotificationParametersChangeOfDiscreteValueNewValueChild(
      WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger");

    // Simple Field (integerValue)
    writeSimpleField("integerValue", integerValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger _value = this;

    // Simple field (integerValue)
    lengthInBits += integerValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetNotificationParametersChangeOfDiscreteValueNewValueIntegerBuilder
      staticParseBuilder(ReadBuffer readBuffer, Short tagNumber) throws ParseException {
    readBuffer.pullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetApplicationTagSignedInteger integerValue =
        readSimpleField(
            "integerValue",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetApplicationTagSignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger");
    // Create the instance
    return new BACnetNotificationParametersChangeOfDiscreteValueNewValueIntegerBuilder(
        integerValue, tagNumber);
  }

  public static class BACnetNotificationParametersChangeOfDiscreteValueNewValueIntegerBuilder
      implements BACnetNotificationParametersChangeOfDiscreteValueNewValue
          .BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
    private final BACnetApplicationTagSignedInteger integerValue;
    private final Short tagNumber;

    public BACnetNotificationParametersChangeOfDiscreteValueNewValueIntegerBuilder(
        BACnetApplicationTagSignedInteger integerValue, Short tagNumber) {

      this.integerValue = integerValue;
      this.tagNumber = tagNumber;
    }

    public BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber) {
      BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger
          bACnetNotificationParametersChangeOfDiscreteValueNewValueInteger =
              new BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger(
                  openingTag, peekedTagHeader, closingTag, integerValue, tagNumber);
      return bACnetNotificationParametersChangeOfDiscreteValueNewValueInteger;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger)) {
      return false;
    }
    BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger that =
        (BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger) o;
    return (getIntegerValue() == that.getIntegerValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getIntegerValue());
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
