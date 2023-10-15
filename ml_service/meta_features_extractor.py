#! venv/bin/python3

import json
import pandas as pd
import numpy as np
from sklearn.preprocessing import LabelEncoder
from catboost import CatBoostRegressor
from datetime import datetime
from realtime_logs_collector import get_daytime_period, get_week_day

departments_count = 279

def wait_time_prediction(data_path):
    model = CatBoostRegressor()
    model.load_model('./models/wait_time_predictor.cbm')

    with open( "./logs.json", "r", encoding = "utf-8") as data:
        data = json.load(data)
    
    full_data = pd.DataFrame(data)
    label_encoder = LabelEncoder()
    full_data['service'] = label_encoder.fit_transform(full_data['service'])
    full_data['week_day'] = label_encoder.fit_transform(full_data['week_day'])
    avg_queu_len_total = full_data["queue_len"].mean()

    timestamp_now = int(datetime.timestamp(datetime.now()))
    predictions = pd.DataFrame(list(range(1,departments_count)), columns=["id"])
    predicted_wait_time = []

    for department_id in range(1, departments_count):
        timestamp_threshold = timestamp_now-30*60
        fresh_data = full_data[(full_data["timestamp"]>timestamp_threshold) & (full_data["id"] == department_id)]
        if len(fresh_data["queue_len"]):
            avg_queu_len = fresh_data["queue_len"].mean() 
            mode_day_time = fresh_data["day_time"].mode()[0]
            mode_week_day = fresh_data["week_day"].mode()[0].astype(int)
            mode_service = fresh_data["service"].mode()[0].astype(int)
        else:
            avg_queu_len = avg_queu_len_total
            mode_day_time = get_daytime_period(timestamp_now)
            mode_week_day = label_encoder.transform([get_week_day(timestamp_now)])[0].astype(int)
            mode_service = ""
        predicted_wait_time.append(model.predict(np.array([mode_week_day, avg_queu_len, mode_service, mode_day_time])))
    predictions["predicted_wait_time"] = predicted_wait_time
    
    # прокинуть в табличку с метафичой
    return predictions

# Запуск генерации событий каждые 15 минут
schedule.every(1).minutes.do(wait_time_prediction)

def run_schedule():
    while True:
        schedule.run_pending()
        time.sleep(1)

if __name__ == "__main__":
    predictions = wait_time_prediction()
    





