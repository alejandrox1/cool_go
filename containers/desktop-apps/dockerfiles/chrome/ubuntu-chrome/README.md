# Chrome on Docker


* To make the microphone work chose the `built-in audio analog stereo`

* Other things to try? Try apt-get-ing `paprefs` and on it do `Network Server >
 [X] Enable network access to local sound devices`.

* ```
  docker run \
  -e PULSE_SERVER=tcp:$(hostname -i):4713 \
  -e PULSE_COOKIE=/run/pulse/cookie \
  -v ~/.config/pulse/cookie:/run/pulse/cookie \
  ...
  ```
