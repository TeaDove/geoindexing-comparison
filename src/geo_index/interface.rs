trait GeoIndex{
    fn new(points: Points) -> Self;
    fn points(self) -> Points;
    fn knn(self, point: Point, k: i32) -> Points;
}
