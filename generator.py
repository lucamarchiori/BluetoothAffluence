import sqlite3
import random
import datetime

def scan(d):
    scanner = {
        "address": "00:00:00:00:00:00",
        "alias": "Raspberry Pi 3",
        "name": "Raspberry Pi 3"
    }

    # Insert scanner into database if it doesn't exist
    cursor.execute("SELECT * FROM scanners WHERE address = ?", (scanner["address"],))
    if(cursor.fetchone() is None):
        cursor.execute("INSERT INTO scanners (address, alias, name) VALUES (?, ?, ?)", (scanner["address"], scanner["alias"], scanner["name"]))
        conn.commit()
    
    # Get scanner id
    cursor.execute("SELECT id FROM scanners WHERE address = ?", (scanner["address"],))
    scannerId = cursor.fetchone()[0]


    scan = {
        "scanTime": d,
        "scannerId": scannerId,
        "scanTime": d.strftime("%Y-%m-%d %H:%M:%S")
    }

    # Insert scan into database
    cursor.execute("INSERT INTO scan (scanTime, scannerId) VALUES (?, ?)", (scan["scanTime"], scan["scannerId"]))
    conn.commit()

    # Get scan id
    cursor.execute("SELECT id FROM scan WHERE scanTime = ?", (scan["scanTime"],))
    scanId = cursor.fetchone()[0]


    devices = []
    min = 0
    max = 0

    # If datetime is between 19:00 and 08:00 generate a random number of devices between 0 and 3
    if(d.hour >= 19 or d.hour <= 8):
        min = 0 
        max = 3

    # If datetime is between 18:00 and 19:00 or datetime is between 8:00 and 9:00 generate a random number of devices between 1 and 5
    elif((d.hour >= 18 and d.hour < 19) or (d.hour >= 8 and d.hour < 9)):
        min = 1
        max = 5
    
    # If datetime is between 9:00 and 11:00 or datetime is between 14:00 and 18:00 generate a random number of devices between 8 and 15
    elif((d.hour >= 9 and d.hour < 11) or (d.hour >= 14 and d.hour < 18)):
        min = 8
        max = 15
        # If datetime is between 12:00 and 14:00 generate a random number of devices between 8 and 10
    elif(d.hour >= 12 and d.hour < 14):
        min = 8
        max = 10

    for i in range(random.randint(min, max)):
        # select a random mac address, alias, name, txPower, rssi and scanID
        mac = random.choice(macs)
        txPower = random.randint(-100, -20)
        rssi = random.randint(-100, -20)

        device = {
            "address": mac,
            "txPower": txPower,
            "rssi": rssi,
            "scanId": scanId
        }
        devices.append(device)
        print(device)

    # Insert devices into database
    for device in devices:
        cursor.execute("INSERT INTO devices (address, txPower, rssi, scanId) VALUES (?, ?, ?, ?)", (device["address"], device["txPower"], device["rssi"], device["scanId"]))
        conn.commit()


# Create a connection to the database
conn = sqlite3.connect('ServerDB.db')
cursor = conn.cursor()

# Create the tables
cursor.execute(("CREATE TABLE IF NOT EXISTS devices (id INTEGER PRIMARY KEY, address TEXT, alias TEXT, name TEXT, txPower REAL, rssi REAL, scanID INTEGER)"))
cursor.execute("CREATE TABLE IF NOT EXISTS scanners (id INTEGER PRIMARY KEY, address TEXT, alias TEXT, name TEXT, UNIQUE(address))")
cursor.execute("CREATE TABLE IF NOT EXISTS scan (id INTEGER PRIMARY KEY, scanTime TEXT, scannerId INTEGER)")
conn.commit()

# Generate 100 mac addresses
macs = []
for i in range(100):
    mac = [random.randint(0x00, 0xff) for _ in range(6)]
    mac = ':'.join(map(lambda x: "%02x" % x, mac))
    macs.append(mac)


dstart = datetime.datetime(2023,12,12,0,0,0,0)
dend = datetime.datetime(2023,12,20,0,0,0,0)

# Run scan every 5 minutes between start and end
while(dstart < dend):
    scan(dstart)
    dstart += datetime.timedelta(minutes=5)

conn.close()