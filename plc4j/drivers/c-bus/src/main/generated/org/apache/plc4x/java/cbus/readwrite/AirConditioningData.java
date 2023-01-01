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
package org.apache.plc4x.java.cbus.readwrite;

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

public abstract class AirConditioningData implements Message {

  // Abstract accessors for discriminator values.

  // Properties.
  protected final AirConditioningCommandTypeContainer commandTypeContainer;

  public AirConditioningData(AirConditioningCommandTypeContainer commandTypeContainer) {
    super();
    this.commandTypeContainer = commandTypeContainer;
  }

  public AirConditioningCommandTypeContainer getCommandTypeContainer() {
    return commandTypeContainer;
  }

  public AirConditioningCommandType getCommandType() {
    return (AirConditioningCommandType) (getCommandTypeContainer().getCommandType());
  }

  protected abstract void serializeAirConditioningDataChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("AirConditioningData");

    // Simple Field (commandTypeContainer)
    writeSimpleEnumField(
        "commandTypeContainer",
        "AirConditioningCommandTypeContainer",
        commandTypeContainer,
        new DataWriterEnumDefault<>(
            AirConditioningCommandTypeContainer::getValue,
            AirConditioningCommandTypeContainer::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    AirConditioningCommandType commandType = getCommandType();
    writeBuffer.writeVirtual("commandType", commandType);

    // Switch field (Serialize the sub-type)
    serializeAirConditioningDataChild(writeBuffer);

    writeBuffer.popContext("AirConditioningData");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    AirConditioningData _value = this;

    // Simple field (commandTypeContainer)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static AirConditioningData staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static AirConditioningData staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("AirConditioningData");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    // Validation
    if (!(org.apache.plc4x.java.cbus.readwrite.utils.StaticHelper
        .knowsAirConditioningCommandTypeContainer(readBuffer))) {
      throw new ParseAssertException("no command type could be found");
    }

    AirConditioningCommandTypeContainer commandTypeContainer =
        readEnumField(
            "commandTypeContainer",
            "AirConditioningCommandTypeContainer",
            new DataReaderEnumDefault<>(
                AirConditioningCommandTypeContainer::enumForValue,
                readUnsignedShort(readBuffer, 8)));
    AirConditioningCommandType commandType =
        readVirtualField(
            "commandType", AirConditioningCommandType.class, commandTypeContainer.getCommandType());

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    AirConditioningDataBuilder builder = null;
    if (EvaluationHelper.equals(commandType, AirConditioningCommandType.HVAC_SCHEDULE_ENTRY)) {
      builder = AirConditioningDataHvacScheduleEntry.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, AirConditioningCommandType.HUMIDITY_SCHEDULE_ENTRY)) {
      builder = AirConditioningDataHumidityScheduleEntry.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, AirConditioningCommandType.REFRESH)) {
      builder = AirConditioningDataRefresh.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, AirConditioningCommandType.ZONE_HVAC_PLANT_STATUS)) {
      builder = AirConditioningDataZoneHvacPlantStatus.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, AirConditioningCommandType.ZONE_HUMIDITY_PLANT_STATUS)) {
      builder = AirConditioningDataZoneHumidityPlantStatus.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, AirConditioningCommandType.ZONE_TEMPERATURE)) {
      builder = AirConditioningDataZoneTemperature.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, AirConditioningCommandType.ZONE_HUMIDITY)) {
      builder = AirConditioningDataZoneHumidity.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, AirConditioningCommandType.SET_ZONE_GROUP_OFF)) {
      builder = AirConditioningDataSetZoneGroupOff.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, AirConditioningCommandType.SET_ZONE_GROUP_ON)) {
      builder = AirConditioningDataSetZoneGroupOn.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, AirConditioningCommandType.SET_ZONE_HVAC_MODE)) {
      builder = AirConditioningDataSetZoneHvacMode.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, AirConditioningCommandType.SET_PLANT_HVAC_LEVEL)) {
      builder = AirConditioningDataSetPlantHvacLevel.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, AirConditioningCommandType.SET_ZONE_HUMIDITY_MODE)) {
      builder = AirConditioningDataSetZoneHumidityMode.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, AirConditioningCommandType.SET_PLANT_HUMIDITY_LEVEL)) {
      builder = AirConditioningDataSetPlantHumidityLevel.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, AirConditioningCommandType.SET_HVAC_UPPER_GUARD_LIMIT)) {
      builder = AirConditioningDataSetHvacUpperGuardLimit.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, AirConditioningCommandType.SET_HVAC_LOWER_GUARD_LIMIT)) {
      builder = AirConditioningDataSetHvacLowerGuardLimit.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, AirConditioningCommandType.SET_HVAC_SETBACK_LIMIT)) {
      builder = AirConditioningDataSetHvacSetbackLimit.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, AirConditioningCommandType.SET_HUMIDITY_UPPER_GUARD_LIMIT)) {
      builder = AirConditioningDataSetHumidityUpperGuardLimit.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, AirConditioningCommandType.SET_HUMIDITY_LOWER_GUARD_LIMIT)) {
      builder = AirConditioningDataSetHumidityLowerGuardLimit.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, AirConditioningCommandType.SET_HUMIDITY_SETBACK_LIMIT)) {
      builder = AirConditioningDataSetHumiditySetbackLimit.staticParseBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "commandType="
              + commandType
              + "]");
    }

    readBuffer.closeContext("AirConditioningData");
    // Create the instance
    AirConditioningData _airConditioningData = builder.build(commandTypeContainer);
    return _airConditioningData;
  }

  public static interface AirConditioningDataBuilder {
    AirConditioningData build(AirConditioningCommandTypeContainer commandTypeContainer);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AirConditioningData)) {
      return false;
    }
    AirConditioningData that = (AirConditioningData) o;
    return (getCommandTypeContainer() == that.getCommandTypeContainer()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getCommandTypeContainer());
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