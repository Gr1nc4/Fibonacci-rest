# Fibonacci-rest
Rest fibonacci sequence application

Данное приложение позводяет получить числа последовательность фибоначии в диапозоне от X до Y.
Отслеживаются два два эндпоинта:
GET "/" - Отображается описание приложения
POST "fibo" - Отпраувляет диапазон заданный в виде:
{
		"x":"10"
		"y":"20"
	}
  и отдает числа принадлежащия ряду фибоначи в этом диапозоне.
  
  # Параметры
  По умолчанию приложение запускается на порту ":8080"
  На компьютере должен быть установлен go v1.16
  # Запуск
  Скопируйте этот репозитоорий с помощью git clone github.com/Gr1nc4/Fibonacci-rest .
  Находясь в дериктории с репозиторием выполните go build или go run cmd/main.go
  Разрешите приложению доступ к сети
  С помощью любого интсрумента для тестирования http api (например postman) протестируйте приложение.
  
