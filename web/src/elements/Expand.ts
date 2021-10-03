import { t } from "@lingui/macro";

import { CSSResult, html, LitElement, TemplateResult } from "lit";
import { customElement, property } from "lit/decorators";

import PFExpandableSection from "../../node_modules/@patternfly/patternfly/components/ExpandableSection/expandable-section.css";

@customElement("ak-expand")
export class Expand extends LitElement {
    @property({ type: Boolean })
    expanded = false;

    @property()
    textOpen = t`Show less`;

    @property()
    textClosed = t`Show more`;

    static get styles(): CSSResult[] {
        return [PFExpandableSection];
    }

    render(): TemplateResult {
        return html`<div class="pf-c-expandable-section ${this.expanded ? "pf-m-expanded" : ""}">
            <button
                type="button"
                class="pf-c-expandable-section__toggle"
                aria-expanded="${this.expanded}"
                @click=${() => {
                    this.expanded = !this.expanded;
                }}
            >
                <span class="pf-c-expandable-section__toggle-icon">
                    <i class="fas fa-angle-right" aria-hidden="true"></i>
                </span>
                <span class="pf-c-expandable-section__toggle-text"
                    >${this.expanded ? t`${this.textOpen}` : t`${this.textClosed}`}</span
                >
            </button>
            <slot ?hidden=${!this.expanded} class="pf-c-expandable-section__content"></slot>
        </div>`;
    }
}
