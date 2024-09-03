import { Burger, Group, AppShell as MAppShell } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { PropsWithChildren } from 'react';
import { NavRow } from './nav/NavRow';

export const AppShell = ({ children }: PropsWithChildren) => {
  const [opened, { toggle }] = useDisclosure();

  return (
    <MAppShell
      header={{ height: 60 }}
      navbar={{ width: 300, breakpoint: 'sm', collapsed: { desktop: true, mobile: !opened } }}
      padding="md"
    >
      <MAppShell.Header>
        <Group h="100%" px="md">
          <Burger opened={opened} onClick={toggle} hiddenFrom="sm" size="sm" />
          <img src="/images/logo.png" className="object-contain h-8 pr-2" />
          <NavRow />
        </Group>
      </MAppShell.Header>

      <MAppShell.Navbar py="md" px={4}>
      </MAppShell.Navbar>

      <MAppShell.Main>
        {children}
      </MAppShell.Main>
    </MAppShell>
  )
}
