# Требования к коду
* Открывающаяся фигурная скобка на той же строке, кроме функции main
* Максимальная длина строки - 120 символов
* По смысловым фрагментам код должен быть разбит на модули
* Наличие поясняющих комментариев ТОЛЬКО там, где без них сложно понять код
* Комментарий выносится на новую строку и отделяется одной пустой строкой ПЕРЕД
* Наименования структур и переменных - существительные, методов и функций - глаголы
* Наименования должны быть понятны СРАЗУ, нести смысл и заменять собой комментарии
* Префиксы в наименоавниях использовать нельзя, нужно постфиксы
* Все методы должны по длине помещаться в экран ноубтука без прокрутки
* ОДИН МЕТОД - ОДНО ДЕЙСТВИЕ
* Публичные методы должны иметь комментарий-пояснение к аргументам, результату и побочным эффектам
* Приватные методы (вспомогательные функции) не должны
* В привытных методах не использовать bool флаги, лучше сделать 2^n функций или ParameterObject
* Константы не должны содержать no / not
* Неочевидные bool выражения выносятся на отдельную строку и называются именем, отражающим суть

# Требования к веткам и коммитам
* Наименование - фамилия_дата_фича
* Отдельные ветки для багфиксов и фичей
* Никаких мерджей в мастер бранч
* Коммантарии при пулл реквесте
* Коммиты ДОЛЖНЫ быть разбиты по смыслу
* Коммиты производятся после каждого смыслового изменения
