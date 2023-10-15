# эти библиотеки нам уже знакомы
import pandas as pd
import matplotlib.pyplot as plt 
import numpy as np
 
queue = pd.read_csv('services/predict/data/queue.csv')

# превратим дату в индекс и сделаем изменение постоянным
queue.set_index('datetime', inplace = True)

# превратим дату (наш индекс) в объект datetime
queue.index = pd.to_datetime(queue.index)
print(queue)

# обучающая выборка будет включать данные до 2023-09 года включительно
train = queue[:'2023-09']
 
# тестовая выборка начнется с 2023-10 года (по сути, один месяц)
test = queue['2023-10':]

# принудительно отключим предупреждения системы
import warnings
warnings.simplefilter(action = 'ignore', category = Warning)
 
 # обучим модель с соответствующими параметрами, SARIMAX(3, 0, 0)x(0, 1, 0, 12)
# импортируем класс модели
from statsmodels.tsa.statespace.sarimax import SARIMAX
 
# создадим объект этой модели
model = SARIMAX(train, 
                order = (3, 0, 0), 
                seasonal_order = (0, 1, 0, 12))
 
# применим метод fit
result = model.fit()

# тестовый прогнозный период начнется с конца обучающего периода
start = len(train)
 
# и закончится в конце тестового
end = len(train) + len(test) - 1
  
# применим метод predict
predictions = result.predict(start, end)

# прогнозный период с конца имеющихся данных
start = len(queue)
 
# и закончится 36 месяцев спустя
end = (len(queue) - 1) + 3 * 12
 
# теперь построим прогноз на три года вперед
forecast = result.predict(start, end)
 
# посмотрим на весь 1963 год
forecast[-12:]