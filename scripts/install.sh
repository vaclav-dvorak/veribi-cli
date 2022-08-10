#!/usr/bin/env sh
set -e

uname_os() {
  os=$(uname -s | tr '[:upper:]' '[:lower:]')
  case "$os" in
    cygwin_nt*) os="windows" ;;
    mingw*) os="windows" ;;
    msys_nt*) os="windows" ;;
  esac
  echo "$os"
}

uname_arch() {
  arch=$(uname -m)
  case $arch in
    x86_64) arch="amd64" ;;
    x86) arch="386" ;;
    i686) arch="386" ;;
    i386) arch="386" ;;
    aarch64) arch="arm64" ;;
    armv5*) arch="armv5" ;;
    armv6*) arch="armv6" ;;
    armv7*) arch="armv7" ;;
  esac
  echo ${arch}
}

uname_os_check() {
  os=$(uname_os)
  case "$os" in
    darwin) return 0 ;;
    linux) return 0 ;;
    windows) return 0 ;;
  esac
  echo "Unsupported operation system '$(uname -s)'."
  return 1
}

uname_arch_check() {
  arch=$(uname_arch)
  case "$arch" in
    # 386) return 0 ;;
    amd64) return 0 ;;
    # arm64) return 0 ;;
    # armv5) return 0 ;;
    # armv6) return 0 ;;
    # armv7) return 0 ;;
  esac
  echo "Unsupported architecture '$(uname -m)'."
  return 1
}

uname_os_check
uname_arch_check

name="veribi_${os}_${arch}"
tar="$name.tar.gz"
url="https://github.com/vaclav-dvorak/veribi-cli/releases/latest/download"
echo "Downloading latest release of ${name}..."
curl -sL ${url}/$tar -o "/tmp/$tar" || exit 1
tar xzf "/tmp/${tar}" -C /tmp
rm "/tmp/$tar" "/tmp/LICENSE"
chmod +x "/tmp/veribi"

echo "Moving /tmp/veribi to /usr/local/bin/veribi (you might be asked for your password due to sudo)"
if [ -x "$(command -v sudo)" ]; then
  sudo mv "/tmp/veribi" "/usr/local/bin/veribi"
else
  mv "/tmp/veribi" "/usr/local/bin/veribi"
fi
echo
echo "Completed installing $(veribi --version)"
