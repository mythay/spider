# spider

A high performance modbus agent to collect data from different devices

## steps

- implement a agent similar with collectd modbus input plugin.

- support read multiple register at one time to improve performance

- support customer logic expression/script to process complex registers

- support complex register mapping like Zigbee in SID.


## Appendix

### collectd modbus manual page

Plugin modbus

The modbus plugin connects to a Modbus "slave" via Modbus/TCP or Modbus/RTU and reads register values. It supports reading single registers (unsigned 16 bit values), large integer values (unsigned 32 bit values) and floating point values (two registers interpreted as IEEE floats in big endian notation).

Synopsis:

 <Data "voltage-input-1">
   RegisterBase 0
   RegisterType float
   RegisterCmd ReadHolding
   Type voltage
   Instance "input-1"
 </Data>

 <Data "voltage-input-2">
   RegisterBase 2
   RegisterType float
   RegisterCmd ReadHolding
   Type voltage
   Instance "input-2"
 </Data>

 <Data "supply-temperature-1">
   RegisterBase 0
   RegisterType Int16
   RegisterCmd ReadHolding
   Type temperature
   Instance "temp-1"
 </Data>

 <Host "modbus.example.com">
   Address "192.168.0.42"
   Port    "502"
   Interval 60

   <Slave 1>
     Instance "power-supply"
     Collect  "voltage-input-1"
     Collect  "voltage-input-2"
   </Slave>
 </Host>

 <Host "localhost">
   Device "/dev/ttyUSB0"
   Baudrate 38400
   Interval 20

   <Slave 1>
     Instance "temperature"
     Collect  "supply-temperature-1"
   </Slave>
 </Host>

<Data Name> blocks

    Data blocks define a mapping between register numbers and the "types" used by collectd.

    Within <Data /> blocks, the following options are allowed:

    RegisterBase Number

        Configures the base register to read from the device. If the option RegisterType has been set to Uint32 or Float, this and the next register will be read (the register number is increased by one).
    RegisterType Int16|Int32|Uint16|Uint32|Float

        Specifies what kind of data is returned by the device. If the type is Int32, Uint32 or Float, two 16 bit registers will be read and the data is combined into one value. Defaults to Uint16.
    RegisterCmd ReadHolding|ReadInput

        Specifies register type to be collected from device. Works only with libmodbus 2.9.2 or higher. Defaults to ReadHolding.
    Type Type

        Specifies the "type" (data set) to use when dispatching the value to collectd. Currently, only data sets with exactly one data source are supported.
    Instance Instance

        Sets the type instance to use when dispatching the value to collectd. If unset, an empty string (no type instance) is used.

<Host Name> blocks

    Host blocks are used to specify to which hosts to connect and what data to read from their "slaves". The string argument Name is used as hostname when dispatching the values to collectd.

    Within <Host /> blocks, the following options are allowed:

    Address Hostname

        For Modbus/TCP, specifies the node name (the actual network address) used to connect to the host. This may be an IP address or a hostname. Please note that the used libmodbus library only supports IPv4 at the moment.
    Port Service

        for Modbus/TCP, specifies the port used to connect to the host. The port can either be given as a number or as a service name. Please note that the Service argument must be a string, even if ports are given in their numerical form. Defaults to "502".
    Device Devicenode

        For Modbus/RTU, specifies the path to the serial device being used.
    Baudrate Baudrate

        For Modbus/RTU, specifies the baud rate of the serial device. Note, connections currently support only 8/N/1.
    Interval Interval

        Sets the interval (in seconds) in which the values will be collected from this host. By default the global Interval setting will be used.
    <Slave ID>

        Over each connection, multiple Modbus devices may be reached. The slave ID is used to specify which device should be addressed. For each device you want to query, one Slave block must be given.

        Within <Slave /> blocks, the following options are allowed:

        Instance Instance

            Specify the plugin instance to use when dispatching the values to collectd. By default "slave_ID" is used.
        Collect DataName

            Specifies which data to retrieve from the device. DataName must be the same string as the Name argument passed to a Data block. You can specify this option multiple times to collect more than one value from a slave. At least one Collect option is mandatory.

```