Преза: https://docs.google.com/presentation/d/1kSzTfOWmmxW3PNvUgm5oYcwbPynePBmc/edit#slide=id.p2

## актуальность проблематики работы/исследования/проекта;
- KDtree - Apache Cassandra, Elasticsearch 
- GeoHash(S2) + btree -  mongo, dynamo
- R-Tree - Postgresql(Postgis), redis, oracle, tile38
- R*-Tree - sqlite
- H3 - clickhouse

## предметная область;

## содержательная постановка задачи;
Рассмотрим исследуемые операции:
- Insert(p origin) - вставка нового объекта p в индекс
- KNN(p origin, k int) - поиск k ближайших объектов от точки p
- RangeSearch(p origin, r float) - поиск всех объектов, которых входят в круг, определяющийся центром p и радиусом r.
- RectangleSearch(p origin, r float) - поиск всех объектов, которых входят в прямоугольник с нижним левым углом в plat - r, plng - r и верхним правым углом в  plat + r; plng + r

## математическая постановка задачи (или задач, в случае если их более одной);

## использование методов и средств ИКТ;
- 

## сведения о полученных результатах;
### Графики
Анализ графиков:
- Очень часто "грубая сила" (пояснить что это) работает лучше всего
- KDTree на малых значениях работает хорошо 
## дополнительный иллюстративный материал (при необходимости);
## вывод
- R(*)Tree - почти всегда
- KD(B)Tree - на KNN
- Geohash/H3+Btree - 