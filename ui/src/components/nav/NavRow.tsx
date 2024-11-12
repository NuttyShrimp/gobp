import { Button, Center, Group, Menu } from "@mantine/core"
import { navRoutes } from "./config";
import { HiChevronDown } from 'react-icons/hi';
import { Link, NavLink } from "react-router-dom";
import { ProtectedLink } from "./ProtectedLink";

export const NavRow = () => {
  const items = navRoutes.map((route) => {
    const menuItems = route.children?.map((item) => (
      item.permission ? (
        <Menu.Item component={ProtectedLink} permission={item.permission} to={item.href} key={item.href.toString()}>{item.title}</Menu.Item>
      ) : (
        <Menu.Item component={Link} to={item.href} key={item.href.toString()}>{item.title}</Menu.Item>
      )
    ));

    if (menuItems) {
      return (
        <Menu key={route.title} trigger="hover" transitionProps={{ exitDuration: 0 }} withinPortal>
          <Menu.Target>
            {route.permission ? (
              <ProtectedLink permission={route.permission} to={route.href}>
                {({ isActive }) => (
                  <Button
                    onClick={(event) => event.preventDefault()}
                    variant={isActive ? 'light' : "subtle"}
                    color="red"
                    className="text-gray-800 hover:text-red-500"
                  >
                    <Center>
                      <span className={"mr-1"}>{route.title}</span>
                      <HiChevronDown size="0.9rem" />
                    </Center>
                  </Button>
                )}
              </ProtectedLink>
            )
              : (
                <NavLink to={route.href}>
                  {({ isActive }) => (
                    <Button
                      onClick={(event) => event.preventDefault()}
                      variant={isActive ? 'light' : "subtle"}
                      color="red"
                      className="text-gray-800 hover:text-red-500"
                    >
                      <Center>
                        <span className={"mr-1"}>{route.title}</span>
                        <HiChevronDown size="0.9rem" />
                      </Center>
                    </Button>
                  )}
                </NavLink>
              )}
          </Menu.Target>
          <Menu.Dropdown>{menuItems}</Menu.Dropdown>
        </Menu>
      );
    }

    return route.permission ? (
      <ProtectedLink permission={route.permission} to={route.href}>
        {({ isActive }) => (
          <Button variant={isActive ? "light" : 'subtle'} className="text-gray-800 hover:text-red-500">{route.title}</Button>
        )}
      </ProtectedLink>
    ) : (
      <NavLink to={route.href}>
        {({ isActive }) => (
          <Button variant={isActive ? "light" : 'subtle'} className="text-gray-800 hover:text-red-500">{route.title}</Button>
        )}
      </NavLink>
    );
  });
  return (
    <Group gap={"md"} visibleFrom="sm">
      {items}
    </Group>
  )
}
