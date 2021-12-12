/*
Copyright © 2021 Guy Barnhart-Magen

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
    joonix "github.com/joonix/log"
    log "github.com/sirupsen/logrus"
    "io"
    "log4jScanner/cmd"
    "log4jScanner/utils"
    "os"
    "strings"
)

var (
    Version   string
    BuildTime string
)

// TODO: add version to the version command

func setupLog(logFormat, logLevel string, file io.Writer) {
    switch strings.ToLower(logFormat) {
    case "text":
        {
            log.SetFormatter(&log.TextFormatter{}) // normal output
        }
    case "json":
        {
            log.SetFormatter(&log.JSONFormatter{}) // simple json output
        }
    case "fluentd":
        {
            log.SetFormatter(joonix.NewFormatter()) //Fluentd compatible
        }
    default:
        log.SetFormatter(joonix.NewFormatter()) //Fluentd compatible
    }

    if file != nil {
        log.SetOutput(file)
    } else {
        log.Fatal("unable to log to file")
    }
    switch strings.ToLower(logLevel) {
    case "debug":
        log.SetLevel(log.DebugLevel)
    case "warning":
        log.SetLevel(log.WarnLevel)
    default:
        log.SetLevel(log.InfoLevel)
    }
}

// TODO: log to file
// TODO: add header/pterm
// TODO: add context/cancel/done when done scanning

func main() {
    Version = "undefined"
    BuildTime = "undefined"
    utils.SetVersion(Version)
    //utils.PrintHeader()
    file, err := os.OpenFile("log4jScanner.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Error("Failed to log to file")
    }

    setupLog("text", "info", file)
    log.WithFields(log.Fields{"buildTime": BuildTime}).Debugf("Version: ", Version)

    //cmd.SetVersionTemplate("test")
    //cmd.SetHelpFunc()
    cmd.Execute()
    defer file.Close()
}