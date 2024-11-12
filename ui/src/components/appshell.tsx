import { Burger, Button, Group, AppShell as MAppShell, Stack } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { PropsWithChildren } from 'react';
import { NavRow } from './nav/NavRow';
import { useAuth } from '@/lib/context/auth';
import { ProtectedLink } from './nav/ProtectedLink';
import { Link, NavLink } from 'react-router-dom';
import { navRoutes } from './nav/config';

export const AppShell = ({ children }: PropsWithChildren) => {
  const [opened, { toggle }] = useDisclosure();
  const { name, logout } = useAuth();

  return (
    <MAppShell
      header={{ height: 60 }}
      navbar={{ width: 300, breakpoint: 'sm', collapsed: { desktop: true, mobile: !opened } }}
      padding="md"
    >
      <MAppShell.Header>
        <Group justify='space-between' h="100%" px="md">
          <Group h="100%" px="md">
            <Burger opened={opened} onClick={toggle} hiddenFrom="sm" size="sm" />
            <Link to="/">
              <img src="/images/logo.png" className="object-contain h-8 pr-2" />
            </Link>
            <NavRow />
          </Group>
          <Button variant="subtle" onClick={logout}>
            {name}
          </Button>
        </Group>
      </MAppShell.Header>

      <MAppShell.Navbar p="md">
        <Stack
          align="stretch"
          justify="center"
        >
          {navRoutes.map(route => {
            return route.permission ? (
              <ProtectedLink key={route.href.toString()} permission={route.permission} to={route.href}>
                {({ isActive }) => (
                  <Button variant={isActive ? "light" : 'subtle'} className="text-gray-800 hover:text-red-500">{route.title}</Button>
                )}
              </ProtectedLink>
            ) : (
              <NavLink key={route.href.toString()} to={route.href}>
                {({ isActive }) => (
                  <Button variant={isActive ? "light" : 'subtle'} className="text-gray-800 hover:text-red-500">{route.title}</Button>
                )}
              </NavLink>
            );
          })}
        </Stack>
      </MAppShell.Navbar>

      <MAppShell.Main>
        {children}
      </MAppShell.Main>
    </MAppShell>
  )
}
