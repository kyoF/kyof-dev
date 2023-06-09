from django.conf import settings
from django.db import models


class Snippet(models.Model):
    title = models.CharField('title', max_length=128)
    code = models.TextField('code', blank=True)
    description = models.TextField('description', blank=True)
    created_by = models.ForeignKey(settings.AUTH_USER_MODEL, verbose_name='create_user', on_delete=models.CASCADE)
    created_at = models.DateTimeField('create_datetime', auto_now_add=True)
    updated_at = models.DateTimeField('update_datetime', auto_now=True)

    def __str__(self):
        return self.title
