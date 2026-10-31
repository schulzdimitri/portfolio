import os
from django.conf import settings
from django.http import FileResponse, Http404
from django.views import View
from rest_framework import status
from rest_framework.request import Request
from rest_framework.response import Response
from rest_framework.views import APIView
from portfolio.models import About, Experience, Project
from portfolio.serializers import (
    AboutSerializer,
    ExperienceSerializer,
    ProjectSerializer,
)


class FrontendIndexView(View):
    def get(self, _request):
        index_path = os.path.join(settings.BASE_DIR, "templates", "index.html")
        if not os.path.isfile(index_path):
            raise Http404("Frontend index.html not found")
        return FileResponse(open(index_path, "rb"), content_type="text/html")


class HealthCheckView(APIView):
    def get(self, _request: Request) -> Response:
        return Response({"status": "ok"}, status=status.HTTP_200_OK)


class AboutListCreateView(APIView):
    def get(self, _request: Request) -> Response:
        try:
            queryset = About.objects.all()
            serializer = AboutSerializer(queryset, many=True)
            return Response(
                {
                    "type": "ABOUT",
                    "count": len(serializer.data),
                    "attributes": serializer.data,
                },
                status=status.HTTP_200_OK,
            )
        except Exception:
            return Response(
                {"detail": "Internal server error"},
                status=status.HTTP_500_INTERNAL_SERVER_ERROR,
            )

    def post(self, request: Request) -> Response:
        try:
            serializer = AboutSerializer(data=request.data)
            if not serializer.is_valid():
                return Response(
                    {"detail": serializer.errors},
                    status=status.HTTP_400_BAD_REQUEST,
                )
            serializer.save()
            return Response(
                {
                    "type": "ABOUT",
                    "count": 1,
                    "attributes": serializer.data,
                },
                status=status.HTTP_201_CREATED,
            )
        except Exception:
            return Response(
                {"detail": "Internal server error"},
                status=status.HTTP_500_INTERNAL_SERVER_ERROR,
            )


class ExperienceListCreateView(APIView):
    def get(self, _request: Request) -> Response:
        try:
            queryset = Experience.objects.all()
            serializer = ExperienceSerializer(queryset, many=True)
            return Response(
                {
                    "type": "EXPERIENCE",
                    "count": len(serializer.data),
                    "attributes": serializer.data,
                },
                status=status.HTTP_200_OK,
            )
        except Exception:
            return Response(
                {"detail": "Internal server error"},
                status=status.HTTP_500_INTERNAL_SERVER_ERROR,
            )

    def post(self, request: Request) -> Response:
        try:
            serializer = ExperienceSerializer(data=request.data)
            if not serializer.is_valid():
                return Response(
                    {"detail": serializer.errors},
                    status=status.HTTP_400_BAD_REQUEST,
                )
            serializer.save()
            return Response(
                {
                    "type": "EXPERIENCE",
                    "count": 1,
                    "attributes": serializer.data,
                },
                status=status.HTTP_201_CREATED,
            )
        except Exception:
            return Response(
                {"detail": "Internal server error"},
                status=status.HTTP_500_INTERNAL_SERVER_ERROR,
            )


class ProjectListCreateView(APIView):
    def get(self, _request: Request) -> Response:
        try:
            queryset = Project.objects.all()
            serializer = ProjectSerializer(queryset, many=True)
            return Response(
                {
                    "type": "PROJECT",
                    "count": len(serializer.data),
                    "attributes": serializer.data,
                },
                status=status.HTTP_200_OK,
            )
        except Exception:
            return Response(
                {"detail": "Internal server error"},
                status=status.HTTP_500_INTERNAL_SERVER_ERROR,
            )

    def post(self, request: Request) -> Response:
        try:
            serializer = ProjectSerializer(data=request.data)
            if not serializer.is_valid():
                return Response(
                    {"detail": serializer.errors},
                    status=status.HTTP_400_BAD_REQUEST,
                )
            serializer.save()
            return Response(
                {
                    "type": "PROJECT",
                    "count": 1,
                    "attributes": serializer.data,
                },
                status=status.HTTP_201_CREATED,
            )
        except Exception:
            return Response(
                {"detail": "Internal server error"},
                status=status.HTTP_500_INTERNAL_SERVER_ERROR,
            )


class ProjectDetailView(APIView):
    def get(self, _request: Request, project_name: str) -> Response:
        try:
            queryset = Project.objects.filter(name=project_name)
            serializer = ProjectSerializer(queryset, many=True)
            return Response(
                {
                    "type": "PROJECT",
                    "count": len(serializer.data),
                    "attributes": serializer.data,
                },
                status=status.HTTP_200_OK,
            )
        except Exception:
            return Response(
                {"detail": "Internal server error"},
                status=status.HTTP_500_INTERNAL_SERVER_ERROR,
            )
