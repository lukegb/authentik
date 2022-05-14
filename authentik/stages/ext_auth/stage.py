"""ExtAuth stage logic"""
from django.contrib import messages
from django.http import HttpResponse
from django.utils.translation import gettext_lazy as _
from structlog.stdlib import get_logger

from authentik.core.models import Source, User
from authentik.flows.challenge import Challenge, ChallengeResponse, ChallengeTypes
from authentik.flows.planner import PLAN_CONTEXT_PENDING_USER
from authentik.flows.stage import ChallengeStageView

LOGGER = get_logger()
USER_ATTRIBUTE_SOURCE_IDP = "lukegb.com/source-idp"


class ExtAuthStageView(ChallengeStageView):
    """Redirects the user to their external authentication provider."""

    def error(self, message) -> Challenge:
        """Logs an error and displays it to the user."""
        messages.error(self.request, message)
        LOGGER.debug(message)

    def get_challenge(self, *args, **kwargs) -> Challenge:
        """Redirects the user to their external authentication provider."""
        if PLAN_CONTEXT_PENDING_USER not in self.executor.plan.context:
            self.error(_("No pending user."))
            return self.executor.stage_invalid()
        user: User = self.executor.plan.context[PLAN_CONTEXT_PENDING_USER]
        if USER_ATTRIBUTE_SOURCE_IDP not in user.attributes:
            # Skip this stage.
            return self.executor.stage_ok()
        try:
            source: Source = Source.objects.select_subclasses().get(
                slug=user.attributes[USER_ATTRIBUTE_SOURCE_IDP]
            )
        except Source.DoesNotExist:
            self.error(_("User's IdP does not exist."))
            return self.executor.stage_ok()

        ui_login_button = source.ui_login_button(self.request)
        challenge = ui_login_button.challenge
        challenge = type(challenge)(data=challenge.instance)

        return challenge

    def challenge_valid(self, response: ChallengeResponse) -> HttpResponse:
        return self.executor.stage_invalid()
