# Generated by Django 2.0 on 2018-07-26 01:06

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('users', '0002_auto_20180724_0400'),
    ]

    operations = [
        migrations.AddField(
            model_name='user',
            name='role',
            field=models.CharField(choices=[('farmer', '农场主'), ('consumer', '消费者')], default='consumer', max_length=100),
        ),
    ]