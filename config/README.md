# Settings

This directory will be dedicated to holding files required to store preset data such as log file locations and command abilities.

## Filelog (Logger)
The filelog package will utilize settings saved in [.pondLog.toml](https://github.com/Syssos/Go_Shell/blob/main/settings/cmds.toml) in the this directory.

**In order for Filelog to utilize this file it must be stored in the home directory**

If you are unfamiliar with the TOML format, more information can be found on their [github](https://github.com/toml-lang/toml) page.

`Greeting` and `Salute` are strings ment to be printed before and after Loops run method is called.

`LogFile` will tell the logger where to store/retrieve the log file from.

`DtFormat` is the format for the datetime string. [Yourbasic.org](https://yourbasic.org/golang/format-parse-string-time-date-example/) has great example's on how to change the format to a more desired output.

`DtTimeZone` and `DtOffset` will be used when setting the timezone for the logger to base its datetime strings off of.

To change the timezone of the logger to `UTC-8` the following config settings will need to be changed
  - `DtTimeZone = "UTC -8"`
  - `DtOffset   = "-8"`

**Note the** `Greeting` **and** `Salute` **Logging settings are likely to get removed for lack of relevance to the logger.**

## hostme_test.html
The `hostme` command will allow a user to host a single html file on a webserver running on localhost. To test this feature the hostme_test.html file was created to use during the process.

Viewing this file can be done via the host me command by running the following.

```
$ hostme settings/hostme_test.html
```

## example_log.txt
This file is an example for the logger. It's sole purpose is demenstrating the output of a log file.

## images/
The images folder does not contain a readme, because its only purpose is to store images used by other readme's. 