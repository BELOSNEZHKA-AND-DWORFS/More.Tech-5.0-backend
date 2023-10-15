#! venv/bin/python3

import random
import json
import time
import schedule
import numpy as np
from datetime import datetime

minutes_diff = 1
def get_week_day(timestamp):
    return time.strftime("%a", time.gmtime(timestamp))

def generate_wait_time():
    return int(np.random.poisson(1200))  # Время ожидания в секундах (от 0 до 3600)

def get_daytime_period(timestamp):
    return np.clip(((time.gmtime(timestamp).tm_hour+3) - 8) // 2, 1, 6)

def generate_event_packet():
    timestamp_now = int(datetime.timestamp(datetime.now()))
    logs = []
    for department_id in range(1, 278):
        num_events = random.randint(0, 8)
        for _ in range(num_events):
            wait_time = generate_wait_time()
            timestamp =  timestamp_now - random.randint(0, minutes_diff*60) - wait_time
            log = {
                "timestamp": timestamp,
                "id": department_id,
                "week_day": get_week_day(timestamp),
                "queue_len": random.randint(0, 35),
                "service": random.choice([
                    "takemoney", "putmoney", "userubles", "usedollars", "usereuros",
                    "disabledpersons", "nfc", "qr", "biometry", "investory", "credit",
                    "vtbprime", "vtbprivilege", "ramp"
                ]),
                "day_time": get_daytime_period(timestamp),
                "wait_time": wait_time
            }
            logs.append(log)
    with open('logs.json', 'r') as json_file:
        existing_logs = json.load(json_file)
    existing_logs.extend(logs)
    with open('logs.json', 'w') as json_file:
        json.dump(logs, json_file)

# Запуск генерации событий каждые 5 минут
schedule.every(minutes_diff).minutes.do(generate_event_packet)

def run_schedule():
    while True:
        schedule.run_pending()
        time.sleep(1)

if __name__ == "__main__":
    run_schedule()
