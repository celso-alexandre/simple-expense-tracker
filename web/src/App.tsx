import { Route, Routes } from 'react-router-dom';
import dayjs from 'dayjs';
import localeData from 'dayjs/plugin/localeData';
import localizedFormat from 'dayjs/plugin/localizedFormat';
import { QueryParamProvider } from 'use-query-params';
import { ReactRouter6Adapter } from 'use-query-params/adapters/react-router-6';
import { SideMenu } from './components/side-menu';
import { NoMatch } from './not-found';
import { ExpensePlans } from './pages/expense-plans';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ExpensePlanDetails } from './pages/expense-plans/details';
import { ExpensePlanRecords } from './pages/expense-plan-records';

import('dayjs/locale/pt-br');

dayjs.extend(localizedFormat);
dayjs.extend(localeData);
dayjs.locale('pt-br');

const queryClient = new QueryClient();

export function App() {
  return (
    <div style={{ display: 'flex' }}>
      <SideMenu />

      <div
        style={{
          justifyContent: 'flex-end',
          marginLeft: '30px',
          width: '100%',
        }}>
          <QueryClientProvider client={queryClient}>
            <QueryParamProvider adapter={ReactRouter6Adapter}>
              <Routes>
                <Route path="/" element={<ExpensePlans />} />
                <Route path="/expense-plans" element={<ExpensePlans />} />

                <Route path="/expense-plans/new" element={<ExpensePlanDetails />} />
                <Route path="/expense-plans/:ID" element={<ExpensePlanDetails />} />

                <Route path="/expense-plan-records" element={<ExpensePlanRecords />} />

                <Route path="*" element={<NoMatch />} />
              </Routes>
            </QueryParamProvider>
        </QueryClientProvider>
      </div>
    </div>
  );
}
