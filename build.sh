#!/bin/bash
#
#
# Use shell script build go program applications.
# Author: helloshaohua.
# Date: 2021.6.13.

BUILD_COMPILE_DIR="./compile"
PLATFORM_AND_CPU_ARCH=("darwin/amd64" "linux/amd64" "linux/arm64" "linux/mips64" "windows/amd64" "windows/386" "windows/arm")
PLATFORM=""
CPU_ARCH=""
readonly WINDOWS="windows"
readonly WINDOWS_EXECUTABLE_FILE_EXT="exe"


# PrintLogFunc Format the print log function.
# This function takes two arguments：
# This first parameter: Description of the log information.
# The second parameter: Output text rendering colors.
function PrintLogFunc()
{
  log_info=$1
  use_color=$2

  # If the length of {use_color} is 0, the default color will be used.
  if [ -z "$use_color" ]; then
    use_color=0
  fi

  if [ $use_color = "WARNING" ]; then
    echo -e "\033[31m[$(date +"%Y-%m-%d %H:%M:%S")] WARNING -  ${log_info} \033[0m"
  elif [ $use_color = "INFO" ]; then
    echo -e "\033[32m[$(date +"%Y-%m-%d %H:%M:%S")] INFO -  ${log_info} \033[0m"
  else
    echo "[$(date +"%Y-%m-%d %H:%M:%S")] DEBUG - ${log_info}"
  fi

  # 输出两行空行.
  echo;
}

# GetPlatformAndCPUArchFunc get platfrom and cpu arch info.
function GetPlatformAndCPUArchFunc()
{
  info=$1
  if [ ${#info} == 0 ]; then
    PrintLogFunc "platfrom and cpu arch info can't empty" WARNING
    exit 1
  fi

  INFO_ARR=(${info//// })

  for (( i = 0; i < ${#INFO_ARR[@]}; i++ )); do
    case $i in
    0)
      PLATFORM=${INFO_ARR[i]};;
    1)
      CPU_ARCH=${INFO_ARR[i]};;
    esac
  done
}

# Build on the loop.
for INFO in "${PLATFORM_AND_CPU_ARCH[@]}"; do
  GetPlatformAndCPUArchFunc "$INFO"

  BUILD_PROGRAM_NAME="speedtestx.$PLATFORM.$CPU_ARCH"

  if [ $PLATFORM = $WINDOWS ]; then
    BUILD_PROGRAM_NAME="${BUILD_PROGRAM_NAME}.${WINDOWS_EXECUTABLE_FILE_EXT}"
  fi

  GOOS="$PLATFORM" GOARCH="$CPU_ARCH" go build -o "$BUILD_COMPILE_DIR/$BUILD_PROGRAM_NAME" ./cmd/speedtestx

  PrintLogFunc "Build $BUILD_PROGRAM_NAME successfully~" INFO
done