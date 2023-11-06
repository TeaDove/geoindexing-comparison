#[derive(Clone, Debug, PartialEq)]
pub struct Point{
    pub lat: f64,
    pub lng: f64
}

impl Point{
    pub fn new(lat: f64, lng: f64) -> Self{
        return Point{lat, lng};
    }
}

mod tests{
    #[allow(unused_imports)]
    use super::*;

    #[test]
    fn test_point_new_ok(){
        let point = Point::new(57.0, 38.0);

        assert_eq!(point, Point{lat:57.0, lng:38.0})
    }
}
