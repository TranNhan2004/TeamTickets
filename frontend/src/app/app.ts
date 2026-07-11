import { Component, signal } from '@angular/core';
import { Avatar } from './common/avatar/avatar';
import { BellNotification } from './common/bell-notification/bell-notification';
import { Button } from './common/button/button';
import { Card } from './common/card/card';
import { Chip } from './common/chip/chip';
import { Dialog } from './common/dialog/dialog';
import { SearchBar } from './common/search-bar/search-bar';
import { Skeleton } from './common/skeleton/skeleton';
import { HasAccess } from './core/directives/has-access';

@Component({
  selector: 'app-root',
  imports: [Button, Chip, SearchBar, Avatar, BellNotification, Card, Dialog, Skeleton, HasAccess],
  templateUrl: './app.html',
  styleUrl: './app.scss',
})
export class App {
  protected readonly title = signal('Team Tickets');
  protected readonly latestSearch = signal('');
  protected readonly hasAccess = signal(false);

  protected handleSearch(query: string): void {
    this.latestSearch.set(query);
  }
}
