generate:
	swagger generate client --target=./netbox --spec=./swagger.json --copyright-file=./copyright_header.txt
	@ag bool -G ".go" |egrep "netbox.*.go:"|awk -F":" '{print $$1}'  |sort |uniq|grep -Ev "writable_circuit_termination.go|writable_console_port.go|writable_console_server_port.go|writable_device_interface.go|writable_power_outlet.go|writable_power_port.go"|while read line; do sed -e 's# bool `json:# *bool `json:#g' -i  $$line;done
	#echo " ==> There are still few bugs to correct"  

clean:
	rm -rf netbox/client netbox/models

integration:
	go test ./... -tags=integration

