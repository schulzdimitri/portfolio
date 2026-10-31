from django.db import models


class Project(models.Model):
    name = models.CharField(max_length=150)
    description = models.CharField(max_length=500)
    url = models.CharField(max_length=255)

    class Meta:
        db_table = "projects"
        ordering = ["id"]

    def __str__(self) -> str:
        return str(self.name)


class Experience(models.Model):
    role = models.CharField(max_length=150)
    company = models.CharField(max_length=150)
    period = models.CharField(max_length=150)
    description = models.CharField(max_length=255)
    technologies = models.JSONField(default=list)

    class Meta:
        db_table = "experiences"
        ordering = ["id"]

    def __str__(self) -> str:
        return f"{self.role} at {self.company}"
