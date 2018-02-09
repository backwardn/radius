// Generated by radius-dict-gen. DO NOT EDIT.

package rfc2867

import (
	"net"
	"strconv"
	"time"

	"layeh.com/radius"

	. "layeh.com/radius/rfc2866"
)

var _ = radius.Type(0)
var _ = strconv.Itoa
var _ = net.ParseIP
var _ = time.Time{}

const (
	AcctTunnelConnection_Type  radius.Type = 68
	AcctTunnelPacketsLost_Type radius.Type = 86
)

func init() {
	AcctStatusType_Strings[AcctStatusType_Value_TunnelStart] = "Tunnel-Start"
	AcctStatusType_Strings[AcctStatusType_Value_TunnelStop] = "Tunnel-Stop"
	AcctStatusType_Strings[AcctStatusType_Value_TunnelReject] = "Tunnel-Reject"
	AcctStatusType_Strings[AcctStatusType_Value_TunnelLinkStart] = "Tunnel-Link-Start"
	AcctStatusType_Strings[AcctStatusType_Value_TunnelLinkStop] = "Tunnel-Link-Stop"
	AcctStatusType_Strings[AcctStatusType_Value_TunnelLinkReject] = "Tunnel-Link-Reject"
}

const (
	AcctStatusType_Value_TunnelStart      AcctStatusType = 9
	AcctStatusType_Value_TunnelStop       AcctStatusType = 10
	AcctStatusType_Value_TunnelReject     AcctStatusType = 11
	AcctStatusType_Value_TunnelLinkStart  AcctStatusType = 12
	AcctStatusType_Value_TunnelLinkStop   AcctStatusType = 13
	AcctStatusType_Value_TunnelLinkReject AcctStatusType = 14
)

func AcctTunnelConnection_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(AcctTunnelConnection_Type, a)
	return nil
}

func AcctTunnelConnection_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(AcctTunnelConnection_Type, a)
	return nil
}

func AcctTunnelConnection_Get(p *radius.Packet) (value []byte) {
	value, _ = AcctTunnelConnection_Lookup(p)
	return
}

func AcctTunnelConnection_GetString(p *radius.Packet) (value string) {
	return string(AcctTunnelConnection_Get(p))
}

func AcctTunnelConnection_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[AcctTunnelConnection_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AcctTunnelConnection_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[AcctTunnelConnection_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AcctTunnelConnection_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(AcctTunnelConnection_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func AcctTunnelConnection_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(AcctTunnelConnection_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func AcctTunnelConnection_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(AcctTunnelConnection_Type, a)
	return
}

func AcctTunnelConnection_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(AcctTunnelConnection_Type, a)
	return
}

type AcctTunnelPacketsLost uint32

var AcctTunnelPacketsLost_Strings = map[AcctTunnelPacketsLost]string{}

func (a AcctTunnelPacketsLost) String() string {
	if str, ok := AcctTunnelPacketsLost_Strings[a]; ok {
		return str
	}
	return "AcctTunnelPacketsLost(" + strconv.Itoa(int(a)) + ")"
}

func AcctTunnelPacketsLost_Add(p *radius.Packet, value AcctTunnelPacketsLost) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctTunnelPacketsLost_Type, a)
	return nil
}

func AcctTunnelPacketsLost_Get(p *radius.Packet) (value AcctTunnelPacketsLost) {
	value, _ = AcctTunnelPacketsLost_Lookup(p)
	return
}

func AcctTunnelPacketsLost_Gets(p *radius.Packet) (values []AcctTunnelPacketsLost, err error) {
	var i uint32
	for _, attr := range p.Attributes[AcctTunnelPacketsLost_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctTunnelPacketsLost(i))
	}
	return
}

func AcctTunnelPacketsLost_Lookup(p *radius.Packet) (value AcctTunnelPacketsLost, err error) {
	a, ok := p.Lookup(AcctTunnelPacketsLost_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctTunnelPacketsLost(i)
	return
}

func AcctTunnelPacketsLost_Set(p *radius.Packet, value AcctTunnelPacketsLost) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctTunnelPacketsLost_Type, a)
	return nil
}
