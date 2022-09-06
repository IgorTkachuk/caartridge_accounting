***[I used golang-migrate](https://github.com/golang-migrate/migrate)***
for best manipulate changes in database schema.

You can install cli in your Windows system by [scoop](https://scoop.sh/). After install you need to restart your IDE for apply paths.

[There](https://github.com/golang-migrate/migrate/issues/282#issuecomment-530743258) are a scenario when you have trouble with your _dirty_  mirgation and how you can fix it.
Mirgations placed in migration directory.

Makefile consist several target about that. Firsts - "create_mirgation" (prepare new migration file), second-"migrate" (apply mirgations on you database.You need to change DSN for correct usage)