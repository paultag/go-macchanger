// {{{ Copyright (c) Paul R. Tagliamonte <paultag@gmail.com>, 2017
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE. }}}

package macchanger

import (
	"net"
	"unsafe"

	"golang.org/x/sys/unix"
)

type ifreq_hwaddr struct {
	name   [unix.IFNAMSIZ]byte
	hwaddr unix.RawSockaddr
}

// Change the Hardware Address (MAC) for the interface `iface`
// to `mac`.
func ChangeHardwareAddr(iface net.Interface, mac net.HardwareAddr) error {
	var ifr ifreq_hwaddr
	copy(ifr.name[:], append([]byte(iface.Name), 0))
	ifr.hwaddr.Family = unix.ARPHRD_ETHER
	if len(mac) > len(ifr.hwaddr.Data) {
		mac = mac[:len(ifr.hwaddr.Data)]
	}
	for i, _ := range mac {
		ifr.hwaddr.Data[i] = int8(mac[i])
	}

	s, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)
	if err != nil {
		return err
	}
	defer unix.Close(s)
	return unix.IoctlSetInt(s, unix.SIOCSIFHWADDR, int(uintptr(unsafe.Pointer(&ifr))))
}

// vim: foldmethod=marker
