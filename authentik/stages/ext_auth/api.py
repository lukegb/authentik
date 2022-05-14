"""ext_auth Stage API Views"""
from rest_framework.viewsets import ModelViewSet

from authentik.core.api.used_by import UsedByMixin
from authentik.flows.api.stages import StageSerializer
from authentik.stages.ext_auth.models import ExtAuthStage


class ExtAuthStageSerializer(StageSerializer):
    """ExtAuthStage Serializer"""

    class Meta:

        model = ExtAuthStage
        fields = StageSerializer.Meta.fields + ["silent"]


class ExtAuthStageViewSet(UsedByMixin, ModelViewSet):
    """ExtAuthStage Viewset"""

    queryset = ExtAuthStage.objects.all()
    serializer_class = ExtAuthStageSerializer
    filterset_fields = "__all__"
    ordering = ["name"]
