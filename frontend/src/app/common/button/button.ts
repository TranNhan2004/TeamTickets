import { NgClass, NgStyle } from '@angular/common';
import { ChangeDetectionStrategy, Component, computed, input } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { getContrastColor } from '../../core/utils/color';

type ButtonColor = 'primary' | 'success' | 'warning' | 'danger' | 'info' | 'neutral';
type ButtonSize = 'sm' | 'md' | 'lg';
type ButtonVariant = 'filled' | 'outlined' | 'tonal' | 'text';
type ButtonAfterPressEffect = 'underline' | 'none';

const COLOR_MAP: Record<ButtonColor, string> = {
  primary: '#2563eb',
  success: '#16a34a',
  warning: '#d97706',
  danger: '#dc2626',
  info: '#0891b2',
  neutral: '#475569',
};

@Component({
  selector: 'app-button',
  imports: [MatButtonModule, MatTooltipModule, NgClass, NgStyle],
  templateUrl: './button.html',
  styleUrl: './button.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class Button {
  color = input<ButtonColor>('primary');
  size = input<ButtonSize>('md');
  variant = input<ButtonVariant>('filled');
  afterPressEffect = input<ButtonAfterPressEffect>('none');
  disabled = input(false);
  fullWidth = input(false);
  tooltipTitle = input<string | null>(null);

  protected readonly buttonColor = computed(() => COLOR_MAP[this.color()]);
  protected readonly buttonLabelColor = computed(() => getContrastColor(this.buttonColor()));

  protected readonly buttonStyle = computed(() => {
    const color = this.buttonColor();
    const labelColor = this.buttonLabelColor();
    const textHoverColor = '#60a5fa';

    return {
      '--tt-button-color': color,
      '--tt-button-label-color': labelColor,
      '--tt-button-tonal-bg': `${color}14`,
      '--mat-button-filled-container-color': color,
      '--mat-button-filled-label-text-color': labelColor,
      '--mat-button-filled-state-layer-color': labelColor,
      '--mat-button-filled-ripple-color': `${labelColor}1f`,
      '--mat-button-outlined-label-text-color': color,
      '--mat-button-outlined-outline-color': color,
      '--mat-button-outlined-state-layer-color': color,
      '--mat-button-outlined-ripple-color': `${color}1f`,
      '--mat-button-text-label-text-color': color,
      '--mat-button-text-state-layer-color': textHoverColor,
      '--mat-button-text-ripple-color': `${textHoverColor}1f`,
    };
  });

  protected readonly buttonClass = computed(() => ({
    'tt-button--sm': this.size() === 'sm',
    'tt-button--md': this.size() === 'md',
    'tt-button--lg': this.size() === 'lg',
    'tt-button--full': this.fullWidth(),
    'tt-button--filled': this.variant() === 'filled',
    'tt-button--outlined': this.variant() === 'outlined',
    'tt-button--tonal': this.variant() === 'tonal',
    'tt-button--text': this.variant() === 'text',
    'tt-button--text-effect-underline':
      this.variant() === 'text' && this.afterPressEffect() === 'underline',
  }));
}
