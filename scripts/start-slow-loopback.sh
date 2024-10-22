#!/usr/bin/env bash

#tc qdisc add dev lo root netem delay 10ms rate 100Mbit
tc qdisc change dev lo root netem delay 10ms rate 1000Mbit
