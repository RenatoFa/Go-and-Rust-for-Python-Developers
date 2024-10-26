from dataclasses import dataclass
from datetime import datetime, timezone

PER_HOUR = 5


@dataclass
class VM:
    start_time: datetime
    end_time: datetime

    def cost(self):
        end_time = self.end_time
        if end_time is None:
            end_time = timezone.utc()

        duration = end_time - self.start_time
        hours = duration.total_seconds() / (60*60)

        return hours * PER_HOUR


class Spot(VM):
    def cost(self):
        p = super().cost()

        return p * .8


def total_cost(vms):
    return sum(vm.cost() for vm in vms)


if __name__ == '__main__':
    vms = [
        VM(start_time=datetime(2022, 4, 12, 17, 30, tzinfo=timezone.utc),
           end_time=datetime(2022, 4, 12, 19, 54, tzinfo=timezone.utc)),
        Spot(start_time=datetime(2022, 4, 13, 9, 46, tzinfo=timezone.utc),
             end_time=datetime(2022, 4, 15, 12, 7, tzinfo=timezone.utc))
    ]

    print(total_cost(vms))
