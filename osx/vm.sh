#! /bin/sh

declare -a vms=("Deadpool" "Wolverine")

for vm in "${vms[@]}"
do
    VBoxManage controlvm $vm poweroff
    VBoxManage unregistervm $vm --delete
    VBoxManage import "$OVA_FILE" --vsys 0 --vmname $vm
    VBoxManage startvm $vm
done