#! venv/bin/python3

import json
import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import LabelEncoder
from catboost import CatBoostRegressor, Pool

with open('../logs.json', 'r') as json_file:
    data = json.load(json_file)

df = pd.DataFrame(data)

features = ["week_day", "queue_len", "service", "day_time"]
target = "wait_time"
label_encoder = LabelEncoder()

df['week_day'] = label_encoder.fit_transform(df['week_day'])
df['service'] = label_encoder.fit_transform(df['service'])
X = df[features]
y = df[target]

X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)


train_pool = Pool(X_train, label=y_train, cat_features=["week_day", "service"])
eval_set = Pool(X_test, label=y_test, cat_features=["week_day", "service"])

model = CatBoostRegressor(iterations=500, depth=6, learning_rate=0.1, loss_function='RMSE', cat_features=["week_day", "service"])
model.fit(train_pool, eval_set=eval_set, verbose=100)
model.save_model('../models/wait_time_predictor.cbm')