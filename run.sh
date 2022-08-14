#!/bin/bash

go build -o bookings cmd/web/*.go
./bookings -dbname=booking -dbuser=root -cache=false -production=false -dbpass=root