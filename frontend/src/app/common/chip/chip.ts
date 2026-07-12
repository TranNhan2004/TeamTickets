import { NgClass, NgStyle } from '@angular/common';
import { ChangeDetectionStrategy, Component, computed, input } from '@angular/core';
import { MatChipsModule } from '@angular/material/chips';
import { MatTooltipModule } from '@angular/material/tooltip';
import { getContrastColor } from '../../core/utils/color';

type ChipColor = 'primary' | 'success' | 'warning' | 'danger' | 'info' | 'neutral';
type ChipSize = 'sm' | 'md' | 'lg';
type ChipVariant = 'filled' | 'outlined' | 'tonal' | 'text';

const COLOR_MAP: Record<ChipColor, string> = {
  primary: '#2563eb',
  success: '#16a34a',
  warning: '#d97706',
  danger: '#dc2626',
  info: '#0891b2',
  neutral: '#475569',
};

@Component({
  selector: 'app-chip',
  imports: [MatChipsModule, MatTooltipModule, NgClass, NgStyle],
  templateUrl: './chip.html',
  styleUrl: './chip.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class Chip {
  color = input<ChipColor>('primary');
  size = input<ChipSize>('md');
  variant = input<ChipVariant>('filled');
  tooltipTitle = input<string | null>(null);

  protected readonly chipColor = computed(() => COLOR_MAP[this.color()]);
  protected readonly chipLabelColor = computed(() => getContrastColor(this.chipColor()));

  protected readonly chipStyle = computed(() => {
    const color = this.chipColor();
    const labelColor = this.chipLabelColor();

    return {
      '--tt-chip-color': color,
      '--tt-chip-label-color': labelColor,
      '--tt-chip-tonal-bg': `${color}14`,
    };
  });

  protected readonly chipClass = computed(() => ({
    'tt-chip--sm': this.size() === 'sm',
    'tt-chip--md': this.size() === 'md',
    'tt-chip--lg': this.size() === 'lg',
    'tt-chip--filled': this.variant() === 'filled',
    'tt-chip--outlined': this.variant() === 'outlined',
    'tt-chip--tonal': this.variant() === 'tonal',
    'tt-chip--text': this.variant() === 'text',
  }));
}
