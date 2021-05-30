import uuid

import datetime as dt
from django.db import models


class Room(models.Model):
    uuid = models.UUIDField(primary_key=True,
                            default=uuid.uuid4(),
                            db_column='uuid')
    name = models.CharField(max_length=100,
                            default='unnamed',
                            blank=None,
                            db_column='name')
    create_at = models.DateTimeField(default=dt.datetime.now(),
                                     db_column='create_at')

    def __str__(self):
        return self.name + ": " + str(self.uuid)

    class Meta:
        db_table = 'room'
