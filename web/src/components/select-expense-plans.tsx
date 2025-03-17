import { useQuery } from '@tanstack/react-query';
import { Select } from 'antd';
import { restListExpensePlan } from '../util/api';

type Props = Parameters<typeof Select>[0];

export function SelectExpensePlans(props: Props) {
   const { data, isLoading } = useQuery({
      queryKey: ['expense-plan-record-list'],
      refetchOnMount: true,
      queryFn: async () => {
         return restListExpensePlan({});
      },
   });
   return (
      <Select loading={isLoading} {...props}>
         {data?.items.map((item) => (
            <Select.Option key={item.expense_plan_id} value={item.expense_plan_id}>
               {item.title} {item.category}
            </Select.Option>
         ))}
      </Select>
   );
}