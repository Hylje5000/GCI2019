from shodan import Shodan
import os
import socket

apikey = input("Input your Shodan API-key")
api = Shodan(apikey)
#titlebar
def titlebar():
    os.system('clear')
              
    print("**********************************************")
    print("  Shodan Scanner for GCI 2019 - MiskaKyto  ")
    print("**********************************************")
# Main menu
def getchoice():
    print("1 What is my IP")
    print("2 Scan specified host")
    print("3 Shodan search")
    print("0 Exit")
    return input()

def getIP():
    print("Your IP address is: ")
    print([l for l in ([ip for ip in socket.gethostbyname_ex(socket.gethostname())[2] if not ip.startswith("127.")][:1], [[(s.connect(('8.8.8.8', 53)), s.getsockname()[0], s.close()) for s in [socket.socket(socket.AF_INET, socket.SOCK_DGRAM)]][0][1]]) if l][0][0])
    print("\n")
    print('*' * 15)

def scanHost():
    ip = input("Input an IP address: ")
    host = api.host(ip)
    #List Ip, Org and Operating System
    print("""
        IP: {}
        Organization: {}
        Operating System: {}
""".format(host['ip_str'], host.get('org', 'n/a'), host.get('os', 'n/a')))

# List open ports
    for item in host['data']:
            print("""       Port: {}""".format(item['port']))
    print("\n")

def shodanSearch():
    searchterm = input("Input search term: ")
    try:
        # Search Shodan
        results = api.search(searchterm)

        # Show the results
        print('Results found: {}'.format(results['total']))
        for result in results['matches']:
                print('IP: {}'.format(result['ip_str']))
                print('')
    except shodan.APIError:
            print('Error: {}'.format(e))
#Main program
answer = ''
titlebar()
while answer != 'q':
    answer = getchoice()
    titlebar()
    if answer == "1":
        getIP()
    elif answer == "2":
        scanHost()
    elif answer == "3":
        shodanSearch()
    if answer == "0":
        print("Exiting app...")
        quit()



