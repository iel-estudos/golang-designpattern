package main

import "fmt"

type SohneeTV struct {
	vol     int
	channel int
	isOn    bool
}

func (st *SohneeTV) turnOn() {
	fmt.Println("SohneeTV is now on")
	st.isOn = true
}

func (st *SohneeTV) turnOff() {
	fmt.Println("SohneeTV is now off")
	st.isOn = false
}

func (st *SohneeTV) volumeUp() int {
	st.vol++
	fmt.Println("Increasing SohneeTV volume to", st.vol)
	return st.vol
}

func (st *SohneeTV) volumeDown() int {
	st.vol--
	fmt.Println("Increasing SohneeTV volume to", st.vol)
	return st.vol
}

func (st *SohneeTV) channelUp() int {
	st.channel++
	fmt.Println("Increasing SohneeTV channel to", st.vol)
	return st.vol
}

func (st *SohneeTV) channelDown() int {
	st.channel--
	fmt.Println("Increasing SohneeTV channel to", st.vol)
	return st.vol
}