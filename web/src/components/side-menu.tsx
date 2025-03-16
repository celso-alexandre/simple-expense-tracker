import { useEffect, useMemo, useState } from 'react';
import type { MenuProps } from 'antd';
import { Button, Menu } from 'antd';
import { BiListUl } from 'react-icons/bi';
import {  MdOutlinePointOfSale } from 'react-icons/md';
import { AiOutlineMenuFold, AiOutlineMenuUnfold } from 'react-icons/ai';
import { useMediaQuery } from 'react-responsive';
import { matchRoutes, useLocation } from 'react-router-dom';

type MenuItem = Required<MenuProps>['items'][number];

export function SideMenu() {
  const isDesktopOrLaptop = useMediaQuery({
    query: '(min-width: 1360px)',
  });
  const [collapsed, setCollapsed] = useState(true);
  const location = useLocation();

  const toggleCollapsed = () => {
    setCollapsed((bol) => {
      return !bol;
    });
  };

  useEffect(() => {
    setCollapsed(!isDesktopOrLaptop);
  }, [isDesktopOrLaptop]);

  const items: (MenuItem & {
    children?: MenuItem[];
  })[] = useMemo(() => [
    {
      key: '1',
      label: 'Despesas',
      icon: <MdOutlinePointOfSale />,
      children: [
        {
          key: '/expense-plans',
          label: <a href="/expense-plans">Planejamento Despesas</a>,
          icon: <BiListUl />,
        },
      ],
    },
  ], []);

  const routes = useMemo(() => {
    return items.reduce(
      (prev, cur) => {
        if (!(cur.key as string)?.includes('/') && !cur.children?.length)
          return prev;
        return [
          ...prev,
          ...(!cur.children
            ? [{ path: cur.key as string }]
            : cur.children
                ?.filter((child) => {
                  return child?.key;
                })
                ?.flatMap((child) => {
                  return {
                    path: child?.key as string,
                    parent: cur.key as string,
                  };
                }) || []),
        ];
      },
      [] as { path: string; parent?: string }[]
    );
  }, [items]);

  const matches = matchRoutes(routes, location);
  const { route } = matches?.[0] || {};

  console.log({ matches, routes, location });

  return (
    <div style={{ width: !collapsed ? 256 : 80, justifyContent: 'flex-start' }}>
      <Menu
        defaultOpenKeys={[route?.parent ?? '1']}
        defaultSelectedKeys={[route?.path ?? '/sales']}
        mode="inline"
        theme="light"
        inlineCollapsed={collapsed}
        items={items}
        contentEditable={false}
      />
      <Button
        type="primary"
        onClick={toggleCollapsed}
        style={{ width: '100%' }}>
        {collapsed ? <AiOutlineMenuUnfold /> : <AiOutlineMenuFold />}
      </Button>
    </div>
  );
}
