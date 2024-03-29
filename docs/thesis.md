## Сравнительный анализ алгоритмов геопоиска и пространственных индексов в высоконагруженных системах

Ибрагимов П.И. (m1803003@edu.misis.ru)

Работа с геоданными в высоконагруженных системах вынуждает использовать пространственные индексы и алгоритмы геопоиска.
Хоть самые популярные методы индексации были разработаны более чем 40 лет назад, консенсуса в вопросе индексов нет как в научной среде,
так и среди разработчиков систем управления баз данных. Например, СУБД Redis и Postgresql используют дерево R-Tree, а
AWS DynamoDB и MongoDB - geohash в связке с B-tree или префиксным деревом. При этом каждый год выходят новые научные статьи с новыми подходами к
анализу геоданных. Под геоданными в данной работе подразумевается пара координат широта-долгота.
Целью данной работы ставится исследование применимости разных пространственных индексов для решения базовых задач работы с геоданными.
Для достижения поставленной цели были сформулированы следующие задачи:
- Проанализировать существующие реализации индексов и их применимость в разных системах
- Проанализировать узкие места указанных индексов
- Выбрать индексы для тестирования
- Разработать систему тестирования производительности индексов с учетом разных задач и разных входных данных
- Проанализировать результаты работы разработанной системы
В результате проведенного теоретического анализа были сделаны следующие выводы:
- Наиболее популярным решением является индекс R-Tree, он имеет большое количество как реализаций на разных языков, так и
  способен оптимально выполнять большое количество задач. Минусом данного индекса является то, что он не подходит для обработки больших
  объемов данных из-за того, что он не содержит в себе функциональности к кластеризации.
- H3 и Geohash в связке с B-Tree и другими скалярными индексами заменяют R-Tree на тех задачах, где требование
  к кластеризации индекса есть.
- Существует множество оптимизаций тех или иных индексов. Например, SD-RTree - улучшенная версия R-Tree, которая
  позволяет индексу кластеризироваться. Но такие оптимизации редко доходят до стадии интеграции в какие-либо
  высоконагруженные системы.
  Разработанная аналитическая система тестирования индексов может помочь как и исследователям, так и разработчикам в анализе
  применимости индексов в разных задачах.

Научный руководитель - cт. преп. Тагиев Э. Р