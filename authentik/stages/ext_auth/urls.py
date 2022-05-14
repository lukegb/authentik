"""API URLs"""
from authentik.stages.ext_auth.api import ExtAuthStageViewSet

api_urlpatterns = [("stages/ext_auth", ExtAuthStageViewSet)]
