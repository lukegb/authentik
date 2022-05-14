import "#components/ak-secret-text-input";
import "#elements/forms/FormGroup";
import "#elements/forms/HorizontalFormElement";
import "#elements/utils/TimeDeltaHelp";

import { DEFAULT_CONFIG } from "#common/api/config";

import { BaseStageForm } from "#admin/stages/BaseStageForm";

import { ExtAuthStage, StagesApi } from "@goauthentik/api";

import { msg } from "@lit/localize";
import { html, TemplateResult } from "lit";
import { customElement } from "lit/decorators.js";
import { ifDefined } from "lit/directives/if-defined.js";

@customElement("ak-stage-ext-auth-form")
export class ExtAuthStageForm extends BaseStageForm<ExtAuthStage> {
    async loadInstance(pk: string): Promise<ExtAuthStage> {
        const stage = await new StagesApi(DEFAULT_CONFIG).stagesExtAuthRetrieve({
            stageUuid: pk,
        });
        return stage;
    }

    async send(data: ExtAuthStage): Promise<ExtAuthStage> {
        if (this.instance) {
            return new StagesApi(DEFAULT_CONFIG).stagesExtAuthPartialUpdate({
                stageUuid: this.instance.pk || "",
                patchedExtAuthStageRequest: data,
            });
        }
        return new StagesApi(DEFAULT_CONFIG).stagesExtAuthCreate({
            extAuthStageRequest: data,
        });
    }

    renderForm(): TemplateResult {
        return html`<form class="pf-c-form pf-m-horizontal">
            <span> ${msg("Redirect the user to their selected external authentication provider.")} </span>
            <ak-form-element-horizontal label=${msg("Name")} ?required=${true} name="name">
                <input
                    type="text"
                    value="${ifDefined(this.instance?.name || "")}"
                    class="pf-c-form-control"
                    required
                />
            </ak-form-element-horizontal>
            <ak-form-group open label="${msg("Stage-specific settings")}">
                <div class="pf-c-form">
                    <ak-form-element-horizontal name="silent">
                        <label class="pf-c-switch">
                            <input
                                class="pf-c-switch__input"
                                type="checkbox"
                                ?checked=${this.instance?.silent ?? true}
                            />
                            <span class="pf-c-switch__toggle">
                                <span class="pf-c-switch__toggle-icon">
                                    <i class="fas fa-check" aria-hidden="true"></i>
                                </span>
                            </span>
                            <span class="pf-c-switch__label"
                                >${msg("Redirect silently")}</span
                            >
                        </label>
                        <p class="pf-c-form__helper-text">
                            ${msg(
				"Redirect the user silently without an interstitial page.",
                            )}
                        </p>
                    </ak-form-element-horizontal>
		</div>
            </ak-form-group>
        </form>`;
    }
}

declare global {
    interface HTMLElementTagNameMap {
        "ak-stage-ext-auth-form": ExtAuthStageForm;
    }
}
