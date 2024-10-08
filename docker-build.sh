#!/usr/bin/env bash
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
export PATH

init_var() {
  ECHO_TYPE="echo -e"

  trojan_core_version=3.0.0

  arch_arr="linux/386,linux/amd64,linux/arm/v6,linux/arm/v7,linux/arm64,linux/ppc64le,linux/s390x"
}

echo_content() {
  case $1 in
  "red")
    ${ECHO_TYPE} "\033[31m$2\033[0m"
    ;;
  "green")
    ${ECHO_TYPE} "\033[32m$2\033[0m"
    ;;
  "yellow")
    ${ECHO_TYPE} "\033[33m$2\033[0m"
    ;;
  "blue")
    ${ECHO_TYPE} "\033[34m$2\033[0m"
    ;;
  "purple")
    ${ECHO_TYPE} "\033[35m$2\033[0m"
    ;;
  "skyBlue")
    ${ECHO_TYPE} "\033[36m$2\033[0m"
    ;;
  "white")
    ${ECHO_TYPE} "\033[37m$2\033[0m"
    ;;
  esac
}

main() {
  echo_content skyBlue "start build trojan-core CPU：${arch_arr}"

  docker buildx build -t jonssonyan/trojan-core:latest --platform ${arch_arr} --push .
  if [[ "$?" == "0" ]]; then
    echo_content green "trojan-core Version：latest CPU：${arch_arr} build success"

    if [[ ${trojan_core_version} != "latest" ]]; then
      docker buildx build -t jonssonyan/trojan-core:${trojan_core_version} --platform ${arch_arr} --push .
      if [[ "$?" == "0" ]]; then
        echo_content green "trojan-core-linux Version：${trojan_core_version} CPU：${arch_arr} build success"
      else
        echo_content red "trojan-core-linux Version：${trojan_core_version} CPU：${arch_arr} build failed"
      fi
    fi
  else
    echo_content red "trojan-core-linux Version：latest CPU：${arch_arr} build failed"
  fi

  echo_content skyBlue "trojan-core CPU：${arch_arr} build finished"
}

init_var
main
