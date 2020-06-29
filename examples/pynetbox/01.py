#!/usr/bin/env python
"""
Example (in your terminal):
    $./netbox.py -nm 24 -e dev -t default carve
    $./netbox.py -nm 28 -e prod -t americas carve
    $./netbox.py -nm 22 -e prod -t non_routable carve
    $./netbox.py -pfx 1.2.3.4/24 delete

"""

import argparse, pynetbox, os, sys, urllib3
from prettytable import PrettyTable

NETBOX_MAP = {
    "dev": {
        "default"       : 0,
        "non_routable"  : 9182,
    },
    "corp": {
        "default"      : 9058,
        "non_routable" : 9116,
    },
    "prod": {
        "non_routable" : 9115,
        "americas"     : 7850,
        "europe"       : 7851,
        "asia"         : 7852,
    }
}

def types_help():
    text = "\n"
    for env, types in NETBOX_MAP.items():
        text += f"{env}: {list(types.keys())}\n"
    return text

def cmdline_args(nb):
    FUNCTION_MAP = {'carve'  : carve_block,
                    'update' : update_block,
                    'delete' : delete_block }

    p = argparse.ArgumentParser(description=__doc__,
        formatter_class=argparse.RawDescriptionHelpFormatter)

    p.add_argument('-nm', '--netmask', action="store", required=True,
                    help="The netmask to be used (i.e 24)")
    p.add_argument('-e','--env', action="store", choices=NETBOX_MAP.keys(), required=True,
                    help="The enviornment to operate on CHOICES: %(choices)s")
    p.add_argument('-pfx','--prefix', action="store",
                    help="The prefix to be deleted (i.e 1.2.3.4/8)")
    p.add_argument('-t','--type', action="store", help=f"The type of subnet to carve CHOICES: {types_help()}")
    p.add_argument('command', choices=FUNCTION_MAP.keys())



    args = p.parse_args()

    # Check for a valid types
    valid_types = NETBOX_MAP[args.env].keys()
    if not args.env or (args.type not in valid_types) :
        p.print_help()
        print()
        print()
        print(f"Type is a required parameter, possible values: {list(valid_types)}", file=sys.stderr)
        sys.exit(1)

    func = FUNCTION_MAP[args.command]
    func(nb, args)

def carve_block(nb, args):
    pt = PrettyTable()
    pt.field_names = ["Available Blocks"]
    prefix = nb.ipam.prefixes.get(NETBOX_MAP[args.env][args.type])
    for avail_pref in prefix.available_prefixes.list():
        pt.add_row([avail_pref['prefix']])

    print(pt)
    print("By proceeding, you will reserve a /%s block" % args.netmask)
    answer = input('Do you want to continue? [y/n]: ')
    if not answer or answer[0].lower() != 'y':
        exit(1)

    descr = input('Subnet Description (i.e project_name || subnet name): ')
    new_prefix = prefix.available_prefixes.create({
        'prefix_length': int(args.netmask),
        'description' : descr
        }
    )
    print("Successfully reserved new prefix %s" %new_prefix['prefix'])

def update_block():
    pass


def delete_block(nb, args):
    ip = args.prefix.split("/")[0]
    net_mask = args.prefix.split("/")[1]
    record = nb.ipam.prefixes.get(q=ip, mask_length=net_mask)

    if record != None:
        print("By proceeding, you will DELETE prefix [%s]" %args.prefix)
        answer = input('Do you want to continue? [y/n]: ')
        if not answer or answer[0].lower() != 'y':
            exit(1)

        if record.delete():
            print("Successfully deleted prefix %s" %args.prefix)
    else:
        print("Unable to locate prefix, double check the prefix.")

if __name__ == '__main__':

    MIN_PYTHON = (3, 0)
    if sys.version_info < MIN_PYTHON:
        sys.exit("Python %s.%s or later is required.\n" % MIN_PYTHON)

    urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

    try:
        nb_api_token = os.environ['NETBOX_API_TOKEN']
    except KeyError:
        print ("Netbox token missing. Please set it. 'export NETBOX_API_TOKEN=12345'")
        exit(1)

    nb = pynetbox.api(
        token=nb_api_token,
       # ssl_verify=False,
      #  url='https://se1-network-netbox01.se1.cloud.blizzard.net'
        url='http://netbox.k8s.me'
    )

    cmdline_args(nb)
