from faker import Faker
from datetime import datetime
import pandas as pd
import numpy as np
from gen_utils import *

fake = Faker()

nrow = 100000

df = pd.DataFrame()

df['id'] = range(1, nrow+1)
df['office_id'] = np.random.randint(1, 278, nrow)
#df['user_id'] = 12
df['queue'] = np.random.randint(1, 12, nrow)


df['datetime'] = [fake.date_time_between(
    start_date=datetime(2023, 9, 13),
    end_date=datetime(2023, 10, 13)
    ).strftime('%Y-%m-%d %I:%M')
    for _ in range(nrow)]

df.to_csv("services/predict/data/queue.csv", index=False)



np.random.seed(12345)
nrow = 6000

np.set_printoptions(precision=2)

df2 = pd.DataFrame()

df2['id'] = range(1, nrow+1)
df2['user_id'] = np.random.randint(1, 20, nrow)
df2['office_id'] = np.random.randint(1, 278, nrow)
df2['rating'] = np.random.randint(1, 5, nrow)


df2['datetime'] = [fake.date_time_between(
    start_date=datetime(2023, 9, 13),
    end_date=datetime(2023, 10, 13)
    ).strftime('%m/%d/%Y %I:%M ')
    for _ in range(nrow)]

df2.to_csv("services/predict/data/offices_ratings.csv", index=False)


