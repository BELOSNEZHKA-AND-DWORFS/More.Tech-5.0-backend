import json
from random import randint
from random import uniform

allTypesList = ["office", "atm"]

atm_services = ["isVtb", "supportsRub", "supportsUsd", "supportsEur", "serviceActivity", "allDay", "weelchair",
                    "blind", "nfcForBankCards", "qrRead", "supportsChargeRub"]
office_services = ["isOrganization", "isBiometric", "isInvestment", "isUniversal", "isCashOffice",
                   "isAvailableNow", "isAvailableWeekends", "isAvailableToday", "isFreeOffice",
                   "isRamp", "isCurrency", "isVtbPrime", "isVtbPrivilege", "supportsChargeRub"]

allWorking_time_fl = ["mo-fr 9:00-18:00, sa 9:00-16:00",
                      "mo-fr 9:00-14:00",
                      "mo-fr 9:00-18:00",
                      "mo-fr 9:00-18:00, sa-su 9:00-16:00",
                      "mo-sa 9:00-18:00",
                      "mo-su 9:00-18:00",
                      "mo-sa 9:00-18:00, su 9:00-17:00",
                      "mo-sa 0:00-23:59"]
allWorking_time_ul = allWorking_time_fl

id = 0

def getLatitude():
    return 0

def getLongitude():
    return 0

def generate_random_data(obj, i):
    global id
    objType = allTypesList[randint(0, 1)]
    services = set()
    allServices = office_services if objType == "office" else atm_services
    for i in range(randint(0, len(allServices) - 1)):
        services.add(allServices[randint(0, len(allServices) - 1)])

    if len(services) == 0:
        services.add(allServices[3])

    latitude = obj[i]["latitude"]
    longitude = obj[i]["longitude"]

    working_time_fl = allWorking_time_fl[randint(0, len(allWorking_time_fl) - 1)]
    working_time_ul = allWorking_time_fl[randint(0, len(allWorking_time_ul) - 1)]

    dataList = { "id": id, "type": objType, "services": list(services), 
     "location": (latitude, longitude), 
     "working_time_fl": working_time_fl,
     "working_time_ul": working_time_ul }
    id += 1

    return json.dumps(dataList)

def gen_all_data(count, objects):
    result = []
    for i in range(count):
        result.append(generate_random_data(objects, i))
    return result

with open("data.txt", encoding='utf-8') as file:
    objects = json.load(file)

data = gen_all_data(278, objects)

with open("file.txt", "w") as file:
    file.write("[")
    for line in data:
        file.write(line + ",\n")
    file.write("]")
