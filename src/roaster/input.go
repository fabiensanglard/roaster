/*
Copyright 2019 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package roaster

import (
  "bufio"
  "errors"
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)

func intFromString(values map[string]string, key string) (int, error) {
  value := values[key]
  number, err := strconv.ParseInt(value, 0, 64)
  if err != nil {
    return 0, err
  }
  return int(number), nil
}

func uint16FromString(values map[string]string, key string) (uint16, error) {
  value := values[key]

  if value == "" {
    err := fmt.Sprintf("Missing expected key '%s'", key)
    return 0, errors.New(err)
  }

  number, err := strconv.ParseInt(value, 0, 16)
  if err != nil {
    return 0, err
  }
  return uint16(number), nil
}

func stringFromString(values map[string]string, key string) (string, error) {
  value := values[key]
  if value == "" {
    err := fmt.Sprintf("Missing expected key '%s'", key)
    return "", errors.New(err)
  }
  return value, nil
}

func readPairs(scanner * bufio.Scanner, lineNumber *int) map[string]string {
  pairs := make(map[string]string)
  for line := scanner.Text(); scanner.Scan() ; line = scanner.Text(){
    key, value := readPair(line, lineNumber)
    *lineNumber++
    if key == "" {
      break
    }
    pairs[key] = value

  }
  return pairs
}

func readPair(line string, lineNumber *int) (string, string) {
  tokens := strings.Split(line, "=")

  if len(tokens) != 2 {
    return "", ""
  }

  key := strings.Trim(tokens[0], " ")
  value := strings.Trim(tokens[1], " ")

  return key, value
}

func Dispatch() {
  file, err := os.Open("feeds.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  lineNumber := 1
  for ;scanner.Scan(); lineNumber++ {
    line := scanner.Text()

    // Comment
    if strings.HasPrefix(line,"//") {
      continue;
    }

    if len(line) == 0 {
      continue
    }

    key, value := readPair(line, &lineNumber)
    if (key != "type") {
      err := fmt.Sprintf("Error on line %d, expected 'type = X' format.", lineNumber)
      panic(err)
    }

    if dispatch[value] == nil {
      err := fmt.Sprintf("Unknown type '%s' on line %d", value, lineNumber)
      panic(err)
    }

    blockFirstLine := lineNumber
    pairs := readPairs(scanner, &lineNumber)

    err := dispatch[value].ParseParameters(pairs)
    if err != nil {
      errorString := fmt.Sprintf("Parsing error in block starting on line %d: '%s'", blockFirstLine, err)
      panic(errorString)
    }

    err = dispatch[value].Run()
    if err != nil {
      errorString := fmt.Sprintf("Change block starting on line %d failed: '%s'", blockFirstLine, err)
      panic(errorString)
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}

