# zimt

> Operate your IoT network from terminal.

## Installation

    @TODO

## Configuration

Configuration file is optional, `zimt` will use its default values for the options in case file is missing

Configuration file is expected to be found either at default location: `~/.zimt.yaml` or specified with `--config` flag.

It might be useful to check the [configuration file example](./docs/.zimt.yaml.example).

The configuration options are:

```yaml
mqtt:
  # the broker host, optional, by defaut set to `localhost`
  broker: localhost

  # the broker port, optional, by default set to `1883`
  port: 1883

  # zigbee2mqtt basic topic, optional, by default set to `zigbee2mqtt`
  base-topic: zigbee2mqtt

  # broker auth user, optional, zimt doesn't send it to the broker if missing
  user: <mqtt-user>

  # broker auth password, optional, zimt doesn't send it to the broker if missing
  password: <mqtt-password>

  # client idendifier visible in connection list, optional, by default set to `zimt`
  client-id: zimt
```

## Usage

`zimt` is a CLI tool to manage your zigbee connected devices via [zigbee2mqtt bridge](https://github.com/Koenkk/zigbee2mqtt).

### Synopsis

    zimt [flags] <command> <subcommand> [parameters]

### Flags

    --help

Print help information

    --verbose

Turn on verbose/debug logging

    --config /path/to/config/file

Take configuration file from not default location

### Commands

    zimt bridge

Prints bridge version

    zimt broker

Prints broker details

    zimt config

Prints zimt configuration

    zimt device list

Prints all connected devices

    zimt version

Prints zimt version and build information

## License

`zimt` is licensed under MIT license. (see [LICENSE](./LICENSE))

Some parts of `zimt` use third party libraries which are licensed under different licenses:

Library | License
---|---
https://github.com/eclipse/paho.mqtt.golang | [Eclipse Public License 1.0](https://github.com/eclipse/paho.mqtt.golang/blob/master/LICENSE)
https://github.com/mitchellh/go-homedir | [MIT License](https://github.com/mitchellh/go-homedir/blob/master/LICENSE)
https://github.com/spf13/cobra | [Apache License 2.0](https://github.com/spf13/cobra/blob/master/LICENSE.txt)
https://github.com/spf13/viper | [MIT License](https://github.com/spf13/viper/blob/master/LICENSE)
