"""ext_auth stage models"""

from django.db import models
from django.utils.translation import gettext_lazy as _
from django.views import View
from rest_framework.serializers import BaseSerializer

from authentik.flows.models import Stage


class ExtAuthStage(Stage):
    """Redirects the user to their external authenticator."""

    silent = models.BooleanField(
        default=True,
        help_text=_(
            (
                "When enabled, the HTML flow executor will immediately initiate logging in with "
                "the user's external authentication service."
            )
        ),
    )

    @property
    def serializer(self) -> BaseSerializer:
        from authentik.stages.ext_auth.api import ExtAuthStageSerializer

        return ExtAuthStageSerializer

    @property
    def type(self) -> type[View]:
        from authentik.stages.ext_auth.stage import ExtAuthStageView

        return ExtAuthStageView

    @property
    def component(self) -> str:
        return "ak-stage-ext-auth-form"

    class Meta:

        verbose_name = _("ExtAuth Stage")
        verbose_name_plural = _("ExtAuth Stages")
