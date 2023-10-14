# эти библиотеки нам уже знакомы
import pandas as pd
import numpy as np
 
# модуль sparse библиотеки scipy понадобится 
# для работы с разреженными матрицами (об этом ниже)
from scipy.sparse import csr_matrix
 
# из sklearn мы импортируем алгоритм k-ближайших соседей
from sklearn.neighbors import NearestNeighbors

offices = pd.read_csv('services/predict/data/offices.csv')
offices_ratings = pd.read_csv('services/predict/data/offices_ratings.csv')

offices_ratings.drop(['datetime','id'], axis = 1, inplace = True)

# по горизонтали будут фильмы, по вертикали - пользователи, значения - оценки
user_item_matrix = offices_ratings.pivot_table(index = 'office_id', columns = 'user_id', values = 'rating', aggfunc='mean')

# параметр inplace = True опять же поможет сохранить результат
user_item_matrix.fillna(0, inplace = True)

# вначале сгруппируем (объединим) пользователей, возьмем только столбец rating 
# и посчитаем, сколько было оценок у каждого пользователя
users_votes = offices_ratings.groupby('user_id')['rating'].agg('count')
 
# сделаем то же самое, только для office
offices_votes = offices_ratings.groupby('office_id')['rating'].agg('count')
 
# теперь создадим фильтр (mask)
user_mask = users_votes[users_votes > 5].index
office_mask = offices_votes[offices_votes > 1].index
 
# применим фильтры и отберем фильмы с достаточным количеством оценок
user_item_matrix = user_item_matrix.loc[office_mask,:]
 
# а также активных пользователей
user_item_matrix = user_item_matrix.loc[:,user_mask]

# атрибут values передаст функции csr_matrix только значения датафрейма
csr_data = csr_matrix(user_item_matrix.values)

user_item_matrix = user_item_matrix.rename_axis(None, axis = 1).reset_index()

# создадим объект класса NearestNeighbors
knn = NearestNeighbors(metric = 'cosine', algorithm = 'brute', n_neighbors = 50, n_jobs = -1)
 
# обучим модель
knn.fit(csr_data)

recommendations = 250
search_word = 'ДО'

# для начала найдем офис в заголовках датафрейма office
office_search = offices[offices['salePointName'].str.contains(search_word)]

# вариантов может быть несколько, для простоты всегда будем брать первый вариант
# через iloc[0] мы берем первую строку столбца ['id']
office_id = office_search.iloc[0]['id']
 
# далее по индексу офиса в датасете offices найдем соответствующий индекс
# в матрице предпочтений
office_id = user_item_matrix[user_item_matrix['office_id'] == office_id].index[0]

distances, indices = knn.kneighbors(csr_data[office_id], n_neighbors = recommendations + 1)

# уберем лишние измерения через squeeze() и преобразуем массивы в списки с помощью tolist()
indices_list = indices.squeeze().tolist()
distances_list = distances.squeeze().tolist()
 
# далее с помощью функций zip и list преобразуем наши списки
indices_distances = list(zip(indices_list, distances_list))
 
# остается отсортировать список по расстояниям через key = lambda x: x[1] (то есть по второму элементу)
# в возрастающем порядке reverse = False
indices_distances_sorted = sorted(indices_distances, key = lambda x: x[1], reverse = False)
 
# и убрать первый элемент с индексом 901 (потому что это и есть "Матрица")
indices_distances_sorted = indices_distances_sorted[1:]

# создаем пустой список, в который будем помещать название офиса и расстояние до него
recom_list = []
 
# теперь в цикле будем поочередно проходить по кортежам
for ind_dist in indices_distances_sorted:
 
    # искать office_id в матрице предпочтений
    matrix_office_id = user_item_matrix.iloc[ind_dist[0]]['office_id']
 
    # выяснять индекс этого office_id в датафрейме office
    id = offices[offices['id'] == matrix_office_id].index
 
    # брать название офиса и расстояние до него
    #title = offices.iloc[id]['salePointName'].values[0]
    try:
      idd = offices.iloc[id]['id'].values[0]
    except IndexError:
      continue  
    dist = ind_dist[1]
 
    # помещать каждую пару в питоновский словарь
    # который, в свою очередь, станет элементом списка recom_list
    recom_list.append({'office_id' : idd, 'distance' : dist})

recom_df = pd.DataFrame(recom_list)

recom_df.to_json("services/predict/data/office_prediction.json", orient='index', indent=2 )