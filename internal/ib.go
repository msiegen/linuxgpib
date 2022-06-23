// Copyright 2022 Google LLC
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// version 2 as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

package internal

/*
#cgo linux LDFLAGS: -lgpib
#include <stdlib.h>
#include <gpib/ib.h>
*/
import "C"

import (
	"unsafe"
)

func AllSPoll(board_desc int, addressList []Address) (resultList []int) {
	addressList2 := append(addressList, NOADDR)
	resultList = make([]int, len(addressList))
	C.AllSPoll(C.int(board_desc), (*C.Addr4882_t)(unsafe.Pointer(&addressList2[0])), (*C.short)(unsafe.Pointer(&resultList[0])))
	return
}

func DevClear(board_desc int, address Address) {
	C.DevClear(C.int(board_desc), C.Addr4882_t(address))
	return
}

func DevClearList(board_desc int, addressList []Address) {
	addressList2 := append(addressList, NOADDR)
	C.DevClearList(C.int(board_desc), (*C.Addr4882_t)(unsafe.Pointer(&addressList2[0])))
	return
}

func EnableLocal(board_desc int, addressList []Address) {
	addressList2 := append(addressList, NOADDR)
	C.EnableLocal(C.int(board_desc), (*C.Addr4882_t)(unsafe.Pointer(&addressList2[0])))
	return
}

func EnableRemote(board_desc int, addressList []Address) {
	addressList2 := append(addressList, NOADDR)
	C.EnableRemote(C.int(board_desc), (*C.Addr4882_t)(unsafe.Pointer(&addressList2[0])))
	return
}

func FindLstn(board_desc int, padList []Address) (resultList []Address) {
	padList2 := append(padList, NOADDR)
	resultList = make([]Address, len(padList))
	C.FindLstn(C.int(board_desc), (*C.Addr4882_t)(unsafe.Pointer(&padList2[0])), (*C.Addr4882_t)(unsafe.Pointer(&resultList[0])), C.int(len(padList)))
	return
}

func FindRQS(board_desc int, addressList []Address) (result int) {
	addressList2 := append(addressList, NOADDR)
	C.FindRQS(C.int(board_desc), (*C.Addr4882_t)(unsafe.Pointer(&addressList2[0])), (*C.short)(unsafe.Pointer(&result)))
	return
}

func PassControl(board_desc int, address Address) {
	C.PassControl(C.int(board_desc), C.Addr4882_t(address))
	return
}

func PPoll(board_desc int) (result int) {
	C.PPoll(C.int(board_desc), (*C.short)(unsafe.Pointer(&result)))
	return
}

func PPollConfig(board_desc int, address Address, dataLine int, lineSense int) {
	C.PPollConfig(C.int(board_desc), C.Addr4882_t(address), C.int(dataLine), C.int(lineSense))
	return
}

func PPollUnconfig(board_desc int, addressList []Address) {
	addressList2 := append(addressList, NOADDR)
	C.PPollUnconfig(C.int(board_desc), (*C.Addr4882_t)(unsafe.Pointer(&addressList2[0])))
	return
}

func RcvRespMsg(board_desc int, buffer []byte, termination int) {
	C.RcvRespMsg(C.int(board_desc), unsafe.Pointer(&buffer[0]), C.long(len(buffer)), C.int(termination))
	return
}

func ReadStatusByte(board_desc int, address Address) (result int) {
	C.ReadStatusByte(C.int(board_desc), C.Addr4882_t(address), (*C.short)(unsafe.Pointer(&result)))
	return
}

func Receive(board_desc int, address Address, buffer []byte, termination int) {
	C.Receive(C.int(board_desc), C.Addr4882_t(address), unsafe.Pointer(&buffer[0]), C.long(len(buffer)), C.int(termination))
	return
}

func ReceiveSetup(board_desc int, address Address) {
	C.ReceiveSetup(C.int(board_desc), C.Addr4882_t(address))
	return
}

func ResetSys(board_desc int, addressList []Address) {
	addressList2 := append(addressList, NOADDR)
	C.ResetSys(C.int(board_desc), (*C.Addr4882_t)(unsafe.Pointer(&addressList2[0])))
	return
}

func Send(board_desc int, address Address, buffer []byte, eot_mode int) {
	bufferPtr := C.CBytes(buffer)
	defer C.free(unsafe.Pointer(bufferPtr))
	C.Send(C.int(board_desc), C.Addr4882_t(address), unsafe.Pointer(bufferPtr), C.long(len(buffer)), C.int(eot_mode))
	return
}

func SendCmds(board_desc int, cmds []byte) {
	cmdsPtr := C.CBytes(cmds)
	defer C.free(unsafe.Pointer(cmdsPtr))
	C.SendCmds(C.int(board_desc), unsafe.Pointer(cmdsPtr), C.long(len(cmds)))
	return
}

func SendDataBytes(board_desc int, buffer []byte, eotmode int) {
	bufferPtr := C.CBytes(buffer)
	defer C.free(unsafe.Pointer(bufferPtr))
	C.SendDataBytes(C.int(board_desc), unsafe.Pointer(bufferPtr), C.long(len(buffer)), C.int(eotmode))
	return
}

