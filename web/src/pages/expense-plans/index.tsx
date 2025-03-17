import { useQuery } from '@tanstack/react-query';
import { restDeleteExpensePlan, restListExpensePlan } from '../../util/api';
import { Button, Table } from 'antd';
import { centsToCurrency } from '../../util/brlFormat';
import { formatDate } from '../../util/formatDate';
import { Link } from 'react-router-dom';
import { BiPencil, BiTrash } from 'react-icons/bi';
import useModal from 'antd/es/modal/useModal';
import useNotification from 'antd/es/notification/useNotification';
import { Config } from '../../config';

export function ExpensePlans() {
   const [modal, modalCtx] = useModal();
   const [noti, notiCtx] = useNotification();
   const { data, isLoading, refetch } = useQuery({
      queryKey: ['expense-plans-list'],
      refetchOnMount: true,
      refetchInterval: Config.DETAULT_REFETCH_INTERVAL,
      queryFn: async () => {
         return restListExpensePlan({});
      },
   });

   return (
      <div className='flex mt-10 flex-col gap-6 w-[98%]'>
         {modalCtx}
         {notiCtx}
         <div className='flex justify-between items-center'>
            <h1 className='font-semibold text-2xl'>Planejamento Despesas</h1>
            <Link to="/expense-plans/new" className="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600">
               Criar
            </Link>
         </div>
         <div>
            <Table
               dataSource={data?.items}
               loading={isLoading}
               rowKey={rec => rec.expense_plan_id!}
               columns={[
                  {
                     dataIndex: 'expense_plan_id',
                     title: '#',
                  },
                  {
                     dataIndex: 'category',
                     title: 'Categoria',
                  },
                  {
                     dataIndex: 'title',
                     title: 'Título',
                  },
                  {
                     dataIndex: 'amount_planned',
                     title: 'Gasto Planejado',
                     render(value) {
                        return centsToCurrency(value);
                     }
                  },
                  {
                     dataIndex: 'last_amount_spent',
                     title: 'Último Gasto',
                     render(value, rec) {
                        if (!value) return null;
                        const diff = (rec.amount_planned || 0) - value;
                        const last = centsToCurrency(value);
                        if (!diff) return last;
                        return (
                           <div className='flex flex-row gap-1'>
                              <span>{formatDate(rec.last_payment_date)}</span>
                              <span>{last}</span>
                              <span className={`${diff < 0 ? 'text-red-500': 'text-green-500'}`}>({centsToCurrency(diff)})</span>
                           </div>
                        );
                     }
                  },
                  {
                     dataIndex: 'recurrency_type',
                     title: 'Recorrência',
                     render(_, rec) {
                        switch (rec.recurrency_type) {
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
                     dataIndex: 'paid_count',
                     title: 'Parcela',
                     render(value, rec) {
                        if (!rec.recurrency_type) return null;
                        
                        let c = `${value || 0}`;
                        if (rec.recurrency_interval) c = `${c} de ${rec.recurrency_interval}`;
                        return c;
                     }
                  },
                  {
                     title: 'Ações',
                     dataIndex: 'actions',
                     render: (_, rec) => (
                        <div className='flex flex-col gap-2'>
                           <div className='flex flex-row'>
                              <Link to={`/expense-plans/${rec.expense_plan_id}`} className="text-blue-500 hover:underline">
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
                                             <p>Deseja remover o plano de despesas #{rec.expense_plan_id} <span className='font-bold'>{rec.title}</span>?</p>
                                             
                                             <div className='mt-4'>
                                                <p>Esta ação não poderá ser desfeita.</p>
                                                <p>O registro será removido permanentemente.</p>
                                             </div>
                                          </div>
                                       ),
                                       onOk: async () => {
                                          try {
                                             await restDeleteExpensePlan({ expense_plan_id: rec.expense_plan_id! });
                                             noti.success({
                                                message: 'Plano de despesas removido com sucesso',
                                             });
                                          } catch(e) {
                                             console.error('restDeleteExpensePlan', e);
                                             noti.error({
                                                message: 'Erro ao remover plano de despesas',
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