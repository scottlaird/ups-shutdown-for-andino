# ups-shutdown-for-andino

ups-shutdown-for-andino is a utility for shutting down an [Andino
X1](https://andino.systems/andino-x1/) when the [Andino
UPS](https://andino.systems/andino-ups-uninterruptible-power-supply/) signals
low power.

The Andino X1 is a Raspberry Pi in a industrial case, with a 24 VDC
power supply and an embedded Arduino.  The Andino UPS is a small 24
VDC uninteruptable power supply that can power the Andino X1 for a
minute or two after power loss, and signal when it's running low on
power.  This code reads that low power signal and triggers a clean
shutdown.  That means that you can ignore the Raspberry Pi in the
Andino X1 and simply power the entire enclosure up and down as needed,
knowing that it'll shut down cleanly when power is lost.

To wire this up, follow [Andino's directions on
GitHub](https://github.com/andino-systems/Andino-UPS).  Basically,
wire the "Power Fail Relay" on the UPS to Digital Input 1 on the
Andino X1, following the ground and power wiring as shown.  Then
install [Andino's
drivers](https://github.com/andino-systems/Andino-X1/tree/master/doc/BaseBoard)
inside of Raspian, so that you can talk to their Arduino on
`/dev/ttyAMA0`.

Once that is complete, compile `ups-shutdown-for-andino` by running
`go build` and install it in `/sbin`.  For systems that use `systemd`
(like Raspian and Ubuntu), copy `ups-shutdown-for-andino.service` to
`/etc/systemd/system/`.  Run `systemctl enable
ups-shutdown-for-andino` to enable to start on reboot, and `systemctl
start ups-shutdown-for-andino` to start it immediately.

### Disclaimer

This is not an officially supported Google product.

