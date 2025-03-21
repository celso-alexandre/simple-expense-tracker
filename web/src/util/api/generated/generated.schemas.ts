/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * My API
 * This is a sample API using Swagger.
 * OpenAPI spec version: 1.0
 */
export interface HandlerCreateExpensePlanRecordRequest {
  amount_paid?: number;
  expense_plan_id?: number;
  paid_date?: string;
  payment_date?: string;
}

export interface HandlerCreateExpensePlanRecordResponse {
  expense_plan_record_id?: number;
}

export interface HandlerCreateExpensePlanRequest {
  amount_planned: number;
  category: QueryExpensePlanCategory;
  recurrency_interval?: number;
  recurrency_type?: QueryRecurrencyType;
  title: string;
}

export interface HandlerCreateExpensePlanResponse {
  expense_plan_id?: number;
}

export interface HandlerDeleteExpensePlanRecordsRequest {
  expense_plan_record_id: number;
}

export interface HandlerDeleteExpensePlanRecordsResponse { [key: string]: unknown }

export interface HandlerDeleteExpensePlanRequest {
  expense_plan_id: number;
}

export interface HandlerDeleteExpensePlanResponse { [key: string]: unknown }

export interface HandlerGetExpensePlanRecordRequest {
  expense_plan_id: number;
}

export interface HandlerGetExpensePlanRecordResponse {
  amount_paid?: number;
  created_at?: string;
  expense_plan?: HandlerGetExpensePlanRecordResponseExpensePlan;
  expense_plan_id?: number;
  expense_plan_record_id?: number;
  expense_plan_sequence?: number;
  paid_date?: string;
  payment_date?: string;
  updated_at?: string;
}

export interface HandlerGetExpensePlanRecordResponseExpensePlan {
  amount_planned?: number;
  category?: string;
  expense_plan_id?: number;
  recurrency_interval?: number;
  recurrency_type?: QueryRecurrencyType;
  title?: string;
}

export interface HandlerGetExpensePlanRequest {
  expense_plan_id: number;
}

export interface HandlerGetExpensePlanResponse {
  amount_planned?: number;
  category?: QueryExpensePlanCategory;
  created_at?: string;
  expense_plan_id?: number;
  first_expense_plan_record?: HandlerGetExpensePlanResponseExpensePlanRecord;
  first_expense_plan_record_id?: number;
  last_expense_plan_record?: HandlerGetExpensePlanResponseExpensePlanRecord;
  last_expense_plan_record_id?: number;
  recurrency_interval?: number;
  recurrency_type?: QueryRecurrencyType;
  title?: string;
  updated_at?: string;
}

export interface HandlerGetExpensePlanResponseExpensePlanRecord {
  amount_paid?: number;
  expense_plan_record_id?: number;
  expense_plan_sequence?: number;
  paid_date?: string;
  payment_date?: string;
}

export interface HandlerListExpensePlanRecordRequest { [key: string]: unknown }

export interface HandlerListExpensePlanRecordResponse {
  items?: HandlerListExpensePlanRecordResponseListExpensePlanRecord[];
}

export interface HandlerListExpensePlanRecordResponseListExpensePlan {
  amount_planned?: number;
  category?: string;
  expense_plan_id?: number;
  recurrency_interval?: number;
  recurrency_type?: QueryRecurrencyType;
  title?: string;
}

export interface HandlerListExpensePlanRecordResponseListExpensePlanRecord {
  amount_paid?: number;
  created_at?: string;
  expense_plan?: HandlerListExpensePlanRecordResponseListExpensePlan;
  expense_plan_id?: number;
  expense_plan_record_id?: number;
  expense_plan_sequence?: number;
  paid_date?: string;
  payment_date?: string;
  updated_at?: string;
}

export interface HandlerListExpensePlanRequest { [key: string]: unknown }

export interface HandlerListExpensePlanResponse {
  items?: HandlerListExpensePlanResponseListExpensePlan[];
}

export interface HandlerListExpensePlanResponseListExpensePlan {
  amount_planned?: number;
  category?: QueryExpensePlanCategory;
  created_at?: string;
  expense_plan_id?: number;
  first_expense_plan_record?: HandlerGetExpensePlanResponseExpensePlanRecord;
  first_expense_plan_record_id?: number;
  last_expense_plan_record?: HandlerGetExpensePlanResponseExpensePlanRecord;
  last_expense_plan_record_id?: number;
  recurrency_interval?: number;
  recurrency_type?: QueryRecurrencyType;
  title?: string;
  updated_at?: string;
}

export interface HandlerUpdateExpensePlanRecordRequest {
  amount_paid?: number;
  expense_plan_id?: number;
  expense_plan_record_id?: number;
  paid_date?: string;
  payment_date?: string;
}

export interface HandlerUpdateExpensePlanRecordResponse {
  expense_plan_record_id?: number;
}

export interface HandlerUpdateExpensePlanRequest {
  amount_planned: number;
  category: QueryExpensePlanCategory;
  expense_plan_id: number;
  recurrency_interval?: number;
  recurrency_type?: QueryRecurrencyType;
  title: string;
}

export interface HandlerUpdateExpensePlanResponse {
  expense_plan_id?: number;
}

export type QueryExpensePlanCategory = typeof QueryExpensePlanCategory[keyof typeof QueryExpensePlanCategory];


 
export const QueryExpensePlanCategory = {
  FOOD: 'FOOD',
  TRANSPORT: 'TRANSPORT',
  PROPERTY: 'PROPERTY',
  TAX: 'TAX',
  ENTERTAINMENT: 'ENTERTAINMENT',
  OTHER: 'OTHER',
} as const;

export type QueryRecurrencyType = typeof QueryRecurrencyType[keyof typeof QueryRecurrencyType];


 
export const QueryRecurrencyType = {
  MONTHLY: 'MONTHLY',
  YEARLY: 'YEARLY',
} as const;

