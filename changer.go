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

/*
#include <sys/ioctl.h>
#include <net/if_arp.h>
#include <net/if.h>
#include <malloc.h>
#include <string.h>


int _set_hardware_address(char* name, char addr[]) {
	struct ifreq ifr;
	int s;

	memcpy(ifr.ifr_hwaddr.sa_data, addr, 6);
	strcpy(ifr.ifr_name, name);

	s = socket(AF_INET, SOCK_DGRAM, 0);
	if (s == -1) {
		return -1;
	}

	ifr.ifr_hwaddr.sa_family = ARPHRD_ETHER;
	return ioctl(s, SIOCSIFHWADDR, &ifr);
}

*/
import "C"

import (
	"fmt"
	"net"
	"unsafe"
)

// Change the Hardware Address (MAC) for the interface `iface`
// to `mac`.
func ChangeHardwareAddr(iface net.Interface, mac net.HardwareAddr) error {
	cName := C.CString(iface.Name)
	defer C.free(unsafe.Pointer(cName))

	cAddr := (*C.char)(C.CBytes(mac))
	defer C.free(unsafe.Pointer(cAddr))

	if i := C._set_hardware_address(cName, cAddr); i != 0 {
		return fmt.Errorf("Failed to set Hardware Address")
	}
	return nil
}

// vim: foldmethod=marker
