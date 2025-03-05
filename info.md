# Проект 'proglog'
#### Книги: 
1. [distributed-services-with-go](../books/distributed-services-with-go.pdf)


## Создание проекта:
```
mkdir proglog
cd proglog
go mod init github.com/KornilovLN/proglog
```

## Этапы отправки проекта на GitHub
### 1. Создание репозитория на GitHub
* Войдите в свой аккаунт на GitHub.
* Нажмите на "+" в правом верхнем углу и выберите "New repository".
* Введите имя репозитория: proglog
* Добавьте описание (опционально).
* Оставьте репозиторий публичным или сделайте приватным.
* Не инициализируйте репозиторий с README, .gitignore или лицензией (так как мы уже создали локальный проект).
* Нажмите "Create repository".

### 2. Инициализация Git в локальном проекте
```
cd proglog
git init
```

### 3. Добавление файлов в индекс Git
```
git add .
```

### 4. Создание первого коммита
```
git commit -m "Инициализация проекта proglog"
```

### 5. Подключение к удаленному репозиторию
```
git remote add origin https://github.com/KornilovLN/proglog.git
```

### 6. Отправка кода на GitHub
```
git push -u origin main
```

### 7. Если ваша основная ветка называется "master", а не "main", используйте:
```
git push -u origin master
```

### 8. Дополнительные шаги (если необходимо)
* Создание README.md (если его еще нет)
```
echo "# Проект proglog" > README.md
git add README.md
git commit -m "Добавлен README.md"
git push
```

### 9. Настройка .gitignore для Go-проекта
```
curl -o .gitignore https://raw.githubusercontent.com/github/gitignore/master/Go.gitignore
```
```
git add .gitignore
git commit -m "Добавлен .gitignore для Go"
git push
```

### 10. Если GitHub просит аутентификацию
```
Начиная с августа 2021 года, GitHub не принимает обычные пароли при push через HTTPS. 
Вместо этого используйте токен доступа (Personal Access Token)
или настройте SSH-ключи для аутентификации.
```