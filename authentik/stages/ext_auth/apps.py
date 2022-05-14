"""authentik ext_auth stage app config"""
from authentik.blueprints.apps import ManagedAppConfig


class AuthentikStageExtAuthConfig(ManagedAppConfig):
    """authentik ext_auth stage config"""

    name = "authentik.stages.ext_auth"
    label = "authentik_stages_ext_auth"
    verbose_name = "authentik Stages.ExtAuth"
    default = True
