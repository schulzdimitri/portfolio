import os
import sys
from pathlib import Path
import dj_database_url

if os.getenv("TESTING") != "1" and "pytest" not in sys.modules:
    try:
        from dotenv import load_dotenv

        load_dotenv()
    except ImportError:
        pass

BASE_DIR = Path(__file__).resolve().parent.parent

SECRET_KEY = os.getenv(
    "DJANGO_SECRET_KEY",
    "django-insecure-02imp%=0382)(zn16_062cnw*u@9skqa)6y-t7at44%7*-$i86",
)

DEBUG = os.getenv("DEBUG", "True").lower() in ("true", "1", "yes")

ALLOWED_HOSTS = ["*"]

INSTALLED_APPS = [
    "django.contrib.admin",
    "django.contrib.auth",
    "django.contrib.contenttypes",
    "django.contrib.sessions",
    "django.contrib.messages",
    "django.contrib.staticfiles",
    "corsheaders",
    "rest_framework",
    "portfolio",
]

MIDDLEWARE = [
    "corsheaders.middleware.CorsMiddleware",
    "django.middleware.security.SecurityMiddleware",
    "whitenoise.middleware.WhiteNoiseMiddleware",
    "django.contrib.sessions.middleware.SessionMiddleware",
    "django.middleware.common.CommonMiddleware",
    "django.middleware.csrf.CsrfViewMiddleware",
    "django.contrib.auth.middleware.AuthenticationMiddleware",
    "django.contrib.messages.middleware.MessageMiddleware",
    "django.middleware.clickjacking.XFrameOptionsMiddleware",
]

ROOT_URLCONF = "config.urls"

TEMPLATES = [
    {
        "BACKEND": "django.template.backends.django.DjangoTemplates",
        "DIRS": [BASE_DIR / "templates"],
        "APP_DIRS": True,
        "OPTIONS": {
            "context_processors": [
                "django.template.context_processors.request",
                "django.contrib.auth.context_processors.auth",
                "django.contrib.messages.context_processors.messages",
            ],
        },
    },
]

WSGI_APPLICATION = "config.wsgi.application"


def get_database_config() -> dict:
    database_url = os.getenv("DATABASE_URL")
    if database_url and os.getenv("TESTING") != "1":
        return dj_database_url.parse(
            database_url,
            conn_max_age=0,
            ssl_require=True,
        )

    is_vercel = os.getenv("VERCEL") == "1" or "VERCEL" in os.environ
    db_dir = "/tmp" if is_vercel else str(BASE_DIR)
    db_path = os.getenv("DB_PATH", os.path.join(db_dir, "schema.db"))
    return {
        "ENGINE": "django.db.backends.sqlite3",
        "NAME": db_path,
    }


DATABASES = {
    "default": get_database_config(),
}

AUTH_PASSWORD_VALIDATORS = [
    {
        "NAME": (
            "django.contrib.auth.password_validation."
            "UserAttributeSimilarityValidator"
        ),
    },
    {
        "NAME": (
            "django.contrib.auth.password_validation."
            "MinimumLengthValidator"
        ),
    },
    {
        "NAME": (
            "django.contrib.auth.password_validation."
            "CommonPasswordValidator"
        ),
    },
    {
        "NAME": (
            "django.contrib.auth.password_validation."
            "NumericPasswordValidator"
        ),
    },
]

allowed_origins_env = os.getenv("ALLOWED_ORIGINS")
if allowed_origins_env:
    CORS_ALLOWED_ORIGINS = [
        origin.strip()
        for origin in allowed_origins_env.split(",")
        if origin.strip()
    ]
    CORS_ALLOW_ALL_ORIGINS = False
    CORS_ALLOW_CREDENTIALS = True
else:
    CORS_ALLOW_ALL_ORIGINS = True
    CORS_ALLOW_CREDENTIALS = False

LANGUAGE_CODE = "en-us"

TIME_ZONE = "UTC"

USE_I18N = True

USE_TZ = True

STATIC_URL = "/static/"
STATICFILES_DIRS = [BASE_DIR / "static"]
STATIC_ROOT = BASE_DIR / "staticfiles"
WHITENOISE_ROOT = BASE_DIR / "static"
WHITENOISE_USE_FINDERS = True

DEFAULT_AUTO_FIELD = "django.db.models.BigAutoField"

REST_FRAMEWORK = {
    "UNAUTHENTICATED_USER": None,
}