func SendIFC(board_desc int) {
	C.SendIFC(C.int(board_desc))
	return
}

func SendLLO(board_desc int) {
	C.SendLLO(C.int(board_desc))
	return
}

func SendList(board_desc int, addressList []Address, buffer []byte, eotmode int) {
	addressList2 := append(addressList, NOADDR)
	bufferPtr := C.CBytes(buffer)
	defer C.free(unsafe.Pointer(bufferPtr))
	C.SendList(C.int(board_desc), (*C.Addr4882_t)(unsafe.Pointer(&addressList2[0])), unsafe.Pointer(bufferPtr), C.long(len(buffer)), C.int(eotmode))
	return
}

func SendSetup(board_desc int, addressList []Address) {
	addressList2 := append(addressList, NOADDR)
	C.SendSetup(C.int(board_desc), (*C.Addr4882_t)(unsafe.Pointer(&addressList2[0])))
	return
}

func SetRWLS(board_desc int, addressList []Address) {
	addressList2 := append(addressList, NOADDR)
	C.SetRWLS(C.int(board_desc), (*C.Addr4882_t)(unsafe.Pointer(&addressList2[0])))
	return
}

func TestSRQ(board_desc int) (result int) {
	C.TestSRQ(C.int(board_desc), (*C.short)(unsafe.Pointer(&result)))
	return
}

func TestSys(board_desc int, addressList []Address) (resultList []int) {
	addressList2 := append(addressList, NOADDR)
	resultList = make([]int, len(addressList))
	C.TestSys(C.int(board_desc), (*C.Addr4882_t)(unsafe.Pointer(&addressList2[0])), (*C.short)(unsafe.Pointer(&resultList[0])))
	return
}

func Trigger(board_desc int, address Address) {
	C.Trigger(C.int(board_desc), C.Addr4882_t(address))
	return
}

func TriggerList(board_desc int, addressList []Address) {
	addressList2 := append(addressList, NOADDR)
	C.TriggerList(C.int(board_desc), (*C.Addr4882_t)(unsafe.Pointer(&addressList2[0])))
	return
}

func WaitSRQ(board_desc int) (result int) {
	C.WaitSRQ(C.int(board_desc), (*C.short)(unsafe.Pointer(&result)))
	return
}

func Ibask(ud int, option int) (ibsta int, value int) {
	ibsta = int(C.ibask(C.int(ud), C.int(option), (*C.int)(unsafe.Pointer(&value))))
	return
}

func Ibbna(ud int, board_name string) (ibsta int) {
	board_namePtr := C.CString(board_name)
	defer C.free(unsafe.Pointer(board_namePtr))
	ibsta = int(C.ibbna(C.int(ud), board_namePtr))
	return
}

func Ibcac(ud int, synchronous int) (ibsta int) {
	ibsta = int(C.ibcac(C.int(ud), C.int(synchronous)))
	return
}

func Ibclr(ud int) (ibsta int) {
	ibsta = int(C.ibclr(C.int(ud)))
	return
}

func Ibcmd(ud int, cmd []byte) (ibsta int) {
	cmdPtr := C.CBytes(cmd)
	defer C.free(unsafe.Pointer(cmdPtr))
	ibsta = int(C.ibcmd(C.int(ud), unsafe.Pointer(cmdPtr), C.long(len(cmd))))
	return
}

func Ibcmda(ud int, cmd []byte) (ibsta int) {
	cmdPtr := C.CBytes(cmd)
	defer C.free(unsafe.Pointer(cmdPtr))
	ibsta = int(C.ibcmda(C.int(ud), unsafe.Pointer(cmdPtr), C.long(len(cmd))))
	return
}

func Ibconfig(ud int, option int, value int) (ibsta int) {
	ibsta = int(C.ibconfig(C.int(ud), C.int(option), C.int(value)))
	return
}

func Ibdev(board_index int, pad int, sad int, timo int, send_eoi int, eosmode int) (ud int) {
	ud = int(C.ibdev(C.int(board_index), C.int(pad), C.int(sad), C.int(timo), C.int(send_eoi), C.int(eosmode)))
	return
}

func Ibdma(ud int, v int) (ibsta int) {
	ibsta = int(C.ibdma(C.int(ud), C.int(v)))
	return
}

func Ibeot(ud int, v int) (ibsta int) {
	ibsta = int(C.ibeot(C.int(ud), C.int(v)))
	return
}

func Ibeos(ud int, v int) (ibsta int) {
	ibsta = int(C.ibeos(C.int(ud), C.int(v)))
	return
}

func Ibevent(ud int) (ibsta int, event int) {
	ibsta = int(C.ibevent(C.int(ud), (*C.short)(unsafe.Pointer(&event))))
	return
}

func Ibfind(dev string) (ibsta int) {
	devPtr := C.CString(dev)
	defer C.free(unsafe.Pointer(devPtr))
	ibsta = int(C.ibfind(devPtr))
	return
}

