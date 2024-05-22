# skillfactory_task_30.8.1

# Запуск  

1. Сколнировать репозиторий:
```bash   
git clone https://github.com/MaksimovDenis/Skillfactory_task_30.8.1.git
```

2. Перейти в директорию проекта (если Вы не в ней):  
```bash    
cd skillfactory_task_30.8.1 
```

3. Поднимет базу данных:  
```bash      
make up 
```

4. Локальная установка утилиты миграции:  
```bash      
make install-deps 
```

5. Накатить миграции:  
```bash      
make down 
```

6. Запустить сервис:  
```bash      
make run 
```

# Примеры запросов  

 - Получать список всех задач  
 ![image](https://github.com/MaksimovDenis/vk_restAPI/assets/44647373/d2b88593-6dec-4fea-93f3-6ca89df895ad)  
 - Получать список задач по автору  
 ![image](https://github.com/MaksimovDenis/vk_restAPI/assets/44647373/27a14a5c-2867-4d38-a5d9-5fa1c87a80ed)  
 - Получать список задач по метке  
 ![image](https://github.com/MaksimovDenis/vk_restAPI/assets/44647373/c94d9d91-ed13-48a0-8fbc-6af39696491d)  
 - Обновлять задачу по id
 ![image](https://github.com/MaksimovDenis/vk_restAPI/assets/44647373/03d76568-81f9-4a8a-a7bf-a5d053f9c3d4)  
 - Удалять задачу по id
 ![image](https://github.com/MaksimovDenis/vk_restAPI/assets/44647373/423e53ca-1420-45cc-9be1-769fa5a20a60)  
 ![image](https://github.com/MaksimovDenis/vk_restAPI/assets/44647373/2448f42b-b1ac-44fc-b0f9-30a2cd2c6ce1)  
  - Создавать задачи
 ![image](https://github.com/MaksimovDenis/vk_restAPI/assets/44647373/012307fa-f6c7-4177-a0ad-223e420a5349)  
 ![image](https://github.com/MaksimovDenis/vk_restAPI/assets/44647373/a66ee3f6-6750-4c35-9d44-1b2e9122fd8a)  