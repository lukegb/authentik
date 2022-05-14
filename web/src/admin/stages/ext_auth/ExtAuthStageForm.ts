import { DEFAULT_CONFIG } from "@goauthentik/common/api/config";
import { first } from "@goauthentik/common/utils";
import "@goauthentik/elements/forms/FormGroup";
import "@goauthentik/elements/forms/HorizontalFormElement";
import { ModelForm } from "@goauthentik/elements/forms/ModelForm";

import { t } from "@lingui/macro";

import { TemplateResult, html } from "lit";
import { customElement, property } from "lit/decorators.js";
import { ifDefined } from "lit/directives/if-defined.js";

import { ExtAuthStage, StagesApi } from "@goauthentik/api";

@customElement("ak-stage-ext-auth-form")
export class ExtAuthStageForm extends ModelForm<ExtAuthStage, string> {
    loadInstance(pk: string): Promise<ExtAuth> {
        return new StagesApi(DEFAULT_CONFIG).stagesExtAuthRetrieve({
            stageUuid: pk,
        });
    }

    @property({ type: Boolean })
    showConnectionSettings = false;

    getSuccessMessage(): string {
        if (this.instance) {
            return t`Successfully updated stage.`;
        } else {
            return t`Successfully created stage.`;
        }
    }

    send = (data: ExtAuth): Promise<ExtAuth> => {
        if (this.instance) {
            return new StagesApi(DEFAULT_CONFIG).stagesExtAuthPartialUpdate({
                stageUuid: this.instance.pk || "",
                patchedExtAuthRequest: data,
            });
        } else {
            return new StagesApi(DEFAULT_CONFIG).stagesExtAuthCreate({
                extAuthStageRequest: data,
            });
        }
    };

    renderForm(): TemplateResult {
        return html`<form class="pf-c-form pf-m-horizontal">
            <div class="form-help-text">
                ${t`Redirect the user to their selected external authentication provider.`}
            </div>
            <ak-form-element-horizontal label=${t`Name`} ?required=${true} name="name">
                <input
                    type="text"
                    value="${ifDefined(this.instance?.name || "")}"
                    class="pf-c-form-control"
                    required
                />
            </ak-form-element-horizontal>
            <ak-form-group .expanded=${true}>
                <span slot="header"> ${t`Stage-specific settings`} </span>
                <div slot="body" class="pf-c-form">
                    <ak-form-element-horizontal name="silent">
                        <div class="pf-c-check">
                            <input
                                type="checkbox"
                                class="pf-c-check__input"
                                ?checked=${first(this.instance?.silent, true)}
                            />
                            <label class="pf-c-check__label"> ${t`Redirect silently`} </label>
                        </div>
                        <p class="pf-c-form__helper-text">
                            ${t`Redirect the user silently without an interstitial page.`}
                        </p>
                    </ak-form-element-horizontal>
                </div>
            </ak-form-group>
        </form>`;
    }
}
