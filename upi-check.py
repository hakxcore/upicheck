import requests
target=input("\n[+] Enter the target Phone Number or UPI ID:')
ids=["apl", "ybl", "oksbi", "okhdfcbank", "axl", "paytm","ibl", "upi", "icici", "sbi", "kotak", "postbank", "axisbank", "okicici", "okaxis", "dbs", "barodampay", "idfcbank"]
url="https://upibankvalidator.com/api/upiValidation?upi="
print(")
if target.find('@')!=-1:
response=requests.post(url+target).json()
if response['isUpiRegistered']: print("[+]UPI ID registered with name " + response['name'])
else:
print("[-] UPI ID not registered!")
else:
for i in ids:
response=requests.post(url+target+'@'+i).json() if response['isUpiRegistered"]:
print("[+] UPI ID registered with name " + response['name']+" with`'+i)
else:
print("[-]UPI ID not registered with "+i)
