from dataclasses import dataclass


@dataclass
class Location:
    lat: float
    lng: float

    def __post_init__(self):
        if self.lat < -90 or self.lat > 90:
            raise ValueError(f'invalid lat: {self.lat}')
        if self.lng < -100 or self.lng > 180:
            raise ValueError(f'invalid lng: {self.lng}')

    def move(self, lat, lng):
        self.lat = lat
        self.lng = lng


@dataclass
class Car(Location):
    id: str


def move_all(items, lat, lng):
    for item in items:
        item.move(lat, lng)


if __name__ == '__main__':
    location = Location(lat=32.5252837, lng=34.9427434)
    location.move(32.0641339, 34.8742343)
    print(location)

    car = Car(id='python', lat=32.5253837, lng=34.9427434)

    car.move(32.0641339, 34.8742343)

    print(car)

    items = [Location(32.0641339, 34.9160405),
             Car(id='python', lat=32.5253837, lng=34.9427434)]

    move_all(items, 32.0641339, 34.8742343)

    for item in items:
        print(item)
