generate:
	swagger generate client --target=./netbox --spec=./swagger.json --copyright-file=./copyright_header.txt
	@ag bool |grep ".go" |awk -F":" '{print $$1}'  |sort |uniq|while read line; do sed -e 's# *bool `json:# *bool `json:#g' -i  $$line;done
	echo " ==> There are still few bugs to correct"  

clean:
	rm -rf netbox/client netbox/models

integration:
	go test ./... -tags=integration

