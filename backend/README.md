Есть пользователь, у него есть name,password, id and time Created.
Так же у него на аккаунте есть задачи. Задачи разделены по приоритету.
1)Paramount - tasks с самым высоким приоритетом. Срочные и важные
2)Hidh - tasks wich hidh importance. Срочные и не важные
3)Middl - tasks wich middl importance. Не ссрочные и важные
4)Low - tasks wich low importance. Не срочные и не важные
Веб-сервис должен позволять пользователям создавать свои аккаунты и заполнять их задачами, пользователь может удалять и править только задачами своего аккунта. 
Прежде чем добавить или отредактировать задачу пользователь должен сказать какой у нее приоретет. 

next request in grok:

ты хочешь чтобы я использовал go get github.com/lib/pq этот Postgres, но разве этот самый популярный и востребованный? там же вроде другой, не?
Ты же помнишь, что я хочу сделать работу не через терминал, а через страничку.
version: '3.8'
services:
db:
image: postgres:16-alpine
restart: always
environment:
POSTGRES_USER: todo
POSTGRES_PASSWORD: secret123
POSTGRES_DB: todo_db
ports:
- "5432:5432"
volumes:
- pgdata:/var/lib/postgresql/data

volumes:
pgdata:   
предпоследнюю строчку (volumes: ) подчеркивает как ошибку.
