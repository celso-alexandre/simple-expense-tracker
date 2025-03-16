/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * My API
 * This is a sample API using Swagger.
 * OpenAPI spec version: 1.0
 */
import type {
  HandlerCreateExpensePlanRequest,
  HandlerCreateExpensePlanResponse,
  HandlerGetExpensePlanRequest,
  HandlerGetExpensePlanResponse,
  HandlerListExpensePlanRequest,
  HandlerListExpensePlanResponse
} from './generated.schemas';


/**
 * Create expense-plan item
 * @summary Create expense-plan item
 */
export type postExpensePlanCreateResponse200 = {
  data: HandlerCreateExpensePlanResponse
  status: 200
}
    
export type postExpensePlanCreateResponseComposite = postExpensePlanCreateResponse200;
    
export type postExpensePlanCreateResponse = postExpensePlanCreateResponseComposite & {
  headers: Headers;
}

export const getPostExpensePlanCreateUrl = () => {


  

  return `/expense-plan/create`
}

export const postExpensePlanCreate = async (handlerCreateExpensePlanRequest: HandlerCreateExpensePlanRequest, options?: RequestInit): Promise<postExpensePlanCreateResponse> => {
  
  const res = await fetch(getPostExpensePlanCreateUrl(),
  {      
    ...options,
    method: 'POST',
    headers: { 'Content-Type': 'application/json', ...options?.headers },
    body: JSON.stringify(
      handlerCreateExpensePlanRequest,)
  }
)

  const body = [204, 205, 304].includes(res.status) ? null : await res.text()
  const data: postExpensePlanCreateResponse['data'] = body ? JSON.parse(body) : {}

  return { data, status: res.status, headers: res.headers } as postExpensePlanCreateResponse
}



/**
 * Get all expense-plan items (using cursor-based pagination)
 * @summary Get all expense-plan items
 */
export type postExpensePlanGetResponse200 = {
  data: HandlerGetExpensePlanResponse
  status: 200
}
    
export type postExpensePlanGetResponseComposite = postExpensePlanGetResponse200;
    
export type postExpensePlanGetResponse = postExpensePlanGetResponseComposite & {
  headers: Headers;
}

export const getPostExpensePlanGetUrl = () => {


  

  return `/expense-plan/get`
}

export const postExpensePlanGet = async (handlerGetExpensePlanRequest: HandlerGetExpensePlanRequest, options?: RequestInit): Promise<postExpensePlanGetResponse> => {
  
  const res = await fetch(getPostExpensePlanGetUrl(),
  {      
    ...options,
    method: 'POST',
    headers: { 'Content-Type': 'application/json', ...options?.headers },
    body: JSON.stringify(
      handlerGetExpensePlanRequest,)
  }
)

  const body = [204, 205, 304].includes(res.status) ? null : await res.text()
  const data: postExpensePlanGetResponse['data'] = body ? JSON.parse(body) : {}

  return { data, status: res.status, headers: res.headers } as postExpensePlanGetResponse
}



/**
 * List all expense-plan items (using cursor-based pagination)
 * @summary List all expense-plan items
 */
export type postExpensePlanListResponse200 = {
  data: HandlerListExpensePlanResponse
  status: 200
}
    
export type postExpensePlanListResponseComposite = postExpensePlanListResponse200;
    
export type postExpensePlanListResponse = postExpensePlanListResponseComposite & {
  headers: Headers;
}

export const getPostExpensePlanListUrl = () => {


  

  return `/expense-plan/list`
}

export const postExpensePlanList = async (handlerListExpensePlanRequest: HandlerListExpensePlanRequest, options?: RequestInit): Promise<postExpensePlanListResponse> => {
  
  const res = await fetch(getPostExpensePlanListUrl(),
  {      
    ...options,
    method: 'POST',
    headers: { 'Content-Type': 'application/json', ...options?.headers },
    body: JSON.stringify(
      handlerListExpensePlanRequest,)
  }
)

  const body = [204, 205, 304].includes(res.status) ? null : await res.text()
  const data: postExpensePlanListResponse['data'] = body ? JSON.parse(body) : {}

  return { data, status: res.status, headers: res.headers } as postExpensePlanListResponse
}



