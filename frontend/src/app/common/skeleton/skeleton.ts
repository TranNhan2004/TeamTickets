import { NgClass, NgStyle } from '@angular/common';
import { ChangeDetectionStrategy, Component, computed, input } from '@angular/core';

type SkeletonVariant = 'text' | 'circle' | 'rectangle';

@Component({
  selector: 'app-skeleton',
  imports: [NgClass, NgStyle],
  templateUrl: './skeleton.html',
  styleUrl: './skeleton.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class Skeleton {
  variant = input<SkeletonVariant>('text');
  width = input<string | null>(null);
  height = input<string | null>(null);

  protected readonly skeletonClass = computed(() => ({
    [`tt-skeleton--${this.variant()}`]: true,
  }));

  protected readonly skeletonStyle = computed(() => {
    const variant = this.variant();
    const defaultSize = variant === 'circle' ? '40px' : '100%';

    return {
      width: this.width() ?? defaultSize,
      height: this.height() ?? (variant === 'text' ? '16px' : variant === 'circle' ? '40px' : '120px'),
    };
  });
}
