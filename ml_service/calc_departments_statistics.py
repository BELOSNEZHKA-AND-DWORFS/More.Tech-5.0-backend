#! venv/bin/python3

import pandas as pd
import numpy as np
import json

with open( "./logs.json", "r", encoding = "utf-8") as data:
    data = json.load(data)
    
df = pd.DataFrame(data)
df['timestamp'] = pd.to_datetime(df['timestamp'], unit='s')

end_date = df['timestamp'].max()
start_date = end_date - pd.DateOffset(weeks=2)

filtered_data = df[(df['timestamp'] >= start_date) & (df['timestamp'] <= end_date)]

grouped_data = filtered_data.groupby(['id', 'day_time'])

result = grouped_data.agg({'wait_time': 'mean', 'id': 'count'})
result = result.rename(columns={'wait_time': 'avg_wait_time', 'id': 'attendance'})
print(result.to_json())