# Eat The Rich. Правила

## Роли
### **Богач**
Особенности роли:

    - Пассивный доход: 2 зм/раунд
    - Скорость: 2 клетки
    - Ёмкость инвентаря: 9 ячеек
    - Покупок в раунд: 3
    - Стартовый капитал: 30 зм  
    - Стартовый предмет: "Поднятие цены" 

Возможности:

    - Использовать предмет (Использовать предмет из инвентаря)
    - Стать бедняком (Если любой игрок стал богаче вас - вы становитесь бедняком)

### **Бедняк**
Особенности роли:

    - Пассивный доход: 0 зм/раунд
    - Скорость: 1 клетка
    - Ёмкость инвентаря: 2 ячейки
    - Покупок в раунд: 1
    - Стартовый капитал: 12 зм
    - Стартовый предмет: "Нож"

Возможности:

    - Использовать предмет (Использовать предмет из инвентаря)
    - Смерть (Если баланс опускается до нуля - вы умираете)
    - Стать богачом (Если ваш баланс больше чем у богача, вы становитесь богачом)

## Поле
Всего на поле 26 клеток. Из них 14 Пустых клеток, 3 Хороших клетки, 3 Плохих клетки, 4 Магазина и 2 Казино

### **Пустая клетка**
При попадании на пустую клетку, ничего не происходит
### **Хорошая клетка**
При попадании на хорошую клетку, с равным шансом происходит одно случайное событие из этого списка:

    - Игрок получает бесплатный предмет 
      (Он может выбрать любой предмет из магазина, кроме тех что недоступны его роли)
    - Игрок получает 1 золотых монеты
    - Игрок получает 2 золотых монеты
    - Игрок получает 3 золотых монеты
    - Игрок получает 4 золотых монеты
    - Игрок получает 5 золотых монеты
    - Игрок получает "Нож"
    - Игрок может украсть 2 золота у другого игрока
    - Происходит событие из списка событий Плохой клетки
    - Игрок может выбрать другого игрока, для которого произойдёт событие из списка событий Плохой клетки
    - Игрок может переместиться на любую клетку поля
      (Активность клетки, на которую попал игрок, не может быть применена в этом раунде)


### **Плохая клетка**
При попадании на плохую клетку, с равным шансом происходит одно случайное событие из этого списка:

    - Игрок теряет половину своего золота
    - Богач получает 2 золотых монеты
    - Игрок теряет 3 золота
    - Игрок отдаёт другому игрок 2 золота
    - Игрок отдаёт другому игрок 3 золота
    - Игрок отдаёт другому игрок 5 золота
    - Игрок теряет 1 предмет на свой выбор
    - Богач контроллирует твой следующий ход ?
    - Игрок меняется местами со случайным игроком
    
### **Казино**
При попадании на клетку казино, с указанным шансом происходит одно случайное событие из этого списка:

    - Ничего (Шанс 50%)
    - Золото игрока удваивается (Шанс 10%)
    - Золото игрока уменьшается вдвое (Шанс 10%)
    - Игрок телепортируется на случайную клетку (Шанс 10%)
      (Активность клетки, на которую попал игрок, не может быть применена в этом раунде)
    - Происходит событие из списка событий Хорошей клетки (Шанс 10%)
    - Происходит событие из списка событий Плохой клетки (Шанс 10%) 
    
### **Магазин**
При попадании на клетку магазина, игрок попадает в магазин где может потратить свой баланс на покупку предметов

## Предметы
### **Только для богача**
#### "Поднятие цены"
Позволяет поднять цену любого предмета в магазине на 1.
#### "Бродвей"
Устанавливается на пустую клетку. Игрок попавший на Бродвей, платит богачу 2 монеты

### **Общие**
#### "Пушка"
Позволяет уничтожить стену.
#### "Безголовая лошадь"
+1 к скорости.
#### "Ловушка"
Поставить на клетку. Игрок попавший подвергается эффекту Плохой клетки
#### "Амулет восставшего"
Через раунд после смерти, восстанавливает игроку 5 монет
#### "Стена"
Поставить стену на путь. Больше им нельзя воспользоваться
#### "Шляпа телепортации"
Телепортирует на случайную клетку
#### "Шляпа волшебника"
Меняет тебя местами со случайным игроком
#### "Фальшивые монеты"
Временно увеличивает баланс игрока на 5 монет до следующего хода. В следующем ходу баланс уменьшается на 10 монет.

### **Только для бедняков**
#### "Нож"
Украсть 5 монет у игрока на той же клетке что и вы
#### "Гоблин"
Помещается на ту же клетку где стоит игрок. 
Перемещается случайно. При попадании на клетку с богачом, крадёт одну монету в пользу владельца