/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * My API
 * This is a sample API using Swagger.
 * OpenAPI spec version: 1.0
 */
export interface HandlerCreateExpensePlanRequest {
  amount_planned: number;
  category: QueryExpensePlanCategory;
  recurrency_type?: QueryRecurrencyType;
  title: string;
}

export interface HandlerCreateExpensePlanResponse {
  expense_plan_id?: number;
}

export interface HandlerDeleteExpensePlanRequest {
  expense_plan_id: number;
}

export interface HandlerDeleteExpensePlanResponse { [key: string]: unknown }

export interface HandlerGetExpensePlanRequest {
  expense_plan_id: number;
}

export interface HandlerGetExpensePlanResponse {
  amount_planned?: number;
  category?: QueryExpensePlanCategory;
  created_at?: string;
  expense_plan_id?: number;
  first_payment_date?: string;
  last_amount_spent?: number;
  last_paid_date?: string;
  last_payment_date?: string;
  paid_count?: number;
  recurrency_interval?: number;
  recurrency_type?: QueryRecurrencyType;
  title?: string;
  updated_at?: string;
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
  previous_payment_amount?: number;
  previous_payment_date?: string;
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
  first_payment_date?: string;
  last_amount_spent?: number;
  last_paid_date?: string;
  last_payment_date?: string;
  paid_count?: number;
  recurrency_interval?: number;
  recurrency_type?: QueryRecurrencyType;
  title?: string;
  updated_at?: string;
}

export interface HandlerUpdateExpensePlanRequest {
  amount_planned: number;
  category: QueryExpensePlanCategory;
  expense_plan_id: number;
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

