import { useQuery } from '@tanstack/react-query';
import { Button, Table } from 'antd';
import { centsToCurrency } from '../../util/brlFormat';
import { formatDate } from '../../util/formatDate';
import { Link } from 'react-router-dom';
import { BiPencil, BiTrash } from 'react-icons/bi';
import { Config } from '../../config';
import { restDeleteExpensePlanRecord, restListExpensePlanRecord } from '../../util/api';
import useModal from 'antd/es/modal/useModal';
import useNotification from 'antd/es/notification/useNotification';

export function ExpensePlanRecords() {
   const [modal, modalCtx] = useModal();
   const [noti, notiCtx] = useNotification();
   const { data, isLoading, refetch } = useQuery({
      queryKey: ['expense-plan-record-list'],
      refetchOnMount: true,
      refetchInterval: Config.DETAULT_REFETCH_INTERVAL,
      queryFn: async () => {
         return restListExpensePlanRecord({});
      },
   });

   return (
      <div className='flex mt-10 flex-col gap-6 w-[98%]'>
         {modalCtx}
         {notiCtx}
         <div className='flex justify-between items-center'>
            <h1 className='font-semibold text-2xl'>Despesas</h1>
            <Link to="/expense-plan-records/new" className="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600">
               Criar
            </Link>
         </div>
         <div>
            <Table
               dataSource={data?.items}
               loading={isLoading}
               rowKey={rec => rec.expense_plan_record_id!}
               columns={[
                  {
                     dataIndex: 'expense_plan_record_id',
                     title: '#',
                  },
                  {
                     dataIndex: 'expense_plan_id',
                     title: 'Planejamento Despesa',
                     align: 'center',
                     render(value, rec) {
                        return <Link to={`/expense-plans/${value}`}>{value} {rec.expense_plan?.title}</Link>;
                     }
                  },
                  {
                     dataIndex: 'payment_date',
                     title: 'Data',
                     render(value) {
                        return formatDate(value);
                     }
                  },
                  {
                     dataIndex: 'paid_date',
                     title: 'Pago em',
                     render(value) {
                        return formatDate(value);
                     }
                  },
                  {
                     dataIndex: 'amount_paid',
                     title: 'Gasto',
                     render(value, rec) {
                        const v = centsToCurrency(value);
                        const p = rec.expense_plan?.amount_planned || 0;
                        if (p === value || !p) {
                           return v;
                        }
                        const diff = value - p;
                        return (
                           <div className='flex flex-row gap-2'>
                              <span>{v}</span>
                              <span 
                                 className={diff > 0 ? 'text-red-500' : 'text-green-500'}
                              >
                                 ({centsToCurrency(diff * -1)})
                              </span>
                           </div>
                        );
                     }
                  },
                  {
                     dataIndex: 'recurrency_type',
                     title: 'Recorrência',
                     render(_, rec) {
                        switch (rec.expense_plan?.recurrency_type) {
                           case 'MONTHLY':
                              return 'Mensal';
                           case 'YEARLY':
                              return 'Anual';
                           default:
                              return 'Única';
                        }
                     }
                  },
                  {
                     dataIndex: 'expense_plan_record_sequence',
                     title: 'Parcelas',
                     align: 'center',
                     render(value, rec) {
                        if (!rec.expense_plan?.recurrency_type) return null;
                        
                        let c = `${value || 0}`;
                        if (rec.expense_plan?.recurrency_interval) c = `${c} de ${rec.expense_plan?.recurrency_interval}`;
                        return c;
                     }
                  },
                  {
                     title: 'Ações',
                     dataIndex: 'actions',
                     render: (_, rec) => (
                        <div className='flex flex-col gap-2'>
                           <div className='flex flex-row'>
                              <Link to={`/expense-plan-records/${rec.expense_plan_record_id}`} className="text-blue-500 hover:underline">
                                 <Button icon={<BiPencil />} type="link">
                                    Editar
                                 </Button>
                              </Link>
                           </div>
                           <div>
                              <Button 
                                 icon={<BiTrash />} 
                                 type="link" 
                                 className="!text-red-500 hover:underline"
                                 onClick={() => {
                                    modal.confirm({
                                       title: 'Deseja continuar?',
                                       content: (
                                          <div>
                                             <p>Deseja remover a despesa #{rec.expense_plan_record_id} <span className='font-bold'>{rec.expense_plan?.title}</span>?</p>
                                             
                                             <div className='mt-4'>
                                                <p>Esta ação não poderá ser desfeita.</p>
                                                <p>O registro será removido permanentemente.</p>
                                             </div>
                                          </div>
                                       ),
                                       onOk: async () => {
                                          try {
                                             await restDeleteExpensePlanRecord({ expense_plan_record_id: rec.expense_plan_record_id! });
                                             noti.success({
                                                message: 'Pagamento removido com sucesso',
                                             });
                                          } catch(e) {
                                             console.error('restDeleteExpenseRecord', e);
                                             noti.error({
                                                message: 'Erro ao remover Recordo de despesas',
                                             });
                                          }
                                          refetch();                                     
                                       }
                                    })
                                 }}
                              >
                                 Remover
                              </Button>
                           </div>
                        </div>
                     ),
                  }, 
               ]}
            />
         </div>
      </div>
   );
}