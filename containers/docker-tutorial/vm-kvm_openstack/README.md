# Openstack on KVM

1. Create Linux brudge from physical interface:
    Create a bridge from physical interface, **enp...**, to let virtual OpenStack nodes in
    KVM communicate with external networks.
    The bridge will be attached to KVM OpenStack public network **pub_net**.
    Pub_net will be attached to the first interface, **eth0**, on each virtual
    node.

    ```
    virt-manager
    ```
    Right click QEMU -> Details -> Network Interfaces -> Add Interfaces ->
    Bridge -> Forward:
    ```
    Interface Name: br0enp0
    Start Mode: None
    Active Now: Enable
    IP Settings: 
    Bridge Settings: STP on, delay 0.00 sec
    Interface to Bridge: ecp0
    ```

2. Create an isolated virtual network
    Create a network isolated from the outside world, whose purpose is to
    provide physical connection (carrier, ISO/OSI Layer 2: Data Link) between
    virtual nodes.
    This network will be attached to the second interface (eth1) on each
    virtual node to provide an exchange of messages.

    QEMU -> Details -> Virtual Networks -> Add Network

    ```
    Network Name: openstack_net0
    Enable DHCPv4: Disable
    EnableStatic Route Definition: Disable
    IPv6 Space Definition: Disable
    Connection to Physical Network: Isolated Virtual Network
    Enable IPv6 Internal Routing/Networking: Disable
    DNS DOmain Name: openstack_net0
    ```
    It will also create a virtual bridge - look at the Device field.


# Useful Commands
```
sudo lshw -C network
```
