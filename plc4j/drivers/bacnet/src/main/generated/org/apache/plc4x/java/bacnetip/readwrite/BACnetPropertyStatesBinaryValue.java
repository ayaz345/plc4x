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

public class BACnetPropertyStatesBinaryValue extends BACnetPropertyStates implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetBinaryPVTagged binaryValue;

  public BACnetPropertyStatesBinaryValue(
      BACnetTagHeader peekedTagHeader, BACnetBinaryPVTagged binaryValue) {
    super(peekedTagHeader);
    this.binaryValue = binaryValue;
  }

  public BACnetBinaryPVTagged getBinaryValue() {
    return binaryValue;
  }

  @Override
  protected void serializeBACnetPropertyStatesChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetPropertyStatesBinaryValue");

    // Simple Field (binaryValue)
    writeSimpleField("binaryValue", binaryValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetPropertyStatesBinaryValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetPropertyStatesBinaryValue _value = this;

    // Simple field (binaryValue)
    lengthInBits += binaryValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetPropertyStatesBinaryValueBuilder staticParseBuilder(
      ReadBuffer readBuffer, Short peekedTagNumber) throws ParseException {
    readBuffer.pullContext("BACnetPropertyStatesBinaryValue");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetBinaryPVTagged binaryValue =
        readSimpleField(
            "binaryValue",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetBinaryPVTagged.staticParse(
                        readBuffer,
                        (short) (peekedTagNumber),
                        (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    readBuffer.closeContext("BACnetPropertyStatesBinaryValue");
    // Create the instance
    return new BACnetPropertyStatesBinaryValueBuilder(binaryValue);
  }

  public static class BACnetPropertyStatesBinaryValueBuilder
      implements BACnetPropertyStates.BACnetPropertyStatesBuilder {
    private final BACnetBinaryPVTagged binaryValue;

    public BACnetPropertyStatesBinaryValueBuilder(BACnetBinaryPVTagged binaryValue) {

      this.binaryValue = binaryValue;
    }

    public BACnetPropertyStatesBinaryValue build(BACnetTagHeader peekedTagHeader) {
      BACnetPropertyStatesBinaryValue bACnetPropertyStatesBinaryValue =
          new BACnetPropertyStatesBinaryValue(peekedTagHeader, binaryValue);
      return bACnetPropertyStatesBinaryValue;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPropertyStatesBinaryValue)) {
      return false;
    }
    BACnetPropertyStatesBinaryValue that = (BACnetPropertyStatesBinaryValue) o;
    return (getBinaryValue() == that.getBinaryValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getBinaryValue());
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
