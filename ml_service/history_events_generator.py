import random
import json
import time
import numpy as np

def generate_timestamp():
    return random.randint(1685566800, 1696107600)

def get_week_day(timestamp):
    return time.strftime("%a", time.gmtime(timestamp))

def get_daytime_period(timestamp):
    return np.clip(((time.gmtime(timestamp).tm_hour+3) - 8) // 2, 1, 6)

def generate_wait_time():
    return int(np.random.poisson(3600))  # Время ожидания в секундах (от 0 до 3600)

def generate_logs(log_count):
    logs = []


    for _ in range(log_count):
        timestamp = generate_timestamp()
        log = {
            "timestamp": timestamp,
            "id": random.randint(1, 50),
            "week_day": get_week_day(timestamp),
            "queue_len": random.randint(0, 35),
            "service": random.choice([
                "takemoney", "putmoney", "userubles", "usedollars", "usereuros",
                "disabledpersons", "nfc", "qr", "biometry", "investory", "credit",
                "vtbprime", "vtbprivilege", "ramp"
            ]),
            "day_time": get_daytime_period(timestamp),
            "wait_time": generate_wait_time()
        }
        logs.append(log)

    return logs

# Генерация и сохранение JSON-логов
log_count = 5000  # желаемое количество логов
logs = generate_logs(log_count)

with open("logs.json", "w") as file:
    json.dump(logs, file, indent=2)