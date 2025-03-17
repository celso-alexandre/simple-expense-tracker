import { useNavigate, useParams } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { Form, Input, InputNumber, Select, Button, Skeleton } from 'antd';
import { restCreateExpensePlan, restGetExpensePlan, restUpdateExpensePlan } from '../../../util/api';
import { useState } from 'react';
import { centsToDecimal, decimalToCents } from '../../../util/brlFormat';
import { QueryExpensePlanCategory } from '../../../util/api/generated/generated.schemas';
import useNotification from 'antd/es/notification/useNotification';
import { GoBackButton } from '../../../components/go-back-button';

export function ExpensePlanDetails() {
   const { ID } = useParams();
   const id = parseInt(ID ?? '0', 10);
   const isNew = !id;
   const [loadingMutation, setLoadingMutation] = useState(false);
   const navigate = useNavigate();

   const { data, isLoading, refetch } = useQuery({
      queryKey: ['expense-plan-details', id],
      queryFn: async () => {
         if (isNew) return null;
         return restGetExpensePlan({ expense_plan_id: id });
      },
   });

   const [noti, notiCtx] = useNotification();

   const [form] = Form.useForm<typeof data>();
   const recurrencyType = Form.useWatch('recurrency_type', form);

   async function onFinish(values: typeof data) {
      console.log('onFinish', values);
      if (!values) return;

      try {
         setLoadingMutation(true);
         if (isNew) {
            const res = await restCreateExpensePlan({
               title: values.title,
               category: values.category,
               amount_planned: decimalToCents(values.amount_planned),
               recurrency_type: values.recurrency_type,
               recurrency_interval: values.recurrency_interval,
            });
            noti.success({ message: 'Planejamento de despesas criado com sucesso. Redirecionando...' });
            navigate(`/expense-plans/${res.expense_plan_id}`);

            return;
         }

         await restUpdateExpensePlan({
            expense_plan_id: id,
            title: values.title,
            category: values.category,
            amount_planned: decimalToCents(values.amount_planned),
            recurrency_type: values.recurrency_type,
            recurrency_interval: values.recurrency_interval,
         });
         refetch();
         noti.success({ message: 'Planejamento de despesas atualizado com sucesso.' });

         return;
      } catch (error) {
         console.error(error);
      } finally {
         setLoadingMutation(false);
      }
   }

   return (
      <div className="flex mt-10 flex-col w-full gap-6">
         {notiCtx}
         <div className='flex flex-row gap-2 items-center'>
            <GoBackButton />
            <h1 className="font-semibold text-2xl">
               {isNew ? 'Novo Planejamento Despesas' : `Planejamento Despesas #${id}`}
            </h1>
         </div>

         <div className="w-[98%] bg-white p-6 rounded-lg shadow">
            {(isLoading || loadingMutation) && !isNew ? (
               <Skeleton active />
            ) : (
               <Form
                  onFinish={onFinish}
                  form={form}
                  layout="vertical"
                  initialValues={{
                     ...data,
                     category: data?.category || '',
                     recurrency_type: data?.recurrency_type || '',
                     amount_planned: centsToDecimal(data?.amount_planned),
                     last_amount_spent: centsToDecimal(data?.last_amount_spent),
                  }}
                  className="grid grid-cols-2 gap-4"
               >
                  <Form.Item name="title" label="Título" rules={[{ required: true, message: 'Campo obrigatório' }]}>
                     <Input />
                  </Form.Item>

                  <Form.Item name="category" label="Categoria" rules={[{ required: true, message: 'Campo obrigatório' }]}>
                     <Select value={data?.category}>
                        {Object.entries(QueryExpensePlanCategory).map(([, value]) => (
                           <Select.Option key={value} value={value}>
                              {value}
                           </Select.Option>
                        ))}
                     </Select>
                  </Form.Item>

                  <Form.Item name="amount_planned" label="Gasto Planejado" rules={[{ required: true }]}>
                     <InputNumber decimalSeparator=',' precision={2} className="w-full" />
                  </Form.Item>

                  <Form.Item name="last_amount_spent" label="Último Gasto">
                     <InputNumber decimalSeparator=',' precision={2} disabled className="w-full" />
                  </Form.Item>

                  <Form.Item name="recurrency_type" label="Recorrência">
                     <Select>
                        <Select.Option value="">Única</Select.Option>
                        <Select.Option value="MONTHLY">Mensal</Select.Option>
                        <Select.Option value="YEARLY">Anual</Select.Option>
                     </Select>
                  </Form.Item>

                  <Form.Item name="paid_count" label="Parcelas quitadas">
                     <InputNumber disabled className="w-full" />
                  </Form.Item>

                  <Form.Item name="recurrency_interval" label="Parcelas totais" rules={[{ required: !!recurrencyType, message: 'Campo obrigatório' }]}>
                     <InputNumber disabled={!recurrencyType} className="w-full" />
                  </Form.Item>

                  <div className="col-span-2 flex justify-end gap-4">
                     <Button type="default" href="/expense-plans">
                        Cancelar
                     </Button>
                     <Button type="primary" htmlType="submit">
                        Salvar
                     </Button>
                  </div>
               </Form>
            )}
         </div>
      </div>
   );
}
