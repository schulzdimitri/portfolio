from rest_framework import serializers
from portfolio.models import Experience, Project


class ProjectSerializer(serializers.ModelSerializer):
    name = serializers.CharField(
        max_length=100, required=True, allow_blank=False
    )
    description = serializers.CharField(
        max_length=500, required=True, allow_blank=False
    )
    url = serializers.CharField(
        max_length=200, required=True, allow_blank=False
    )

    class Meta:
        model = Project
        fields = ["id", "name", "description", "url"]

    def validate_url(self, value: str) -> str:
        if not value.startswith(("http://", "https://")):
            raise serializers.ValidationError(
                "Project URL must be a valid URL"
            )
        return value


class ExperienceSerializer(serializers.ModelSerializer):
    role = serializers.CharField(
        max_length=150, required=True, allow_blank=False
    )
    company = serializers.CharField(
        max_length=150, required=True, allow_blank=False
    )
    period = serializers.CharField(
        max_length=150, required=True, allow_blank=False
    )
    description = serializers.CharField(
        max_length=255, required=True, allow_blank=False
    )
    technologies = serializers.ListField(
        child=serializers.CharField(), required=True
    )

    class Meta:
        model = Experience
        fields = [
            "id",
            "role",
            "company",
            "period",
            "description",
            "technologies",
        ]
