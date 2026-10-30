from django.contrib import admin
from portfolio.models import About, Experience, Project


@admin.register(Project)
class ProjectAdmin(admin.ModelAdmin):
    list_display = ("id", "name", "url")
    search_fields = ("name", "description")


@admin.register(About)
class AboutAdmin(admin.ModelAdmin):
    list_display = ("id", "title")


@admin.register(Experience)
class ExperienceAdmin(admin.ModelAdmin):
    list_display = ("id", "role", "company", "period")
    search_fields = ("role", "company")
