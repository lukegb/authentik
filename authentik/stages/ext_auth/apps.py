"""authentik ext_auth stage app config"""
from django.apps import AppConfig


class AuthentikStageExtAuthConfig(AppConfig):
    """authentik ext_auth stage config"""

    name = "authentik.stages.ext_auth"
    label = "authentik_stages_ext_auth"
    verbose_name = "authentik Stages.ExtAuth"