func Ibgts(ud int, shadow_handshake int) (ibsta int) {
	ibsta = int(C.ibgts(C.int(ud), C.int(shadow_handshake)))
	return
}

func Ibist(ud int, ist int) (ibsta int) {
	ibsta = int(C.ibist(C.int(ud), C.int(ist)))
	return
}

func Iblines(ud int) (ibsta int, line_status int) {
	ibsta = int(C.iblines(C.int(ud), (*C.short)(unsafe.Pointer(&line_status))))
	return
}

func Ibln(ud int, pad int, sad int) (ibsta int, found_listener int) {
	ibsta = int(C.ibln(C.int(ud), C.int(pad), C.int(sad), (*C.short)(unsafe.Pointer(&found_listener))))
	return
}

func Ibloc(ud int) (ibsta int) {
	ibsta = int(C.ibloc(C.int(ud)))
	return
}

func Ibonl(ud int, onl int) (ibsta int) {
	ibsta = int(C.ibonl(C.int(ud), C.int(onl)))
	return
}

func Ibpad(ud int, v int) (ibsta int) {
	ibsta = int(C.ibpad(C.int(ud), C.int(v)))
	return
}

func Ibpct(ud int) (ibsta int) {
	ibsta = int(C.ibpct(C.int(ud)))
	return
}

func Ibppc(ud int, v int) (ibsta int) {
	ibsta = int(C.ibppc(C.int(ud), C.int(v)))
	return
}

func Ibrd(ud int, buf []byte) (ibsta int) {
	ibsta = int(C.ibrd(C.int(ud), unsafe.Pointer(&buf[0]), C.long(len(buf))))
	return
}

func Ibrda(ud int, buf []byte) (ibsta int) {
	ibsta = int(C.ibrda(C.int(ud), unsafe.Pointer(&buf[0]), C.long(len(buf))))
	return
}

func Ibrdf(ud int, file_path string) (ibsta int) {
	file_pathPtr := C.CString(file_path)
	defer C.free(unsafe.Pointer(file_pathPtr))
	ibsta = int(C.ibrdf(C.int(ud), file_pathPtr))
	return
}

func Ibrpp(ud int) (ibsta int, ppr byte) {
	ibsta = int(C.ibrpp(C.int(ud), (*C.char)(unsafe.Pointer(&ppr))))
	return
}

func Ibrsc(ud int, v int) (ibsta int) {
	ibsta = int(C.ibrsc(C.int(ud), C.int(v)))
	return
}

func Ibrsp(ud int) (ibsta int, spr byte) {
	ibsta = int(C.ibrsp(C.int(ud), (*C.char)(unsafe.Pointer(&spr))))
	return
}

func Ibrsv(ud int, v int) (ibsta int) {
	ibsta = int(C.ibrsv(C.int(ud), C.int(v)))
	return
}

func Ibsad(ud int, v int) (ibsta int) {
	ibsta = int(C.ibsad(C.int(ud), C.int(v)))
	return
}

func Ibsic(ud int) (ibsta int) {
	ibsta = int(C.ibsic(C.int(ud)))
	return
}

func Ibspb(ud int) (ibsta int, sp_bytes int) {
	ibsta = int(C.ibspb(C.int(ud), (*C.short)(unsafe.Pointer(&sp_bytes))))
	return
}

func Ibsre(ud int, v int) (ibsta int) {
	ibsta = int(C.ibsre(C.int(ud), C.int(v)))
	return
}

func Ibstop(ud int) (ibsta int) {
	ibsta = int(C.ibstop(C.int(ud)))
	return
}

func Ibtmo(ud int, v int) (ibsta int) {
	ibsta = int(C.ibtmo(C.int(ud), C.int(v)))
	return
}

func Ibtrg(ud int) (ibsta int) {
	ibsta = int(C.ibtrg(C.int(ud)))
	return
}

func Ibvers() (version string) {
	var versionPtr *C.char
	C.ibvers((**C.char)(unsafe.Pointer(&versionPtr)))
	version = C.GoString(versionPtr)
	return
}

func Ibwait(ud int, mask int) (ibsta int) {
	ibsta = int(C.ibwait(C.int(ud), C.int(mask)))
	return
}

func Ibwrt(ud int, buf []byte) (ibsta int) {
	bufPtr := C.CBytes(buf)
	defer C.free(unsafe.Pointer(bufPtr))
	ibsta = int(C.ibwrt(C.int(ud), unsafe.Pointer(bufPtr), C.long(len(buf))))
	return
}

func Ibwrta(ud int, buf []byte) (ibsta int) {
	bufPtr := C.CBytes(buf)
	defer C.free(unsafe.Pointer(bufPtr))
	ibsta = int(C.ibwrta(C.int(ud), unsafe.Pointer(bufPtr), C.long(len(buf))))
	return
}

func Ibwrtf(ud int, file_path string) (ibsta int) {
	file_pathPtr := C.CString(file_path)
	defer C.free(unsafe.Pointer(file_pathPtr))
	ibsta = int(C.ibwrtf(C.int(ud), file_pathPtr))
	return
}
