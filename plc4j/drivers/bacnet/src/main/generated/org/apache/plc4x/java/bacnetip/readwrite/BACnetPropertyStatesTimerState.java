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

public class BACnetPropertyStatesTimerState extends BACnetPropertyStates implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetTimerStateTagged timerState;

  public BACnetPropertyStatesTimerState(
      BACnetTagHeader peekedTagHeader, BACnetTimerStateTagged timerState) {
    super(peekedTagHeader);
    this.timerState = timerState;
  }

  public BACnetTimerStateTagged getTimerState() {
    return timerState;
  }

  @Override
  protected void serializeBACnetPropertyStatesChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetPropertyStatesTimerState");

    // Simple Field (timerState)
    writeSimpleField("timerState", timerState, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetPropertyStatesTimerState");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetPropertyStatesTimerState _value = this;

    // Simple field (timerState)
    lengthInBits += timerState.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetPropertyStatesTimerStateBuilder staticParseBuilder(
      ReadBuffer readBuffer, Short peekedTagNumber) throws ParseException {
    readBuffer.pullContext("BACnetPropertyStatesTimerState");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetTimerStateTagged timerState =
        readSimpleField(
            "timerState",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetTimerStateTagged.staticParse(
                        readBuffer,
                        (short) (peekedTagNumber),
                        (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    readBuffer.closeContext("BACnetPropertyStatesTimerState");
    // Create the instance
    return new BACnetPropertyStatesTimerStateBuilder(timerState);
  }

  public static class BACnetPropertyStatesTimerStateBuilder
      implements BACnetPropertyStates.BACnetPropertyStatesBuilder {
    private final BACnetTimerStateTagged timerState;

    public BACnetPropertyStatesTimerStateBuilder(BACnetTimerStateTagged timerState) {

      this.timerState = timerState;
    }

    public BACnetPropertyStatesTimerState build(BACnetTagHeader peekedTagHeader) {
      BACnetPropertyStatesTimerState bACnetPropertyStatesTimerState =
          new BACnetPropertyStatesTimerState(peekedTagHeader, timerState);
      return bACnetPropertyStatesTimerState;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPropertyStatesTimerState)) {
      return false;
    }
    BACnetPropertyStatesTimerState that = (BACnetPropertyStatesTimerState) o;
    return (getTimerState() == that.getTimerState()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getTimerState());
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
