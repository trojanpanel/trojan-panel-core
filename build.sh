#!/usr/bin/env bash
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
export PATH

init_var() {
  ECHO_TYPE="echo -e"

  trojan_panel_core_version=2.3.1

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
  echo_content skyBlue "start build trojan-panel-core CPU：${arch_arr}"

  docker buildx build -t jonssonyan/trojan-panel-core:latest --platform ${arch_arr} --push .
  if [[ "$?" == "0" ]]; then
    echo_content green "trojan-panel-core Version：latest CPU：${arch_arr} build success"

    if [[ ${trojan_panel_core_version} != "latest" ]]; then
      docker buildx build -t jonssonyan/trojan-panel-core:${trojan_panel_core_version} --platform ${arch_arr} --push .
      if [[ "$?" == "0" ]]; then
        echo_content green "trojan-panel-core-linux Version：${trojan_panel_core_version} CPU：${arch_arr} build success"
      else
        echo_content red "trojan-panel-core-linux Version：${trojan_panel_core_version} CPU：${arch_arr} build failed"
      fi
    fi
  else
    echo_content red "trojan-panel-core-linux Version：latest CPU：${arch_arr} build failed"
  fi

  echo_content skyBlue "trojan-panel-core CPU：${arch_arr} build finished"
}

init_var
main
