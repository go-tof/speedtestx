#!/bin/bash
#
#
# Use shell script build go program applications.
# Author: helloshaohua.
# Date: 2021.6.13.

BUILD_COMPILE_DIR="./compile"
PLATFORM_AND_CPU_ARCH=("linux/amd64" "linux/arm64" "linux/mips64")
PLATFORM=""
CPU_ARCH=""


# PrintLogFunc Format the print log function.
# This function takes two arguments：
# This first parameter: Description of the log information.
# The second parameter: Output text rendering colors.
function PrintLogFunc()
{
  log_info=$1
  use_color=$2

  # 如果{输出文字渲染颜色}长度为0时，则使用默认颜色.
  if [ -z "$use_color" ]; then
    use_color=0
  fi

  # 如果{输出文字渲染颜色}参数为1时，时使用{蓝色}文字，否则默认文字颜色.
  if [ $use_color = "WARNING" ]; then
    echo -e "\033[31m[`date +"%Y-%m-%d %H:%M:%S"`] WARNING -  ${log_info} \033[0m"
  elif [ $use_color = "INFO" ]; then
    echo -e "\033[32m[`date +"%Y-%m-%d %H:%M:%S"`] INFO -  ${log_info} \033[0m"
  else
    echo "[`date +"%Y-%m-%d %H:%M:%S"`] DEBUG - ${log_info}"
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

  BUILD_PROGRAM_NAME="nettop.$PLATFORM.$CPU_ARCH"
  GOOS="$PLATFORM" GOARCH="$CPU_ARCH" go build -o "$BUILD_COMPILE_DIR/$BUILD_PROGRAM_NAME" ./cmd/nettop

  PrintLogFunc "Build $BUILD_PROGRAM_NAME successfully~" INFO
done