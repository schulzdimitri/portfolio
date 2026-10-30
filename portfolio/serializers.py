from rest_framework import serializers
from portfolio.models import Project


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
