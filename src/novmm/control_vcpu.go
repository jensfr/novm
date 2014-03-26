package main

import (
    "syscall"
)

type VcpuSettings struct {
    // Which vcpu?
    Id  int `json:"id"`

    // Single stepping?
    Step bool `json:"step"`

    // Paused?
    Paused bool `json:"paused"`
}

func (control *Control) Vcpu(settings *VcpuSettings, ok *bool) error {
    // A valid vcpu?
    vcpus := control.vm.GetVcpus()
    if settings.Id >= len(vcpus) {
        *ok = false
        return syscall.EINVAL
    }

    // Grab our specific vcpu.
    vcpu := vcpus[settings.Id]

    // Ensure steping is as expected.
    err := vcpu.SetStepping(settings.Step)
    if err != nil {
        *ok = false
        return err
    }

    // Ensure that the vcpu is paused/unpaused.
    if settings.Paused {
        err = vcpu.Pause(true)
    } else {
        err = vcpu.Unpause(true)
    }

    // Done.
    *ok = (err == nil)
    return err
}
