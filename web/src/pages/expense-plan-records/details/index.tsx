import { useNavigate, useParams } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { Form, InputNumber, Button, Skeleton, DatePicker } from 'antd';
import { useEffect, useState } from 'react';
import { centsToDecimal, decimalToCents } from '../../../util/brlFormat';
import useNotification from 'antd/es/notification/useNotification';
import { GoBackButton } from '../../../components/go-back-button';
import { parseDateOrNull } from '../../../util/formatDate';
import { SelectExpensePlans } from '../../../components/select-expense-plans';
import { restCreateExpensePlanRecord, restGetExpensePlan, restGetExpensePlanRecord, restUpdateExpensePlanRecord } from '../../../util/api';
import { QueryRecurrencyType } from '../../../util/api/generated/generated.schemas';
import { ManipulateType } from 'dayjs';

export function ExpensePlanRecordDetails() {
   const { ID } = useParams();
   const id = parseInt(ID ?? '0', 10);
   const isNew = !id;
   const [loadingMutation, setLoadingMutation] = useState(false);
   const navigate = useNavigate();

   const { data, isLoading, refetch } = useQuery({
      queryKey: ['expense-plan-details', id],
      queryFn: async () => {
         if (isNew) return null;
         return restGetExpensePlanRecord({ expense_plan_id: id });
      },
   });


   const [noti, notiCtx] = useNotification();

   const [form] = Form.useForm();
   const expensePlanId = Form.useWatch('expense_plan_id', form);

   const { data: dataExpensePlan } = useQuery({
      queryKey: ['expense-plan-record-get', expensePlanId],
      refetchOnMount: true,
      queryFn: async () => {
         if (!expensePlanId) return null;
         return restGetExpensePlan({
            expense_plan_id: expensePlanId,
         });
      },
   });

   useEffect(() => {
      const rTypeToDjsUnit: { [key in QueryRecurrencyType]: ManipulateType } = {
         MONTHLY: 'months',
         YEARLY: 'years',
      };
      const djsUnit = rTypeToDjsUnit[dataExpensePlan?.recurrency_type || 'MONTHLY'];
      form.setFieldsValue({
         expense_plan_sequence: (dataExpensePlan?.recurrency_interval || 0) + 1,
         amount_paid: 0,
         payment_date: parseDateOrNull(dataExpensePlan?.first_expense_plan_record?.payment_date)?.add(1, djsUnit),
         paid_date: parseDateOrNull(dataExpensePlan?.first_expense_plan_record?.payment_date)?.add(1, djsUnit),
      });
   }, [dataExpensePlan, form]);

   async function onFinish(values: typeof data) {
      console.log('onFinish', values);
      if (!values) return;

      try {
         setLoadingMutation(true);
         if (isNew) {
            const res = await restCreateExpensePlanRecord({
               expense_plan_id: values.expense_plan_id,
               amount_paid: decimalToCents(values.amount_paid),
               payment_date: values.payment_date,
               paid_date: values.paid_date,
            });
            noti.success({ message: 'Despesa registrada com sucesso. Redirecionando...' });
            navigate(`/expense-plan-records/${res.expense_plan_record_id}`);

            return;
         }

         await restUpdateExpensePlanRecord({
            expense_plan_record_id: id,
            expense_plan_id: values.expense_plan_id,
            amount_paid: decimalToCents(values.amount_paid),
            payment_date: values.payment_date,
            paid_date: values.paid_date,
         });
         refetch();
         noti.success({ message: 'Despesa atualizada com sucesso.' });

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
               {isNew ? 'Nova Despesa' : `Despesa #${id}`}
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
                     expense_plan_id: data?.expense_plan_id,
                     amount_paid: centsToDecimal(data?.amount_paid),
                     payment_date: parseDateOrNull(data?.payment_date),
                     paid_date: parseDateOrNull(data?.paid_date),
                     expense_plan_sequence: data?.expense_plan_sequence || 1,
                  }}
                  className="grid grid-cols-2 gap-4"
               >
                  <Form.Item name="expense_plan_id" label="Planejamento Despesa" rules={[{ required: true, message: 'Campo obrigatório' }]}>
                     <SelectExpensePlans />
                  </Form.Item>

                  <Form.Item name="amount_paid" label="Gasto" rules={[{ required: true }]}>
                     <InputNumber decimalSeparator=',' precision={2} className="w-full" />
                  </Form.Item>

                  <Form.Item name="payment_date" label="Data">
                     <DatePicker />
                  </Form.Item>

                  <Form.Item name="paid_date" label="Data Pagamento">
                     <DatePicker />
                  </Form.Item>

                  <Form.Item name="expense_plan_sequence" label="Parcela">
                     <InputNumber disabled className="w-full" />
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
